// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtosEthernetSwitchDto dtos ethernet switch dto
//
// swagger:model dtos.EthernetSwitchDto
type DtosEthernetSwitchDto struct {

	// Address - switch ip address
	Address string `json:"address,omitempty"`

	// CreatedAt - entity create time
	CreatedAt string `json:"createdAt,omitempty"`

	// ID - unique identifier
	ID string `json:"id,omitempty"`

	// Name - switch name
	Name string `json:"name,omitempty"`

	// Serial - switch serial number
	Serial string `json:"serial,omitempty"`

	// SwitchModel - switch model
	SwitchModel string `json:"switchModel,omitempty"`

	// UpdatedAt - entity update time
	UpdatedAt string `json:"updatedAt,omitempty"`

	// Username - switch admin username
	Username string `json:"username,omitempty"`
}

// Validate validates this dtos ethernet switch dto
func (m *DtosEthernetSwitchDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dtos ethernet switch dto based on context it is used
func (m *DtosEthernetSwitchDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtosEthernetSwitchDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtosEthernetSwitchDto) UnmarshalBinary(b []byte) error {
	var res DtosEthernetSwitchDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
