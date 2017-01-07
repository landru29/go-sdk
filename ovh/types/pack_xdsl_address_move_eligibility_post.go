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

// PackXdslAddressMoveEligibilityPost ...
type PackXdslAddressMoveEligibilityPost struct {
	Address *XdslEligibilityAddress `json:"address,omitempty"`

	LineNumber string `json:"lineNumber,omitempty"`
}