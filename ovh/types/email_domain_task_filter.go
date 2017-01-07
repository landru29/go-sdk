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

// EmailDomainTaskFilter Task filter List
type EmailDomainTaskFilter struct {

	// Account Account name of task
	Account string `json:"account,omitempty"`

	// Action Action of task
	Action string `json:"action,omitempty"`

	// Domain Domain name of task
	Domain string `json:"domain,omitempty"`

	// ID Id of task
	ID int64 `json:"id,omitempty"`

	// Timestamp Creation date of task
	Timestamp *time.Time `json:"timestamp,omitempty"`
}