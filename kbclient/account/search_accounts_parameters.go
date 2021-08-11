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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSearchAccountsParams creates a new SearchAccountsParams object
// with the default values initialized.
func NewSearchAccountsParams() *SearchAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
		limitDefault                    = int64(100)
		offsetDefault                   = int64(0)
	)
	return &SearchAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:                    &auditDefault,
		Limit:                    &limitDefault,
		Offset:                   &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchAccountsParamsWithTimeout creates a new SearchAccountsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchAccountsParamsWithTimeout(timeout time.Duration) *SearchAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
		limitDefault                    = int64(100)
		offsetDefault                   = int64(0)
	)
	return &SearchAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:                    &auditDefault,
		Limit:                    &limitDefault,
		Offset:                   &offsetDefault,

		timeout: timeout,
	}
}

// NewSearchAccountsParamsWithContext creates a new SearchAccountsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchAccountsParamsWithContext(ctx context.Context) *SearchAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
		limitDefault                    = int64(100)
		offsetDefault                   = int64(0)
	)
	return &SearchAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:                    &auditDefault,
		Limit:                    &limitDefault,
		Offset:                   &offsetDefault,

		Context: ctx,
	}
}

// NewSearchAccountsParamsWithHTTPClient creates a new SearchAccountsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchAccountsParamsWithHTTPClient(client *http.Client) *SearchAccountsParams {
	var (
		accountWithBalanceDefault       = bool(false)
		accountWithBalanceAndCBADefault = bool(false)
		auditDefault                    = string("NONE")
		limitDefault                    = int64(100)
		offsetDefault                   = int64(0)
	)
	return &SearchAccountsParams{
		AccountWithBalance:       &accountWithBalanceDefault,
		AccountWithBalanceAndCBA: &accountWithBalanceAndCBADefault,
		Audit:                    &auditDefault,
		Limit:                    &limitDefault,
		Offset:                   &offsetDefault,
		HTTPClient:               client,
	}
}

/*SearchAccountsParams contains all the parameters to send to the API endpoint
for the search accounts operation typically these are written to a http.Request
*/
type SearchAccountsParams struct {

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
	/*SearchKey*/
	SearchKey string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the search accounts params
func (o *SearchAccountsParams) WithTimeout(timeout time.Duration) *SearchAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search accounts params
func (o *SearchAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search accounts params
func (o *SearchAccountsParams) WithContext(ctx context.Context) *SearchAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search accounts params
func (o *SearchAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search accounts params
func (o *SearchAccountsParams) WithHTTPClient(client *http.Client) *SearchAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search accounts params
func (o *SearchAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountWithBalance adds the accountWithBalance to the search accounts params
func (o *SearchAccountsParams) WithAccountWithBalance(accountWithBalance *bool) *SearchAccountsParams {
	o.SetAccountWithBalance(accountWithBalance)
	return o
}

// SetAccountWithBalance adds the accountWithBalance to the search accounts params
func (o *SearchAccountsParams) SetAccountWithBalance(accountWithBalance *bool) {
	o.AccountWithBalance = accountWithBalance
}

// WithAccountWithBalanceAndCBA adds the accountWithBalanceAndCBA to the search accounts params
func (o *SearchAccountsParams) WithAccountWithBalanceAndCBA(accountWithBalanceAndCBA *bool) *SearchAccountsParams {
	o.SetAccountWithBalanceAndCBA(accountWithBalanceAndCBA)
	return o
}

// SetAccountWithBalanceAndCBA adds the accountWithBalanceAndCBA to the search accounts params
func (o *SearchAccountsParams) SetAccountWithBalanceAndCBA(accountWithBalanceAndCBA *bool) {
	o.AccountWithBalanceAndCBA = accountWithBalanceAndCBA
}

// WithAudit adds the audit to the search accounts params
func (o *SearchAccountsParams) WithAudit(audit *string) *SearchAccountsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the search accounts params
func (o *SearchAccountsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the search accounts params
func (o *SearchAccountsParams) WithLimit(limit *int64) *SearchAccountsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search accounts params
func (o *SearchAccountsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search accounts params
func (o *SearchAccountsParams) WithOffset(offset *int64) *SearchAccountsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search accounts params
func (o *SearchAccountsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithSearchKey adds the searchKey to the search accounts params
func (o *SearchAccountsParams) WithSearchKey(searchKey string) *SearchAccountsParams {
	o.SetSearchKey(searchKey)
	return o
}

// SetSearchKey adds the searchKey to the search accounts params
func (o *SearchAccountsParams) SetSearchKey(searchKey string) {
	o.SearchKey = searchKey
}

// WriteToRequest writes these params to a swagger request
func (o *SearchAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param searchKey
	if err := r.SetPathParam("searchKey", o.SearchKey); err != nil {
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
