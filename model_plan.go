/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package gosb

type Plan struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Metadata *Metadata `json:"metadata,omitempty"`
	MaintenanceInfo *MaintenanceInfo `json:"maintenance_info,omitempty"`
	Free bool `json:"free,omitempty"`
	Bindable bool `json:"bindable,omitempty"`
	Schemas *Schemas `json:"schemas,omitempty"`
}
