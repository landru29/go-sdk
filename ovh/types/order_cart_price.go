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

// OrderCartPrice Price informations with label
type OrderCartPrice struct {

	// Label Label corresponding to a price
	Label string `json:"label,omitempty"`

	Price *OrderPrice `json:"price,omitempty"`
}