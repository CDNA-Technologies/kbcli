// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tag API client.
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *Client {

	return &Client{transport: transport, formats: formats, authInfo: authInfo, defaults: defaults}
}

// killbill default values. When a call is made to an operation, these values are used
// if params doesn't specify them.
type KillbillDefaults interface {
	// Default API Key. If not set explicitly in params, this will be used.
	XKillbillAPIKey() *string
	// Default API Secret. If not set explicitly in params, this will be used.
	XKillbillAPISecret() *string
	// Default CreatedBy. If not set explicitly in params, this will be used.
	XKillbillCreatedBy() *string
	// Default Comment. If not set explicitly in params, this will be used.
	XKillbillComment() *string
	// Default Reason. If not set explicitly in params, this will be used.
	XKillbillReason() *string
	// Default WithStackTrace. If not set explicitly in params, this will be used.
	KillbillWithStackTrace() *bool
}

/*
Client for tag API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

/*
GetTags lists tags
*/
func (a *Client) GetTags(ctx context.Context, params *GetTagsParams) (*GetTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTagsParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTags",
		Method:             "GET",
		PathPattern:        "/1.0/kb/tags/pagination",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTagsOK), nil

}

/*
SearchTags searches tags
*/
func (a *Client) SearchTags(ctx context.Context, params *SearchTagsParams) (*SearchTagsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchTagsParams()
	}
	params.Context = ctx
	if params.XKillbillAPIKey == "" && a.defaults.XKillbillAPIKey() != nil {
		params.XKillbillAPIKey = *a.defaults.XKillbillAPIKey()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	if params.XKillbillAPISecret == "" && a.defaults.XKillbillAPISecret() != nil {
		params.XKillbillAPISecret = *a.defaults.XKillbillAPISecret()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "searchTags",
		Method:             "GET",
		PathPattern:        "/1.0/kb/tags/search/{searchKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchTagsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchTagsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
