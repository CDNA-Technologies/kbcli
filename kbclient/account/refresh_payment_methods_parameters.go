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
	"github.com/go-openapi/swag"
)

// NewRefreshPaymentMethodsParams creates a new RefreshPaymentMethodsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRefreshPaymentMethodsParams() *RefreshPaymentMethodsParams {
	return &RefreshPaymentMethodsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshPaymentMethodsParamsWithTimeout creates a new RefreshPaymentMethodsParams object
// with the ability to set a timeout on a request.
func NewRefreshPaymentMethodsParamsWithTimeout(timeout time.Duration) *RefreshPaymentMethodsParams {
	return &RefreshPaymentMethodsParams{
		timeout: timeout,
	}
}

// NewRefreshPaymentMethodsParamsWithContext creates a new RefreshPaymentMethodsParams object
// with the ability to set a context for a request.
func NewRefreshPaymentMethodsParamsWithContext(ctx context.Context) *RefreshPaymentMethodsParams {
	return &RefreshPaymentMethodsParams{
		Context: ctx,
	}
}

// NewRefreshPaymentMethodsParamsWithHTTPClient creates a new RefreshPaymentMethodsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRefreshPaymentMethodsParamsWithHTTPClient(client *http.Client) *RefreshPaymentMethodsParams {
	return &RefreshPaymentMethodsParams{
		HTTPClient: client,
	}
}

/* RefreshPaymentMethodsParams contains all the parameters to send to the API endpoint
   for the refresh payment methods operation.

   Typically these are written to a http.Request.
*/
type RefreshPaymentMethodsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// AccountID.
	//
	// Format: uuid
	AccountID strfmt.UUID

	// PluginName.
	PluginName *string

	// PluginProperty.
	PluginProperty []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the refresh payment methods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshPaymentMethodsParams) WithDefaults() *RefreshPaymentMethodsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the refresh payment methods params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RefreshPaymentMethodsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) WithTimeout(timeout time.Duration) *RefreshPaymentMethodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) WithContext(ctx context.Context) *RefreshPaymentMethodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) WithHTTPClient(client *http.Client) *RefreshPaymentMethodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) WithXKillbillComment(xKillbillComment *string) *RefreshPaymentMethodsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *RefreshPaymentMethodsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) WithXKillbillReason(xKillbillReason *string) *RefreshPaymentMethodsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) WithAccountID(accountID strfmt.UUID) *RefreshPaymentMethodsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithPluginName adds the pluginName to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) WithPluginName(pluginName *string) *RefreshPaymentMethodsParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) SetPluginName(pluginName *string) {
	o.PluginName = pluginName
}

// WithPluginProperty adds the pluginProperty to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) WithPluginProperty(pluginProperty []string) *RefreshPaymentMethodsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the refresh payment methods params
func (o *RefreshPaymentMethodsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshPaymentMethodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if o.PluginName != nil {

		// query param pluginName
		var qrPluginName string

		if o.PluginName != nil {
			qrPluginName = *o.PluginName
		}
		qPluginName := qrPluginName
		if qPluginName != "" {

			if err := r.SetQueryParam("pluginName", qPluginName); err != nil {
				return err
			}
		}
	}

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamRefreshPaymentMethods binds the parameter pluginProperty
func (o *RefreshPaymentMethodsParams) bindParamPluginProperty(formats strfmt.Registry) []string {
	pluginPropertyIR := o.PluginProperty

	var pluginPropertyIC []string
	for _, pluginPropertyIIR := range pluginPropertyIR { // explode []string

		pluginPropertyIIV := pluginPropertyIIR // string as string
		pluginPropertyIC = append(pluginPropertyIC, pluginPropertyIIV)
	}

	// items.CollectionFormat: "multi"
	pluginPropertyIS := swag.JoinByFormat(pluginPropertyIC, "multi")

	return pluginPropertyIS
}
