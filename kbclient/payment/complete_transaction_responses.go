// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CompleteTransactionReader is a Reader for the CompleteTransaction structure.
type CompleteTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCompleteTransactionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCompleteTransactionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 402:
		result := NewCompleteTransactionPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCompleteTransactionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCompleteTransactionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 502:
		result := NewCompleteTransactionBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewCompleteTransactionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewCompleteTransactionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCompleteTransactionNoContent creates a CompleteTransactionNoContent with default headers values
func NewCompleteTransactionNoContent() *CompleteTransactionNoContent {
	return &CompleteTransactionNoContent{}
}

/* CompleteTransactionNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type CompleteTransactionNoContent struct {
}

func (o *CompleteTransactionNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}][%d] completeTransactionNoContent ", 204)
}

func (o *CompleteTransactionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionBadRequest creates a CompleteTransactionBadRequest with default headers values
func NewCompleteTransactionBadRequest() *CompleteTransactionBadRequest {
	return &CompleteTransactionBadRequest{}
}

/* CompleteTransactionBadRequest describes a response with status code 400, with default header values.

Invalid paymentId supplied
*/
type CompleteTransactionBadRequest struct {
}

func (o *CompleteTransactionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}][%d] completeTransactionBadRequest ", 400)
}

func (o *CompleteTransactionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionPaymentRequired creates a CompleteTransactionPaymentRequired with default headers values
func NewCompleteTransactionPaymentRequired() *CompleteTransactionPaymentRequired {
	return &CompleteTransactionPaymentRequired{}
}

/* CompleteTransactionPaymentRequired describes a response with status code 402, with default header values.

Transaction declined by gateway
*/
type CompleteTransactionPaymentRequired struct {
}

func (o *CompleteTransactionPaymentRequired) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}][%d] completeTransactionPaymentRequired ", 402)
}

func (o *CompleteTransactionPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionNotFound creates a CompleteTransactionNotFound with default headers values
func NewCompleteTransactionNotFound() *CompleteTransactionNotFound {
	return &CompleteTransactionNotFound{}
}

/* CompleteTransactionNotFound describes a response with status code 404, with default header values.

Account or payment not found
*/
type CompleteTransactionNotFound struct {
}

func (o *CompleteTransactionNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}][%d] completeTransactionNotFound ", 404)
}

func (o *CompleteTransactionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionUnprocessableEntity creates a CompleteTransactionUnprocessableEntity with default headers values
func NewCompleteTransactionUnprocessableEntity() *CompleteTransactionUnprocessableEntity {
	return &CompleteTransactionUnprocessableEntity{}
}

/* CompleteTransactionUnprocessableEntity describes a response with status code 422, with default header values.

Payment is aborted by a control plugin
*/
type CompleteTransactionUnprocessableEntity struct {
}

func (o *CompleteTransactionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}][%d] completeTransactionUnprocessableEntity ", 422)
}

func (o *CompleteTransactionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionBadGateway creates a CompleteTransactionBadGateway with default headers values
func NewCompleteTransactionBadGateway() *CompleteTransactionBadGateway {
	return &CompleteTransactionBadGateway{}
}

/* CompleteTransactionBadGateway describes a response with status code 502, with default header values.

Failed to submit payment transaction
*/
type CompleteTransactionBadGateway struct {
}

func (o *CompleteTransactionBadGateway) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}][%d] completeTransactionBadGateway ", 502)
}

func (o *CompleteTransactionBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionServiceUnavailable creates a CompleteTransactionServiceUnavailable with default headers values
func NewCompleteTransactionServiceUnavailable() *CompleteTransactionServiceUnavailable {
	return &CompleteTransactionServiceUnavailable{}
}

/* CompleteTransactionServiceUnavailable describes a response with status code 503, with default header values.

Payment in unknown status, failed to receive gateway response
*/
type CompleteTransactionServiceUnavailable struct {
}

func (o *CompleteTransactionServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}][%d] completeTransactionServiceUnavailable ", 503)
}

func (o *CompleteTransactionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionGatewayTimeout creates a CompleteTransactionGatewayTimeout with default headers values
func NewCompleteTransactionGatewayTimeout() *CompleteTransactionGatewayTimeout {
	return &CompleteTransactionGatewayTimeout{}
}

/* CompleteTransactionGatewayTimeout describes a response with status code 504, with default header values.

Payment operation timeout
*/
type CompleteTransactionGatewayTimeout struct {
}

func (o *CompleteTransactionGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}][%d] completeTransactionGatewayTimeout ", 504)
}

func (o *CompleteTransactionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
