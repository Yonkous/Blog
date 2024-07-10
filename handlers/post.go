package handlers

import (
	"Blog/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post
	db.Find(&posts)
	err := json.NewEncoder(w).Encode(posts)
	if err != nil {
		return
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	db.First(&post, params["id"])
	json.NewEncoder(w).Encode(post)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Perform validation if needed

	if err := db.Create(&post).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var post models.Post
	json.NewDecoder(r.Body).Decode(&post)
	var dbPost models.Post
	db.First(&dbPost, params["id"])

	dbPost.Title = post.Title
	dbPost.Body = post.Body
	db.Save(&dbPost)
	json.NewEncoder(w).Encode(dbPost)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var post models.Post
	db.Delete(&post, params["id"])
	json.NewEncoder(w).Encode(post)
}
