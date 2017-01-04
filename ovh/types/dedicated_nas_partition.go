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

// Storage partition
type DedicatedNasPartition struct {

	// the given name of partition
	PartitionName string `json:"partitionName,omitempty"`

	// must be nfs cifs or both
	Protocol string `json:"protocol,omitempty"`

	// Partition size
	Size int64 `json:"size,omitempty"`
}