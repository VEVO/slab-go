package slab

import (
	"reflect"
	"testing"
	"time"
)

func TestUserService_List(t *testing.T) {
	insertDate := &DateTime{time.Date(2019, time.May, 1, 22, 44, 33, 78957000, time.UTC)}
	updateDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 16, 733422000, time.UTC)}
	want := &[]User{
		{ID: "abc123", Name: "Homer S.", Email: "homer@example.com", Type: "user", Title: "Beer belly filler", InsertedAt: insertDate, UpdatedAt: updateDate, Avatar: &Image{}},
		{ID: "efg456", Name: "Wild dude", Email: "wild@example.com", Type: "user", Title: "Bar owner", InsertedAt: insertDate, UpdatedAt: updateDate, Avatar: &Image{}},
	}
	expectedResp := `{"data":{"organization":{"users":[
	{"type":"user","title":"Beer belly filler","name":"Homer S.","insertedAt":"2019-05-01T22:44:33.078957Z","updatedAt":"2019-06-18T22:40:16.733422Z","id":"abc123","email":"homer@example.com","description":"","deactivatedAt":null,"avatar":{"thumb":null,"original":null}},
	{"type":"user","title":"Bar owner","name":"Wild dude","insertedAt":"2019-05-01T22:44:33.078957Z","updatedAt":"2019-06-18T22:40:16.733422Z","id":"efg456","email":"wild@example.com","description":"","deactivatedAt":null,"avatar":{"thumb":null,"original":null}}
	]}}}`

	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.User.List()
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveFromPost returned: %#v\nwant %#v", got, want)
	}
}

func TestUserService_Get(t *testing.T) {
	insertDate := &DateTime{time.Date(2019, time.May, 1, 22, 44, 33, 78957000, time.UTC)}
	updateDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 16, 733422000, time.UTC)}
	want := &User{ID: "abc123", Name: "Homer S.", Email: "homer@example.com", Type: "user", Title: "Beer belly filler", InsertedAt: insertDate, UpdatedAt: updateDate, Avatar: &Image{}}
	expectedResp := `{"data":{"user": {"type":"user","title":"Beer belly filler","name":"Homer S.","insertedAt":"2019-05-01T22:44:33.078957Z","updatedAt":"2019-06-18T22:40:16.733422Z","id":"abc123","email":"homer@example.com","description":"","deactivatedAt":null,"avatar":{"thumb":null,"original":null}} }}`

	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.User.Get("abc123")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveFromPost returned: %#v\nwant %#v", got, want)
	}
}
