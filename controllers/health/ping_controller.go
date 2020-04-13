package health

import (
	"io"
	"net/http"
)
// This function will be used by Cloud Provider to ensure our webserver is up and running
// Add all functions required to control the health of your microservice (database, Redis,..)
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `{"alive": true}`)
}
