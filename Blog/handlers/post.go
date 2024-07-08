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
	var post []models.Post
	json.NewDecoder(r.Body).Decode(&post)
	err := db.Create(&post).Error
	if err != nil {
		return
	}
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
