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

func TestCreateCategory(t *testing.T) {
    // Set up the database
    database.DB, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    database.DB.AutoMigrate(&models.Category{})

    // Create a new category
    category := models.Category{
        Name: "testcategory",
    }
    categoryJSON, _ := json.Marshal(category)

    req, err := http.NewRequest("POST", "/categories", bytes.NewBuffer(categoryJSON))
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CreateCategory)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusCreated {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusCreated)
    }

    var createdCategory models.Category
    json.Unmarshal(rr.Body.Bytes(), &createdCategory)
    if createdCategory.Name != category.Name {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), categoryJSON)
    }
}

func TestGetCategory(t *testing.T) {
    // Set up the database
    database.DB, _ = gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
    database.DB.AutoMigrate(&models.Category{})

    // Create a new category
    category := models.Category{
        Name: "testcategory",
    }
    database.DB.Create(&category)

    req, err := http.NewRequest("GET", "/categories/1", nil)
    if err != nil {
        t.Fatal(err)
    }
    rr := httptest.NewRecorder()

    // Set up the router
    router := mux.NewRouter()
    router.HandleFunc("/categories/{id}", GetCategory).Methods("GET")
    router.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    var returnedCategory models.Category
    json.Unmarshal(rr.Body.Bytes(), &returnedCategory)
    if returnedCategory.Name != category.Name {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), category)
    }
}
