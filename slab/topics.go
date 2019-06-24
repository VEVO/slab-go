package slab

import (
	"context"
)

// TopicService is an implementation of the service to interact with the posts
type TopicService service

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

// List retrieves all the cwtopicsposts available in the organization including their details
func (p *TopicService) List() (*[]Topic, error) {
	query := `{
        organization {
            topics{
                id,
                name,
                description,
				posts{id, title},
				hierarchy,
				parent{id},
				ancestors{id},
				children{id},
                insertedAt,
                updatedAt
            }
        }
    }`
	var resp struct {
		Organization *Organization `json:"organization"`
	}
	err := p.client.Do(context.Background(), query, nil, &resp)
	return resp.Organization.Topics, err
}

// Get retrieves the details of a specific topic
func (p *TopicService) Get(id string) (*Topic, error) {
	query := `
    query ($id: ID){
        topic(id: $id){
                id,
                name,
                description,
				posts{id, title},
				hierarchy,
				parent{id},
				ancestors{id},
				children{id},
                insertedAt,
                updatedAt
        }
    }`
	var resp struct {
		Topic *Topic `json:"topic"`
	}
	vars := map[string]interface{}{"id": id}
	err := p.client.Do(context.Background(), query, vars, &resp)
	return resp.Topic, err

}
