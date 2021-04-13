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

// Plan struct for Plan
type Plan struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	// See [Service Metadata Conventions](https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#service-metadata) for more details.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	MaintenanceInfo *MaintenanceInfo `json:"maintenance_info,omitempty"`
	Free *bool `json:"free,omitempty"`
	Bindable *bool `json:"bindable,omitempty"`
	Schemas *Schemas `json:"schemas,omitempty"`
}

// NewPlan instantiates a new Plan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlan(id string, name string, description string, ) *Plan {
	this := Plan{}
	this.Id = id
	this.Name = name
	this.Description = description
	var free bool = true
	this.Free = &free
	return &this
}

// NewPlanWithDefaults instantiates a new Plan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanWithDefaults() *Plan {
	this := Plan{}
	var free bool = true
	this.Free = &free
	return &this
}

// GetId returns the Id field value
func (o *Plan) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Plan) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Plan) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Plan) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Plan) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Plan) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *Plan) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Plan) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Plan) SetDescription(v string) {
	o.Description = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Plan) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Plan) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Plan) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetMaintenanceInfo returns the MaintenanceInfo field value if set, zero value otherwise.
func (o *Plan) GetMaintenanceInfo() MaintenanceInfo {
	if o == nil || o.MaintenanceInfo == nil {
		var ret MaintenanceInfo
		return ret
	}
	return *o.MaintenanceInfo
}

// GetMaintenanceInfoOk returns a tuple with the MaintenanceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetMaintenanceInfoOk() (*MaintenanceInfo, bool) {
	if o == nil || o.MaintenanceInfo == nil {
		return nil, false
	}
	return o.MaintenanceInfo, true
}

// HasMaintenanceInfo returns a boolean if a field has been set.
func (o *Plan) HasMaintenanceInfo() bool {
	if o != nil && o.MaintenanceInfo != nil {
		return true
	}

	return false
}

// SetMaintenanceInfo gets a reference to the given MaintenanceInfo and assigns it to the MaintenanceInfo field.
func (o *Plan) SetMaintenanceInfo(v MaintenanceInfo) {
	o.MaintenanceInfo = &v
}

// GetFree returns the Free field value if set, zero value otherwise.
func (o *Plan) GetFree() bool {
	if o == nil || o.Free == nil {
		var ret bool
		return ret
	}
	return *o.Free
}

// GetFreeOk returns a tuple with the Free field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetFreeOk() (*bool, bool) {
	if o == nil || o.Free == nil {
		return nil, false
	}
	return o.Free, true
}

// HasFree returns a boolean if a field has been set.
func (o *Plan) HasFree() bool {
	if o != nil && o.Free != nil {
		return true
	}

	return false
}

// SetFree gets a reference to the given bool and assigns it to the Free field.
func (o *Plan) SetFree(v bool) {
	o.Free = &v
}

// GetBindable returns the Bindable field value if set, zero value otherwise.
func (o *Plan) GetBindable() bool {
	if o == nil || o.Bindable == nil {
		var ret bool
		return ret
	}
	return *o.Bindable
}

// GetBindableOk returns a tuple with the Bindable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetBindableOk() (*bool, bool) {
	if o == nil || o.Bindable == nil {
		return nil, false
	}
	return o.Bindable, true
}

// HasBindable returns a boolean if a field has been set.
func (o *Plan) HasBindable() bool {
	if o != nil && o.Bindable != nil {
		return true
	}

	return false
}

// SetBindable gets a reference to the given bool and assigns it to the Bindable field.
func (o *Plan) SetBindable(v bool) {
	o.Bindable = &v
}

// GetSchemas returns the Schemas field value if set, zero value otherwise.
func (o *Plan) GetSchemas() Schemas {
	if o == nil || o.Schemas == nil {
		var ret Schemas
		return ret
	}
	return *o.Schemas
}

// GetSchemasOk returns a tuple with the Schemas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Plan) GetSchemasOk() (*Schemas, bool) {
	if o == nil || o.Schemas == nil {
		return nil, false
	}
	return o.Schemas, true
}

// HasSchemas returns a boolean if a field has been set.
func (o *Plan) HasSchemas() bool {
	if o != nil && o.Schemas != nil {
		return true
	}

	return false
}

// SetSchemas gets a reference to the given Schemas and assigns it to the Schemas field.
func (o *Plan) SetSchemas(v Schemas) {
	o.Schemas = &v
}

func (o Plan) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.MaintenanceInfo != nil {
		toSerialize["maintenance_info"] = o.MaintenanceInfo
	}
	if o.Free != nil {
		toSerialize["free"] = o.Free
	}
	if o.Bindable != nil {
		toSerialize["bindable"] = o.Bindable
	}
	if o.Schemas != nil {
		toSerialize["schemas"] = o.Schemas
	}
	return json.Marshal(toSerialize)
}

type NullablePlan struct {
	value *Plan
	isSet bool
}

func (v NullablePlan) Get() *Plan {
	return v.value
}

func (v *NullablePlan) Set(val *Plan) {
	v.value = val
	v.isSet = true
}

func (v NullablePlan) IsSet() bool {
	return v.isSet
}

func (v *NullablePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlan(val *Plan) *NullablePlan {
	return &NullablePlan{value: val, isSet: true}
}

func (v NullablePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

