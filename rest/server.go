package rest

import (
	"ecommerce/config"
	"ecommerce/db"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/reviews"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
	reviewsHandler *reviews.Handler
}

func NewServer(productHandler *product.Handler,
	userHandler *user.Handler,
	reviewsHandler *reviews.Handler,
) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
		reviewsHandler: reviewsHandler,
	}
}

func (server *Server) Start(cnf config.Config) {
	product.Db = db.ConnectDb()
	manager := middleware.NewManager()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)
	mux := http.NewServeMux()
	wrapedMux := manager.WrapMux(mux)
	server.productHandler.RegRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)
	server.reviewsHandler.RegReviews(mux, *manager)

	addr := ":" + strconv.Itoa(int(cnf.HttpPort))
	fmt.Println("Listening on port ", addr)
	err := http.ListenAndServe(addr, wrapedMux)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
