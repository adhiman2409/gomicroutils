package tracer

import (
	// ...existing imports...

	"net/http"

	"github.com/gorilla/mux"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

// ZipkinZapMiddleware logs trace/span info with zap after each request
func ZipkinMiddleware(tracer *zipkin.Tracer) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			zipkinhttp.NewServerMiddleware(
				tracer,
				zipkinhttp.SpanName(mux.CurrentRoute(r).GetName()),
				zipkinhttp.ServerTags(map[string]string{"component": "rest-api"}),
			)(next)
		})
	}
}
