// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Workspace workspace
//
// swagger:model Workspace
type Workspace struct {

	// Workspace Capabilities
	// Required: true
	Capabilities map[string]bool `json:"capabilities"`

	// The Workspace information
	// Required: true
	Details *WorkspaceDetails `json:"details"`

	// Workspace ID
	// Required: true
	ID *string `json:"id"`

	// The Workspace location
	// Required: true
	Location *WorkspaceLocation `json:"location"`

	// The Workspace name
	// Required: true
	Name *string `json:"name"`

	// The Workspace status
	// Required: true
	Status *string `json:"status"`

	// The Workspace type
	// Required: true
	// Enum: [off-premises on-premises]
	Type *string `json:"type"`
}

// Validate validates this workspace
func (m *Workspace) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workspace) validateCapabilities(formats strfmt.Registry) error {

	if err := validate.Required("capabilities", "body", m.Capabilities); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("details", "body", m.Details); err != nil {
		return err
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *Workspace) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *Workspace) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Workspace) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var workspaceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["off-premises","on-premises"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workspaceTypeTypePropEnum = append(workspaceTypeTypePropEnum, v)
	}
}

const (

	// WorkspaceTypeOffDashPremises captures enum value "off-premises"
	WorkspaceTypeOffDashPremises string = "off-premises"

	// WorkspaceTypeOnDashPremises captures enum value "on-premises"
	WorkspaceTypeOnDashPremises string = "on-premises"
)

// prop value enum
func (m *Workspace) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, workspaceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Workspace) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this workspace based on the context it is used
func (m *Workspace) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Workspace) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.Details != nil {

		if err := m.Details.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *Workspace) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if m.Location != nil {

		if err := m.Location.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Workspace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Workspace) UnmarshalBinary(b []byte) error {
	var res Workspace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
