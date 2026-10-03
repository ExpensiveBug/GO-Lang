package main

import (
	"BLogPost/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := routes.SetupRoutes()

	fmt.Println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
