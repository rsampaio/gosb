// Code generated by go-swagger; DO NOT EDIT.

package service_bindings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rsampaio/gosb/models"
)

// ServiceBindingLastOperationGetReader is a Reader for the ServiceBindingLastOperationGet structure.
type ServiceBindingLastOperationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBindingLastOperationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBindingLastOperationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBindingLastOperationGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBindingLastOperationGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBindingLastOperationGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewServiceBindingLastOperationGetGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewServiceBindingLastOperationGetPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewServiceBindingLastOperationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceBindingLastOperationGetOK creates a ServiceBindingLastOperationGetOK with default headers values
func NewServiceBindingLastOperationGetOK() *ServiceBindingLastOperationGetOK {
	return &ServiceBindingLastOperationGetOK{}
}

/* ServiceBindingLastOperationGetOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBindingLastOperationGetOK struct {

	/* Indicates when to retry the request
	 */
	RetryAfter string

	Payload *models.LastOperationResource
}

func (o *ServiceBindingLastOperationGetOK) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetOK  %+v", 200, o.Payload)
}
func (o *ServiceBindingLastOperationGetOK) GetPayload() *models.LastOperationResource {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header RetryAfter
	hdrRetryAfter := response.GetHeader("RetryAfter")

	if hdrRetryAfter != "" {
		o.RetryAfter = hdrRetryAfter
	}

	o.Payload = new(models.LastOperationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetBadRequest creates a ServiceBindingLastOperationGetBadRequest with default headers values
func NewServiceBindingLastOperationGetBadRequest() *ServiceBindingLastOperationGetBadRequest {
	return &ServiceBindingLastOperationGetBadRequest{}
}

/* ServiceBindingLastOperationGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBindingLastOperationGetBadRequest struct {
	Payload *models.Error
}

func (o *ServiceBindingLastOperationGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetBadRequest  %+v", 400, o.Payload)
}
func (o *ServiceBindingLastOperationGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetUnauthorized creates a ServiceBindingLastOperationGetUnauthorized with default headers values
func NewServiceBindingLastOperationGetUnauthorized() *ServiceBindingLastOperationGetUnauthorized {
	return &ServiceBindingLastOperationGetUnauthorized{}
}

/* ServiceBindingLastOperationGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBindingLastOperationGetUnauthorized struct {
	Payload *models.Error
}

func (o *ServiceBindingLastOperationGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetUnauthorized  %+v", 401, o.Payload)
}
func (o *ServiceBindingLastOperationGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetNotFound creates a ServiceBindingLastOperationGetNotFound with default headers values
func NewServiceBindingLastOperationGetNotFound() *ServiceBindingLastOperationGetNotFound {
	return &ServiceBindingLastOperationGetNotFound{}
}

/* ServiceBindingLastOperationGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBindingLastOperationGetNotFound struct {
	Payload *models.Error
}

func (o *ServiceBindingLastOperationGetNotFound) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetNotFound  %+v", 404, o.Payload)
}
func (o *ServiceBindingLastOperationGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetGone creates a ServiceBindingLastOperationGetGone with default headers values
func NewServiceBindingLastOperationGetGone() *ServiceBindingLastOperationGetGone {
	return &ServiceBindingLastOperationGetGone{}
}

/* ServiceBindingLastOperationGetGone describes a response with status code 410, with default header values.

Gone
*/
type ServiceBindingLastOperationGetGone struct {
	Payload *models.Error
}

func (o *ServiceBindingLastOperationGetGone) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetGone  %+v", 410, o.Payload)
}
func (o *ServiceBindingLastOperationGetGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetPreconditionFailed creates a ServiceBindingLastOperationGetPreconditionFailed with default headers values
func NewServiceBindingLastOperationGetPreconditionFailed() *ServiceBindingLastOperationGetPreconditionFailed {
	return &ServiceBindingLastOperationGetPreconditionFailed{}
}

/* ServiceBindingLastOperationGetPreconditionFailed describes a response with status code 412, with default header values.

Precondition Failed
*/
type ServiceBindingLastOperationGetPreconditionFailed struct {
	Payload *models.Error
}

func (o *ServiceBindingLastOperationGetPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBindingLastOperationGetPreconditionFailed  %+v", 412, o.Payload)
}
func (o *ServiceBindingLastOperationGetPreconditionFailed) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBindingLastOperationGetDefault creates a ServiceBindingLastOperationGetDefault with default headers values
func NewServiceBindingLastOperationGetDefault(code int) *ServiceBindingLastOperationGetDefault {
	return &ServiceBindingLastOperationGetDefault{
		_statusCode: code,
	}
}

/* ServiceBindingLastOperationGetDefault describes a response with status code -1, with default header values.

Unexpected
*/
type ServiceBindingLastOperationGetDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the service binding last operation get default response
func (o *ServiceBindingLastOperationGetDefault) Code() int {
	return o._statusCode
}

func (o *ServiceBindingLastOperationGetDefault) Error() string {
	return fmt.Sprintf("[GET /v2/service_instances/{instance_id}/service_bindings/{binding_id}/last_operation][%d] serviceBinding.lastOperation.get default  %+v", o._statusCode, o.Payload)
}
func (o *ServiceBindingLastOperationGetDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBindingLastOperationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}