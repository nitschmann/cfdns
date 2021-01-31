package model

import (
	"github.com/go-playground/validator/v10"
)

// ConfigProfile is a single profile item in the config file
type ConfigProfile struct {
	Name        string `mapstructure:"-" toml:"-" validate:"required,lowercase,alphanum"`
	ApiKey      string `mapstructure:"api_key" toml:"api_key" validate:"required"`
	Email       string `mapstructure:"email" toml:"email" validate:"required,email"`
	DefaultZone string `mapstructure:"default_zone" toml:"default_zone" validate:"omitempty,lowercase"`
}

// Validate checks if the ConfigProfile struct is valid
func (p ConfigProfile) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
