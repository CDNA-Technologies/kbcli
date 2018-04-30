// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

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

// NewCreateInvoiceItemCustomFieldsParams creates a new CreateInvoiceItemCustomFieldsParams object
// with the default values initialized.
func NewCreateInvoiceItemCustomFieldsParams() *CreateInvoiceItemCustomFieldsParams {
	var ()
	return &CreateInvoiceItemCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateInvoiceItemCustomFieldsParamsWithTimeout creates a new CreateInvoiceItemCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateInvoiceItemCustomFieldsParamsWithTimeout(timeout time.Duration) *CreateInvoiceItemCustomFieldsParams {
	var ()
	return &CreateInvoiceItemCustomFieldsParams{

		timeout: timeout,
	}
}

// NewCreateInvoiceItemCustomFieldsParamsWithContext creates a new CreateInvoiceItemCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateInvoiceItemCustomFieldsParamsWithContext(ctx context.Context) *CreateInvoiceItemCustomFieldsParams {
	var ()
	return &CreateInvoiceItemCustomFieldsParams{

		Context: ctx,
	}
}

// NewCreateInvoiceItemCustomFieldsParamsWithHTTPClient creates a new CreateInvoiceItemCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateInvoiceItemCustomFieldsParamsWithHTTPClient(client *http.Client) *CreateInvoiceItemCustomFieldsParams {
	var ()
	return &CreateInvoiceItemCustomFieldsParams{
		HTTPClient: client,
	}
}

/*CreateInvoiceItemCustomFieldsParams contains all the parameters to send to the API endpoint
for the create invoice item custom fields operation typically these are written to a http.Request
*/
type CreateInvoiceItemCustomFieldsParams struct {

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
	/*InvoiceItemID*/
	InvoiceItemID strfmt.UUID

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithTimeout(timeout time.Duration) *CreateInvoiceItemCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithContext(ctx context.Context) *CreateInvoiceItemCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithHTTPClient(client *http.Client) *CreateInvoiceItemCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreateInvoiceItemCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreateInvoiceItemCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreateInvoiceItemCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateInvoiceItemCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreateInvoiceItemCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreateInvoiceItemCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithInvoiceItemID adds the invoiceItemID to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) WithInvoiceItemID(invoiceItemID strfmt.UUID) *CreateInvoiceItemCustomFieldsParams {
	o.SetInvoiceItemID(invoiceItemID)
	return o
}

// SetInvoiceItemID adds the invoiceItemId to the create invoice item custom fields params
func (o *CreateInvoiceItemCustomFieldsParams) SetInvoiceItemID(invoiceItemID strfmt.UUID) {
	o.InvoiceItemID = invoiceItemID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateInvoiceItemCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param invoiceItemId
	if err := r.SetPathParam("invoiceItemId", o.InvoiceItemID.String()); err != nil {
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
