package httpclient

import "net/http"

// Client interface
type Client interface {
	Do(req *http.Request) (*http.Response, error)
}

// New returns an intance of *http.Client, which meets the requirements of the Client interface
func New() *http.Client {
	return &http.Client{}
}
