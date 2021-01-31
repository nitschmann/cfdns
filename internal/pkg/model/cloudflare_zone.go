package model

import "time"

// CloudflareZone describes a Cloudflare zone
type CloudflareZone struct {
	ID         string
	Name       string
	Type       string
	Status     string
	CreatedOn  time.Time
	ModifiedOn time.Time
}
