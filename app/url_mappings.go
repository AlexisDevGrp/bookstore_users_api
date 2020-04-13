package app

import (
	"github.com/AlexisDevGrp/bookstore_users_api/controllers/health"
	"github.com/AlexisDevGrp/bookstore_users_api/controllers/users"
)
// RegisterRoutes - Controller url mapping provides information on the authorised catalog of routes
func mapURLs() {
	r.HandleFunc("/ping", health.Ping)
	r.HandleFunc("/users/{user_id:[0-9]+}", users.GetUser).Methods("GET")
	r.HandleFunc("/users", users.CreateUser).Methods("POST")
}
