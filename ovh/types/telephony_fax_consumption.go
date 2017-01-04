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

// Fax delivery record
type TelephonyFaxConsumption struct {

	Called string `json:"called,omitempty"`

	Calling string `json:"calling,omitempty"`

	ConsumptionId int64 `json:"consumptionId,omitempty"`

	CreationDatetime time.Time `json:"creationDatetime,omitempty"`

	Pages int64 `json:"pages,omitempty"`

	PriceWithoutTax OrderPrice `json:"priceWithoutTax,omitempty"`

	WayType string `json:"wayType,omitempty"`
}