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

// Service struct for Service
type Service struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Description string `json:"description"`
	Tags *[]string `json:"tags,omitempty"`
	Requires *[]string `json:"requires,omitempty"`
	Bindable bool `json:"bindable"`
	// See [Service Metadata Conventions](https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#service-metadata) for more details.
	Metadata *map[string]interface{} `json:"metadata,omitempty"`
	DashboardClient *DashboardClient `json:"dashboard_client,omitempty"`
	BindingRotatable *bool `json:"binding_rotatable,omitempty"`
	PlanUpdateable *bool `json:"plan_updateable,omitempty"`
	Plans []Plan `json:"plans"`
}

// NewService instantiates a new Service object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewService(name string, id string, description string, bindable bool, plans []Plan, ) *Service {
	this := Service{}
	this.Name = name
	this.Id = id
	this.Description = description
	this.Bindable = bindable
	this.Plans = plans
	return &this
}

// NewServiceWithDefaults instantiates a new Service object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceWithDefaults() *Service {
	this := Service{}
	return &this
}

// GetName returns the Name field value
func (o *Service) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Service) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Service) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *Service) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Service) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Service) SetId(v string) {
	o.Id = v
}

// GetDescription returns the Description field value
func (o *Service) GetDescription() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Service) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Service) SetDescription(v string) {
	o.Description = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Service) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Service) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Service) SetTags(v []string) {
	o.Tags = &v
}

// GetRequires returns the Requires field value if set, zero value otherwise.
func (o *Service) GetRequires() []string {
	if o == nil || o.Requires == nil {
		var ret []string
		return ret
	}
	return *o.Requires
}

// GetRequiresOk returns a tuple with the Requires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetRequiresOk() (*[]string, bool) {
	if o == nil || o.Requires == nil {
		return nil, false
	}
	return o.Requires, true
}

// HasRequires returns a boolean if a field has been set.
func (o *Service) HasRequires() bool {
	if o != nil && o.Requires != nil {
		return true
	}

	return false
}

// SetRequires gets a reference to the given []string and assigns it to the Requires field.
func (o *Service) SetRequires(v []string) {
	o.Requires = &v
}

// GetBindable returns the Bindable field value
func (o *Service) GetBindable() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Bindable
}

// GetBindableOk returns a tuple with the Bindable field value
// and a boolean to check if the value has been set.
func (o *Service) GetBindableOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bindable, true
}

// SetBindable sets field value
func (o *Service) SetBindable(v bool) {
	o.Bindable = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Service) GetMetadata() map[string]interface{} {
	if o == nil || o.Metadata == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetMetadataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Service) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *Service) SetMetadata(v map[string]interface{}) {
	o.Metadata = &v
}

// GetDashboardClient returns the DashboardClient field value if set, zero value otherwise.
func (o *Service) GetDashboardClient() DashboardClient {
	if o == nil || o.DashboardClient == nil {
		var ret DashboardClient
		return ret
	}
	return *o.DashboardClient
}

// GetDashboardClientOk returns a tuple with the DashboardClient field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetDashboardClientOk() (*DashboardClient, bool) {
	if o == nil || o.DashboardClient == nil {
		return nil, false
	}
	return o.DashboardClient, true
}

// HasDashboardClient returns a boolean if a field has been set.
func (o *Service) HasDashboardClient() bool {
	if o != nil && o.DashboardClient != nil {
		return true
	}

	return false
}

// SetDashboardClient gets a reference to the given DashboardClient and assigns it to the DashboardClient field.
func (o *Service) SetDashboardClient(v DashboardClient) {
	o.DashboardClient = &v
}

// GetBindingRotatable returns the BindingRotatable field value if set, zero value otherwise.
func (o *Service) GetBindingRotatable() bool {
	if o == nil || o.BindingRotatable == nil {
		var ret bool
		return ret
	}
	return *o.BindingRotatable
}

// GetBindingRotatableOk returns a tuple with the BindingRotatable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetBindingRotatableOk() (*bool, bool) {
	if o == nil || o.BindingRotatable == nil {
		return nil, false
	}
	return o.BindingRotatable, true
}

// HasBindingRotatable returns a boolean if a field has been set.
func (o *Service) HasBindingRotatable() bool {
	if o != nil && o.BindingRotatable != nil {
		return true
	}

	return false
}

// SetBindingRotatable gets a reference to the given bool and assigns it to the BindingRotatable field.
func (o *Service) SetBindingRotatable(v bool) {
	o.BindingRotatable = &v
}

// GetPlanUpdateable returns the PlanUpdateable field value if set, zero value otherwise.
func (o *Service) GetPlanUpdateable() bool {
	if o == nil || o.PlanUpdateable == nil {
		var ret bool
		return ret
	}
	return *o.PlanUpdateable
}

// GetPlanUpdateableOk returns a tuple with the PlanUpdateable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Service) GetPlanUpdateableOk() (*bool, bool) {
	if o == nil || o.PlanUpdateable == nil {
		return nil, false
	}
	return o.PlanUpdateable, true
}

// HasPlanUpdateable returns a boolean if a field has been set.
func (o *Service) HasPlanUpdateable() bool {
	if o != nil && o.PlanUpdateable != nil {
		return true
	}

	return false
}

// SetPlanUpdateable gets a reference to the given bool and assigns it to the PlanUpdateable field.
func (o *Service) SetPlanUpdateable(v bool) {
	o.PlanUpdateable = &v
}

// GetPlans returns the Plans field value
func (o *Service) GetPlans() []Plan {
	if o == nil  {
		var ret []Plan
		return ret
	}

	return o.Plans
}

// GetPlansOk returns a tuple with the Plans field value
// and a boolean to check if the value has been set.
func (o *Service) GetPlansOk() (*[]Plan, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Plans, true
}

// SetPlans sets field value
func (o *Service) SetPlans(v []Plan) {
	o.Plans = v
}

func (o Service) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Requires != nil {
		toSerialize["requires"] = o.Requires
	}
	if true {
		toSerialize["bindable"] = o.Bindable
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.DashboardClient != nil {
		toSerialize["dashboard_client"] = o.DashboardClient
	}
	if o.BindingRotatable != nil {
		toSerialize["binding_rotatable"] = o.BindingRotatable
	}
	if o.PlanUpdateable != nil {
		toSerialize["plan_updateable"] = o.PlanUpdateable
	}
	if true {
		toSerialize["plans"] = o.Plans
	}
	return json.Marshal(toSerialize)
}

type NullableService struct {
	value *Service
	isSet bool
}

func (v NullableService) Get() *Service {
	return v.value
}

func (v *NullableService) Set(val *Service) {
	v.value = val
	v.isSet = true
}

func (v NullableService) IsSet() bool {
	return v.isSet
}

func (v *NullableService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableService(val *Service) *NullableService {
	return &NullableService{value: val, isSet: true}
}

func (v NullableService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


