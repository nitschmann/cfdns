package config

import (
	"fmt"

	"github.com/BurntSushi/toml"

	configPkg "github.com/nitschmann/cfdns/internal/pkg/config"
	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// ProfileAddService is the service interface for operations to add profiles to config files
type ProfileAddService interface {
	AddNewProfile(profile *model.ConfigProfile, forceOverwrite bool) error
}

// ProfileAddServiceObj implements the ProfileAddService interface per default
type ProfileAddServiceObj struct {
	configFilepath string
}

// NewProfileAddService returns a new pointer instance of ProfileAddServiceObj default values
func NewProfileAddService(configFilepath string) *ProfileAddServiceObj {
	return &ProfileAddServiceObj{configFilepath: configFilepath}
}

// AddNewProfile will add a new profile to the specified config file
func (serv *ProfileAddServiceObj) AddNewProfile(profile *model.ConfigProfile, forceOverwrite bool) error {
	err := profile.Validate()
	if err != nil {
		return err
	}

	configProfiles := make(map[string]*model.ConfigProfile)
	_, err = toml.DecodeFile(serv.configFilepath, &configProfiles)
	if err != nil {
		return err
	}

	if _, ok := configProfiles[profile.Name]; ok && !forceOverwrite {
		return fmt.Errorf("Profile '%s' already exists", profile.Name)
	}

	configProfiles[profile.Name] = profile

	err = configPkg.UpdateConfigFileProfileContents(serv.configFilepath, configProfiles)
	if err != nil {
		return err
	}

	return nil
}
