package config

import (
	"github.com/BurntSushi/toml"

	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// ProfileListService is the service interface for operations to list profiles from config files
type ProfileListService interface {
	Get() (*map[string]*model.ConfigProfile, error)
}

// ProfileListServiceObj implements the ProfileListService  per default
type ProfileListServiceObj struct {
	configFilepath string
}

// NewProfileListService returns a new pointer instance of ProfileListServiceObj with default values
func NewProfileListService(configFilepath string) *ProfileListServiceObj {
	return &ProfileListServiceObj{configFilepath: configFilepath}
}

// Get the list of available profiles in the given config file
func (serv *ProfileListServiceObj) Get() (map[string]*model.ConfigProfile, error) {
	list := make(map[string]*model.ConfigProfile)
	_, err := toml.DecodeFile(serv.configFilepath, &list)

	return list, err
}
