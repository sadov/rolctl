// Code generated by go-swagger; DO NOT EDIT.

package log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sadov/rolctl/models"
)

// GetLogHTTPIDReader is a Reader for the GetLogHTTPID structure.
type GetLogHTTPIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogHTTPIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogHTTPIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetLogHTTPIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLogHTTPIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLogHTTPIDOK creates a GetLogHTTPIDOK with default headers values
func NewGetLogHTTPIDOK() *GetLogHTTPIDOK {
	return &GetLogHTTPIDOK{}
}

/*
GetLogHTTPIDOK describes a response with status code 200, with default header values.

OK
*/
type GetLogHTTPIDOK struct {
	Payload *models.DtosHTTPLogDto
}

// IsSuccess returns true when this get log Http Id o k response has a 2xx status code
func (o *GetLogHTTPIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get log Http Id o k response has a 3xx status code
func (o *GetLogHTTPIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log Http Id o k response has a 4xx status code
func (o *GetLogHTTPIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get log Http Id o k response has a 5xx status code
func (o *GetLogHTTPIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get log Http Id o k response a status code equal to that given
func (o *GetLogHTTPIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLogHTTPIDOK) Error() string {
	return fmt.Sprintf("[GET /log/http/{id}][%d] getLogHttpIdOK  %+v", 200, o.Payload)
}

func (o *GetLogHTTPIDOK) String() string {
	return fmt.Sprintf("[GET /log/http/{id}][%d] getLogHttpIdOK  %+v", 200, o.Payload)
}

func (o *GetLogHTTPIDOK) GetPayload() *models.DtosHTTPLogDto {
	return o.Payload
}

func (o *GetLogHTTPIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosHTTPLogDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLogHTTPIDNotFound creates a GetLogHTTPIDNotFound with default headers values
func NewGetLogHTTPIDNotFound() *GetLogHTTPIDNotFound {
	return &GetLogHTTPIDNotFound{}
}

/*
GetLogHTTPIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetLogHTTPIDNotFound struct {
}

// IsSuccess returns true when this get log Http Id not found response has a 2xx status code
func (o *GetLogHTTPIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get log Http Id not found response has a 3xx status code
func (o *GetLogHTTPIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log Http Id not found response has a 4xx status code
func (o *GetLogHTTPIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get log Http Id not found response has a 5xx status code
func (o *GetLogHTTPIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get log Http Id not found response a status code equal to that given
func (o *GetLogHTTPIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetLogHTTPIDNotFound) Error() string {
	return fmt.Sprintf("[GET /log/http/{id}][%d] getLogHttpIdNotFound ", 404)
}

func (o *GetLogHTTPIDNotFound) String() string {
	return fmt.Sprintf("[GET /log/http/{id}][%d] getLogHttpIdNotFound ", 404)
}

func (o *GetLogHTTPIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetLogHTTPIDInternalServerError creates a GetLogHTTPIDInternalServerError with default headers values
func NewGetLogHTTPIDInternalServerError() *GetLogHTTPIDInternalServerError {
	return &GetLogHTTPIDInternalServerError{}
}

/*
GetLogHTTPIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetLogHTTPIDInternalServerError struct {
}

// IsSuccess returns true when this get log Http Id internal server error response has a 2xx status code
func (o *GetLogHTTPIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get log Http Id internal server error response has a 3xx status code
func (o *GetLogHTTPIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get log Http Id internal server error response has a 4xx status code
func (o *GetLogHTTPIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get log Http Id internal server error response has a 5xx status code
func (o *GetLogHTTPIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get log Http Id internal server error response a status code equal to that given
func (o *GetLogHTTPIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLogHTTPIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /log/http/{id}][%d] getLogHttpIdInternalServerError ", 500)
}

func (o *GetLogHTTPIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /log/http/{id}][%d] getLogHttpIdInternalServerError ", 500)
}

func (o *GetLogHTTPIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}