package models

import "time"

type Post struct {
	ID          int       `json:"id,omitempty"`
	UserID      string    `json:"userid,omitempty"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Image       string    `json:"image"`
	PublishedAt time.Time `json:"publishedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
