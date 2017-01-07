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

// EmailDomainResponderAccount Responder of account
type EmailDomainResponderAccount struct {

	// Account Name of account
	Account string `json:"account,omitempty"`

	// Content Content of responder
	Content string `json:"content,omitempty"`

	// Copy If true, emails will be copy to emailToCopy address
	Copy bool `json:"copy,omitempty"`

	// CopyTo Account where copy emails
	CopyTo string `json:"copyTo,omitempty"`

	// From Date of start responder
	From *time.Time `json:"from,omitempty"`

	// To Date of end responder
	To *time.Time `json:"to,omitempty"`
}