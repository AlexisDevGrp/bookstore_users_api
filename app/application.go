package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	r = mux.NewRouter()
)
func StartApplication() {
	mapURLs()
	log.Fatal(http.ListenAndServe(":8091", r))
}