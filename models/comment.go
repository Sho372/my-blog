package models

import "time"

type Comment struct {
    ID         int       `json:"id" gorm:"primary_key"`
    PostID     int       `json:"post_id"`
    AuthorName string    `json:"author_name"`
    Content    string    `json:"content"`
    CreatedAt  time.Time `json:"created_at"`
}
