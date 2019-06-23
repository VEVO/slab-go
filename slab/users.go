package slab

// User represents a given slab user
type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Email         string    `json:"email"`
	Title         string    `json:"title"`
	Type          string    `json:"type"`
	Avatar        *Image    `json:"avatar,omitempty"`
	InsertedAt    *DateTime `json:"insertedAt"`
	UpdatedAt     *DateTime `json:"updatedAt"`
	DeactivatedAt *DateTime `json:"deactivatedAt"`
}

// Image is the type of the avatar which is a url to the avatar if any
type Image struct {
	Original string `json:"original"`
	Thumb    string `json:"thumb"`
}
