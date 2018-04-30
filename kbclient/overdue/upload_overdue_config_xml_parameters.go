// Code generated by go-swagger; DO NOT EDIT.

package overdue

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

// NewUploadOverdueConfigXMLParams creates a new UploadOverdueConfigXMLParams object
// with the default values initialized.
func NewUploadOverdueConfigXMLParams() *UploadOverdueConfigXMLParams {
	var ()
	return &UploadOverdueConfigXMLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadOverdueConfigXMLParamsWithTimeout creates a new UploadOverdueConfigXMLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadOverdueConfigXMLParamsWithTimeout(timeout time.Duration) *UploadOverdueConfigXMLParams {
	var ()
	return &UploadOverdueConfigXMLParams{

		timeout: timeout,
	}
}

// NewUploadOverdueConfigXMLParamsWithContext creates a new UploadOverdueConfigXMLParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadOverdueConfigXMLParamsWithContext(ctx context.Context) *UploadOverdueConfigXMLParams {
	var ()
	return &UploadOverdueConfigXMLParams{

		Context: ctx,
	}
}

// NewUploadOverdueConfigXMLParamsWithHTTPClient creates a new UploadOverdueConfigXMLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadOverdueConfigXMLParamsWithHTTPClient(client *http.Client) *UploadOverdueConfigXMLParams {
	var ()
	return &UploadOverdueConfigXMLParams{
		HTTPClient: client,
	}
}

/*UploadOverdueConfigXMLParams contains all the parameters to send to the API endpoint
for the upload overdue config Xml operation typically these are written to a http.Request
*/
type UploadOverdueConfigXMLParams struct {

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
	Body *string

	WithStackTrace *bool
	timeout        time.Duration
	Context        context.Context
	HTTPClient     *http.Client
}

// WithTimeout adds the timeout to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithTimeout(timeout time.Duration) *UploadOverdueConfigXMLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithContext(ctx context.Context) *UploadOverdueConfigXMLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithHTTPClient(client *http.Client) *UploadOverdueConfigXMLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithXKillbillAPIKey(xKillbillAPIKey string) *UploadOverdueConfigXMLParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithXKillbillAPISecret(xKillbillAPISecret string) *UploadOverdueConfigXMLParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithXKillbillComment(xKillbillComment *string) *UploadOverdueConfigXMLParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *UploadOverdueConfigXMLParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithXKillbillReason(xKillbillReason *string) *UploadOverdueConfigXMLParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) WithBody(body *string) *UploadOverdueConfigXMLParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload overdue config Xml params
func (o *UploadOverdueConfigXMLParams) SetBody(body *string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UploadOverdueConfigXMLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
