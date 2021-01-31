package cloudflare

import (
	"errors"
	"fmt"

	cloudflareSDK "github.com/cloudflare/cloudflare-go"

	"github.com/nitschmann/cfdns/internal/pkg/customerror"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// DnsRecordRepository is the data interface for the Cloudflare DNS record API
type DnsRecordRepository interface {
	Create(obj model.CloudflareDnsRecord) (model.CloudflareDnsRecord, error)
	Delete(zoneID string, dnsRecord model.CloudflareDnsRecord) (model.CloudflareDnsRecord, error)
	FetchList(zoneID string) ([]model.CloudflareDnsRecord, error)
	Find(zoneID, id string) (model.CloudflareDnsRecord, error)
	FindSingleByNameAndType(zoneID, name, t string) (model.CloudflareDnsRecord, error)
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

// Create a new DNS record for a zone
func (repo *DnsRecordRepositoryObj) Create(obj model.CloudflareDnsRecord) (model.CloudflareDnsRecord, error) {
	response, err := repo.Connector.CreateDNSRecord(obj.ZoneID, cloudflareSDK.DNSRecord{
		Type:     obj.Type,
		Name:     obj.Name,
		Content:  obj.Content,
		TTL:      obj.TTL,
		Priority: obj.Priority,
		Proxied:  obj.Proxied,
	})
	if err != nil {
		return obj, err
	}

	obj.ID = response.Result.ID
	obj.CreatedOn = response.Result.CreatedOn
	obj.ModifiedOn = response.Result.ModifiedOn

	return obj, nil
}

// Delete a DNS record with given ID from a zone
func (repo *DnsRecordRepositoryObj) Delete(zoneID string, dnsRecord model.CloudflareDnsRecord) (model.CloudflareDnsRecord, error) {
	err := repo.Connector.DeleteDNSRecord(zoneID, dnsRecord.ID)
	return dnsRecord, err
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
			ZoneID:     d.ZoneID,
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

// Find a single DNS record for a given zone ID and record ID
func (repo *DnsRecordRepositoryObj) Find(zoneID, id string) (model.CloudflareDnsRecord, error) {
	var dnsRecord model.CloudflareDnsRecord

	result, err := repo.Connector.DNSRecord(zoneID, id)
	if err != nil {
		return dnsRecord, &customerror.RecordNotFound{
			Type:             "CloudflareDnsRecord",
			IdentifierColumn: "ID",
			Identifier:       id,
			Err:              err,
			PrintOriginalErr: true,
		}
	}

	dnsRecord = model.CloudflareDnsRecord{
		ID:         result.ID,
		ZoneID:     result.ZoneID,
		Type:       result.Type,
		Name:       result.Name,
		Content:    result.Content,
		TTL:        result.TTL,
		Priority:   result.Priority,
		Proxied:    result.Proxied,
		CreatedOn:  result.CreatedOn,
		ModifiedOn: result.ModifiedOn,
	}

	return dnsRecord, nil
}

// FindSingleByNameAndType tries to find a single DNS record for a zone with its name and type.
// It will return an error in case multiple records match with the conditions
func (repo *DnsRecordRepositoryObj) FindSingleByNameAndType(zoneID, name, t string) (model.CloudflareDnsRecord, error) {
	var dnsRecord model.CloudflareDnsRecord

	results, err := repo.Connector.DNSRecords(zoneID, cloudflareSDK.DNSRecord{
		Name: name,
		Type: t,
	})
	if err != nil {
		return dnsRecord, err
	}

	if len(results) == 0 {
		return dnsRecord, &customerror.RecordNotFound{
			Type:             "CloudflareDnsRecord",
			IdentifierColumn: "Name",
			Identifier:       name,
		}
	}

	if len(results) > 1 {
		errMsg := fmt.Sprintf("Multiple DNS records matched for name '%s' and type '%s'", name, t)
		return dnsRecord, errors.New(errMsg)
	}

	result := results[0]
	dnsRecord = model.CloudflareDnsRecord{
		ID:         result.ID,
		ZoneID:     result.ZoneID,
		Type:       result.Type,
		Name:       result.Name,
		Content:    result.Content,
		TTL:        result.TTL,
		Priority:   result.Priority,
		Proxied:    result.Proxied,
		CreatedOn:  result.CreatedOn,
		ModifiedOn: result.ModifiedOn,
	}

	return dnsRecord, nil
}
