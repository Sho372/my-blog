package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"my-blog/database"
	"my-blog/models"

	"github.com/gorilla/mux"
)

// CreatePost creates a new blog post
func CreatePost(w http.ResponseWriter, r *http.Request) {
    var post models.Post
    if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := database.DB.Create(&post).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(post)
}

// GetPost retrieves a post by ID
func GetPost(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    var post models.Post
    if err := database.DB.First(&post, id).Error; err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(post)
}

// GetPosts retrieves all posts
func GetPosts(w http.ResponseWriter, r *http.Request) {
    var posts []models.Post
    if err := database.DB.Find(&posts).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(posts)
}
