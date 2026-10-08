package cmd

import (
	"ecommerce/config"
	"ecommerce/rest"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/reviews"
	"ecommerce/rest/handlers/user"
	"ecommerce/rest/middleware"
)

func Serve() {
	cnf := config.GetConfig()
	mid := middleware.NewMiddlewares(cnf)
	server := rest.NewServer(product.NewHandler(mid), user.NewHandler(), reviews.NewHandler())
	server.Start(cnf)
}
