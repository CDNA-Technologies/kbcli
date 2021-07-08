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
)

// NewGetAccountEmailAuditLogsWithHistoryParams creates a new GetAccountEmailAuditLogsWithHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetAccountEmailAuditLogsWithHistoryParams() *GetAccountEmailAuditLogsWithHistoryParams {
	return &GetAccountEmailAuditLogsWithHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountEmailAuditLogsWithHistoryParamsWithTimeout creates a new GetAccountEmailAuditLogsWithHistoryParams object
// with the ability to set a timeout on a request.
func NewGetAccountEmailAuditLogsWithHistoryParamsWithTimeout(timeout time.Duration) *GetAccountEmailAuditLogsWithHistoryParams {
	return &GetAccountEmailAuditLogsWithHistoryParams{
		timeout: timeout,
	}
}

// NewGetAccountEmailAuditLogsWithHistoryParamsWithContext creates a new GetAccountEmailAuditLogsWithHistoryParams object
// with the ability to set a context for a request.
func NewGetAccountEmailAuditLogsWithHistoryParamsWithContext(ctx context.Context) *GetAccountEmailAuditLogsWithHistoryParams {
	return &GetAccountEmailAuditLogsWithHistoryParams{
		Context: ctx,
	}
}

// NewGetAccountEmailAuditLogsWithHistoryParamsWithHTTPClient creates a new GetAccountEmailAuditLogsWithHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetAccountEmailAuditLogsWithHistoryParamsWithHTTPClient(client *http.Client) *GetAccountEmailAuditLogsWithHistoryParams {
	return &GetAccountEmailAuditLogsWithHistoryParams{
		HTTPClient: client,
	}
}

/* GetAccountEmailAuditLogsWithHistoryParams contains all the parameters to send to the API endpoint
   for the get account email audit logs with history operation.

   Typically these are written to a http.Request.
*/
type GetAccountEmailAuditLogsWithHistoryParams struct {

	// AccountEmailID.
	//
	// Format: uuid
	AccountEmailID strfmt.UUID

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get account email audit logs with history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountEmailAuditLogsWithHistoryParams) WithDefaults() *GetAccountEmailAuditLogsWithHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get account email audit logs with history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetAccountEmailAuditLogsWithHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) WithTimeout(timeout time.Duration) *GetAccountEmailAuditLogsWithHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) WithContext(ctx context.Context) *GetAccountEmailAuditLogsWithHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) WithHTTPClient(client *http.Client) *GetAccountEmailAuditLogsWithHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountEmailID adds the accountEmailID to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) WithAccountEmailID(accountEmailID strfmt.UUID) *GetAccountEmailAuditLogsWithHistoryParams {
	o.SetAccountEmailID(accountEmailID)
	return o
}

// SetAccountEmailID adds the accountEmailId to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) SetAccountEmailID(accountEmailID strfmt.UUID) {
	o.AccountEmailID = accountEmailID
}

// WithAccountID adds the accountID to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) WithAccountID(accountID strfmt.UUID) *GetAccountEmailAuditLogsWithHistoryParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get account email audit logs with history params
func (o *GetAccountEmailAuditLogsWithHistoryParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountEmailAuditLogsWithHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountEmailId
	if err := r.SetPathParam("accountEmailId", o.AccountEmailID.String()); err != nil {
		return err
	}

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
