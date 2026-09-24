package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// HandleCORS(w)
	fmt.Println("METHOD:", r.Method)
	fmt.Println("URL PATH:", r.URL.Path)
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
		return
	}
	if r.Method != http.MethodDelete {
		return
	}
	idString := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	fmt.Println(id)
	// if r.Method != "POST" {
	// 	return
	// }
}
