package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/adhiman2409/gomicroutils/errs"
	"github.com/adhiman2409/gomicroutils/grpcclient"
	"github.com/adhiman2409/gomicroutils/logger"
	"github.com/rs/xid"
	"go.uber.org/zap"
)

type contextKey string

const (
	correlationIDCtxKey contextKey = "correlation_id"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	errMsg     string
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK, ""}
}

func (lrw *loggingResponseWriter) Header() http.Header {
	return lrw.ResponseWriter.Header()
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(p []byte) (int, error) {
	if lrw.statusCode >= http.StatusBadRequest {
		aerr := errs.RestError{}
		if err := json.Unmarshal(p, &aerr); err == nil {
			lrw.errMsg = aerr.Message
		}
	}
	return lrw.ResponseWriter.Write(p)
}

func getCurrentISTTime() time.Time {
	return time.Now().Local().Add(time.Hour*time.Duration(5) + time.Minute*time.Duration(30))
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := logger.Get()

		correlationID := xid.New().String()

		ctx := context.WithValue(
			r.Context(),
			correlationIDCtxKey,
			correlationID,
		)

		r = r.WithContext(ctx)

		l = l.With(zap.String(string(correlationIDCtxKey), correlationID))

		w.Header().Add("X-Correlation-ID", correlationID)

		lrw := newLoggingResponseWriter(w)

		r = r.WithContext(logger.WithCtx(ctx, l))
		authInfo := grpcclient.GetAuthInfo(r)
		defer func(start time.Time, domain string) {
			if domain == "" {
				return
			}
			if lrw.statusCode >= 400 {
				l.Error(
					fmt.Sprintf(
						"%s request to %s failed",
						r.Method,
						r.RequestURI,
					), domain,
					zap.String("method", r.Method),
					zap.String("url", r.RequestURI),
					zap.String("user_agent", r.UserAgent()),
					zap.Int("status_code", lrw.statusCode),
					zap.String("err_msg", lrw.errMsg),
					zap.String("IST", getCurrentISTTime().Format("2006-01-02 15:04:05")),
					zap.Duration("elapsed_ms", time.Since(start)),
				)
				return
			} else if time.Since(start).Seconds() >= 1.0 {
				if strings.Contains(r.RequestURI, "/download") || strings.Contains(r.RequestURI, "/upload") {
					return
				}
				l.Warn(
					fmt.Sprintf(
						"%s request to %s completed",
						r.Method,
						r.RequestURI,
					), domain,
					zap.String("method", r.Method),
					zap.String("url", r.RequestURI),
					zap.String("user_agent", r.UserAgent()),
					zap.Int("status_code", lrw.statusCode),
					zap.String("err_msg", lrw.errMsg),
					zap.String("IST", getCurrentISTTime().Format("2006-01-02 15:04:05")),
					zap.Duration("elapsed_ms", time.Since(start)),
				)
				return
			}
		}(time.Now(), authInfo.Domain)

		next.ServeHTTP(lrw, r)
	})
}
