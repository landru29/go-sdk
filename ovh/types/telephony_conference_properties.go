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

// TelephonyConferenceProperties Conference properties
type TelephonyConferenceProperties struct {

	// AnnounceFile Whether or not announce file is played before entrance
	AnnounceFile bool `json:"announceFile,omitempty"`

	// AnonymousRejection Whether or not anonmyous participants are allowed
	AnonymousRejection bool `json:"anonymousRejection,omitempty"`

	// EnterMuted Whether or not participants enter conference room muted
	EnterMuted bool `json:"enterMuted,omitempty"`

	// EventsChannel The events channel hash
	EventsChannel string `json:"eventsChannel,omitempty"`

	// Language The conference sounds language
	Language string `json:"language,omitempty"`

	// Pin The conference pin number
	Pin string `json:"pin,omitempty"`

	// RecordStatus Whether or not conference is recorded
	RecordStatus bool `json:"recordStatus,omitempty"`

	// ReportEmail The email address to send conference report to
	ReportEmail string `json:"reportEmail,omitempty"`

	// ReportStatus The status of the reporting
	ReportStatus string `json:"reportStatus,omitempty"`
}
