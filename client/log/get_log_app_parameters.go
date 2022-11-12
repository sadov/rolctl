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
	"github.com/go-openapi/swag"
)

// NewGetLogAppParams creates a new GetLogAppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLogAppParams() *GetLogAppParams {
	return &GetLogAppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLogAppParamsWithTimeout creates a new GetLogAppParams object
// with the ability to set a timeout on a request.
func NewGetLogAppParamsWithTimeout(timeout time.Duration) *GetLogAppParams {
	return &GetLogAppParams{
		timeout: timeout,
	}
}

// NewGetLogAppParamsWithContext creates a new GetLogAppParams object
// with the ability to set a context for a request.
func NewGetLogAppParamsWithContext(ctx context.Context) *GetLogAppParams {
	return &GetLogAppParams{
		Context: ctx,
	}
}

// NewGetLogAppParamsWithHTTPClient creates a new GetLogAppParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLogAppParamsWithHTTPClient(client *http.Client) *GetLogAppParams {
	return &GetLogAppParams{
		HTTPClient: client,
	}
}

/*
GetLogAppParams contains all the parameters to send to the API endpoint

	for the get log app operation.

	Typically these are written to a http.Request.
*/
type GetLogAppParams struct {

	/* OrderBy.

	   Order by field
	*/
	OrderBy *string

	/* OrderDirection.

	   'asc' or 'desc' for ascending or descending order
	*/
	OrderDirection *string

	/* Page.

	   page number
	*/
	Page *int64

	/* PageSize.

	   number of entities per page
	*/
	PageSize *int64

	/* Search.

	   searchable value in entity
	*/
	Search *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get log app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogAppParams) WithDefaults() *GetLogAppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get log app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLogAppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get log app params
func (o *GetLogAppParams) WithTimeout(timeout time.Duration) *GetLogAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get log app params
func (o *GetLogAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get log app params
func (o *GetLogAppParams) WithContext(ctx context.Context) *GetLogAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get log app params
func (o *GetLogAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get log app params
func (o *GetLogAppParams) WithHTTPClient(client *http.Client) *GetLogAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get log app params
func (o *GetLogAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the get log app params
func (o *GetLogAppParams) WithOrderBy(orderBy *string) *GetLogAppParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get log app params
func (o *GetLogAppParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrderDirection adds the orderDirection to the get log app params
func (o *GetLogAppParams) WithOrderDirection(orderDirection *string) *GetLogAppParams {
	o.SetOrderDirection(orderDirection)
	return o
}

// SetOrderDirection adds the orderDirection to the get log app params
func (o *GetLogAppParams) SetOrderDirection(orderDirection *string) {
	o.OrderDirection = orderDirection
}

// WithPage adds the page to the get log app params
func (o *GetLogAppParams) WithPage(page *int64) *GetLogAppParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get log app params
func (o *GetLogAppParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get log app params
func (o *GetLogAppParams) WithPageSize(pageSize *int64) *GetLogAppParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get log app params
func (o *GetLogAppParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the get log app params
func (o *GetLogAppParams) WithSearch(search *string) *GetLogAppParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get log app params
func (o *GetLogAppParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetLogAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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