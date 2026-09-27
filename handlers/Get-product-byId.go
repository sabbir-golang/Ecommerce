package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	// fmt.Println("METHOD:", r.Method)
	fmt.Println("URL PATH:", r.URL.Path)

	id, err := strconv.Atoi(productId)
	if err != nil {
		fmt.Println("Invalid ID number")
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(Products[id-1])
	fmt.Println(Products[id-1])

}
