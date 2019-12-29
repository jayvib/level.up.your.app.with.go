package web

import (
	"fmt"
	"github.com/jayvib/golog"
	"gophr/api/v1"
	"gophr/api/v1/message"
	"gophr/api/v1/session"
	"gophr/view"
	"net/http"
)

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Get the login information from from
	username := r.FormValue("username")
	password := r.FormValue("password")
	next := r.FormValue("next")

	// from the service compare the password
	usr, err := h.svc.GetAndComparePassword(r.Context(), username, password)
	if err != nil {
		if err != v1.ErrorInvalidCredentials {
			msg := message.New("en").AddName("UserNotRegistered").Apply()
			view.RenderTemplate(w, r, h.svc, h.sessionCache, "sessions/login", map[string]interface{}{
				"Error": msg,
				"User": usr,
				"Next": next,
			})
			return
		} else {
			msg := message.New("en").AddName("InvalidCredential").Apply()
			view.RenderTemplate(w, r, h.svc, h.sessionCache, "sessions/login", map[string]interface{}{
				"Error": msg,
				"User": usr,
				"Next": next,
			})
		}
		return
	}

	// create a session
	sess := session.GetSession(h.sessionCache, w, r)
	sess.UserID = usr.ID
	golog.Debugf("session: %#v\n", sess)
	err = h.sessionCache.Set(sess.ID, sess, session.DurationInSecond)
	if err != nil {
		golog.Error(err)
		return
	}

	if next == "" {
		next = "/"
	}

	s,_  := h.sessionCache.Get(sess.ID)
	golog.Debugf("session result: %#v\n", s)

	// redirect to next url
	http.Redirect(w, r, fmt.Sprintf("%s?flash=Signed+in", next), http.StatusFound)
}