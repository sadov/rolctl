// Code generated by go-swagger; DO NOT EDIT.

package log

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

// NewGetLogAppIDParams creates a new GetLogAppIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogAppIDParams() *GetLogAppIDParams {
	return &GetLogAppIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogAppIDParamsWithTimeout creates a new GetLogAppIDParams object
// with the ability to set a timeout on a request.
func NewGetLogAppIDParamsWithTimeout(timeout time.Duration) *GetLogAppIDParams {
	return &GetLogAppIDParams{
		timeout: timeout,
	}
}

// NewGetLogAppIDParamsWithContext creates a new GetLogAppIDParams object
// with the ability to set a context for a request.
func NewGetLogAppIDParamsWithContext(ctx context.Context) *GetLogAppIDParams {
	return &GetLogAppIDParams{
		Context: ctx,
	}
}

// NewGetLogAppIDParamsWithHTTPClient creates a new GetLogAppIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogAppIDParamsWithHTTPClient(client *http.Client) *GetLogAppIDParams {
	return &GetLogAppIDParams{
		HTTPClient: client,
	}
}

/*
GetLogAppIDParams contains all the parameters to send to the API endpoint

	for the get log app ID operation.

	Typically these are written to a http.Request.
*/
type GetLogAppIDParams struct {

	/* ID.

	   log id
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log app ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogAppIDParams) WithDefaults() *GetLogAppIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log app ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogAppIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get log app ID params
func (o *GetLogAppIDParams) WithTimeout(timeout time.Duration) *GetLogAppIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log app ID params
func (o *GetLogAppIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log app ID params
func (o *GetLogAppIDParams) WithContext(ctx context.Context) *GetLogAppIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log app ID params
func (o *GetLogAppIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log app ID params
func (o *GetLogAppIDParams) WithHTTPClient(client *http.Client) *GetLogAppIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log app ID params
func (o *GetLogAppIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get log app ID params
func (o *GetLogAppIDParams) WithID(id string) *GetLogAppIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get log app ID params
func (o *GetLogAppIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogAppIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
