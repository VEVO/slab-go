package slab

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func setup(t *testing.T, response string) (c *Client, mux *http.ServeMux, teardown func()) {
	// mux is the test server router
	mux = http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		_, err := ioutil.ReadAll(r.Body)
		assert.NoError(t, err)
		_, err = io.WriteString(w, response)
		assert.NoError(t, err)
	})

	// srv is the test server that will serve the endpoints
	srv := httptest.NewServer(mux)
	apiEndpoint = srv.URL

	c = NewClient(&http.Client{}, "dummy_token")

	return c, mux, srv.Close
}
