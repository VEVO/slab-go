package slab

import (
	"reflect"
	"testing"
	"time"
)

func TestOrganizationService_Get(t *testing.T) {
	insertDate := &DateTime{time.Date(2019, time.May, 1, 22, 44, 33, 78957000, time.UTC)}
	updateDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 16, 733422000, time.UTC)}
	want := &Organization{
		ID: "zzzmy0rgzzz", Name: "MyOrg", Host: "myorg.slab.com", InsertedAt: insertDate, UpdatedAt: updateDate,
		Users: &[]User{
			{ID: "abc123", Name: "Homer S."},
			{ID: "efg345", Name: "Marge S."},
		},
		Topics: &[]Topic{
			{
				ID:          "abc123",
				Description: "Description of topic A.",
				Name:        "Topic A",
				Posts:       &[]Post{{ID: "postid1", Title: "Post 1 from topic A"}, {ID: "postid2", Title: "Post 2 from topic A"}},
			},
			{
				ID:          "zzzblabla",
				Description: "Description of topic B.",
				Name:        "Topic B",
				Posts:       &[]Post{{ID: "postid3", Title: "Post 3 from topic B"}, {ID: "postid4", Title: "Post 4 from topic B"}},
			},
		},
	}
	expectedResp := `{"data":{"organization":{
		"id":"zzzmy0rgzzz",
		"name":"MyOrg",
		"host":"myorg.slab.com",
		"insertedAt":"2019-05-01T22:44:33.078957Z",
		"updatedAt":"2019-06-18T22:40:16.733422Z",
		"users":[{"name":"Homer S.","id":"abc123"},{"name":"Marge S.","id":"efg345"}],
		"topics":[
		{
			"posts":[{"id": "postid1","title": "Post 1 from topic A"},{ "id": "postid2", "title": "Post 2 from topic A" } ],
			"name": "Topic A",
			"description": "Description of topic A.",
			"id": "abc123"
		},{
			"description": "Description of topic B.",
			"id": "zzzblabla",
			"name": "Topic B",
			"posts": [{ "id": "postid3", "title": "Post 3 from topic B" },{ "id": "postid4", "title": "Post 4 from topic B" }]
		}
		]
	}}}`

	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Organization.Get()
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Get returned: %#v\nwant %#v", got, want)
	}
}
