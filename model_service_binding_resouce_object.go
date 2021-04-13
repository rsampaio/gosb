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

// ServiceBindingResouceObject struct for ServiceBindingResouceObject
type ServiceBindingResouceObject struct {
	AppGuid *string `json:"app_guid,omitempty"`
	Route *string `json:"route,omitempty"`
}

// NewServiceBindingResouceObject instantiates a new ServiceBindingResouceObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBindingResouceObject() *ServiceBindingResouceObject {
	this := ServiceBindingResouceObject{}
	return &this
}

// NewServiceBindingResouceObjectWithDefaults instantiates a new ServiceBindingResouceObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindingResouceObjectWithDefaults() *ServiceBindingResouceObject {
	this := ServiceBindingResouceObject{}
	return &this
}

// GetAppGuid returns the AppGuid field value if set, zero value otherwise.
func (o *ServiceBindingResouceObject) GetAppGuid() string {
	if o == nil || o.AppGuid == nil {
		var ret string
		return ret
	}
	return *o.AppGuid
}

// GetAppGuidOk returns a tuple with the AppGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResouceObject) GetAppGuidOk() (*string, bool) {
	if o == nil || o.AppGuid == nil {
		return nil, false
	}
	return o.AppGuid, true
}

// HasAppGuid returns a boolean if a field has been set.
func (o *ServiceBindingResouceObject) HasAppGuid() bool {
	if o != nil && o.AppGuid != nil {
		return true
	}

	return false
}

// SetAppGuid gets a reference to the given string and assigns it to the AppGuid field.
func (o *ServiceBindingResouceObject) SetAppGuid(v string) {
	o.AppGuid = &v
}

// GetRoute returns the Route field value if set, zero value otherwise.
func (o *ServiceBindingResouceObject) GetRoute() string {
	if o == nil || o.Route == nil {
		var ret string
		return ret
	}
	return *o.Route
}

// GetRouteOk returns a tuple with the Route field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingResouceObject) GetRouteOk() (*string, bool) {
	if o == nil || o.Route == nil {
		return nil, false
	}
	return o.Route, true
}

// HasRoute returns a boolean if a field has been set.
func (o *ServiceBindingResouceObject) HasRoute() bool {
	if o != nil && o.Route != nil {
		return true
	}

	return false
}

// SetRoute gets a reference to the given string and assigns it to the Route field.
func (o *ServiceBindingResouceObject) SetRoute(v string) {
	o.Route = &v
}

func (o ServiceBindingResouceObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppGuid != nil {
		toSerialize["app_guid"] = o.AppGuid
	}
	if o.Route != nil {
		toSerialize["route"] = o.Route
	}
	return json.Marshal(toSerialize)
}

type NullableServiceBindingResouceObject struct {
	value *ServiceBindingResouceObject
	isSet bool
}

func (v NullableServiceBindingResouceObject) Get() *ServiceBindingResouceObject {
	return v.value
}

func (v *NullableServiceBindingResouceObject) Set(val *ServiceBindingResouceObject) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBindingResouceObject) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBindingResouceObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBindingResouceObject(val *ServiceBindingResouceObject) *NullableServiceBindingResouceObject {
	return &NullableServiceBindingResouceObject{value: val, isSet: true}
}

func (v NullableServiceBindingResouceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBindingResouceObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


