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

// PutEthernetSwitchIDVlanVlanUUIDReader is a Reader for the PutEthernetSwitchIDVlanVlanUUID structure.
type PutEthernetSwitchIDVlanVlanUUIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutEthernetSwitchIDVlanVlanUUIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutEthernetSwitchIDVlanVlanUUIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutEthernetSwitchIDVlanVlanUUIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutEthernetSwitchIDVlanVlanUUIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutEthernetSwitchIDVlanVlanUUIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutEthernetSwitchIDVlanVlanUUIDOK creates a PutEthernetSwitchIDVlanVlanUUIDOK with default headers values
func NewPutEthernetSwitchIDVlanVlanUUIDOK() *PutEthernetSwitchIDVlanVlanUUIDOK {
	return &PutEthernetSwitchIDVlanVlanUUIDOK{}
}

/*
PutEthernetSwitchIDVlanVlanUUIDOK describes a response with status code 200, with default header values.

OK
*/
type PutEthernetSwitchIDVlanVlanUUIDOK struct {
	Payload *models.DtosEthernetSwitchVLANDto
}

// IsSuccess returns true when this put ethernet switch Id vlan vlan Uuid o k response has a 2xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put ethernet switch Id vlan vlan Uuid o k response has a 3xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put ethernet switch Id vlan vlan Uuid o k response has a 4xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put ethernet switch Id vlan vlan Uuid o k response has a 5xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put ethernet switch Id vlan vlan Uuid o k response a status code equal to that given
func (o *PutEthernetSwitchIDVlanVlanUUIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutEthernetSwitchIDVlanVlanUUIDOK) Error() string {
	return fmt.Sprintf("[PUT /ethernet-switch/{id}/vlan/{vlanUUID}][%d] putEthernetSwitchIdVlanVlanUuidOK  %+v", 200, o.Payload)
}

func (o *PutEthernetSwitchIDVlanVlanUUIDOK) String() string {
	return fmt.Sprintf("[PUT /ethernet-switch/{id}/vlan/{vlanUUID}][%d] putEthernetSwitchIdVlanVlanUuidOK  %+v", 200, o.Payload)
}

func (o *PutEthernetSwitchIDVlanVlanUUIDOK) GetPayload() *models.DtosEthernetSwitchVLANDto {
	return o.Payload
}

func (o *PutEthernetSwitchIDVlanVlanUUIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosEthernetSwitchVLANDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEthernetSwitchIDVlanVlanUUIDBadRequest creates a PutEthernetSwitchIDVlanVlanUUIDBadRequest with default headers values
func NewPutEthernetSwitchIDVlanVlanUUIDBadRequest() *PutEthernetSwitchIDVlanVlanUUIDBadRequest {
	return &PutEthernetSwitchIDVlanVlanUUIDBadRequest{}
}

/*
PutEthernetSwitchIDVlanVlanUUIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PutEthernetSwitchIDVlanVlanUUIDBadRequest struct {
	Payload *models.DtosValidationErrorDto
}

// IsSuccess returns true when this put ethernet switch Id vlan vlan Uuid bad request response has a 2xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put ethernet switch Id vlan vlan Uuid bad request response has a 3xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put ethernet switch Id vlan vlan Uuid bad request response has a 4xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put ethernet switch Id vlan vlan Uuid bad request response has a 5xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put ethernet switch Id vlan vlan Uuid bad request response a status code equal to that given
func (o *PutEthernetSwitchIDVlanVlanUUIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutEthernetSwitchIDVlanVlanUUIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ethernet-switch/{id}/vlan/{vlanUUID}][%d] putEthernetSwitchIdVlanVlanUuidBadRequest  %+v", 400, o.Payload)
}

func (o *PutEthernetSwitchIDVlanVlanUUIDBadRequest) String() string {
	return fmt.Sprintf("[PUT /ethernet-switch/{id}/vlan/{vlanUUID}][%d] putEthernetSwitchIdVlanVlanUuidBadRequest  %+v", 400, o.Payload)
}

func (o *PutEthernetSwitchIDVlanVlanUUIDBadRequest) GetPayload() *models.DtosValidationErrorDto {
	return o.Payload
}

func (o *PutEthernetSwitchIDVlanVlanUUIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DtosValidationErrorDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutEthernetSwitchIDVlanVlanUUIDNotFound creates a PutEthernetSwitchIDVlanVlanUUIDNotFound with default headers values
func NewPutEthernetSwitchIDVlanVlanUUIDNotFound() *PutEthernetSwitchIDVlanVlanUUIDNotFound {
	return &PutEthernetSwitchIDVlanVlanUUIDNotFound{}
}

/*
PutEthernetSwitchIDVlanVlanUUIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PutEthernetSwitchIDVlanVlanUUIDNotFound struct {
}

// IsSuccess returns true when this put ethernet switch Id vlan vlan Uuid not found response has a 2xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put ethernet switch Id vlan vlan Uuid not found response has a 3xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put ethernet switch Id vlan vlan Uuid not found response has a 4xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put ethernet switch Id vlan vlan Uuid not found response has a 5xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put ethernet switch Id vlan vlan Uuid not found response a status code equal to that given
func (o *PutEthernetSwitchIDVlanVlanUUIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutEthernetSwitchIDVlanVlanUUIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /ethernet-switch/{id}/vlan/{vlanUUID}][%d] putEthernetSwitchIdVlanVlanUuidNotFound ", 404)
}

func (o *PutEthernetSwitchIDVlanVlanUUIDNotFound) String() string {
	return fmt.Sprintf("[PUT /ethernet-switch/{id}/vlan/{vlanUUID}][%d] putEthernetSwitchIdVlanVlanUuidNotFound ", 404)
}

func (o *PutEthernetSwitchIDVlanVlanUUIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutEthernetSwitchIDVlanVlanUUIDInternalServerError creates a PutEthernetSwitchIDVlanVlanUUIDInternalServerError with default headers values
func NewPutEthernetSwitchIDVlanVlanUUIDInternalServerError() *PutEthernetSwitchIDVlanVlanUUIDInternalServerError {
	return &PutEthernetSwitchIDVlanVlanUUIDInternalServerError{}
}

/*
PutEthernetSwitchIDVlanVlanUUIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PutEthernetSwitchIDVlanVlanUUIDInternalServerError struct {
}

// IsSuccess returns true when this put ethernet switch Id vlan vlan Uuid internal server error response has a 2xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put ethernet switch Id vlan vlan Uuid internal server error response has a 3xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put ethernet switch Id vlan vlan Uuid internal server error response has a 4xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put ethernet switch Id vlan vlan Uuid internal server error response has a 5xx status code
func (o *PutEthernetSwitchIDVlanVlanUUIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put ethernet switch Id vlan vlan Uuid internal server error response a status code equal to that given
func (o *PutEthernetSwitchIDVlanVlanUUIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutEthernetSwitchIDVlanVlanUUIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /ethernet-switch/{id}/vlan/{vlanUUID}][%d] putEthernetSwitchIdVlanVlanUuidInternalServerError ", 500)
}

func (o *PutEthernetSwitchIDVlanVlanUUIDInternalServerError) String() string {
	return fmt.Sprintf("[PUT /ethernet-switch/{id}/vlan/{vlanUUID}][%d] putEthernetSwitchIdVlanVlanUuidInternalServerError ", 500)
}

func (o *PutEthernetSwitchIDVlanVlanUUIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}