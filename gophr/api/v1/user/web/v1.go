package web

import (
	"github.com/gorilla/mux"
	"gophr/api/v1/session"
	"gophr/api/v1/user"
	"gophr/middleware"
	"net/http"
)

func RegisterHandlers(r *mux.Router, svc user.Service, cache session.Cache) *mux.Router{
	h :=  New(svc, cache)
	subrouter := r.PathPrefix("/v1").Subrouter()
	subrouter.HandleFunc("/test", func(writer http.ResponseWriter, request *http.Request) {})
	subrouter.HandleFunc("/register", h.CreateUser).Methods(http.MethodPost)

	// with authentication middleware
	securedSubrouter := r.PathPrefix("/v1").Subrouter()
	securedSubrouter.Use(mux.MiddlewareFunc(middleware.AuthenticationMiddleware(svc, cache)))
	return subrouter
}
