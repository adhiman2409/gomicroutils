package tracer

import (
	"os"
	"strconv"

	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"
	reporterhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

const endpointURL = "/api/v2/spans"

func NewTracer(domain string) (*zipkin.Tracer, error) {
	// The reporter sends traces to zipkin server
	localEndPointURL := os.Getenv("ZIPKIN_LOCAL_ENDPOINT_URL")
	if localEndPointURL == "" {
		localEndPointURL = "https://" + domain + endpointURL
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
		return nil, err
	}
	localEndpoint := &model.Endpoint{ServiceName: serviceName, Port: uint16(servicePortInt)}

	// Sampler tells you which traces are going to be sampled or not. In this case we will record 100% (1.00) of traces.
	sampler, err := zipkin.NewCountingSampler(1)
	if err != nil {
		return nil, err
	}

	t, err := zipkin.NewTracer(
		reporter,
		zipkin.WithSampler(sampler),
		zipkin.WithLocalEndpoint(localEndpoint),
		zipkin.WithSharedSpans(true), // Enable shared spans for better trace visibility
	)
	if err != nil {
		return nil, err
	}

	return t, err
}
