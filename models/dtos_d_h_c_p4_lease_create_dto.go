// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtosDHCP4LeaseCreateDto dtos d h c p4 lease create dto
//
// swagger:model dtos.DHCP4LeaseCreateDto
type DtosDHCP4LeaseCreateDto struct {

	// Expires datetime
	Expires string `json:"expires,omitempty"`

	// IP address in ipv4 format
	IP string `json:"ip,omitempty"`

	// MAC address in format like this 00-00-00-00-00
	Mac string `json:"mac,omitempty"`
}

// Validate validates this dtos d h c p4 lease create dto
func (m *DtosDHCP4LeaseCreateDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dtos d h c p4 lease create dto based on context it is used
func (m *DtosDHCP4LeaseCreateDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtosDHCP4LeaseCreateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtosDHCP4LeaseCreateDto) UnmarshalBinary(b []byte) error {
	var res DtosDHCP4LeaseCreateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
