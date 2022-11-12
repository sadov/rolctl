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

// GetEthernetSwitchIDVlanVlanUUIDReader is a Reader for the GetEthernetSwitchIDVlanVlanUUID structure.
type GetEthernetSwitchIDVlanVlanUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEthernetSwitchIDVlanVlanUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEthernetSwitchIDVlanVlanUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEthernetSwitchIDVlanVlanUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEthernetSwitchIDVlanVlanUUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEthernetSwitchIDVlanVlanUUIDOK creates a GetEthernetSwitchIDVlanVlanUUIDOK with default headers values
func NewGetEthernetSwitchIDVlanVlanUUIDOK() *GetEthernetSwitchIDVlanVlanUUIDOK {
	return &GetEthernetSwitchIDVlanVlanUUIDOK{}
}

/*
GetEthernetSwitchIDVlanVlanUUIDOK describes a response with status code 200, with default header values.

OK
*/
type GetEthernetSwitchIDVlanVlanUUIDOK struct {
	Payload *models.DtosEthernetSwitchVLANDto
}

// IsSuccess returns true when this get ethernet switch Id vlan vlan Uuid o k response has a 2xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get ethernet switch Id vlan vlan Uuid o k response has a 3xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ethernet switch Id vlan vlan Uuid o k response has a 4xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ethernet switch Id vlan vlan Uuid o k response has a 5xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get ethernet switch Id vlan vlan Uuid o k response a status code equal to that given
func (o *GetEthernetSwitchIDVlanVlanUUIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEthernetSwitchIDVlanVlanUUIDOK) Error() string {
	return fmt.Sprintf("[GET /ethernet-switch/{id}/vlan/{vlanUUID}][%d] getEthernetSwitchIdVlanVlanUuidOK  %+v", 200, o.Payload)
}

func (o *GetEthernetSwitchIDVlanVlanUUIDOK) String() string {
	return fmt.Sprintf("[GET /ethernet-switch/{id}/vlan/{vlanUUID}][%d] getEthernetSwitchIdVlanVlanUuidOK  %+v", 200, o.Payload)
}

func (o *GetEthernetSwitchIDVlanVlanUUIDOK) GetPayload() *models.DtosEthernetSwitchVLANDto {
	return o.Payload
}

func (o *GetEthernetSwitchIDVlanVlanUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosEthernetSwitchVLANDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEthernetSwitchIDVlanVlanUUIDNotFound creates a GetEthernetSwitchIDVlanVlanUUIDNotFound with default headers values
func NewGetEthernetSwitchIDVlanVlanUUIDNotFound() *GetEthernetSwitchIDVlanVlanUUIDNotFound {
	return &GetEthernetSwitchIDVlanVlanUUIDNotFound{}
}

/*
GetEthernetSwitchIDVlanVlanUUIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetEthernetSwitchIDVlanVlanUUIDNotFound struct {
}

// IsSuccess returns true when this get ethernet switch Id vlan vlan Uuid not found response has a 2xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get ethernet switch Id vlan vlan Uuid not found response has a 3xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ethernet switch Id vlan vlan Uuid not found response has a 4xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get ethernet switch Id vlan vlan Uuid not found response has a 5xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get ethernet switch Id vlan vlan Uuid not found response a status code equal to that given
func (o *GetEthernetSwitchIDVlanVlanUUIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetEthernetSwitchIDVlanVlanUUIDNotFound) Error() string {
	return fmt.Sprintf("[GET /ethernet-switch/{id}/vlan/{vlanUUID}][%d] getEthernetSwitchIdVlanVlanUuidNotFound ", 404)
}

func (o *GetEthernetSwitchIDVlanVlanUUIDNotFound) String() string {
	return fmt.Sprintf("[GET /ethernet-switch/{id}/vlan/{vlanUUID}][%d] getEthernetSwitchIdVlanVlanUuidNotFound ", 404)
}

func (o *GetEthernetSwitchIDVlanVlanUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEthernetSwitchIDVlanVlanUUIDInternalServerError creates a GetEthernetSwitchIDVlanVlanUUIDInternalServerError with default headers values
func NewGetEthernetSwitchIDVlanVlanUUIDInternalServerError() *GetEthernetSwitchIDVlanVlanUUIDInternalServerError {
	return &GetEthernetSwitchIDVlanVlanUUIDInternalServerError{}
}

/*
GetEthernetSwitchIDVlanVlanUUIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetEthernetSwitchIDVlanVlanUUIDInternalServerError struct {
}

// IsSuccess returns true when this get ethernet switch Id vlan vlan Uuid internal server error response has a 2xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get ethernet switch Id vlan vlan Uuid internal server error response has a 3xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ethernet switch Id vlan vlan Uuid internal server error response has a 4xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ethernet switch Id vlan vlan Uuid internal server error response has a 5xx status code
func (o *GetEthernetSwitchIDVlanVlanUUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get ethernet switch Id vlan vlan Uuid internal server error response a status code equal to that given
func (o *GetEthernetSwitchIDVlanVlanUUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetEthernetSwitchIDVlanVlanUUIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /ethernet-switch/{id}/vlan/{vlanUUID}][%d] getEthernetSwitchIdVlanVlanUuidInternalServerError ", 500)
}

func (o *GetEthernetSwitchIDVlanVlanUUIDInternalServerError) String() string {
	return fmt.Sprintf("[GET /ethernet-switch/{id}/vlan/{vlanUUID}][%d] getEthernetSwitchIdVlanVlanUuidInternalServerError ", 500)
}

func (o *GetEthernetSwitchIDVlanVlanUUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}