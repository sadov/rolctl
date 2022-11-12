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

// GetDhcpIDReader is a Reader for the GetDhcpID structure.
type GetDhcpIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDhcpIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDhcpIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetDhcpIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDhcpIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDhcpIDOK creates a GetDhcpIDOK with default headers values
func NewGetDhcpIDOK() *GetDhcpIDOK {
	return &GetDhcpIDOK{}
}

/*
GetDhcpIDOK describes a response with status code 200, with default header values.

OK
*/
type GetDhcpIDOK struct {
	Payload *models.DtosDHCP4ServerDto
}

// IsSuccess returns true when this get dhcp Id o k response has a 2xx status code
func (o *GetDhcpIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get dhcp Id o k response has a 3xx status code
func (o *GetDhcpIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dhcp Id o k response has a 4xx status code
func (o *GetDhcpIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dhcp Id o k response has a 5xx status code
func (o *GetDhcpIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get dhcp Id o k response a status code equal to that given
func (o *GetDhcpIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDhcpIDOK) Error() string {
	return fmt.Sprintf("[GET /dhcp/{id}][%d] getDhcpIdOK  %+v", 200, o.Payload)
}

func (o *GetDhcpIDOK) String() string {
	return fmt.Sprintf("[GET /dhcp/{id}][%d] getDhcpIdOK  %+v", 200, o.Payload)
}

func (o *GetDhcpIDOK) GetPayload() *models.DtosDHCP4ServerDto {
	return o.Payload
}

func (o *GetDhcpIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosDHCP4ServerDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDhcpIDNotFound creates a GetDhcpIDNotFound with default headers values
func NewGetDhcpIDNotFound() *GetDhcpIDNotFound {
	return &GetDhcpIDNotFound{}
}

/*
GetDhcpIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetDhcpIDNotFound struct {
}

// IsSuccess returns true when this get dhcp Id not found response has a 2xx status code
func (o *GetDhcpIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dhcp Id not found response has a 3xx status code
func (o *GetDhcpIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dhcp Id not found response has a 4xx status code
func (o *GetDhcpIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get dhcp Id not found response has a 5xx status code
func (o *GetDhcpIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get dhcp Id not found response a status code equal to that given
func (o *GetDhcpIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetDhcpIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dhcp/{id}][%d] getDhcpIdNotFound ", 404)
}

func (o *GetDhcpIDNotFound) String() string {
	return fmt.Sprintf("[GET /dhcp/{id}][%d] getDhcpIdNotFound ", 404)
}

func (o *GetDhcpIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDhcpIDInternalServerError creates a GetDhcpIDInternalServerError with default headers values
func NewGetDhcpIDInternalServerError() *GetDhcpIDInternalServerError {
	return &GetDhcpIDInternalServerError{}
}

/*
GetDhcpIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDhcpIDInternalServerError struct {
}

// IsSuccess returns true when this get dhcp Id internal server error response has a 2xx status code
func (o *GetDhcpIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get dhcp Id internal server error response has a 3xx status code
func (o *GetDhcpIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get dhcp Id internal server error response has a 4xx status code
func (o *GetDhcpIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get dhcp Id internal server error response has a 5xx status code
func (o *GetDhcpIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get dhcp Id internal server error response a status code equal to that given
func (o *GetDhcpIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetDhcpIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dhcp/{id}][%d] getDhcpIdInternalServerError ", 500)
}

func (o *GetDhcpIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /dhcp/{id}][%d] getDhcpIdInternalServerError ", 500)
}

func (o *GetDhcpIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
