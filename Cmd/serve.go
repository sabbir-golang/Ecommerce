package cmd

import (
	db "ecommerce/Db"
	utils "ecommerce/Utils"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	handlers.Db = db.ConnectDb()
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProducts))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handlers.GetProductById))
	mux.Handle("DELETE /products/{productId}", http.HandlerFunc(handlers.DeleteProduct))
	mux.Handle("PUT /products/{productId}", http.HandlerFunc(handlers.UpdateProducts))
	// mux.HandleFunc("OPTIONS /create-products", http.HandlerFunc(handlers.CreateProducts))
	// mux.HandleFunc("/products/", handlers.DeleteProduct)
	global := utils.GlobalRouter(mux)
	fmt.Println("Listening on port :8080", global)
	err := http.ListenAndServe(":8080", global)
	if err != nil {
		fmt.Println(err)
	}
}
