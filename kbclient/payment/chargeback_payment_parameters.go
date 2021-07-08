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

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// NewChargebackPaymentParams creates a new ChargebackPaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChargebackPaymentParams() *ChargebackPaymentParams {
	return &ChargebackPaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChargebackPaymentParamsWithTimeout creates a new ChargebackPaymentParams object
// with the ability to set a timeout on a request.
func NewChargebackPaymentParamsWithTimeout(timeout time.Duration) *ChargebackPaymentParams {
	return &ChargebackPaymentParams{
		timeout: timeout,
	}
}

// NewChargebackPaymentParamsWithContext creates a new ChargebackPaymentParams object
// with the ability to set a context for a request.
func NewChargebackPaymentParamsWithContext(ctx context.Context) *ChargebackPaymentParams {
	return &ChargebackPaymentParams{
		Context: ctx,
	}
}

// NewChargebackPaymentParamsWithHTTPClient creates a new ChargebackPaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewChargebackPaymentParamsWithHTTPClient(client *http.Client) *ChargebackPaymentParams {
	return &ChargebackPaymentParams{
		HTTPClient: client,
	}
}

/* ChargebackPaymentParams contains all the parameters to send to the API endpoint
   for the chargeback payment operation.

   Typically these are written to a http.Request.
*/
type ChargebackPaymentParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.PaymentTransaction

	// ControlPluginName.
	ControlPluginName []string

	// PaymentID.
	//
	// Format: uuid
	PaymentID strfmt.UUID

	// PluginProperty.
	PluginProperty []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chargeback payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargebackPaymentParams) WithDefaults() *ChargebackPaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chargeback payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargebackPaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the chargeback payment params
func (o *ChargebackPaymentParams) WithTimeout(timeout time.Duration) *ChargebackPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chargeback payment params
func (o *ChargebackPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chargeback payment params
func (o *ChargebackPaymentParams) WithContext(ctx context.Context) *ChargebackPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chargeback payment params
func (o *ChargebackPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chargeback payment params
func (o *ChargebackPaymentParams) WithHTTPClient(client *http.Client) *ChargebackPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chargeback payment params
func (o *ChargebackPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the chargeback payment params
func (o *ChargebackPaymentParams) WithXKillbillComment(xKillbillComment *string) *ChargebackPaymentParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the chargeback payment params
func (o *ChargebackPaymentParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback payment params
func (o *ChargebackPaymentParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ChargebackPaymentParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback payment params
func (o *ChargebackPaymentParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the chargeback payment params
func (o *ChargebackPaymentParams) WithXKillbillReason(xKillbillReason *string) *ChargebackPaymentParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the chargeback payment params
func (o *ChargebackPaymentParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the chargeback payment params
func (o *ChargebackPaymentParams) WithBody(body *kbmodel.PaymentTransaction) *ChargebackPaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the chargeback payment params
func (o *ChargebackPaymentParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the chargeback payment params
func (o *ChargebackPaymentParams) WithControlPluginName(controlPluginName []string) *ChargebackPaymentParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the chargeback payment params
func (o *ChargebackPaymentParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPaymentID adds the paymentID to the chargeback payment params
func (o *ChargebackPaymentParams) WithPaymentID(paymentID strfmt.UUID) *ChargebackPaymentParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the chargeback payment params
func (o *ChargebackPaymentParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the chargeback payment params
func (o *ChargebackPaymentParams) WithPluginProperty(pluginProperty []string) *ChargebackPaymentParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the chargeback payment params
func (o *ChargebackPaymentParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *ChargebackPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ControlPluginName != nil {

		// binding items for controlPluginName
		joinedControlPluginName := o.bindParamControlPluginName(reg)

		// query array param controlPluginName
		if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
			return err
		}
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
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

// bindParamChargebackPayment binds the parameter controlPluginName
func (o *ChargebackPaymentParams) bindParamControlPluginName(formats strfmt.Registry) []string {
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

// bindParamChargebackPayment binds the parameter pluginProperty
func (o *ChargebackPaymentParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
