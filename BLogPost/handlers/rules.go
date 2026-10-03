package handlers

import (
	"html/template"
	"net/http"
)

func RulesHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles(
		"templates/base.html",
		"templates/rules.html",
	)
	if err != nil {
		http.Error(w, "Unable to load template", http.StatusInternalServerError)
		return
	}
	err = templ.ExecuteTemplate(w, "base", nil)
	if err != nil {
		http.Error(w, "Unable to render template", http.StatusInternalServerError)
		return
	}

}
