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

// NewGetHostNetworkBridgeParams creates a new GetHostNetworkBridgeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHostNetworkBridgeParams() *GetHostNetworkBridgeParams {
	return &GetHostNetworkBridgeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostNetworkBridgeParamsWithTimeout creates a new GetHostNetworkBridgeParams object
// with the ability to set a timeout on a request.
func NewGetHostNetworkBridgeParamsWithTimeout(timeout time.Duration) *GetHostNetworkBridgeParams {
	return &GetHostNetworkBridgeParams{
		timeout: timeout,
	}
}

// NewGetHostNetworkBridgeParamsWithContext creates a new GetHostNetworkBridgeParams object
// with the ability to set a context for a request.
func NewGetHostNetworkBridgeParamsWithContext(ctx context.Context) *GetHostNetworkBridgeParams {
	return &GetHostNetworkBridgeParams{
		Context: ctx,
	}
}

// NewGetHostNetworkBridgeParamsWithHTTPClient creates a new GetHostNetworkBridgeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHostNetworkBridgeParamsWithHTTPClient(client *http.Client) *GetHostNetworkBridgeParams {
	return &GetHostNetworkBridgeParams{
		HTTPClient: client,
	}
}

/*
GetHostNetworkBridgeParams contains all the parameters to send to the API endpoint

	for the get host network bridge operation.

	Typically these are written to a http.Request.
*/
type GetHostNetworkBridgeParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get host network bridge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostNetworkBridgeParams) WithDefaults() *GetHostNetworkBridgeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get host network bridge params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostNetworkBridgeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get host network bridge params
func (o *GetHostNetworkBridgeParams) WithTimeout(timeout time.Duration) *GetHostNetworkBridgeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host network bridge params
func (o *GetHostNetworkBridgeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host network bridge params
func (o *GetHostNetworkBridgeParams) WithContext(ctx context.Context) *GetHostNetworkBridgeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host network bridge params
func (o *GetHostNetworkBridgeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host network bridge params
func (o *GetHostNetworkBridgeParams) WithHTTPClient(client *http.Client) *GetHostNetworkBridgeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host network bridge params
func (o *GetHostNetworkBridgeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostNetworkBridgeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
