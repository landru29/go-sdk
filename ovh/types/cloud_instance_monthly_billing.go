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

// CloudInstanceMonthlyBilling MonthlyBilling
type CloudInstanceMonthlyBilling struct {

	// Since Monthly billing activated since
	Since *time.Time `json:"since,omitempty"`

	// Status Monthly billing status
	Status string `json:"status,omitempty"`
}
