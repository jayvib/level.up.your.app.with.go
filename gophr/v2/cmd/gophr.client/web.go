package main

import (
	"github.com/jayvib/golog"
	"net/http"
)


func HomePage(w http.ResponseWriter, r *http.Request) {

	layout, _ := layoutTemplate.Clone()
	err := layout.Execute(w, nil)
	if err != nil {
		golog.Error(err)
		return
	}
}

