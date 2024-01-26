package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	bookRouter := r.PathPrefix("/books").Subrouter()

	bookRouter.HandleFunc("/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	bookRouter.HandleFunc("/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "This a route to create a new book: %s\n", vars["title"])
	}).Methods("POST")

	bookRouter.HandleFunc("/{title}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "This is a route to get a book: %s\n", vars["title"])
	}).Methods("GET")

	bookRouter.HandleFunc("/{id}", UpdateBook).Methods("PUT")
	bookRouter.HandleFunc("/{id}", DeleteBook).Methods("DELETE")

	http.ListenAndServe(":80", r)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "This is a route to update a book: %s\n", vars["id"])
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "This is a route to delete a book: %s\n", vars["id"])
}
