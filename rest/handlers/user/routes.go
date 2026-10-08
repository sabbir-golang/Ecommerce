package user

import (
	"ecommerce/rest/middleware"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /create-user",
		manager.With(
			http.HandlerFunc(h.CreateUser),
		),
	)
	mux.Handle("POST /login", manager.With(http.HandlerFunc(h.LoginUser)))
}
