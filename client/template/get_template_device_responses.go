// Code generated by go-swagger; DO NOT EDIT.

package template

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/sadov/rolctl/models"
)

// GetTemplateDeviceReader is a Reader for the GetTemplateDevice structure.
type GetTemplateDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTemplateDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTemplateDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetTemplateDeviceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTemplateDeviceOK creates a GetTemplateDeviceOK with default headers values
func NewGetTemplateDeviceOK() *GetTemplateDeviceOK {
	return &GetTemplateDeviceOK{}
}

/*
GetTemplateDeviceOK describes a response with status code 200, with default header values.

OK
*/
type GetTemplateDeviceOK struct {
	Payload *models.DtosPaginatedItemsDtoDtosDeviceTemplateDto
}

// IsSuccess returns true when this get template device o k response has a 2xx status code
func (o *GetTemplateDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get template device o k response has a 3xx status code
func (o *GetTemplateDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get template device o k response has a 4xx status code
func (o *GetTemplateDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get template device o k response has a 5xx status code
func (o *GetTemplateDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get template device o k response a status code equal to that given
func (o *GetTemplateDeviceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetTemplateDeviceOK) Error() string {
	return fmt.Sprintf("[GET /template/device/][%d] getTemplateDeviceOK  %+v", 200, o.Payload)
}

func (o *GetTemplateDeviceOK) String() string {
	return fmt.Sprintf("[GET /template/device/][%d] getTemplateDeviceOK  %+v", 200, o.Payload)
}

func (o *GetTemplateDeviceOK) GetPayload() *models.DtosPaginatedItemsDtoDtosDeviceTemplateDto {
	return o.Payload
}

func (o *GetTemplateDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosPaginatedItemsDtoDtosDeviceTemplateDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTemplateDeviceInternalServerError creates a GetTemplateDeviceInternalServerError with default headers values
func NewGetTemplateDeviceInternalServerError() *GetTemplateDeviceInternalServerError {
	return &GetTemplateDeviceInternalServerError{}
}

/*
GetTemplateDeviceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetTemplateDeviceInternalServerError struct {
}

// IsSuccess returns true when this get template device internal server error response has a 2xx status code
func (o *GetTemplateDeviceInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get template device internal server error response has a 3xx status code
func (o *GetTemplateDeviceInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get template device internal server error response has a 4xx status code
func (o *GetTemplateDeviceInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get template device internal server error response has a 5xx status code
func (o *GetTemplateDeviceInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get template device internal server error response a status code equal to that given
func (o *GetTemplateDeviceInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetTemplateDeviceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /template/device/][%d] getTemplateDeviceInternalServerError ", 500)
}

func (o *GetTemplateDeviceInternalServerError) String() string {
	return fmt.Sprintf("[GET /template/device/][%d] getTemplateDeviceInternalServerError ", 500)
}

func (o *GetTemplateDeviceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
