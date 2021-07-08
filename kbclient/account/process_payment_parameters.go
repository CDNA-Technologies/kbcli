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

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// NewProcessPaymentParams creates a new ProcessPaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProcessPaymentParams() *ProcessPaymentParams {
	return &ProcessPaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewProcessPaymentParamsWithTimeout creates a new ProcessPaymentParams object
// with the ability to set a timeout on a request.
func NewProcessPaymentParamsWithTimeout(timeout time.Duration) *ProcessPaymentParams {
	return &ProcessPaymentParams{
		timeout: timeout,
	}
}

// NewProcessPaymentParamsWithContext creates a new ProcessPaymentParams object
// with the ability to set a context for a request.
func NewProcessPaymentParamsWithContext(ctx context.Context) *ProcessPaymentParams {
	return &ProcessPaymentParams{
		Context: ctx,
	}
}

// NewProcessPaymentParamsWithHTTPClient creates a new ProcessPaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewProcessPaymentParamsWithHTTPClient(client *http.Client) *ProcessPaymentParams {
	return &ProcessPaymentParams{
		HTTPClient: client,
	}
}

/* ProcessPaymentParams contains all the parameters to send to the API endpoint
   for the process payment operation.

   Typically these are written to a http.Request.
*/
type ProcessPaymentParams struct {

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

	// Body.
	Body *kbmodel.PaymentTransaction

	// ControlPluginName.
	ControlPluginName []string

	// PaymentMethodID.
	//
	// Format: uuid
	PaymentMethodID *strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the process payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProcessPaymentParams) WithDefaults() *ProcessPaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the process payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProcessPaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the process payment params
func (o *ProcessPaymentParams) WithTimeout(timeout time.Duration) *ProcessPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the process payment params
func (o *ProcessPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the process payment params
func (o *ProcessPaymentParams) WithContext(ctx context.Context) *ProcessPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the process payment params
func (o *ProcessPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the process payment params
func (o *ProcessPaymentParams) WithHTTPClient(client *http.Client) *ProcessPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the process payment params
func (o *ProcessPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the process payment params
func (o *ProcessPaymentParams) WithXKillbillComment(xKillbillComment *string) *ProcessPaymentParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the process payment params
func (o *ProcessPaymentParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the process payment params
func (o *ProcessPaymentParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ProcessPaymentParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the process payment params
func (o *ProcessPaymentParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the process payment params
func (o *ProcessPaymentParams) WithXKillbillReason(xKillbillReason *string) *ProcessPaymentParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the process payment params
func (o *ProcessPaymentParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the process payment params
func (o *ProcessPaymentParams) WithAccountID(accountID strfmt.UUID) *ProcessPaymentParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the process payment params
func (o *ProcessPaymentParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the process payment params
func (o *ProcessPaymentParams) WithBody(body *kbmodel.PaymentTransaction) *ProcessPaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the process payment params
func (o *ProcessPaymentParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the process payment params
func (o *ProcessPaymentParams) WithControlPluginName(controlPluginName []string) *ProcessPaymentParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the process payment params
func (o *ProcessPaymentParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPaymentMethodID adds the paymentMethodID to the process payment params
func (o *ProcessPaymentParams) WithPaymentMethodID(paymentMethodID *strfmt.UUID) *ProcessPaymentParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the process payment params
func (o *ProcessPaymentParams) SetPaymentMethodID(paymentMethodID *strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WithPluginProperty adds the pluginProperty to the process payment params
func (o *ProcessPaymentParams) WithPluginProperty(pluginProperty []string) *ProcessPaymentParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the process payment params
func (o *ProcessPaymentParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *ProcessPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ControlPluginName != nil {

		// binding items for controlPluginName
		joinedControlPluginName := o.bindParamControlPluginName(reg)

		// query array param controlPluginName
		if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
			return err
		}
	}

	if o.PaymentMethodID != nil {

		// query param paymentMethodId
		var qrPaymentMethodID strfmt.UUID

		if o.PaymentMethodID != nil {
			qrPaymentMethodID = *o.PaymentMethodID
		}
		qPaymentMethodID := qrPaymentMethodID.String()
		if qPaymentMethodID != "" {

			if err := r.SetQueryParam("paymentMethodId", qPaymentMethodID); err != nil {
				return err
			}
		}
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

// bindParamProcessPayment binds the parameter controlPluginName
func (o *ProcessPaymentParams) bindParamControlPluginName(formats strfmt.Registry) []string {
	controlPluginNameIR := o.ControlPluginName

	var controlPluginNameIC []string
	for _, controlPluginNameIIR := range controlPluginNameIR { // explode []string

		controlPluginNameIIV := controlPluginNameIIR // string as string
		controlPluginNameIC = append(controlPluginNameIC, controlPluginNameIIV)
	}

	// items.CollectionFormat: "multi"
	controlPluginNameIS := swag.JoinByFormat(controlPluginNameIC, "multi")

	return controlPluginNameIS
}

// bindParamProcessPayment binds the parameter pluginProperty
func (o *ProcessPaymentParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
