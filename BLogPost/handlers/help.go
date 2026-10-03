package handlers

import (
	"html/template"
	"net/http"
)

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles(
		"templates/base.html",
		"templates/help.html",
	)
	if err != nil {
		http.Error(w, "Unable to load page", http.StatusInternalServerError)
		return
	}
	err = templ.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
		return
	}
}
