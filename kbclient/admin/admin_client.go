// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new admin API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for admin API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetQueueEntries(params *GetQueueEntriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetQueueEntriesOK, error)

	InvalidatesCache(params *InvalidatesCacheParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InvalidatesCacheNoContent, error)

	InvalidatesCacheByAccount(params *InvalidatesCacheByAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InvalidatesCacheByAccountNoContent, error)

	InvalidatesCacheByTenant(params *InvalidatesCacheByTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InvalidatesCacheByTenantNoContent, error)

	PutInRotation(params *PutInRotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutInRotationNoContent, error)

	PutOutOfRotation(params *PutOutOfRotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutOutOfRotationNoContent, error)

	TriggerInvoiceGenerationForParkedAccounts(params *TriggerInvoiceGenerationForParkedAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerInvoiceGenerationForParkedAccountsOK, error)

	UpdatePaymentTransactionState(params *UpdatePaymentTransactionStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePaymentTransactionStateNoContent, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetQueueEntries gets queues entries
*/
func (a *Client) GetQueueEntries(params *GetQueueEntriesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetQueueEntriesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetQueueEntriesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getQueueEntries",
		Method:             "GET",
		PathPattern:        "/1.0/kb/admin/queues",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetQueueEntriesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetQueueEntriesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getQueueEntries: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InvalidatesCache invalidates the given cache if specified otherwise invalidates all caches
*/
func (a *Client) InvalidatesCache(params *InvalidatesCacheParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InvalidatesCacheNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInvalidatesCacheParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "invalidatesCache",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/admin/cache",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InvalidatesCacheReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InvalidatesCacheNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for invalidatesCache: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InvalidatesCacheByAccount invalidates caches per account level
*/
func (a *Client) InvalidatesCacheByAccount(params *InvalidatesCacheByAccountParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InvalidatesCacheByAccountNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInvalidatesCacheByAccountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "invalidatesCacheByAccount",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/admin/cache/accounts/{accountId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InvalidatesCacheByAccountReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InvalidatesCacheByAccountNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for invalidatesCacheByAccount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InvalidatesCacheByTenant invalidates caches per tenant level
*/
func (a *Client) InvalidatesCacheByTenant(params *InvalidatesCacheByTenantParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InvalidatesCacheByTenantNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInvalidatesCacheByTenantParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "invalidatesCacheByTenant",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/admin/cache/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InvalidatesCacheByTenantReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InvalidatesCacheByTenantNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for invalidatesCacheByTenant: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutInRotation puts the host back into rotation
*/
func (a *Client) PutInRotation(params *PutInRotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutInRotationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutInRotationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putInRotation",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/admin/healthcheck",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutInRotationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutInRotationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putInRotation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutOutOfRotation puts the host out of rotation
*/
func (a *Client) PutOutOfRotation(params *PutOutOfRotationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PutOutOfRotationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutOutOfRotationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "putOutOfRotation",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/admin/healthcheck",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutOutOfRotationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutOutOfRotationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putOutOfRotation: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TriggerInvoiceGenerationForParkedAccounts triggers an invoice generation for all parked accounts
*/
func (a *Client) TriggerInvoiceGenerationForParkedAccounts(params *TriggerInvoiceGenerationForParkedAccountsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TriggerInvoiceGenerationForParkedAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTriggerInvoiceGenerationForParkedAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "triggerInvoiceGenerationForParkedAccounts",
		Method:             "POST",
		PathPattern:        "/1.0/kb/admin/invoices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TriggerInvoiceGenerationForParkedAccountsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TriggerInvoiceGenerationForParkedAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for triggerInvoiceGenerationForParkedAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePaymentTransactionState updates existing payment transaction and associated payment state
*/
func (a *Client) UpdatePaymentTransactionState(params *UpdatePaymentTransactionStateParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdatePaymentTransactionStateNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePaymentTransactionStateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updatePaymentTransactionState",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/admin/payments/{paymentId}/transactions/{paymentTransactionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePaymentTransactionStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePaymentTransactionStateNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePaymentTransactionState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
