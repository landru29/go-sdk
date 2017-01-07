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

// DedicatedCloudBackupJob Dedicated Cloud Backup Job
type DedicatedCloudBackupJob struct {

	// AllocatedDisk Disk space allocated to the virtual machine
	AllocatedDisk float64 `json:"allocatedDisk,omitempty"`

	// BackupDays List of days your Virtual Machine will be backuped
	BackupDays []string `json:"backupDays,omitempty"`

	// Encryption Backup is encrypted
	Encryption bool `json:"encryption,omitempty"`

	// OfferType Offer type of the backup job
	OfferType string `json:"offerType,omitempty"`

	// RetentionTime Number of days before the backup is deleted
	RetentionTime int64 `json:"retentionTime,omitempty"`

	// State State of the backup job
	State string `json:"state,omitempty"`

	// VMName Name of the virtual Machine
	VMName string `json:"vmName,omitempty"`
}