package handlers

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestNotionHandler_GetPages_NoToken(t *testing.T) {
	// Clear environment variables for testing
	originalToken := os.Getenv("NOTION_TOKEN")
	originalDBID := os.Getenv("NOTION_DATABASE_ID")
	defer func() {
		os.Setenv("NOTION_TOKEN", originalToken)
		os.Setenv("NOTION_DATABASE_ID", originalDBID)
	}()
	
	os.Unsetenv("NOTION_TOKEN")
	os.Unsetenv("NOTION_DATABASE_ID")

	handler := NewNotionHandler()
	req := httptest.NewRequest("GET", "/notion/pages", nil)
	w := httptest.NewRecorder()

	handler.GetPages(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestNotionHandler_GetPages_NoDatabaseID(t *testing.T) {
	// Set token but not database ID
	originalToken := os.Getenv("NOTION_TOKEN")
	originalDBID := os.Getenv("NOTION_DATABASE_ID")
	defer func() {
		os.Setenv("NOTION_TOKEN", originalToken)
		os.Setenv("NOTION_DATABASE_ID", originalDBID)
	}()
	
	os.Setenv("NOTION_TOKEN", "test_token")
	os.Unsetenv("NOTION_DATABASE_ID")

	handler := NewNotionHandler()
	req := httptest.NewRequest("GET", "/notion/pages", nil)
	w := httptest.NewRecorder()

	handler.GetPages(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
	}
}