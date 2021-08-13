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

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// NewCreatePaymentCustomFieldsParams creates a new CreatePaymentCustomFieldsParams object
// with the default values initialized.
func NewCreatePaymentCustomFieldsParams() *CreatePaymentCustomFieldsParams {
	var ()
	return &CreatePaymentCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePaymentCustomFieldsParamsWithTimeout creates a new CreatePaymentCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePaymentCustomFieldsParamsWithTimeout(timeout time.Duration) *CreatePaymentCustomFieldsParams {
	var ()
	return &CreatePaymentCustomFieldsParams{

		timeout: timeout,
	}
}

// NewCreatePaymentCustomFieldsParamsWithContext creates a new CreatePaymentCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePaymentCustomFieldsParamsWithContext(ctx context.Context) *CreatePaymentCustomFieldsParams {
	var ()
	return &CreatePaymentCustomFieldsParams{

		Context: ctx,
	}
}

// NewCreatePaymentCustomFieldsParamsWithHTTPClient creates a new CreatePaymentCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePaymentCustomFieldsParamsWithHTTPClient(client *http.Client) *CreatePaymentCustomFieldsParams {
	var ()
	return &CreatePaymentCustomFieldsParams{
		HTTPClient: client,
	}
}

/*CreatePaymentCustomFieldsParams contains all the parameters to send to the API endpoint
for the create payment custom fields operation typically these are written to a http.Request
*/
type CreatePaymentCustomFieldsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body []*kbmodel.CustomField
	/*PaymentID*/
	PaymentID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) WithTimeout(timeout time.Duration) *CreatePaymentCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) WithContext(ctx context.Context) *CreatePaymentCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) WithHTTPClient(client *http.Client) *CreatePaymentCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreatePaymentCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreatePaymentCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreatePaymentCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreatePaymentCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithPaymentID adds the paymentID to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) WithPaymentID(paymentID strfmt.UUID) *CreatePaymentCustomFieldsParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the create payment custom fields params
func (o *CreatePaymentCustomFieldsParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePaymentCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// header param WithProfilingInfo
	if o.WithProfilingInfo != nil && len(*o.WithProfilingInfo) > 0 {
		if err := r.SetHeaderParam("X-Killbill-Profiling-Req", *o.WithProfilingInfo); err != nil {
			return err
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
