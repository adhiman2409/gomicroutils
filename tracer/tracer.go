package tracer

import (
	"context"
	"fmt"
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

var tracer *zipkin.Tracer

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
		fmt.Println("Using service name:", serviceName, "and port:", servicePortInt, "localEndPointURL:", localEndPointURL)
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
			return
		}

		tracer = t
	})

	return tracer
}

func SSFC(ctx context.Context) (zipkin.Span, context.Context) {

	functionName := "Unknown"
	pc, _, _, ok := runtime.Caller(FUNCTION_SKIP_LEVEL)
	if ok {
		function := runtime.FuncForPC(pc).Name()
		functok := strings.Split(function, ".")
		functionName = functok[len(functok)-1]
	}

	span, ctx := GetTracer().StartSpanFromContext(ctx, functionName)
	return span, ctx
}

func STag(span zipkin.Span, key string, value string) {
	if span == nil {
		return
	}
	span.Tag(key, value)
}

func SError(span zipkin.Span, err string) {
	if span == nil || err == "" {
		return
	}
	span.Tag("error", err)
}
