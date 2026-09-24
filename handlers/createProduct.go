package handlers

import (
	"database/sql"
	"ecommerce/models"
	"encoding/json"
	"fmt"
	"net/http"
)

var DB *sql.DB
var Products []models.Product

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	// HandleCORS(w)
	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(200)
	// 	return
	// }
	// id:=r.Body("id")
	// description := r.Body("description")
	// imageUrl := r.Body("imageUrl")
	// price := r.Body("price")
	// title := r.Body("title")

	var newProduct models.Product
	decode := json.NewDecoder(r.Body)
	decode.Decode(&newProduct)
	// fmt.Println(len(Products))
	newProduct.ID = len(Products) + 1
	Products = append(Products, newProduct)
	w.WriteHeader(201)
	encoder := json.NewEncoder(w)
	encoder.Encode(newProduct)
	sqlstat := "insert into products(id,title, description, price, image_url) values ($1,$2,$3,$4,$5)"
	_, err := DB.Exec(sqlstat, newProduct.ID, newProduct.Title, newProduct.Description, newProduct.Price, newProduct.ImgUrl)
	if err != nil {
		fmt.Println("Database instert Failed", err)
		return
	}
	fmt.Println("insert Done ", newProduct)
}
