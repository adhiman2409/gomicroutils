package tracer

import (
	"context"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const endpointURL = "https://zipkin.unirms.com/api/v2/spans"

const FUNCTION_SKIP_LEVEL = 1

var once sync.Once

var T *zipkin.Tracer

func GetTracer() *zipkin.Tracer {
	once.Do(func() {
		// The reporter sends traces to zipkin server
		localEndPointURL := os.Getenv("ZIPKIN_LOCAL_ENDPOINT_URL")
		if localEndPointURL == "" {
			localEndPointURL = endpointURL
		}
		reporter := reporterhttp.NewReporter(localEndPointURL)

		// Local endpoint represent the local service information
		serviceName := os.Getenv("ZIPKIN_SERVICE_NAME")
		if serviceName == "" {
			serviceName = "utils-srv"
		}
		servicePort := os.Getenv("ZIPKIN_SERVICE_PORT")
		if servicePort == "" {
			servicePort = "3000"
		}
		servicePortInt, err := strconv.Atoi(servicePort)
		if err != nil {
			return
		}
		localEndpoint := &model.Endpoint{ServiceName: serviceName, Port: uint16(servicePortInt)}

		// Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00) of traces.
		sampler, err := zipkin.NewCountingSampler(1)
		if err != nil {
			return
		}

		t, err := zipkin.NewTracer(
			reporter,
			zipkin.WithSampler(sampler),
			zipkin.WithLocalEndpoint(localEndpoint),
			zipkin.WithSharedSpans(true), // Enable shared spans for better trace visibility
		)
		if err != nil {
			// Log the error and panic to avoid leaving tracer nil
			panic("failed to initialize zipkin tracer: " + err.Error())
		}

		T = t
	})

	return T
}

func SCSP(ctx context.Context, prefix string) (zipkin.Span, context.Context) {
	functionName := "Unknown"
	pc, _, _, ok := runtime.Caller(FUNCTION_SKIP_LEVEL)
	if ok {
		function := runtime.FuncForPC(pc).Name()
		functok := strings.Split(function, ".")
		functionName = functok[len(functok)-1]
	}

	parentSpan := zipkin.SpanFromContext(ctx)
	if parentSpan == nil {
		// No span found: middleware may not be configured correctly
		if functionName != "Verify" {
			log.Println("No Zipkin span found in context", functionName)
		}
		span, newCtx := T.StartSpanFromContext(ctx, prefix+"-"+functionName)
		return span, newCtx
	}

	childSpan := T.StartSpan(prefix+"-"+functionName, zipkin.Parent(parentSpan.Context()))
	newCtx := zipkin.NewContext(ctx, childSpan)
	return childSpan, newCtx
}
