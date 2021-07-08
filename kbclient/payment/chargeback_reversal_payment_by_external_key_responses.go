// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// ChargebackReversalPaymentByExternalKeyReader is a Reader for the ChargebackReversalPaymentByExternalKey structure.
type ChargebackReversalPaymentByExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChargebackReversalPaymentByExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewChargebackReversalPaymentByExternalKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 402:
		result := NewChargebackReversalPaymentByExternalKeyPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChargebackReversalPaymentByExternalKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewChargebackReversalPaymentByExternalKeyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewChargebackReversalPaymentByExternalKeyBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewChargebackReversalPaymentByExternalKeyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewChargebackReversalPaymentByExternalKeyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChargebackReversalPaymentByExternalKeyCreated creates a ChargebackReversalPaymentByExternalKeyCreated with default headers values
func NewChargebackReversalPaymentByExternalKeyCreated() *ChargebackReversalPaymentByExternalKeyCreated {
	return &ChargebackReversalPaymentByExternalKeyCreated{}
}

/* ChargebackReversalPaymentByExternalKeyCreated describes a response with status code 201, with default header values.

Payment transaction created successfully
*/
type ChargebackReversalPaymentByExternalKeyCreated struct {
	Payload *kbmodel.Payment
}

func (o *ChargebackReversalPaymentByExternalKeyCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebackReversals][%d] chargebackReversalPaymentByExternalKeyCreated  %+v", 201, o.Payload)
}
func (o *ChargebackReversalPaymentByExternalKeyCreated) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *ChargebackReversalPaymentByExternalKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChargebackReversalPaymentByExternalKeyPaymentRequired creates a ChargebackReversalPaymentByExternalKeyPaymentRequired with default headers values
func NewChargebackReversalPaymentByExternalKeyPaymentRequired() *ChargebackReversalPaymentByExternalKeyPaymentRequired {
	return &ChargebackReversalPaymentByExternalKeyPaymentRequired{}
}

/* ChargebackReversalPaymentByExternalKeyPaymentRequired describes a response with status code 402, with default header values.

Transaction declined by gateway
*/
type ChargebackReversalPaymentByExternalKeyPaymentRequired struct {
}

func (o *ChargebackReversalPaymentByExternalKeyPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebackReversals][%d] chargebackReversalPaymentByExternalKeyPaymentRequired ", 402)
}

func (o *ChargebackReversalPaymentByExternalKeyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentByExternalKeyNotFound creates a ChargebackReversalPaymentByExternalKeyNotFound with default headers values
func NewChargebackReversalPaymentByExternalKeyNotFound() *ChargebackReversalPaymentByExternalKeyNotFound {
	return &ChargebackReversalPaymentByExternalKeyNotFound{}
}

/* ChargebackReversalPaymentByExternalKeyNotFound describes a response with status code 404, with default header values.

Account or payment not found
*/
type ChargebackReversalPaymentByExternalKeyNotFound struct {
}

func (o *ChargebackReversalPaymentByExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebackReversals][%d] chargebackReversalPaymentByExternalKeyNotFound ", 404)
}

func (o *ChargebackReversalPaymentByExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentByExternalKeyUnprocessableEntity creates a ChargebackReversalPaymentByExternalKeyUnprocessableEntity with default headers values
func NewChargebackReversalPaymentByExternalKeyUnprocessableEntity() *ChargebackReversalPaymentByExternalKeyUnprocessableEntity {
	return &ChargebackReversalPaymentByExternalKeyUnprocessableEntity{}
}

/* ChargebackReversalPaymentByExternalKeyUnprocessableEntity describes a response with status code 422, with default header values.

Payment is aborted by a control plugin
*/
type ChargebackReversalPaymentByExternalKeyUnprocessableEntity struct {
}

func (o *ChargebackReversalPaymentByExternalKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebackReversals][%d] chargebackReversalPaymentByExternalKeyUnprocessableEntity ", 422)
}

func (o *ChargebackReversalPaymentByExternalKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentByExternalKeyBadGateway creates a ChargebackReversalPaymentByExternalKeyBadGateway with default headers values
func NewChargebackReversalPaymentByExternalKeyBadGateway() *ChargebackReversalPaymentByExternalKeyBadGateway {
	return &ChargebackReversalPaymentByExternalKeyBadGateway{}
}

/* ChargebackReversalPaymentByExternalKeyBadGateway describes a response with status code 502, with default header values.

Failed to submit payment transaction
*/
type ChargebackReversalPaymentByExternalKeyBadGateway struct {
}

func (o *ChargebackReversalPaymentByExternalKeyBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebackReversals][%d] chargebackReversalPaymentByExternalKeyBadGateway ", 502)
}

func (o *ChargebackReversalPaymentByExternalKeyBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentByExternalKeyServiceUnavailable creates a ChargebackReversalPaymentByExternalKeyServiceUnavailable with default headers values
func NewChargebackReversalPaymentByExternalKeyServiceUnavailable() *ChargebackReversalPaymentByExternalKeyServiceUnavailable {
	return &ChargebackReversalPaymentByExternalKeyServiceUnavailable{}
}

/* ChargebackReversalPaymentByExternalKeyServiceUnavailable describes a response with status code 503, with default header values.

Payment in unknown status, failed to receive gateway response
*/
type ChargebackReversalPaymentByExternalKeyServiceUnavailable struct {
}

func (o *ChargebackReversalPaymentByExternalKeyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebackReversals][%d] chargebackReversalPaymentByExternalKeyServiceUnavailable ", 503)
}

func (o *ChargebackReversalPaymentByExternalKeyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentByExternalKeyGatewayTimeout creates a ChargebackReversalPaymentByExternalKeyGatewayTimeout with default headers values
func NewChargebackReversalPaymentByExternalKeyGatewayTimeout() *ChargebackReversalPaymentByExternalKeyGatewayTimeout {
	return &ChargebackReversalPaymentByExternalKeyGatewayTimeout{}
}

/* ChargebackReversalPaymentByExternalKeyGatewayTimeout describes a response with status code 504, with default header values.

Payment operation timeout
*/
type ChargebackReversalPaymentByExternalKeyGatewayTimeout struct {
}

func (o *ChargebackReversalPaymentByExternalKeyGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebackReversals][%d] chargebackReversalPaymentByExternalKeyGatewayTimeout ", 504)
}

func (o *ChargebackReversalPaymentByExternalKeyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
