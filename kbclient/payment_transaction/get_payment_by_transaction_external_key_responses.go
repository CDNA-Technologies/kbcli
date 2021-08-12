// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// GetPaymentByTransactionExternalKeyReader is a Reader for the GetPaymentByTransactionExternalKey structure.
type GetPaymentByTransactionExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentByTransactionExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentByTransactionExternalKeyOK()
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

// NewGetPaymentByTransactionExternalKeyOK creates a GetPaymentByTransactionExternalKeyOK with default headers values
func NewGetPaymentByTransactionExternalKeyOK() *GetPaymentByTransactionExternalKeyOK {
	return &GetPaymentByTransactionExternalKeyOK{}
}

/*GetPaymentByTransactionExternalKeyOK handles this case with default header values.

successful operation
*/
type GetPaymentByTransactionExternalKeyOK struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentByTransactionExternalKeyOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions][%d] getPaymentByTransactionExternalKeyOK  %+v", 200, o.Payload)
}

func (o *GetPaymentByTransactionExternalKeyOK) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *GetPaymentByTransactionExternalKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentByTransactionExternalKeyNotFound creates a GetPaymentByTransactionExternalKeyNotFound with default headers values
func NewGetPaymentByTransactionExternalKeyNotFound() *GetPaymentByTransactionExternalKeyNotFound {
	return &GetPaymentByTransactionExternalKeyNotFound{}
}

/*GetPaymentByTransactionExternalKeyNotFound handles this case with default header values.

Payment not found
*/
type GetPaymentByTransactionExternalKeyNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentByTransactionExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions][%d] getPaymentByTransactionExternalKeyNotFound ", 404)
}

func (o *GetPaymentByTransactionExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
