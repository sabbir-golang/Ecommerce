package product

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(h.GetProducts),
	),
	)
	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(h.CreateProducts),
		h.middlewares.AuthenticationJwt,
	),
	)
	mux.Handle("GET /products/{productId}", manager.With(
		http.HandlerFunc(h.GetProductById),
	),
	)
	mux.Handle("DELETE /products/{productId}", manager.With(
		http.HandlerFunc(h.DeleteProduct),
		// middleware.AuthenticationJwt,
	),
	)
	mux.Handle("PUT /products/{productId}", manager.With(
		http.HandlerFunc(h.UpdateProducts),
		h.middlewares.AuthenticationJwt,
	),
	)
}
