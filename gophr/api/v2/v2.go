package v2

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHandlers(r *mux.Router) *mux.Router{
	subrouter := r.PathPrefix("/v2").Subrouter()
	subrouter.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = fmt.Fprintf(writer, "api version 2")
	})
	return subrouter
}
