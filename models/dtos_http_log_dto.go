// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DtosHTTPLogDto dtos HTTP log dto
//
// swagger:model dtos.HTTPLogDto
type DtosHTTPLogDto struct {

	// ClientIP - client IP address
	ClientIP string `json:"clientIP,omitempty"`

	// CreatedAt - entity create time
	CreatedAt string `json:"createdAt,omitempty"`

	// CustomRequestHeaders - custom headers of the request
	CustomRequestHeaders string `json:"customRequestHeaders,omitempty"`

	// Domain - domain that processed the request
	Domain string `json:"domain,omitempty"`

	// HTTPMethod - http method
	Httpmethod string `json:"httpmethod,omitempty"`

	// ID - unique identifier
	ID string `json:"id,omitempty"`

	// Latency - latency in milliseconds
	Latency int64 `json:"latency,omitempty"`

	// QueryParams - query params passed
	QueryParams string `json:"queryParams,omitempty"`

	// RelativePath - path to the endpoint
	RelativePath string `json:"relativePath,omitempty"`

	// RequestBody - body of the request
	RequestBody string `json:"requestBody,omitempty"`

	// RequestHeaders - headers of the request
	RequestHeaders string `json:"requestHeaders,omitempty"`

	// ResponseBody - body of the response
	ResponseBody string `json:"responseBody,omitempty"`

	// UpdatedAt - entity update time
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this dtos HTTP log dto
func (m *DtosHTTPLogDto) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this dtos HTTP log dto based on context it is used
func (m *DtosHTTPLogDto) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DtosHTTPLogDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DtosHTTPLogDto) UnmarshalBinary(b []byte) error {
	var res DtosHTTPLogDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
