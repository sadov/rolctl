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

// PutDhcpIDReader is a Reader for the PutDhcpID structure.
type PutDhcpIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDhcpIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutDhcpIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutDhcpIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutDhcpIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutDhcpIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutDhcpIDOK creates a PutDhcpIDOK with default headers values
func NewPutDhcpIDOK() *PutDhcpIDOK {
	return &PutDhcpIDOK{}
}

/*
PutDhcpIDOK describes a response with status code 200, with default header values.

OK
*/
type PutDhcpIDOK struct {
	Payload *models.DtosDHCP4ServerDto
}

// IsSuccess returns true when this put dhcp Id o k response has a 2xx status code
func (o *PutDhcpIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put dhcp Id o k response has a 3xx status code
func (o *PutDhcpIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put dhcp Id o k response has a 4xx status code
func (o *PutDhcpIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put dhcp Id o k response has a 5xx status code
func (o *PutDhcpIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put dhcp Id o k response a status code equal to that given
func (o *PutDhcpIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutDhcpIDOK) Error() string {
	return fmt.Sprintf("[PUT /dhcp/{id}][%d] putDhcpIdOK  %+v", 200, o.Payload)
}

func (o *PutDhcpIDOK) String() string {
	return fmt.Sprintf("[PUT /dhcp/{id}][%d] putDhcpIdOK  %+v", 200, o.Payload)
}

func (o *PutDhcpIDOK) GetPayload() *models.DtosDHCP4ServerDto {
	return o.Payload
}

func (o *PutDhcpIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosDHCP4ServerDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDhcpIDBadRequest creates a PutDhcpIDBadRequest with default headers values
func NewPutDhcpIDBadRequest() *PutDhcpIDBadRequest {
	return &PutDhcpIDBadRequest{}
}

/*
PutDhcpIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutDhcpIDBadRequest struct {
	Payload *models.DtosValidationErrorDto
}

// IsSuccess returns true when this put dhcp Id bad request response has a 2xx status code
func (o *PutDhcpIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put dhcp Id bad request response has a 3xx status code
func (o *PutDhcpIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put dhcp Id bad request response has a 4xx status code
func (o *PutDhcpIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put dhcp Id bad request response has a 5xx status code
func (o *PutDhcpIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put dhcp Id bad request response a status code equal to that given
func (o *PutDhcpIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutDhcpIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dhcp/{id}][%d] putDhcpIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutDhcpIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /dhcp/{id}][%d] putDhcpIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutDhcpIDBadRequest) GetPayload() *models.DtosValidationErrorDto {
	return o.Payload
}

func (o *PutDhcpIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosValidationErrorDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDhcpIDNotFound creates a PutDhcpIDNotFound with default headers values
func NewPutDhcpIDNotFound() *PutDhcpIDNotFound {
	return &PutDhcpIDNotFound{}
}

/*
PutDhcpIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutDhcpIDNotFound struct {
}

// IsSuccess returns true when this put dhcp Id not found response has a 2xx status code
func (o *PutDhcpIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put dhcp Id not found response has a 3xx status code
func (o *PutDhcpIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put dhcp Id not found response has a 4xx status code
func (o *PutDhcpIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put dhcp Id not found response has a 5xx status code
func (o *PutDhcpIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put dhcp Id not found response a status code equal to that given
func (o *PutDhcpIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutDhcpIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /dhcp/{id}][%d] putDhcpIdNotFound ", 404)
}

func (o *PutDhcpIDNotFound) String() string {
	return fmt.Sprintf("[PUT /dhcp/{id}][%d] putDhcpIdNotFound ", 404)
}

func (o *PutDhcpIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDhcpIDInternalServerError creates a PutDhcpIDInternalServerError with default headers values
func NewPutDhcpIDInternalServerError() *PutDhcpIDInternalServerError {
	return &PutDhcpIDInternalServerError{}
}

/*
PutDhcpIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutDhcpIDInternalServerError struct {
}

// IsSuccess returns true when this put dhcp Id internal server error response has a 2xx status code
func (o *PutDhcpIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put dhcp Id internal server error response has a 3xx status code
func (o *PutDhcpIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put dhcp Id internal server error response has a 4xx status code
func (o *PutDhcpIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put dhcp Id internal server error response has a 5xx status code
func (o *PutDhcpIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put dhcp Id internal server error response a status code equal to that given
func (o *PutDhcpIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutDhcpIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /dhcp/{id}][%d] putDhcpIdInternalServerError ", 500)
}

func (o *PutDhcpIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /dhcp/{id}][%d] putDhcpIdInternalServerError ", 500)
}

func (o *PutDhcpIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}