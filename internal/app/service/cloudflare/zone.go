package cloudflare

import (
	cloudflareRepo "github.com/nitschmann/cfdns/internal/app/repository/cloudflare"
	"github.com/nitschmann/cfdns/internal/pkg/customerror"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// ZoneService is the interface to manage Cloudflare zones service logic
type ZoneService interface {
	FindByIdOrName(id string) (model.CloudflareZone, error)
	List() ([]model.CloudflareZone, error)
}

// ZoneServiceObj implements the ZoneService interface per default
type ZoneServiceObj struct {
	Config     *model.CloudflareConfig
	Repository cloudflareRepo.ZoneRepository
}

// NewZoneService returns a new pointer instance of ZoneServiceObj with default values
func NewZoneService(config *model.CloudflareConfig) (*ZoneServiceObj, error) {
	var service *ZoneServiceObj

	err := config.Validate()
	if err != nil {
		return service, err
	}

	repo, err := cloudflareRepo.NewZoneRepository(config)
	if err != nil {
		return service, err
	}

	service = &ZoneServiceObj{
		Config:     config,
		Repository: repo,
	}

	return service, nil
}

// FindByIdOrName tries to find a Cloudflare zone through the API via its ID or name
func (serv *ZoneServiceObj) FindByIdOrName(id string) (model.CloudflareZone, error) {
	var zone model.CloudflareZone

	zone, err := serv.Repository.FindByName(id)
	if err != nil {
		if _, ok := err.(*customerror.RecordNotFound); ok {
			zone, err = serv.Repository.Find(id)
			if err != nil {
				return zone, err
			}
		} else {
			return zone, err
		}
	}

	return zone, nil
}

// List returns a full list of zones
func (serv *ZoneServiceObj) List() ([]model.CloudflareZone, error) {
	list, err := serv.Repository.FetchList()
	return list, err
}
