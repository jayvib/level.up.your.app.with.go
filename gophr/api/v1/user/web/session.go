package web

import (
	"gophr/api/v1/session"
	"gophr/view"
	"net/http"
)

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	// Get the session from the request
	sess, _ := session.GetRequestSession(h.sessionCache, r)
	if sess != nil {
		// delete the session from the cache
		err := h.sessionCache.Delete(sess.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// redirect to sign out page
	view.RenderTemplate(w, r, h.svc, h.sessionCache,"sessions/destroy", nil)
}

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	// TODO: Please implement me

	// Get the login information from from

	// from the service compare the password

	// create a session

	// redirect to next url

}