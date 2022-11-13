// Code generated by go-swagger; DO NOT EDIT.

package template

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

// NewGetTemplateDeviceNameParams creates a new GetTemplateDeviceNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTemplateDeviceNameParams() *GetTemplateDeviceNameParams {
	return &GetTemplateDeviceNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTemplateDeviceNameParamsWithTimeout creates a new GetTemplateDeviceNameParams object
// with the ability to set a timeout on a request.
func NewGetTemplateDeviceNameParamsWithTimeout(timeout time.Duration) *GetTemplateDeviceNameParams {
	return &GetTemplateDeviceNameParams{
		timeout: timeout,
	}
}

// NewGetTemplateDeviceNameParamsWithContext creates a new GetTemplateDeviceNameParams object
// with the ability to set a context for a request.
func NewGetTemplateDeviceNameParamsWithContext(ctx context.Context) *GetTemplateDeviceNameParams {
	return &GetTemplateDeviceNameParams{
		Context: ctx,
	}
}

// NewGetTemplateDeviceNameParamsWithHTTPClient creates a new GetTemplateDeviceNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTemplateDeviceNameParamsWithHTTPClient(client *http.Client) *GetTemplateDeviceNameParams {
	return &GetTemplateDeviceNameParams{
		HTTPClient: client,
	}
}

/*
GetTemplateDeviceNameParams contains all the parameters to send to the API endpoint

	for the get template device name operation.

	Typically these are written to a http.Request.
*/
type GetTemplateDeviceNameParams struct {

	/* Name.

	   device template name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get template device name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTemplateDeviceNameParams) WithDefaults() *GetTemplateDeviceNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get template device name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTemplateDeviceNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get template device name params
func (o *GetTemplateDeviceNameParams) WithTimeout(timeout time.Duration) *GetTemplateDeviceNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get template device name params
func (o *GetTemplateDeviceNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get template device name params
func (o *GetTemplateDeviceNameParams) WithContext(ctx context.Context) *GetTemplateDeviceNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get template device name params
func (o *GetTemplateDeviceNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get template device name params
func (o *GetTemplateDeviceNameParams) WithHTTPClient(client *http.Client) *GetTemplateDeviceNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get template device name params
func (o *GetTemplateDeviceNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get template device name params
func (o *GetTemplateDeviceNameParams) WithName(name string) *GetTemplateDeviceNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get template device name params
func (o *GetTemplateDeviceNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetTemplateDeviceNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
