// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Plan plan
//
// swagger:model Plan
type Plan struct {

	// bindable
	Bindable bool `json:"bindable,omitempty"`

	// binding rotatable
	BindingRotatable *bool `json:"binding_rotatable,omitempty"`

	// description
	// Required: true
	Description *string `json:"description"`

	// free
	Free *bool `json:"free,omitempty"`

	// id
	// Required: true
	ID *string `json:"id"`

	// maintenance info
	MaintenanceInfo *MaintenanceInfo `json:"maintenance_info,omitempty"`

	// maximum polling duration
	MaximumPollingDuration int64 `json:"maximum_polling_duration,omitempty"`

	// metadata
	Metadata Metadata `json:"metadata,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// plan updateable
	PlanUpdateable bool `json:"plan_updateable,omitempty"`

	// schemas
	Schemas *SchemasObject `json:"schemas,omitempty"`
}

// Validate validates this plan
func (m *Plan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaintenanceInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plan) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validateMaintenanceInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.MaintenanceInfo) { // not required
		return nil
	}

	if m.MaintenanceInfo != nil {
		if err := m.MaintenanceInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance_info")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validateSchemas(formats strfmt.Registry) error {
	if swag.IsZero(m.Schemas) { // not required
		return nil
	}

	if m.Schemas != nil {
		if err := m.Schemas.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schemas")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this plan based on the context it is used
func (m *Plan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMaintenanceInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchemas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plan) contextValidateMaintenanceInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.MaintenanceInfo != nil {
		if err := m.MaintenanceInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("maintenance_info")
			}
			return err
		}
	}

	return nil
}

func (m *Plan) contextValidateSchemas(ctx context.Context, formats strfmt.Registry) error {

	if m.Schemas != nil {
		if err := m.Schemas.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schemas")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plan) UnmarshalBinary(b []byte) error {
	var res Plan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
