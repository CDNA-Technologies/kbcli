// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewDeleteInvoiceCustomFieldsParams creates a new DeleteInvoiceCustomFieldsParams object
// with the default values initialized.
func NewDeleteInvoiceCustomFieldsParams() *DeleteInvoiceCustomFieldsParams {
	var ()
	return &DeleteInvoiceCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteInvoiceCustomFieldsParamsWithTimeout creates a new DeleteInvoiceCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteInvoiceCustomFieldsParamsWithTimeout(timeout time.Duration) *DeleteInvoiceCustomFieldsParams {
	var ()
	return &DeleteInvoiceCustomFieldsParams{

		timeout: timeout,
	}
}

// NewDeleteInvoiceCustomFieldsParamsWithContext creates a new DeleteInvoiceCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteInvoiceCustomFieldsParamsWithContext(ctx context.Context) *DeleteInvoiceCustomFieldsParams {
	var ()
	return &DeleteInvoiceCustomFieldsParams{

		Context: ctx,
	}
}

// NewDeleteInvoiceCustomFieldsParamsWithHTTPClient creates a new DeleteInvoiceCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteInvoiceCustomFieldsParamsWithHTTPClient(client *http.Client) *DeleteInvoiceCustomFieldsParams {
	var ()
	return &DeleteInvoiceCustomFieldsParams{
		HTTPClient: client,
	}
}

/*DeleteInvoiceCustomFieldsParams contains all the parameters to send to the API endpoint
for the delete invoice custom fields operation typically these are written to a http.Request
*/
type DeleteInvoiceCustomFieldsParams struct {

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
	/*CustomField*/
	CustomField []strfmt.UUID
	/*InvoiceID*/
	InvoiceID strfmt.UUID

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithTimeout(timeout time.Duration) *DeleteInvoiceCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithContext(ctx context.Context) *DeleteInvoiceCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithHTTPClient(client *http.Client) *DeleteInvoiceCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *DeleteInvoiceCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *DeleteInvoiceCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *DeleteInvoiceCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteInvoiceCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *DeleteInvoiceCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithCustomField adds the customField to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithCustomField(customField []strfmt.UUID) *DeleteInvoiceCustomFieldsParams {
	o.SetCustomField(customField)
	return o
}

// SetCustomField adds the customField to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetCustomField(customField []strfmt.UUID) {
	o.CustomField = customField
}

// WithInvoiceID adds the invoiceID to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) WithInvoiceID(invoiceID strfmt.UUID) *DeleteInvoiceCustomFieldsParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the delete invoice custom fields params
func (o *DeleteInvoiceCustomFieldsParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInvoiceCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesCustomField []string
	for _, v := range o.CustomField {
		valuesCustomField = append(valuesCustomField, v.String())
	}

	joinedCustomField := swag.JoinByFormat(valuesCustomField, "multi")
	// query array param customField
	if err := r.SetQueryParam("customField", joinedCustomField...); err != nil {
		return err
	}

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
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
