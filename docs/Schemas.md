# Schemas

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceInstance** | Pointer to [**ServiceInstanceSchema**](ServiceInstanceSchema.md) |  | [optional] 
**ServiceBinding** | Pointer to [**ServiceBindingSchema**](ServiceBindingSchema.md) |  | [optional] 

## Methods

### NewSchemas

`func NewSchemas() *Schemas`

NewSchemas instantiates a new Schemas object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemasWithDefaults

`func NewSchemasWithDefaults() *Schemas`

NewSchemasWithDefaults instantiates a new Schemas object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceInstance

`func (o *Schemas) GetServiceInstance() ServiceInstanceSchema`

GetServiceInstance returns the ServiceInstance field if non-nil, zero value otherwise.

### GetServiceInstanceOk

`func (o *Schemas) GetServiceInstanceOk() (*ServiceInstanceSchema, bool)`

GetServiceInstanceOk returns a tuple with the ServiceInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceInstance

`func (o *Schemas) SetServiceInstance(v ServiceInstanceSchema)`

SetServiceInstance sets ServiceInstance field to given value.

### HasServiceInstance

`func (o *Schemas) HasServiceInstance() bool`

HasServiceInstance returns a boolean if a field has been set.

### GetServiceBinding

`func (o *Schemas) GetServiceBinding() ServiceBindingSchema`

GetServiceBinding returns the ServiceBinding field if non-nil, zero value otherwise.

### GetServiceBindingOk

`func (o *Schemas) GetServiceBindingOk() (*ServiceBindingSchema, bool)`

GetServiceBindingOk returns a tuple with the ServiceBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceBinding

`func (o *Schemas) SetServiceBinding(v ServiceBindingSchema)`

SetServiceBinding sets ServiceBinding field to given value.

### HasServiceBinding

`func (o *Schemas) HasServiceBinding() bool`

HasServiceBinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


