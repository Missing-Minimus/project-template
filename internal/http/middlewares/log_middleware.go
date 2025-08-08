package middlewares

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/Missing-Minimus/projects-template/internal/infra/thirdparty/logger"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &responseRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		start := time.Now()
		next.ServeHTTP(recorder, r)
		duration := time.Since(start)

		logger.Info(
			fmt.Sprintf(
				"[REQUEST]\n[Method] %s\n[Path] %s\n[Headers] %v\n[RESPONSE]\n[Status] %d\n[Duration] %s\n[Body] %s",
				r.Method,
				r.URL.Path,
				r.Header,
				recorder.statusCode,
				duration.String(),
				recorder.body.String(),
			),
		)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
	body       bytes.Buffer
}

func (r *responseRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
