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

// CloudQuotaInstanceQuotas InstanceQuotas
type CloudQuotaInstanceQuotas struct {

	// MaxCores Maximum total cores allowed in your project
	MaxCores int64 `json:"maxCores,omitempty"`

	// MaxInstances Maximum total cores allowed in your project
	MaxInstances int64 `json:"maxInstances,omitempty"`

	MaxRAM int64 `json:"maxRam,omitempty"`

	// UsedCores Current used cores number
	UsedCores int64 `json:"usedCores,omitempty"`

	// UsedInstances Current used instances
	UsedInstances int64 `json:"usedInstances,omitempty"`

	// UsedRAM Current used ram
	UsedRAM int64 `json:"usedRAM,omitempty"`
}
