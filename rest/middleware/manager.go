package middleware

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	mngr := Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
	return &mngr
}
func (mngr *Manager) Use(middleware ...Middleware) {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	for _, middleware := range middlewares {
		n = middleware(n)
	}
	return n
}

//	for _, globalmiddleware := range mngr.globalMiddlewares {
//		n = globalmiddleware(n)
//	}
//
//	for i := len(middlewares) - 1; i >= 0; i-- {
//		middleware := middlewares[i]
//		n = middleware(n)
//	}
//
// return n
func (mngr *Manager) WrapMux(handler http.Handler) http.Handler {
	h := handler

	for _, middleware := range mngr.globalMiddlewares {
		h = middleware(h)
	}
	return h
}
