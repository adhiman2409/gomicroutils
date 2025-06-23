package tracer

import (
	// ...existing imports...

	"net/http"

	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

// ZipkinMiddleware injects Zipkin tracing into each HTTP request
func ZipkinMiddleware() func(http.Handler) http.Handler {
	return zipkinhttp.NewServerMiddleware(T)
}
