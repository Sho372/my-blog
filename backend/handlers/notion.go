package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/jomei/notionapi"
)

type NotionHandler struct {
	client *notionapi.Client
}

func NewNotionHandler() *NotionHandler {
	token := os.Getenv("NOTION_TOKEN")
	if token == "" {
		return &NotionHandler{client: nil}
	}
	
	client := notionapi.NewClient(notionapi.Token(token))
	return &NotionHandler{client: client}
}

type NotionPage struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	CreatedTime string `json:"created_time"`
	LastEdited  string `json:"last_edited_time"`
	Icon        string `json:"icon,omitempty"`
}

func (h *NotionHandler) GetPages(w http.ResponseWriter, r *http.Request) {
	if h.client == nil {
		http.Error(w, "Notion API not configured", http.StatusInternalServerError)
		return
	}

	databaseID := os.Getenv("NOTION_DATABASE_ID")
	if databaseID == "" {
		http.Error(w, "Notion database ID not configured", http.StatusInternalServerError)
		return
	}

	// Query the database
	query := &notionapi.DatabaseQueryRequest{
		PageSize: 100,
		Sorts: []notionapi.SortObject{
			{
				Property:  "Last edited time",
				Direction: "descending",
			},
		},
	}

	resp, err := h.client.Database.Query(r.Context(), notionapi.DatabaseID(databaseID), query)
	if err != nil {
		http.Error(w, "Failed to fetch pages from Notion", http.StatusInternalServerError)
		return
	}

	var pages []NotionPage
	for _, page := range resp.Results {
		notionPage := NotionPage{
			ID:          string(page.ID),
			URL:         string(page.URL),
			CreatedTime: page.CreatedTime.String(),
			LastEdited:  page.LastEditedTime.String(),
		}

		// Extract title from properties
		if titleProp, exists := page.Properties["Title"]; exists {
			if titleProp.GetType() == "title" {
				title := titleProp.(*notionapi.TitleProperty)
				if len(title.Title) > 0 {
					notionPage.Title = title.Title[0].PlainText
				}
			}
		}

		// Extract icon if exists (simplified for now)
		if page.Icon != nil {
			// Icon handling can be added later if needed
		}

		pages = append(pages, notionPage)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"pages": pages,
		"total": len(pages),
	})
}