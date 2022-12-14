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
	"github.com/go-openapi/swag"
)

// NewGetDhcpIDLeaseParams creates a new GetDhcpIDLeaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetDhcpIDLeaseParams() *GetDhcpIDLeaseParams {
	return &GetDhcpIDLeaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetDhcpIDLeaseParamsWithTimeout creates a new GetDhcpIDLeaseParams object
// with the ability to set a timeout on a request.
func NewGetDhcpIDLeaseParamsWithTimeout(timeout time.Duration) *GetDhcpIDLeaseParams {
	return &GetDhcpIDLeaseParams{
		timeout: timeout,
	}
}

// NewGetDhcpIDLeaseParamsWithContext creates a new GetDhcpIDLeaseParams object
// with the ability to set a context for a request.
func NewGetDhcpIDLeaseParamsWithContext(ctx context.Context) *GetDhcpIDLeaseParams {
	return &GetDhcpIDLeaseParams{
		Context: ctx,
	}
}

// NewGetDhcpIDLeaseParamsWithHTTPClient creates a new GetDhcpIDLeaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetDhcpIDLeaseParamsWithHTTPClient(client *http.Client) *GetDhcpIDLeaseParams {
	return &GetDhcpIDLeaseParams{
		HTTPClient: client,
	}
}

/*
GetDhcpIDLeaseParams contains all the parameters to send to the API endpoint

	for the get dhcp ID lease operation.

	Typically these are written to a http.Request.
*/
type GetDhcpIDLeaseParams struct {

	/* ID.

	   DHCP v4 server ID
	*/
	ID string

	/* OrderBy.

	   Order by field
	*/
	OrderBy *string

	/* OrderDirection.

	   'asc' or 'desc' for ascending or descending order
	*/
	OrderDirection *string

	/* Page.

	   Page number
	*/
	Page *int64

	/* PageSize.

	   Number of entities per page
	*/
	PageSize *int64

	/* Search.

	   Searchable value in entity
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get dhcp ID lease params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDhcpIDLeaseParams) WithDefaults() *GetDhcpIDLeaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get dhcp ID lease params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetDhcpIDLeaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) WithTimeout(timeout time.Duration) *GetDhcpIDLeaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) WithContext(ctx context.Context) *GetDhcpIDLeaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) WithHTTPClient(client *http.Client) *GetDhcpIDLeaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) WithID(id string) *GetDhcpIDLeaseParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) SetID(id string) {
	o.ID = id
}

// WithOrderBy adds the orderBy to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) WithOrderBy(orderBy *string) *GetDhcpIDLeaseParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrderDirection adds the orderDirection to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) WithOrderDirection(orderDirection *string) *GetDhcpIDLeaseParams {
	o.SetOrderDirection(orderDirection)
	return o
}

// SetOrderDirection adds the orderDirection to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) SetOrderDirection(orderDirection *string) {
	o.OrderDirection = orderDirection
}

// WithPage adds the page to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) WithPage(page *int64) *GetDhcpIDLeaseParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) WithPageSize(pageSize *int64) *GetDhcpIDLeaseParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) WithSearch(search *string) *GetDhcpIDLeaseParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get dhcp ID lease params
func (o *GetDhcpIDLeaseParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetDhcpIDLeaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.OrderBy != nil {

		// query param orderBy
		var qrOrderBy string

		if o.OrderBy != nil {
			qrOrderBy = *o.OrderBy
		}
		qOrderBy := qrOrderBy
		if qOrderBy != "" {

			if err := r.SetQueryParam("orderBy", qOrderBy); err != nil {
				return err
			}
		}
	}

	if o.OrderDirection != nil {

		// query param orderDirection
		var qrOrderDirection string

		if o.OrderDirection != nil {
			qrOrderDirection = *o.OrderDirection
		}
		qOrderDirection := qrOrderDirection
		if qOrderDirection != "" {

			if err := r.SetQueryParam("orderDirection", qOrderDirection); err != nil {
				return err
			}
		}
	}

	if o.Page != nil {

		// query param page
		var qrPage int64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int64

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt64(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.Search != nil {

		// query param search
		var qrSearch string

		if o.Search != nil {
			qrSearch = *o.Search
		}
		qSearch := qrSearch
		if qSearch != "" {

			if err := r.SetQueryParam("search", qSearch); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
