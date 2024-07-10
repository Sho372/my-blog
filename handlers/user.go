package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"my-blog/database"
	"my-blog/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := database.DB.Create(&user).Error; err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

// GetUser retrieves a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    fmt.Println("xxxx")
    var user models.User
    if err := database.DB.First(&user, id).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            fmt.Println("hhhh")
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }

    json.NewEncoder(w).Encode(user)
}
