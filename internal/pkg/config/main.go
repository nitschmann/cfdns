package config

import (
	"os/user"
	"path"

	"github.com/nitschmann/cfdns/internal/pkg/model"
)

var profiles *map[string]*model.ConfigProfile

// DefaultFilepath returns the default in the system for the config file. It should be used if
// no explicit path to a config file is given.
func DefaultFilepath() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return path.Join(user.HomeDir, ".cfdns"), nil
}

// GetProfiles returns a map list of currently loaded and available profiles
func GetProfiles() *map[string]*model.ConfigProfile {
	return profiles
}
