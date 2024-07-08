package handlers

import (
	"Blog/auth"
	"encoding/json"
	"net/http"
	"time"
)

var users = map[string]string{
	"user":   "userpassword",
	"admin":  "adminpassword",
	"editor": "editorpassword",
}

// LoginHandler handles user login and JWT generation
func GenerateTokenHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate user credentials
	storedPassword, ok := users[creds.Username]
	if !ok || storedPassword != creds.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Determine user role
	role := "user"
	if creds.Username == "admin" {
		role = "admin"
	} else if creds.Username == "editor" {
		role = "editor"
	}

	// Generate JWT token
	token, err := auth.GenerateJWT(creds.Username, role)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Set the token in the response
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(30 * time.Minute),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}
