package slab

import (
	"context"
)

// OrganizationService is an implementation of service for the organization
type OrganizationService service

// Organization represents the organization the token is currently assigned to.
type Organization struct {
	ID         string    `json:"id"`
	Host       string    `json:"host"`
	Name       string    `json:"name"`
	Posts      []Post    `json:"posts"`
	Topics     []Topic   `json:"topics"`
	Users      []User    `json:"users"`
	InsertedAt *DateTime `json:"insertedAt"`
	UpdatedAt  *DateTime `json:"updatedAt"`
}

// Get fetches the organization details and populate the struct with the details.
// Note that it does not fetch the hierarchies of the topics or the content of the posts
// or the user details to avoid issues with too much data
func (o *OrganizationService) Get() (*Organization, error) {
	query := `{
		organization {
			id,
			host,
			name,
			posts{id, title},
			topics{id,name, description, posts{id, title}},
			users{id, name},
			insertedAt,
			updatedAt
		}
	}`
	var resp struct {
		Organization *Organization `json:"organization"`
	}
	if err := o.client.Do(context.Background(), query, &resp); err != nil {
		return nil, err
	}
	return resp.Organization, nil

}
