package slab

// Topic is a level of the tree hierarchy to which posts are attached.
type Topic struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Posts       []Post    `json:"posts"`
	Hierarchy   *[]string `json:"hierarchy"`
	Parent      *Topic    `json:"parent,omitempty"`
	Ancestors   *[]Topic  `json:"ancestors"`
	Children    *[]Topic  `json:"children"`
	InsertedAt  *DateTime `json:"insertedAt"`
	UpdatedAt   *DateTime `json:"updatedAt"`
}
