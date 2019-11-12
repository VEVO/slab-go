package slab

import (
	"context"
	"strings"
)

// TopicService is an implementation of the service to interact with the posts
type TopicService service

// Topic is a level of the tree hierarchy to which posts are attached.
type Topic struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Posts       *[]Post   `json:"posts,omitempty"`
	Hierarchy   *[]string `json:"hierarchy,omitempty"`
	Parent      *Topic    `json:"parent,omitempty"`
	Ancestors   *[]Topic  `json:"ancestors,omitempty"`
	Children    *[]Topic  `json:"children,omitempty"`
	InsertedAt  *DateTime `json:"insertedAt,omitempty"`
	UpdatedAt   *DateTime `json:"updatedAt,omitempty"`
}

// List retrieves all the cwtopicsposts available in the organization including their details
func (t *TopicService) List() (*[]Topic, error) {
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
	err := t.client.Do(context.Background(), query, nil, &resp)
	if resp.Organization != nil {
		return resp.Organization.Topics, err
	}
	return nil, err
}

// Get retrieves the details of a specific topic
func (t *TopicService) Get(id string) (*Topic, error) {
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
	err := t.client.Do(context.Background(), query, vars, &resp)
	return resp.Topic, err
}

// Create inserts a new topic in slab
func (t *TopicService) Create(name, description, parentID string) (*Topic, error) {
	query := `
	mutation (
		$name: String!,
		$description: String,
		$parentId: ID
	){
		createTopic(
			name: $name, description: $description, parentId: $parentId
		){ id, name, description }
	}`
	var resp struct {
		Topic *Topic `json:"createTopic"`
	}
	vars := map[string]interface{}{"name": name, "description": description, "parentId": parentID}
	err := t.client.Do(context.Background(), query, vars, &resp)
	return resp.Topic, err
}

// AddToPost attaches a topic to a post
func (t *TopicService) AddToPost(topicID, postID string) (*Topic, error) {
	query := `
	mutation($postId: ID!, $topicId: ID!){
		addTopicToPost(postId: $postId, topicId: $topicId){ id, name, description }
	}`
	var resp struct {
		Topic *Topic `json:"addTopicToPost"`
	}
	vars := map[string]interface{}{"postId": postID, "topicId": topicID}
	err := t.client.Do(context.Background(), query, vars, &resp)
	return resp.Topic, err
}

// RemoveFromPost detaches a topic from a post
func (t *TopicService) RemoveFromPost(topicID, postID string) (*Topic, error) {
	query := `
	mutation($postId: ID!, $topicId: ID!){
		removeTopicFromPost(postId: $postId, topicId: $topicId){ id, name, description }
	}`
	var resp struct {
		Topic *Topic `json:"removeTopicFromPost"`
	}
	vars := map[string]interface{}{"postId": postID, "topicId": topicID}
	err := t.client.Do(context.Background(), query, vars, &resp)
	return resp.Topic, err
}

// AutoGenerate split the given `topicHierarcy` using the `separator` provided
// and returns the id of the leaf topic.
// This is handy to auto-generate topics based on a path of topic names.
//
// Example: AutoGenerate("Engineering/Services/slab-go", "/") will first search for the "Engineering"
// topic at the top level of the topics. If not found, create it.
// Then in its children it, will look for the "Services" topic and create it if not found.
// Then in the children topics of the "Services" topic, it will look for "slab-go" and create it of not found.
// Then it will return the topic ID of the "slab-go" topic we just mentionned.
//
// Note that the topic names are compared without taking the case in account so "EngineeRing" is the same as
// "engineering" in our example.
func (t *TopicService) AutoGenerate(topicHierarchy, separator string) (topicID string, err error) {
	topics, err := t.List()
	if err != nil {
		return "", err
	}

	items := strings.Split(topicHierarchy, separator)
	return t.autoCreate(nil, topics, items)
}

// autoCreate is where the real work of the AutoGenerate function happens. It just needed to be recursive to be
// efficient.
func (t *TopicService) autoCreate(parent *Topic, tree *[]Topic, hierarchy []string) (topicID string, err error) {
	item := hierarchy[0]
	var found *Topic

	for _, topic := range *tree {
		if ((parent == nil && topic.Parent == nil) || (parent != nil && topic.Parent != nil && topic.Parent.ID == parent.ID)) && strings.ToLower(topic.Name) == strings.ToLower(item) {
			found = &topic
			break
		}
	}

	// Topic not found, create it and all its children
	if found == nil {
		parentID := ""
		if parent != nil {
			parentID = parent.ID
		}
		for _, i := range hierarchy {
			if found, err = t.Create(i, "", parentID); err != nil {
				return "", err
			}
			parentID = found.ID
		}

		return found.ID, nil
	}
	if len(hierarchy) == 1 { // Last loop, just return the topic
		return found.ID, nil
	}
	return t.autoCreate(found, tree, hierarchy[1:])
}
