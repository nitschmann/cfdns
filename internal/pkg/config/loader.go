package config

import (
	"github.com/spf13/viper"
)

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

// AutoFilepath is the automatically detected path of the currently used config file
func AutoFilepath() string {
	return viper.ConfigFileUsed()
}
