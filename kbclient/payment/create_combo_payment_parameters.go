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

// NewCreateComboPaymentParams creates a new CreateComboPaymentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateComboPaymentParams() *CreateComboPaymentParams {
	return &CreateComboPaymentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateComboPaymentParamsWithTimeout creates a new CreateComboPaymentParams object
// with the ability to set a timeout on a request.
func NewCreateComboPaymentParamsWithTimeout(timeout time.Duration) *CreateComboPaymentParams {
	return &CreateComboPaymentParams{
		timeout: timeout,
	}
}

// NewCreateComboPaymentParamsWithContext creates a new CreateComboPaymentParams object
// with the ability to set a context for a request.
func NewCreateComboPaymentParamsWithContext(ctx context.Context) *CreateComboPaymentParams {
	return &CreateComboPaymentParams{
		Context: ctx,
	}
}

// NewCreateComboPaymentParamsWithHTTPClient creates a new CreateComboPaymentParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateComboPaymentParamsWithHTTPClient(client *http.Client) *CreateComboPaymentParams {
	return &CreateComboPaymentParams{
		HTTPClient: client,
	}
}

/* CreateComboPaymentParams contains all the parameters to send to the API endpoint
   for the create combo payment operation.

   Typically these are written to a http.Request.
*/
type CreateComboPaymentParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.ComboPaymentTransaction

	// ControlPluginName.
	ControlPluginName []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create combo payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateComboPaymentParams) WithDefaults() *CreateComboPaymentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create combo payment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateComboPaymentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create combo payment params
func (o *CreateComboPaymentParams) WithTimeout(timeout time.Duration) *CreateComboPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create combo payment params
func (o *CreateComboPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create combo payment params
func (o *CreateComboPaymentParams) WithContext(ctx context.Context) *CreateComboPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create combo payment params
func (o *CreateComboPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create combo payment params
func (o *CreateComboPaymentParams) WithHTTPClient(client *http.Client) *CreateComboPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create combo payment params
func (o *CreateComboPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create combo payment params
func (o *CreateComboPaymentParams) WithXKillbillComment(xKillbillComment *string) *CreateComboPaymentParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create combo payment params
func (o *CreateComboPaymentParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create combo payment params
func (o *CreateComboPaymentParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateComboPaymentParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create combo payment params
func (o *CreateComboPaymentParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create combo payment params
func (o *CreateComboPaymentParams) WithXKillbillReason(xKillbillReason *string) *CreateComboPaymentParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create combo payment params
func (o *CreateComboPaymentParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create combo payment params
func (o *CreateComboPaymentParams) WithBody(body *kbmodel.ComboPaymentTransaction) *CreateComboPaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create combo payment params
func (o *CreateComboPaymentParams) SetBody(body *kbmodel.ComboPaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the create combo payment params
func (o *CreateComboPaymentParams) WithControlPluginName(controlPluginName []string) *CreateComboPaymentParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the create combo payment params
func (o *CreateComboPaymentParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WriteToRequest writes these params to a swagger request
func (o *CreateComboPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCreateComboPayment binds the parameter controlPluginName
func (o *CreateComboPaymentParams) bindParamControlPluginName(formats strfmt.Registry) []string {
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
