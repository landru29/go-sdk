/* 
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

import (
	"time"
)

// Domain service
type EmailDomainDomainService struct {

	// List of allowed sizes for this domain in bytes
	AllowedAccountSize []int64 `json:"allowedAccountSize,omitempty"`

	// Creation date of domain
	CreationDate time.Time `json:"creationDate,omitempty"`

	// Name of domain
	Domain string `json:"domain,omitempty"`

	// Filerz of domain
	Filerz int64 `json:"filerz,omitempty"`

	// Name of servicelinked with this domain
	LinkTo string `json:"linkTo,omitempty"`

	// Offer of email service
	Offer string `json:"offer,omitempty"`

	// Domain Status
	Status string `json:"status,omitempty"`
}