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

// NewChargebackPaymentByExternalKeyParams creates a new ChargebackPaymentByExternalKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChargebackPaymentByExternalKeyParams() *ChargebackPaymentByExternalKeyParams {
	return &ChargebackPaymentByExternalKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChargebackPaymentByExternalKeyParamsWithTimeout creates a new ChargebackPaymentByExternalKeyParams object
// with the ability to set a timeout on a request.
func NewChargebackPaymentByExternalKeyParamsWithTimeout(timeout time.Duration) *ChargebackPaymentByExternalKeyParams {
	return &ChargebackPaymentByExternalKeyParams{
		timeout: timeout,
	}
}

// NewChargebackPaymentByExternalKeyParamsWithContext creates a new ChargebackPaymentByExternalKeyParams object
// with the ability to set a context for a request.
func NewChargebackPaymentByExternalKeyParamsWithContext(ctx context.Context) *ChargebackPaymentByExternalKeyParams {
	return &ChargebackPaymentByExternalKeyParams{
		Context: ctx,
	}
}

// NewChargebackPaymentByExternalKeyParamsWithHTTPClient creates a new ChargebackPaymentByExternalKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewChargebackPaymentByExternalKeyParamsWithHTTPClient(client *http.Client) *ChargebackPaymentByExternalKeyParams {
	return &ChargebackPaymentByExternalKeyParams{
		HTTPClient: client,
	}
}

/* ChargebackPaymentByExternalKeyParams contains all the parameters to send to the API endpoint
   for the chargeback payment by external key operation.

   Typically these are written to a http.Request.
*/
type ChargebackPaymentByExternalKeyParams struct {

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

	// PluginProperty.
	PluginProperty []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chargeback payment by external key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargebackPaymentByExternalKeyParams) WithDefaults() *ChargebackPaymentByExternalKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chargeback payment by external key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChargebackPaymentByExternalKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) WithTimeout(timeout time.Duration) *ChargebackPaymentByExternalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) WithContext(ctx context.Context) *ChargebackPaymentByExternalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) WithHTTPClient(client *http.Client) *ChargebackPaymentByExternalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) WithXKillbillComment(xKillbillComment *string) *ChargebackPaymentByExternalKeyParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ChargebackPaymentByExternalKeyParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) WithXKillbillReason(xKillbillReason *string) *ChargebackPaymentByExternalKeyParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) WithBody(body *kbmodel.PaymentTransaction) *ChargebackPaymentByExternalKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) WithControlPluginName(controlPluginName []string) *ChargebackPaymentByExternalKeyParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPluginProperty adds the pluginProperty to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) WithPluginProperty(pluginProperty []string) *ChargebackPaymentByExternalKeyParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the chargeback payment by external key params
func (o *ChargebackPaymentByExternalKeyParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *ChargebackPaymentByExternalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamChargebackPaymentByExternalKey binds the parameter controlPluginName
func (o *ChargebackPaymentByExternalKeyParams) bindParamControlPluginName(formats strfmt.Registry) []string {
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

// bindParamChargebackPaymentByExternalKey binds the parameter pluginProperty
func (o *ChargebackPaymentByExternalKeyParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
