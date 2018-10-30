package handler

import (
	"net/http"

	funcv1 "github.com/golang-boilerplate/pkg/functions/v1"
	"github.com/golang-boilerplate/pkg/server/middleware"
	"github.com/julienschmidt/httprouter"
)

//IHandler is an interface that have all func related to route
type IHandler interface {
	HandleHelloWorld(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

//Handler is a struct that will manage the function of the handler each route
type Handler struct {
	Funcs      funcv1.IFuncV1
	Middleware middleware.IMiddleware
}

//InitHandler is a function that will instantiate handler
func InitHandler(fn *funcv1.FuncsV1, mw *middleware.Middleware) *Handler {
	return &Handler{
		Funcs:      fn,
		Middleware: mw,
	}
}
