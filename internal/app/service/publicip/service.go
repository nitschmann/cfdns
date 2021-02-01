package publicip

import (
	"github.com/nitschmann/cfdns/pkg/checkip"
	"github.com/nitschmann/cfdns/pkg/util/httpclient"
)

// Service is the interface for the publicip service package
type Service interface {
	FetchPublicIPV4() (string, error)
}

// ServiceObj implements the Service interface per default
type ServiceObj struct {
	checkIPClient checkip.Client
}

// New returns a new pointer instance of ServiceObj
func New(httpClient httpclient.Client) *ServiceObj {
	return &ServiceObj{
		checkIPClient: checkip.NewWithHTTPClient(httpClient),
	}
}

// FetchPublicIPV4 fetches the public IPv4 address of the machine
func (serv *ServiceObj) FetchPublicIPV4() (string, error) {
	ip, err := serv.checkIPClient.GetPublicIPV4()
	return ip, err
}
