package config

import (
	"bytes"
	"os"

	"github.com/BurntSushi/toml"

	"github.com/nitschmann/cfdns/internal/pkg/model"
)

// UpdateConfigFileProfileContents overwrites the profiles in the given config filepath with the given list of profiles
func UpdateConfigFileProfileContents(configFilepath string, configProfiles map[string]*model.ConfigProfile) error {
	err := os.Truncate(configFilepath, 0)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(configFilepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	err = toml.NewEncoder(buf).Encode(configProfiles)
	if err != nil {
		return err
	}

	_, err = f.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
