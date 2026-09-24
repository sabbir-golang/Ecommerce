package handlers

import (
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// HandleCORS(w)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(Products)
}
