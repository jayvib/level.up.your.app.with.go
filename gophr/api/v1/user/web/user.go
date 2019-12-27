package web

import (
	"github.com/jayvib/golog"
	"gophr/api/v1/message"
	"gophr/api/v1/session"
	"gophr/api/v1/user"
	"gophr/model"
	"gophr/view"
	"net/http"
	"time"
)

func New(svc user.Service, cache session.Cache) *Handler {
	return &Handler{
		svc: svc,
		sessionCache: cache,
	}
}

type Handler struct {
	svc user.Service
	sessionCache session.Cache
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	usr, err := model.NewUser(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"))
	if err != nil {
		view.RenderTemplate(w, r, "users/new", map[string]interface{}{
			"Error": err.Error(),
			"User": usr,
		})
		return
	}

	// Do business logic here
	err = h.svc.Save(r.Context(), usr)
	if err != nil {

		msg := message.New("en")
		switch err {
		case user.ErrorUsernameExists:
			msg.AddName("UsernameExists")
		case user.ErrorEmailExists:
			msg.AddName("EmailExists")
		}


		view.RenderTemplate(w, r, "users/new", map[string]interface{}{
			"Error": msg.Apply(),
			"User": usr,
		})
		return
	}

	session := model.NewSession(w)
	session.UserID = usr.ID

	err = h.sessionCache.Set(session.ID, session, model.SessionDuration*time.Second)
	if err != nil {
		view.RenderTemplate(w, r, "users/new", map[string]interface{}{
			"Error": err.Error(),
			"User": usr,
		})
		return
	}

	golog.Debugf("%#v\n",usr)
	http.Redirect(w, r, "/?flash=User+created", http.StatusFound)
}
