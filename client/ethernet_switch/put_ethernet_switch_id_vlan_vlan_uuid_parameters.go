// Code generated by go-swagger; DO NOT EDIT.

package ethernet_switch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/sadov/rolctl/models"
)

// NewPutEthernetSwitchIDVlanVlanUUIDParams creates a new PutEthernetSwitchIDVlanVlanUUIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutEthernetSwitchIDVlanVlanUUIDParams() *PutEthernetSwitchIDVlanVlanUUIDParams {
	return &PutEthernetSwitchIDVlanVlanUUIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutEthernetSwitchIDVlanVlanUUIDParamsWithTimeout creates a new PutEthernetSwitchIDVlanVlanUUIDParams object
// with the ability to set a timeout on a request.
func NewPutEthernetSwitchIDVlanVlanUUIDParamsWithTimeout(timeout time.Duration) *PutEthernetSwitchIDVlanVlanUUIDParams {
	return &PutEthernetSwitchIDVlanVlanUUIDParams{
		timeout: timeout,
	}
}

// NewPutEthernetSwitchIDVlanVlanUUIDParamsWithContext creates a new PutEthernetSwitchIDVlanVlanUUIDParams object
// with the ability to set a context for a request.
func NewPutEthernetSwitchIDVlanVlanUUIDParamsWithContext(ctx context.Context) *PutEthernetSwitchIDVlanVlanUUIDParams {
	return &PutEthernetSwitchIDVlanVlanUUIDParams{
		Context: ctx,
	}
}

// NewPutEthernetSwitchIDVlanVlanUUIDParamsWithHTTPClient creates a new PutEthernetSwitchIDVlanVlanUUIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutEthernetSwitchIDVlanVlanUUIDParamsWithHTTPClient(client *http.Client) *PutEthernetSwitchIDVlanVlanUUIDParams {
	return &PutEthernetSwitchIDVlanVlanUUIDParams{
		HTTPClient: client,
	}
}

/*
PutEthernetSwitchIDVlanVlanUUIDParams contains all the parameters to send to the API endpoint

	for the put ethernet switch ID vlan vlan UUID operation.

	Typically these are written to a http.Request.
*/
type PutEthernetSwitchIDVlanVlanUUIDParams struct {

	/* ID.

	   Ethernet switch ID
	*/
	ID string

	/* Request.

	   Ethernet switch fields
	*/
	Request *models.DtosEthernetSwitchVLANUpdateDto

	/* VlanUUID.

	   Ethernet switch VLAN UUID
	*/
	VlanUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put ethernet switch ID vlan vlan UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) WithDefaults() *PutEthernetSwitchIDVlanVlanUUIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put ethernet switch ID vlan vlan UUID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) WithTimeout(timeout time.Duration) *PutEthernetSwitchIDVlanVlanUUIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) WithContext(ctx context.Context) *PutEthernetSwitchIDVlanVlanUUIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) WithHTTPClient(client *http.Client) *PutEthernetSwitchIDVlanVlanUUIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) WithID(id string) *PutEthernetSwitchIDVlanVlanUUIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) WithRequest(request *models.DtosEthernetSwitchVLANUpdateDto) *PutEthernetSwitchIDVlanVlanUUIDParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) SetRequest(request *models.DtosEthernetSwitchVLANUpdateDto) {
	o.Request = request
}

// WithVlanUUID adds the vlanUUID to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) WithVlanUUID(vlanUUID string) *PutEthernetSwitchIDVlanVlanUUIDParams {
	o.SetVlanUUID(vlanUUID)
	return o
}

// SetVlanUUID adds the vlanUuid to the put ethernet switch ID vlan vlan UUID params
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) SetVlanUUID(vlanUUID string) {
	o.VlanUUID = vlanUUID
}

// WriteToRequest writes these params to a swagger request
func (o *PutEthernetSwitchIDVlanVlanUUIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}
	if o.Request != nil {
		if err := r.SetBodyParam(o.Request); err != nil {
			return err
		}
	}

	// path param vlanUUID
	if err := r.SetPathParam("vlanUUID", o.VlanUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}