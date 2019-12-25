package v1

import (
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHandlers(r *mux.Router) *mux.Router{
	subrouter := r.PathPrefix("/v1").Subrouter()
	subrouter.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {})
	return subrouter
}
