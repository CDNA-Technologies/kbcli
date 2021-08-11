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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInvoiceCustomFieldsParams creates a new GetInvoiceCustomFieldsParams object
// with the default values initialized.
func NewGetInvoiceCustomFieldsParams() *GetInvoiceCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetInvoiceCustomFieldsParams{
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceCustomFieldsParamsWithTimeout creates a new GetInvoiceCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoiceCustomFieldsParamsWithTimeout(timeout time.Duration) *GetInvoiceCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetInvoiceCustomFieldsParams{
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetInvoiceCustomFieldsParamsWithContext creates a new GetInvoiceCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoiceCustomFieldsParamsWithContext(ctx context.Context) *GetInvoiceCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetInvoiceCustomFieldsParams{
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetInvoiceCustomFieldsParamsWithHTTPClient creates a new GetInvoiceCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoiceCustomFieldsParamsWithHTTPClient(client *http.Client) *GetInvoiceCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetInvoiceCustomFieldsParams{
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetInvoiceCustomFieldsParams contains all the parameters to send to the API endpoint
for the get invoice custom fields operation typically these are written to a http.Request
*/
type GetInvoiceCustomFieldsParams struct {

	/*Audit*/
	Audit *string
	/*InvoiceID*/
	InvoiceID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) WithTimeout(timeout time.Duration) *GetInvoiceCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) WithContext(ctx context.Context) *GetInvoiceCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) WithHTTPClient(client *http.Client) *GetInvoiceCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) WithAudit(audit *string) *GetInvoiceCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithInvoiceID adds the invoiceID to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) WithInvoiceID(invoiceID strfmt.UUID) *GetInvoiceCustomFieldsParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the get invoice custom fields params
func (o *GetInvoiceCustomFieldsParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
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
