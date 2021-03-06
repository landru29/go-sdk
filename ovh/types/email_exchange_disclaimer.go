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

// EmailExchangeDisclaimer Exchange organization disclaimer
type EmailExchangeDisclaimer struct {

	// Content Signature, added at the bottom of your organization emails
	Content string `json:"content,omitempty"`

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// Name Disclaimer name
	Name string `json:"name,omitempty"`

	// OutsideOnly Activate the disclaimer only for external emails
	OutsideOnly bool `json:"outsideOnly,omitempty"`

	// TaskPendingID task pending id
	TaskPendingID int64 `json:"taskPendingId,omitempty"`
}
