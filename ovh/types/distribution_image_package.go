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

// An image package description
type DistributionImagePackage struct {

	// Package alias
	Alias string `json:"alias,omitempty"`

	// Package name
	Name string `json:"name,omitempty"`

	// Package version
	Version string `json:"version,omitempty"`
}