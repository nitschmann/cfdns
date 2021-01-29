package publicip

import (
	"github.com/nitschmann/cfdns/pkg/checkip"
	"github.com/nitschmann/cfdns/pkg/util/httpclient"
)

// Service is the interface for the publicip service package
type Service interface {
	FetchPublicIpV4() (string, error)
}

// ServiceObj implements the Service interface per default
type ServiceObj struct {
	checkIpClient checkip.Client
}

// New returns a new pointer instance of ServiceObj
func New(httpClient httpclient.Client) *ServiceObj {
	return &ServiceObj{
		checkIpClient: checkip.NewWithHttpClient(httpClient),
	}
}

// FetchPublicIpV4 fetches the public IPv4 address of the machine
func (serv *ServiceObj) FetchPublicIpV4() (string, error) {
	ip, err := serv.checkIpClient.GetPublicIpV4()
	return ip, err
}
