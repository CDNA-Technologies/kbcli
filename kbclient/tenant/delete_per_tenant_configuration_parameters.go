// Code generated by go-swagger; DO NOT EDIT.

package tenant

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
)

// NewDeletePerTenantConfigurationParams creates a new DeletePerTenantConfigurationParams object
// with the default values initialized.
func NewDeletePerTenantConfigurationParams() *DeletePerTenantConfigurationParams {
	var ()
	return &DeletePerTenantConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePerTenantConfigurationParamsWithTimeout creates a new DeletePerTenantConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePerTenantConfigurationParamsWithTimeout(timeout time.Duration) *DeletePerTenantConfigurationParams {
	var ()
	return &DeletePerTenantConfigurationParams{

		timeout: timeout,
	}
}

// NewDeletePerTenantConfigurationParamsWithContext creates a new DeletePerTenantConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePerTenantConfigurationParamsWithContext(ctx context.Context) *DeletePerTenantConfigurationParams {
	var ()
	return &DeletePerTenantConfigurationParams{

		Context: ctx,
	}
}

// NewDeletePerTenantConfigurationParamsWithHTTPClient creates a new DeletePerTenantConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePerTenantConfigurationParamsWithHTTPClient(client *http.Client) *DeletePerTenantConfigurationParams {
	var ()
	return &DeletePerTenantConfigurationParams{
		HTTPClient: client,
	}
}

/*DeletePerTenantConfigurationParams contains all the parameters to send to the API endpoint
for the delete per tenant configuration operation typically these are written to a http.Request
*/
type DeletePerTenantConfigurationParams struct {

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

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) WithTimeout(timeout time.Duration) *DeletePerTenantConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) WithContext(ctx context.Context) *DeletePerTenantConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) WithHTTPClient(client *http.Client) *DeletePerTenantConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) WithXKillbillAPIKey(xKillbillAPIKey string) *DeletePerTenantConfigurationParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) WithXKillbillAPISecret(xKillbillAPISecret string) *DeletePerTenantConfigurationParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) WithXKillbillComment(xKillbillComment *string) *DeletePerTenantConfigurationParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeletePerTenantConfigurationParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) WithXKillbillReason(xKillbillReason *string) *DeletePerTenantConfigurationParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete per tenant configuration params
func (o *DeletePerTenantConfigurationParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePerTenantConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
