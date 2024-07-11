package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"my-blog/database"
	"my-blog/models"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
    // Set up the database
    database.DB, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    database.DB.AutoMigrate(&models.User{})

    // Create a new user
    user := models.User{
        Username: "testuser",
        Email:    "testuser@example.com",
    }
    userJSON, _ := json.Marshal(user)

    req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(userJSON))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateUser)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    var createdUser models.User
    json.Unmarshal(rr.Body.Bytes(), &createdUser)
    if createdUser.Username != user.Username || createdUser.Email != user.Email {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), userJSON)
    }
}

func TestGetUser(t *testing.T) {
    // Set up the database
    database.DB, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    database.DB.AutoMigrate(&models.User{})

    // Create a new user
    user := models.User{
        Username: "testuser",
        Email:    "testuser@example.com",
    }
    database.DB.Create(&user)

    req, err := http.NewRequest("GET", "/users/1", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()

    // Set up the router
    router := mux.NewRouter()
    router.HandleFunc("/users/{id}", GetUser).Methods("GET")
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var returnedUser models.User
    json.Unmarshal(rr.Body.Bytes(), &returnedUser)
    if returnedUser.Username != user.Username || returnedUser.Email != user.Email {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), user)
    }
}
