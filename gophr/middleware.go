package main

import (
	"github.com/jayvib/golog"
	"net/http"
)

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		golog.Infof("%s | %s", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func AuthenticateMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.WriteHeader(http.StatusUnauthorized)
		RenderTemplate(w, r, "others/unauthorized", nil)
	})
}
