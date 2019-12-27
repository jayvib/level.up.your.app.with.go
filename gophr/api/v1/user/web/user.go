package web

import (
	"github.com/jayvib/golog"
	"gophr/api/v1/message"
	"gophr/api/v1/user"
	"gophr/model"
	"gophr/view"
	"net/http"
)

func New(svc user.Service) *Handler {
	return &Handler{ svc: svc }
}

type Handler struct {
	svc user.Service
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

	golog.Debugf("%#v\n",usr)
	http.Redirect(w, r, "/?flash=User+created", http.StatusFound)
}
