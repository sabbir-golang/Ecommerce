package cmd

import (
	db "ecommerce/Db"
	utils "ecommerce/Utils"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	handlers.DB = db.ConnectDb()
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProducts))
	// mux.HandleFunc("OPTIONS /create-products", http.HandlerFunc(handlers.CreateProducts))
	// mux.HandleFunc("/products/", handlers.DeleteProduct)
	fmt.Println("Listening on port :8080")
	global := utils.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", global)
	if err != nil {
		fmt.Println(err)
	}
}
