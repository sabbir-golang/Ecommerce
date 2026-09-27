package middleware

import (
	"log"
	"net/http"
)

func Hudai(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Ami Hudai aktaa middleware")
		next.ServeHTTP(w, r)
	})
}
