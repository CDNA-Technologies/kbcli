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

	"github.com/killbill/kbcli/v2/kbmodel"
)

// NewCreateAccountParams creates a new CreateAccountParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateAccountParams() *CreateAccountParams {
	return &CreateAccountParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccountParamsWithTimeout creates a new CreateAccountParams object
// with the ability to set a timeout on a request.
func NewCreateAccountParamsWithTimeout(timeout time.Duration) *CreateAccountParams {
	return &CreateAccountParams{
		timeout: timeout,
	}
}

// NewCreateAccountParamsWithContext creates a new CreateAccountParams object
// with the ability to set a context for a request.
func NewCreateAccountParamsWithContext(ctx context.Context) *CreateAccountParams {
	return &CreateAccountParams{
		Context: ctx,
	}
}

// NewCreateAccountParamsWithHTTPClient creates a new CreateAccountParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateAccountParamsWithHTTPClient(client *http.Client) *CreateAccountParams {
	return &CreateAccountParams{
		HTTPClient: client,
	}
}

/* CreateAccountParams contains all the parameters to send to the API endpoint
   for the create account operation.

   Typically these are written to a http.Request.
*/
type CreateAccountParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// Body.
	Body *kbmodel.Account

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAccountParams) WithDefaults() *CreateAccountParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create account params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateAccountParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create account params
func (o *CreateAccountParams) WithTimeout(timeout time.Duration) *CreateAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create account params
func (o *CreateAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create account params
func (o *CreateAccountParams) WithContext(ctx context.Context) *CreateAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create account params
func (o *CreateAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create account params
func (o *CreateAccountParams) WithHTTPClient(client *http.Client) *CreateAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create account params
func (o *CreateAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create account params
func (o *CreateAccountParams) WithXKillbillComment(xKillbillComment *string) *CreateAccountParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create account params
func (o *CreateAccountParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create account params
func (o *CreateAccountParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateAccountParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create account params
func (o *CreateAccountParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create account params
func (o *CreateAccountParams) WithXKillbillReason(xKillbillReason *string) *CreateAccountParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create account params
func (o *CreateAccountParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create account params
func (o *CreateAccountParams) WithBody(body *kbmodel.Account) *CreateAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create account params
func (o *CreateAccountParams) SetBody(body *kbmodel.Account) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
