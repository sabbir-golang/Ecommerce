package cmd

import (
	db "ecommerce/Db"
	utils "ecommerce/Utils"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	handlers.Db = db.ConnectDb()
	mux := http.NewServeMux()
	cntrl := func(w http.ResponseWriter, r *http.Request) {
	}
	handler := http.HandlerFunc(cntrl)
	mux.Handle("GET /route", middleware.Logger(handler))
	mux.Handle("GET /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProducts))))
	mux.Handle("POST /products", middleware.Logger(http.HandlerFunc(handlers.CreateProducts)))
	mux.Handle("GET /products/{productId}", middleware.Logger(http.HandlerFunc(handlers.GetProductById)))
	mux.Handle("DELETE /products/{productId}", middleware.Logger(http.HandlerFunc(handlers.DeleteProduct)))
	mux.Handle("PUT /products/{productId}", middleware.Logger(http.HandlerFunc(handlers.UpdateProducts)))
	// mux.HandleFunc("OPTIONS /create-products", http.HandlerFunc(handlers.CreateProducts))
	// mux.HandleFunc("/products/", handlers.DeleteProduct)
	global := utils.GlobalRouter(mux)
	fmt.Println("Listening on port :8080", global)
	err := http.ListenAndServe(":8080", global)
	if err != nil {
		fmt.Println(err)
	}
}
