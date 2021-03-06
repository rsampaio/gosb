// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/rsampaio/gosb/models"
)

// ServiceInstanceUpdateReader is a Reader for the ServiceInstanceUpdate structure.
type ServiceInstanceUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceInstanceUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceInstanceUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewServiceInstanceUpdateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceInstanceUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceInstanceUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewServiceInstanceUpdatePreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewServiceInstanceUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewServiceInstanceUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewServiceInstanceUpdateOK creates a ServiceInstanceUpdateOK with default headers values
func NewServiceInstanceUpdateOK() *ServiceInstanceUpdateOK {
	return &ServiceInstanceUpdateOK{}
}

/* ServiceInstanceUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ServiceInstanceUpdateOK struct {
	Payload interface{}
}

func (o *ServiceInstanceUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateOK  %+v", 200, o.Payload)
}
func (o *ServiceInstanceUpdateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *ServiceInstanceUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceUpdateAccepted creates a ServiceInstanceUpdateAccepted with default headers values
func NewServiceInstanceUpdateAccepted() *ServiceInstanceUpdateAccepted {
	return &ServiceInstanceUpdateAccepted{}
}

/* ServiceInstanceUpdateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type ServiceInstanceUpdateAccepted struct {
	Payload *models.ServiceInstanceAsyncOperation
}

func (o *ServiceInstanceUpdateAccepted) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateAccepted  %+v", 202, o.Payload)
}
func (o *ServiceInstanceUpdateAccepted) GetPayload() *models.ServiceInstanceAsyncOperation {
	return o.Payload
}

func (o *ServiceInstanceUpdateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceInstanceAsyncOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceUpdateBadRequest creates a ServiceInstanceUpdateBadRequest with default headers values
func NewServiceInstanceUpdateBadRequest() *ServiceInstanceUpdateBadRequest {
	return &ServiceInstanceUpdateBadRequest{}
}

/* ServiceInstanceUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceInstanceUpdateBadRequest struct {
	Payload *models.Error
}

func (o *ServiceInstanceUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateBadRequest  %+v", 400, o.Payload)
}
func (o *ServiceInstanceUpdateBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceUpdateUnauthorized creates a ServiceInstanceUpdateUnauthorized with default headers values
func NewServiceInstanceUpdateUnauthorized() *ServiceInstanceUpdateUnauthorized {
	return &ServiceInstanceUpdateUnauthorized{}
}

/* ServiceInstanceUpdateUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceInstanceUpdateUnauthorized struct {
	Payload *models.Error
}

func (o *ServiceInstanceUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateUnauthorized  %+v", 401, o.Payload)
}
func (o *ServiceInstanceUpdateUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceUpdatePreconditionFailed creates a ServiceInstanceUpdatePreconditionFailed with default headers values
func NewServiceInstanceUpdatePreconditionFailed() *ServiceInstanceUpdatePreconditionFailed {
	return &ServiceInstanceUpdatePreconditionFailed{}
}

/* ServiceInstanceUpdatePreconditionFailed describes a response with status code 412, with default header values.

Precondition Failed
*/
type ServiceInstanceUpdatePreconditionFailed struct {
	Payload *models.Error
}

func (o *ServiceInstanceUpdatePreconditionFailed) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdatePreconditionFailed  %+v", 412, o.Payload)
}
func (o *ServiceInstanceUpdatePreconditionFailed) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceUpdatePreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceUpdateUnprocessableEntity creates a ServiceInstanceUpdateUnprocessableEntity with default headers values
func NewServiceInstanceUpdateUnprocessableEntity() *ServiceInstanceUpdateUnprocessableEntity {
	return &ServiceInstanceUpdateUnprocessableEntity{}
}

/* ServiceInstanceUpdateUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ServiceInstanceUpdateUnprocessableEntity struct {
	Payload *models.Error
}

func (o *ServiceInstanceUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstanceUpdateUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ServiceInstanceUpdateUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceInstanceUpdateDefault creates a ServiceInstanceUpdateDefault with default headers values
func NewServiceInstanceUpdateDefault(code int) *ServiceInstanceUpdateDefault {
	return &ServiceInstanceUpdateDefault{
		_statusCode: code,
	}
}

/* ServiceInstanceUpdateDefault describes a response with status code -1, with default header values.

Unexpected
*/
type ServiceInstanceUpdateDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the service instance update default response
func (o *ServiceInstanceUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ServiceInstanceUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /v2/service_instances/{instance_id}][%d] serviceInstance.update default  %+v", o._statusCode, o.Payload)
}
func (o *ServiceInstanceUpdateDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceInstanceUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
