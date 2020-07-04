package router

import (
	"github.com/bubeha/go-test-rest-api/app/controller"
	"github.com/gorilla/mux"
)

// IRouter ...
type IRouter interface {
	InitRoutes()
	RouteMultiplexer() *mux.Router
}

type router struct {
	mux *mux.Router
}

func (r *router) RouteMultiplexer() *mux.Router {
	return r.mux
}

func (r *router) InitRoutes() {
	r.mux.HandleFunc("/", controller.Greetings).Methods("GET")
}

// New ...
func New() IRouter {
	return &router{
		mux: mux.NewRouter(),
	}
}
