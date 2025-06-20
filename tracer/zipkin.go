package tracer

import (
	// ...existing imports...

	"net/http"

	"github.com/gorilla/mux"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

// ZipkinZapMiddleware logs trace/span info with zap after each request
func ZipkinMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, _ := NewTracer("vms.unirms.com")
		zipkinhttp.NewServerMiddleware(
			t,
			zipkinhttp.SpanName(mux.CurrentRoute(r).GetName()),
			zipkinhttp.ServerTags(map[string]string{"component": "rest-api"}),
		)
		next.ServeHTTP(w, r)
	})
}
