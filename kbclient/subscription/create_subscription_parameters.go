// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

	"github.com/killbill/kbcli/v2/kbmodel"
)

// NewCreateSubscriptionParams creates a new CreateSubscriptionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSubscriptionParams() *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSubscriptionParamsWithTimeout creates a new CreateSubscriptionParams object
// with the ability to set a timeout on a request.
func NewCreateSubscriptionParamsWithTimeout(timeout time.Duration) *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		timeout: timeout,
	}
}

// NewCreateSubscriptionParamsWithContext creates a new CreateSubscriptionParams object
// with the ability to set a context for a request.
func NewCreateSubscriptionParamsWithContext(ctx context.Context) *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		Context: ctx,
	}
}

// NewCreateSubscriptionParamsWithHTTPClient creates a new CreateSubscriptionParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSubscriptionParamsWithHTTPClient(client *http.Client) *CreateSubscriptionParams {
	return &CreateSubscriptionParams{
		HTTPClient: client,
	}
}

/* CreateSubscriptionParams contains all the parameters to send to the API endpoint
   for the create subscription operation.

   Typically these are written to a http.Request.
*/
type CreateSubscriptionParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// BillingDate.
	//
	// Format: date
	BillingDate *strfmt.Date

	// Body.
	Body *kbmodel.Subscription

	// CallCompletion.
	CallCompletion *bool

	// CallTimeoutSec.
	//
	// Format: int64
	// Default: 3
	CallTimeoutSec *int64

	// EntitlementDate.
	//
	// Format: date
	EntitlementDate *strfmt.Date

	// Migrated.
	Migrated *bool

	// PluginProperty.
	PluginProperty []string

	// RenameKeyIfExistsAndUnused.
	//
	// Default: true
	RenameKeyIfExistsAndUnused *bool

	// SkipResponse.
	SkipResponse *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSubscriptionParams) WithDefaults() *CreateSubscriptionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create subscription params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSubscriptionParams) SetDefaults() {
	var (
		callCompletionDefault = bool(false)

		callTimeoutSecDefault = int64(3)

		migratedDefault = bool(false)

		renameKeyIfExistsAndUnusedDefault = bool(true)

		skipResponseDefault = bool(false)
	)

	val := CreateSubscriptionParams{
		CallCompletion:             &callCompletionDefault,
		CallTimeoutSec:             &callTimeoutSecDefault,
		Migrated:                   &migratedDefault,
		RenameKeyIfExistsAndUnused: &renameKeyIfExistsAndUnusedDefault,
		SkipResponse:               &skipResponseDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create subscription params
func (o *CreateSubscriptionParams) WithTimeout(timeout time.Duration) *CreateSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create subscription params
func (o *CreateSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create subscription params
func (o *CreateSubscriptionParams) WithContext(ctx context.Context) *CreateSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create subscription params
func (o *CreateSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create subscription params
func (o *CreateSubscriptionParams) WithHTTPClient(client *http.Client) *CreateSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create subscription params
func (o *CreateSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the create subscription params
func (o *CreateSubscriptionParams) WithXKillbillComment(xKillbillComment *string) *CreateSubscriptionParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create subscription params
func (o *CreateSubscriptionParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create subscription params
func (o *CreateSubscriptionParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateSubscriptionParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create subscription params
func (o *CreateSubscriptionParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create subscription params
func (o *CreateSubscriptionParams) WithXKillbillReason(xKillbillReason *string) *CreateSubscriptionParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create subscription params
func (o *CreateSubscriptionParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBillingDate adds the billingDate to the create subscription params
func (o *CreateSubscriptionParams) WithBillingDate(billingDate *strfmt.Date) *CreateSubscriptionParams {
	o.SetBillingDate(billingDate)
	return o
}

// SetBillingDate adds the billingDate to the create subscription params
func (o *CreateSubscriptionParams) SetBillingDate(billingDate *strfmt.Date) {
	o.BillingDate = billingDate
}

// WithBody adds the body to the create subscription params
func (o *CreateSubscriptionParams) WithBody(body *kbmodel.Subscription) *CreateSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create subscription params
func (o *CreateSubscriptionParams) SetBody(body *kbmodel.Subscription) {
	o.Body = body
}

// WithCallCompletion adds the callCompletion to the create subscription params
func (o *CreateSubscriptionParams) WithCallCompletion(callCompletion *bool) *CreateSubscriptionParams {
	o.SetCallCompletion(callCompletion)
	return o
}

// SetCallCompletion adds the callCompletion to the create subscription params
func (o *CreateSubscriptionParams) SetCallCompletion(callCompletion *bool) {
	o.CallCompletion = callCompletion
}

// WithCallTimeoutSec adds the callTimeoutSec to the create subscription params
func (o *CreateSubscriptionParams) WithCallTimeoutSec(callTimeoutSec *int64) *CreateSubscriptionParams {
	o.SetCallTimeoutSec(callTimeoutSec)
	return o
}

// SetCallTimeoutSec adds the callTimeoutSec to the create subscription params
func (o *CreateSubscriptionParams) SetCallTimeoutSec(callTimeoutSec *int64) {
	o.CallTimeoutSec = callTimeoutSec
}

// WithEntitlementDate adds the entitlementDate to the create subscription params
func (o *CreateSubscriptionParams) WithEntitlementDate(entitlementDate *strfmt.Date) *CreateSubscriptionParams {
	o.SetEntitlementDate(entitlementDate)
	return o
}

// SetEntitlementDate adds the entitlementDate to the create subscription params
func (o *CreateSubscriptionParams) SetEntitlementDate(entitlementDate *strfmt.Date) {
	o.EntitlementDate = entitlementDate
}

// WithMigrated adds the migrated to the create subscription params
func (o *CreateSubscriptionParams) WithMigrated(migrated *bool) *CreateSubscriptionParams {
	o.SetMigrated(migrated)
	return o
}

// SetMigrated adds the migrated to the create subscription params
func (o *CreateSubscriptionParams) SetMigrated(migrated *bool) {
	o.Migrated = migrated
}

// WithPluginProperty adds the pluginProperty to the create subscription params
func (o *CreateSubscriptionParams) WithPluginProperty(pluginProperty []string) *CreateSubscriptionParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the create subscription params
func (o *CreateSubscriptionParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRenameKeyIfExistsAndUnused adds the renameKeyIfExistsAndUnused to the create subscription params
func (o *CreateSubscriptionParams) WithRenameKeyIfExistsAndUnused(renameKeyIfExistsAndUnused *bool) *CreateSubscriptionParams {
	o.SetRenameKeyIfExistsAndUnused(renameKeyIfExistsAndUnused)
	return o
}

// SetRenameKeyIfExistsAndUnused adds the renameKeyIfExistsAndUnused to the create subscription params
func (o *CreateSubscriptionParams) SetRenameKeyIfExistsAndUnused(renameKeyIfExistsAndUnused *bool) {
	o.RenameKeyIfExistsAndUnused = renameKeyIfExistsAndUnused
}

// WithSkipResponse adds the skipResponse to the create subscription params
func (o *CreateSubscriptionParams) WithSkipResponse(skipResponse *bool) *CreateSubscriptionParams {
	o.SetSkipResponse(skipResponse)
	return o
}

// SetSkipResponse adds the skipResponse to the create subscription params
func (o *CreateSubscriptionParams) SetSkipResponse(skipResponse *bool) {
	o.SkipResponse = skipResponse
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.BillingDate != nil {

		// query param billingDate
		var qrBillingDate strfmt.Date

		if o.BillingDate != nil {
			qrBillingDate = *o.BillingDate
		}
		qBillingDate := qrBillingDate.String()
		if qBillingDate != "" {

			if err := r.SetQueryParam("billingDate", qBillingDate); err != nil {
				return err
			}
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CallCompletion != nil {

		// query param callCompletion
		var qrCallCompletion bool

		if o.CallCompletion != nil {
			qrCallCompletion = *o.CallCompletion
		}
		qCallCompletion := swag.FormatBool(qrCallCompletion)
		if qCallCompletion != "" {

			if err := r.SetQueryParam("callCompletion", qCallCompletion); err != nil {
				return err
			}
		}
	}

	if o.CallTimeoutSec != nil {

		// query param callTimeoutSec
		var qrCallTimeoutSec int64

		if o.CallTimeoutSec != nil {
			qrCallTimeoutSec = *o.CallTimeoutSec
		}
		qCallTimeoutSec := swag.FormatInt64(qrCallTimeoutSec)
		if qCallTimeoutSec != "" {

			if err := r.SetQueryParam("callTimeoutSec", qCallTimeoutSec); err != nil {
				return err
			}
		}
	}

	if o.EntitlementDate != nil {

		// query param entitlementDate
		var qrEntitlementDate strfmt.Date

		if o.EntitlementDate != nil {
			qrEntitlementDate = *o.EntitlementDate
		}
		qEntitlementDate := qrEntitlementDate.String()
		if qEntitlementDate != "" {

			if err := r.SetQueryParam("entitlementDate", qEntitlementDate); err != nil {
				return err
			}
		}
	}

	if o.Migrated != nil {

		// query param migrated
		var qrMigrated bool

		if o.Migrated != nil {
			qrMigrated = *o.Migrated
		}
		qMigrated := swag.FormatBool(qrMigrated)
		if qMigrated != "" {

			if err := r.SetQueryParam("migrated", qMigrated); err != nil {
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

	if o.RenameKeyIfExistsAndUnused != nil {

		// query param renameKeyIfExistsAndUnused
		var qrRenameKeyIfExistsAndUnused bool

		if o.RenameKeyIfExistsAndUnused != nil {
			qrRenameKeyIfExistsAndUnused = *o.RenameKeyIfExistsAndUnused
		}
		qRenameKeyIfExistsAndUnused := swag.FormatBool(qrRenameKeyIfExistsAndUnused)
		if qRenameKeyIfExistsAndUnused != "" {

			if err := r.SetQueryParam("renameKeyIfExistsAndUnused", qRenameKeyIfExistsAndUnused); err != nil {
				return err
			}
		}
	}

	if o.SkipResponse != nil {

		// query param skipResponse
		var qrSkipResponse bool

		if o.SkipResponse != nil {
			qrSkipResponse = *o.SkipResponse
		}
		qSkipResponse := swag.FormatBool(qrSkipResponse)
		if qSkipResponse != "" {

			if err := r.SetQueryParam("skipResponse", qSkipResponse); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCreateSubscription binds the parameter pluginProperty
func (o *CreateSubscriptionParams) bindParamPluginProperty(formats strfmt.Registry) []string {
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
