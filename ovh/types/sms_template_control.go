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

// SmsTemplateControl Sms template for moderation (Needed to send in US country)
type SmsTemplateControl struct {

	// Activity Specify the kind of template
	Activity string `json:"activity,omitempty"`

	// Comment Message sent by the moderator
	Comment string `json:"comment,omitempty"`

	// Datetime Template creation datetime
	Datetime *time.Time `json:"datetime,omitempty"`

	// Description Template description
	Description string `json:"description,omitempty"`

	// Message Message pattern to be moderated. Use \"#VALUE#\" format for dynamic text area.
	Message string `json:"message,omitempty"`

	// Name Name of the template
	Name string `json:"name,omitempty"`

	// Status Template status
	Status string `json:"status,omitempty"`
}