package handlers

import (
	"BLogPost/database"
	"BLogPost/repositories"
	"html/template"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Println("Database connection error:", err)
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	posts, err := repositories.GetPosts(db)
	if err != nil {
		http.Error(w, "Unable to get posts: "+err.Error(), http.StatusInternalServerError)
		return
	}
	templ, err := template.ParseFiles(
		"templates/base.html",
		"templates/index.html",
	)
	if err != nil {
		http.Error(w, "Unable to Load Home Page", http.StatusInternalServerError)
		return
	}
	err = templ.ExecuteTemplate(w, "base", posts)
	if err != nil {
		http.Error(w, "unable to render Home page", http.StatusInternalServerError)
		return
	}
}
