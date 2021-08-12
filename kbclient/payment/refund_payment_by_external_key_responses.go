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

// RefundPaymentByExternalKeyReader is a Reader for the RefundPaymentByExternalKey structure.
type RefundPaymentByExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefundPaymentByExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewRefundPaymentByExternalKeyCreated()
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

// NewRefundPaymentByExternalKeyCreated creates a RefundPaymentByExternalKeyCreated with default headers values
func NewRefundPaymentByExternalKeyCreated() *RefundPaymentByExternalKeyCreated {
	return &RefundPaymentByExternalKeyCreated{}
}

/*RefundPaymentByExternalKeyCreated handles this case with default header values.

Payment transaction created successfully
*/
type RefundPaymentByExternalKeyCreated struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *RefundPaymentByExternalKeyCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/refunds][%d] refundPaymentByExternalKeyCreated  %+v", 201, o.Payload)
}

func (o *RefundPaymentByExternalKeyCreated) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *RefundPaymentByExternalKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRefundPaymentByExternalKeyPaymentRequired creates a RefundPaymentByExternalKeyPaymentRequired with default headers values
func NewRefundPaymentByExternalKeyPaymentRequired() *RefundPaymentByExternalKeyPaymentRequired {
	return &RefundPaymentByExternalKeyPaymentRequired{}
}

/*RefundPaymentByExternalKeyPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type RefundPaymentByExternalKeyPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

func (o *RefundPaymentByExternalKeyPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/refunds][%d] refundPaymentByExternalKeyPaymentRequired ", 402)
}

func (o *RefundPaymentByExternalKeyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefundPaymentByExternalKeyNotFound creates a RefundPaymentByExternalKeyNotFound with default headers values
func NewRefundPaymentByExternalKeyNotFound() *RefundPaymentByExternalKeyNotFound {
	return &RefundPaymentByExternalKeyNotFound{}
}

/*RefundPaymentByExternalKeyNotFound handles this case with default header values.

Account or payment not found
*/
type RefundPaymentByExternalKeyNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *RefundPaymentByExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/refunds][%d] refundPaymentByExternalKeyNotFound ", 404)
}

func (o *RefundPaymentByExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefundPaymentByExternalKeyUnprocessableEntity creates a RefundPaymentByExternalKeyUnprocessableEntity with default headers values
func NewRefundPaymentByExternalKeyUnprocessableEntity() *RefundPaymentByExternalKeyUnprocessableEntity {
	return &RefundPaymentByExternalKeyUnprocessableEntity{}
}

/*RefundPaymentByExternalKeyUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type RefundPaymentByExternalKeyUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

func (o *RefundPaymentByExternalKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/refunds][%d] refundPaymentByExternalKeyUnprocessableEntity ", 422)
}

func (o *RefundPaymentByExternalKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefundPaymentByExternalKeyBadGateway creates a RefundPaymentByExternalKeyBadGateway with default headers values
func NewRefundPaymentByExternalKeyBadGateway() *RefundPaymentByExternalKeyBadGateway {
	return &RefundPaymentByExternalKeyBadGateway{}
}

/*RefundPaymentByExternalKeyBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type RefundPaymentByExternalKeyBadGateway struct {
	HttpResponse runtime.ClientResponse
}

func (o *RefundPaymentByExternalKeyBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/refunds][%d] refundPaymentByExternalKeyBadGateway ", 502)
}

func (o *RefundPaymentByExternalKeyBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefundPaymentByExternalKeyServiceUnavailable creates a RefundPaymentByExternalKeyServiceUnavailable with default headers values
func NewRefundPaymentByExternalKeyServiceUnavailable() *RefundPaymentByExternalKeyServiceUnavailable {
	return &RefundPaymentByExternalKeyServiceUnavailable{}
}

/*RefundPaymentByExternalKeyServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type RefundPaymentByExternalKeyServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

func (o *RefundPaymentByExternalKeyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/refunds][%d] refundPaymentByExternalKeyServiceUnavailable ", 503)
}

func (o *RefundPaymentByExternalKeyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefundPaymentByExternalKeyGatewayTimeout creates a RefundPaymentByExternalKeyGatewayTimeout with default headers values
func NewRefundPaymentByExternalKeyGatewayTimeout() *RefundPaymentByExternalKeyGatewayTimeout {
	return &RefundPaymentByExternalKeyGatewayTimeout{}
}

/*RefundPaymentByExternalKeyGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type RefundPaymentByExternalKeyGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

func (o *RefundPaymentByExternalKeyGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/refunds][%d] refundPaymentByExternalKeyGatewayTimeout ", 504)
}

func (o *RefundPaymentByExternalKeyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
