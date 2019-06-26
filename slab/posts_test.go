package slab

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestPostService_List(t *testing.T) {
	insertDate := &DateTime{time.Date(2019, time.May, 1, 22, 44, 33, 78957000, time.UTC)}
	updateDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 16, 733422000, time.UTC)}
	publishedDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 15, 733422000, time.UTC)}
	want := &[]Post{
		{ID: "postid1", Title: "Post 1 from topic A", InsertedAt: insertDate, UpdatedAt: updateDate, PublishedAt: publishedDate, Version: 0},
		{ID: "postid2", Title: "Post 2 from topic A", InsertedAt: insertDate, UpdatedAt: updateDate, PublishedAt: publishedDate, Version: 0},
		{ID: "postid3", Title: "Post 3 from topic B", InsertedAt: insertDate, UpdatedAt: updateDate, PublishedAt: publishedDate, Version: 3},
		{ID: "postid4", Title: "Post 4 from topic B", InsertedAt: insertDate, UpdatedAt: updateDate, PublishedAt: publishedDate, Version: 1},
	}
	expectedResp := `{"data":{"organization":{"posts":[
	{"id": "postid1", "title": "Post 1 from topic A", "insertedAt":"2019-05-01T22:44:33.078957Z", "updatedAt":"2019-06-18T22:40:16.733422Z", "publishedAt":"2019-06-18T22:40:15.733422Z","version": 0},
	{"id": "postid2", "title": "Post 2 from topic A", "insertedAt":"2019-05-01T22:44:33.078957Z", "updatedAt":"2019-06-18T22:40:16.733422Z", "publishedAt":"2019-06-18T22:40:15.733422Z","version": 0},
	{"id": "postid3", "title": "Post 3 from topic B", "insertedAt":"2019-05-01T22:44:33.078957Z", "updatedAt":"2019-06-18T22:40:16.733422Z", "publishedAt":"2019-06-18T22:40:15.733422Z","version": 3},
	{"id": "postid4", "title": "Post 4 from topic B", "insertedAt":"2019-05-01T22:44:33.078957Z", "updatedAt":"2019-06-18T22:40:16.733422Z", "publishedAt":"2019-06-18T22:40:15.733422Z","version": 1}
	]}}}`
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Post.List()
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Get returned: %#v\nwant %#v", got, want)
	}
}
func TestPostService_Get(t *testing.T) {
	insertDate := &DateTime{time.Date(2019, time.May, 1, 22, 44, 33, 78957000, time.UTC)}
	updateDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 16, 733422000, time.UTC)}
	publishedDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 15, 733422000, time.UTC)}
	contentOut := `[{\"insert\":\"slab-go\"},{\"attributes\":{\"header\":1},\"insert\":\"\\n\"},{\"insert\":\"slab-go is a Go client library for accessing the \"},{\"attributes\":{\"link\":\"https://the.slab.com/public/slab-api-vk0o0i33\"},\"insert\":\"slab.com API\"},{\"insert\":\".\\nUsage examples can be found in the \"},{\"attributes\":{\"code\":true},\"insert\":\"examples\"},{\"insert\":\" folder of this repository.\\n\"}]`

	unescapeContent := strings.Replace(strings.Replace(contentOut, `\"`, `"`, -1), `\\`, `\`, -1)
	want := &Post{
		ID: "abc123", Title: "slab-go", Version: 0, Content: &unescapeContent,
		InsertedAt: insertDate, UpdatedAt: updateDate, PublishedAt: publishedDate,
	}

	expectedResp := fmt.Sprintf(`{"data":{"post":{
		"id":"abc123",
		"title":"slab-go",
		"insertedAt":"2019-05-01T22:44:33.078957Z",
		"updatedAt":"2019-06-18T22:40:16.733422Z",
		"publishedAt":"2019-06-18T22:40:15.733422Z",
		"version":0,
		"content":"%s"
	}}}`, contentOut)
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Post.Get("abc123")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Get returned: %#v\nwant %#v", got, want)
	}
}

func TestPostService_Create(t *testing.T) {
	want := &Post{ID: "abc123"}
	expectedResp := `{"data":{"createPost":{"id":"abc123"}}}`

	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Post.Create("")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Create returned: %#v\nwant %#v", got, want)
	}
}

func TestPostService_Delete(t *testing.T) {
	want := &Post{ID: "abc123"}
	expectedResp := `{"data":{"deletePost":{"id":"abc123"}}}`

	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Post.Delete("abc123", "")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Delete returned: %#v\nwant %#v", got, want)
	}
}

func TestPostService_Sync(t *testing.T) {
	insertDate := &DateTime{time.Date(2019, time.May, 1, 22, 44, 33, 78957000, time.UTC)}
	updateDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 16, 733422000, time.UTC)}
	publishedDate := &DateTime{time.Date(2019, time.June, 18, 22, 40, 15, 733422000, time.UTC)}
	contentIn := `# slab-go

	slab-go is a Go client library for accessing the [slab.com API](https://the.slab.com/public/slab-api-vk0o0i33).

	Usage examples can be found in the ` + "`examples`" + ` folder of this repository.`
	contentOut := `[{\"insert\":\"slab-go\"},{\"attributes\":{\"header\":1},\"insert\":\"\\n\"},{\"insert\":\"slab-go is a Go client library for accessing the \"},{\"attributes\":{\"link\":\"https://the.slab.com/public/slab-api-vk0o0i33\"},\"insert\":\"slab.com API\"},{\"insert\":\".\\nUsage examples can be found in the \"},{\"attributes\":{\"code\":true},\"insert\":\"examples\"},{\"insert\":\" folder of this repository.\\n\"}]`

	unescapeContent := strings.Replace(strings.Replace(contentOut, `\"`, `"`, -1), `\\`, `\`, -1)
	want := &Post{
		ID: "abc123", Title: "slab-go", Version: 0, Content: &unescapeContent,
		InsertedAt: insertDate, UpdatedAt: updateDate, PublishedAt: publishedDate,
	}

	expectedResp := fmt.Sprintf(`{"data":{"syncPost":{
		"id":"abc123",
		"title":"slab-go",
		"insertedAt":"2019-05-01T22:44:33.078957Z",
		"updatedAt":"2019-06-18T22:40:16.733422Z",
		"publishedAt":"2019-06-18T22:40:15.733422Z",
		"version":0,
		"content":"%s"
	}}}`, contentOut)
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	got, err := c.Post.Sync("dummypostid", contentIn, "https://github.com/VEVO/slab-go/blob/master/README.md", "", "MARKDOWN")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Sync returned: %#v\nwant %#v", got, want)
	}
}

func TestAddTopic(t *testing.T) {
	expectedResp := `{"data":{"addTopicToPost":{"name":"foo","id":"abc123","description":"bar"}}}`
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	err := c.Post.AddTopic("dummypostid", "abc123")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
}

func TestRemoveTopic(t *testing.T) {
	expectedResp := `{"data":{"removeTopicFromPost":{"name":"foo","id":"abc123","description":"bar"}}}`
	c, _, teardown := setup(t, expectedResp)
	defer teardown()

	err := c.Post.RemoveTopic("dummypostid", "abc123")
	if err != nil {
		t.Errorf("Expecting no error, got: %v", err)
	}
}
