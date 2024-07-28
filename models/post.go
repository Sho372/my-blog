package models

import (
	"database/sql"
	"time"
)

type Post struct {
    ID         int       `json:"id" gorm:"primary_key"`
    Title      string    `json:"title"`
    Content    string    `json:"content"`
    AuthorID   int       `json:"author_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
    PublishedAt NullTime `json:"published_at"`
}

type NullTime struct {
    sql.NullTime
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
    if !nt.Valid {
        return []byte(`null`), nil
    }
    return []byte(`"` + nt.Time.Format(time.RFC3339) + `"`), nil
}
