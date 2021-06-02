// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service bindings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service bindings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ServiceBindingBinding(params *ServiceBindingBindingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingBindingOK, *ServiceBindingBindingCreated, *ServiceBindingBindingAccepted, error)

	ServiceBindingGet(params *ServiceBindingGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingGetOK, error)

	ServiceBindingLastOperationGet(params *ServiceBindingLastOperationGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingLastOperationGetOK, error)

	ServiceBindingUnbinding(params *ServiceBindingUnbindingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingUnbindingOK, *ServiceBindingUnbindingAccepted, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ServiceBindingBinding generations of a service binding
*/
func (a *Client) ServiceBindingBinding(params *ServiceBindingBindingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingBindingOK, *ServiceBindingBindingCreated, *ServiceBindingBindingAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingBindingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBinding.binding",
		Method:             "PUT",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceBindingBindingReader{formats: a.formats},
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
	case *ServiceBindingBindingOK:
		return value, nil, nil, nil
	case *ServiceBindingBindingCreated:
		return nil, value, nil, nil
	case *ServiceBindingBindingAccepted:
		return nil, nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceBindingBindingDefault)
	return nil, nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ServiceBindingGet gets a service binding
*/
func (a *Client) ServiceBindingGet(params *ServiceBindingGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBinding.get",
		Method:             "GET",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceBindingGetReader{formats: a.formats},
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
	success, ok := result.(*ServiceBindingGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceBindingGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ServiceBindingLastOperationGet lasts requested operation state for service binding
*/
func (a *Client) ServiceBindingLastOperationGet(params *ServiceBindingLastOperationGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingLastOperationGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingLastOperationGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBinding.lastOperation.get",
		Method:             "GET",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceBindingLastOperationGetReader{formats: a.formats},
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
	success, ok := result.(*ServiceBindingLastOperationGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceBindingLastOperationGetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ServiceBindingUnbinding deprovisions of a service binding
*/
func (a *Client) ServiceBindingUnbinding(params *ServiceBindingUnbindingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBindingUnbindingOK, *ServiceBindingUnbindingAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBindingUnbindingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBinding.unbinding",
		Method:             "DELETE",
		PathPattern:        "/v2/service_instances/{instance_id}/service_bindings/{binding_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServiceBindingUnbindingReader{formats: a.formats},
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
	case *ServiceBindingUnbindingOK:
		return value, nil, nil
	case *ServiceBindingUnbindingAccepted:
		return nil, value, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServiceBindingUnbindingDefault)
	return nil, nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}