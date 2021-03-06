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

// DedicatedServerHardwareSpecifications A structure describing informations about this dedicated server
type DedicatedServerHardwareSpecifications struct {

	// CoresPerProcessor number of cores per processor
	CoresPerProcessor int64 `json:"coresPerProcessor,omitempty"`

	DefaultHardwareRaidSize *DedicatedServerHardwareSpecificationsDefaultHardwareRaidSize `json:"defaultHardwareRaidSize,omitempty"`

	// DefaultHardwareRaidType Default hardware raid type configured on this server
	DefaultHardwareRaidType string `json:"defaultHardwareRaidType,omitempty"`

	// Description commercial name of this server
	Description string `json:"description,omitempty"`

	// DiskGroups details about the groups of disks in the server
	DiskGroups []*DedicatedServerHardwareSpecificationsDisk `json:"diskGroups,omitempty"`

	MemorySize *DedicatedServerHardwareSpecificationsMemorySize `json:"memorySize,omitempty"`

	// Motherboard server motherboard
	Motherboard string `json:"motherboard,omitempty"`

	// NumberOfProcessors number of processors in this dedicated server
	NumberOfProcessors int64 `json:"numberOfProcessors,omitempty"`

	// ProcessorArchitecture processor architecture bit
	ProcessorArchitecture string `json:"processorArchitecture,omitempty"`

	// ProcessorName processor name
	ProcessorName string `json:"processorName,omitempty"`

	// ThreadsPerProcessor number of threads per processor
	ThreadsPerProcessor int64 `json:"threadsPerProcessor,omitempty"`

	// UsbKeys Capacity of the USB keys installed on your server, if any
	UsbKeys []*ComplexTypeUnitAndValueLong `json:"usbKeys,omitempty"`
}
