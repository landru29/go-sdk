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

// Application volumes
type DockerSlaveFrameworkAppVolume struct {

	// Container path
	ContainerPath int64 `json:"containerPath,omitempty"`

	// Host path
	HostPath int64 `json:"hostPath,omitempty"`

	// Volume mode
	Mode string `json:"mode,omitempty"`
}