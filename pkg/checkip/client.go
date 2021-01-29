package checkip

import "github.com/nitschmann/cfdns/pkg/util/httpclient"

// Client is the interface for the checkip package
type Client interface {
	GetPublicIpV4() (string, error)
}

// ClientObj implements the Cliet interface per default
type ClientObj struct {
	httpClient httpclient.Client
}

// New returns a new pointer instance of ClientObj with default values
func New() *ClientObj {
	return &ClientObj{httpClient: httpclient.New()}
}

// NewWithHttpClient returns a new pointer instance of ClientObj with a custom http.Client
func NewWithHttpClient(httpClient httpclient.Client) *ClientObj {
	return &ClientObj{httpClient: httpClient}
}
