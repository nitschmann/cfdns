package cloudflare

import (
	cloudflareRepo "github.com/nitschmann/cfdns/internal/app/repository/cloudflare"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// DnsService is the interface to manage Cloudflare zones service logic
type DnsService interface {
	// FindByIdOrName(id string) (model.CloudflareZone, error)
	List(zoneID string) ([]model.CloudflareDnsRecord, error)
}

// DnsServiceObj implements the DnsService interface per default
type DnsServiceObj struct {
	Config     *model.CloudflareConfig
	Repository cloudflareRepo.DnsRecordRepository
}

// NewDnsService returns a new pointer instance of DnsServiceObj with default values
func NewDnsService(config *model.CloudflareConfig) (*DnsServiceObj, error) {
	var service *DnsServiceObj

	err := config.Validate()
	if err != nil {
		return service, err
	}

	repo, err := cloudflareRepo.NewDnsRecordRepository(config)
	if err != nil {
		return service, err
	}

	service = &DnsServiceObj{
		Config:     config,
		Repository: repo,
	}

	return service, nil
}

// List returns a full list of DNS records for a zone
func (serv *DnsServiceObj) List(zoneID string) ([]model.CloudflareDnsRecord, error) {
	list, err := serv.Repository.FetchList(zoneID)
	return list, err
}
