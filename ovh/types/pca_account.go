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

// Cloud Archives Account
type PcaAccount struct {

	// account domain
	Domain string `json:"domain,omitempty"`

	// host where to upload files to
	Host string `json:"host,omitempty"`

	// login for cloud archives
	Login string `json:"login,omitempty"`

	// encrypted password
	Password string `json:"password,omitempty"`

	// ssh key to be used for upload and download
	Sshkey string `json:"sshkey,omitempty"`
}