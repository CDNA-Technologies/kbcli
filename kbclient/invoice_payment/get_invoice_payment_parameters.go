// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetInvoicePaymentParams creates a new GetInvoicePaymentParams object
// with the default values initialized.
func NewGetInvoicePaymentParams() *GetInvoicePaymentParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetInvoicePaymentParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoicePaymentParamsWithTimeout creates a new GetInvoicePaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoicePaymentParamsWithTimeout(timeout time.Duration) *GetInvoicePaymentParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetInvoicePaymentParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: timeout,
	}
}

// NewGetInvoicePaymentParamsWithContext creates a new GetInvoicePaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoicePaymentParamsWithContext(ctx context.Context) *GetInvoicePaymentParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetInvoicePaymentParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		Context: ctx,
	}
}

// NewGetInvoicePaymentParamsWithHTTPClient creates a new GetInvoicePaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoicePaymentParamsWithHTTPClient(client *http.Client) *GetInvoicePaymentParams {
	var (
		auditDefault          = string("NONE")
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &GetInvoicePaymentParams{
		Audit:          &auditDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,
		HTTPClient:     client,
	}
}

/*GetInvoicePaymentParams contains all the parameters to send to the API endpoint
for the get invoice payment operation typically these are written to a http.Request
*/
type GetInvoicePaymentParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*PaymentID*/
	PaymentID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string
	/*WithAttempts*/
	WithAttempts *bool
	/*WithPluginInfo*/
	WithPluginInfo *bool

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the get invoice payment params
func (o *GetInvoicePaymentParams) WithTimeout(timeout time.Duration) *GetInvoicePaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice payment params
func (o *GetInvoicePaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice payment params
func (o *GetInvoicePaymentParams) WithContext(ctx context.Context) *GetInvoicePaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice payment params
func (o *GetInvoicePaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice payment params
func (o *GetInvoicePaymentParams) WithHTTPClient(client *http.Client) *GetInvoicePaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice payment params
func (o *GetInvoicePaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get invoice payment params
func (o *GetInvoicePaymentParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetInvoicePaymentParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get invoice payment params
func (o *GetInvoicePaymentParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get invoice payment params
func (o *GetInvoicePaymentParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetInvoicePaymentParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get invoice payment params
func (o *GetInvoicePaymentParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get invoice payment params
func (o *GetInvoicePaymentParams) WithAudit(audit *string) *GetInvoicePaymentParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice payment params
func (o *GetInvoicePaymentParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithPaymentID adds the paymentID to the get invoice payment params
func (o *GetInvoicePaymentParams) WithPaymentID(paymentID strfmt.UUID) *GetInvoicePaymentParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the get invoice payment params
func (o *GetInvoicePaymentParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the get invoice payment params
func (o *GetInvoicePaymentParams) WithPluginProperty(pluginProperty []string) *GetInvoicePaymentParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the get invoice payment params
func (o *GetInvoicePaymentParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithWithAttempts adds the withAttempts to the get invoice payment params
func (o *GetInvoicePaymentParams) WithWithAttempts(withAttempts *bool) *GetInvoicePaymentParams {
	o.SetWithAttempts(withAttempts)
	return o
}

// SetWithAttempts adds the withAttempts to the get invoice payment params
func (o *GetInvoicePaymentParams) SetWithAttempts(withAttempts *bool) {
	o.WithAttempts = withAttempts
}

// WithWithPluginInfo adds the withPluginInfo to the get invoice payment params
func (o *GetInvoicePaymentParams) WithWithPluginInfo(withPluginInfo *bool) *GetInvoicePaymentParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the get invoice payment params
func (o *GetInvoicePaymentParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoicePaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

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

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if o.WithAttempts != nil {

		// query param withAttempts
		var qrWithAttempts bool
		if o.WithAttempts != nil {
			qrWithAttempts = *o.WithAttempts
		}
		qWithAttempts := swag.FormatBool(qrWithAttempts)
		if qWithAttempts != "" {
			if err := r.SetQueryParam("withAttempts", qWithAttempts); err != nil {
				return err
			}
		}

	}

	if o.WithPluginInfo != nil {

		// query param withPluginInfo
		var qrWithPluginInfo bool
		if o.WithPluginInfo != nil {
			qrWithPluginInfo = *o.WithPluginInfo
		}
		qWithPluginInfo := swag.FormatBool(qrWithPluginInfo)
		if qWithPluginInfo != "" {
			if err := r.SetQueryParam("withPluginInfo", qWithPluginInfo); err != nil {
				return err
			}
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
