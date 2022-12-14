// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtosValidationErrorElemDto dtos validation error elem dto
//
// swagger:model dtos.ValidationErrorElemDto
type DtosValidationErrorElemDto struct {

	// Value for Key
	Error string `json:"error,omitempty"`

	// Field with validation error
	Field string `json:"field,omitempty"`

	// Source for the field, for example query or body
	Source string `json:"source,omitempty"`
}

// Validate validates this dtos validation error elem dto
func (m *DtosValidationErrorElemDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dtos validation error elem dto based on context it is used
func (m *DtosValidationErrorElemDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtosValidationErrorElemDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtosValidationErrorElemDto) UnmarshalBinary(b []byte) error {
	var res DtosValidationErrorElemDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
