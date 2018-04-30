// Code generated by go-swagger; DO NOT EDIT.

package payment_method

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new payment method API client.
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
Client for payment method API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

/*
CreatePaymentMethodCustomFields adds custom fields to payment method
*/
func (a *Client) CreatePaymentMethodCustomFields(ctx context.Context, params *CreatePaymentMethodCustomFieldsParams) (*CreatePaymentMethodCustomFieldsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePaymentMethodCustomFieldsParams()
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

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
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
		ID:                 "createPaymentMethodCustomFields",
		Method:             "POST",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePaymentMethodCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentMethodCustomFieldsCreated), nil

}

/*
DeletePaymentMethod deletes a payment method
*/
func (a *Client) DeletePaymentMethod(ctx context.Context, params *DeletePaymentMethodParams) (*DeletePaymentMethodNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePaymentMethodParams()
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

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
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
		ID:                 "deletePaymentMethod",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePaymentMethodReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePaymentMethodNoContent), nil

}

/*
DeletePaymentMethodCustomFields removes custom fields from payment method
*/
func (a *Client) DeletePaymentMethodCustomFields(ctx context.Context, params *DeletePaymentMethodCustomFieldsParams) (*DeletePaymentMethodCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePaymentMethodCustomFieldsParams()
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

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
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
		ID:                 "deletePaymentMethodCustomFields",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePaymentMethodCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeletePaymentMethodCustomFieldsNoContent), nil

}

/*
GetPaymentMethod retrieves a payment method by id
*/
func (a *Client) GetPaymentMethod(ctx context.Context, params *GetPaymentMethodParams) (*GetPaymentMethodOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodParams()
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

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentMethod",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentMethodReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentMethodOK), nil

}

/*
GetPaymentMethodByKey retrieves a payment method by external key
*/
func (a *Client) GetPaymentMethodByKey(ctx context.Context, params *GetPaymentMethodByKeyParams) (*GetPaymentMethodByKeyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodByKeyParams()
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

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentMethodByKey",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentMethodByKeyReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentMethodByKeyOK), nil

}

/*
GetPaymentMethodCustomFields retrieves payment method custom fields
*/
func (a *Client) GetPaymentMethodCustomFields(ctx context.Context, params *GetPaymentMethodCustomFieldsParams) (*GetPaymentMethodCustomFieldsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodCustomFieldsParams()
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
	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentMethodCustomFields",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentMethodCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentMethodCustomFieldsOK), nil

}

/*
GetPaymentMethods lists payment methods
*/
func (a *Client) GetPaymentMethods(ctx context.Context, params *GetPaymentMethodsParams) (*GetPaymentMethodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPaymentMethodsParams()
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

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPaymentMethods",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods/pagination",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPaymentMethodsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentMethodsOK), nil

}

/*
ModifyPaymentMethodCustomFields modifies custom fields to payment method
*/
func (a *Client) ModifyPaymentMethodCustomFields(ctx context.Context, params *ModifyPaymentMethodCustomFieldsParams) (*ModifyPaymentMethodCustomFieldsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyPaymentMethodCustomFieldsParams()
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

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
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
		ID:                 "modifyPaymentMethodCustomFields",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/paymentMethods/{paymentMethodId}/customFields",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyPaymentMethodCustomFieldsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyPaymentMethodCustomFieldsNoContent), nil

}

/*
SearchPaymentMethods searches payment methods
*/
func (a *Client) SearchPaymentMethods(ctx context.Context, params *SearchPaymentMethodsParams) (*SearchPaymentMethodsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchPaymentMethodsParams()
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
		ID:                 "searchPaymentMethods",
		Method:             "GET",
		PathPattern:        "/1.0/kb/paymentMethods/search/{searchKey}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SearchPaymentMethodsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SearchPaymentMethodsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
