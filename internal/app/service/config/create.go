package config

import (
	"bytes"
	"fmt"
	"os"
	"path"

	"github.com/BurntSushi/toml"

	configPkg "github.com/nitschmann/cfdns/internal/pkg/config"
	"github.com/nitschmann/cfdns/internal/pkg/model"
	"github.com/nitschmann/cfdns/internal/pkg/validator"
)

// CreateService is the service interface to create config files
type CreateService interface {
	Create(p string, force bool) (string, error)
}

// CreateServiceObj implements the CreateService interface per default
type CreateServiceObj struct {
	apiKey string
	email  string
	zone   string
}

// NewCreateService returns an pointer instance of CreateServiceObj with default values
func NewCreateService(apiKey, email, zone string) *CreateServiceObj {
	return &CreateServiceObj{
		apiKey: apiKey,
		email:  email,
		zone:   zone,
	}
}

// Create takes care to create a new config file with a default profile under the given path
func (serv *CreateServiceObj) Create(p string, force bool) (string, error) {
	var configFilepath string

	configProfile, err := serv.initAndValidateProfileObj()
	if err != nil {
		return configFilepath, err
	}

	if p == "" {
		defaultPath, err := configPkg.DefaultFilepath()
		if err != nil {
			return configFilepath, err
		}

		p = defaultPath
	}

	err = serv.createPathDirectoryIfNotExists(p)
	if err != nil {
		return configFilepath, err
	}

	configFilepath = path.Join(p, "config")
	var f *os.File

	if _, err := os.Stat(configFilepath); err == nil {
		if !force {
			return "", fmt.Errorf("Config file '%s' already exists", configFilepath)
		}

		err := os.Truncate(configFilepath, 0)
		if err != nil {
			return "", err
		}

		f, err = os.OpenFile(configFilepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			return "", err
		}
	} else {
		f, err = os.Create(configFilepath)
		if err != nil {
			return "", err
		}
	}
	defer f.Close()

	err = serv.writeConfigProfileToFile(f, configProfile)
	if err != nil {
		return "", err
	}

	return configFilepath, nil
}

func (serv *CreateServiceObj) createPathDirectoryIfNotExists(p string) error {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return os.Mkdir(p, os.ModeDir|0755)
	}

	return nil
}

func (serv *CreateServiceObj) initAndValidateProfileObj() (*model.ConfigProfile, error) {
	profile := &model.ConfigProfile{
		Name:   "default",
		APIKey: serv.apiKey,
		Email:  serv.email,
	}

	validator := validator.NewModelValidator(profile)
	err := validator.Validate()

	return profile, err
}

func (serv *CreateServiceObj) writeConfigProfileToFile(f *os.File, configProfile *model.ConfigProfile) error {
	list := make(map[string]*model.ConfigProfile)
	list[configProfile.Name] = configProfile

	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(list)
	if err != nil {
		return err
	}

	_, err = f.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
