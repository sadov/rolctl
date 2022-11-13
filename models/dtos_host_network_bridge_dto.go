// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtosHostNetworkBridgeDto dtos host network bridge dto
//
// swagger:model dtos.HostNetworkBridgeDto
type DtosHostNetworkBridgeDto struct {

	// Addresses list
	Addresses []string `json:"addresses"`

	// Name interface full name
	Name string `json:"name,omitempty"`

	// Slaves slice of slaves interfaces names
	Slaves []string `json:"slaves"`
}

// Validate validates this dtos host network bridge dto
func (m *DtosHostNetworkBridgeDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dtos host network bridge dto based on context it is used
func (m *DtosHostNetworkBridgeDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtosHostNetworkBridgeDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtosHostNetworkBridgeDto) UnmarshalBinary(b []byte) error {
	var res DtosHostNetworkBridgeDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
