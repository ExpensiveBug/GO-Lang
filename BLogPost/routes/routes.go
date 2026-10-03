package routes

import (
	"net/http"

	"BLogPost/handlers"
)

func SetupRoutes() *http.ServeMux {

	// own servermux
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/post", handlers.PostHandler)
	mux.HandleFunc("/create_post", handlers.CreatePostHandler)
	mux.HandleFunc("/edit_post", handlers.EditPostHandler)
	mux.HandleFunc("/delete_post", handlers.DeletePostHandler)
	mux.HandleFunc("/about", handlers.AboutHandler)
	mux.HandleFunc("/help", handlers.HelpHandler)
	mux.HandleFunc("/rules", handlers.RulesHandler)

	staticfileserver := http.FileServer(http.Dir("./static"))
	// mux.Handle("/static/",http.StripPrefix("/static/",staticFileServer))
	mux.Handle("/static/", staticfileserver)

	return mux
}
