# \ServiceBindingsApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceBindingBinding**](ServiceBindingsApi.md#ServiceBindingBinding) | **Put** /v2/service_instances/{instance_id}/service_bindings/{binding_id} | generate a service binding
[**ServiceBindingGet**](ServiceBindingsApi.md#ServiceBindingGet) | **Get** /v2/service_instances/{instance_id}/service_bindings/{binding_id} | get a service binding
[**ServiceBindingLastOperationGet**](ServiceBindingsApi.md#ServiceBindingLastOperationGet) | **Get** /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation | get the last requested operation state for service binding
[**ServiceBindingUnbinding**](ServiceBindingsApi.md#ServiceBindingUnbinding) | **Delete** /v2/service_instances/{instance_id}/service_bindings/{binding_id} | deprovision a service binding



## ServiceBindingBinding

> ServiceBindingResponse ServiceBindingBinding(ctx, instanceId, bindingId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceBindingRequest(serviceBindingRequest).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()

generate a service binding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xBrokerAPIVersion := "xBrokerAPIVersion_example" // string | version number of the Service Broker API that the Platform will use (default to "2.13")
    instanceId := "instanceId_example" // string | instance id of instance to create a binding on
    bindingId := "bindingId_example" // string | binding id of binding to create
    serviceBindingRequest := *openapiclient.NewServiceBindingRequest("ServiceId_example", "PlanId_example") // ServiceBindingRequest | parameters for the requested service binding
    xBrokerAPIOriginatingIdentity := "xBrokerAPIOriginatingIdentity_example" // string | identity of the user that initiated the request from the Platform (optional)
    acceptsIncomplete := true // bool | asynchronous operations supported (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceBindingsApi.ServiceBindingBinding(context.Background(), instanceId, bindingId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceBindingRequest(serviceBindingRequest).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceBindingsApi.ServiceBindingBinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceBindingBinding`: ServiceBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceBindingsApi.ServiceBindingBinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | instance id of instance to create a binding on | 
**bindingId** | **string** | binding id of binding to create | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceBindingBindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBrokerAPIVersion** | **string** | version number of the Service Broker API that the Platform will use | [default to &quot;2.13&quot;]


 **serviceBindingRequest** | [**ServiceBindingRequest**](ServiceBindingRequest.md) | parameters for the requested service binding | 
 **xBrokerAPIOriginatingIdentity** | **string** | identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **bool** | asynchronous operations supported | 

### Return type

[**ServiceBindingResponse**](ServiceBindingResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceBindingGet

> ServiceBindingResource ServiceBindingGet(ctx, instanceId, bindingId).XBrokerAPIVersion(xBrokerAPIVersion).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).ServiceId(serviceId).PlanId(planId).Execute()

get a service binding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xBrokerAPIVersion := "xBrokerAPIVersion_example" // string | version number of the Service Broker API that the Platform will use (default to "2.13")
    instanceId := "instanceId_example" // string | instance id of instance associated with the binding
    bindingId := "bindingId_example" // string | binding id of binding to fetch
    xBrokerAPIOriginatingIdentity := "xBrokerAPIOriginatingIdentity_example" // string | identity of the user that initiated the request from the Platform (optional)
    serviceId := "serviceId_example" // string | id of the service associated with the instance (optional)
    planId := "planId_example" // string | id of the plan associated with the instance (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceBindingsApi.ServiceBindingGet(context.Background(), instanceId, bindingId).XBrokerAPIVersion(xBrokerAPIVersion).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).ServiceId(serviceId).PlanId(planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceBindingsApi.ServiceBindingGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceBindingGet`: ServiceBindingResource
    fmt.Fprintf(os.Stdout, "Response from `ServiceBindingsApi.ServiceBindingGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | instance id of instance associated with the binding | 
**bindingId** | **string** | binding id of binding to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceBindingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBrokerAPIVersion** | **string** | version number of the Service Broker API that the Platform will use | [default to &quot;2.13&quot;]


 **xBrokerAPIOriginatingIdentity** | **string** | identity of the user that initiated the request from the Platform | 
 **serviceId** | **string** | id of the service associated with the instance | 
 **planId** | **string** | id of the plan associated with the instance | 

### Return type

[**ServiceBindingResource**](ServiceBindingResource.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceBindingLastOperationGet

> LastOperationResource ServiceBindingLastOperationGet(ctx, instanceId, bindingId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceId(serviceId).PlanId(planId).Operation(operation).Execute()

get the last requested operation state for service binding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xBrokerAPIVersion := "xBrokerAPIVersion_example" // string | version number of the Service Broker API that the Platform will use (default to "2.13")
    instanceId := "instanceId_example" // string | instance id of instance to find last operation applied to it
    bindingId := "bindingId_example" // string | binding id of service binding to find last operation applied to it
    serviceId := "serviceId_example" // string | id of the service associated with the instance (optional)
    planId := "planId_example" // string | id of the plan associated with the instance (optional)
    operation := "operation_example" // string | a provided identifier for the operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceBindingsApi.ServiceBindingLastOperationGet(context.Background(), instanceId, bindingId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceId(serviceId).PlanId(planId).Operation(operation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceBindingsApi.ServiceBindingLastOperationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceBindingLastOperationGet`: LastOperationResource
    fmt.Fprintf(os.Stdout, "Response from `ServiceBindingsApi.ServiceBindingLastOperationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | instance id of instance to find last operation applied to it | 
**bindingId** | **string** | binding id of service binding to find last operation applied to it | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceBindingLastOperationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBrokerAPIVersion** | **string** | version number of the Service Broker API that the Platform will use | [default to &quot;2.13&quot;]


 **serviceId** | **string** | id of the service associated with the instance | 
 **planId** | **string** | id of the plan associated with the instance | 
 **operation** | **string** | a provided identifier for the operation | 

### Return type

[**LastOperationResource**](LastOperationResource.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceBindingUnbinding

> map[string]interface{} ServiceBindingUnbinding(ctx, instanceId, bindingId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceId(serviceId).PlanId(planId).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()

deprovision a service binding

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xBrokerAPIVersion := "xBrokerAPIVersion_example" // string | version number of the Service Broker API that the Platform will use (default to "2.13")
    instanceId := "instanceId_example" // string | id of the instance associated with the binding being deleted
    bindingId := "bindingId_example" // string | id of the binding being deleted
    serviceId := "serviceId_example" // string | id of the service associated with the instance for which a binding is being deleted
    planId := "planId_example" // string | id of the plan associated with the instance for which a binding is being deleted
    xBrokerAPIOriginatingIdentity := "xBrokerAPIOriginatingIdentity_example" // string | identity of the user that initiated the request from the Platform (optional)
    acceptsIncomplete := true // bool | asynchronous operations supported (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceBindingsApi.ServiceBindingUnbinding(context.Background(), instanceId, bindingId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceId(serviceId).PlanId(planId).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceBindingsApi.ServiceBindingUnbinding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceBindingUnbinding`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceBindingsApi.ServiceBindingUnbinding`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | id of the instance associated with the binding being deleted | 
**bindingId** | **string** | id of the binding being deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceBindingUnbindingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBrokerAPIVersion** | **string** | version number of the Service Broker API that the Platform will use | [default to &quot;2.13&quot;]


 **serviceId** | **string** | id of the service associated with the instance for which a binding is being deleted | 
 **planId** | **string** | id of the plan associated with the instance for which a binding is being deleted | 
 **xBrokerAPIOriginatingIdentity** | **string** | identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **bool** | asynchronous operations supported | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

