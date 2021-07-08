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

	"github.com/killbill/kbcli/v2/kbmodel"
)

// NewModifyPaymentCustomFieldsParams creates a new ModifyPaymentCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewModifyPaymentCustomFieldsParams() *ModifyPaymentCustomFieldsParams {
	return &ModifyPaymentCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewModifyPaymentCustomFieldsParamsWithTimeout creates a new ModifyPaymentCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewModifyPaymentCustomFieldsParamsWithTimeout(timeout time.Duration) *ModifyPaymentCustomFieldsParams {
	return &ModifyPaymentCustomFieldsParams{
		timeout: timeout,
	}
}

// NewModifyPaymentCustomFieldsParamsWithContext creates a new ModifyPaymentCustomFieldsParams object
// with the ability to set a context for a request.
func NewModifyPaymentCustomFieldsParamsWithContext(ctx context.Context) *ModifyPaymentCustomFieldsParams {
	return &ModifyPaymentCustomFieldsParams{
		Context: ctx,
	}
}

// NewModifyPaymentCustomFieldsParamsWithHTTPClient creates a new ModifyPaymentCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewModifyPaymentCustomFieldsParamsWithHTTPClient(client *http.Client) *ModifyPaymentCustomFieldsParams {
	return &ModifyPaymentCustomFieldsParams{
		HTTPClient: client,
	}
}

/* ModifyPaymentCustomFieldsParams contains all the parameters to send to the API endpoint
   for the modify payment custom fields operation.

   Typically these are written to a http.Request.
*/
type ModifyPaymentCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body []*kbmodel.CustomField

	// PaymentID.
	//
	// Format: uuid
	PaymentID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the modify payment custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyPaymentCustomFieldsParams) WithDefaults() *ModifyPaymentCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the modify payment custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ModifyPaymentCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) WithTimeout(timeout time.Duration) *ModifyPaymentCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) WithContext(ctx context.Context) *ModifyPaymentCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) WithHTTPClient(client *http.Client) *ModifyPaymentCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *ModifyPaymentCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ModifyPaymentCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *ModifyPaymentCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *ModifyPaymentCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithPaymentID adds the paymentID to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) WithPaymentID(paymentID strfmt.UUID) *ModifyPaymentCustomFieldsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the modify payment custom fields params
func (o *ModifyPaymentCustomFieldsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyPaymentCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
