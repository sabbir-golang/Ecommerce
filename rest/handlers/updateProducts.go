package handlers

import (
	"ecommerce/database"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("productId")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid id path", 404)
		return
	}

	var UpdateProduct database.Product
	decode := json.NewDecoder(r.Body)
	decode.Decode(&UpdateProduct)
	UpdateProduct.ID = id
	sqlState := `update products set title=$1, description=$2, image_url=$3, price=$4 where id=$5`
	Db.Exec(sqlState, UpdateProduct.Title, UpdateProduct.Description, UpdateProduct.ImgUrl, UpdateProduct.Price, id)
	database.Update(UpdateProduct)

}
