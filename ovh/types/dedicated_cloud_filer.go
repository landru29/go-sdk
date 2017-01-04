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

// Dedicated Cloud Filer
type DedicatedCloudFiler struct {

	// Billing type of this filer
	BillingType string `json:"billingType,omitempty"`

	// Filer Id
	FilerId int64 `json:"filerId,omitempty"`

	// Human-Readable profile name
	FullProfile string `json:"fullProfile,omitempty"`

	// Filer name
	Name string `json:"name,omitempty"`

	// Commercial profile name
	Profile string `json:"profile,omitempty"`

	Size DedicatedCloudFilerSize `json:"size,omitempty"`

	// Available space of this datastore, in GB
	SpaceFree float64 `json:"spaceFree,omitempty"`

	// Provisionned space of this datastore, in GB
	SpaceProvisionned float64 `json:"spaceProvisionned,omitempty"`

	// Used Space of this filer, in GB
	SpaceUsed float64 `json:"spaceUsed,omitempty"`

	// State of the filer
	State string `json:"state,omitempty"`

	// Number of virtual machine on the filer
	VmTotal int64 `json:"vmTotal,omitempty"`
}