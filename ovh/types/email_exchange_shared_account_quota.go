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

// Get shared account quota usage in total available space
type EmailExchangeSharedAccountQuota struct {

	// total amount of space in MB for shared accounts within organization
	QuotaLimit int64 `json:"quotaLimit,omitempty"`

	// space in MB already reserved from the quota limit
	QuotaReserved int64 `json:"quotaReserved,omitempty"`

	// currently used space in KB within all shared accounts
	QuotaUsed int64 `json:"quotaUsed,omitempty"`
}