// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

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

// CreateRefundWithAdjustmentsReader is a Reader for the CreateRefundWithAdjustments structure.
type CreateRefundWithAdjustmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRefundWithAdjustmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateRefundWithAdjustmentsCreated()
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

// NewCreateRefundWithAdjustmentsCreated creates a CreateRefundWithAdjustmentsCreated with default headers values
func NewCreateRefundWithAdjustmentsCreated() *CreateRefundWithAdjustmentsCreated {
	return &CreateRefundWithAdjustmentsCreated{}
}

/*CreateRefundWithAdjustmentsCreated handles this case with default header values.

Created refund successfully
*/
type CreateRefundWithAdjustmentsCreated struct {
	Payload *kbmodel.InvoicePayment

	HttpResponse runtime.ClientResponse
}

func (o *CreateRefundWithAdjustmentsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/refunds][%d] createRefundWithAdjustmentsCreated  %+v", 201, o.Payload)
}

func (o *CreateRefundWithAdjustmentsCreated) GetPayload() *kbmodel.InvoicePayment {
	return o.Payload
}

func (o *CreateRefundWithAdjustmentsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.InvoicePayment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRefundWithAdjustmentsBadRequest creates a CreateRefundWithAdjustmentsBadRequest with default headers values
func NewCreateRefundWithAdjustmentsBadRequest() *CreateRefundWithAdjustmentsBadRequest {
	return &CreateRefundWithAdjustmentsBadRequest{}
}

/*CreateRefundWithAdjustmentsBadRequest handles this case with default header values.

Invalid payment id supplied
*/
type CreateRefundWithAdjustmentsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateRefundWithAdjustmentsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/refunds][%d] createRefundWithAdjustmentsBadRequest ", 400)
}

func (o *CreateRefundWithAdjustmentsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateRefundWithAdjustmentsNotFound creates a CreateRefundWithAdjustmentsNotFound with default headers values
func NewCreateRefundWithAdjustmentsNotFound() *CreateRefundWithAdjustmentsNotFound {
	return &CreateRefundWithAdjustmentsNotFound{}
}

/*CreateRefundWithAdjustmentsNotFound handles this case with default header values.

Account or payment not found
*/
type CreateRefundWithAdjustmentsNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateRefundWithAdjustmentsNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/refunds][%d] createRefundWithAdjustmentsNotFound ", 404)
}

func (o *CreateRefundWithAdjustmentsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
