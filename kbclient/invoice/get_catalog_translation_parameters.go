// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewGetCatalogTranslationParams creates a new GetCatalogTranslationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCatalogTranslationParams() *GetCatalogTranslationParams {
	return &GetCatalogTranslationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCatalogTranslationParamsWithTimeout creates a new GetCatalogTranslationParams object
// with the ability to set a timeout on a request.
func NewGetCatalogTranslationParamsWithTimeout(timeout time.Duration) *GetCatalogTranslationParams {
	return &GetCatalogTranslationParams{
		timeout: timeout,
	}
}

// NewGetCatalogTranslationParamsWithContext creates a new GetCatalogTranslationParams object
// with the ability to set a context for a request.
func NewGetCatalogTranslationParamsWithContext(ctx context.Context) *GetCatalogTranslationParams {
	return &GetCatalogTranslationParams{
		Context: ctx,
	}
}

// NewGetCatalogTranslationParamsWithHTTPClient creates a new GetCatalogTranslationParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCatalogTranslationParamsWithHTTPClient(client *http.Client) *GetCatalogTranslationParams {
	return &GetCatalogTranslationParams{
		HTTPClient: client,
	}
}

/* GetCatalogTranslationParams contains all the parameters to send to the API endpoint
   for the get catalog translation operation.

   Typically these are written to a http.Request.
*/
type GetCatalogTranslationParams struct {

	// Locale.
	Locale string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get catalog translation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogTranslationParams) WithDefaults() *GetCatalogTranslationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get catalog translation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCatalogTranslationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get catalog translation params
func (o *GetCatalogTranslationParams) WithTimeout(timeout time.Duration) *GetCatalogTranslationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get catalog translation params
func (o *GetCatalogTranslationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get catalog translation params
func (o *GetCatalogTranslationParams) WithContext(ctx context.Context) *GetCatalogTranslationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get catalog translation params
func (o *GetCatalogTranslationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get catalog translation params
func (o *GetCatalogTranslationParams) WithHTTPClient(client *http.Client) *GetCatalogTranslationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get catalog translation params
func (o *GetCatalogTranslationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocale adds the locale to the get catalog translation params
func (o *GetCatalogTranslationParams) WithLocale(locale string) *GetCatalogTranslationParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get catalog translation params
func (o *GetCatalogTranslationParams) SetLocale(locale string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogTranslationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param locale
	if err := r.SetPathParam("locale", o.Locale); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
