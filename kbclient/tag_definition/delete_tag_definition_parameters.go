// Code generated by go-swagger; DO NOT EDIT.

package tag_definition

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
)

// NewDeleteTagDefinitionParams creates a new DeleteTagDefinitionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTagDefinitionParams() *DeleteTagDefinitionParams {
	return &DeleteTagDefinitionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTagDefinitionParamsWithTimeout creates a new DeleteTagDefinitionParams object
// with the ability to set a timeout on a request.
func NewDeleteTagDefinitionParamsWithTimeout(timeout time.Duration) *DeleteTagDefinitionParams {
	return &DeleteTagDefinitionParams{
		timeout: timeout,
	}
}

// NewDeleteTagDefinitionParamsWithContext creates a new DeleteTagDefinitionParams object
// with the ability to set a context for a request.
func NewDeleteTagDefinitionParamsWithContext(ctx context.Context) *DeleteTagDefinitionParams {
	return &DeleteTagDefinitionParams{
		Context: ctx,
	}
}

// NewDeleteTagDefinitionParamsWithHTTPClient creates a new DeleteTagDefinitionParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTagDefinitionParamsWithHTTPClient(client *http.Client) *DeleteTagDefinitionParams {
	return &DeleteTagDefinitionParams{
		HTTPClient: client,
	}
}

/* DeleteTagDefinitionParams contains all the parameters to send to the API endpoint
   for the delete tag definition operation.

   Typically these are written to a http.Request.
*/
type DeleteTagDefinitionParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// TagDefinitionID.
	//
	// Format: uuid
	TagDefinitionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete tag definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTagDefinitionParams) WithDefaults() *DeleteTagDefinitionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete tag definition params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTagDefinitionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete tag definition params
func (o *DeleteTagDefinitionParams) WithTimeout(timeout time.Duration) *DeleteTagDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete tag definition params
func (o *DeleteTagDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete tag definition params
func (o *DeleteTagDefinitionParams) WithContext(ctx context.Context) *DeleteTagDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete tag definition params
func (o *DeleteTagDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete tag definition params
func (o *DeleteTagDefinitionParams) WithHTTPClient(client *http.Client) *DeleteTagDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete tag definition params
func (o *DeleteTagDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete tag definition params
func (o *DeleteTagDefinitionParams) WithXKillbillComment(xKillbillComment *string) *DeleteTagDefinitionParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete tag definition params
func (o *DeleteTagDefinitionParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete tag definition params
func (o *DeleteTagDefinitionParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteTagDefinitionParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete tag definition params
func (o *DeleteTagDefinitionParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete tag definition params
func (o *DeleteTagDefinitionParams) WithXKillbillReason(xKillbillReason *string) *DeleteTagDefinitionParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete tag definition params
func (o *DeleteTagDefinitionParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithTagDefinitionID adds the tagDefinitionID to the delete tag definition params
func (o *DeleteTagDefinitionParams) WithTagDefinitionID(tagDefinitionID strfmt.UUID) *DeleteTagDefinitionParams {
	o.SetTagDefinitionID(tagDefinitionID)
	return o
}

// SetTagDefinitionID adds the tagDefinitionId to the delete tag definition params
func (o *DeleteTagDefinitionParams) SetTagDefinitionID(tagDefinitionID strfmt.UUID) {
	o.TagDefinitionID = tagDefinitionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTagDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tagDefinitionId
	if err := r.SetPathParam("tagDefinitionId", o.TagDefinitionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
