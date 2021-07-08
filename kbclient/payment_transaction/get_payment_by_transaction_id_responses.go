// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetPaymentByTransactionIDReader is a Reader for the GetPaymentByTransactionID structure.
type GetPaymentByTransactionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentByTransactionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPaymentByTransactionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetPaymentByTransactionIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPaymentByTransactionIDOK creates a GetPaymentByTransactionIDOK with default headers values
func NewGetPaymentByTransactionIDOK() *GetPaymentByTransactionIDOK {
	return &GetPaymentByTransactionIDOK{}
}

/* GetPaymentByTransactionIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPaymentByTransactionIDOK struct {
	Payload *kbmodel.Payment
}

func (o *GetPaymentByTransactionIDOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}][%d] getPaymentByTransactionIdOK  %+v", 200, o.Payload)
}
func (o *GetPaymentByTransactionIDOK) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *GetPaymentByTransactionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentByTransactionIDNotFound creates a GetPaymentByTransactionIDNotFound with default headers values
func NewGetPaymentByTransactionIDNotFound() *GetPaymentByTransactionIDNotFound {
	return &GetPaymentByTransactionIDNotFound{}
}

/* GetPaymentByTransactionIDNotFound describes a response with status code 404, with default header values.

Payment not found
*/
type GetPaymentByTransactionIDNotFound struct {
}

func (o *GetPaymentByTransactionIDNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}][%d] getPaymentByTransactionIdNotFound ", 404)
}

func (o *GetPaymentByTransactionIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
