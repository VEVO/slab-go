// The postFromMarkdown command shows how to use the syncPost command to upload externally-managed
// markdown posts to slab.
// The token is expected to be located in and environment variable called `SLAB_TOKEN`
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/aerostitch/slab-go/slab"
)

func main() {
	slabToken := os.Getenv("SLAB_TOKEN")
	c := slab.NewClient(&http.Client{Timeout: time.Duration(10 * time.Second)}, slabToken)

	// Pulling some content from a url just to have an example
	resp, err := http.Get("https://raw.githubusercontent.com/aerostitch/slab-go/master/README.md")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	p, err := c.Post.Sync("slabgoREADME", string(body), "https://github.com/aerostitch/slab-go/blob/master/README.md", "https://github.com/aerostitch/slab-go/blob/master/README.md", "MARKDOWN")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Post id is: %s\nPost title is: %s\nPost content is:\n%s\nPost version: %d\n", p.ID, p.Title, *p.Content, *p.Version)
}
