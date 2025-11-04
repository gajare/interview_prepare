package main

import (
	"encoding/json"
	"net/http"
)

// =============================================
// 2. Authorization Middleware (Role-based)
// =============================================
func RequireRole(requiredRole string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value("user").(*Claims)
			if !ok || user.Role != requiredRole {
				http.Error(w, "Forbidden: insufficient role", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

// Optional: Scope-based authorization
func RequireScope(requiredScope string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, ok := r.Context().Value("user").(*Claims)
			if !ok {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}
			for _, s := range user.Scopes {
				if s == requiredScope {
					next.ServeHTTP(w, r)
					return
				}
			}
			http.Error(w, "Forbidden: missing scope", http.StatusForbidden)
		})
	}
}

// =============================================
// 3. Login Endpoint (Generate JWT)
// =============================================
func loginHandler(w http.ResponseWriter, r *http.Request) {
	// In real app: validate username/password from DB
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Dummy auth
	if username == "admin" && password == "secret" {
		token, _ := generateToken(1, "admin", []string{"read:secret", "write:users"})
		writeJSON(w, http.StatusOK, map[string]string{"token": token}, "Login successful")
		return
	}
	if username == "user" && password == "pass" {
		token, _ := generateToken(2, "user", []string{"read:public"})
		writeJSON(w, http.StatusOK, map[string]string{"token": token}, "Login successful")
		return
	}

	http.Error(w, "Invalid credentials", http.StatusUnauthorized)
}

// =============================================
// 4. Protected GET Route (Demo)
// =============================================
func secretHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(*Claims)
	writeJSON(w, http.StatusOK, map[string]any{
		"user_id": user.UserID,
		"role":    user.Role,
		"message": "You accessed the secret area!",
	}, "Authenticated & Authorized")
}

// =============================================
// JSON Helper
// =============================================
func writeJSON(w http.ResponseWriter, status int, data any, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]any{
		"message": msg,
		"data":    data,
	})
}
