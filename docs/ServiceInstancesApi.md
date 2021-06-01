# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceInstanceDeprovision**](ServiceInstancesApi.md#ServiceInstanceDeprovision) | **Delete** /v2/service_instances/{instance_id} | deprovision a service instance
[**ServiceInstanceGet**](ServiceInstancesApi.md#ServiceInstanceGet) | **Get** /v2/service_instances/{instance_id} | get a service instance
[**ServiceInstanceLastOperationGet**](ServiceInstancesApi.md#ServiceInstanceLastOperationGet) | **Get** /v2/service_instances/{instance_id}/last_operation | get the last requested operation state for service instance
[**ServiceInstanceProvision**](ServiceInstancesApi.md#ServiceInstanceProvision) | **Put** /v2/service_instances/{instance_id} | provision a service instance
[**ServiceInstanceUpdate**](ServiceInstancesApi.md#ServiceInstanceUpdate) | **Patch** /v2/service_instances/{instance_id} | update a service instance

# **ServiceInstanceDeprovision**
> Object ServiceInstanceDeprovision(ctx, xBrokerAPIVersion, instanceId, serviceId, planId, optional)
deprovision a service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]
  **instanceId** | **string**| id of instance being deleted | 
  **serviceId** | **string**| id of the service associated with the instance being deleted | 
  **planId** | **string**| id of the plan associated with the instance being deleted | 
 **optional** | ***ServiceInstancesApiServiceInstanceDeprovisionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceInstancesApiServiceInstanceDeprovisionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xBrokerAPIOriginatingIdentity** | **optional.String**| identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **optional.Bool**| asynchronous deprovision supported | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceGet**
> ServiceInstanceResource ServiceInstanceGet(ctx, xBrokerAPIVersion, instanceId, optional)
get a service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]
  **instanceId** | **string**| instance id of instance to fetch | 
 **optional** | ***ServiceInstancesApiServiceInstanceGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceInstancesApiServiceInstanceGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xBrokerAPIOriginatingIdentity** | **optional.String**| identity of the user that initiated the request from the Platform | 
 **serviceId** | **optional.String**| id of the service associated with the instance | 
 **planId** | **optional.String**| id of the plan associated with the instance | 

### Return type

[**ServiceInstanceResource**](ServiceInstanceResource.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceLastOperationGet**
> LastOperationResource ServiceInstanceLastOperationGet(ctx, xBrokerAPIVersion, instanceId, optional)
get the last requested operation state for service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]
  **instanceId** | **string**| instance id of instance to find last operation applied to it | 
 **optional** | ***ServiceInstancesApiServiceInstanceLastOperationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceInstancesApiServiceInstanceLastOperationGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **serviceId** | **optional.String**| id of the service associated with the instance | 
 **planId** | **optional.String**| id of the plan associated with the instance | 
 **operation** | **optional.String**| a provided identifier for the operation | 

### Return type

[**LastOperationResource**](LastOperationResource.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceProvision**
> ServiceInstanceProvisionResponse ServiceInstanceProvision(ctx, body, xBrokerAPIVersion, instanceId, optional)
provision a service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInstanceProvisionRequestBody**](ServiceInstanceProvisionRequestBody.md)| parameters for the requested service instance provision | 
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]
  **instanceId** | **string**| instance id of instance to provision | 
 **optional** | ***ServiceInstancesApiServiceInstanceProvisionOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceInstancesApiServiceInstanceProvisionOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xBrokerAPIOriginatingIdentity** | **optional.**| identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **optional.**| asynchronous operations supported | 

### Return type

[**ServiceInstanceProvisionResponse**](ServiceInstanceProvisionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceInstanceUpdate**
> Object ServiceInstanceUpdate(ctx, body, xBrokerAPIVersion, instanceId, optional)
update a service instance

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceInstanceUpdateRequestBody**](ServiceInstanceUpdateRequestBody.md)| parameters for the requested service instance update | 
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]
  **instanceId** | **string**| instance id of instance to update | 
 **optional** | ***ServiceInstancesApiServiceInstanceUpdateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceInstancesApiServiceInstanceUpdateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xBrokerAPIOriginatingIdentity** | **optional.**| identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **optional.**| asynchronous operations supported | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

