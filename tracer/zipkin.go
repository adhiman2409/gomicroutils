package tracer

import (
	// ...existing imports...

	"net/http"

	"github.com/gorilla/mux"
	zipkinhttp "github.com/openzipkin/zipkin-go/middleware/http"
)

// ZipkinZapMiddleware logs trace/span info with zap after each request
func ZipkinMiddleware(next http.Handler) http.Handler {
	t, _ := NewTracer("vms.unirms.com")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zipkinhttp.NewServerMiddleware(
			t,
			zipkinhttp.SpanName(mux.CurrentRoute(r).GetName()),
			zipkinhttp.ServerTags(map[string]string{"component": "rest-api"}),
		)(next)
	})
}
