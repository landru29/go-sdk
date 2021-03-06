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

// IPFirewallNetworkRule Rule on ip
type IPFirewallNetworkRule struct {

	// Action Action on this rule
	Action string `json:"action,omitempty"`

	CreationDate *time.Time `json:"creationDate,omitempty"`

	// Destination Destination ip for your rule
	Destination string `json:"destination,omitempty"`

	// DestinationPort Destination port range for your rule. Only with TCP/UDP protocol
	DestinationPort string `json:"destinationPort,omitempty"`

	// Fragments Fragments option
	Fragments bool `json:"fragments,omitempty"`

	// Protocol Network protocol
	Protocol string `json:"protocol,omitempty"`

	Rule string `json:"rule,omitempty"`

	Sequence int64 `json:"sequence,omitempty"`

	// Source Source ip for your rule
	Source string `json:"source,omitempty"`

	// SourcePort Source port range for your rule. Only with TCP/UDP protocol
	SourcePort string `json:"sourcePort,omitempty"`

	// State Current state of your rule
	State string `json:"state,omitempty"`

	// TCPOption TCP option on your rule
	TCPOption string `json:"tcpOption,omitempty"`
}
