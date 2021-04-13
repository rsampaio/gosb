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

// ServiceBindingVolumeMountDevice struct for ServiceBindingVolumeMountDevice
type ServiceBindingVolumeMountDevice struct {
	VolumeId string `json:"volume_id"`
	MountConfig *map[string]interface{} `json:"mount_config,omitempty"`
}

// NewServiceBindingVolumeMountDevice instantiates a new ServiceBindingVolumeMountDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBindingVolumeMountDevice(volumeId string, ) *ServiceBindingVolumeMountDevice {
	this := ServiceBindingVolumeMountDevice{}
	this.VolumeId = volumeId
	return &this
}

// NewServiceBindingVolumeMountDeviceWithDefaults instantiates a new ServiceBindingVolumeMountDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindingVolumeMountDeviceWithDefaults() *ServiceBindingVolumeMountDevice {
	this := ServiceBindingVolumeMountDevice{}
	return &this
}

// GetVolumeId returns the VolumeId field value
func (o *ServiceBindingVolumeMountDevice) GetVolumeId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.VolumeId
}

// GetVolumeIdOk returns a tuple with the VolumeId field value
// and a boolean to check if the value has been set.
func (o *ServiceBindingVolumeMountDevice) GetVolumeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.VolumeId, true
}

// SetVolumeId sets field value
func (o *ServiceBindingVolumeMountDevice) SetVolumeId(v string) {
	o.VolumeId = v
}

// GetMountConfig returns the MountConfig field value if set, zero value otherwise.
func (o *ServiceBindingVolumeMountDevice) GetMountConfig() map[string]interface{} {
	if o == nil || o.MountConfig == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.MountConfig
}

// GetMountConfigOk returns a tuple with the MountConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingVolumeMountDevice) GetMountConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.MountConfig == nil {
		return nil, false
	}
	return o.MountConfig, true
}

// HasMountConfig returns a boolean if a field has been set.
func (o *ServiceBindingVolumeMountDevice) HasMountConfig() bool {
	if o != nil && o.MountConfig != nil {
		return true
	}

	return false
}

// SetMountConfig gets a reference to the given map[string]interface{} and assigns it to the MountConfig field.
func (o *ServiceBindingVolumeMountDevice) SetMountConfig(v map[string]interface{}) {
	o.MountConfig = &v
}

func (o ServiceBindingVolumeMountDevice) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["volume_id"] = o.VolumeId
	}
	if o.MountConfig != nil {
		toSerialize["mount_config"] = o.MountConfig
	}
	return json.Marshal(toSerialize)
}

type NullableServiceBindingVolumeMountDevice struct {
	value *ServiceBindingVolumeMountDevice
	isSet bool
}

func (v NullableServiceBindingVolumeMountDevice) Get() *ServiceBindingVolumeMountDevice {
	return v.value
}

func (v *NullableServiceBindingVolumeMountDevice) Set(val *ServiceBindingVolumeMountDevice) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBindingVolumeMountDevice) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBindingVolumeMountDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBindingVolumeMountDevice(val *ServiceBindingVolumeMountDevice) *NullableServiceBindingVolumeMountDevice {
	return &NullableServiceBindingVolumeMountDevice{value: val, isSet: true}
}

func (v NullableServiceBindingVolumeMountDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBindingVolumeMountDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


