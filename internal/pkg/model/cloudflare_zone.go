package model

import (
	"strings"
	"time"
)

// CloudflareZone describes a Cloudflare zone
type CloudflareZone struct {
	ID         string
	Name       string
	Type       string
	Status     string
	CreatedOn  time.Time
	ModifiedOn time.Time
}

// NormalizeDnsRecordName joins a name with the name of the zone.
// Example: test will become test.example.com for z with Name example.com
func (z *CloudflareZone) NormalizeDnsRecordName(name string) string {
	return strings.Join([]string{name, z.Name}, ".")
}
