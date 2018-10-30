package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-boilerplate/pkg/server/handler"
	"github.com/golang-boilerplate/pkg/server/routes"
	"github.com/julienschmidt/httprouter"
)

//Config is a struct that hold server configuration
type Config struct {
	Port string
}

//ServeWeb is a function to serve the route
func ServeWeb(config Config, env string) {
	// go startNonTLSServer()

	r := httprouter.New()

	h := handler.InitHandler(nil, nil)

	router := routes.InitRouter(r, h)
	router.LoadRoute()
	fmt.Printf("Running titip-print.id on %s environment\n", env)
	fmt.Printf("Starting web server at https://localhost%s\n", config.Port)
	srv := &http.Server{
		Addr:         config.Port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// err := srv.ListenAndServeTLS("server.crt", "server.key")
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

//startNonTLSServer to redirect all non tls to tls server
func startNonTLSServer() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Redirecting to https://localhost/")
		http.Redirect(w, r, "https://localhost:8888", http.StatusTemporaryRedirect)
	}))
	http.ListenAndServe(":80", mux)
}
