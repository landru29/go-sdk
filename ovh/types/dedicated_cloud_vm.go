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

// Dedicated Cloud Virtual Machine
type DedicatedCloudVm struct {

	// Cluster of the virtual machine.
	ClusterName string `json:"clusterName,omitempty"`

	// Maximum CPU performance, in MHz.
	CpuMax float64 `json:"cpuMax,omitempty"`

	// Number of processors in the virtual machine.
	CpuNum int64 `json:"cpuNum,omitempty"`

	// Time that the virtual machine was ready, but could not get scheduled to run on the physical CPU. In millisecond
	CpuReady float64 `json:"cpuReady,omitempty"`

	// Percentage of time that the virtual machine was ready, but could not get scheduled to run on the physical CPU. In percent
	CpuReadyPercent float64 `json:"cpuReadyPercent,omitempty"`

	// Current CPU performance, in MHz.
	CpuUsed int64 `json:"cpuUsed,omitempty"`

	// List of filers in use by the virtual machine.
	Filers []DedicatedCloudVmFiler `json:"filers,omitempty"`

	// Folder of the virtual machine.
	FolderName string `json:"folderName,omitempty"`

	// Name of the host hosting the virtual machine.
	HostName string `json:"hostName,omitempty"`

	// Memory size of the virtual machine, in MB
	MemoryMax int64 `json:"memoryMax,omitempty"`

	// Amount of guest memory that is shared with other virtual machines, in Mb
	MemoryTps string `json:"memoryTps,omitempty"`

	// Current memory utilization, in MB
	MemoryUsed int64 `json:"memoryUsed,omitempty"`

	// moRef of the virtual machine.
	MoRef string `json:"moRef,omitempty"`

	// Name of the virtual machine.
	Name string `json:"name,omitempty"`

	// Number of packets received.
	NetPacketRx float64 `json:"netPacketRx,omitempty"`

	// Number of packets transmitted.
	NetPacketTx float64 `json:"netPacketTx,omitempty"`

	// Rate at which data is received. In KB/s
	NetRx float64 `json:"netRx,omitempty"`

	// Rate at which data is transmitted. In KB/s
	NetTx float64 `json:"netTx,omitempty"`

	// List of the networks link to this virtual machine
	Networks []DedicatedCloudVmNetwork `json:"networks,omitempty"`

	// Power state of the virtual machine.
	PowerState string `json:"powerState,omitempty"`

	// Amount of time for a read operation from the virtual disk. In millisecond
	ReadLatency float64 `json:"readLatency,omitempty"`

	// Number of read issued per second to the virtual disk. In millisecond
	ReadPerSecond float64 `json:"readPerSecond,omitempty"`

	// Rate of reading data from the virtual disk. In KB/s
	ReadRate float64 `json:"readRate,omitempty"`

	// The index of the current VM in instanceUuids array starting from 1, so 1 means that it is the primary VM.
	RoleFt string `json:"roleFt,omitempty"`

	// Number of snapshot of the virtual machine.
	SnapshotNum int64 `json:"snapshotNum,omitempty"`

	// The fault tolerance state of the virtual machine.
	StateFt string `json:"stateFt,omitempty"`

	// Id of the virtual machine.
	VmId int64 `json:"vmId,omitempty"`

	// Current version status of VMware Tools in the guest operating system.
	VmwareTools string `json:"vmwareTools,omitempty"`

	// Current version of VMware Tools
	VmwareToolsVersion string `json:"vmwareToolsVersion,omitempty"`

	// Amount of time for a write operation from the virtual disk. In millisecond
	WriteLatency float64 `json:"writeLatency,omitempty"`

	// Number of write issued per second to the virtual disk. In millisecond
	WritePerSecond float64 `json:"writePerSecond,omitempty"`

	// Rate of writing data from the virtual disk. In KB/s
	WriteRate float64 `json:"writeRate,omitempty"`
}