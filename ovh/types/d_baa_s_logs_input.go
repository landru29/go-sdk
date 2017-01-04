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

// Input
type DBaaSLogsInput struct {

	// Input creation
	CreatedAt time.Time `json:"createdAt,omitempty"`

	// Input description
	Description string `json:"description,omitempty"`

	// Input engine UUID
	EngineId string `json:"engineId,omitempty"`

	// Port
	ExposedPort string `json:"exposedPort,omitempty"`

	// Hostname
	Hostname string `json:"hostname,omitempty"`

	// Input ID
	InputId string `json:"inputId,omitempty"`

	// Associated DBaaS Logs option
	OptionId string `json:"optionId,omitempty"`

	// Input IP address
	PublicAddress string `json:"publicAddress,omitempty"`

	// Indicate if input need to be restarted
	RestartRequired bool `json:"restartRequired,omitempty"`

	// Force only one instance
	SingleInstance bool `json:"singleInstance,omitempty"`

	// Input SSL certificate
	SslCertificate string `json:"sslCertificate,omitempty"`

	// init: configuration required, pending: ready to start, running: available
	Status string `json:"status,omitempty"`

	// Associated Graylog stream
	StreamId string `json:"streamId,omitempty"`

	// Input title
	Title string `json:"title,omitempty"`

	// Input last update
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}