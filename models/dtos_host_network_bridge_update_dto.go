// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtosHostNetworkBridgeUpdateDto dtos host network bridge update dto
//
// swagger:model dtos.HostNetworkBridgeUpdateDto
type DtosHostNetworkBridgeUpdateDto struct {

	// Addresses list
	Addresses []string `json:"addresses"`

	// Slaves slice of slaves interfaces names
	Slaves []string `json:"slaves"`
}

// Validate validates this dtos host network bridge update dto
func (m *DtosHostNetworkBridgeUpdateDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dtos host network bridge update dto based on context it is used
func (m *DtosHostNetworkBridgeUpdateDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtosHostNetworkBridgeUpdateDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtosHostNetworkBridgeUpdateDto) UnmarshalBinary(b []byte) error {
	var res DtosHostNetworkBridgeUpdateDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}