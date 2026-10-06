package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func TestMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Println("TestMiddleware", start)
		next.ServeHTTP(w, r)
	})
}
