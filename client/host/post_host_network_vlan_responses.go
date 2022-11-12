// Code generated by go-swagger; DO NOT EDIT.

package host

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sadov/rolctl/models"
)

// PostHostNetworkVlanReader is a Reader for the PostHostNetworkVlan structure.
type PostHostNetworkVlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostHostNetworkVlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostHostNetworkVlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostHostNetworkVlanBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostHostNetworkVlanInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostHostNetworkVlanOK creates a PostHostNetworkVlanOK with default headers values
func NewPostHostNetworkVlanOK() *PostHostNetworkVlanOK {
	return &PostHostNetworkVlanOK{}
}

/*
PostHostNetworkVlanOK describes a response with status code 200, with default header values.

OK
*/
type PostHostNetworkVlanOK struct {
	Payload *models.DtosHostNetworkVlanDto
}

// IsSuccess returns true when this post host network vlan o k response has a 2xx status code
func (o *PostHostNetworkVlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post host network vlan o k response has a 3xx status code
func (o *PostHostNetworkVlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post host network vlan o k response has a 4xx status code
func (o *PostHostNetworkVlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post host network vlan o k response has a 5xx status code
func (o *PostHostNetworkVlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post host network vlan o k response a status code equal to that given
func (o *PostHostNetworkVlanOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostHostNetworkVlanOK) Error() string {
	return fmt.Sprintf("[POST /host/network/vlan/][%d] postHostNetworkVlanOK  %+v", 200, o.Payload)
}

func (o *PostHostNetworkVlanOK) String() string {
	return fmt.Sprintf("[POST /host/network/vlan/][%d] postHostNetworkVlanOK  %+v", 200, o.Payload)
}

func (o *PostHostNetworkVlanOK) GetPayload() *models.DtosHostNetworkVlanDto {
	return o.Payload
}

func (o *PostHostNetworkVlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosHostNetworkVlanDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHostNetworkVlanBadRequest creates a PostHostNetworkVlanBadRequest with default headers values
func NewPostHostNetworkVlanBadRequest() *PostHostNetworkVlanBadRequest {
	return &PostHostNetworkVlanBadRequest{}
}

/*
PostHostNetworkVlanBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostHostNetworkVlanBadRequest struct {
	Payload *models.DtosValidationErrorDto
}

// IsSuccess returns true when this post host network vlan bad request response has a 2xx status code
func (o *PostHostNetworkVlanBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post host network vlan bad request response has a 3xx status code
func (o *PostHostNetworkVlanBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post host network vlan bad request response has a 4xx status code
func (o *PostHostNetworkVlanBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post host network vlan bad request response has a 5xx status code
func (o *PostHostNetworkVlanBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post host network vlan bad request response a status code equal to that given
func (o *PostHostNetworkVlanBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostHostNetworkVlanBadRequest) Error() string {
	return fmt.Sprintf("[POST /host/network/vlan/][%d] postHostNetworkVlanBadRequest  %+v", 400, o.Payload)
}

func (o *PostHostNetworkVlanBadRequest) String() string {
	return fmt.Sprintf("[POST /host/network/vlan/][%d] postHostNetworkVlanBadRequest  %+v", 400, o.Payload)
}

func (o *PostHostNetworkVlanBadRequest) GetPayload() *models.DtosValidationErrorDto {
	return o.Payload
}

func (o *PostHostNetworkVlanBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosValidationErrorDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostHostNetworkVlanInternalServerError creates a PostHostNetworkVlanInternalServerError with default headers values
func NewPostHostNetworkVlanInternalServerError() *PostHostNetworkVlanInternalServerError {
	return &PostHostNetworkVlanInternalServerError{}
}

/*
PostHostNetworkVlanInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostHostNetworkVlanInternalServerError struct {
}

// IsSuccess returns true when this post host network vlan internal server error response has a 2xx status code
func (o *PostHostNetworkVlanInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post host network vlan internal server error response has a 3xx status code
func (o *PostHostNetworkVlanInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post host network vlan internal server error response has a 4xx status code
func (o *PostHostNetworkVlanInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post host network vlan internal server error response has a 5xx status code
func (o *PostHostNetworkVlanInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post host network vlan internal server error response a status code equal to that given
func (o *PostHostNetworkVlanInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostHostNetworkVlanInternalServerError) Error() string {
	return fmt.Sprintf("[POST /host/network/vlan/][%d] postHostNetworkVlanInternalServerError ", 500)
}

func (o *PostHostNetworkVlanInternalServerError) String() string {
	return fmt.Sprintf("[POST /host/network/vlan/][%d] postHostNetworkVlanInternalServerError ", 500)
}

func (o *PostHostNetworkVlanInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}