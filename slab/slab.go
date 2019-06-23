// This code is inspired by https://github.com/shurcooL/githubv4/blob/master/githubv4.go

package slab

import (
	"context"
	"net/http"

	"github.com/shurcooL/graphql"
)

const apiEndpoint = "https://api.slab.com/v1/graphql"

// Client is the client used for the graphql api
type Client struct {
	client *graphql.Client
	// APIToken is the authentication token to use when talking to the slab API
	APIToken string
}

// NewClient creates a new slab client with the provided http.Client.
// If httpClient is nil, then http.DefaultClient is used.
func NewClient(httpClient *http.Client, apiToken string) *Client {
	return &Client{
		client:   graphql.NewClient(apiEndpoint, httpClient),
		APIToken: apiToken,
	}
}

// Query executes a single GraphQL query request,
// with a query derived from q, populating the response into it.
// q should be a pointer to struct that corresponds to the GraphQL schema.
func (c *Client) Query(ctx context.Context, q interface{}, variables map[string]interface{}) error {
	return c.client.Query(ctx, q, variables)
}

// Mutate executes a single GraphQL mutation request,
// with a mutation derived from m, populating the response into it.
// m should be a pointer to struct that corresponds to the GraphQL schema.
// Provided input will be set as a variable named "input".
func (c *Client) Mutate(ctx context.Context, m interface{}, input Input, variables map[string]interface{}) error {
	if variables == nil {
		variables = map[string]interface{}{"input": input}
	} else {
		variables["input"] = input
	}
	return c.client.Mutate(ctx, m, variables)
}
