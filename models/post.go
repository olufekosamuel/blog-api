package models

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
}
