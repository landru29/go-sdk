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

// Scheduled event
type TelephonySchedulerEvent struct {

	// The category of the event
	Categories string `json:"categories,omitempty"`

	// The ending date of the event
	DateEnd time.Time `json:"dateEnd,omitempty"`

	// The beginning date of the event
	DateStart time.Time `json:"dateStart,omitempty"`

	Description string `json:"description,omitempty"`

	Title string `json:"title,omitempty"`

	// The unique ICS event identifier
	Uid string `json:"uid,omitempty"`
}