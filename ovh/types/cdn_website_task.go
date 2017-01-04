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

// Task on CDN
type CdnWebsiteTask struct {

	Comment string `json:"comment,omitempty"`

	Function string `json:"function,omitempty"`

	Status string `json:"status,omitempty"`

	TaskId int64 `json:"taskId,omitempty"`
}