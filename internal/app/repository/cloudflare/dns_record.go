package cloudflare

import (
	cloudflareSDK "github.com/cloudflare/cloudflare-go"

	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// DnsRecordRepository is the data interface for the Cloudflare DNS record API
type DnsRecordRepository interface {
	FetchList(zoneID string) ([]model.CloudflareDnsRecord, error)
}

// DnsRecordRepositoryObj implements the DnsRecordRepository interface per default
type DnsRecordRepositoryObj struct {
	Connector *cloudflareSDK.API
}

// NewDnsRecordRepository returns a new pointer instance of DnsRecordRepositoryObj with default values
func NewDnsRecordRepository(config *model.CloudflareConfig) (*DnsRecordRepositoryObj, error) {
	var repository *DnsRecordRepositoryObj

	connector, err := cloudflareSDK.New(config.ApiKey, config.Email)
	if err != nil {
		return repository, err
	}

	repository = &DnsRecordRepositoryObj{Connector: connector}

	return repository, nil
}

// FetchList fetches the list of all DNS records for a zone
func (repo *DnsRecordRepositoryObj) FetchList(zoneID string) ([]model.CloudflareDnsRecord, error) {
	var list []model.CloudflareDnsRecord

	dnsRecords, err := repo.Connector.DNSRecords(zoneID, cloudflareSDK.DNSRecord{})
	if err != nil {
		return list, err
	}

	for _, d := range dnsRecords {
		list = append(list, model.CloudflareDnsRecord{
			ID:         d.ID,
			Type:       d.Type,
			Name:       d.Name,
			Content:    d.Content,
			TTL:        d.TTL,
			Priority:   d.Priority,
			Proxied:    d.Proxied,
			CreatedOn:  d.CreatedOn,
			ModifiedOn: d.ModifiedOn,
		})
	}

	return list, nil
}
