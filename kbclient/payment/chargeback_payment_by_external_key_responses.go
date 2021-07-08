// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// ChargebackPaymentByExternalKeyReader is a Reader for the ChargebackPaymentByExternalKey structure.
type ChargebackPaymentByExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChargebackPaymentByExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewChargebackPaymentByExternalKeyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 402:
		result := NewChargebackPaymentByExternalKeyPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewChargebackPaymentByExternalKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewChargebackPaymentByExternalKeyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewChargebackPaymentByExternalKeyBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewChargebackPaymentByExternalKeyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewChargebackPaymentByExternalKeyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewChargebackPaymentByExternalKeyCreated creates a ChargebackPaymentByExternalKeyCreated with default headers values
func NewChargebackPaymentByExternalKeyCreated() *ChargebackPaymentByExternalKeyCreated {
	return &ChargebackPaymentByExternalKeyCreated{}
}

/* ChargebackPaymentByExternalKeyCreated describes a response with status code 201, with default header values.

Payment transaction created successfully
*/
type ChargebackPaymentByExternalKeyCreated struct {
	Payload *kbmodel.Payment
}

func (o *ChargebackPaymentByExternalKeyCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyCreated  %+v", 201, o.Payload)
}
func (o *ChargebackPaymentByExternalKeyCreated) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *ChargebackPaymentByExternalKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChargebackPaymentByExternalKeyPaymentRequired creates a ChargebackPaymentByExternalKeyPaymentRequired with default headers values
func NewChargebackPaymentByExternalKeyPaymentRequired() *ChargebackPaymentByExternalKeyPaymentRequired {
	return &ChargebackPaymentByExternalKeyPaymentRequired{}
}

/* ChargebackPaymentByExternalKeyPaymentRequired describes a response with status code 402, with default header values.

Transaction declined by gateway
*/
type ChargebackPaymentByExternalKeyPaymentRequired struct {
}

func (o *ChargebackPaymentByExternalKeyPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyPaymentRequired ", 402)
}

func (o *ChargebackPaymentByExternalKeyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyNotFound creates a ChargebackPaymentByExternalKeyNotFound with default headers values
func NewChargebackPaymentByExternalKeyNotFound() *ChargebackPaymentByExternalKeyNotFound {
	return &ChargebackPaymentByExternalKeyNotFound{}
}

/* ChargebackPaymentByExternalKeyNotFound describes a response with status code 404, with default header values.

Account or payment not found
*/
type ChargebackPaymentByExternalKeyNotFound struct {
}

func (o *ChargebackPaymentByExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyNotFound ", 404)
}

func (o *ChargebackPaymentByExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyUnprocessableEntity creates a ChargebackPaymentByExternalKeyUnprocessableEntity with default headers values
func NewChargebackPaymentByExternalKeyUnprocessableEntity() *ChargebackPaymentByExternalKeyUnprocessableEntity {
	return &ChargebackPaymentByExternalKeyUnprocessableEntity{}
}

/* ChargebackPaymentByExternalKeyUnprocessableEntity describes a response with status code 422, with default header values.

Payment is aborted by a control plugin
*/
type ChargebackPaymentByExternalKeyUnprocessableEntity struct {
}

func (o *ChargebackPaymentByExternalKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyUnprocessableEntity ", 422)
}

func (o *ChargebackPaymentByExternalKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyBadGateway creates a ChargebackPaymentByExternalKeyBadGateway with default headers values
func NewChargebackPaymentByExternalKeyBadGateway() *ChargebackPaymentByExternalKeyBadGateway {
	return &ChargebackPaymentByExternalKeyBadGateway{}
}

/* ChargebackPaymentByExternalKeyBadGateway describes a response with status code 502, with default header values.

Failed to submit payment transaction
*/
type ChargebackPaymentByExternalKeyBadGateway struct {
}

func (o *ChargebackPaymentByExternalKeyBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyBadGateway ", 502)
}

func (o *ChargebackPaymentByExternalKeyBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyServiceUnavailable creates a ChargebackPaymentByExternalKeyServiceUnavailable with default headers values
func NewChargebackPaymentByExternalKeyServiceUnavailable() *ChargebackPaymentByExternalKeyServiceUnavailable {
	return &ChargebackPaymentByExternalKeyServiceUnavailable{}
}

/* ChargebackPaymentByExternalKeyServiceUnavailable describes a response with status code 503, with default header values.

Payment in unknown status, failed to receive gateway response
*/
type ChargebackPaymentByExternalKeyServiceUnavailable struct {
}

func (o *ChargebackPaymentByExternalKeyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyServiceUnavailable ", 503)
}

func (o *ChargebackPaymentByExternalKeyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackPaymentByExternalKeyGatewayTimeout creates a ChargebackPaymentByExternalKeyGatewayTimeout with default headers values
func NewChargebackPaymentByExternalKeyGatewayTimeout() *ChargebackPaymentByExternalKeyGatewayTimeout {
	return &ChargebackPaymentByExternalKeyGatewayTimeout{}
}

/* ChargebackPaymentByExternalKeyGatewayTimeout describes a response with status code 504, with default header values.

Payment operation timeout
*/
type ChargebackPaymentByExternalKeyGatewayTimeout struct {
}

func (o *ChargebackPaymentByExternalKeyGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/chargebacks][%d] chargebackPaymentByExternalKeyGatewayTimeout ", 504)
}

func (o *ChargebackPaymentByExternalKeyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
