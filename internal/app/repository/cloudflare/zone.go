package cloudflare

import (
	cloudflareSDK "github.com/cloudflare/cloudflare-go"

	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// ZoneRepository is the data interface for the Cloudflare zone API
type ZoneRepository interface {
	FetchList() ([]model.CloudflareZone, error)
}

// ZoneRepositoryObj implements the ZoneRepository per default
type ZoneRepositoryObj struct {
	Connector *cloudflareSDK.API
}

// NewZoneRepository
func NewZoneRepository(config *model.CloudflareConfig) (*ZoneRepositoryObj, error) {
	var obj *ZoneRepositoryObj

	connector, err := cloudflareSDK.New(config.ApiKey, config.Email)
	if err != nil {
		return obj, err
	}

	obj = &ZoneRepositoryObj{Connector: connector}

	return obj, nil
}

// FetchList
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
