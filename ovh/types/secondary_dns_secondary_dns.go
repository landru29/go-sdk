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

import (
	"time"
)

// Secondary dns infos
type SecondaryDnsSecondaryDns struct {

	CreationDate time.Time `json:"creationDate,omitempty"`

	// secondary dns server
	Dns string `json:"dns,omitempty"`

	// domain on slave server
	Domain string `json:"domain,omitempty"`

	// master ip
	IpMaster string `json:"ipMaster,omitempty"`
}