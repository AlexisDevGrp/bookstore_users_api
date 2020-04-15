package app

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	Client *sql.DB
	r = mux.NewRouter()
)
func StartApplication() {
	mapURLs()
	log.Fatal(http.ListenAndServe(":8091", r))
}