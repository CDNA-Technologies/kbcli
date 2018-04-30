// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewInvalidatesCacheByTenantParams creates a new InvalidatesCacheByTenantParams object
// with the default values initialized.
func NewInvalidatesCacheByTenantParams() *InvalidatesCacheByTenantParams {
	var ()
	return &InvalidatesCacheByTenantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInvalidatesCacheByTenantParamsWithTimeout creates a new InvalidatesCacheByTenantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInvalidatesCacheByTenantParamsWithTimeout(timeout time.Duration) *InvalidatesCacheByTenantParams {
	var ()
	return &InvalidatesCacheByTenantParams{

		timeout: timeout,
	}
}

// NewInvalidatesCacheByTenantParamsWithContext creates a new InvalidatesCacheByTenantParams object
// with the default values initialized, and the ability to set a context for a request
func NewInvalidatesCacheByTenantParamsWithContext(ctx context.Context) *InvalidatesCacheByTenantParams {
	var ()
	return &InvalidatesCacheByTenantParams{

		Context: ctx,
	}
}

// NewInvalidatesCacheByTenantParamsWithHTTPClient creates a new InvalidatesCacheByTenantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInvalidatesCacheByTenantParamsWithHTTPClient(client *http.Client) *InvalidatesCacheByTenantParams {
	var ()
	return &InvalidatesCacheByTenantParams{
		HTTPClient: client,
	}
}

/*InvalidatesCacheByTenantParams contains all the parameters to send to the API endpoint
for the invalidates cache by tenant operation typically these are written to a http.Request
*/
type InvalidatesCacheByTenantParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) WithTimeout(timeout time.Duration) *InvalidatesCacheByTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) WithContext(ctx context.Context) *InvalidatesCacheByTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) WithHTTPClient(client *http.Client) *InvalidatesCacheByTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) WithXKillbillAPIKey(xKillbillAPIKey string) *InvalidatesCacheByTenantParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) WithXKillbillAPISecret(xKillbillAPISecret string) *InvalidatesCacheByTenantParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the invalidates cache by tenant params
func (o *InvalidatesCacheByTenantParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WriteToRequest writes these params to a swagger request
func (o *InvalidatesCacheByTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
