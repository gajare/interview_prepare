// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

// ===========================================
// 1. Role Middleware (simulates auth)
// ===========================================
func roleMiddleware(role string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Simulate getting role from JWT/session
			userRole := r.Header.Get("X-User-Role")

			if userRole != role {
				http.Error(w, fmt.Sprintf("Forbidden: requires %s role", role), http.StatusForbidden)
				return
			}

			log.Printf("[AUTH] %s accessed %s", userRole, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	}
}

// ===========================================
// 2. Handlers (one per role + path var)
// ===========================================

// ADMIN: /admin/users/{id}
func adminGetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "ADMIN: Viewing user %s (sensitive data)\n", id)
}

// EMPLOYEE: /employee/tasks/{taskID}
func employeeGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["taskID"]
	fmt.Fprintf(w, "EMPLOYEE: Working on task %s\n", taskID)
}

// CUSTOMER: /customer/orders/{orderID}
func customerGetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["orderID"]
	status := r.URL.Query().Get("status")
	fmt.Fprintf(w, "CUSTOMER: Order %s, Status: %s\n", orderID, status)
}

// ===========================================
// 3. Setup Subrouters by Role
// ===========================================
func setupAdminRouter() *mux.Router {
	admin := mux.NewRouter().PathPrefix("/admin").Subrouter()
	admin.Use(roleMiddleware("admin")) // Only admins

	admin.HandleFunc("/users/{id}", adminGetUser).Methods("GET")
	admin.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "ADMIN: Dashboard loaded\n")
	}).Methods("GET")

	return admin
}

func setupEmployeeRouter() *mux.Router {
	emp := mux.NewRouter().PathPrefix("/employee").Subrouter()
	emp.Use(roleMiddleware("employee"))

	emp.HandleFunc("/tasks/{taskID}", employeeGetTask).Methods("GET")
	emp.HandleFunc("/schedule", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "EMPLOYEE: Schedule view\n")
	}).Methods("GET")

	return emp
}

func setupCustomerRouter() *mux.Router {
	cust := mux.NewRouter().PathPrefix("/customer").Subrouter()
	cust.Use(roleMiddleware("customer"))

	cust.HandleFunc("/orders/{orderID}", customerGetOrder).Methods("GET")
	cust.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "CUSTOMER: Profile page\n")
	}).Methods("GET")

	return cust
}

// ===========================================
// 4. Main – Wire All Subrouters
// ===========================================
func main() {
	mainRouter := mux.NewRouter()

	// Register all subrouters
	mainRouter.PathPrefix("/admin").Handler(setupAdminRouter())
	mainRouter.PathPrefix("/employee").Handler(setupEmployeeRouter())
	mainRouter.PathPrefix("/customer").Handler(setupCustomerRouter())

	// Optional: Root
	mainRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome! Use /admin, /employee, or /customer\n")
	}).Methods("GET")

	// Server with graceful shutdown
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mainRouter,
	}

	go func() {
		log.Println("Server starting on http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Server stopped.")
}
