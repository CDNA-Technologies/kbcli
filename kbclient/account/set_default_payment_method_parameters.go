// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewSetDefaultPaymentMethodParams creates a new SetDefaultPaymentMethodParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetDefaultPaymentMethodParams() *SetDefaultPaymentMethodParams {
	return &SetDefaultPaymentMethodParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetDefaultPaymentMethodParamsWithTimeout creates a new SetDefaultPaymentMethodParams object
// with the ability to set a timeout on a request.
func NewSetDefaultPaymentMethodParamsWithTimeout(timeout time.Duration) *SetDefaultPaymentMethodParams {
	return &SetDefaultPaymentMethodParams{
		timeout: timeout,
	}
}

// NewSetDefaultPaymentMethodParamsWithContext creates a new SetDefaultPaymentMethodParams object
// with the ability to set a context for a request.
func NewSetDefaultPaymentMethodParamsWithContext(ctx context.Context) *SetDefaultPaymentMethodParams {
	return &SetDefaultPaymentMethodParams{
		Context: ctx,
	}
}

// NewSetDefaultPaymentMethodParamsWithHTTPClient creates a new SetDefaultPaymentMethodParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetDefaultPaymentMethodParamsWithHTTPClient(client *http.Client) *SetDefaultPaymentMethodParams {
	return &SetDefaultPaymentMethodParams{
		HTTPClient: client,
	}
}

/* SetDefaultPaymentMethodParams contains all the parameters to send to the API endpoint
   for the set default payment method operation.

   Typically these are written to a http.Request.
*/
type SetDefaultPaymentMethodParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	// PayAllUnpaidInvoices.
	PayAllUnpaidInvoices *bool

	// PaymentMethodID.
	//
	// Format: uuid
	PaymentMethodID strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set default payment method params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDefaultPaymentMethodParams) WithDefaults() *SetDefaultPaymentMethodParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set default payment method params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetDefaultPaymentMethodParams) SetDefaults() {
	var (
		payAllUnpaidInvoicesDefault = bool(false)
	)

	val := SetDefaultPaymentMethodParams{
		PayAllUnpaidInvoices: &payAllUnpaidInvoicesDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithTimeout(timeout time.Duration) *SetDefaultPaymentMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithContext(ctx context.Context) *SetDefaultPaymentMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithHTTPClient(client *http.Client) *SetDefaultPaymentMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithXKillbillComment(xKillbillComment *string) *SetDefaultPaymentMethodParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *SetDefaultPaymentMethodParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithXKillbillReason(xKillbillReason *string) *SetDefaultPaymentMethodParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithAccountID(accountID strfmt.UUID) *SetDefaultPaymentMethodParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithPayAllUnpaidInvoices adds the payAllUnpaidInvoices to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithPayAllUnpaidInvoices(payAllUnpaidInvoices *bool) *SetDefaultPaymentMethodParams {
	o.SetPayAllUnpaidInvoices(payAllUnpaidInvoices)
	return o
}

// SetPayAllUnpaidInvoices adds the payAllUnpaidInvoices to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetPayAllUnpaidInvoices(payAllUnpaidInvoices *bool) {
	o.PayAllUnpaidInvoices = payAllUnpaidInvoices
}

// WithPaymentMethodID adds the paymentMethodID to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithPaymentMethodID(paymentMethodID strfmt.UUID) *SetDefaultPaymentMethodParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetPaymentMethodID(paymentMethodID strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WithPluginProperty adds the pluginProperty to the set default payment method params
func (o *SetDefaultPaymentMethodParams) WithPluginProperty(pluginProperty []string) *SetDefaultPaymentMethodParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the set default payment method params
func (o *SetDefaultPaymentMethodParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *SetDefaultPaymentMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if o.PayAllUnpaidInvoices != nil {

		// query param payAllUnpaidInvoices
		var qrPayAllUnpaidInvoices bool

		if o.PayAllUnpaidInvoices != nil {
			qrPayAllUnpaidInvoices = *o.PayAllUnpaidInvoices
		}
		qPayAllUnpaidInvoices := swag.FormatBool(qrPayAllUnpaidInvoices)
		if qPayAllUnpaidInvoices != "" {

			if err := r.SetQueryParam("payAllUnpaidInvoices", qPayAllUnpaidInvoices); err != nil {
				return err
			}
		}
	}

	// path param paymentMethodId
	if err := r.SetPathParam("paymentMethodId", o.PaymentMethodID.String()); err != nil {
		return err
	}

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSetDefaultPaymentMethod binds the parameter pluginProperty
func (o *SetDefaultPaymentMethodParams) bindParamPluginProperty(formats strfmt.Registry) []string {
	pluginPropertyIR := o.PluginProperty

	var pluginPropertyIC []string
	for _, pluginPropertyIIR := range pluginPropertyIR { // explode []string

		pluginPropertyIIV := pluginPropertyIIR // string as string
		pluginPropertyIC = append(pluginPropertyIC, pluginPropertyIIV)
	}

	// items.CollectionFormat: "multi"
	pluginPropertyIS := swag.JoinByFormat(pluginPropertyIC, "multi")

	return pluginPropertyIS
}
