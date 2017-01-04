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

// Active Directory organizational unit
type MsServicesActiveDirectoryOrganizationalUnit struct {

	// Flag indicating if passwords should be forced to follow Microsoft's password guidelines
	ComplexityEnabled bool `json:"complexityEnabled,omitempty"`

	// Account lock time (in minutes) when too much passwords have been tried
	LockoutDuration int64 `json:"lockoutDuration,omitempty"`

	// Time (in minutes) before the password attempts counter is reset
	LockoutObservationWindow int64 `json:"lockoutObservationWindow,omitempty"`

	// Maximum number of password tries before account locking
	LockoutThreshold int64 `json:"lockoutThreshold,omitempty"`

	// Maximum lifespan of passwords, in days
	MaxPasswordAge int64 `json:"maxPasswordAge,omitempty"`

	// Minimum lifespan of passwords, in days (0 = unlimited)
	MinPasswordAge int64 `json:"minPasswordAge,omitempty"`

	// Minimum number of characters passwords must contain
	MinPasswordLength int64 `json:"minPasswordLength,omitempty"`

	// Name of the Active Directory organizational unit
	Name string `json:"name,omitempty"`

	// Current state of the Active Directory organizational unit
	State string `json:"state,omitempty"`

	// Task pending id
	TaskPendingId int64 `json:"taskPendingId,omitempty"`
}