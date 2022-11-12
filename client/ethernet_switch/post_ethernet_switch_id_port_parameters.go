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

	"github.com/sadov/rolctl/models"
)

// NewPostEthernetSwitchIDPortParams creates a new PostEthernetSwitchIDPortParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostEthernetSwitchIDPortParams() *PostEthernetSwitchIDPortParams {
	return &PostEthernetSwitchIDPortParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostEthernetSwitchIDPortParamsWithTimeout creates a new PostEthernetSwitchIDPortParams object
// with the ability to set a timeout on a request.
func NewPostEthernetSwitchIDPortParamsWithTimeout(timeout time.Duration) *PostEthernetSwitchIDPortParams {
	return &PostEthernetSwitchIDPortParams{
		timeout: timeout,
	}
}

// NewPostEthernetSwitchIDPortParamsWithContext creates a new PostEthernetSwitchIDPortParams object
// with the ability to set a context for a request.
func NewPostEthernetSwitchIDPortParamsWithContext(ctx context.Context) *PostEthernetSwitchIDPortParams {
	return &PostEthernetSwitchIDPortParams{
		Context: ctx,
	}
}

// NewPostEthernetSwitchIDPortParamsWithHTTPClient creates a new PostEthernetSwitchIDPortParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostEthernetSwitchIDPortParamsWithHTTPClient(client *http.Client) *PostEthernetSwitchIDPortParams {
	return &PostEthernetSwitchIDPortParams{
		HTTPClient: client,
	}
}

/*
PostEthernetSwitchIDPortParams contains all the parameters to send to the API endpoint

	for the post ethernet switch ID port operation.

	Typically these are written to a http.Request.
*/
type PostEthernetSwitchIDPortParams struct {

	/* ID.

	   Ethernet switch ID
	*/
	ID string

	/* Request.

	   Ethernet switch port fields
	*/
	Request *models.DtosEthernetSwitchPortCreateDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post ethernet switch ID port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEthernetSwitchIDPortParams) WithDefaults() *PostEthernetSwitchIDPortParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post ethernet switch ID port params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostEthernetSwitchIDPortParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) WithTimeout(timeout time.Duration) *PostEthernetSwitchIDPortParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) WithContext(ctx context.Context) *PostEthernetSwitchIDPortParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) WithHTTPClient(client *http.Client) *PostEthernetSwitchIDPortParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) WithID(id string) *PostEthernetSwitchIDPortParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) SetID(id string) {
	o.ID = id
}

// WithRequest adds the request to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) WithRequest(request *models.DtosEthernetSwitchPortCreateDto) *PostEthernetSwitchIDPortParams {
	o.SetRequest(request)
	return o
}

// SetRequest adds the request to the post ethernet switch ID port params
func (o *PostEthernetSwitchIDPortParams) SetRequest(request *models.DtosEthernetSwitchPortCreateDto) {
	o.Request = request
}

// WriteToRequest writes these params to a swagger request
func (o *PostEthernetSwitchIDPortParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
