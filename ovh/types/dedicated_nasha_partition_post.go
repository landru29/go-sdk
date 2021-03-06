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

// DedicatedNashaPartitionPost ...
type DedicatedNashaPartitionPost struct {
	PartitionName string `json:"partitionName,omitempty"`

	Protocol string `json:"protocol,omitempty"`

	Size int64 `json:"size,omitempty"`
}
