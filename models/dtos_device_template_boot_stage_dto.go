// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtosDeviceTemplateBootStageDto dtos device template boot stage dto
//
// swagger:model dtos.DeviceTemplateBootStageDto
type DtosDeviceTemplateBootStageDto struct {

	// Action for this boot stage.
	// Can be: File, CheckPowerSwitch, EmergencyPowerOff,
	// PowerOff, EmergencyPowerOn, PowerOn,
	// CheckManagement
	//
	// For File action:
	// 	A stage can only be marked complete if all files have
	// 	been downloaded by the device via TFTP or DHCP,
	// 	after which the next step can be loaded.
	Action string `json:"action,omitempty"`

	// Description of boot stage template
	Description string `json:"description,omitempty"`

	// Files slice of files for boot stage
	Files []*DtosDeviceTemplateBootStageFileDto `json:"files"`

	// Name of boot stage template
	Name string `json:"name,omitempty"`
}

// Validate validates this dtos device template boot stage dto
func (m *DtosDeviceTemplateBootStageDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtosDeviceTemplateBootStageDto) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {
		if swag.IsZero(m.Files[i]) { // not required
			continue
		}

		if m.Files[i] != nil {
			if err := m.Files[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dtos device template boot stage dto based on the context it is used
func (m *DtosDeviceTemplateBootStageDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DtosDeviceTemplateBootStageDto) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if m.Files[i] != nil {
			if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("files" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("files" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DtosDeviceTemplateBootStageDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtosDeviceTemplateBootStageDto) UnmarshalBinary(b []byte) error {
	var res DtosDeviceTemplateBootStageDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}