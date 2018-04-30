// Code generated by go-swagger; DO NOT EDIT.

package invoice

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
)

// NewCreateFutureInvoiceParams creates a new CreateFutureInvoiceParams object
// with the default values initialized.
func NewCreateFutureInvoiceParams() *CreateFutureInvoiceParams {
	var ()
	return &CreateFutureInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFutureInvoiceParamsWithTimeout creates a new CreateFutureInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFutureInvoiceParamsWithTimeout(timeout time.Duration) *CreateFutureInvoiceParams {
	var ()
	return &CreateFutureInvoiceParams{

		timeout: timeout,
	}
}

// NewCreateFutureInvoiceParamsWithContext creates a new CreateFutureInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFutureInvoiceParamsWithContext(ctx context.Context) *CreateFutureInvoiceParams {
	var ()
	return &CreateFutureInvoiceParams{

		Context: ctx,
	}
}

// NewCreateFutureInvoiceParamsWithHTTPClient creates a new CreateFutureInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFutureInvoiceParamsWithHTTPClient(client *http.Client) *CreateFutureInvoiceParams {
	var ()
	return &CreateFutureInvoiceParams{
		HTTPClient: client,
	}
}

/*CreateFutureInvoiceParams contains all the parameters to send to the API endpoint
for the create future invoice operation typically these are written to a http.Request
*/
type CreateFutureInvoiceParams struct {

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
	/*AccountID*/
	AccountID strfmt.UUID
	/*TargetDate*/
	TargetDate *strfmt.Date

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the create future invoice params
func (o *CreateFutureInvoiceParams) WithTimeout(timeout time.Duration) *CreateFutureInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create future invoice params
func (o *CreateFutureInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create future invoice params
func (o *CreateFutureInvoiceParams) WithContext(ctx context.Context) *CreateFutureInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create future invoice params
func (o *CreateFutureInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create future invoice params
func (o *CreateFutureInvoiceParams) WithHTTPClient(client *http.Client) *CreateFutureInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create future invoice params
func (o *CreateFutureInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create future invoice params
func (o *CreateFutureInvoiceParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreateFutureInvoiceParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create future invoice params
func (o *CreateFutureInvoiceParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create future invoice params
func (o *CreateFutureInvoiceParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreateFutureInvoiceParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create future invoice params
func (o *CreateFutureInvoiceParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create future invoice params
func (o *CreateFutureInvoiceParams) WithXKillbillComment(xKillbillComment *string) *CreateFutureInvoiceParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create future invoice params
func (o *CreateFutureInvoiceParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create future invoice params
func (o *CreateFutureInvoiceParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateFutureInvoiceParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create future invoice params
func (o *CreateFutureInvoiceParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create future invoice params
func (o *CreateFutureInvoiceParams) WithXKillbillReason(xKillbillReason *string) *CreateFutureInvoiceParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create future invoice params
func (o *CreateFutureInvoiceParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the create future invoice params
func (o *CreateFutureInvoiceParams) WithAccountID(accountID strfmt.UUID) *CreateFutureInvoiceParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the create future invoice params
func (o *CreateFutureInvoiceParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithTargetDate adds the targetDate to the create future invoice params
func (o *CreateFutureInvoiceParams) WithTargetDate(targetDate *strfmt.Date) *CreateFutureInvoiceParams {
	o.SetTargetDate(targetDate)
	return o
}

// SetTargetDate adds the targetDate to the create future invoice params
func (o *CreateFutureInvoiceParams) SetTargetDate(targetDate *strfmt.Date) {
	o.TargetDate = targetDate
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFutureInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param accountId
	qrAccountID := o.AccountID
	qAccountID := qrAccountID.String()
	if qAccountID != "" {
		if err := r.SetQueryParam("accountId", qAccountID); err != nil {
			return err
		}
	}

	if o.TargetDate != nil {

		// query param targetDate
		var qrTargetDate strfmt.Date
		if o.TargetDate != nil {
			qrTargetDate = *o.TargetDate
		}
		qTargetDate := qrTargetDate.String()
		if qTargetDate != "" {
			if err := r.SetQueryParam("targetDate", qTargetDate); err != nil {
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
