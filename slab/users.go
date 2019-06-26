package slab

import "context"

// UserService is an implementation of the service to interact with the users
type UserService service

// User represents a given slab user
type User struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Email         string    `json:"email"`
	Title         string    `json:"title"`
	Type          string    `json:"type"`
	Avatar        *Image    `json:"avatar,omitempty"`
	InsertedAt    *DateTime `json:"insertedAt,omitempty"`
	UpdatedAt     *DateTime `json:"updatedAt,omitempty"`
	DeactivatedAt *DateTime `json:"deactivatedAt,omitempty"`
}

// Image is the type of the avatar which is a url to the avatar if any
type Image struct {
	Original string `json:"original"`
	Thumb    string `json:"thumb"`
}

// List retrieves all the users available in the organization including their details
func (p *UserService) List() (*[]User, error) {
	query := `{
		organization {
			users{
				id,
				name,
				description,
				email,
				title,
				type,
				avatar{original, thumb},
				insertedAt,
				deactivatedAt,
				updatedAt
			}
		}
	}`
	var resp struct {
		Organization *Organization `json:"organization"`
	}
	err := p.client.Do(context.Background(), query, nil, &resp)
	return resp.Organization.Users, err
}

// Get retrieves the details of a specific user including its content
func (p *UserService) Get(id string) (*User, error) {
	query := `
    query ($id: ID){
		user(id: $id){
			id,
			name,
			description,
			email,
			title,
			type,
			avatar{original, thumb},
			insertedAt,
			deactivatedAt,
			updatedAt
		}
    }`
	var resp struct {
		User *User `json:"user"`
	}
	vars := map[string]interface{}{"id": id}
	err := p.client.Do(context.Background(), query, vars, &resp)
	return resp.User, err

}
