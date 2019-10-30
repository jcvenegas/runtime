/*
 * Cloud Hypervisor API
 *
 * Local HTTP based API for managing and inspecting a cloud-hypervisor virtual machine.
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// MemoryConfig struct for MemoryConfig
type MemoryConfig struct {
	Size int32 `json:"size"`
	File string `json:"file,omitempty"`
}
