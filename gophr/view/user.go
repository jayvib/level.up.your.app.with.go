package view

import "net/http"

func UserNewViewHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, r, "users/new", nil)
}