package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /route",
		manager.With(
			http.HandlerFunc(handlers.GetProducts),
		),
	)
	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProducts),
	),
	)
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(handlers.CreateProducts),
	),
	)
	mux.Handle("GET /products/{productId}", manager.With(
		http.HandlerFunc(handlers.GetProductById),
	),
	)
	mux.Handle("DELETE /products/{productId}", manager.With(
		http.HandlerFunc(handlers.DeleteProduct),
	),
	)
	mux.Handle("PUT /products/{productId}", manager.With(
		http.HandlerFunc(handlers.UpdateProducts),
	),
	)
}
