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

// Sharepoint account information
type MsServicesSharepointInformation struct {

	// The rights assigned to the sharepoint account
	AccessRights string `json:"accessRights,omitempty"`

	// Active Directory Account id
	ActiveDirectoryAccountId int64 `json:"activeDirectoryAccountId,omitempty"`

	// OneDrive usage in byte
	CurrentUsage int64 `json:"currentUsage,omitempty"`

	// delete at expiration
	DeleteAtExpiration bool `json:"deleteAtExpiration,omitempty"`

	// Sharepoint account id
	Id int64 `json:"id,omitempty"`

	// Sharepoint account license
	License string `json:"license,omitempty"`

	// office license is available
	OfficeLicense bool `json:"officeLicense,omitempty"`

	// OneDrive maximum size in byte
	Quota int64 `json:"quota,omitempty"`

	// Sharepoint account state
	State string `json:"state,omitempty"`

	// Pending task for this sharepoint account
	TaskPendingId int64 `json:"taskPendingId,omitempty"`
}