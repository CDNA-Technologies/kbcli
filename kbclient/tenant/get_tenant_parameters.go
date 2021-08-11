// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTenantParams creates a new GetTenantParams object
// with the default values initialized.
func NewGetTenantParams() *GetTenantParams {
	var ()
	return &GetTenantParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTenantParamsWithTimeout creates a new GetTenantParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTenantParamsWithTimeout(timeout time.Duration) *GetTenantParams {
	var ()
	return &GetTenantParams{

		timeout: timeout,
	}
}

// NewGetTenantParamsWithContext creates a new GetTenantParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTenantParamsWithContext(ctx context.Context) *GetTenantParams {
	var ()
	return &GetTenantParams{

		Context: ctx,
	}
}

// NewGetTenantParamsWithHTTPClient creates a new GetTenantParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTenantParamsWithHTTPClient(client *http.Client) *GetTenantParams {
	var ()
	return &GetTenantParams{
		HTTPClient: client,
	}
}

/*GetTenantParams contains all the parameters to send to the API endpoint
for the get tenant operation typically these are written to a http.Request
*/
type GetTenantParams struct {

	/*TenantID*/
	TenantID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get tenant params
func (o *GetTenantParams) WithTimeout(timeout time.Duration) *GetTenantParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tenant params
func (o *GetTenantParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tenant params
func (o *GetTenantParams) WithContext(ctx context.Context) *GetTenantParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tenant params
func (o *GetTenantParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tenant params
func (o *GetTenantParams) WithHTTPClient(client *http.Client) *GetTenantParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tenant params
func (o *GetTenantParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTenantID adds the tenantID to the get tenant params
func (o *GetTenantParams) WithTenantID(tenantID strfmt.UUID) *GetTenantParams {
	o.SetTenantID(tenantID)
	return o
}

// SetTenantID adds the tenantId to the get tenant params
func (o *GetTenantParams) SetTenantID(tenantID strfmt.UUID) {
	o.TenantID = tenantID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTenantParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param tenantId
	if err := r.SetPathParam("tenantId", o.TenantID.String()); err != nil {
		return err
	}

	// header param WithProfilingInfo
	if o.WithProfilingInfo != nil && len(*o.WithProfilingInfo) > 0 {
		if err := r.SetHeaderParam("X-Killbill-Profiling-Req", *o.WithProfilingInfo); err != nil {
			return err
		}
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
