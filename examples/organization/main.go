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
	fmt.Printf("Org: %#v\nErr: %s\n", o, err)
}
