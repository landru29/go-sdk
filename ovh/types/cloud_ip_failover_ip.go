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

// CloudIPFailoverIP FailoverIp
type CloudIPFailoverIP struct {

	// Block IP block
	Block string `json:"block,omitempty"`

	// ContinentCode Ip continent
	ContinentCode string `json:"continentCode,omitempty"`

	// Geoloc Ip location
	Geoloc string `json:"geoloc,omitempty"`

	// ID Ip id
	ID string `json:"id,omitempty"`

	// IP Ip
	IP string `json:"ip,omitempty"`

	// Progress Current operation progress in percent
	Progress int64 `json:"progress,omitempty"`

	// RoutedTo Instance where ip is routed to
	RoutedTo string `json:"routedTo,omitempty"`

	// Status Ip status
	Status string `json:"status,omitempty"`

	// SubType IP sub type
	SubType string `json:"subType,omitempty"`
}
