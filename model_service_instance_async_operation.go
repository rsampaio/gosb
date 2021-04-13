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

// ServiceInstanceAsyncOperation struct for ServiceInstanceAsyncOperation
type ServiceInstanceAsyncOperation struct {
	DashboardUrl *string `json:"dashboard_url,omitempty"`
	Operation *string `json:"operation,omitempty"`
	Metadata *ServiceInstanceMetadata `json:"metadata,omitempty"`
}

// NewServiceInstanceAsyncOperation instantiates a new ServiceInstanceAsyncOperation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceInstanceAsyncOperation() *ServiceInstanceAsyncOperation {
	this := ServiceInstanceAsyncOperation{}
	return &this
}

// NewServiceInstanceAsyncOperationWithDefaults instantiates a new ServiceInstanceAsyncOperation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceInstanceAsyncOperationWithDefaults() *ServiceInstanceAsyncOperation {
	this := ServiceInstanceAsyncOperation{}
	return &this
}

// GetDashboardUrl returns the DashboardUrl field value if set, zero value otherwise.
func (o *ServiceInstanceAsyncOperation) GetDashboardUrl() string {
	if o == nil || o.DashboardUrl == nil {
		var ret string
		return ret
	}
	return *o.DashboardUrl
}

// GetDashboardUrlOk returns a tuple with the DashboardUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstanceAsyncOperation) GetDashboardUrlOk() (*string, bool) {
	if o == nil || o.DashboardUrl == nil {
		return nil, false
	}
	return o.DashboardUrl, true
}

// HasDashboardUrl returns a boolean if a field has been set.
func (o *ServiceInstanceAsyncOperation) HasDashboardUrl() bool {
	if o != nil && o.DashboardUrl != nil {
		return true
	}

	return false
}

// SetDashboardUrl gets a reference to the given string and assigns it to the DashboardUrl field.
func (o *ServiceInstanceAsyncOperation) SetDashboardUrl(v string) {
	o.DashboardUrl = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ServiceInstanceAsyncOperation) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstanceAsyncOperation) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ServiceInstanceAsyncOperation) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ServiceInstanceAsyncOperation) SetOperation(v string) {
	o.Operation = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ServiceInstanceAsyncOperation) GetMetadata() ServiceInstanceMetadata {
	if o == nil || o.Metadata == nil {
		var ret ServiceInstanceMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceInstanceAsyncOperation) GetMetadataOk() (*ServiceInstanceMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ServiceInstanceAsyncOperation) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ServiceInstanceMetadata and assigns it to the Metadata field.
func (o *ServiceInstanceAsyncOperation) SetMetadata(v ServiceInstanceMetadata) {
	o.Metadata = &v
}

func (o ServiceInstanceAsyncOperation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DashboardUrl != nil {
		toSerialize["dashboard_url"] = o.DashboardUrl
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableServiceInstanceAsyncOperation struct {
	value *ServiceInstanceAsyncOperation
	isSet bool
}

func (v NullableServiceInstanceAsyncOperation) Get() *ServiceInstanceAsyncOperation {
	return v.value
}

func (v *NullableServiceInstanceAsyncOperation) Set(val *ServiceInstanceAsyncOperation) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceInstanceAsyncOperation) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceInstanceAsyncOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceInstanceAsyncOperation(val *ServiceInstanceAsyncOperation) *NullableServiceInstanceAsyncOperation {
	return &NullableServiceInstanceAsyncOperation{value: val, isSet: true}
}

func (v NullableServiceInstanceAsyncOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceInstanceAsyncOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

