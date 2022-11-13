// Code generated by go-swagger; DO NOT EDIT.

package ethernet_switch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sadov/rolctl/models"
)

// PostEthernetSwitchReader is a Reader for the PostEthernetSwitch structure.
type PostEthernetSwitchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostEthernetSwitchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostEthernetSwitchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostEthernetSwitchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostEthernetSwitchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostEthernetSwitchOK creates a PostEthernetSwitchOK with default headers values
func NewPostEthernetSwitchOK() *PostEthernetSwitchOK {
	return &PostEthernetSwitchOK{}
}

/*
PostEthernetSwitchOK describes a response with status code 200, with default header values.

OK
*/
type PostEthernetSwitchOK struct {
	Payload *models.DtosEthernetSwitchDto
}

// IsSuccess returns true when this post ethernet switch o k response has a 2xx status code
func (o *PostEthernetSwitchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post ethernet switch o k response has a 3xx status code
func (o *PostEthernetSwitchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post ethernet switch o k response has a 4xx status code
func (o *PostEthernetSwitchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post ethernet switch o k response has a 5xx status code
func (o *PostEthernetSwitchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post ethernet switch o k response a status code equal to that given
func (o *PostEthernetSwitchOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostEthernetSwitchOK) Error() string {
	return fmt.Sprintf("[POST /ethernet-switch/][%d] postEthernetSwitchOK  %+v", 200, o.Payload)
}

func (o *PostEthernetSwitchOK) String() string {
	return fmt.Sprintf("[POST /ethernet-switch/][%d] postEthernetSwitchOK  %+v", 200, o.Payload)
}

func (o *PostEthernetSwitchOK) GetPayload() *models.DtosEthernetSwitchDto {
	return o.Payload
}

func (o *PostEthernetSwitchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosEthernetSwitchDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEthernetSwitchBadRequest creates a PostEthernetSwitchBadRequest with default headers values
func NewPostEthernetSwitchBadRequest() *PostEthernetSwitchBadRequest {
	return &PostEthernetSwitchBadRequest{}
}

/*
PostEthernetSwitchBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PostEthernetSwitchBadRequest struct {
	Payload *models.DtosValidationErrorDto
}

// IsSuccess returns true when this post ethernet switch bad request response has a 2xx status code
func (o *PostEthernetSwitchBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post ethernet switch bad request response has a 3xx status code
func (o *PostEthernetSwitchBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post ethernet switch bad request response has a 4xx status code
func (o *PostEthernetSwitchBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post ethernet switch bad request response has a 5xx status code
func (o *PostEthernetSwitchBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post ethernet switch bad request response a status code equal to that given
func (o *PostEthernetSwitchBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostEthernetSwitchBadRequest) Error() string {
	return fmt.Sprintf("[POST /ethernet-switch/][%d] postEthernetSwitchBadRequest  %+v", 400, o.Payload)
}

func (o *PostEthernetSwitchBadRequest) String() string {
	return fmt.Sprintf("[POST /ethernet-switch/][%d] postEthernetSwitchBadRequest  %+v", 400, o.Payload)
}

func (o *PostEthernetSwitchBadRequest) GetPayload() *models.DtosValidationErrorDto {
	return o.Payload
}

func (o *PostEthernetSwitchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosValidationErrorDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostEthernetSwitchInternalServerError creates a PostEthernetSwitchInternalServerError with default headers values
func NewPostEthernetSwitchInternalServerError() *PostEthernetSwitchInternalServerError {
	return &PostEthernetSwitchInternalServerError{}
}

/*
PostEthernetSwitchInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostEthernetSwitchInternalServerError struct {
}

// IsSuccess returns true when this post ethernet switch internal server error response has a 2xx status code
func (o *PostEthernetSwitchInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post ethernet switch internal server error response has a 3xx status code
func (o *PostEthernetSwitchInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post ethernet switch internal server error response has a 4xx status code
func (o *PostEthernetSwitchInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post ethernet switch internal server error response has a 5xx status code
func (o *PostEthernetSwitchInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post ethernet switch internal server error response a status code equal to that given
func (o *PostEthernetSwitchInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostEthernetSwitchInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ethernet-switch/][%d] postEthernetSwitchInternalServerError ", 500)
}

func (o *PostEthernetSwitchInternalServerError) String() string {
	return fmt.Sprintf("[POST /ethernet-switch/][%d] postEthernetSwitchInternalServerError ", 500)
}

func (o *PostEthernetSwitchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
