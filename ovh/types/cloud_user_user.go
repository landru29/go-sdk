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

// CloudUser User
type CloudUser struct {

	// CreationDate User creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// Description User description
	Description string `json:"description,omitempty"`

	// ID User id
	ID int64 `json:"id,omitempty"`

	// Status User status
	Status string `json:"status,omitempty"`

	// Username Username
	Username string `json:"username,omitempty"`
}
