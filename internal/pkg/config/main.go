package config

import (
	"os/user"
	"path"
)

// DefaultFilepath returns the default in the system for the config file. It should be used if
// no explicit path to a config file is given.
func DefaultFilepath() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}

	return path.Join(user.HomeDir, ".cfdns"), nil
}
