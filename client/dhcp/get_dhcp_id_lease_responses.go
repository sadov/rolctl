// Code generated by go-swagger; DO NOT EDIT.

package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sadov/rolctl/models"
)

// GetDhcpIDLeaseReader is a Reader for the GetDhcpIDLease structure.
type GetDhcpIDLeaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDhcpIDLeaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDhcpIDLeaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetDhcpIDLeaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDhcpIDLeaseOK creates a GetDhcpIDLeaseOK with default headers values
func NewGetDhcpIDLeaseOK() *GetDhcpIDLeaseOK {
	return &GetDhcpIDLeaseOK{}
}

/*
GetDhcpIDLeaseOK describes a response with status code 200, with default header values.

OK
*/
type GetDhcpIDLeaseOK struct {
	Payload *models.DtosPaginatedItemsDtoDtosDHCP4LeaseDto
}

// IsSuccess returns true when this get dhcp Id lease o k response has a 2xx status code
func (o *GetDhcpIDLeaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dhcp Id lease o k response has a 3xx status code
func (o *GetDhcpIDLeaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dhcp Id lease o k response has a 4xx status code
func (o *GetDhcpIDLeaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dhcp Id lease o k response has a 5xx status code
func (o *GetDhcpIDLeaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dhcp Id lease o k response a status code equal to that given
func (o *GetDhcpIDLeaseOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDhcpIDLeaseOK) Error() string {
	return fmt.Sprintf("[GET /dhcp/{id}/lease][%d] getDhcpIdLeaseOK  %+v", 200, o.Payload)
}

func (o *GetDhcpIDLeaseOK) String() string {
	return fmt.Sprintf("[GET /dhcp/{id}/lease][%d] getDhcpIdLeaseOK  %+v", 200, o.Payload)
}

func (o *GetDhcpIDLeaseOK) GetPayload() *models.DtosPaginatedItemsDtoDtosDHCP4LeaseDto {
	return o.Payload
}

func (o *GetDhcpIDLeaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosPaginatedItemsDtoDtosDHCP4LeaseDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDhcpIDLeaseInternalServerError creates a GetDhcpIDLeaseInternalServerError with default headers values
func NewGetDhcpIDLeaseInternalServerError() *GetDhcpIDLeaseInternalServerError {
	return &GetDhcpIDLeaseInternalServerError{}
}

/*
GetDhcpIDLeaseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDhcpIDLeaseInternalServerError struct {
}

// IsSuccess returns true when this get dhcp Id lease internal server error response has a 2xx status code
func (o *GetDhcpIDLeaseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dhcp Id lease internal server error response has a 3xx status code
func (o *GetDhcpIDLeaseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dhcp Id lease internal server error response has a 4xx status code
func (o *GetDhcpIDLeaseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dhcp Id lease internal server error response has a 5xx status code
func (o *GetDhcpIDLeaseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get dhcp Id lease internal server error response a status code equal to that given
func (o *GetDhcpIDLeaseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDhcpIDLeaseInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dhcp/{id}/lease][%d] getDhcpIdLeaseInternalServerError ", 500)
}

func (o *GetDhcpIDLeaseInternalServerError) String() string {
	return fmt.Sprintf("[GET /dhcp/{id}/lease][%d] getDhcpIdLeaseInternalServerError ", 500)
}

func (o *GetDhcpIDLeaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
