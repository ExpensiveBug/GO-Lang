package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm Error! %v\n", err)
		return
	}
	fmt.Fprintf(w, "Post Request Successful :) ")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Address : %s\n", address)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 NOT FOUND", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello")
}

func main() {
	// setting up the server
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
