package models

import "time"

type Post struct {
    ID         int       `json:"id" gorm:"primary_key"`
    Title      string    `json:"title"`
    Content    string    `json:"content"`
    AuthorID   int       `json:"author_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    PublishedAt time.Time `json:"published_at"`
}
