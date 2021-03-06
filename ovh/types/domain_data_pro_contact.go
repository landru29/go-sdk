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

// DomainDataProContact Representation of an .pro Contact Resource
type DomainDataProContact struct {

	// Authority Authority that certify your profesional status
	Authority string `json:"authority,omitempty"`

	// AuthorityWebsite Website of the authority that certify your profesional status
	AuthorityWebsite string `json:"authorityWebsite,omitempty"`

	// ID .pro Contact ID
	ID int64 `json:"id,omitempty"`

	// JobDescription Description of your job
	JobDescription string `json:"jobDescription,omitempty"`

	// LicenseNumber License number given by the authority
	LicenseNumber string `json:"licenseNumber,omitempty"`
}
