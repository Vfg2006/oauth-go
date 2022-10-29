package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	WithRoutes = func(routes ...Route) ConfigRouter {
		return func(router *Router) {
			router.AddRoutes(routes...)
		}
	}
)

type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}

type Router struct {
	router *httprouter.Router
}

type ConfigRouter func(router *Router)

func New(configs ...ConfigRouter) Router {
	router := &Router{
		router: httprouter.New(),
	}

	for _, config := range configs {
		config(router)
	}

	return *router
}

func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}

func (r Router) AddRoutes(routes ...Route) {
	for _, route := range routes {
		r.router.Handler(route.Method, route.Path, route.Handler)
	}
}
