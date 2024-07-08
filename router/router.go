package router

import (
	"Blog/auth"
	"Blog/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	// Root endpoint
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")

	// Auth routes
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")

	// Posts endpoints
	router.HandleFunc("/posts", handlers.GetPosts).Methods("GET")
	router.HandleFunc("/posts", handlers.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", handlers.GetPost).Methods("GET")

	// Admin and editor routes
	adminEditorRouter := router.PathPrefix("/posts").Subrouter()
	adminEditorRouter.Use(auth.Authorize("admin", "editor", "user"))
	adminEditorRouter.HandleFunc("/{id}", handlers.UpdatePost).Methods("PUT")
	adminEditorRouter.HandleFunc("/{id}", handlers.DeletePost).Methods("DELETE")

	return router
}
