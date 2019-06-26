// The topics command provides a way to add / remove / list topics
// The token is expected to be located in and environment variable called `SLAB_TOKEN`
//
// Usage examples:
//
// ./topics -action create -name "Test Joseph" -desc "Delete me" -parent "foo1234"
// ID of the newly created topic: bar2345
//
// ./topics -action attach -id "bar2345" -postID "duc3sw1ld"
// Your post is now attached to the topic: 'Test Joseph'
//
// ./topics -action detach -id "bar2345" -postID "duc3sw1ld"
// Your post is now detached from the topic: 'Test Joseph'
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/VEVO/slab-go/slab"
)

var c *slab.Client

func main() {
	slabToken := os.Getenv("SLAB_TOKEN")

	var action = flag.String("action", "list", `actions on the topic:
	* list: lists all the available topics
	* get: shows the details of a specific topic
	* create: creates a topic using the -desc and -name flags
	* attach: attaches the specified topic id from -id to the post specified by -postID
	* detach: removes the specified topic from -id from the post specified by -postID
	`)
	var topicID = flag.String("id", "", `is the topic ID to provide when working on a specific topic`)
	var topicName = flag.String("name", "", "is the name to use when creating a topic")
	var topicDesc = flag.String("desc", "", "is the description to use when creating a topic")
	var parent = flag.String("parent", "", "is the topic ID of the parent to attach the topic during its creation")
	var postID = flag.String("postID", "", "is the ID of the post to attach the topic to or detach it from")
	flag.Parse()

	c = slab.NewClient(&http.Client{Timeout: time.Duration(10 * time.Second)}, slabToken)

	switch *action {
	case "list":
		list()
	case "get":
		get(topicID)
	case "create":
		create(topicName, topicDesc, parent)
	case "attach":
		attach(topicID, postID)
	case "detach":
		detach(topicID, postID)
	default:
		fmt.Printf("Unrecognized action: %s\n", *action)
	}
}

// list shows an example of listing all the available topics
func list() {
	t, err := c.Topic.List()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Topic: %#v\n", t)
}

// get is an example on retrieving the details of a single topic
func get(id *string) {
	t, err := c.Topic.Get(*id)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID: %s\n", t.ID)
	fmt.Printf("Name: %s\n", t.Name)
	fmt.Printf("Description: %s\n", t.Description)
	fmt.Printf("Posts: %#v\n", t.Posts)
	fmt.Printf("Hierarchy: %#v\n", t.Hierarchy)
	fmt.Printf("Parent: %#v\n", t.Parent)
	fmt.Printf("Ancestors: %#v\n", t.Ancestors)
	fmt.Printf("Number of children: %d\n", len(*t.Children))
}

// create shows how to create a new topic
func create(name, desc, parentID *string) {
	t, err := c.Topic.Create(*name, *desc, *parentID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ID of the newly created topic: %s\n", t.ID)
}

// attach adds a topic to a post
func attach(topicID, postID *string) {
	t, err := c.Topic.AddToPost(*topicID, *postID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Your post is now attached to the topic: '%s'\n", t.Name)
}

// detach removes a topic from a post
func detach(topicID, postID *string) {
	t, err := c.Topic.RemoveFromPost(*topicID, *postID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Your post is now detached from the topic: '%s'\n", t.Name)
}
