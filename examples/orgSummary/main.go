// The orgSummary command shows how to retrieve the details of your organization
// from the slab.com API.
// The token is expected to be located in and environment variable called `SLAB_TOKEN`
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aerostitch/slab-go/slab"
)

func main() {
	slabToken := os.Getenv("SLAB_TOKEN")
	c := slab.NewClient(&http.Client{Timeout: time.Duration(10 * time.Second)}, slabToken)
	o, err := c.Organization.Get()
	if err != nil {
		panic(err)
	}
	fmt.Printf("My organization's name is %s and its id is: %s\nTo connect to slab I go to: https://%s\n", o.Name, o.ID, o.Host)
	fmt.Printf("There are %d posts and %d topics available in my organization.\n", len(*o.Posts), len(*o.Topics))
	fmt.Printf("Currently %d users are attached to my organization.\n", len(*o.Users))
}
