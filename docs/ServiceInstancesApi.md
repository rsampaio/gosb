# \ServiceInstancesApi

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceInstanceDeprovision**](ServiceInstancesApi.md#ServiceInstanceDeprovision) | **Delete** /v2/service_instances/{instance_id} | deprovision a service instance
[**ServiceInstanceGet**](ServiceInstancesApi.md#ServiceInstanceGet) | **Get** /v2/service_instances/{instance_id} | get a service instance
[**ServiceInstanceLastOperationGet**](ServiceInstancesApi.md#ServiceInstanceLastOperationGet) | **Get** /v2/service_instances/{instance_id}/last_operation | get the last requested operation state for service instance
[**ServiceInstanceProvision**](ServiceInstancesApi.md#ServiceInstanceProvision) | **Put** /v2/service_instances/{instance_id} | provision a service instance
[**ServiceInstanceUpdate**](ServiceInstancesApi.md#ServiceInstanceUpdate) | **Patch** /v2/service_instances/{instance_id} | update a service instance



## ServiceInstanceDeprovision

> map[string]interface{} ServiceInstanceDeprovision(ctx, instanceId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceId(serviceId).PlanId(planId).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()

deprovision a service instance

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
    instanceId := "instanceId_example" // string | id of instance being deleted
    serviceId := "serviceId_example" // string | id of the service associated with the instance being deleted
    planId := "planId_example" // string | id of the plan associated with the instance being deleted
    xBrokerAPIOriginatingIdentity := "xBrokerAPIOriginatingIdentity_example" // string | identity of the user that initiated the request from the Platform (optional)
    acceptsIncomplete := true // bool | asynchronous deprovision supported (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceInstancesApi.ServiceInstanceDeprovision(context.Background(), instanceId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceId(serviceId).PlanId(planId).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceInstancesApi.ServiceInstanceDeprovision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceInstanceDeprovision`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceInstancesApi.ServiceInstanceDeprovision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | id of instance being deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceInstanceDeprovisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBrokerAPIVersion** | **string** | version number of the Service Broker API that the Platform will use | [default to &quot;2.13&quot;]

 **serviceId** | **string** | id of the service associated with the instance being deleted | 
 **planId** | **string** | id of the plan associated with the instance being deleted | 
 **xBrokerAPIOriginatingIdentity** | **string** | identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **bool** | asynchronous deprovision supported | 

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


## ServiceInstanceGet

> ServiceInstanceResource ServiceInstanceGet(ctx, instanceId).XBrokerAPIVersion(xBrokerAPIVersion).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).ServiceId(serviceId).PlanId(planId).Execute()

get a service instance

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
    instanceId := "instanceId_example" // string | instance id of instance to fetch
    xBrokerAPIOriginatingIdentity := "xBrokerAPIOriginatingIdentity_example" // string | identity of the user that initiated the request from the Platform (optional)
    serviceId := "serviceId_example" // string | id of the service associated with the instance (optional)
    planId := "planId_example" // string | id of the plan associated with the instance (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceInstancesApi.ServiceInstanceGet(context.Background(), instanceId).XBrokerAPIVersion(xBrokerAPIVersion).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).ServiceId(serviceId).PlanId(planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceInstancesApi.ServiceInstanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceInstanceGet`: ServiceInstanceResource
    fmt.Fprintf(os.Stdout, "Response from `ServiceInstancesApi.ServiceInstanceGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | instance id of instance to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceInstanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBrokerAPIVersion** | **string** | version number of the Service Broker API that the Platform will use | [default to &quot;2.13&quot;]

 **xBrokerAPIOriginatingIdentity** | **string** | identity of the user that initiated the request from the Platform | 
 **serviceId** | **string** | id of the service associated with the instance | 
 **planId** | **string** | id of the plan associated with the instance | 

### Return type

[**ServiceInstanceResource**](ServiceInstanceResource.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceInstanceLastOperationGet

> LastOperationResource ServiceInstanceLastOperationGet(ctx, instanceId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceId(serviceId).PlanId(planId).Operation(operation).Execute()

get the last requested operation state for service instance

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
    serviceId := "serviceId_example" // string | id of the service associated with the instance (optional)
    planId := "planId_example" // string | id of the plan associated with the instance (optional)
    operation := "operation_example" // string | a provided identifier for the operation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceInstancesApi.ServiceInstanceLastOperationGet(context.Background(), instanceId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceId(serviceId).PlanId(planId).Operation(operation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceInstancesApi.ServiceInstanceLastOperationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceInstanceLastOperationGet`: LastOperationResource
    fmt.Fprintf(os.Stdout, "Response from `ServiceInstancesApi.ServiceInstanceLastOperationGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | instance id of instance to find last operation applied to it | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceInstanceLastOperationGetRequest struct via the builder pattern


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


## ServiceInstanceProvision

> ServiceInstanceProvisionResponse ServiceInstanceProvision(ctx, instanceId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceInstanceProvisionRequestBody(serviceInstanceProvisionRequestBody).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()

provision a service instance

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
    instanceId := "instanceId_example" // string | instance id of instance to provision
    serviceInstanceProvisionRequestBody := *openapiclient.NewServiceInstanceProvisionRequestBody("ServiceId_example", "PlanId_example", "OrganizationGuid_example", "SpaceGuid_example") // ServiceInstanceProvisionRequestBody | parameters for the requested service instance provision
    xBrokerAPIOriginatingIdentity := "xBrokerAPIOriginatingIdentity_example" // string | identity of the user that initiated the request from the Platform (optional)
    acceptsIncomplete := true // bool | asynchronous operations supported (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceInstancesApi.ServiceInstanceProvision(context.Background(), instanceId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceInstanceProvisionRequestBody(serviceInstanceProvisionRequestBody).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceInstancesApi.ServiceInstanceProvision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceInstanceProvision`: ServiceInstanceProvisionResponse
    fmt.Fprintf(os.Stdout, "Response from `ServiceInstancesApi.ServiceInstanceProvision`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | instance id of instance to provision | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceInstanceProvisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBrokerAPIVersion** | **string** | version number of the Service Broker API that the Platform will use | [default to &quot;2.13&quot;]

 **serviceInstanceProvisionRequestBody** | [**ServiceInstanceProvisionRequestBody**](ServiceInstanceProvisionRequestBody.md) | parameters for the requested service instance provision | 
 **xBrokerAPIOriginatingIdentity** | **string** | identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **bool** | asynchronous operations supported | 

### Return type

[**ServiceInstanceProvisionResponse**](ServiceInstanceProvisionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceInstanceUpdate

> map[string]interface{} ServiceInstanceUpdate(ctx, instanceId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceInstanceUpdateRequestBody(serviceInstanceUpdateRequestBody).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()

update a service instance

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
    instanceId := "instanceId_example" // string | instance id of instance to update
    serviceInstanceUpdateRequestBody := *openapiclient.NewServiceInstanceUpdateRequestBody("ServiceId_example") // ServiceInstanceUpdateRequestBody | parameters for the requested service instance update
    xBrokerAPIOriginatingIdentity := "xBrokerAPIOriginatingIdentity_example" // string | identity of the user that initiated the request from the Platform (optional)
    acceptsIncomplete := true // bool | asynchronous operations supported (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceInstancesApi.ServiceInstanceUpdate(context.Background(), instanceId).XBrokerAPIVersion(xBrokerAPIVersion).ServiceInstanceUpdateRequestBody(serviceInstanceUpdateRequestBody).XBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity).AcceptsIncomplete(acceptsIncomplete).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceInstancesApi.ServiceInstanceUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceInstanceUpdate`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `ServiceInstancesApi.ServiceInstanceUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**instanceId** | **string** | instance id of instance to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiServiceInstanceUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xBrokerAPIVersion** | **string** | version number of the Service Broker API that the Platform will use | [default to &quot;2.13&quot;]

 **serviceInstanceUpdateRequestBody** | [**ServiceInstanceUpdateRequestBody**](ServiceInstanceUpdateRequestBody.md) | parameters for the requested service instance update | 
 **xBrokerAPIOriginatingIdentity** | **string** | identity of the user that initiated the request from the Platform | 
 **acceptsIncomplete** | **bool** | asynchronous operations supported | 

### Return type

**map[string]interface{}**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

