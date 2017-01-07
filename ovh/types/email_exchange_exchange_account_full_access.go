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

// EmailExchangeAccountFullAccess Users having full access on this mailbox
type EmailExchangeAccountFullAccess struct {

	// AllowedAccountID Account id to give full access
	AllowedAccountID int64 `json:"allowedAccountId,omitempty"`

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// TaskPendingID Pending task id
	TaskPendingID int64 `json:"taskPendingId,omitempty"`
}