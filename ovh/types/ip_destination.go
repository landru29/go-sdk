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

// IPDestination A structure given service and its nexthops as a destination for failover ips
type IPDestination struct {

	// Nexthop Nexthops available on this service
	Nexthop []string `json:"nexthop,omitempty"`

	// Service Service destination
	Service string `json:"service,omitempty"`
}