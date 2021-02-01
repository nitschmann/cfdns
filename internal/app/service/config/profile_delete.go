package config

import (
	"errors"
	"fmt"

	"github.com/BurntSushi/toml"

	configPkg "github.com/nitschmann/cfdns/internal/pkg/config"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// ProfileDeleteService is the service interface for operations to delete profiles from config files
type ProfileDeleteService interface {
	DeleteProfile(name string) error
}

// ProfileDeleteServiceObj implements the ProfileDeleteService interface per default
type ProfileDeleteServiceObj struct {
	configFilepath string
}

// NewProfileDeleteService returns a new pointer instance of ProfileDeleteServiceObj default values
func NewProfileDeleteService(configFilepath string) *ProfileDeleteServiceObj {
	return &ProfileDeleteServiceObj{configFilepath: configFilepath}
}

// DeleteProfile will delete a config profile with the given name from the config file if present
func (serv *ProfileDeleteServiceObj) DeleteProfile(name string) error {
	configProfiles := make(map[string]*model.ConfigProfile)
	_, err := toml.DecodeFile(serv.configFilepath, &configProfiles)
	if err != nil {
		return err
	}

	if _, ok := configProfiles[name]; !ok {
		return errors.New(fmt.Sprintf("No profile '%s' available", name))
	}

	delete(configProfiles, name)

	err = configPkg.UpdateConfigFileProfileContents(serv.configFilepath, configProfiles)
	if err != nil {
		return err
	}

	return nil
}
