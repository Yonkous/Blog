package handlers

import (
	"Blog/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetPosts handles fetching all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	db.Find(&posts)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		http.Error(w, "Failed to encode posts", http.StatusInternalServerError)
		return
	}
}

// GetPost handles fetching a single post by ID
func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post models.Post
	result := db.First(&post, postID)
	if result.Error != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, "Failed to encode post", http.StatusInternalServerError)
		return
	}
}

// CreatePost handles creating a new post
func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// Perform validation if needed

	result := db.Create(&post)
	if result.Error != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// UpdatePost handles updating an existing post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var updatedPost models.Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var existingPost models.Post
	result := db.First(&existingPost, postID)
	if result.Error != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	existingPost.Title = updatedPost.Title
	existingPost.Body = updatedPost.Body
	db.Save(&existingPost)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingPost)
}

// DeletePost handles deleting a post by ID
func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	var post models.Post
	result := db.Delete(&post, postID)
	if result.Error != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
