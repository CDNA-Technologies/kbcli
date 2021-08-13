// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// NewRenameExternalKeyParams creates a new RenameExternalKeyParams object
// with the default values initialized.
func NewRenameExternalKeyParams() *RenameExternalKeyParams {
	var ()
	return &RenameExternalKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRenameExternalKeyParamsWithTimeout creates a new RenameExternalKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRenameExternalKeyParamsWithTimeout(timeout time.Duration) *RenameExternalKeyParams {
	var ()
	return &RenameExternalKeyParams{

		timeout: timeout,
	}
}

// NewRenameExternalKeyParamsWithContext creates a new RenameExternalKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewRenameExternalKeyParamsWithContext(ctx context.Context) *RenameExternalKeyParams {
	var ()
	return &RenameExternalKeyParams{

		Context: ctx,
	}
}

// NewRenameExternalKeyParamsWithHTTPClient creates a new RenameExternalKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRenameExternalKeyParamsWithHTTPClient(client *http.Client) *RenameExternalKeyParams {
	var ()
	return &RenameExternalKeyParams{
		HTTPClient: client,
	}
}

/*RenameExternalKeyParams contains all the parameters to send to the API endpoint
for the rename external key operation typically these are written to a http.Request
*/
type RenameExternalKeyParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.Bundle
	/*BundleID*/
	BundleID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the rename external key params
func (o *RenameExternalKeyParams) WithTimeout(timeout time.Duration) *RenameExternalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rename external key params
func (o *RenameExternalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rename external key params
func (o *RenameExternalKeyParams) WithContext(ctx context.Context) *RenameExternalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rename external key params
func (o *RenameExternalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rename external key params
func (o *RenameExternalKeyParams) WithHTTPClient(client *http.Client) *RenameExternalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rename external key params
func (o *RenameExternalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the rename external key params
func (o *RenameExternalKeyParams) WithXKillbillComment(xKillbillComment *string) *RenameExternalKeyParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the rename external key params
func (o *RenameExternalKeyParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the rename external key params
func (o *RenameExternalKeyParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *RenameExternalKeyParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the rename external key params
func (o *RenameExternalKeyParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the rename external key params
func (o *RenameExternalKeyParams) WithXKillbillReason(xKillbillReason *string) *RenameExternalKeyParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the rename external key params
func (o *RenameExternalKeyParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the rename external key params
func (o *RenameExternalKeyParams) WithBody(body *kbmodel.Bundle) *RenameExternalKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the rename external key params
func (o *RenameExternalKeyParams) SetBody(body *kbmodel.Bundle) {
	o.Body = body
}

// WithBundleID adds the bundleID to the rename external key params
func (o *RenameExternalKeyParams) WithBundleID(bundleID strfmt.UUID) *RenameExternalKeyParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the rename external key params
func (o *RenameExternalKeyParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WriteToRequest writes these params to a swagger request
func (o *RenameExternalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
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
