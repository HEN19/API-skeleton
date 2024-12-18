package routes

import (
	"github.com/api-skeleton/endpoint"
	"github.com/gorilla/mux"
)

func Controller() *mux.Router {
	routes := mux.NewRouter()

	user := routes.PathPrefix("/user").Subrouter()
	user.HandleFunc("/register", endpoint.RegistrationEndpoint).Methods("POST", "OPTIONS")
	user.HandleFunc("/login", endpoint.LoginEndpoint).Methods("POST", "OPTIONS")

	return routes
}
