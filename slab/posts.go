package slab

import (
	"context"
	"fmt"
)

// PostService is an implementation of the service to interact with the posts
type PostService service

// Post represent the structure of a post.
type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     []byte    `json:"content,omitempty"`
	Version     *int      `json:"version,omitempty"`
	PublishedAt *DateTime `json:"publishedAt,omitempty"`
	UpdatedAt   *DateTime `json:"updatedAt,omitempty"`
}

// GetAll retrieves all the posts available in the organization including their details
// but their content stays empty. Content is only available for filtered queries for now.
func (p *PostService) GetAll() (*[]Post, error) {
	query := `{
		organization {
			posts{
				id,
				title,
				version,
				insertedAt,
				publishedAt,
				updatedAt
			}
		}
	}`
	var resp struct {
		Organization *Organization `json:"organization"`
	}
	if err := p.client.Do(context.Background(), query, &resp); err != nil {
		return nil, err
	}
	return resp.Organization.Posts, nil
}

// Get retrieves the details of a specific post including its content
func (p *PostService) Get(id string) (*Post, error) {
	query := fmt.Sprintf(`{
		post(id: "%s"){
			id,
			title,
			version,
			content,
			insertedAt,
			publishedAt,
			updatedAt
		}
	}`, id)
	var resp struct {
		Post *Post `json:"post"`
	}
	if err := p.client.Do(context.Background(), query, &resp); err != nil {
		return nil, err
	}
	return resp.Post, nil

}
