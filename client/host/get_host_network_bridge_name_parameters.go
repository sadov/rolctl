// Code generated by go-swagger; DO NOT EDIT.

package host

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

// NewGetHostNetworkBridgeNameParams creates a new GetHostNetworkBridgeNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHostNetworkBridgeNameParams() *GetHostNetworkBridgeNameParams {
	return &GetHostNetworkBridgeNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostNetworkBridgeNameParamsWithTimeout creates a new GetHostNetworkBridgeNameParams object
// with the ability to set a timeout on a request.
func NewGetHostNetworkBridgeNameParamsWithTimeout(timeout time.Duration) *GetHostNetworkBridgeNameParams {
	return &GetHostNetworkBridgeNameParams{
		timeout: timeout,
	}
}

// NewGetHostNetworkBridgeNameParamsWithContext creates a new GetHostNetworkBridgeNameParams object
// with the ability to set a context for a request.
func NewGetHostNetworkBridgeNameParamsWithContext(ctx context.Context) *GetHostNetworkBridgeNameParams {
	return &GetHostNetworkBridgeNameParams{
		Context: ctx,
	}
}

// NewGetHostNetworkBridgeNameParamsWithHTTPClient creates a new GetHostNetworkBridgeNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHostNetworkBridgeNameParamsWithHTTPClient(client *http.Client) *GetHostNetworkBridgeNameParams {
	return &GetHostNetworkBridgeNameParams{
		HTTPClient: client,
	}
}

/*
GetHostNetworkBridgeNameParams contains all the parameters to send to the API endpoint

	for the get host network bridge name operation.

	Typically these are written to a http.Request.
*/
type GetHostNetworkBridgeNameParams struct {

	/* Name.

	   Bridge name
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get host network bridge name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostNetworkBridgeNameParams) WithDefaults() *GetHostNetworkBridgeNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get host network bridge name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostNetworkBridgeNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get host network bridge name params
func (o *GetHostNetworkBridgeNameParams) WithTimeout(timeout time.Duration) *GetHostNetworkBridgeNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host network bridge name params
func (o *GetHostNetworkBridgeNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host network bridge name params
func (o *GetHostNetworkBridgeNameParams) WithContext(ctx context.Context) *GetHostNetworkBridgeNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host network bridge name params
func (o *GetHostNetworkBridgeNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host network bridge name params
func (o *GetHostNetworkBridgeNameParams) WithHTTPClient(client *http.Client) *GetHostNetworkBridgeNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host network bridge name params
func (o *GetHostNetworkBridgeNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get host network bridge name params
func (o *GetHostNetworkBridgeNameParams) WithName(name string) *GetHostNetworkBridgeNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get host network bridge name params
func (o *GetHostNetworkBridgeNameParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostNetworkBridgeNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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