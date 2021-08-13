// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/CDNA-Technologies/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// ChargebackReversalPaymentReader is a Reader for the ChargebackReversalPayment structure.
type ChargebackReversalPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChargebackReversalPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewChargebackReversalPaymentCreated()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewChargebackReversalPaymentCreated creates a ChargebackReversalPaymentCreated with default headers values
func NewChargebackReversalPaymentCreated() *ChargebackReversalPaymentCreated {
	return &ChargebackReversalPaymentCreated{}
}

/*ChargebackReversalPaymentCreated handles this case with default header values.

Payment transaction created successfully
*/
type ChargebackReversalPaymentCreated struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *ChargebackReversalPaymentCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentCreated  %+v", 201, o.Payload)
}

func (o *ChargebackReversalPaymentCreated) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *ChargebackReversalPaymentCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChargebackReversalPaymentBadRequest creates a ChargebackReversalPaymentBadRequest with default headers values
func NewChargebackReversalPaymentBadRequest() *ChargebackReversalPaymentBadRequest {
	return &ChargebackReversalPaymentBadRequest{}
}

/*ChargebackReversalPaymentBadRequest handles this case with default header values.

Invalid paymentId supplied
*/
type ChargebackReversalPaymentBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackReversalPaymentBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentBadRequest ", 400)
}

func (o *ChargebackReversalPaymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentPaymentRequired creates a ChargebackReversalPaymentPaymentRequired with default headers values
func NewChargebackReversalPaymentPaymentRequired() *ChargebackReversalPaymentPaymentRequired {
	return &ChargebackReversalPaymentPaymentRequired{}
}

/*ChargebackReversalPaymentPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type ChargebackReversalPaymentPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackReversalPaymentPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentPaymentRequired ", 402)
}

func (o *ChargebackReversalPaymentPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentNotFound creates a ChargebackReversalPaymentNotFound with default headers values
func NewChargebackReversalPaymentNotFound() *ChargebackReversalPaymentNotFound {
	return &ChargebackReversalPaymentNotFound{}
}

/*ChargebackReversalPaymentNotFound handles this case with default header values.

Account or payment not found
*/
type ChargebackReversalPaymentNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackReversalPaymentNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentNotFound ", 404)
}

func (o *ChargebackReversalPaymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentUnprocessableEntity creates a ChargebackReversalPaymentUnprocessableEntity with default headers values
func NewChargebackReversalPaymentUnprocessableEntity() *ChargebackReversalPaymentUnprocessableEntity {
	return &ChargebackReversalPaymentUnprocessableEntity{}
}

/*ChargebackReversalPaymentUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type ChargebackReversalPaymentUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackReversalPaymentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentUnprocessableEntity ", 422)
}

func (o *ChargebackReversalPaymentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentBadGateway creates a ChargebackReversalPaymentBadGateway with default headers values
func NewChargebackReversalPaymentBadGateway() *ChargebackReversalPaymentBadGateway {
	return &ChargebackReversalPaymentBadGateway{}
}

/*ChargebackReversalPaymentBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type ChargebackReversalPaymentBadGateway struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackReversalPaymentBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentBadGateway ", 502)
}

func (o *ChargebackReversalPaymentBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentServiceUnavailable creates a ChargebackReversalPaymentServiceUnavailable with default headers values
func NewChargebackReversalPaymentServiceUnavailable() *ChargebackReversalPaymentServiceUnavailable {
	return &ChargebackReversalPaymentServiceUnavailable{}
}

/*ChargebackReversalPaymentServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type ChargebackReversalPaymentServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackReversalPaymentServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentServiceUnavailable ", 503)
}

func (o *ChargebackReversalPaymentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewChargebackReversalPaymentGatewayTimeout creates a ChargebackReversalPaymentGatewayTimeout with default headers values
func NewChargebackReversalPaymentGatewayTimeout() *ChargebackReversalPaymentGatewayTimeout {
	return &ChargebackReversalPaymentGatewayTimeout{}
}

/*ChargebackReversalPaymentGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type ChargebackReversalPaymentGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

func (o *ChargebackReversalPaymentGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/chargebackReversals][%d] chargebackReversalPaymentGatewayTimeout ", 504)
}

func (o *ChargebackReversalPaymentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
