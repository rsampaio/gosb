# ServiceInstanceProvisionRequestBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceId** | **string** |  | 
**PlanId** | **string** |  | 
**Context** | Pointer to **map[string]interface{}** | See [Context Conventions](https://github.com/openservicebrokerapi/servicebroker/blob/master/profile.md#context-object) for more details. | [optional] 
**OrganizationGuid** | **string** |  | 
**SpaceGuid** | **string** |  | 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewServiceInstanceProvisionRequestBody

`func NewServiceInstanceProvisionRequestBody(serviceId string, planId string, organizationGuid string, spaceGuid string, ) *ServiceInstanceProvisionRequestBody`

NewServiceInstanceProvisionRequestBody instantiates a new ServiceInstanceProvisionRequestBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceProvisionRequestBodyWithDefaults

`func NewServiceInstanceProvisionRequestBodyWithDefaults() *ServiceInstanceProvisionRequestBody`

NewServiceInstanceProvisionRequestBodyWithDefaults instantiates a new ServiceInstanceProvisionRequestBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceId

`func (o *ServiceInstanceProvisionRequestBody) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *ServiceInstanceProvisionRequestBody) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *ServiceInstanceProvisionRequestBody) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetPlanId

`func (o *ServiceInstanceProvisionRequestBody) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *ServiceInstanceProvisionRequestBody) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *ServiceInstanceProvisionRequestBody) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.


### GetContext

`func (o *ServiceInstanceProvisionRequestBody) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ServiceInstanceProvisionRequestBody) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ServiceInstanceProvisionRequestBody) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *ServiceInstanceProvisionRequestBody) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetOrganizationGuid

`func (o *ServiceInstanceProvisionRequestBody) GetOrganizationGuid() string`

GetOrganizationGuid returns the OrganizationGuid field if non-nil, zero value otherwise.

### GetOrganizationGuidOk

`func (o *ServiceInstanceProvisionRequestBody) GetOrganizationGuidOk() (*string, bool)`

GetOrganizationGuidOk returns a tuple with the OrganizationGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationGuid

`func (o *ServiceInstanceProvisionRequestBody) SetOrganizationGuid(v string)`

SetOrganizationGuid sets OrganizationGuid field to given value.


### GetSpaceGuid

`func (o *ServiceInstanceProvisionRequestBody) GetSpaceGuid() string`

GetSpaceGuid returns the SpaceGuid field if non-nil, zero value otherwise.

### GetSpaceGuidOk

`func (o *ServiceInstanceProvisionRequestBody) GetSpaceGuidOk() (*string, bool)`

GetSpaceGuidOk returns a tuple with the SpaceGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceGuid

`func (o *ServiceInstanceProvisionRequestBody) SetSpaceGuid(v string)`

SetSpaceGuid sets SpaceGuid field to given value.


### GetParameters

`func (o *ServiceInstanceProvisionRequestBody) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ServiceInstanceProvisionRequestBody) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ServiceInstanceProvisionRequestBody) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ServiceInstanceProvisionRequestBody) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


