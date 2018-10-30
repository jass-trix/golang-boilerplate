package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//HandleHelloWorld is a function that hold the hello world route
func (h Handler) HandleHelloWorld(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Write([]byte("Hello"))
}
