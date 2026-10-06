package handlers

import (
	utils "ecommerce/Utils"
	"ecommerce/database"
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
	product := database.Get(id)
	utils.SendData(w, product, 200)
	fmt.Println(product)

}
