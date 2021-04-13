/*
 * Open Service Broker API
 *
 * The Open Service Broker API defines an HTTP(S) interface between Platforms and Service Brokers.
 *
 * API version: master - might contain changes that are not yet released
 * Contact: open-service-broker-api@googlegroups.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gosb

import (
	"encoding/json"
)

// ServiceBindingResource struct for ServiceBindingResource
type ServiceBindingResource struct {
	Metadata *ServiceBindingMetadata `json:"metadata,omitempty"`
	Credentials *map[string]interface{} `json:"credentials,omitempty"`
	SyslogDrainUrl *string `json:"syslog_drain_url,omitempty"`
	RouteServiceUrl *string `json:"route_service_url,omitempty"`
	VolumeMounts *[]ServiceBindingVolumeMount `json:"volume_mounts,omitempty"`
	Endpoints *[]ServiceBindingEndpoint `json:"endpoints,omitempty"`
	Parameters *map[string]interface{} `json:"parameters,omitempty"`
}

// NewServiceBindingResource instantiates a new ServiceBindingResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBindingResource() *ServiceBindingResource {
	this := ServiceBindingResource{}
	return &this
}

// NewServiceBindingResourceWithDefaults instantiates a new ServiceBindingResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindingResourceWithDefaults() *ServiceBindingResource {
	this := ServiceBindingResource{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServiceBindingResource) GetMetadata() ServiceBindingMetadata {
	if o == nil || o.Metadata == nil {
		var ret ServiceBindingMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResource) GetMetadataOk() (*ServiceBindingMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServiceBindingResource) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ServiceBindingMetadata and assigns it to the Metadata field.
func (o *ServiceBindingResource) SetMetadata(v ServiceBindingMetadata) {
	o.Metadata = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *ServiceBindingResource) GetCredentials() map[string]interface{} {
	if o == nil || o.Credentials == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResource) GetCredentialsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *ServiceBindingResource) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given map[string]interface{} and assigns it to the Credentials field.
func (o *ServiceBindingResource) SetCredentials(v map[string]interface{}) {
	o.Credentials = &v
}

// GetSyslogDrainUrl returns the SyslogDrainUrl field value if set, zero value otherwise.
func (o *ServiceBindingResource) GetSyslogDrainUrl() string {
	if o == nil || o.SyslogDrainUrl == nil {
		var ret string
		return ret
	}
	return *o.SyslogDrainUrl
}

// GetSyslogDrainUrlOk returns a tuple with the SyslogDrainUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResource) GetSyslogDrainUrlOk() (*string, bool) {
	if o == nil || o.SyslogDrainUrl == nil {
		return nil, false
	}
	return o.SyslogDrainUrl, true
}

// HasSyslogDrainUrl returns a boolean if a field has been set.
func (o *ServiceBindingResource) HasSyslogDrainUrl() bool {
	if o != nil && o.SyslogDrainUrl != nil {
		return true
	}

	return false
}

// SetSyslogDrainUrl gets a reference to the given string and assigns it to the SyslogDrainUrl field.
func (o *ServiceBindingResource) SetSyslogDrainUrl(v string) {
	o.SyslogDrainUrl = &v
}

// GetRouteServiceUrl returns the RouteServiceUrl field value if set, zero value otherwise.
func (o *ServiceBindingResource) GetRouteServiceUrl() string {
	if o == nil || o.RouteServiceUrl == nil {
		var ret string
		return ret
	}
	return *o.RouteServiceUrl
}

// GetRouteServiceUrlOk returns a tuple with the RouteServiceUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResource) GetRouteServiceUrlOk() (*string, bool) {
	if o == nil || o.RouteServiceUrl == nil {
		return nil, false
	}
	return o.RouteServiceUrl, true
}

// HasRouteServiceUrl returns a boolean if a field has been set.
func (o *ServiceBindingResource) HasRouteServiceUrl() bool {
	if o != nil && o.RouteServiceUrl != nil {
		return true
	}

	return false
}

// SetRouteServiceUrl gets a reference to the given string and assigns it to the RouteServiceUrl field.
func (o *ServiceBindingResource) SetRouteServiceUrl(v string) {
	o.RouteServiceUrl = &v
}

// GetVolumeMounts returns the VolumeMounts field value if set, zero value otherwise.
func (o *ServiceBindingResource) GetVolumeMounts() []ServiceBindingVolumeMount {
	if o == nil || o.VolumeMounts == nil {
		var ret []ServiceBindingVolumeMount
		return ret
	}
	return *o.VolumeMounts
}

// GetVolumeMountsOk returns a tuple with the VolumeMounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResource) GetVolumeMountsOk() (*[]ServiceBindingVolumeMount, bool) {
	if o == nil || o.VolumeMounts == nil {
		return nil, false
	}
	return o.VolumeMounts, true
}

// HasVolumeMounts returns a boolean if a field has been set.
func (o *ServiceBindingResource) HasVolumeMounts() bool {
	if o != nil && o.VolumeMounts != nil {
		return true
	}

	return false
}

// SetVolumeMounts gets a reference to the given []ServiceBindingVolumeMount and assigns it to the VolumeMounts field.
func (o *ServiceBindingResource) SetVolumeMounts(v []ServiceBindingVolumeMount) {
	o.VolumeMounts = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *ServiceBindingResource) GetEndpoints() []ServiceBindingEndpoint {
	if o == nil || o.Endpoints == nil {
		var ret []ServiceBindingEndpoint
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResource) GetEndpointsOk() (*[]ServiceBindingEndpoint, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *ServiceBindingResource) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []ServiceBindingEndpoint and assigns it to the Endpoints field.
func (o *ServiceBindingResource) SetEndpoints(v []ServiceBindingEndpoint) {
	o.Endpoints = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ServiceBindingResource) GetParameters() map[string]interface{} {
	if o == nil || o.Parameters == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResource) GetParametersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ServiceBindingResource) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *ServiceBindingResource) SetParameters(v map[string]interface{}) {
	o.Parameters = &v
}

func (o ServiceBindingResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if o.SyslogDrainUrl != nil {
		toSerialize["syslog_drain_url"] = o.SyslogDrainUrl
	}
	if o.RouteServiceUrl != nil {
		toSerialize["route_service_url"] = o.RouteServiceUrl
	}
	if o.VolumeMounts != nil {
		toSerialize["volume_mounts"] = o.VolumeMounts
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableServiceBindingResource struct {
	value *ServiceBindingResource
	isSet bool
}

func (v NullableServiceBindingResource) Get() *ServiceBindingResource {
	return v.value
}

func (v *NullableServiceBindingResource) Set(val *ServiceBindingResource) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBindingResource) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBindingResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBindingResource(val *ServiceBindingResource) *NullableServiceBindingResource {
	return &NullableServiceBindingResource{value: val, isSet: true}
}

func (v NullableServiceBindingResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBindingResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

