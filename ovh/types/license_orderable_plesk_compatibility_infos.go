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

// LicenseOrderablePleskCompatibilityInfos All versions available for Plesk products
type LicenseOrderablePleskCompatibilityInfos struct {
	CanHavePowerPack bool `json:"canHavePowerPack,omitempty"`

	CanHaveResellerManagement bool `json:"canHaveResellerManagement,omitempty"`

	CanHaveWordpressToolkit bool `json:"canHaveWordpressToolkit,omitempty"`

	CompliantAntivirus []string `json:"compliantAntivirus,omitempty"`

	CompliantApplicationSets []string `json:"compliantApplicationSets,omitempty"`

	CompliantDomains []string `json:"compliantDomains,omitempty"`

	CompliantLanguagePack []string `json:"compliantLanguagePack,omitempty"`

	PotentialProblems []string `json:"potentialProblems,omitempty"`

	Version string `json:"version,omitempty"`
}
