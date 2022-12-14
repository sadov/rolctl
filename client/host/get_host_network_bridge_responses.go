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

// GetHostNetworkBridgeReader is a Reader for the GetHostNetworkBridge structure.
type GetHostNetworkBridgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHostNetworkBridgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetHostNetworkBridgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetHostNetworkBridgeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetHostNetworkBridgeOK creates a GetHostNetworkBridgeOK with default headers values
func NewGetHostNetworkBridgeOK() *GetHostNetworkBridgeOK {
	return &GetHostNetworkBridgeOK{}
}

/*
GetHostNetworkBridgeOK describes a response with status code 200, with default header values.

OK
*/
type GetHostNetworkBridgeOK struct {
	Payload []*models.DtosHostNetworkBridgeDto
}

// IsSuccess returns true when this get host network bridge o k response has a 2xx status code
func (o *GetHostNetworkBridgeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get host network bridge o k response has a 3xx status code
func (o *GetHostNetworkBridgeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host network bridge o k response has a 4xx status code
func (o *GetHostNetworkBridgeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host network bridge o k response has a 5xx status code
func (o *GetHostNetworkBridgeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get host network bridge o k response a status code equal to that given
func (o *GetHostNetworkBridgeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetHostNetworkBridgeOK) Error() string {
	return fmt.Sprintf("[GET /host/network/bridge/][%d] getHostNetworkBridgeOK  %+v", 200, o.Payload)
}

func (o *GetHostNetworkBridgeOK) String() string {
	return fmt.Sprintf("[GET /host/network/bridge/][%d] getHostNetworkBridgeOK  %+v", 200, o.Payload)
}

func (o *GetHostNetworkBridgeOK) GetPayload() []*models.DtosHostNetworkBridgeDto {
	return o.Payload
}

func (o *GetHostNetworkBridgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHostNetworkBridgeInternalServerError creates a GetHostNetworkBridgeInternalServerError with default headers values
func NewGetHostNetworkBridgeInternalServerError() *GetHostNetworkBridgeInternalServerError {
	return &GetHostNetworkBridgeInternalServerError{}
}

/*
GetHostNetworkBridgeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetHostNetworkBridgeInternalServerError struct {
}

// IsSuccess returns true when this get host network bridge internal server error response has a 2xx status code
func (o *GetHostNetworkBridgeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get host network bridge internal server error response has a 3xx status code
func (o *GetHostNetworkBridgeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get host network bridge internal server error response has a 4xx status code
func (o *GetHostNetworkBridgeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get host network bridge internal server error response has a 5xx status code
func (o *GetHostNetworkBridgeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get host network bridge internal server error response a status code equal to that given
func (o *GetHostNetworkBridgeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetHostNetworkBridgeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /host/network/bridge/][%d] getHostNetworkBridgeInternalServerError ", 500)
}

func (o *GetHostNetworkBridgeInternalServerError) String() string {
	return fmt.Sprintf("[GET /host/network/bridge/][%d] getHostNetworkBridgeInternalServerError ", 500)
}

func (o *GetHostNetworkBridgeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
