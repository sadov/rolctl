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

// NewPutDhcpIDLeaseLeaseIDParams creates a new PutDhcpIDLeaseLeaseIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutDhcpIDLeaseLeaseIDParams() *PutDhcpIDLeaseLeaseIDParams {
	return &PutDhcpIDLeaseLeaseIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutDhcpIDLeaseLeaseIDParamsWithTimeout creates a new PutDhcpIDLeaseLeaseIDParams object
// with the ability to set a timeout on a request.
func NewPutDhcpIDLeaseLeaseIDParamsWithTimeout(timeout time.Duration) *PutDhcpIDLeaseLeaseIDParams {
	return &PutDhcpIDLeaseLeaseIDParams{
		timeout: timeout,
	}
}

// NewPutDhcpIDLeaseLeaseIDParamsWithContext creates a new PutDhcpIDLeaseLeaseIDParams object
// with the ability to set a context for a request.
func NewPutDhcpIDLeaseLeaseIDParamsWithContext(ctx context.Context) *PutDhcpIDLeaseLeaseIDParams {
	return &PutDhcpIDLeaseLeaseIDParams{
		Context: ctx,
	}
}

// NewPutDhcpIDLeaseLeaseIDParamsWithHTTPClient creates a new PutDhcpIDLeaseLeaseIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutDhcpIDLeaseLeaseIDParamsWithHTTPClient(client *http.Client) *PutDhcpIDLeaseLeaseIDParams {
	return &PutDhcpIDLeaseLeaseIDParams{
		HTTPClient: client,
	}
}

/*
PutDhcpIDLeaseLeaseIDParams contains all the parameters to send to the API endpoint

	for the put dhcp ID lease lease ID operation.

	Typically these are written to a http.Request.
*/
type PutDhcpIDLeaseLeaseIDParams struct {

	/* ID.

	   DHCP v4 server ID
	*/
	ID string

	/* LeaseID.

	   DHCP v4 lease ID
	*/
	LeaseID string

	/* Request.

	   DHCP v4 lease fields
	*/
	Request *models.DtosDHCP4LeaseUpdateDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put dhcp ID lease lease ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDhcpIDLeaseLeaseIDParams) WithDefaults() *PutDhcpIDLeaseLeaseIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put dhcp ID lease lease ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutDhcpIDLeaseLeaseIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) WithTimeout(timeout time.Duration) *PutDhcpIDLeaseLeaseIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) WithContext(ctx context.Context) *PutDhcpIDLeaseLeaseIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) WithHTTPClient(client *http.Client) *PutDhcpIDLeaseLeaseIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) WithID(id string) *PutDhcpIDLeaseLeaseIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) SetID(id string) {
	o.ID = id
}

// WithLeaseID adds the leaseID to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) WithLeaseID(leaseID string) *PutDhcpIDLeaseLeaseIDParams {
	o.SetLeaseID(leaseID)
	return o
}

// SetLeaseID adds the leaseId to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) SetLeaseID(leaseID string) {
	o.LeaseID = leaseID
}

// WithRequest adds the request to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) WithRequest(request *models.DtosDHCP4LeaseUpdateDto) *PutDhcpIDLeaseLeaseIDParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the put dhcp ID lease lease ID params
func (o *PutDhcpIDLeaseLeaseIDParams) SetRequest(request *models.DtosDHCP4LeaseUpdateDto) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PutDhcpIDLeaseLeaseIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param leaseID
	if err := r.SetPathParam("leaseID", o.LeaseID); err != nil {
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
