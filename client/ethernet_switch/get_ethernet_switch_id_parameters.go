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
)

// NewGetEthernetSwitchIDParams creates a new GetEthernetSwitchIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEthernetSwitchIDParams() *GetEthernetSwitchIDParams {
	return &GetEthernetSwitchIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEthernetSwitchIDParamsWithTimeout creates a new GetEthernetSwitchIDParams object
// with the ability to set a timeout on a request.
func NewGetEthernetSwitchIDParamsWithTimeout(timeout time.Duration) *GetEthernetSwitchIDParams {
	return &GetEthernetSwitchIDParams{
		timeout: timeout,
	}
}

// NewGetEthernetSwitchIDParamsWithContext creates a new GetEthernetSwitchIDParams object
// with the ability to set a context for a request.
func NewGetEthernetSwitchIDParamsWithContext(ctx context.Context) *GetEthernetSwitchIDParams {
	return &GetEthernetSwitchIDParams{
		Context: ctx,
	}
}

// NewGetEthernetSwitchIDParamsWithHTTPClient creates a new GetEthernetSwitchIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEthernetSwitchIDParamsWithHTTPClient(client *http.Client) *GetEthernetSwitchIDParams {
	return &GetEthernetSwitchIDParams{
		HTTPClient: client,
	}
}

/*
GetEthernetSwitchIDParams contains all the parameters to send to the API endpoint

	for the get ethernet switch ID operation.

	Typically these are written to a http.Request.
*/
type GetEthernetSwitchIDParams struct {

	/* ID.

	   Ethernet switch ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get ethernet switch ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEthernetSwitchIDParams) WithDefaults() *GetEthernetSwitchIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get ethernet switch ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEthernetSwitchIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get ethernet switch ID params
func (o *GetEthernetSwitchIDParams) WithTimeout(timeout time.Duration) *GetEthernetSwitchIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get ethernet switch ID params
func (o *GetEthernetSwitchIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get ethernet switch ID params
func (o *GetEthernetSwitchIDParams) WithContext(ctx context.Context) *GetEthernetSwitchIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get ethernet switch ID params
func (o *GetEthernetSwitchIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get ethernet switch ID params
func (o *GetEthernetSwitchIDParams) WithHTTPClient(client *http.Client) *GetEthernetSwitchIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get ethernet switch ID params
func (o *GetEthernetSwitchIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get ethernet switch ID params
func (o *GetEthernetSwitchIDParams) WithID(id string) *GetEthernetSwitchIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get ethernet switch ID params
func (o *GetEthernetSwitchIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEthernetSwitchIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}