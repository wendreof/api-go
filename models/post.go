package models

//Post represents a BLOG POST
type Post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Text   string `json:"body"`
}
