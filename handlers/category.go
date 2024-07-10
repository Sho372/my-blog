package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"my-blog/database"
	"my-blog/models"

	"github.com/gorilla/mux"
)

// CreateCategory creates a new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
    var category models.Category
    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := database.DB.Create(&category).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(category)
}

// GetCategory retrieves a category by ID
func GetCategory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid category ID", http.StatusBadRequest)
        return
    }

    var category models.Category
    if err := database.DB.First(&category, id).Error; err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(category)
}

// GetCategories retrieves all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
    var categories []models.Category
    if err := database.DB.Find(&categories).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(categories)
}
