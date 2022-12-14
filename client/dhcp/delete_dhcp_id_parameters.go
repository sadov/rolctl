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
)

// NewDeleteDhcpIDParams creates a new DeleteDhcpIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteDhcpIDParams() *DeleteDhcpIDParams {
	return &DeleteDhcpIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDhcpIDParamsWithTimeout creates a new DeleteDhcpIDParams object
// with the ability to set a timeout on a request.
func NewDeleteDhcpIDParamsWithTimeout(timeout time.Duration) *DeleteDhcpIDParams {
	return &DeleteDhcpIDParams{
		timeout: timeout,
	}
}

// NewDeleteDhcpIDParamsWithContext creates a new DeleteDhcpIDParams object
// with the ability to set a context for a request.
func NewDeleteDhcpIDParamsWithContext(ctx context.Context) *DeleteDhcpIDParams {
	return &DeleteDhcpIDParams{
		Context: ctx,
	}
}

// NewDeleteDhcpIDParamsWithHTTPClient creates a new DeleteDhcpIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteDhcpIDParamsWithHTTPClient(client *http.Client) *DeleteDhcpIDParams {
	return &DeleteDhcpIDParams{
		HTTPClient: client,
	}
}

/*
DeleteDhcpIDParams contains all the parameters to send to the API endpoint

	for the delete dhcp ID operation.

	Typically these are written to a http.Request.
*/
type DeleteDhcpIDParams struct {

	/* ID.

	   DHCP v4 server ID
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete dhcp ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDhcpIDParams) WithDefaults() *DeleteDhcpIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete dhcp ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteDhcpIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete dhcp ID params
func (o *DeleteDhcpIDParams) WithTimeout(timeout time.Duration) *DeleteDhcpIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete dhcp ID params
func (o *DeleteDhcpIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete dhcp ID params
func (o *DeleteDhcpIDParams) WithContext(ctx context.Context) *DeleteDhcpIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete dhcp ID params
func (o *DeleteDhcpIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete dhcp ID params
func (o *DeleteDhcpIDParams) WithHTTPClient(client *http.Client) *DeleteDhcpIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete dhcp ID params
func (o *DeleteDhcpIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete dhcp ID params
func (o *DeleteDhcpIDParams) WithID(id string) *DeleteDhcpIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete dhcp ID params
func (o *DeleteDhcpIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDhcpIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
