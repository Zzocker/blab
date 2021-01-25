package health

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterHandlers registers the handlers that perform healthchecks
func RegisterHandlers(r *mux.Router) {
	r.HandleFunc("/check", check).Methods("GET")
}

func check(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("OK"))
}
