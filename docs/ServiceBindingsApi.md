# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceBindingBinding**](ServiceBindingsApi.md#ServiceBindingBinding) | **Put** /v2/service_instances/{instance_id}/service_bindings/{binding_id} | generate a service binding
[**ServiceBindingGet**](ServiceBindingsApi.md#ServiceBindingGet) | **Get** /v2/service_instances/{instance_id}/service_bindings/{binding_id} | get a service binding
[**ServiceBindingLastOperationGet**](ServiceBindingsApi.md#ServiceBindingLastOperationGet) | **Get** /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation | get the last requested operation state for service binding
[**ServiceBindingUnbinding**](ServiceBindingsApi.md#ServiceBindingUnbinding) | **Delete** /v2/service_instances/{instance_id}/service_bindings/{binding_id} | deprovision a service binding

# **ServiceBindingBinding**
> ServiceBindingResponse ServiceBindingBinding(ctx, body, xBrokerAPIVersion, instanceId, bindingId, optional)
generate a service binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ServiceBindingRequest**](ServiceBindingRequest.md)| parameters for the requested service binding | 
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]
  **instanceId** | **string**| instance id of instance to create a binding on | 
  **bindingId** | **string**| binding id of binding to create | 
 **optional** | ***ServiceBindingsApiServiceBindingBindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceBindingsApiServiceBindingBindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xBrokerAPIOriginatingIdentity** | **optional.**| identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **optional.**| asynchronous operations supported | 

### Return type

[**ServiceBindingResponse**](ServiceBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceBindingGet**
> ServiceBindingResource ServiceBindingGet(ctx, xBrokerAPIVersion, instanceId, bindingId, optional)
get a service binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]
  **instanceId** | **string**| instance id of instance associated with the binding | 
  **bindingId** | **string**| binding id of binding to fetch | 
 **optional** | ***ServiceBindingsApiServiceBindingGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceBindingsApiServiceBindingGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xBrokerAPIOriginatingIdentity** | **optional.String**| identity of the user that initiated the request from the Platform | 
 **serviceId** | **optional.String**| id of the service associated with the instance | 
 **planId** | **optional.String**| id of the plan associated with the instance | 

### Return type

[**ServiceBindingResource**](ServiceBindingResource.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ServiceBindingLastOperationGet**
> LastOperationResource ServiceBindingLastOperationGet(ctx, xBrokerAPIVersion, instanceId, bindingId, optional)
get the last requested operation state for service binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]
  **instanceId** | **string**| instance id of instance to find last operation applied to it | 
  **bindingId** | **string**| binding id of service binding to find last operation applied to it | 
 **optional** | ***ServiceBindingsApiServiceBindingLastOperationGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceBindingsApiServiceBindingLastOperationGetOpts struct
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

# **ServiceBindingUnbinding**
> Object ServiceBindingUnbinding(ctx, xBrokerAPIVersion, instanceId, bindingId, serviceId, planId, optional)
deprovision a service binding

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]
  **instanceId** | **string**| id of the instance associated with the binding being deleted | 
  **bindingId** | **string**| id of the binding being deleted | 
  **serviceId** | **string**| id of the service associated with the instance for which a binding is being deleted | 
  **planId** | **string**| id of the plan associated with the instance for which a binding is being deleted | 
 **optional** | ***ServiceBindingsApiServiceBindingUnbindingOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ServiceBindingsApiServiceBindingUnbindingOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **xBrokerAPIOriginatingIdentity** | **optional.String**| identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **optional.Bool**| asynchronous operations supported | 

### Return type

[**Object**](Object.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

