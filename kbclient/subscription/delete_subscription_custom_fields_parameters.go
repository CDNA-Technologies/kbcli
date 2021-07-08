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
)

// NewDeleteSubscriptionCustomFieldsParams creates a new DeleteSubscriptionCustomFieldsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSubscriptionCustomFieldsParams() *DeleteSubscriptionCustomFieldsParams {
	return &DeleteSubscriptionCustomFieldsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSubscriptionCustomFieldsParamsWithTimeout creates a new DeleteSubscriptionCustomFieldsParams object
// with the ability to set a timeout on a request.
func NewDeleteSubscriptionCustomFieldsParamsWithTimeout(timeout time.Duration) *DeleteSubscriptionCustomFieldsParams {
	return &DeleteSubscriptionCustomFieldsParams{
		timeout: timeout,
	}
}

// NewDeleteSubscriptionCustomFieldsParamsWithContext creates a new DeleteSubscriptionCustomFieldsParams object
// with the ability to set a context for a request.
func NewDeleteSubscriptionCustomFieldsParamsWithContext(ctx context.Context) *DeleteSubscriptionCustomFieldsParams {
	return &DeleteSubscriptionCustomFieldsParams{
		Context: ctx,
	}
}

// NewDeleteSubscriptionCustomFieldsParamsWithHTTPClient creates a new DeleteSubscriptionCustomFieldsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSubscriptionCustomFieldsParamsWithHTTPClient(client *http.Client) *DeleteSubscriptionCustomFieldsParams {
	return &DeleteSubscriptionCustomFieldsParams{
		HTTPClient: client,
	}
}

/* DeleteSubscriptionCustomFieldsParams contains all the parameters to send to the API endpoint
   for the delete subscription custom fields operation.

   Typically these are written to a http.Request.
*/
type DeleteSubscriptionCustomFieldsParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// CustomField.
	CustomField []strfmt.UUID

	// SubscriptionID.
	//
	// Format: uuid
	SubscriptionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete subscription custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSubscriptionCustomFieldsParams) WithDefaults() *DeleteSubscriptionCustomFieldsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete subscription custom fields params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSubscriptionCustomFieldsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) WithTimeout(timeout time.Duration) *DeleteSubscriptionCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) WithContext(ctx context.Context) *DeleteSubscriptionCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) WithHTTPClient(client *http.Client) *DeleteSubscriptionCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *DeleteSubscriptionCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteSubscriptionCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *DeleteSubscriptionCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithCustomField adds the customField to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) WithCustomField(customField []strfmt.UUID) *DeleteSubscriptionCustomFieldsParams {
	o.SetCustomField(customField)
	return o
}

// SetCustomField adds the customField to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) SetCustomField(customField []strfmt.UUID) {
	o.CustomField = customField
}

// WithSubscriptionID adds the subscriptionID to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) WithSubscriptionID(subscriptionID strfmt.UUID) *DeleteSubscriptionCustomFieldsParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the delete subscription custom fields params
func (o *DeleteSubscriptionCustomFieldsParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSubscriptionCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.CustomField != nil {

		// binding items for customField
		joinedCustomField := o.bindParamCustomField(reg)

		// query array param customField
		if err := r.SetQueryParam("customField", joinedCustomField...); err != nil {
			return err
		}
	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamDeleteSubscriptionCustomFields binds the parameter customField
func (o *DeleteSubscriptionCustomFieldsParams) bindParamCustomField(formats strfmt.Registry) []string {
	customFieldIR := o.CustomField

	var customFieldIC []string
	for _, customFieldIIR := range customFieldIR { // explode []strfmt.UUID

		customFieldIIV := customFieldIIR.String() // strfmt.UUID as string
		customFieldIC = append(customFieldIC, customFieldIIV)
	}

	// items.CollectionFormat: "multi"
	customFieldIS := swag.JoinByFormat(customFieldIC, "multi")

	return customFieldIS
}
