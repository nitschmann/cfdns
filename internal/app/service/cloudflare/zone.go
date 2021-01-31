package cloudflare

import (
	cloudflareRepo "github.com/nitschmann/cfdns/internal/app/repository/cloudflare"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// ZoneService is the interface to manage Cloudflare zones service logic
type ZoneService interface {
	List() ([]model.CloudflareZone, error)
}

// ZoneServiceObj implements the ZoneService interface per default
type ZoneServiceObj struct {
	Config     *model.CloudflareConfig
	Repository cloudflareRepo.ZoneRepository
}

// NewZoneService returns a new pointer instance of ZoneServiceObj with default values
func NewZoneService(config *model.CloudflareConfig) (*ZoneServiceObj, error) {
	var obj *ZoneServiceObj

	err := config.Validate()
	if err != nil {
		return obj, err
	}

	repo, err := cloudflareRepo.NewZoneRepository(config)
	if err != nil {
		return obj, err
	}

	obj = &ZoneServiceObj{
		Config:     config,
		Repository: repo,
	}

	return obj, nil
}

// List returns a full list of zones
func (serv *ZoneServiceObj) List() ([]model.CloudflareZone, error) {
	list, err := serv.Repository.FetchList()
	return list, err
}
