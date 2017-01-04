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

// Structure of usage account
type DomainDomainUsageAccountStruct struct {

	// Timestamp
	Date time.Time `json:"date,omitempty"`

	// Number of message in mailbox
	EmailCount int64 `json:"emailCount,omitempty"`

	// Size of mailbox (bytes)
	Quota int64 `json:"quota,omitempty"`
}