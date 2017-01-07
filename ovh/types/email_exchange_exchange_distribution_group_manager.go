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

// EmailExchangeDistributionGroupManager Mailing list managers
type EmailExchangeDistributionGroupManager struct {

	// CreationDate Creation date
	CreationDate *time.Time `json:"creationDate,omitempty"`

	// ManagerAccountID Manager account id
	ManagerAccountID int64 `json:"managerAccountId,omitempty"`

	// ManagerEmailAddress Member account primaryEmailAddress
	ManagerEmailAddress string `json:"managerEmailAddress,omitempty"`

	// TaskPendingID Pending task id
	TaskPendingID int64 `json:"taskPendingId,omitempty"`
}