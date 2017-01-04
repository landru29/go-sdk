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

// Server partitions informations
type DedicatedServerRtmPartition struct {

	InodeUsage DedicatedServerRtmPartitionInodeUsage `json:"inodeUsage,omitempty"`

	// Partition mount point
	MountPoint string `json:"mountPoint,omitempty"`

	// Partition
	Partition string `json:"partition,omitempty"`

	Usage DedicatedServerRtmPartitionUsage `json:"usage,omitempty"`
}