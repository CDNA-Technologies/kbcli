// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

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

	"github.com/killbill/kbcli/v2/kbmodel"
)

// NewModifyInvoiceItemCustomFieldsParams creates a new ModifyInvoiceItemCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyInvoiceItemCustomFieldsParams() *ModifyInvoiceItemCustomFieldsParams {
	return &ModifyInvoiceItemCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyInvoiceItemCustomFieldsParamsWithTimeout creates a new ModifyInvoiceItemCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewModifyInvoiceItemCustomFieldsParamsWithTimeout(timeout time.Duration) *ModifyInvoiceItemCustomFieldsParams {
	return &ModifyInvoiceItemCustomFieldsParams{
		timeout: timeout,
	}
}

// NewModifyInvoiceItemCustomFieldsParamsWithContext creates a new ModifyInvoiceItemCustomFieldsParams object
// with the ability to set a context for a request.
func NewModifyInvoiceItemCustomFieldsParamsWithContext(ctx context.Context) *ModifyInvoiceItemCustomFieldsParams {
	return &ModifyInvoiceItemCustomFieldsParams{
		Context: ctx,
	}
}

// NewModifyInvoiceItemCustomFieldsParamsWithHTTPClient creates a new ModifyInvoiceItemCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyInvoiceItemCustomFieldsParamsWithHTTPClient(client *http.Client) *ModifyInvoiceItemCustomFieldsParams {
	return &ModifyInvoiceItemCustomFieldsParams{
		HTTPClient: client,
	}
}

/* ModifyInvoiceItemCustomFieldsParams contains all the parameters to send to the API endpoint
   for the modify invoice item custom fields operation.

   Typically these are written to a http.Request.
*/
type ModifyInvoiceItemCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body []*kbmodel.CustomField

	// InvoiceItemID.
	//
	// Format: uuid
	InvoiceItemID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the modify invoice item custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyInvoiceItemCustomFieldsParams) WithDefaults() *ModifyInvoiceItemCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify invoice item custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyInvoiceItemCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) WithTimeout(timeout time.Duration) *ModifyInvoiceItemCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) WithContext(ctx context.Context) *ModifyInvoiceItemCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) WithHTTPClient(client *http.Client) *ModifyInvoiceItemCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *ModifyInvoiceItemCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ModifyInvoiceItemCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *ModifyInvoiceItemCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *ModifyInvoiceItemCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithInvoiceItemID adds the invoiceItemID to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) WithInvoiceItemID(invoiceItemID strfmt.UUID) *ModifyInvoiceItemCustomFieldsParams {
	o.SetInvoiceItemID(invoiceItemID)
	return o
}

// SetInvoiceItemID adds the invoiceItemId to the modify invoice item custom fields params
func (o *ModifyInvoiceItemCustomFieldsParams) SetInvoiceItemID(invoiceItemID strfmt.UUID) {
	o.InvoiceItemID = invoiceItemID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyInvoiceItemCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}
	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param invoiceItemId
	if err := r.SetPathParam("invoiceItemId", o.InvoiceItemID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
