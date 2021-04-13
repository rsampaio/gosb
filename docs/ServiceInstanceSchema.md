# ServiceInstanceSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Create** | Pointer to [**ServiceInstanceSchemaCreate**](ServiceInstanceSchema_create.md) |  | [optional] 
**Update** | Pointer to [**ServiceInstanceSchemaCreate**](ServiceInstanceSchema_create.md) |  | [optional] 

## Methods

### NewServiceInstanceSchema

`func NewServiceInstanceSchema() *ServiceInstanceSchema`

NewServiceInstanceSchema instantiates a new ServiceInstanceSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceInstanceSchemaWithDefaults

`func NewServiceInstanceSchemaWithDefaults() *ServiceInstanceSchema`

NewServiceInstanceSchemaWithDefaults instantiates a new ServiceInstanceSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreate

`func (o *ServiceInstanceSchema) GetCreate() ServiceInstanceSchemaCreate`

GetCreate returns the Create field if non-nil, zero value otherwise.

### GetCreateOk

`func (o *ServiceInstanceSchema) GetCreateOk() (*ServiceInstanceSchemaCreate, bool)`

GetCreateOk returns a tuple with the Create field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreate

`func (o *ServiceInstanceSchema) SetCreate(v ServiceInstanceSchemaCreate)`

SetCreate sets Create field to given value.

### HasCreate

`func (o *ServiceInstanceSchema) HasCreate() bool`

HasCreate returns a boolean if a field has been set.

### GetUpdate

`func (o *ServiceInstanceSchema) GetUpdate() ServiceInstanceSchemaCreate`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *ServiceInstanceSchema) GetUpdateOk() (*ServiceInstanceSchemaCreate, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *ServiceInstanceSchema) SetUpdate(v ServiceInstanceSchemaCreate)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *ServiceInstanceSchema) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


