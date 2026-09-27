package handlers

import (
	"fmt"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// HandleCORS(w)
	fmt.Println("METHOD: DELETE")

	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	sqlState := `delete from products where id= $1`
	Db.Exec(sqlState, id)
	fmt.Println(Products[id-1])
	for i, product := range Products {
		if product.ID == id {
			Products = append(Products[:i], Products[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Product %d deleted successfully", id)

			break
		}
	}
	http.Error(w, "Product not Found !", http.StatusNotFound)

	// if r.Method != "POST" {
	// 	return
	// }
}
