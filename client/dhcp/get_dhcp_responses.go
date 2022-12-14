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

// GetDhcpReader is a Reader for the GetDhcp structure.
type GetDhcpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDhcpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDhcpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetDhcpInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDhcpOK creates a GetDhcpOK with default headers values
func NewGetDhcpOK() *GetDhcpOK {
	return &GetDhcpOK{}
}

/*
GetDhcpOK describes a response with status code 200, with default header values.

OK
*/
type GetDhcpOK struct {
	Payload *models.DtosPaginatedItemsDtoDtosDHCP4ServerDto
}

// IsSuccess returns true when this get dhcp o k response has a 2xx status code
func (o *GetDhcpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dhcp o k response has a 3xx status code
func (o *GetDhcpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dhcp o k response has a 4xx status code
func (o *GetDhcpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dhcp o k response has a 5xx status code
func (o *GetDhcpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dhcp o k response a status code equal to that given
func (o *GetDhcpOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDhcpOK) Error() string {
	return fmt.Sprintf("[GET /dhcp/][%d] getDhcpOK  %+v", 200, o.Payload)
}

func (o *GetDhcpOK) String() string {
	return fmt.Sprintf("[GET /dhcp/][%d] getDhcpOK  %+v", 200, o.Payload)
}

func (o *GetDhcpOK) GetPayload() *models.DtosPaginatedItemsDtoDtosDHCP4ServerDto {
	return o.Payload
}

func (o *GetDhcpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosPaginatedItemsDtoDtosDHCP4ServerDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDhcpInternalServerError creates a GetDhcpInternalServerError with default headers values
func NewGetDhcpInternalServerError() *GetDhcpInternalServerError {
	return &GetDhcpInternalServerError{}
}

/*
GetDhcpInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDhcpInternalServerError struct {
}

// IsSuccess returns true when this get dhcp internal server error response has a 2xx status code
func (o *GetDhcpInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dhcp internal server error response has a 3xx status code
func (o *GetDhcpInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dhcp internal server error response has a 4xx status code
func (o *GetDhcpInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dhcp internal server error response has a 5xx status code
func (o *GetDhcpInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get dhcp internal server error response a status code equal to that given
func (o *GetDhcpInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDhcpInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dhcp/][%d] getDhcpInternalServerError ", 500)
}

func (o *GetDhcpInternalServerError) String() string {
	return fmt.Sprintf("[GET /dhcp/][%d] getDhcpInternalServerError ", 500)
}

func (o *GetDhcpInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
