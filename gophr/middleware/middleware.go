package middleware

import (
	"github.com/jayvib/golog"
	"gophr/api/v1/session"
	"gophr/api/v1/user"
	"net/http"
	"net/url"
)

type MiddlewareFunc func(h http.Handler) http.Handler

func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		golog.Infof("%s | %s", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func AuthenticationMiddleware(userRepo user.Service, cache session.Cache) MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return  http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
			usr := session.GetUserFromSession(userRepo, cache, r)
			if usr != nil {
				h.ServeHTTP(w, r)
				return
			}

			query := url.Values{}
			query.Add("next", url.QueryEscape(r.URL.String())	)


			// meaning the session is already expire
			http.Redirect(w, r, "/login?"+query.Encode(), http.StatusFound)
		})
	}
}
