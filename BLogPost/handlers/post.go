package handlers

import (
	"BLogPost/database"
	"BLogPost/models"
	"BLogPost/repositories"
	"html/template"
	"log"
	"net/http"
	"uuid"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	if idString == "" {
		http.Error(w, "Post ID is missing", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	post, err := repositories.GetPostByID(db, id)
	if err != nil {
		http.Error(w, "Unable to get post: "+err.Error(), http.StatusInternalServerError)
		return
	}
	templ, err := template.ParseFiles(
		"templates/base.html",
		"templates/post.html",
	)
	if err != nil {
		http.Error(w, "unable to load page", http.StatusInternalServerError)
		return
	}
	err = templ.ExecuteTemplate(w, "base", post)
	if err != nil {
		http.Error(w, "Unable to render page", http.StatusInternalServerError)
		return
	}
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		templ, err := template.ParseFiles(
			"templates/base.html",
			"templates/create_post.html",
		)
		if err != nil {
			http.Error(w, "unable to load page", http.StatusInternalServerError)
			return
		}

		err = templ.ExecuteTemplate(w, "base", nil)
		if err != nil {
			http.Error(w, "Unable to render page", http.StatusInternalServerError)
			return
		}
		return
	}
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		category := r.FormValue("category")
		author := r.FormValue("author")
		content := r.FormValue("content")

		post := models.Post{
			Title:    title,
			Category: category,
			Author:   author,
			Content:  content,
		}
		db, err := database.ConnectDB()
		if err != nil {
			log.Fatal(err)
		}
		repo := repositories.PostRepository{
			DB: db,
		}
		err = repo.CreatePost(post)
		if err != nil {
			http.Error(w, "Failed to create post", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")

	if idString == "" {
		http.Error(w, "Post ID is missing", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	if r.Method == http.MethodGet {
		post, err := repositories.GetPostByID(db, id)
		if err != nil {
			http.Error(w, "Unable to get post: "+err.Error(), http.StatusInternalServerError)
			return
		}
		templ, err := template.ParseFiles(
			"templates/base.html",
			"templates/edit_post.html",
		)
		if err != nil {
			http.Error(w, "Unable to load template", http.StatusInternalServerError)
			return
		}
		err = templ.ExecuteTemplate(w, "base", post)
		if err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		category := r.FormValue("category")
		content := r.FormValue("content")

		err = repositories.UpdatePost(db, id, title, category, content)
		if err != nil {
			http.Error(w, "Unable to update post: "+err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/post?id="+id.String(), http.StatusSeeOther)
		return
	}
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}
func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	if idString == "" {
		http.Error(w, "Post ID is missing", http.StatusBadRequest)
		return
	}
	id, err := uuid.Parse(idString)
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, "Unable to connect to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	err = repositories.DeletePost(db, id)
	if err != nil {
		http.Error(
			w,
			"Unable to delete post: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
