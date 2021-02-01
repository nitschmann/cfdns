package cloudflare

import (
	"errors"
	"fmt"
	"time"

	cloudflareSDK "github.com/cloudflare/cloudflare-go"

	"github.com/nitschmann/cfdns/internal/pkg/customerror"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// DNSRecordRepository is the data interface for the Cloudflare DNS record API
type DNSRecordRepository interface {
	Create(obj model.CloudflareDNSRecord) (model.CloudflareDNSRecord, error)
	Delete(zoneID string, dnsRecord model.CloudflareDNSRecord) (model.CloudflareDNSRecord, error)
	FetchList(zoneID string) ([]model.CloudflareDNSRecord, error)
	Find(zoneID, id string) (model.CloudflareDNSRecord, error)
	FindSingleByNameAndType(zoneID, name, t string) (model.CloudflareDNSRecord, error)
	Update(zoneID string, dnsRecord model.CloudflareDNSRecord) (model.CloudflareDNSRecord, error)
}

// DNSRecordRepositoryObj implements the DNSRecordRepository interface per default
type DNSRecordRepositoryObj struct {
	Connector *cloudflareSDK.API
}

// NewDNSRecordRepository returns a new pointer instance of DNSRecordRepositoryObj with default values
func NewDNSRecordRepository(config *model.CloudflareConfig) (*DNSRecordRepositoryObj, error) {
	var repository *DNSRecordRepositoryObj

	connector, err := cloudflareSDK.New(config.APIKey, config.Email)
	if err != nil {
		return repository, err
	}

	repository = &DNSRecordRepositoryObj{Connector: connector}

	return repository, nil
}

// Create a new DNS record for a zone
func (repo *DNSRecordRepositoryObj) Create(obj model.CloudflareDNSRecord) (model.CloudflareDNSRecord, error) {
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
func (repo *DNSRecordRepositoryObj) Delete(zoneID string, dnsRecord model.CloudflareDNSRecord) (model.CloudflareDNSRecord, error) {
	err := repo.Connector.DeleteDNSRecord(zoneID, dnsRecord.ID)
	return dnsRecord, err
}

// FetchList fetches the list of all DNS records for a zone
func (repo *DNSRecordRepositoryObj) FetchList(zoneID string) ([]model.CloudflareDNSRecord, error) {
	var list []model.CloudflareDNSRecord

	dnsRecords, err := repo.Connector.DNSRecords(zoneID, cloudflareSDK.DNSRecord{})
	if err != nil {
		return list, err
	}

	for _, d := range dnsRecords {
		list = append(list, model.CloudflareDNSRecord{
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
func (repo *DNSRecordRepositoryObj) Find(zoneID, id string) (model.CloudflareDNSRecord, error) {
	var dnsRecord model.CloudflareDNSRecord

	result, err := repo.Connector.DNSRecord(zoneID, id)
	if err != nil {
		return dnsRecord, &customerror.RecordNotFound{
			Type:             "CloudflareDNSRecord",
			IdentifierColumn: "ID",
			Identifier:       id,
			Err:              err,
			PrintOriginalErr: true,
		}
	}

	dnsRecord = model.CloudflareDNSRecord{
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
func (repo *DNSRecordRepositoryObj) FindSingleByNameAndType(zoneID, name, t string) (model.CloudflareDNSRecord, error) {
	var dnsRecord model.CloudflareDNSRecord

	results, err := repo.Connector.DNSRecords(zoneID, cloudflareSDK.DNSRecord{
		Name: name,
		Type: t,
	})
	if err != nil {
		return dnsRecord, err
	}

	if len(results) == 0 {
		return dnsRecord, &customerror.RecordNotFound{
			Type:             "CloudflareDNSRecord",
			IdentifierColumn: "Name",
			Identifier:       name,
		}
	}

	if len(results) > 1 {
		errMsg := fmt.Sprintf("Multiple DNS records matched for name '%s' and type '%s'", name, t)
		return dnsRecord, errors.New(errMsg)
	}

	result := results[0]
	dnsRecord = model.CloudflareDNSRecord{
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

// Update changes attributes of a single DNS record in a given zone
func (repo *DNSRecordRepositoryObj) Update(zoneID string, dnsRecord model.CloudflareDNSRecord) (model.CloudflareDNSRecord, error) {
	err := repo.Connector.UpdateDNSRecord(zoneID, dnsRecord.ID, cloudflareSDK.DNSRecord{
		Type:    dnsRecord.Type,
		Name:    dnsRecord.Name,
		Content: dnsRecord.Content,
		TTL:     dnsRecord.TTL,
		Proxied: dnsRecord.Proxied,
	})

	if err == nil {
		dnsRecord.ModifiedOn = time.Now()
	}

	return dnsRecord, err
}
