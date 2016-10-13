package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	// START1 OMIT
	log.Println("Setting up routes...")

	http.HandleFunc("/world", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Starting server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
	// STOP1 OMIT
}
