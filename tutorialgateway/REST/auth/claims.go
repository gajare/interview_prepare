package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// =============================================
// JWT Config
// =============================================
var (
	JWT_SECRET = []byte("your-super-secret-jwt-key-change-in-prod") // Use env var!
	JWT_EXPIRY = 15 * time.Minute
)

// User claims
type Claims struct {
	UserID int      `json:"user_id"`
	Role   string   `json:"role"`
	Scopes []string `json:"scopes,omitempty"`
	jwt.RegisteredClaims
}

// =============================================
// Helper: Generate JWT
// =============================================
func generateToken(userID int, role string, scopes []string) (string, error) {
	claims := Claims{
		UserID: userID,
		Role:   role,
		Scopes: scopes,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JWT_EXPIRY)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "demo-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT_SECRET)
}

// =============================================
// 1. JWT Authentication Middleware
// =============================================
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization format. Use: Bearer <token>", http.StatusUnauthorized)
			return
		}

		tokenStr := parts[1]
		claims := &Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Inject user into context
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
