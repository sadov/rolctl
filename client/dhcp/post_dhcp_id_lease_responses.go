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

// PostDhcpIDLeaseReader is a Reader for the PostDhcpIDLease structure.
type PostDhcpIDLeaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostDhcpIDLeaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostDhcpIDLeaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostDhcpIDLeaseBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostDhcpIDLeaseInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostDhcpIDLeaseOK creates a PostDhcpIDLeaseOK with default headers values
func NewPostDhcpIDLeaseOK() *PostDhcpIDLeaseOK {
	return &PostDhcpIDLeaseOK{}
}

/*
PostDhcpIDLeaseOK describes a response with status code 200, with default header values.

OK
*/
type PostDhcpIDLeaseOK struct {
	Payload *models.DtosDHCP4LeaseDto
}

// IsSuccess returns true when this post dhcp Id lease o k response has a 2xx status code
func (o *PostDhcpIDLeaseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post dhcp Id lease o k response has a 3xx status code
func (o *PostDhcpIDLeaseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post dhcp Id lease o k response has a 4xx status code
func (o *PostDhcpIDLeaseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post dhcp Id lease o k response has a 5xx status code
func (o *PostDhcpIDLeaseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post dhcp Id lease o k response a status code equal to that given
func (o *PostDhcpIDLeaseOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostDhcpIDLeaseOK) Error() string {
	return fmt.Sprintf("[POST /dhcp/{id}/lease][%d] postDhcpIdLeaseOK  %+v", 200, o.Payload)
}

func (o *PostDhcpIDLeaseOK) String() string {
	return fmt.Sprintf("[POST /dhcp/{id}/lease][%d] postDhcpIdLeaseOK  %+v", 200, o.Payload)
}

func (o *PostDhcpIDLeaseOK) GetPayload() *models.DtosDHCP4LeaseDto {
	return o.Payload
}

func (o *PostDhcpIDLeaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosDHCP4LeaseDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDhcpIDLeaseBadRequest creates a PostDhcpIDLeaseBadRequest with default headers values
func NewPostDhcpIDLeaseBadRequest() *PostDhcpIDLeaseBadRequest {
	return &PostDhcpIDLeaseBadRequest{}
}

/*
PostDhcpIDLeaseBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostDhcpIDLeaseBadRequest struct {
	Payload *models.DtosValidationErrorDto
}

// IsSuccess returns true when this post dhcp Id lease bad request response has a 2xx status code
func (o *PostDhcpIDLeaseBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post dhcp Id lease bad request response has a 3xx status code
func (o *PostDhcpIDLeaseBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post dhcp Id lease bad request response has a 4xx status code
func (o *PostDhcpIDLeaseBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post dhcp Id lease bad request response has a 5xx status code
func (o *PostDhcpIDLeaseBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post dhcp Id lease bad request response a status code equal to that given
func (o *PostDhcpIDLeaseBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostDhcpIDLeaseBadRequest) Error() string {
	return fmt.Sprintf("[POST /dhcp/{id}/lease][%d] postDhcpIdLeaseBadRequest  %+v", 400, o.Payload)
}

func (o *PostDhcpIDLeaseBadRequest) String() string {
	return fmt.Sprintf("[POST /dhcp/{id}/lease][%d] postDhcpIdLeaseBadRequest  %+v", 400, o.Payload)
}

func (o *PostDhcpIDLeaseBadRequest) GetPayload() *models.DtosValidationErrorDto {
	return o.Payload
}

func (o *PostDhcpIDLeaseBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosValidationErrorDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostDhcpIDLeaseInternalServerError creates a PostDhcpIDLeaseInternalServerError with default headers values
func NewPostDhcpIDLeaseInternalServerError() *PostDhcpIDLeaseInternalServerError {
	return &PostDhcpIDLeaseInternalServerError{}
}

/*
PostDhcpIDLeaseInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostDhcpIDLeaseInternalServerError struct {
}

// IsSuccess returns true when this post dhcp Id lease internal server error response has a 2xx status code
func (o *PostDhcpIDLeaseInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post dhcp Id lease internal server error response has a 3xx status code
func (o *PostDhcpIDLeaseInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post dhcp Id lease internal server error response has a 4xx status code
func (o *PostDhcpIDLeaseInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post dhcp Id lease internal server error response has a 5xx status code
func (o *PostDhcpIDLeaseInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post dhcp Id lease internal server error response a status code equal to that given
func (o *PostDhcpIDLeaseInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostDhcpIDLeaseInternalServerError) Error() string {
	return fmt.Sprintf("[POST /dhcp/{id}/lease][%d] postDhcpIdLeaseInternalServerError ", 500)
}

func (o *PostDhcpIDLeaseInternalServerError) String() string {
	return fmt.Sprintf("[POST /dhcp/{id}/lease][%d] postDhcpIdLeaseInternalServerError ", 500)
}

func (o *PostDhcpIDLeaseInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
