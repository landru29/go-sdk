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

// Storage task
type DedicatedNasTaskTask struct {

	// information about operation
	Details string `json:"details,omitempty"`

	// the date when the task finished
	DoneDate time.Time `json:"doneDate,omitempty"`

	// last modification of task
	LastUpdate time.Time `json:"lastUpdate,omitempty"`

	// Task type of operation
	Operation string `json:"operation,omitempty"`

	// name of the partition
	PartitionName string `json:"partitionName,omitempty"`

	// The actual state of the task
	Status string `json:"status,omitempty"`

	// the name of your service
	StorageName string `json:"storageName,omitempty"`

	// id of the task
	TaskId int64 `json:"taskId,omitempty"`

	// Insertion of task in the todo
	TodoDate time.Time `json:"todoDate,omitempty"`
}