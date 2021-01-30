package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	configPkg "github.com/nitschmann/cfdns/internal/pkg/config"
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

	if p == "" {
		defaultPath, err := configPkg.DefaultFilepath()
		if err != nil {
			return configFilepath, err
		}

		p = defaultPath
	}

	err := serv.createPathDirectoryIfNotExists(p)
	if err != nil {
		return configFilepath, err
	}

	configFilepath = path.Join(p, "config")

	if _, err := os.Stat(configFilepath); err == nil {
		if !force {
			return "", errors.New(fmt.Sprintf("Config file '%s' already exists", configFilepath))
		}

		err := os.Truncate(configFilepath, 0)
		if err != nil {
			return "", err
		}
	}

	fileContent := []byte(serv.fileContent())
	err = ioutil.WriteFile(configFilepath, fileContent, 0644)

	return configFilepath, nil
}

func (serv *CreateServiceObj) createPathDirectoryIfNotExists(p string) error {
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return os.Mkdir(p, os.ModeDir|0755)
	}

	return nil
}

func (serv *CreateServiceObj) fileContent() string {
	contentStr := `
[default]
api_key = %s
email = %s
default_zone = %s`

	content := fmt.Sprintf(contentStr, serv.apiKey, serv.email, serv.zone)
	content = strings.TrimSuffix(content, "\n")
	content = strings.TrimRight(content, " ")

	return content
}
