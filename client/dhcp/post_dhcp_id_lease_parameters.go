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

// NewPostDhcpIDLeaseParams creates a new PostDhcpIDLeaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDhcpIDLeaseParams() *PostDhcpIDLeaseParams {
	return &PostDhcpIDLeaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDhcpIDLeaseParamsWithTimeout creates a new PostDhcpIDLeaseParams object
// with the ability to set a timeout on a request.
func NewPostDhcpIDLeaseParamsWithTimeout(timeout time.Duration) *PostDhcpIDLeaseParams {
	return &PostDhcpIDLeaseParams{
		timeout: timeout,
	}
}

// NewPostDhcpIDLeaseParamsWithContext creates a new PostDhcpIDLeaseParams object
// with the ability to set a context for a request.
func NewPostDhcpIDLeaseParamsWithContext(ctx context.Context) *PostDhcpIDLeaseParams {
	return &PostDhcpIDLeaseParams{
		Context: ctx,
	}
}

// NewPostDhcpIDLeaseParamsWithHTTPClient creates a new PostDhcpIDLeaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDhcpIDLeaseParamsWithHTTPClient(client *http.Client) *PostDhcpIDLeaseParams {
	return &PostDhcpIDLeaseParams{
		HTTPClient: client,
	}
}

/*
PostDhcpIDLeaseParams contains all the parameters to send to the API endpoint

	for the post dhcp ID lease operation.

	Typically these are written to a http.Request.
*/
type PostDhcpIDLeaseParams struct {

	/* ID.

	   DHCP v4 server ID
	*/
	ID string

	/* Request.

	   DHCP v4 lease fields
	*/
	Request *models.DtosDHCP4LeaseCreateDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post dhcp ID lease params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDhcpIDLeaseParams) WithDefaults() *PostDhcpIDLeaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post dhcp ID lease params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDhcpIDLeaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) WithTimeout(timeout time.Duration) *PostDhcpIDLeaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) WithContext(ctx context.Context) *PostDhcpIDLeaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) WithHTTPClient(client *http.Client) *PostDhcpIDLeaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) WithID(id string) *PostDhcpIDLeaseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) WithRequest(request *models.DtosDHCP4LeaseCreateDto) *PostDhcpIDLeaseParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post dhcp ID lease params
func (o *PostDhcpIDLeaseParams) SetRequest(request *models.DtosDHCP4LeaseCreateDto) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostDhcpIDLeaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
