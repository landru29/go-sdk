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

// Information about the port on the DSLAM
type XdslDslamPort struct {

	// Last time the port lost the synchronization
	LastDesyncDate time.Time `json:"lastDesyncDate,omitempty"`

	// Last time the port synchronized
	LastSyncDate time.Time `json:"lastSyncDate,omitempty"`

	Profile XdslDslamLineProfile `json:"profile,omitempty"`

	Status string `json:"status,omitempty"`
}