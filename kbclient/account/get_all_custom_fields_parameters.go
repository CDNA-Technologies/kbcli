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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAllCustomFieldsParams creates a new GetAllCustomFieldsParams object
// with the default values initialized.
func NewGetAllCustomFieldsParams() *GetAllCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetAllCustomFieldsParams{
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAllCustomFieldsParamsWithTimeout creates a new GetAllCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAllCustomFieldsParamsWithTimeout(timeout time.Duration) *GetAllCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetAllCustomFieldsParams{
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetAllCustomFieldsParamsWithContext creates a new GetAllCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAllCustomFieldsParamsWithContext(ctx context.Context) *GetAllCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetAllCustomFieldsParams{
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetAllCustomFieldsParamsWithHTTPClient creates a new GetAllCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAllCustomFieldsParamsWithHTTPClient(client *http.Client) *GetAllCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetAllCustomFieldsParams{
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetAllCustomFieldsParams contains all the parameters to send to the API endpoint
for the get all custom fields operation typically these are written to a http.Request
*/
type GetAllCustomFieldsParams struct {

	/*AccountID*/
	AccountID strfmt.UUID
	/*Audit*/
	Audit *string
	/*ObjectType*/
	ObjectType *string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get all custom fields params
func (o *GetAllCustomFieldsParams) WithTimeout(timeout time.Duration) *GetAllCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get all custom fields params
func (o *GetAllCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get all custom fields params
func (o *GetAllCustomFieldsParams) WithContext(ctx context.Context) *GetAllCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get all custom fields params
func (o *GetAllCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get all custom fields params
func (o *GetAllCustomFieldsParams) WithHTTPClient(client *http.Client) *GetAllCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get all custom fields params
func (o *GetAllCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get all custom fields params
func (o *GetAllCustomFieldsParams) WithAccountID(accountID strfmt.UUID) *GetAllCustomFieldsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get all custom fields params
func (o *GetAllCustomFieldsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAudit adds the audit to the get all custom fields params
func (o *GetAllCustomFieldsParams) WithAudit(audit *string) *GetAllCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get all custom fields params
func (o *GetAllCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithObjectType adds the objectType to the get all custom fields params
func (o *GetAllCustomFieldsParams) WithObjectType(objectType *string) *GetAllCustomFieldsParams {
	o.SetObjectType(objectType)
	return o
}

// SetObjectType adds the objectType to the get all custom fields params
func (o *GetAllCustomFieldsParams) SetObjectType(objectType *string) {
	o.ObjectType = objectType
}

// WriteToRequest writes these params to a swagger request
func (o *GetAllCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
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

	if o.ObjectType != nil {

		// query param objectType
		var qrObjectType string
		if o.ObjectType != nil {
			qrObjectType = *o.ObjectType
		}
		qObjectType := qrObjectType
		if qObjectType != "" {
			if err := r.SetQueryParam("objectType", qObjectType); err != nil {
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
