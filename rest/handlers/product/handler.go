package product

type Handler struct {
	middlewares *middleware
}

func NewHandler(middlewares *middleware) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
