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

func TestCreatePost(t *testing.T) {
    // Set up the database
    database.DB, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    database.DB.AutoMigrate(&models.Post{})

    // Create a new post
    post := models.Post{
        Title:   "test post",
        Content: "test content",
    }
    postJSON, _ := json.Marshal(post)

    req, err := http.NewRequest("POST", "/posts", bytes.NewBuffer(postJSON))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreatePost)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    var createdPost models.Post
    json.Unmarshal(rr.Body.Bytes(), &createdPost)
    if createdPost.Title != post.Title || createdPost.Content != post.Content {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), postJSON)
    }
}

func TestGetPost(t *testing.T) {
    // Set up the database
    database.DB, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    database.DB.AutoMigrate(&models.Post{})

    // Create a new post
    post := models.Post{
        Title:   "test post",
        Content: "test content",
    }
    database.DB.Create(&post)

    req, err := http.NewRequest("GET", "/posts/1", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()

    // Set up the router
    router := mux.NewRouter()
    router.HandleFunc("/posts/{id}", GetPost).Methods("GET")
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var returnedPost models.Post
    json.Unmarshal(rr.Body.Bytes(), &returnedPost)
    if returnedPost.Title != post.Title || returnedPost.Content != post.Content {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), post)
    }
}
