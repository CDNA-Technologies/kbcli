// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTriggerInvoiceGenerationForParkedAccountsParams creates a new TriggerInvoiceGenerationForParkedAccountsParams object
// with the default values initialized.
func NewTriggerInvoiceGenerationForParkedAccountsParams() *TriggerInvoiceGenerationForParkedAccountsParams {
	var (
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &TriggerInvoiceGenerationForParkedAccountsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTriggerInvoiceGenerationForParkedAccountsParamsWithTimeout creates a new TriggerInvoiceGenerationForParkedAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTriggerInvoiceGenerationForParkedAccountsParamsWithTimeout(timeout time.Duration) *TriggerInvoiceGenerationForParkedAccountsParams {
	var (
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &TriggerInvoiceGenerationForParkedAccountsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewTriggerInvoiceGenerationForParkedAccountsParamsWithContext creates a new TriggerInvoiceGenerationForParkedAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewTriggerInvoiceGenerationForParkedAccountsParamsWithContext(ctx context.Context) *TriggerInvoiceGenerationForParkedAccountsParams {
	var (
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &TriggerInvoiceGenerationForParkedAccountsParams{
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewTriggerInvoiceGenerationForParkedAccountsParamsWithHTTPClient creates a new TriggerInvoiceGenerationForParkedAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTriggerInvoiceGenerationForParkedAccountsParamsWithHTTPClient(client *http.Client) *TriggerInvoiceGenerationForParkedAccountsParams {
	var (
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &TriggerInvoiceGenerationForParkedAccountsParams{
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*TriggerInvoiceGenerationForParkedAccountsParams contains all the parameters to send to the API endpoint
for the trigger invoice generation for parked accounts operation typically these are written to a http.Request
*/
type TriggerInvoiceGenerationForParkedAccountsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) WithTimeout(timeout time.Duration) *TriggerInvoiceGenerationForParkedAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) WithContext(ctx context.Context) *TriggerInvoiceGenerationForParkedAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) WithHTTPClient(client *http.Client) *TriggerInvoiceGenerationForParkedAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) WithXKillbillComment(xKillbillComment *string) *TriggerInvoiceGenerationForParkedAccountsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *TriggerInvoiceGenerationForParkedAccountsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) WithXKillbillReason(xKillbillReason *string) *TriggerInvoiceGenerationForParkedAccountsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithLimit adds the limit to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) WithLimit(limit *int64) *TriggerInvoiceGenerationForParkedAccountsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) WithOffset(offset *int64) *TriggerInvoiceGenerationForParkedAccountsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the trigger invoice generation for parked accounts params
func (o *TriggerInvoiceGenerationForParkedAccountsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *TriggerInvoiceGenerationForParkedAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

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
