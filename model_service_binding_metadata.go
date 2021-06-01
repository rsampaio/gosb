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

type ServiceBindingMetadata struct {
	ExpiresAt string `json:"expires_at,omitempty"`
	RenewBefore string `json:"renew_before,omitempty"`
}
