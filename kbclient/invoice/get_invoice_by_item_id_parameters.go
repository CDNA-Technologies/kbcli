// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewGetInvoiceByItemIDParams creates a new GetInvoiceByItemIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInvoiceByItemIDParams() *GetInvoiceByItemIDParams {
	return &GetInvoiceByItemIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceByItemIDParamsWithTimeout creates a new GetInvoiceByItemIDParams object
// with the ability to set a timeout on a request.
func NewGetInvoiceByItemIDParamsWithTimeout(timeout time.Duration) *GetInvoiceByItemIDParams {
	return &GetInvoiceByItemIDParams{
		timeout: timeout,
	}
}

// NewGetInvoiceByItemIDParamsWithContext creates a new GetInvoiceByItemIDParams object
// with the ability to set a context for a request.
func NewGetInvoiceByItemIDParamsWithContext(ctx context.Context) *GetInvoiceByItemIDParams {
	return &GetInvoiceByItemIDParams{
		Context: ctx,
	}
}

// NewGetInvoiceByItemIDParamsWithHTTPClient creates a new GetInvoiceByItemIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInvoiceByItemIDParamsWithHTTPClient(client *http.Client) *GetInvoiceByItemIDParams {
	return &GetInvoiceByItemIDParams{
		HTTPClient: client,
	}
}

/* GetInvoiceByItemIDParams contains all the parameters to send to the API endpoint
   for the get invoice by item Id operation.

   Typically these are written to a http.Request.
*/
type GetInvoiceByItemIDParams struct {

	// Audit.
	//
	// Default: "NONE"
	Audit *string

	// ItemID.
	//
	// Format: uuid
	ItemID strfmt.UUID

	// WithChildrenItems.
	WithChildrenItems *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get invoice by item Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceByItemIDParams) WithDefaults() *GetInvoiceByItemIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get invoice by item Id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInvoiceByItemIDParams) SetDefaults() {
	var (
		auditDefault = string("NONE")

		withChildrenItemsDefault = bool(false)
	)

	val := GetInvoiceByItemIDParams{
		Audit:             &auditDefault,
		WithChildrenItems: &withChildrenItemsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithTimeout(timeout time.Duration) *GetInvoiceByItemIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithContext(ctx context.Context) *GetInvoiceByItemIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithHTTPClient(client *http.Client) *GetInvoiceByItemIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithAudit(audit *string) *GetInvoiceByItemIDParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithItemID adds the itemID to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithItemID(itemID strfmt.UUID) *GetInvoiceByItemIDParams {
	o.SetItemID(itemID)
	return o
}

// SetItemID adds the itemId to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetItemID(itemID strfmt.UUID) {
	o.ItemID = itemID
}

// WithWithChildrenItems adds the withChildrenItems to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) WithWithChildrenItems(withChildrenItems *bool) *GetInvoiceByItemIDParams {
	o.SetWithChildrenItems(withChildrenItems)
	return o
}

// SetWithChildrenItems adds the withChildrenItems to the get invoice by item Id params
func (o *GetInvoiceByItemIDParams) SetWithChildrenItems(withChildrenItems *bool) {
	o.WithChildrenItems = withChildrenItems
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceByItemIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param itemId
	if err := r.SetPathParam("itemId", o.ItemID.String()); err != nil {
		return err
	}

	if o.WithChildrenItems != nil {

		// query param withChildrenItems
		var qrWithChildrenItems bool

		if o.WithChildrenItems != nil {
			qrWithChildrenItems = *o.WithChildrenItems
		}
		qWithChildrenItems := swag.FormatBool(qrWithChildrenItems)
		if qWithChildrenItems != "" {

			if err := r.SetQueryParam("withChildrenItems", qWithChildrenItems); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
