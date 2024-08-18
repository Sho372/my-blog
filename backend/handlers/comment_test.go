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

func TestCreateComment(t *testing.T) {
    // Set up the database
    database.DB, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    database.DB.AutoMigrate(&models.Comment{})

    // Create a new comment
    comment := models.Comment{
        Content: "test comment",
    }
    commentJSON, _ := json.Marshal(comment)

    req, err := http.NewRequest("POST", "/comments", bytes.NewBuffer(commentJSON))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateComment)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    var createdComment models.Comment
    json.Unmarshal(rr.Body.Bytes(), &createdComment)
    if createdComment.Content != comment.Content {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), commentJSON)
    }
}

func TestGetComment(t *testing.T) {
    // Set up the database
    database.DB, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    database.DB.AutoMigrate(&models.Comment{})

    // Create a new comment
    comment := models.Comment{
        Content: "test comment",
    }
    database.DB.Create(&comment)

    req, err := http.NewRequest("GET", "/comments/1", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()

    // Set up the router
    router := mux.NewRouter()
    router.HandleFunc("/comments/{id}", GetComment).Methods("GET")
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var returnedComment models.Comment
    json.Unmarshal(rr.Body.Bytes(), &returnedComment)
    if returnedComment.Content != comment.Content {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), comment)
    }
}
