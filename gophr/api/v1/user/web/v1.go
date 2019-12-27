package web

import (
	"github.com/gorilla/mux"
	"gophr/api/v1/user/repository/file"
	"gophr/api/v1/user/service"
	"net/http"
)

func RegisterHandlers(r *mux.Router) *mux.Router{
	repo := file.New("user.db")
	svc := service.New(repo)
	h :=  New(svc)
	subrouter := r.PathPrefix("/v1").Subrouter()
	subrouter.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {})
	subrouter.HandleFunc("/register", h.CreateUser).Methods(http.MethodPost)
	return subrouter
}
