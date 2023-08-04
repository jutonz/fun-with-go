package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func PostsList(repo *gorm.DB) []Post {
	var posts []Post
	repo.Find(&posts)
	return posts
}
