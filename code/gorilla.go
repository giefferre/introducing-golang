// START1 OMIT
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Setting up routes...")

	router := mux.NewRouter()

	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/foo", FooHandler)

	log.Println("Starting server on port 8081")

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// STOP1 OMIT

// START2 OMIT
func RootHandler(w http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprintln(w, "Welcome to my new website")
}

func FooHandler(w http.ResponseWriter, httpRequest *http.Request) {
	fmt.Fprintf(
		w,
		fmt.Sprintf("Your browser is: %s", httpRequest.UserAgent()),
	)
}

// STOP2 OMIT
