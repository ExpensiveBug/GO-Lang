package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

// user exist
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templ, err := template.ParseFiles(
			"templates/base.html",
			"templates/register.html",
		)
		if err != nil {
			http.Error(w, "unable to load page", http.StatusInternalServerError)
			return
		}
		//  data = nil
		err = templ.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, "Unable to render page", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		password := r.FormValue("password")
		// validate
		fmt.Println(w, "Email : ", email)
		fmt.Println(w, "Password : ", password)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

// new user
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templ, err := template.ParseFiles(
			"templates/base.html",
			"templates/index.html",
		)
		if err != nil {
			http.Error(w, "Unable to load page", http.StatusInternalServerError)
			return
		}
		// data = nil
		err = templ.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, "Unable to render page!", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		// verify and save
		fmt.Println(w, "Name : ", name)
		fmt.Println(w, "Email : ", email)
		fmt.Println(w, "password : ", password)

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
}
