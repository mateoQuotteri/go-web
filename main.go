package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/users", UserServer)
	/*
		http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		})
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Hello GET, %q", html.EscapeString(r.URL.Path))
	case http.MethodPost:
		fmt.Fprintf(w, "Hello POST, %q", html.EscapeString(r.URL.Path))

	default:
		fmt.Fprintf(w, "NOT FOUND, %q", html.EscapeString(r.URL.Path))

	}
}
