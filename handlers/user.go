package handlers

import (
	"database/sql"
	"encoding/json"
	"my-blog/database"
	"my-blog/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)

    stmt, err := database.DB.Prepare("INSERT INTO users(username, email, password_hash) VALUES(?, ?, ?)")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    _, err = stmt.Exec(user.Username, user.Email, user.PasswordHash)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user models.User
    err = database.DB.QueryRow("SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?", id).Scan(
        &user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "User not found", http.StatusNotFound)
        } else {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
        return
    }
    json.NewEncoder(w).Encode(user)
}
