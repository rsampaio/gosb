// Code generated by go-swagger; DO NOT EDIT.

package service_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewServiceInstanceDeprovisionParams creates a new ServiceInstanceDeprovisionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceInstanceDeprovisionParams() *ServiceInstanceDeprovisionParams {
	return &ServiceInstanceDeprovisionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceInstanceDeprovisionParamsWithTimeout creates a new ServiceInstanceDeprovisionParams object
// with the ability to set a timeout on a request.
func NewServiceInstanceDeprovisionParamsWithTimeout(timeout time.Duration) *ServiceInstanceDeprovisionParams {
	return &ServiceInstanceDeprovisionParams{
		timeout: timeout,
	}
}

// NewServiceInstanceDeprovisionParamsWithContext creates a new ServiceInstanceDeprovisionParams object
// with the ability to set a context for a request.
func NewServiceInstanceDeprovisionParamsWithContext(ctx context.Context) *ServiceInstanceDeprovisionParams {
	return &ServiceInstanceDeprovisionParams{
		Context: ctx,
	}
}

// NewServiceInstanceDeprovisionParamsWithHTTPClient creates a new ServiceInstanceDeprovisionParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceInstanceDeprovisionParamsWithHTTPClient(client *http.Client) *ServiceInstanceDeprovisionParams {
	return &ServiceInstanceDeprovisionParams{
		HTTPClient: client,
	}
}

/* ServiceInstanceDeprovisionParams contains all the parameters to send to the API endpoint
   for the service instance deprovision operation.

   Typically these are written to a http.Request.
*/
type ServiceInstanceDeprovisionParams struct {

	/* XBrokerAPIOriginatingIdentity.

	   identity of the user that initiated the request from the Platform
	*/
	XBrokerAPIOriginatingIdentity *string

	/* XBrokerAPIRequestIdentity.

	   idenity of the request from the Platform
	*/
	XBrokerAPIRequestIdentity *string

	/* XBrokerAPIVersion.

	   version number of the Service Broker API that the Platform will use
	*/
	XBrokerAPIVersion string

	/* AcceptsIncomplete.

	   asynchronous operations supported
	*/
	AcceptsIncomplete *bool

	/* InstanceID.

	   instance id of instance to provision
	*/
	InstanceID string

	/* PlanID.

	   id of the plan associated with the instance being deleted
	*/
	PlanID string

	/* ServiceID.

	   id of the service associated with the instance being deleted
	*/
	ServiceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service instance deprovision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceInstanceDeprovisionParams) WithDefaults() *ServiceInstanceDeprovisionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service instance deprovision params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceInstanceDeprovisionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithTimeout(timeout time.Duration) *ServiceInstanceDeprovisionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithContext(ctx context.Context) *ServiceInstanceDeprovisionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithHTTPClient(client *http.Client) *ServiceInstanceDeprovisionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXBrokerAPIOriginatingIdentity adds the xBrokerAPIOriginatingIdentity to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithXBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity *string) *ServiceInstanceDeprovisionParams {
	o.SetXBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity)
	return o
}

// SetXBrokerAPIOriginatingIdentity adds the xBrokerApiOriginatingIdentity to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetXBrokerAPIOriginatingIdentity(xBrokerAPIOriginatingIdentity *string) {
	o.XBrokerAPIOriginatingIdentity = xBrokerAPIOriginatingIdentity
}

// WithXBrokerAPIRequestIdentity adds the xBrokerAPIRequestIdentity to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithXBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity *string) *ServiceInstanceDeprovisionParams {
	o.SetXBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity)
	return o
}

// SetXBrokerAPIRequestIdentity adds the xBrokerApiRequestIdentity to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetXBrokerAPIRequestIdentity(xBrokerAPIRequestIdentity *string) {
	o.XBrokerAPIRequestIdentity = xBrokerAPIRequestIdentity
}

// WithXBrokerAPIVersion adds the xBrokerAPIVersion to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithXBrokerAPIVersion(xBrokerAPIVersion string) *ServiceInstanceDeprovisionParams {
	o.SetXBrokerAPIVersion(xBrokerAPIVersion)
	return o
}

// SetXBrokerAPIVersion adds the xBrokerApiVersion to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetXBrokerAPIVersion(xBrokerAPIVersion string) {
	o.XBrokerAPIVersion = xBrokerAPIVersion
}

// WithAcceptsIncomplete adds the acceptsIncomplete to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithAcceptsIncomplete(acceptsIncomplete *bool) *ServiceInstanceDeprovisionParams {
	o.SetAcceptsIncomplete(acceptsIncomplete)
	return o
}

// SetAcceptsIncomplete adds the acceptsIncomplete to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetAcceptsIncomplete(acceptsIncomplete *bool) {
	o.AcceptsIncomplete = acceptsIncomplete
}

// WithInstanceID adds the instanceID to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithInstanceID(instanceID string) *ServiceInstanceDeprovisionParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithPlanID adds the planID to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithPlanID(planID string) *ServiceInstanceDeprovisionParams {
	o.SetPlanID(planID)
	return o
}

// SetPlanID adds the planId to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetPlanID(planID string) {
	o.PlanID = planID
}

// WithServiceID adds the serviceID to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) WithServiceID(serviceID string) *ServiceInstanceDeprovisionParams {
	o.SetServiceID(serviceID)
	return o
}

// SetServiceID adds the serviceId to the service instance deprovision params
func (o *ServiceInstanceDeprovisionParams) SetServiceID(serviceID string) {
	o.ServiceID = serviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceInstanceDeprovisionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XBrokerAPIOriginatingIdentity != nil {

		// header param X-Broker-API-Originating-Identity
		if err := r.SetHeaderParam("X-Broker-API-Originating-Identity", *o.XBrokerAPIOriginatingIdentity); err != nil {
			return err
		}
	}

	if o.XBrokerAPIRequestIdentity != nil {

		// header param X-Broker-API-Request-Identity
		if err := r.SetHeaderParam("X-Broker-API-Request-Identity", *o.XBrokerAPIRequestIdentity); err != nil {
			return err
		}
	}

	// header param X-Broker-API-Version
	if err := r.SetHeaderParam("X-Broker-API-Version", o.XBrokerAPIVersion); err != nil {
		return err
	}

	if o.AcceptsIncomplete != nil {

		// query param accepts_incomplete
		var qrAcceptsIncomplete bool

		if o.AcceptsIncomplete != nil {
			qrAcceptsIncomplete = *o.AcceptsIncomplete
		}
		qAcceptsIncomplete := swag.FormatBool(qrAcceptsIncomplete)
		if qAcceptsIncomplete != "" {

			if err := r.SetQueryParam("accepts_incomplete", qAcceptsIncomplete); err != nil {
				return err
			}
		}
	}

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
		return err
	}

	// query param plan_id
	qrPlanID := o.PlanID
	qPlanID := qrPlanID
	if qPlanID != "" {

		if err := r.SetQueryParam("plan_id", qPlanID); err != nil {
			return err
		}
	}

	// query param service_id
	qrServiceID := o.ServiceID
	qServiceID := qrServiceID
	if qServiceID != "" {

		if err := r.SetQueryParam("service_id", qServiceID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
