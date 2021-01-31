package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// CloudflareZone describes a Cloudflare DNS record
type CloudflareDnsRecord struct {
	ID         string
	Type       string `validate:"required"`
	Name       string `validate:"required,max=255"`
	Content    string `validate:"required"`
	TTL        int    `validate:"required"`
	Priority   int
	Proxied    bool
	CreatedOn  time.Time
	ModifiedOn time.Time
}

// Validate checks if the CloudflareZone struct is valid
func (d CloudflareDnsRecord) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
