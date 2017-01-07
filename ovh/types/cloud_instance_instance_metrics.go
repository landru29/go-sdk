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

// CloudInstanceMetrics InstanceMetrics
type CloudInstanceMetrics struct {
	Unit string `json:"unit,omitempty"`

	Values []*CloudInstanceMetricsValue `json:"values,omitempty"`
}