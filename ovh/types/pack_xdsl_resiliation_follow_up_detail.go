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

// PackXdslResiliationFollowUpDetail Details about the resiliation
type PackXdslResiliationFollowUpDetail struct {

	// DateTodo Date when the resiliation will take effect
	DateTodo *time.Time `json:"dateTodo,omitempty"`

	// NeedModemReturn If the customer needs to return his modem
	NeedModemReturn bool `json:"needModemReturn,omitempty"`

	// RegistrationDate Date when the resiliation was asked
	RegistrationDate *time.Time `json:"registrationDate,omitempty"`

	// Status Status of the resiliation
	Status string `json:"status,omitempty"`
}