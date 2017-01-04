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

type ClusterHadoopUserPost struct {

	ClouderaManager bool `json:"clouderaManager,omitempty"`

	HttpFrontend bool `json:"httpFrontend,omitempty"`

	Hue bool `json:"hue,omitempty"`

	Password string `json:"password,omitempty"`

	Username string `json:"username,omitempty"`
}