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

// VpsVeeamRestorePointsRestorePost ...
type VpsVeeamRestorePointsRestorePost struct {
	ChangePassword bool `json:"changePassword,omitempty"`

	Export string `json:"export,omitempty"`

	Full bool `json:"full,omitempty"`
}
