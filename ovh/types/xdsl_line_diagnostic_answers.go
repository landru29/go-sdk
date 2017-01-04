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

// Customer answers for line diagnostic
type XdslLineDiagnosticAnswers struct {

	// comment will contains all informations needed for intervention and about your access problem
	Comment string `json:"comment,omitempty"`

	// customer knows that he can be charged (150 euros HT) if he is responsible for the problem or if tests have not been done correctly ?
	ConditionsAccepted bool `json:"conditionsAccepted,omitempty"`

	// customer's phone number
	ContactPhone string `json:"contactPhone,omitempty"`

	// approximative datetime of problem happening
	DatetimeOfAppearance time.Time `json:"datetimeOfAppearance,omitempty"`

	// afternoon closing informations or time for the site
	EndAfternoonHours string `json:"endAfternoonHours,omitempty"`

	// morning closing informations or time for the site
	EndMorningHours string `json:"endMorningHours,omitempty"`

	// indicate if customer wants to be informed by sms
	FollowBySms bool `json:"followBySms,omitempty"`

	// Has modem kept his synchronization during line port reset ?
	HasModemKeptSynchronization bool `json:"hasModemKeptSynchronization,omitempty"`

	// id of appointment chosen (\"possibleValues\" contains choices list with id)
	IdAppointment int64 `json:"idAppointment,omitempty"`

	// is non-professional site ?
	IndividualSite bool `json:"individualSite,omitempty"`

	// Is modem synchronized ? (whatever internet connection)
	ModemIsSynchronized bool `json:"modemIsSynchronized,omitempty"`

	// modem mac address
	ModemMac string `json:"modemMac,omitempty"`

	// Modem still synchronized ? Please check once again.
	ModemStillSynchronized bool `json:"modemStillSynchronized,omitempty"`

	// modem brand and reference
	ModemType string `json:"modemType,omitempty"`

	// is access problem solved ?
	ResolvedAfterTests bool `json:"resolvedAfterTests,omitempty"`

	// is secure site ?
	SecureSite bool `json:"secureSite,omitempty"`

	// Has customer several internal connections ? (on the same place)
	SeveralInternetConnections bool `json:"severalInternetConnections,omitempty"`

	// days or period where site access is not possible
	SiteClosedDays string `json:"siteClosedDays,omitempty"`

	// digicode for site entrance
	SiteDigicode string `json:"siteDigicode,omitempty"`

	// site opening hours or informations
	SiteOpening string `json:"siteOpening,omitempty"`

	// afternoon opening informations or time for the site
	StartAfternoonHours string `json:"startAfternoonHours,omitempty"`

	// morning opening informations or time for the site
	StartMorningHours string `json:"startMorningHours,omitempty"`
}