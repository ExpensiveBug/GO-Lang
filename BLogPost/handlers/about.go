package handlers

import (
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles(
		"templates/base.html",
		"templates/about.html",
	)
	if err != nil {
		http.Error(w, "Unable to load About page !", http.StatusInternalServerError)
		return
	}
	err = templ.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Unable to render page !", http.StatusInternalServerError)
		return
	}

}
