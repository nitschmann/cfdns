package model

// Profile is a single profile item in the config file
type Profile struct {
	ApiKey      string `mapstructure:"api_key"`
	Email       string `mapstructure:"email"`
	DefaultZone string `mapstructure:"default_zone"`
}
