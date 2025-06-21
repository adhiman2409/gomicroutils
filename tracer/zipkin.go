package tracer

import (
	// ...existing imports...

	"net/http"

	"github.com/gorilla/mux"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

// ZipkinMiddleware injects Zipkin tracing into each HTTP request
func ZipkinMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			route := mux.CurrentRoute(r)
			spanName := "unknown"
			if route != nil {
				if name := route.GetName(); name != "" {
					spanName = name
				}
			}
			host := r.Host
			if host == "" {
				host = "unknown"
			}
			spanName = host + "-" + spanName
			middleware := zipkinhttp.NewServerMiddleware(
				GetTracer().tracer,
				zipkinhttp.SpanName(spanName),
				zipkinhttp.ServerTags(map[string]string{"component": "rest-api"}),
			)
			middleware(next).ServeHTTP(w, r)
		})
	}
}
