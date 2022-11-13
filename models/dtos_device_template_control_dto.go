// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtosDeviceTemplateControlDto dtos device template control dto
//
// swagger:model dtos.DeviceTemplateControlDto
type DtosDeviceTemplateControlDto struct {

	// Emergency how to control device power in case of emergency. As example: POE(For Rpi4), IPMI, ILO or PowerSwitch
	Emergency string `json:"emergency,omitempty"`

	// NextBoot how to change next boot device. As example: IPMI, ILO or NONE.
	// For example, NONE is used for Rpi4, we control next boot by u-boot files in boot stages.
	NextBoot string `json:"nextBoot,omitempty"`

	// Power how to control device power. As example: POE(For Rpi4), IPMI, ILO or PowerSwitch
	Power string `json:"power,omitempty"`
}

// Validate validates this dtos device template control dto
func (m *DtosDeviceTemplateControlDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dtos device template control dto based on context it is used
func (m *DtosDeviceTemplateControlDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtosDeviceTemplateControlDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtosDeviceTemplateControlDto) UnmarshalBinary(b []byte) error {
	var res DtosDeviceTemplateControlDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
