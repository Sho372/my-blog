package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"my-blog/database"
	"my-blog/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CreateComment creates a new comment
func CreateComment(w http.ResponseWriter, r *http.Request) {
    var comment models.Comment
    if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := database.DB.Create(&comment).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(comment)
}

// GetComment retrieves a comment by ID
func GetComment(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid comment ID", http.StatusBadRequest)
        return
    }

    var comment models.Comment
    if err := database.DB.First(&comment, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "Comment not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(comment)
}

// GetComments retrieves comments for a post
func GetComments(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    postID, err := strconv.Atoi(vars["post_id"])
    if err != nil {
        http.Error(w, "Invalid post ID", http.StatusBadRequest)
        return
    }

    var comments []models.Comment
    if err := database.DB.Where("post_id = ?", postID).Find(&comments).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(comments)
}
