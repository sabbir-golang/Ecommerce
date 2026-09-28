package middleware

import (
	"log"
	"net/http"
	"time"
)

func Arekta(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Print("Ami Arekta middleware", r.Method, time.Since(start))
	})
}
