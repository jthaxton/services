package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"main/handlers"
	// "io/ioutil"
	// "github.com/smahjoub/events-api/handlers"
	// "github.com/smahjoub/events-api/store"
)

// Args args used to run the server
type Args struct {
	// postgres connection string, of the form,
	// e.g "postgres://user:password@localhost:5432/database?sslmode=disable"
	conn string
	// port for the server of the form,
	// e.g ":8080"
	port string
}

func Run(args Args) error {
	// router
	router := mux.NewRouter().
		PathPrefix("/api/v1/"). // add prefix for v1 api `/api/v1/`
		Subrouter()

	RegisterAllRoutes(router)

	// start server
	log.Println("Starting server at port: ", args.port)
	return http.ListenAndServe(args.port, router)
}

// RegisterAllRoutes registers all routes of the api
func RegisterAllRoutes(router *mux.Router) {

	// set content type
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	
	router.HandleFunc("/test", handlers.GetOther).Methods(http.MethodGet)
	router.HandleFunc("/users/get", handlers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/users/create", handlers.CreateUser).Methods(http.MethodGet)
	
}

