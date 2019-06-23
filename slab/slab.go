package slab

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/machinebox/graphql"
)

const apiEndpoint = "https://api.slab.com/v1/graphql"

// Client is the client used for the graphql api
type Client struct {
	client *graphql.Client
	// APIToken is the authentication token to use when talking to the slab API
	APIToken string

	common       service
	Organization *OrganizationService
	Post         *PostService
}

// NewClient creates a new slab client with the provided http.Client.
// If httpClient is nil, then http.DefaultClient is used.
func NewClient(httpClient *http.Client, apiToken string) *Client {
	c := &Client{
		client:   graphql.NewClient(apiEndpoint, graphql.WithHTTPClient(httpClient)),
		APIToken: apiToken,
	}
	// For debugging
	//c.client.Log = func(s string) { log.Println(s) }
	c.common.client = c
	c.Organization = (*OrganizationService)(&c.common)
	c.Post = (*PostService)(&c.common)

	return c
}

type service struct {
	client *Client
}

// Do executes the given query and populates the resp struct
func (c *Client) Do(ctx context.Context, query string, resp interface{}) error {
	req := graphql.NewRequest(query)
	req.Header.Set("Authorization", c.APIToken)
	if err := c.client.Run(ctx, req, resp); err != nil {
		return err
	}
	return nil
}

// DateTime is a struct that allow us to unmarshal the RFC3339 date formats
type DateTime struct {
	time.Time
}

func (t DateTime) String() string {
	return t.Time.String()
}

// UnmarshalJSON is used to unmarshal the date to json
func (t *DateTime) UnmarshalJSON(data []byte) (err error) {
	str := string(data)
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		t.Time = time.Unix(i, 0)
	} else {
		t.Time, err = time.Parse(`"`+time.RFC3339+`"`, str)
	}
	return
}
