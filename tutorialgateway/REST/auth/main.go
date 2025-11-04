// main.go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/login", loginHandler).Methods(http.MethodPost)

	// Protected API with JWT + Role
	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(JWTMiddleware) // All routes under /api/v1 require JWT

	// Admin-only route
	admin := api.NewRoute().Subrouter()
	admin.HandleFunc("/secret", secretHandler).Methods(http.MethodGet)
	admin.Use(RequireRole("admin")) // Only admin role

	// User with scope
	api.HandleFunc("/public", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, nil, "Public data visible to any authenticated user")
	}).Methods(http.MethodGet)

	// Scope-protected
	scoped := api.NewRoute().Subrouter()
	scoped.HandleFunc("/scoped", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, nil, "You have read:secret scope!")
	}).Methods(http.MethodGet)
	scoped.Use(RequireScope("read:secret"))

	// Server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("Server starting on :8080")
		log.Println("Try: POST /login → GET /api/v1/secret (admin only)")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Server stopped.")
}
