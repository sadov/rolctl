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

// GetEthernetSwitchModelsReader is a Reader for the GetEthernetSwitchModels structure.
type GetEthernetSwitchModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEthernetSwitchModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEthernetSwitchModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEthernetSwitchModelsOK creates a GetEthernetSwitchModelsOK with default headers values
func NewGetEthernetSwitchModelsOK() *GetEthernetSwitchModelsOK {
	return &GetEthernetSwitchModelsOK{}
}

/*
GetEthernetSwitchModelsOK describes a response with status code 200, with default header values.

OK
*/
type GetEthernetSwitchModelsOK struct {
	Payload []*models.DtosEthernetSwitchModelDto
}

// IsSuccess returns true when this get ethernet switch models o k response has a 2xx status code
func (o *GetEthernetSwitchModelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get ethernet switch models o k response has a 3xx status code
func (o *GetEthernetSwitchModelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get ethernet switch models o k response has a 4xx status code
func (o *GetEthernetSwitchModelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get ethernet switch models o k response has a 5xx status code
func (o *GetEthernetSwitchModelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get ethernet switch models o k response a status code equal to that given
func (o *GetEthernetSwitchModelsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEthernetSwitchModelsOK) Error() string {
	return fmt.Sprintf("[GET /ethernet-switch/models][%d] getEthernetSwitchModelsOK  %+v", 200, o.Payload)
}

func (o *GetEthernetSwitchModelsOK) String() string {
	return fmt.Sprintf("[GET /ethernet-switch/models][%d] getEthernetSwitchModelsOK  %+v", 200, o.Payload)
}

func (o *GetEthernetSwitchModelsOK) GetPayload() []*models.DtosEthernetSwitchModelDto {
	return o.Payload
}

func (o *GetEthernetSwitchModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}