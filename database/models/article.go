package models

// Article represents a blog post
type Article struct {
	ID            int    `json:"id"`
	UID           string `gorm:"column:uid" json:"uid"`
	Title         string `json:"title"`
	TitleEncoded  string `json:"title_encoded"`
	Content       string `json:"content"`
	ContentFormat string `json:"content_format"`
	Tags          string `json:"tags"`
	Type          string `json:"type"`
	Author        User   `json:"author"`
	AuthorID      int    `json:"author_id"`
}
