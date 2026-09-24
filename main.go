package main

import (
	db "ecommerce/Db"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

// func helloHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello HandleFunc")
// }

//	func AboutHandler(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "I am a software engineer")
//	}

func main() {
	handlers.DB = db.ConnectDb()
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /create-products", http.HandlerFunc(handlers.CreateProducts))
	// mux.HandleFunc("OPTIONS /create-products", http.HandlerFunc(handlers.CreateProducts))
	// mux.HandleFunc("/products/", handlers.DeleteProduct)
	fmt.Println("Listening on port :8080")
	global := handlers.GlobalRouter(mux)
	err := http.ListenAndServe(":8080", global)
	if err != nil {
		fmt.Println(err)
	}
}
