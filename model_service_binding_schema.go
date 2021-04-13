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

// ServiceBindingSchema struct for ServiceBindingSchema
type ServiceBindingSchema struct {
	Create *ServiceInstanceSchemaCreate `json:"create,omitempty"`
}

// NewServiceBindingSchema instantiates a new ServiceBindingSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceBindingSchema() *ServiceBindingSchema {
	this := ServiceBindingSchema{}
	return &this
}

// NewServiceBindingSchemaWithDefaults instantiates a new ServiceBindingSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindingSchemaWithDefaults() *ServiceBindingSchema {
	this := ServiceBindingSchema{}
	return &this
}

// GetCreate returns the Create field value if set, zero value otherwise.
func (o *ServiceBindingSchema) GetCreate() ServiceInstanceSchemaCreate {
	if o == nil || o.Create == nil {
		var ret ServiceInstanceSchemaCreate
		return ret
	}
	return *o.Create
}

// GetCreateOk returns a tuple with the Create field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceBindingSchema) GetCreateOk() (*ServiceInstanceSchemaCreate, bool) {
	if o == nil || o.Create == nil {
		return nil, false
	}
	return o.Create, true
}

// HasCreate returns a boolean if a field has been set.
func (o *ServiceBindingSchema) HasCreate() bool {
	if o != nil && o.Create != nil {
		return true
	}

	return false
}

// SetCreate gets a reference to the given ServiceInstanceSchemaCreate and assigns it to the Create field.
func (o *ServiceBindingSchema) SetCreate(v ServiceInstanceSchemaCreate) {
	o.Create = &v
}

func (o ServiceBindingSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Create != nil {
		toSerialize["create"] = o.Create
	}
	return json.Marshal(toSerialize)
}

type NullableServiceBindingSchema struct {
	value *ServiceBindingSchema
	isSet bool
}

func (v NullableServiceBindingSchema) Get() *ServiceBindingSchema {
	return v.value
}

func (v *NullableServiceBindingSchema) Set(val *ServiceBindingSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceBindingSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBindingSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBindingSchema(val *ServiceBindingSchema) *NullableServiceBindingSchema {
	return &NullableServiceBindingSchema{value: val, isSet: true}
}

func (v NullableServiceBindingSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBindingSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


