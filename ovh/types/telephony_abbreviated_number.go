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

// TelephonyAbbreviatedNumber Abbreviated number
type TelephonyAbbreviatedNumber struct {

	// AbbreviatedNumber The abbreviated number which must start with \"2\" and must have a length of 3 or 4 digits
	AbbreviatedNumber int64 `json:"abbreviatedNumber,omitempty"`

	// DestinationNumber The destination of the abbreviated number
	DestinationNumber string `json:"destinationNumber,omitempty"`

	Name string `json:"name,omitempty"`

	Surname string `json:"surname,omitempty"`
}
