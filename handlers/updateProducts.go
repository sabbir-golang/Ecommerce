package handlers

import (
	"ecommerce/models"
	"encoding/json"
	"fmt"
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

	var UpdateProduct models.Product
	decode := json.NewDecoder(r.Body)
	decode.Decode(&UpdateProduct)
	sqlState := `update products set title=$1, description=$2, image_url=$3, price=$4 where id=$5`
	Db.Exec(sqlState, UpdateProduct.Title, UpdateProduct.Description, UpdateProduct.ImgUrl, UpdateProduct.Price, id)
	for i, product := range Products {
		if product.ID == id {
			Products[i].Title = UpdateProduct.Title
			Products[i].Description = UpdateProduct.Description
			Products[i].ImgUrl = UpdateProduct.ImgUrl
			Products[i].Price = UpdateProduct.Price
			fmt.Println("Product update successfully")
			break
		}
	}

}
