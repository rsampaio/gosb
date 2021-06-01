# {{classname}}

All URIs are relative to *http://example.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CatalogGet**](CatalogApi.md#CatalogGet) | **Get** /v2/catalog | get the catalog of services that the service broker offers

# **CatalogGet**
> Catalog CatalogGet(ctx, xBrokerAPIVersion)
get the catalog of services that the service broker offers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **xBrokerAPIVersion** | **string**| version number of the Service Broker API that the Platform will use | [default to 2.13]

### Return type

[**Catalog**](Catalog.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

