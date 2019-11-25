# slab-go

[![GoDoc](https://godoc.org/github.com/VEVO/slab-go/slab?status.svg)](https://godoc.org/github.com/VEVO/slab-go/slab)
[![Test Status](https://github.com/VEVO/slab-go/workflows/tests/badge.svg)](https://github.com/VEVO/slab-go/actions?query=workflow%3Atests)
[![Go Report Card](https://goreportcard.com/badge/github.com/VEVO/slab-go)](https://goreportcard.com/report/github.com/VEVO/slab-go)

slab-go is a Go client library for accessing the [slab.com API](https://the.slab.com/public/slab-api-vk0o0i33).

The slab.com API being in development, this SDK is going to change over time to try and follow upstream changes.

## Usage

The import to use is:

```go
import "github.com/VEVO/slab-go/slab"
```

General rule is to first contruct a client which will give you access to the different parts of the API.

Note that to connect you will need to provide the API access token that the slab support team will provide you.

```go
slabToken := os.Getenv("SLAB_TOKEN")
client := slab.NewClient(&http.Client{Timeout: 10 * time.Second}, slabToken)
o, err := c.Organization.Get()
if err != nil {
    panic(err)
}
```

Usage examples can be found in the [examples](https://github.com/VEVO/slab-go/tree/master/examples) folder of this repository.
