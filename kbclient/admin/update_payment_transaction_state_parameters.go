// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewUpdatePaymentTransactionStateParams creates a new UpdatePaymentTransactionStateParams object
// with the default values initialized.
func NewUpdatePaymentTransactionStateParams() *UpdatePaymentTransactionStateParams {
	var ()
	return &UpdatePaymentTransactionStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdatePaymentTransactionStateParamsWithTimeout creates a new UpdatePaymentTransactionStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdatePaymentTransactionStateParamsWithTimeout(timeout time.Duration) *UpdatePaymentTransactionStateParams {
	var ()
	return &UpdatePaymentTransactionStateParams{

		timeout: timeout,
	}
}

// NewUpdatePaymentTransactionStateParamsWithContext creates a new UpdatePaymentTransactionStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdatePaymentTransactionStateParamsWithContext(ctx context.Context) *UpdatePaymentTransactionStateParams {
	var ()
	return &UpdatePaymentTransactionStateParams{

		Context: ctx,
	}
}

// NewUpdatePaymentTransactionStateParamsWithHTTPClient creates a new UpdatePaymentTransactionStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdatePaymentTransactionStateParamsWithHTTPClient(client *http.Client) *UpdatePaymentTransactionStateParams {
	var ()
	return &UpdatePaymentTransactionStateParams{
		HTTPClient: client,
	}
}

/*UpdatePaymentTransactionStateParams contains all the parameters to send to the API endpoint
for the update payment transaction state operation typically these are written to a http.Request
*/
type UpdatePaymentTransactionStateParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.AdminPayment
	/*PaymentID*/
	PaymentID strfmt.UUID
	/*PaymentTransactionID*/
	PaymentTransactionID strfmt.UUID

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithTimeout(timeout time.Duration) *UpdatePaymentTransactionStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithContext(ctx context.Context) *UpdatePaymentTransactionStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithHTTPClient(client *http.Client) *UpdatePaymentTransactionStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithXKillbillAPIKey(xKillbillAPIKey string) *UpdatePaymentTransactionStateParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithXKillbillAPISecret(xKillbillAPISecret string) *UpdatePaymentTransactionStateParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithXKillbillComment(xKillbillComment *string) *UpdatePaymentTransactionStateParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UpdatePaymentTransactionStateParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithXKillbillReason(xKillbillReason *string) *UpdatePaymentTransactionStateParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithBody(body *kbmodel.AdminPayment) *UpdatePaymentTransactionStateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetBody(body *kbmodel.AdminPayment) {
	o.Body = body
}

// WithPaymentID adds the paymentID to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithPaymentID(paymentID strfmt.UUID) *UpdatePaymentTransactionStateParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPaymentTransactionID adds the paymentTransactionID to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) WithPaymentTransactionID(paymentTransactionID strfmt.UUID) *UpdatePaymentTransactionStateParams {
	o.SetPaymentTransactionID(paymentTransactionID)
	return o
}

// SetPaymentTransactionID adds the paymentTransactionId to the update payment transaction state params
func (o *UpdatePaymentTransactionStateParams) SetPaymentTransactionID(paymentTransactionID strfmt.UUID) {
	o.PaymentTransactionID = paymentTransactionID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdatePaymentTransactionStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param paymentTransactionId
	if err := r.SetPathParam("paymentTransactionId", o.PaymentTransactionID.String()); err != nil {
		return err
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
