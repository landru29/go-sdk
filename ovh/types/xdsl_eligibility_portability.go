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

// XdslEligibilityPortability Eligibility of the portability of the line number
type XdslEligibilityPortability struct {

	// Comments The reason(s) of the negative portability eligibility
	Comments []*XdslEligibilityCodeAndMessage `json:"comments,omitempty"`

	// Eligible Whether or not it is possible to port the line number. If false, commentList contains the reason(s)
	Eligible bool `json:"eligible,omitempty"`

	// UnderCondition Whether or not the portability is possible under condition. If true, warningList contains the reason(s)
	UnderCondition bool `json:"underCondition,omitempty"`

	// Warnings The special condition(s) of the portability
	Warnings []*XdslEligibilityCodeAndMessage `json:"warnings,omitempty"`
}