// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

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

// CreateTransactionCustomFieldsReader is a Reader for the CreateTransactionCustomFields structure.
type CreateTransactionCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransactionCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateTransactionCustomFieldsCreated()
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

// NewCreateTransactionCustomFieldsCreated creates a CreateTransactionCustomFieldsCreated with default headers values
func NewCreateTransactionCustomFieldsCreated() *CreateTransactionCustomFieldsCreated {
	return &CreateTransactionCustomFieldsCreated{}
}

/*CreateTransactionCustomFieldsCreated handles this case with default header values.

Custom field created successfully
*/
type CreateTransactionCustomFieldsCreated struct {
	Payload []*kbmodel.CustomField

	HttpResponse runtime.ClientResponse
}

func (o *CreateTransactionCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/customFields][%d] createTransactionCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreateTransactionCustomFieldsCreated) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *CreateTransactionCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionCustomFieldsBadRequest creates a CreateTransactionCustomFieldsBadRequest with default headers values
func NewCreateTransactionCustomFieldsBadRequest() *CreateTransactionCustomFieldsBadRequest {
	return &CreateTransactionCustomFieldsBadRequest{}
}

/*CreateTransactionCustomFieldsBadRequest handles this case with default header values.

Invalid transaction id supplied
*/
type CreateTransactionCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateTransactionCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/customFields][%d] createTransactionCustomFieldsBadRequest ", 400)
}

func (o *CreateTransactionCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
