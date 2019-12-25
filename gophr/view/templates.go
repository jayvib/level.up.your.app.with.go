package view

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
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

func RegisterHandlers(r *mux.Router) *mux.Router {
	subrouter := r.PathPrefix("/").Subrouter()
	subrouter.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	subrouter.HandleFunc("/", HomeViewHandler).Methods(http.MethodGet)
	subrouter.HandleFunc("/register", UserNewViewHandler).Methods(http.MethodGet)
	return subrouter
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {

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

func HomeViewHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, "index/home", map[string]string{"title": AppName})
}

