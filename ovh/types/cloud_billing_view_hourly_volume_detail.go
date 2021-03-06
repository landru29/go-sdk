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

// CloudBillingViewHourlyVolumeDetail HourlyVolumeDetail
type CloudBillingViewHourlyVolumeDetail struct {
	Quantity *CloudBillingViewQuantity `json:"quantity,omitempty"`

	// TotalPrice Total price
	TotalPrice float64 `json:"totalPrice,omitempty"`

	// VolumeID Volume ID
	VolumeID string `json:"volumeId,omitempty"`
}
