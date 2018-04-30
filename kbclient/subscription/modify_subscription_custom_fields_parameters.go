// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewModifySubscriptionCustomFieldsParams creates a new ModifySubscriptionCustomFieldsParams object
// with the default values initialized.
func NewModifySubscriptionCustomFieldsParams() *ModifySubscriptionCustomFieldsParams {
	var ()
	return &ModifySubscriptionCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifySubscriptionCustomFieldsParamsWithTimeout creates a new ModifySubscriptionCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifySubscriptionCustomFieldsParamsWithTimeout(timeout time.Duration) *ModifySubscriptionCustomFieldsParams {
	var ()
	return &ModifySubscriptionCustomFieldsParams{

		timeout: timeout,
	}
}

// NewModifySubscriptionCustomFieldsParamsWithContext creates a new ModifySubscriptionCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifySubscriptionCustomFieldsParamsWithContext(ctx context.Context) *ModifySubscriptionCustomFieldsParams {
	var ()
	return &ModifySubscriptionCustomFieldsParams{

		Context: ctx,
	}
}

// NewModifySubscriptionCustomFieldsParamsWithHTTPClient creates a new ModifySubscriptionCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifySubscriptionCustomFieldsParamsWithHTTPClient(client *http.Client) *ModifySubscriptionCustomFieldsParams {
	var ()
	return &ModifySubscriptionCustomFieldsParams{
		HTTPClient: client,
	}
}

/*ModifySubscriptionCustomFieldsParams contains all the parameters to send to the API endpoint
for the modify subscription custom fields operation typically these are written to a http.Request
*/
type ModifySubscriptionCustomFieldsParams struct {

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
	/*Body*/
	Body []*kbmodel.CustomField
	/*SubscriptionID*/
	SubscriptionID strfmt.UUID

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithTimeout(timeout time.Duration) *ModifySubscriptionCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithContext(ctx context.Context) *ModifySubscriptionCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithHTTPClient(client *http.Client) *ModifySubscriptionCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *ModifySubscriptionCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *ModifySubscriptionCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *ModifySubscriptionCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ModifySubscriptionCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *ModifySubscriptionCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *ModifySubscriptionCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithSubscriptionID adds the subscriptionID to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) WithSubscriptionID(subscriptionID strfmt.UUID) *ModifySubscriptionCustomFieldsParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the modify subscription custom fields params
func (o *ModifySubscriptionCustomFieldsParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifySubscriptionCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID.String()); err != nil {
		return err
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
