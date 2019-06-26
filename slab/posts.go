package slab

import (
	"context"
)

// PostService is an implementation of the service to interact with the posts
type PostService service

// Post represent the structure of a post.
type Post struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Content     *string   `json:"content,omitempty"`
	Version     int       `json:"version,omitempty"`
	InsertedAt  *DateTime `json:"insertedAt,omitempty"`
	PublishedAt *DateTime `json:"publishedAt,omitempty"`
	UpdatedAt   *DateTime `json:"updatedAt,omitempty"`
}

// List retrieves all the posts available in the organization including their details
// but their content stays empty. Content is only available for filtered queries for now.
func (p *PostService) List() (*[]Post, error) {
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
	err := p.client.Do(context.Background(), query, nil, &resp)
	return resp.Organization.Posts, err
}

// Get retrieves the details of a specific post including its content
func (p *PostService) Get(id string) (*Post, error) {
	query := `
	query ($id: ID){
		post(id: $id){
			id,
			title,
			version,
			content,
			insertedAt,
			publishedAt,
			updatedAt
		}
	}`
	var resp struct {
		Post *Post `json:"post"`
	}
	vars := map[string]interface{}{"id": id}
	err := p.client.Do(context.Background(), query, vars, &resp)
	return resp.Post, err

}

// Create creates a new blank post, optionally organized in given topicId.
func (p *PostService) Create(topicID string) (*Post, error) {
	query := `mutation ($topicId: ID){ createPost(topicId: $topicId){ id } }`
	var resp struct {
		Post *Post `json:"createPost"`
	}
	vars := map[string]interface{}{"topicId": topicID}
	err := p.client.Do(context.Background(), query, vars, &resp)
	return resp.Post, err
}

/*
updatePost has been disabled for now by the slab team. They're rewriting it so disabling that for now.

// Update changes the post to match the one given in parameter and/or its publication state.
// The content parameter should be a field that matches
func (p *PostService) Update(id, content string, published bool) (*Post, error) {
	query := `mutation ($id: ID, $content: Json, $published: Boolean){
		updatePost(id: $id, content: $content, published: $published){
			id,
			title,
			version,
			content,
			insertedAt,
			publishedAt,
			updatedAt
		}
	}`
	var resp struct {
		Post *Post `json:"updatePost"`
	}
	vars := map[string]interface{}{"id": id, "content": content, "published": published}
	err := p.client.Do(context.Background(), query, vars, &resp)
	return resp.Post, err
}
*/

// Delete deletes a post with given id or externalID. At least one must be supplied. If both are, id is used.
func (p *PostService) Delete(id, externalID string) (*Post, error) {
	query := `mutation($id: ID, $externalId: ID){ deletePost(id: $id, externalId: $externalId){ id } }`
	var resp struct {
		Post *Post `json:"deletePost"`
	}
	vars := make(map[string]interface{})
	if id != "" {
		vars["id"] = id
	} else {
		vars["externalId"] = externalID
	}
	err := p.client.Do(context.Background(), query, vars, &resp)
	return resp.Post, err
}

// Sync Creates or updates a post on Slab that is a readonly copy of a post stored externally. The editUrl and readUrl
// are links to the externally stored source and will be shown and linked in some parts of the Slab UI. Upon creation,
// at least the editUrl must be supplied. If no readUrl is supplied upon creation, the editUrl is used.
//
// To clarify:
// * `externalID` is the identifier that identifies your post and that you will use, for example,
//   when you want to delete the post.
// * `editUR`L is the url you will be redirected to when you hit the "Edit Post" button in the slab UI.
// * currently accepted `format` fields are `HTML` or `MARKDOWN`.
func (p *PostService) Sync(externalID, content, editURL, readURL, format string) (*Post, error) {
	query := `
	mutation(
		$content: String!
		$editUrl: String
		$externalId: ID!
		$format: PostContentFormat!
		$readUrl: String
	){
		syncPost(
			content: $content
			editUrl: $editUrl
			externalId: $externalId
			format: $format
			readUrl: $readUrl
		){
			id,
			title,
			version,
			content,
			insertedAt,
			publishedAt,
			updatedAt
		}
}`
	var resp struct {
		Post *Post `json:"syncPost"`
	}
	vars := map[string]interface{}{
		"content":    content,
		"editUrl":    editURL,
		"externalId": externalID,
		"format":     format,
		"readUrl":    readURL,
	}
	err := p.client.Do(context.Background(), query, vars, &resp)
	return resp.Post, err
}

// AddTopic attaches a given topic to a given post
//
// Note that this is calling `Topic.AddToPost` and is only put here for convenience
func (p *PostService) AddTopic(postID, topicID string) error {
	_, err := p.client.Topic.AddToPost(topicID, postID)
	return err
}

// RemoveTopic detaches a given topic to a given post
//
// Note that this is calling `Topic.RemoveFromPost` and is only put here for convenience
func (p *PostService) RemoveTopic(postID, topicID string) error {
	_, err := p.client.Topic.RemoveFromPost(topicID, postID)
	return err
}
