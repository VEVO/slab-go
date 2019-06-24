// The topics command provides a way to add / remove / list topics
// The token is expected to be located in and environment variable called `SLAB_TOKEN`
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aerostitch/slab-go/slab"
)

var c *slab.Client

func main() {
	slabToken := os.Getenv("SLAB_TOKEN")

	var action = flag.String("action", "list", `actions on the topic:
	* list: lists all the available topics
	* get: shows the details of a specific topic
	`)
	var topicID = flag.String("id", "", `is the topic ID to provide when working on a specific topic`)
	flag.Parse()

	c = slab.NewClient(&http.Client{Timeout: time.Duration(10 * time.Second)}, slabToken)

	switch *action {
	case "list":
		list()
	case "get":
		get(*topicID)
	default:
		fmt.Printf("Unrecognized action: %s\n", *action)
	}
}

func list() {
	t, err := c.Topic.List()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Topic: %#v\n", t)
}

func get(id string) {
	t, err := c.Topic.Get(id)
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
