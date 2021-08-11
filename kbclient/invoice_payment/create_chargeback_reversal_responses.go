// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// CreateChargebackReversalReader is a Reader for the CreateChargebackReversal structure.
type CreateChargebackReversalReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateChargebackReversalReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateChargebackReversalCreated()
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

// NewCreateChargebackReversalCreated creates a CreateChargebackReversalCreated with default headers values
func NewCreateChargebackReversalCreated() *CreateChargebackReversalCreated {
	return &CreateChargebackReversalCreated{}
}

/*CreateChargebackReversalCreated handles this case with default header values.

Created chargeback reversal successfully
*/
type CreateChargebackReversalCreated struct {
	Payload *kbmodel.InvoicePayment

	HttpResponse runtime.ClientResponse
}

func (o *CreateChargebackReversalCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebackReversals][%d] createChargebackReversalCreated  %+v", 201, o.Payload)
}

func (o *CreateChargebackReversalCreated) GetPayload() *kbmodel.InvoicePayment {
	return o.Payload
}

func (o *CreateChargebackReversalCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.InvoicePayment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateChargebackReversalBadRequest creates a CreateChargebackReversalBadRequest with default headers values
func NewCreateChargebackReversalBadRequest() *CreateChargebackReversalBadRequest {
	return &CreateChargebackReversalBadRequest{}
}

/*CreateChargebackReversalBadRequest handles this case with default header values.

Invalid payment id supplied
*/
type CreateChargebackReversalBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateChargebackReversalBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebackReversals][%d] createChargebackReversalBadRequest ", 400)
}

func (o *CreateChargebackReversalBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateChargebackReversalNotFound creates a CreateChargebackReversalNotFound with default headers values
func NewCreateChargebackReversalNotFound() *CreateChargebackReversalNotFound {
	return &CreateChargebackReversalNotFound{}
}

/*CreateChargebackReversalNotFound handles this case with default header values.

Account or payment not found
*/
type CreateChargebackReversalNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateChargebackReversalNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebackReversals][%d] createChargebackReversalNotFound ", 404)
}

func (o *CreateChargebackReversalNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
