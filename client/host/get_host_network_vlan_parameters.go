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

// NewGetHostNetworkVlanParams creates a new GetHostNetworkVlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetHostNetworkVlanParams() *GetHostNetworkVlanParams {
	return &GetHostNetworkVlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetHostNetworkVlanParamsWithTimeout creates a new GetHostNetworkVlanParams object
// with the ability to set a timeout on a request.
func NewGetHostNetworkVlanParamsWithTimeout(timeout time.Duration) *GetHostNetworkVlanParams {
	return &GetHostNetworkVlanParams{
		timeout: timeout,
	}
}

// NewGetHostNetworkVlanParamsWithContext creates a new GetHostNetworkVlanParams object
// with the ability to set a context for a request.
func NewGetHostNetworkVlanParamsWithContext(ctx context.Context) *GetHostNetworkVlanParams {
	return &GetHostNetworkVlanParams{
		Context: ctx,
	}
}

// NewGetHostNetworkVlanParamsWithHTTPClient creates a new GetHostNetworkVlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetHostNetworkVlanParamsWithHTTPClient(client *http.Client) *GetHostNetworkVlanParams {
	return &GetHostNetworkVlanParams{
		HTTPClient: client,
	}
}

/*
GetHostNetworkVlanParams contains all the parameters to send to the API endpoint

	for the get host network vlan operation.

	Typically these are written to a http.Request.
*/
type GetHostNetworkVlanParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get host network vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostNetworkVlanParams) WithDefaults() *GetHostNetworkVlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get host network vlan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetHostNetworkVlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get host network vlan params
func (o *GetHostNetworkVlanParams) WithTimeout(timeout time.Duration) *GetHostNetworkVlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get host network vlan params
func (o *GetHostNetworkVlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get host network vlan params
func (o *GetHostNetworkVlanParams) WithContext(ctx context.Context) *GetHostNetworkVlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get host network vlan params
func (o *GetHostNetworkVlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get host network vlan params
func (o *GetHostNetworkVlanParams) WithHTTPClient(client *http.Client) *GetHostNetworkVlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get host network vlan params
func (o *GetHostNetworkVlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetHostNetworkVlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}