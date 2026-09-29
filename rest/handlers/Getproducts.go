package handlers

import (
	utils "ecommerce/Utils"
	"ecommerce/database"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// HandleCORS(w)
	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(200)
	// 	return
	// }
	product := database.List()
	utils.SendData(w, product, 200)
}
