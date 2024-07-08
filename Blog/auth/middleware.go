package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// Middleware to check if the user is authorized based on their role
func Authorize(allowedRoles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Extract the JWT token from the request
			tokenStr := r.Header.Get("Authorization")
			if tokenStr == "" {
				http.Error(w, "No token provided", http.StatusUnauthorized)
				return
			}
			tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

			claims := &Claims{}
			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
			if err != nil || !token.Valid {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Check if the user's role is allowed
			for _, role := range allowedRoles {
				if claims.Role == role {
					// Add the claims to the context
					ctx := context.WithValue(r.Context(), "username", claims.Username)
					next.ServeHTTP(w, r.WithContext(ctx))
					return
				}
			}
			http.Error(w, "Forbidden", http.StatusForbidden)
		})
	}
}
