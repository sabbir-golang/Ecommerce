package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	utils "ecommerce/Utils"
	"ecommerce/config"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func AuthenticationJwt(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hi iam JWT")
		header := r.Header.Get("Authorization")
		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			fmt.Println("Unauthorized")
			http.Error(w, "Unauthorized header!", http.StatusUnauthorized)
			return
		}
		accessToken := headerArr[1]
		tokenParts := strings.Split(accessToken, ".")
		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSigniture := tokenParts[2]
		byteMessage := []byte(jwtHeader + "." + jwtPayload)
		byteSecretKey := []byte(config.GetConfig().SecretKey)
		h := hmac.New(sha256.New, byteSecretKey)
		h.Write(byteMessage)
		hash := h.Sum(nil)
		newSigniture := utils.Base64Encode(hash)
		if newSigniture != jwtSigniture {
			http.Error(w, "Unauthorized signiture!", http.StatusUnauthorized)
			return
		}
		log.Print("Authentication middleware")
		next.ServeHTTP(w, r)
	})
}
