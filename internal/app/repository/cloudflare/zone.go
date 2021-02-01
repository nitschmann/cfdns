package cloudflare

import (
	cloudflareSDK "github.com/cloudflare/cloudflare-go"

	"github.com/nitschmann/cfdns/internal/pkg/customerror"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// ZoneRepository is the data interface for the Cloudflare zone API
type ZoneRepository interface {
	FetchList() ([]model.CloudflareZone, error)
	Find(id string) (model.CloudflareZone, error)
	FindByName(name string) (model.CloudflareZone, error)
}

// ZoneRepositoryObj implements the ZoneRepository interface per default
type ZoneRepositoryObj struct {
	Connector *cloudflareSDK.API
}

// NewZoneRepository returns a new pointer instance NewZoneRepositoryObj with default values
func NewZoneRepository(config *model.CloudflareConfig) (*ZoneRepositoryObj, error) {
	var repository *ZoneRepositoryObj

	connector, err := cloudflareSDK.New(config.APIKey, config.Email)
	if err != nil {
		return repository, err
	}

	repository = &ZoneRepositoryObj{Connector: connector}

	return repository, nil
}

// FetchList fetches the list of all zones for the account
func (repo *ZoneRepositoryObj) FetchList() ([]model.CloudflareZone, error) {
	var list []model.CloudflareZone

	zones, err := repo.Connector.ListZones()
	if err != nil {
		return list, err
	}

	for _, z := range zones {
		list = append(list, model.CloudflareZone{
			ID:         z.ID,
			Name:       z.Name,
			Type:       z.Type,
			Status:     z.Status,
			CreatedOn:  z.CreatedOn,
			ModifiedOn: z.ModifiedOn,
		})
	}

	return list, nil
}

// Find fetches the details of a Cloudflare zone via the API by a specific zone ID
func (repo *ZoneRepositoryObj) Find(id string) (model.CloudflareZone, error) {
	var zone model.CloudflareZone

	zoneDetails, err := repo.Connector.ZoneDetails(id)
	if err != nil {
		return zone, &customerror.RecordNotFound{
			Type:             "CloudflareZone",
			IdentifierColumn: "ID",
			Identifier:       id,
		}
	}

	zone = model.CloudflareZone{
		ID:         zoneDetails.ID,
		Name:       zoneDetails.Name,
		Type:       zoneDetails.Type,
		Status:     zoneDetails.Status,
		CreatedOn:  zoneDetails.CreatedOn,
		ModifiedOn: zoneDetails.ModifiedOn,
	}

	return zone, nil
}

// FindByName fetches the details of a Cloudflare zone via the API by a specific zone name
func (repo *ZoneRepositoryObj) FindByName(name string) (model.CloudflareZone, error) {
	var zone model.CloudflareZone

	zoneID, err := repo.Connector.ZoneIDByName(name)
	if err != nil {
		if err.Error() == "Zone could not be found" {
			return zone, &customerror.RecordNotFound{
				Type:             "CloudflareZone",
				IdentifierColumn: "Name",
				Identifier:       name,
			}
		}

		return zone, err
	}

	zone, err = repo.Find(zoneID)
	if err != nil {
		return zone, err
	}

	return zone, nil
}
