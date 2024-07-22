// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SharedProcessorPoolServer shared processor pool server
//
// swagger:model SharedProcessorPoolServer
type SharedProcessorPoolServer struct {

	// The amount of cpus for the server
	Cpus float64 `json:"Cpus,omitempty"`

	// Identifies if uncapped or not
	Uncapped bool `json:"Uncapped,omitempty"`

	// Availability zone for the server
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// The server ID
	ID string `json:"id,omitempty"`

	// The amount of memory for the server
	Memory int64 `json:"memory,omitempty"`

	// The server name
	Name string `json:"name,omitempty"`

	// Status of the server
	Status string `json:"status,omitempty"`

	// The amout of vcpus for the server
	Vcpus int64 `json:"vcpus,omitempty"`
}

// Validate validates this shared processor pool server
func (m *SharedProcessorPoolServer) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this shared processor pool server based on context it is used
func (m *SharedProcessorPoolServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SharedProcessorPoolServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SharedProcessorPoolServer) UnmarshalBinary(b []byte) error {
	var res SharedProcessorPoolServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
