package config

import (
	"github.com/nitschmann/cfdns/internal/pkg/model"
	"github.com/spf13/viper"
)

// AutoFilepath is the automatically detected path of the currently used config file
func AutoFilepath() string {
	return viper.ConfigFileUsed()
}

// SetUpLoader sets up the viper config framework settings
func SetUpLoader() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AddConfigPath("$HOME/.cfdns")
	viper.AddConfigPath("/etc/cfdns")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		} else {
			return err
		}
	}

	return nil
}

func Load() error {
	profiles = new(map[string]*model.ConfigProfile)
	err := viper.Unmarshal(profiles)
	if err != nil {
		return err
	}

	return nil
}
