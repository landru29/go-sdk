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

// DNSsecKey Key type
type DNSsecKey struct {

	// Algorithm Algorithm
	Algorithm int64 `json:"algorithm,omitempty"`

	// Flags Flag of the dnssec key
	Flags int64 `json:"flags,omitempty"`

	// PublicKey Public key
	PublicKey string `json:"publicKey,omitempty"`

	// Tag Key tag
	Tag int64 `json:"tag,omitempty"`
}