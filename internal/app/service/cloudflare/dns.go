package cloudflare

import (
	"errors"
	"strings"

	cloudflareRepo "github.com/nitschmann/cfdns/internal/app/repository/cloudflare"
	"github.com/nitschmann/cfdns/internal/pkg/model"
	"github.com/nitschmann/cfdns/pkg/checkip"
)

// DNSService is the interface to manage Cloudflare zones service logic
type DNSService interface {
	Create(
		zoneID string,
		t string,
		name string,
		content string,
		ttl int,
		priority int,
		proxied bool,
	) (model.CloudflareDNSRecord, error)
	DeleteByIDOrNameAndType(zone model.CloudflareZone, id string) (model.CloudflareDNSRecord, error)
	FindSingleByIDOrNameAndType(zone model.CloudflareZone, id, t string) (model.CloudflareDNSRecord, error)
	List(zoneID string) ([]model.CloudflareDNSRecord, error)
	UpdateARecordContentToPublicIPV4(zone model.CloudflareZone, id string) (model.CloudflareDNSRecord, error)
}

// DNSServiceObj implements the DNSService interface per default
type DNSServiceObj struct {
	Config     *model.CloudflareConfig
	Repository cloudflareRepo.DNSRecordRepository
}

// NewDNSService returns a new pointer instance of DNSServiceObj with default values
func NewDNSService(config *model.CloudflareConfig) (*DNSServiceObj, error) {
	var service *DNSServiceObj

	err := config.Validate()
	if err != nil {
		return service, err
	}

	repo, err := cloudflareRepo.NewDNSRecordRepository(config)
	if err != nil {
		return service, err
	}

	service = &DNSServiceObj{
		Config:     config,
		Repository: repo,
	}

	return service, nil
}

// Create a new DNS record
func (serv *DNSServiceObj) Create(
	zoneID string,
	t string,
	name string,
	content string,
	ttl int,
	priority int,
	proxied bool,
) (model.CloudflareDNSRecord, error) {
	dnsRecord := model.CloudflareDNSRecord{
		ZoneID:   zoneID,
		Type:     t,
		Name:     name,
		Content:  content,
		TTL:      ttl,
		Priority: priority,
		Proxied:  proxied,
	}

	err := dnsRecord.Validate()
	if err != nil {
		return dnsRecord, err
	}

	dnsRecord, err = serv.Repository.Create(dnsRecord)
	if err != nil {
		return dnsRecord, err
	}

	return dnsRecord, nil
}

// DeleteByIDOrNameAndType deletes a DNS record from a zone and identifies it either by its ID or name + type
func (serv *DNSServiceObj) DeleteByIDOrNameAndType(zone model.CloudflareZone, id, t string) (model.CloudflareDNSRecord, error) {
	var deletedDNSRecord model.CloudflareDNSRecord

	dnsRecord, err := serv.FindSingleByIDOrNameAndType(zone, id, t)
	if err != nil {
		return deletedDNSRecord, err
	}

	deletedDNSRecord, err = serv.Repository.Delete(zone.ID, dnsRecord)
	if err != nil {
		return deletedDNSRecord, err
	}

	return deletedDNSRecord, nil
}

// FindSingleByIDOrNameAndType tries to find a single DNS record for a zone by ID or with its name and type.
func (serv *DNSServiceObj) FindSingleByIDOrNameAndType(zone model.CloudflareZone, id, t string) (model.CloudflareDNSRecord, error) {
	zoneID := zone.ID
	dnsRecord, err := serv.Repository.Find(zoneID, id)
	if err != nil {
		if !strings.Contains(id, "."+zone.Name) {
			id = zone.NormalizeDNSRecordName(id)
		}

		dnsRecord, err = serv.Repository.FindSingleByNameAndType(zoneID, id, t)
		if err != nil {
			return dnsRecord, err
		}
	}

	return dnsRecord, nil
}

// List returns a full list of DNS records for a zone
func (serv *DNSServiceObj) List(zoneID string) ([]model.CloudflareDNSRecord, error) {
	list, err := serv.Repository.FetchList(zoneID)
	return list, err
}

// UpdateARecordContentToPublicIPV4 updates the DNS record content to fetched public IPv4
func (serv *DNSServiceObj) UpdateARecordContentToPublicIPV4(zone model.CloudflareZone, id string) (model.CloudflareDNSRecord, error) {
	var dnsRecord model.CloudflareDNSRecord

	dnsRecord, err := serv.FindSingleByIDOrNameAndType(zone, id, "A")
	if err != nil {
		return dnsRecord, err
	}

	if dnsRecord.Type != "A" {
		return dnsRecord, errors.New("DNS record has to be type A")
	}

	checkipClient := checkip.New()
	publicIPV4, err := checkipClient.GetPublicIPV4()
	if err != nil {
		return dnsRecord, err
	}

	dnsRecord.Content = publicIPV4
	dnsRecord, err = serv.Repository.Update(zone.ID, dnsRecord)
	if err != nil {
		return dnsRecord, err
	}

	return dnsRecord, nil
}
