package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, Go 1.22+ mux!\n")
	})

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "User ID: %s\n", id)
	})

	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Created user\n")
	})

	mux.HandleFunc("DELETE /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "Deleted user %s\n", id)
	})

	mux.HandleFunc("GET /articles/{category}/{slug}", func(w http.ResponseWriter, r *http.Request) {
		cat := r.PathValue("category")
		slug := r.PathValue("slug")
		fmt.Fprintf(w, "Article: %s/%s\n", cat, slug)
	})

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
