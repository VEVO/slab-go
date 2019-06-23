package slab

import "time"

// User represents a given slab user
type User struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	Email         string     `json:"email"`
	Title         string     `json:"title"`
	Type          string     `json:"type"`
	Avatar        *[]byte    `json:"avatar,omitempty"`
	InsertedAt    *time.Time `json:"insertedAt"`
	UpdatedAt     *time.Time `json:"updatedAt"`
	DeactivatedAt *time.Time `json:"deactivatedAt"`
}
