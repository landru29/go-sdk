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

// HostingWebModulePost ...
type HostingWebModulePost struct {
	AdminName string `json:"adminName,omitempty"`

	AdminPassword string `json:"adminPassword,omitempty"`

	Dependencies []*HostingWebModuleDependencyType `json:"dependencies,omitempty"`

	Domain string `json:"domain,omitempty"`

	Language string `json:"language,omitempty"`

	ModuleID int64 `json:"moduleId,omitempty"`

	Path string `json:"path,omitempty"`
}
