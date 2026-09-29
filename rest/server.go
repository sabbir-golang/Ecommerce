package rest

import (
	"ecommerce/config"
	"ecommerce/db"
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	handlers.Db = db.ConnectDb()
	manager := middleware.NewManager()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)
	mux := http.NewServeMux()
	wrapedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(int(cnf.HttpPort))
	fmt.Println("Listening on port ", addr)
	err := http.ListenAndServe(addr, wrapedMux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
