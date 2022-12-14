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
	"github.com/go-openapi/swag"
)

// NewGetTemplateDeviceParams creates a new GetTemplateDeviceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTemplateDeviceParams() *GetTemplateDeviceParams {
	return &GetTemplateDeviceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTemplateDeviceParamsWithTimeout creates a new GetTemplateDeviceParams object
// with the ability to set a timeout on a request.
func NewGetTemplateDeviceParamsWithTimeout(timeout time.Duration) *GetTemplateDeviceParams {
	return &GetTemplateDeviceParams{
		timeout: timeout,
	}
}

// NewGetTemplateDeviceParamsWithContext creates a new GetTemplateDeviceParams object
// with the ability to set a context for a request.
func NewGetTemplateDeviceParamsWithContext(ctx context.Context) *GetTemplateDeviceParams {
	return &GetTemplateDeviceParams{
		Context: ctx,
	}
}

// NewGetTemplateDeviceParamsWithHTTPClient creates a new GetTemplateDeviceParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTemplateDeviceParamsWithHTTPClient(client *http.Client) *GetTemplateDeviceParams {
	return &GetTemplateDeviceParams{
		HTTPClient: client,
	}
}

/*
GetTemplateDeviceParams contains all the parameters to send to the API endpoint

	for the get template device operation.

	Typically these are written to a http.Request.
*/
type GetTemplateDeviceParams struct {

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

// WithDefaults hydrates default values in the get template device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTemplateDeviceParams) WithDefaults() *GetTemplateDeviceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get template device params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTemplateDeviceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get template device params
func (o *GetTemplateDeviceParams) WithTimeout(timeout time.Duration) *GetTemplateDeviceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get template device params
func (o *GetTemplateDeviceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get template device params
func (o *GetTemplateDeviceParams) WithContext(ctx context.Context) *GetTemplateDeviceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get template device params
func (o *GetTemplateDeviceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get template device params
func (o *GetTemplateDeviceParams) WithHTTPClient(client *http.Client) *GetTemplateDeviceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get template device params
func (o *GetTemplateDeviceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrderBy adds the orderBy to the get template device params
func (o *GetTemplateDeviceParams) WithOrderBy(orderBy *string) *GetTemplateDeviceParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the get template device params
func (o *GetTemplateDeviceParams) SetOrderBy(orderBy *string) {
	o.OrderBy = orderBy
}

// WithOrderDirection adds the orderDirection to the get template device params
func (o *GetTemplateDeviceParams) WithOrderDirection(orderDirection *string) *GetTemplateDeviceParams {
	o.SetOrderDirection(orderDirection)
	return o
}

// SetOrderDirection adds the orderDirection to the get template device params
func (o *GetTemplateDeviceParams) SetOrderDirection(orderDirection *string) {
	o.OrderDirection = orderDirection
}

// WithPage adds the page to the get template device params
func (o *GetTemplateDeviceParams) WithPage(page *int64) *GetTemplateDeviceParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get template device params
func (o *GetTemplateDeviceParams) SetPage(page *int64) {
	o.Page = page
}

// WithPageSize adds the pageSize to the get template device params
func (o *GetTemplateDeviceParams) WithPageSize(pageSize *int64) *GetTemplateDeviceParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get template device params
func (o *GetTemplateDeviceParams) SetPageSize(pageSize *int64) {
	o.PageSize = pageSize
}

// WithSearch adds the search to the get template device params
func (o *GetTemplateDeviceParams) WithSearch(search *string) *GetTemplateDeviceParams {
	o.SetSearch(search)
	return o
}

// SetSearch adds the search to the get template device params
func (o *GetTemplateDeviceParams) SetSearch(search *string) {
	o.Search = search
}

// WriteToRequest writes these params to a swagger request
func (o *GetTemplateDeviceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
