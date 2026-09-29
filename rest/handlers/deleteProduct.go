package handlers

import (
	"ecommerce/database"
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
	_, err = Db.Exec(sqlState, id)
	if err != nil {
		http.Error(w, "Product not Found !", http.StatusNotFound)
	}
	database.Delete(id)

}
