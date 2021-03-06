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

// VeeamCloudConnectTask Operation with the Cloud Tenant Account
type VeeamCloudConnectTask struct {

	// EndDate When was this Task done
	EndDate *time.Time `json:"endDate,omitempty"`

	// Name Task name
	Name string `json:"name,omitempty"`

	// Progress Current progress
	Progress int64 `json:"progress,omitempty"`

	// StartDate When the task has been created
	StartDate *time.Time `json:"startDate,omitempty"`

	// State Current Task state
	State string `json:"state,omitempty"`

	TaskID int64 `json:"taskId,omitempty"`
}
