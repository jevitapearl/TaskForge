package middleware

import (
	"log"
	"net/http"
	"time"

	"github.com/jevitapearl/TaskForge/internal/contextx"
)

type StatusRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func (sr *StatusRecorder) WriteHeader(statuscode int) {
	sr.StatusCode = statuscode
	sr.ResponseWriter.WriteHeader(sr.StatusCode)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		recorder := &StatusRecorder{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}

		start := time.Now()
		next.ServeHTTP(recorder, r)

		userID := contextx.UserID(r.Context())
		if userID == "" {
			userID = "anonymous"
		}

		log.Printf("user_id=%s | method=%s | path=%s | status=%v | time=%v", userID, r.Method, r.URL.Path, recorder.StatusCode, time.Since(start))
	})
}
