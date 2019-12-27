package v1

import (
	"github.com/jayvib/golog"
	"gophr/model"
	"gophr/view"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	usr, err := model.NewUser(r.FormValue("username"), r.FormValue("email"), r.FormValue("password"))
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
