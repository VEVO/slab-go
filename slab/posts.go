package slab

import "time"

// Post represent the structure of a post
type Post struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Content     []byte     `json:"content,omitempty"`
	Version     *int       `json:"version,omitempty"`
	PublishedAt *time.Time `json:"publishedAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}
