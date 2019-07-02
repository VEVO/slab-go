package slab

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTopicService_List(t *testing.T) {
	insertDate := &DateTime{time.Date(2019, time.May, 1, 22, 44, 33, 78957000, time.UTC)}
	updateDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 16, 733422000, time.UTC)}
	want := &[]Topic{
		{
			ID: "abc123", Hierarchy: &[]string{"efg234.abc123"}, Ancestors: &[]Topic{{ID: "efg234"}},
			Description: "Description of topic A.", InsertedAt: insertDate, UpdatedAt: updateDate,
			Name: "Topic A", Parent: &Topic{ID: "efg234"}, Children: &[]Topic{{ID: "zzzblabla"}},
			Posts: &[]Post{{ID: "postid1", Title: "Post 1 from topic A"}, {ID: "postid2", Title: "Post 2 from topic A"}},
		},
		{
			ID: "zzzblabla", Hierarchy: &[]string{"efg234.abc123.zzzblabla"},
			Ancestors:   &[]Topic{{ID: "abc123"}, {ID: "efg234"}},
			Description: "Description of topic B.", InsertedAt: insertDate, UpdatedAt: updateDate,
			Name: "Topic B", Parent: &Topic{ID: "abc123"}, Children: &[]Topic{},
			Posts: &[]Post{{ID: "postid3", Title: "Post 3 from topic B"}, {ID: "postid4", Title: "Post 4 from topic B"}},
		},
	}
	expectedResp := `{"data":{"organization":{"topics": [
                {
                    "id": "abc123",
                    "hierarchy": ["efg234.abc123"],
				    "ancestors": [{"id":"efg234"}],
                    "children": [{ "id": "zzzblabla"}],
                    "description": "Description of topic A.",
					"insertedAt": "2019-05-01T22:44:33.078957Z",
					"updatedAt":"2019-06-18T22:40:16.733422Z",
                    "name": "Topic A",
                    "parent": {"id": "efg234"},
                    "posts": [
					    {"id": "postid1","title": "Post 1 from topic A"},
                        {"id": "postid2", "title": "Post 2 from topic A"}
                    ]
                },
                {
                    "ancestors": [ { "id": "abc123" }, { "id": "efg234" } ],
                    "children": [],
                    "description": "Description of topic B.",
                    "hierarchy": [ "efg234.abc123.zzzblabla" ],
                    "id": "zzzblabla",
					"insertedAt": "2019-05-01T22:44:33.078957Z",
					"updatedAt":"2019-06-18T22:40:16.733422Z",
                    "name": "Topic B",
                    "parent": { "id": "abc123" },
                    "posts": [
                        { "id": "postid3", "title": "Post 3 from topic B" },
                        { "id": "postid4", "title": "Post 4 from topic B" }
                    ]
                }
            ]}}}`
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Topic.List()
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("List returned: %#v\nwant %#v", got, want)
	}
}

func TestTopicService_Get(t *testing.T) {
	insertDate := &DateTime{time.Date(2019, time.May, 1, 22, 44, 33, 78957000, time.UTC)}
	updateDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 16, 733422000, time.UTC)}
	want := &Topic{
		ID: "abc123", Hierarchy: &[]string{"efg234.abc123"}, Ancestors: &[]Topic{{ID: "efg234"}},
		Description: "Description of topic A.", InsertedAt: insertDate, UpdatedAt: updateDate,
		Name: "Topic A", Parent: &Topic{ID: "efg234"}, Children: &[]Topic{{ID: "zzzblabla"}},
		Posts: &[]Post{{ID: "postid1", Title: "Post 1 from topic A"}, {ID: "postid2", Title: "Post 2 from topic A"}},
	}
	expectedResp := `{"data":{"topic":{
		"id": "abc123",
		"hierarchy": ["efg234.abc123"],
		"ancestors": [{"id":"efg234"}],
		"children": [{ "id": "zzzblabla"}],
		"description": "Description of topic A.",
		"insertedAt": "2019-05-01T22:44:33.078957Z",
		"updatedAt":"2019-06-18T22:40:16.733422Z",
		"name": "Topic A",
		"parent": {"id": "efg234"},
		"posts": [
		{"id": "postid1","title": "Post 1 from topic A"},
		{ "id": "postid2", "title": "Post 2 from topic A" }
		]
	}}}`
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Topic.Get(want.ID)
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Get returned: %#v\nwant %#v", got, want)
	}
}

func TestTopicService_Create(t *testing.T) {
	want := &Topic{ID: "abc123", Name: "foo", Description: "bar"}
	expectedResp := fmt.Sprintf(`{"data":{"createTopic":{"name":"%s","id":"%s","description":"%s"}}}`,
		want.Name, want.ID, want.Description)
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Topic.Create(want.Name, want.Description, "dummy_parent")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Create returned: %#v\nwant %#v", got, want)
	}
}

func TestAddToPost(t *testing.T) {
	want := &Topic{ID: "abc123", Name: "foo", Description: "bar"}
	expectedResp := fmt.Sprintf(`{"data":{"addTopicToPost":{"name":"%s","id":"%s","description":"%s"}}}`,
		want.Name, want.ID, want.Description)
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Topic.AddToPost("abc123", "dummypostid")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("AddToPost returned: %#v\nwant %#v", got, want)
	}
}

func TestRemoveFromPost(t *testing.T) {
	want := &Topic{ID: "abc123", Name: "foo", Description: "bar"}
	expectedResp := fmt.Sprintf(`{"data":{"removeTopicFromPost":{"name":"%s","id":"%s","description":"%s"}}}`,
		want.Name, want.ID, want.Description)
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Topic.RemoveFromPost("abc123", "dummypostid")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveFromPost returned: %#v\nwant %#v", got, want)
	}
}

func TestTopicService_AutoGenerate_NoCreate(t *testing.T) {
	want := "def456"
	expectedResp := `{"data":{"organization":{"topics": [
	            { "parent": null, "name": "Company", "id": "abc123", "children": [] },
	            { "parent": null, "name": "Engineering", "id": "bcd234", "children": [{"id": "bcd234", "name": "Services"}] },
	            { "parent": {"id": "bcd234"}, "name": "Services", "id": "cde345", "children": [{"id": "def456", "name": "App 1"}] },
	            { "parent": {"id": "cde345"}, "name": "App 1", "id": "def456", "children": [] }
            ]}}}`
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Topic.AutoGenerate("Engineering/Services/App 1", "/")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if got != want {
		t.Errorf("List returned: %s\nwant %s", got, want)
	}
}

func TestTopicService_AutoGenerate_Create(t *testing.T) {
	// This test is a bit more complex as there are multiple calls and we cannot change the apiEndpoint for now due to upstream not exposing it
	callNbr := 0
	expectedResponses := []string{
		`{"data":{"organization":{"topics": [
	            { "parent": null, "name": "Company", "id": "abc123", "children": [] },
	            { "parent": null, "name": "Engineering", "id": "bcd234", "children": [{"id": "bcd234", "name": "Services"}] },
	            { "parent": {"id": "bcd234"}, "name": "Services", "id": "cde345", "children": [{"id": "def456", "name": "App 1"}] },
	            { "parent": {"id": "cde345"}, "name": "App 1", "id": "def456", "children": [] }
            ]}}}`,
		`{"data":{"createTopic":{"name":"NewTopic","id":"zzzNewTopic","description":""}}}`,
		`{"data":{"createTopic":{"name":"NewSubTopic","id":"zzzNewSub","description":""}}}`,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		_, err = io.WriteString(w, expectedResponses[callNbr])
		assert.NoError(t, err)
		callNbr++
	})

	// srv is the test server that will serve the endpoints
	srv := httptest.NewServer(mux)
	defer srv.Close()
	apiEndpoint = srv.URL

	c := NewClient(&http.Client{}, "dummy_token")

	want := "zzzNewSub"

	got, err := c.Topic.AutoGenerate("Engineering/NewTopic/NewSubTopic", "/")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if got != want {
		t.Errorf("List returned: %s\nwant %s", got, want)
	}
}
