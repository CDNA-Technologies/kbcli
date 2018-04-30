// Code generated by go-swagger; DO NOT EDIT.

package overdue

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

// NewGetOverdueConfigJSONParams creates a new GetOverdueConfigJSONParams object
// with the default values initialized.
func NewGetOverdueConfigJSONParams() *GetOverdueConfigJSONParams {
	var ()
	return &GetOverdueConfigJSONParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOverdueConfigJSONParamsWithTimeout creates a new GetOverdueConfigJSONParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOverdueConfigJSONParamsWithTimeout(timeout time.Duration) *GetOverdueConfigJSONParams {
	var ()
	return &GetOverdueConfigJSONParams{

		timeout: timeout,
	}
}

// NewGetOverdueConfigJSONParamsWithContext creates a new GetOverdueConfigJSONParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOverdueConfigJSONParamsWithContext(ctx context.Context) *GetOverdueConfigJSONParams {
	var ()
	return &GetOverdueConfigJSONParams{

		Context: ctx,
	}
}

// NewGetOverdueConfigJSONParamsWithHTTPClient creates a new GetOverdueConfigJSONParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOverdueConfigJSONParamsWithHTTPClient(client *http.Client) *GetOverdueConfigJSONParams {
	var ()
	return &GetOverdueConfigJSONParams{
		HTTPClient: client,
	}
}

/*GetOverdueConfigJSONParams contains all the parameters to send to the API endpoint
for the get overdue config Json operation typically these are written to a http.Request
*/
type GetOverdueConfigJSONParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) WithTimeout(timeout time.Duration) *GetOverdueConfigJSONParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) WithContext(ctx context.Context) *GetOverdueConfigJSONParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) WithHTTPClient(client *http.Client) *GetOverdueConfigJSONParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetOverdueConfigJSONParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetOverdueConfigJSONParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get overdue config Json params
func (o *GetOverdueConfigJSONParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WriteToRequest writes these params to a swagger request
func (o *GetOverdueConfigJSONParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
