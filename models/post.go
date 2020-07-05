package models

import "time"

type Post struct {
	ID          int       `json:"id,omitempty"`
	UserID      int       `json:"userid,omitempty"`
	Title       string    `json:"title"`
	Content     string    `json:"content,omitempty"`
	PublishedAt time.Time `json:"publishedAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
}
