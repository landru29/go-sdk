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

// DedicatedServerInstallStartPost ...
type DedicatedServerInstallStartPost struct {
	Details *DedicatedServerInstallCustom `json:"details,omitempty"`

	PartitionSchemeName string `json:"partitionSchemeName,omitempty"`

	TemplateName string `json:"templateName,omitempty"`
}
