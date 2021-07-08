// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// NewGetPaymentTagsParams creates a new GetPaymentTagsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetPaymentTagsParams() *GetPaymentTagsParams {
	return &GetPaymentTagsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentTagsParamsWithTimeout creates a new GetPaymentTagsParams object
// with the ability to set a timeout on a request.
func NewGetPaymentTagsParamsWithTimeout(timeout time.Duration) *GetPaymentTagsParams {
	return &GetPaymentTagsParams{
		timeout: timeout,
	}
}

// NewGetPaymentTagsParamsWithContext creates a new GetPaymentTagsParams object
// with the ability to set a context for a request.
func NewGetPaymentTagsParamsWithContext(ctx context.Context) *GetPaymentTagsParams {
	return &GetPaymentTagsParams{
		Context: ctx,
	}
}

// NewGetPaymentTagsParamsWithHTTPClient creates a new GetPaymentTagsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetPaymentTagsParamsWithHTTPClient(client *http.Client) *GetPaymentTagsParams {
	return &GetPaymentTagsParams{
		HTTPClient: client,
	}
}

/* GetPaymentTagsParams contains all the parameters to send to the API endpoint
   for the get payment tags operation.

   Typically these are written to a http.Request.
*/
type GetPaymentTagsParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// IncludedDeleted.
	IncludedDeleted *bool

	// PaymentID.
	//
	// Format: uuid
	PaymentID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get payment tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentTagsParams) WithDefaults() *GetPaymentTagsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get payment tags params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetPaymentTagsParams) SetDefaults() {
	var (
		auditDefault = string("NONE")

		includedDeletedDefault = bool(false)
	)

	val := GetPaymentTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get payment tags params
func (o *GetPaymentTagsParams) WithTimeout(timeout time.Duration) *GetPaymentTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment tags params
func (o *GetPaymentTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment tags params
func (o *GetPaymentTagsParams) WithContext(ctx context.Context) *GetPaymentTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment tags params
func (o *GetPaymentTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment tags params
func (o *GetPaymentTagsParams) WithHTTPClient(client *http.Client) *GetPaymentTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment tags params
func (o *GetPaymentTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get payment tags params
func (o *GetPaymentTagsParams) WithAudit(audit *string) *GetPaymentTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payment tags params
func (o *GetPaymentTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithIncludedDeleted adds the includedDeleted to the get payment tags params
func (o *GetPaymentTagsParams) WithIncludedDeleted(includedDeleted *bool) *GetPaymentTagsParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get payment tags params
func (o *GetPaymentTagsParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WithPaymentID adds the paymentID to the get payment tags params
func (o *GetPaymentTagsParams) WithPaymentID(paymentID strfmt.UUID) *GetPaymentTagsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the get payment tags params
func (o *GetPaymentTagsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IncludedDeleted != nil {

		// query param includedDeleted
		var qrIncludedDeleted bool

		if o.IncludedDeleted != nil {
			qrIncludedDeleted = *o.IncludedDeleted
		}
		qIncludedDeleted := swag.FormatBool(qrIncludedDeleted)
		if qIncludedDeleted != "" {

			if err := r.SetQueryParam("includedDeleted", qIncludedDeleted); err != nil {
				return err
			}
		}
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
