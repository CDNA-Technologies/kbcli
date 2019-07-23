// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewGetAccountsParams creates a new GetAccountsParams object
// with the default values initialized.
func NewGetAccountsParams() *GetAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
		limitDefault                    = int64(100)
		offsetDefault                   = int64(0)
	)
	return &GetAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:                    &auditDefault,
		Limit:                    &limitDefault,
		Offset:                   &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountsParamsWithTimeout creates a new GetAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountsParamsWithTimeout(timeout time.Duration) *GetAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
		limitDefault                    = int64(100)
		offsetDefault                   = int64(0)
	)
	return &GetAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:                    &auditDefault,
		Limit:                    &limitDefault,
		Offset:                   &offsetDefault,

		timeout: timeout,
	}
}

// NewGetAccountsParamsWithContext creates a new GetAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountsParamsWithContext(ctx context.Context) *GetAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
		limitDefault                    = int64(100)
		offsetDefault                   = int64(0)
	)
	return &GetAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:                    &auditDefault,
		Limit:                    &limitDefault,
		Offset:                   &offsetDefault,

		Context: ctx,
	}
}

// NewGetAccountsParamsWithHTTPClient creates a new GetAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountsParamsWithHTTPClient(client *http.Client) *GetAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
		limitDefault                    = int64(100)
		offsetDefault                   = int64(0)
	)
	return &GetAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:                    &auditDefault,
		Limit:                    &limitDefault,
		Offset:                   &offsetDefault,
		HTTPClient:               client,
	}
}

/*GetAccountsParams contains all the parameters to send to the API endpoint
for the get accounts operation typically these are written to a http.Request
*/
type GetAccountsParams struct {

	/*AccountWithBalance*/
	AccountWithBalance *bool
	/*AccountWithBalanceAndCBA*/
	AccountWithBalanceAndCBA *bool
	/*Audit*/
	Audit *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) WithTimeout(timeout time.Duration) *GetAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get accounts params
func (o *GetAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get accounts params
func (o *GetAccountsParams) WithContext(ctx context.Context) *GetAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get accounts params
func (o *GetAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get accounts params
func (o *GetAccountsParams) WithHTTPClient(client *http.Client) *GetAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get accounts params
func (o *GetAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountWithBalance adds the accountWithBalance to the get accounts params
func (o *GetAccountsParams) WithAccountWithBalance(accountWithBalance *bool) *GetAccountsParams {
	o.SetAccountWithBalance(accountWithBalance)
	return o
}

// SetAccountWithBalance adds the accountWithBalance to the get accounts params
func (o *GetAccountsParams) SetAccountWithBalance(accountWithBalance *bool) {
	o.AccountWithBalance = accountWithBalance
}

// WithAccountWithBalanceAndCBA adds the accountWithBalanceAndCBA to the get accounts params
func (o *GetAccountsParams) WithAccountWithBalanceAndCBA(accountWithBalanceAndCBA *bool) *GetAccountsParams {
	o.SetAccountWithBalanceAndCBA(accountWithBalanceAndCBA)
	return o
}

// SetAccountWithBalanceAndCBA adds the accountWithBalanceAndCBA to the get accounts params
func (o *GetAccountsParams) SetAccountWithBalanceAndCBA(accountWithBalanceAndCBA *bool) {
	o.AccountWithBalanceAndCBA = accountWithBalanceAndCBA
}

// WithAudit adds the audit to the get accounts params
func (o *GetAccountsParams) WithAudit(audit *string) *GetAccountsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get accounts params
func (o *GetAccountsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the get accounts params
func (o *GetAccountsParams) WithLimit(limit *int64) *GetAccountsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get accounts params
func (o *GetAccountsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get accounts params
func (o *GetAccountsParams) WithOffset(offset *int64) *GetAccountsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get accounts params
func (o *GetAccountsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountWithBalance != nil {

		// query param accountWithBalance
		var qrAccountWithBalance bool
		if o.AccountWithBalance != nil {
			qrAccountWithBalance = *o.AccountWithBalance
		}
		qAccountWithBalance := swag.FormatBool(qrAccountWithBalance)
		if qAccountWithBalance != "" {
			if err := r.SetQueryParam("accountWithBalance", qAccountWithBalance); err != nil {
				return err
			}
		}

	}

	if o.AccountWithBalanceAndCBA != nil {

		// query param accountWithBalanceAndCBA
		var qrAccountWithBalanceAndCBA bool
		if o.AccountWithBalanceAndCBA != nil {
			qrAccountWithBalanceAndCBA = *o.AccountWithBalanceAndCBA
		}
		qAccountWithBalanceAndCBA := swag.FormatBool(qrAccountWithBalanceAndCBA)
		if qAccountWithBalanceAndCBA != "" {
			if err := r.SetQueryParam("accountWithBalanceAndCBA", qAccountWithBalanceAndCBA); err != nil {
				return err
			}
		}

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
