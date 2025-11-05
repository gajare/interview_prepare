// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler inside the subrouter — uses mux.Vars(r)
func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Works even in subrouter!
	id := vars["id"]
	name := r.URL.Query().Get("name") // Query param also works

	fmt.Fprintf(w, "Subrouter GET: User ID = %s, Name = %s\n", id, name)
}

func main() {
	// 1. Main router
	mainRouter := mux.NewRouter()

	// 2. Create a SUBROUTER with prefix /api/v1
	apiV1 := mainRouter.PathPrefix("/api/v1").Subrouter()

	// 3. Register GET route WITH path variable {id}
	apiV1.HandleFunc("/users/{id}", getUser).Methods("GET")

	// Optional: Add middleware only to /api/v1/*
	apiV1.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("[API Middleware] %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	// 4. Start server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mainRouter))
}
