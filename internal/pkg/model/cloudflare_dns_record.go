package model

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// CloudflareDNSRecord describes a Cloudflare DNS record
type CloudflareDNSRecord struct {
	ID         string
	ZoneID     string `validate:"required"`
	Type       string `validate:"required,oneof='A' 'AAAA' 'CNAME' 'HTTPS' 'TXT' 'SRV' 'LOC' 'MX' 'NS' 'SPF' 'CERT' 'DNSKEY' 'DS' 'NAPTR' 'SMIMEA' 'SSHFP' 'SVCB' 'TLSA' 'URI'"`
	Name       string `validate:"required,max=255"`
	Content    string `validate:"required"`
	TTL        int    `validate:"required"`
	Priority   int
	Proxied    bool
	CreatedOn  time.Time
	ModifiedOn time.Time
}

// Validate checks if the CloudflareDNSRecord struct is valid
func (d CloudflareDNSRecord) Validate() error {
	validate := validator.New()
	return validate.Struct(d)
}
