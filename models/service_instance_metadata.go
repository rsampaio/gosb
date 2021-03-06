// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceInstanceMetadata service instance metadata
//
// swagger:model ServiceInstanceMetadata
type ServiceInstanceMetadata struct {

	// attributes
	Attributes interface{} `json:"attributes,omitempty"`

	// labels
	Labels interface{} `json:"labels,omitempty"`
}

// Validate validates this service instance metadata
func (m *ServiceInstanceMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service instance metadata based on context it is used
func (m *ServiceInstanceMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceInstanceMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceInstanceMetadata) UnmarshalBinary(b []byte) error {
	var res ServiceInstanceMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
