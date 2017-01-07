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

// TelephonyLineNotificationsLogsOptions Error logging notifications options
type TelephonyLineNotificationsLogsOptions struct {

	// Email Email address where to send notifications
	Email string `json:"email,omitempty"`

	// Frequency Frequency at which the notifications will be send
	Frequency string `json:"frequency,omitempty"`

	// SendIfNull Send a blank notification if there is no diagnosticReports entries for the period
	SendIfNull bool `json:"sendIfNull,omitempty"`
}