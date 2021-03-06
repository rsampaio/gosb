// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service instances API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service instances API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ServiceInstanceDeprovision(params *ServiceInstanceDeprovisionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceDeprovisionOK, *ServiceInstanceDeprovisionAccepted, error)

	ServiceInstanceGet(params *ServiceInstanceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceGetOK, error)

	ServiceInstanceLastOperationGet(params *ServiceInstanceLastOperationGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceLastOperationGetOK, error)

	ServiceInstanceProvision(params *ServiceInstanceProvisionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceProvisionOK, *ServiceInstanceProvisionCreated, *ServiceInstanceProvisionAccepted, error)

	ServiceInstanceUpdate(params *ServiceInstanceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceUpdateOK, *ServiceInstanceUpdateAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ServiceInstanceDeprovision deprovisions a service instance
*/
func (a *Client) ServiceInstanceDeprovision(params *ServiceInstanceDeprovisionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceDeprovisionOK, *ServiceInstanceDeprovisionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceInstanceDeprovisionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceInstance.deprovision",
		Method:             "DELETE",
		PathPattern:        "/v2/service_instances/{instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceInstanceDeprovisionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ServiceInstanceDeprovisionOK:
		return value, nil, nil
	case *ServiceInstanceDeprovisionAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceInstanceDeprovisionDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ServiceInstanceGet gets a service instance
*/
func (a *Client) ServiceInstanceGet(params *ServiceInstanceGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceInstanceGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceInstance.get",
		Method:             "GET",
		PathPattern:        "/v2/service_instances/{instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceInstanceGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceInstanceGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceInstanceGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ServiceInstanceLastOperationGet lasts requested operation state for service instance
*/
func (a *Client) ServiceInstanceLastOperationGet(params *ServiceInstanceLastOperationGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceLastOperationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceInstanceLastOperationGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceInstance.lastOperation.get",
		Method:             "GET",
		PathPattern:        "/v2/service_instances/{instance_id}/last_operation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceInstanceLastOperationGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ServiceInstanceLastOperationGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceInstanceLastOperationGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ServiceInstanceProvision provisions a service instance
*/
func (a *Client) ServiceInstanceProvision(params *ServiceInstanceProvisionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceProvisionOK, *ServiceInstanceProvisionCreated, *ServiceInstanceProvisionAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceInstanceProvisionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceInstance.provision",
		Method:             "PUT",
		PathPattern:        "/v2/service_instances/{instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceInstanceProvisionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ServiceInstanceProvisionOK:
		return value, nil, nil, nil
	case *ServiceInstanceProvisionCreated:
		return nil, value, nil, nil
	case *ServiceInstanceProvisionAccepted:
		return nil, nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceInstanceProvisionDefault)
	return nil, nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ServiceInstanceUpdate updates a service instance
*/
func (a *Client) ServiceInstanceUpdate(params *ServiceInstanceUpdateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceInstanceUpdateOK, *ServiceInstanceUpdateAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceInstanceUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceInstance.update",
		Method:             "PATCH",
		PathPattern:        "/v2/service_instances/{instance_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceInstanceUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ServiceInstanceUpdateOK:
		return value, nil, nil
	case *ServiceInstanceUpdateAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceInstanceUpdateDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
