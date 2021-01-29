package checkip

import "github.com/nitschmann/cfd/pkg/util/httpclient"

// Client is the interface for the checkip package
type Client interface {
	GetPublicIpV4() (string, error)
}

// ClientObj implements the Cliet interface per default
type ClientObj struct {
	httpClient httpclient.Client
}

// New returns a new pointer instance of ClientObj
func New() *ClientObj {
	return &ClientObj{httpClient: httpclient.New()}
}
