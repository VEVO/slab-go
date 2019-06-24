package slab

import (
	"fmt"
	"reflect"
	"testing"
)

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
		t.Errorf("RemoveFromPost returned: %v, want %v", got, want)
	}

}
