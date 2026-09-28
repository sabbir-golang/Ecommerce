package cmd

import (
	db "ecommerce/Db"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	handlers.Db = db.ConnectDb()
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.Cors, middleware.Preflight, middleware.Logger)
	initRoutes(mux, manager)
	// global := middleware.CorsWithPrefight(mux)
	fmt.Println("Listening on port :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err)
	}
}
