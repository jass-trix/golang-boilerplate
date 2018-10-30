package routes

import (
	"github.com/golang-boilerplate/pkg/server/handler"
	"github.com/julienschmidt/httprouter"
)

//IRoute is a function that provide all routing
type IRoute interface {
}

//Router is a struct that will manage all the apps routing
type Router struct {
	Router  *httprouter.Router
	Handler handler.IHandler
}

//InitRouter is a function that will instantiate route
func InitRouter(r *httprouter.Router, h *handler.Handler) *Router {
	return &Router{
		Router:  r,
		Handler: h,
	}
}
