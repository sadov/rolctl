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

// PutHostNetworkBridgeNameReader is a Reader for the PutHostNetworkBridgeName structure.
type PutHostNetworkBridgeNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutHostNetworkBridgeNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutHostNetworkBridgeNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutHostNetworkBridgeNameBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutHostNetworkBridgeNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutHostNetworkBridgeNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutHostNetworkBridgeNameOK creates a PutHostNetworkBridgeNameOK with default headers values
func NewPutHostNetworkBridgeNameOK() *PutHostNetworkBridgeNameOK {
	return &PutHostNetworkBridgeNameOK{}
}

/*
PutHostNetworkBridgeNameOK describes a response with status code 200, with default header values.

OK
*/
type PutHostNetworkBridgeNameOK struct {
	Payload *models.DtosHostNetworkBridgeDto
}

// IsSuccess returns true when this put host network bridge name o k response has a 2xx status code
func (o *PutHostNetworkBridgeNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put host network bridge name o k response has a 3xx status code
func (o *PutHostNetworkBridgeNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put host network bridge name o k response has a 4xx status code
func (o *PutHostNetworkBridgeNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put host network bridge name o k response has a 5xx status code
func (o *PutHostNetworkBridgeNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put host network bridge name o k response a status code equal to that given
func (o *PutHostNetworkBridgeNameOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutHostNetworkBridgeNameOK) Error() string {
	return fmt.Sprintf("[PUT /host/network/bridge/{name}][%d] putHostNetworkBridgeNameOK  %+v", 200, o.Payload)
}

func (o *PutHostNetworkBridgeNameOK) String() string {
	return fmt.Sprintf("[PUT /host/network/bridge/{name}][%d] putHostNetworkBridgeNameOK  %+v", 200, o.Payload)
}

func (o *PutHostNetworkBridgeNameOK) GetPayload() *models.DtosHostNetworkBridgeDto {
	return o.Payload
}

func (o *PutHostNetworkBridgeNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosHostNetworkBridgeDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutHostNetworkBridgeNameBadRequest creates a PutHostNetworkBridgeNameBadRequest with default headers values
func NewPutHostNetworkBridgeNameBadRequest() *PutHostNetworkBridgeNameBadRequest {
	return &PutHostNetworkBridgeNameBadRequest{}
}

/*
PutHostNetworkBridgeNameBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutHostNetworkBridgeNameBadRequest struct {
	Payload *models.DtosValidationErrorDto
}

// IsSuccess returns true when this put host network bridge name bad request response has a 2xx status code
func (o *PutHostNetworkBridgeNameBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put host network bridge name bad request response has a 3xx status code
func (o *PutHostNetworkBridgeNameBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put host network bridge name bad request response has a 4xx status code
func (o *PutHostNetworkBridgeNameBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put host network bridge name bad request response has a 5xx status code
func (o *PutHostNetworkBridgeNameBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put host network bridge name bad request response a status code equal to that given
func (o *PutHostNetworkBridgeNameBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutHostNetworkBridgeNameBadRequest) Error() string {
	return fmt.Sprintf("[PUT /host/network/bridge/{name}][%d] putHostNetworkBridgeNameBadRequest  %+v", 400, o.Payload)
}

func (o *PutHostNetworkBridgeNameBadRequest) String() string {
	return fmt.Sprintf("[PUT /host/network/bridge/{name}][%d] putHostNetworkBridgeNameBadRequest  %+v", 400, o.Payload)
}

func (o *PutHostNetworkBridgeNameBadRequest) GetPayload() *models.DtosValidationErrorDto {
	return o.Payload
}

func (o *PutHostNetworkBridgeNameBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosValidationErrorDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutHostNetworkBridgeNameNotFound creates a PutHostNetworkBridgeNameNotFound with default headers values
func NewPutHostNetworkBridgeNameNotFound() *PutHostNetworkBridgeNameNotFound {
	return &PutHostNetworkBridgeNameNotFound{}
}

/*
PutHostNetworkBridgeNameNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutHostNetworkBridgeNameNotFound struct {
}

// IsSuccess returns true when this put host network bridge name not found response has a 2xx status code
func (o *PutHostNetworkBridgeNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put host network bridge name not found response has a 3xx status code
func (o *PutHostNetworkBridgeNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put host network bridge name not found response has a 4xx status code
func (o *PutHostNetworkBridgeNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put host network bridge name not found response has a 5xx status code
func (o *PutHostNetworkBridgeNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put host network bridge name not found response a status code equal to that given
func (o *PutHostNetworkBridgeNameNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutHostNetworkBridgeNameNotFound) Error() string {
	return fmt.Sprintf("[PUT /host/network/bridge/{name}][%d] putHostNetworkBridgeNameNotFound ", 404)
}

func (o *PutHostNetworkBridgeNameNotFound) String() string {
	return fmt.Sprintf("[PUT /host/network/bridge/{name}][%d] putHostNetworkBridgeNameNotFound ", 404)
}

func (o *PutHostNetworkBridgeNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutHostNetworkBridgeNameInternalServerError creates a PutHostNetworkBridgeNameInternalServerError with default headers values
func NewPutHostNetworkBridgeNameInternalServerError() *PutHostNetworkBridgeNameInternalServerError {
	return &PutHostNetworkBridgeNameInternalServerError{}
}

/*
PutHostNetworkBridgeNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutHostNetworkBridgeNameInternalServerError struct {
}

// IsSuccess returns true when this put host network bridge name internal server error response has a 2xx status code
func (o *PutHostNetworkBridgeNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put host network bridge name internal server error response has a 3xx status code
func (o *PutHostNetworkBridgeNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put host network bridge name internal server error response has a 4xx status code
func (o *PutHostNetworkBridgeNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put host network bridge name internal server error response has a 5xx status code
func (o *PutHostNetworkBridgeNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put host network bridge name internal server error response a status code equal to that given
func (o *PutHostNetworkBridgeNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutHostNetworkBridgeNameInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /host/network/bridge/{name}][%d] putHostNetworkBridgeNameInternalServerError ", 500)
}

func (o *PutHostNetworkBridgeNameInternalServerError) String() string {
	return fmt.Sprintf("[PUT /host/network/bridge/{name}][%d] putHostNetworkBridgeNameInternalServerError ", 500)
}

func (o *PutHostNetworkBridgeNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
