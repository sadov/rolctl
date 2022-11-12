// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtosEthernetSwitchVLANDto dtos ethernet switch v l a n dto
//
// swagger:model dtos.EthernetSwitchVLANDto
type DtosEthernetSwitchVLANDto struct {

	// CreatedAt - entity create time
	CreatedAt string `json:"createdAt,omitempty"`

	// EthernetSwitchID ethernet switch ID
	EthernetSwitchID string `json:"ethernetSwitchID,omitempty"`

	// ID - unique identifier
	ID string `json:"id,omitempty"`

	// TaggedPorts slice of tagged ports IDs
	TaggedPorts []string `json:"taggedPorts"`

	// UntaggedPorts slice of untagged ports IDs
	UntaggedPorts []string `json:"untaggedPorts"`

	// UpdatedAt - entity update time
	UpdatedAt string `json:"updatedAt,omitempty"`

	// VlanID VLAN ID
	VlanID int64 `json:"vlanID,omitempty"`
}

// Validate validates this dtos ethernet switch v l a n dto
func (m *DtosEthernetSwitchVLANDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dtos ethernet switch v l a n dto based on context it is used
func (m *DtosEthernetSwitchVLANDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtosEthernetSwitchVLANDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtosEthernetSwitchVLANDto) UnmarshalBinary(b []byte) error {
	var res DtosEthernetSwitchVLANDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
