// Code generated by go-swagger; DO NOT EDIT.

package dhcp

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

// NewPutDhcpIDParams creates a new PutDhcpIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDhcpIDParams() *PutDhcpIDParams {
	return &PutDhcpIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDhcpIDParamsWithTimeout creates a new PutDhcpIDParams object
// with the ability to set a timeout on a request.
func NewPutDhcpIDParamsWithTimeout(timeout time.Duration) *PutDhcpIDParams {
	return &PutDhcpIDParams{
		timeout: timeout,
	}
}

// NewPutDhcpIDParamsWithContext creates a new PutDhcpIDParams object
// with the ability to set a context for a request.
func NewPutDhcpIDParamsWithContext(ctx context.Context) *PutDhcpIDParams {
	return &PutDhcpIDParams{
		Context: ctx,
	}
}

// NewPutDhcpIDParamsWithHTTPClient creates a new PutDhcpIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDhcpIDParamsWithHTTPClient(client *http.Client) *PutDhcpIDParams {
	return &PutDhcpIDParams{
		HTTPClient: client,
	}
}

/*
PutDhcpIDParams contains all the parameters to send to the API endpoint

	for the put dhcp ID operation.

	Typically these are written to a http.Request.
*/
type PutDhcpIDParams struct {

	/* ID.

	   DHCP v4 server ID
	*/
	ID string

	/* Request.

	   DHCP v4 server fields
	*/
	Request *models.DtosDHCP4ServerUpdateDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put dhcp ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDhcpIDParams) WithDefaults() *PutDhcpIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put dhcp ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDhcpIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put dhcp ID params
func (o *PutDhcpIDParams) WithTimeout(timeout time.Duration) *PutDhcpIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put dhcp ID params
func (o *PutDhcpIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put dhcp ID params
func (o *PutDhcpIDParams) WithContext(ctx context.Context) *PutDhcpIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put dhcp ID params
func (o *PutDhcpIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put dhcp ID params
func (o *PutDhcpIDParams) WithHTTPClient(client *http.Client) *PutDhcpIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put dhcp ID params
func (o *PutDhcpIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put dhcp ID params
func (o *PutDhcpIDParams) WithID(id string) *PutDhcpIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put dhcp ID params
func (o *PutDhcpIDParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the put dhcp ID params
func (o *PutDhcpIDParams) WithRequest(request *models.DtosDHCP4ServerUpdateDto) *PutDhcpIDParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put dhcp ID params
func (o *PutDhcpIDParams) SetRequest(request *models.DtosDHCP4ServerUpdateDto) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PutDhcpIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
