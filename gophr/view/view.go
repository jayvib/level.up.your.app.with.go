package view

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jayvib/golog"
	"gophr/api/v1/session"
	"gophr/api/v1/user"
	"html/template"
	"net/http"
)

var AppName = "Gophr"

// both templats/index/home.html and templates/foo/bar/baz.html will match
// the glob, but templates/test.html won't as the ath needs at least one
// directory inside the templates directory
var templates = template.Must(template.ParseGlob("./templates/**/*.html"))

var layoutFuncs = template.FuncMap{
	"yield": func() (string, error) {
		return "", fmt.Errorf("yield calle inappropriately")
	},
}

var layout = template.Must(template.New("layout.html").
	Funcs(layoutFuncs).
	ParseFiles("templates/layout.html"))

// errorTemplate will be the default mark up
// when an error occur during executing a template.
var errorTemplate = `
<html>
	<body>
		<h1>Error rendering template %s</h1>
		<p>%s</p>
	</body>
</html>`

func RegisterHandlers(r *mux.Router, userService user.Service, cache session.Cache) *mux.Router {
	subrouter := r.PathPrefix("/").Subrouter()
	subrouter.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	subrouter.HandleFunc("/", HomeViewHandler(userService, cache)).Methods(http.MethodGet)
	subrouter.HandleFunc("/register", UserNewViewHandler(userService, cache)).Methods(http.MethodGet)
	return subrouter
}

func UserNewViewHandler(svc user.Service, cache session.Cache) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		RenderTemplate(w, r, svc, cache,"users/new", nil)
	}
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, userService user.Service, sessionCache session.Cache,name string, data map[string]interface{}) {
	golog.Debug("Query:", r.URL.Query())
	if data == nil {
		data = make(map[string]interface{})
	}

	data["CurrentUser"] = session.GetUserFromSession(userService, sessionCache, r)
	data["Flash"] = r.URL.Query().Get("flash")

	// create a custom func map
	funcs := template.FuncMap{
		"yield": func() (template.HTML, error) {
			var buf bytes.Buffer
			err := templates.ExecuteTemplate(&buf, name, data)
			return template.HTML(buf.String()), err
		},
	}

	// clone the main layout
	layoutClone, _ := layout.Clone() // to avoid race condition

	// attach the custom funcs
	layoutClone.Funcs(funcs)

	// execute the layout
	err := layoutClone.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf(errorTemplate, name, err), http.StatusInternalServerError)
	}
}


func HomeViewHandler(service user.Service, sessionCache session.Cache) http.HandlerFunc{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		RenderTemplate(w, r, service, sessionCache, "index/home", map[string]interface{}{"title": AppName})
	})
}

