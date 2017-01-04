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

// Definition of a single claim notice
type DomainDataClaimNoticeClaimNoticeDecision struct {

	// Classifications where trademark claim notice apply
	Classifications []DomainDataClaimNoticeClassification `json:"classifications,omitempty"`

	// Array of court decisions related to claim notice
	CourtDecisions []DomainDataClaimNoticeCourtDecision `json:"courtDecisions,omitempty"`

	// Goods and services on which apply claim notice
	GoodsAndServices string `json:"goodsAndServices,omitempty"`

	// Name of jurisdiction
	Jurisdiction string `json:"jurisdiction,omitempty"`

	// Jurisdiction country code
	JurisdictionCountryCode string `json:"jurisdictionCountryCode,omitempty"`

	// Mark name implicated in claim notice
	MarkName string `json:"markName,omitempty"`

	// Trademark contacts
	TrademarkContacts []DomainDataClaimNoticeContact `json:"trademarkContacts,omitempty"`

	// Trademark holders
	TrademarkHolders []DomainDataClaimNoticeContact `json:"trademarkHolders,omitempty"`

	// Trademark UDRP informations
	TrademarkUDRP []DomainDataClaimNoticeUdrp `json:"trademarkUDRP,omitempty"`
}