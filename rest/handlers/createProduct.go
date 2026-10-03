package handlers

import (
	"database/sql"
	utils "ecommerce/Utils"
	"ecommerce/database"

	"encoding/json"
	"fmt"
	"net/http"
)

var Db *sql.DB
var err error

func CreateProducts(w http.ResponseWriter, r *http.Request) {
	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(200)
	// 	return
	// }
	// id:=r.Body("id")
	// description := r.Body("description")
	// imageUrl := r.Body("imageUrl")
	// price := r.Body("price")
	// title := r.Body("title")

	var newProduct database.Product
	decode := json.NewDecoder(r.Body)
	decode.Decode(&newProduct)
	sqlstat := `insert into products(title, description, price, image_url) values ($1,$2,$3,$4) RETURNING id`
	err := Db.QueryRow(sqlstat, newProduct.Title, newProduct.Description, newProduct.Price, newProduct.ImgUrl).Scan(&newProduct.ID)
	if err != nil {
		fmt.Println("Database error ", err)
		utils.SendError(w, 404, "Database instert Failed")
		return
	}
	database.Store(newProduct)
	utils.SendData(w, newProduct, 201)
	fmt.Println("insert Done ", newProduct)
}
