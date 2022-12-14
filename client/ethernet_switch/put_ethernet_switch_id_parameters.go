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

// NewPutEthernetSwitchIDParams creates a new PutEthernetSwitchIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutEthernetSwitchIDParams() *PutEthernetSwitchIDParams {
	return &PutEthernetSwitchIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutEthernetSwitchIDParamsWithTimeout creates a new PutEthernetSwitchIDParams object
// with the ability to set a timeout on a request.
func NewPutEthernetSwitchIDParamsWithTimeout(timeout time.Duration) *PutEthernetSwitchIDParams {
	return &PutEthernetSwitchIDParams{
		timeout: timeout,
	}
}

// NewPutEthernetSwitchIDParamsWithContext creates a new PutEthernetSwitchIDParams object
// with the ability to set a context for a request.
func NewPutEthernetSwitchIDParamsWithContext(ctx context.Context) *PutEthernetSwitchIDParams {
	return &PutEthernetSwitchIDParams{
		Context: ctx,
	}
}

// NewPutEthernetSwitchIDParamsWithHTTPClient creates a new PutEthernetSwitchIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutEthernetSwitchIDParamsWithHTTPClient(client *http.Client) *PutEthernetSwitchIDParams {
	return &PutEthernetSwitchIDParams{
		HTTPClient: client,
	}
}

/*
PutEthernetSwitchIDParams contains all the parameters to send to the API endpoint

	for the put ethernet switch ID operation.

	Typically these are written to a http.Request.
*/
type PutEthernetSwitchIDParams struct {

	/* ID.

	   Ethernet switch ID
	*/
	ID string

	/* Request.

	   Ethernet switch fields
	*/
	Request *models.DtosEthernetSwitchUpdateDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put ethernet switch ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEthernetSwitchIDParams) WithDefaults() *PutEthernetSwitchIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put ethernet switch ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutEthernetSwitchIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) WithTimeout(timeout time.Duration) *PutEthernetSwitchIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) WithContext(ctx context.Context) *PutEthernetSwitchIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) WithHTTPClient(client *http.Client) *PutEthernetSwitchIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) WithID(id string) *PutEthernetSwitchIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) WithRequest(request *models.DtosEthernetSwitchUpdateDto) *PutEthernetSwitchIDParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put ethernet switch ID params
func (o *PutEthernetSwitchIDParams) SetRequest(request *models.DtosEthernetSwitchUpdateDto) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PutEthernetSwitchIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
