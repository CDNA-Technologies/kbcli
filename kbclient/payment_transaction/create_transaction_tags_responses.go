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

// CreateTransactionTagsReader is a Reader for the CreateTransactionTags structure.
type CreateTransactionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransactionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateTransactionTagsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateTransactionTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateTransactionTagsCreated creates a CreateTransactionTagsCreated with default headers values
func NewCreateTransactionTagsCreated() *CreateTransactionTagsCreated {
	return &CreateTransactionTagsCreated{}
}

/* CreateTransactionTagsCreated describes a response with status code 201, with default header values.

Tag created successfully
*/
type CreateTransactionTagsCreated struct {
	Payload []*kbmodel.Tag
}

func (o *CreateTransactionTagsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/tags][%d] createTransactionTagsCreated  %+v", 201, o.Payload)
}
func (o *CreateTransactionTagsCreated) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *CreateTransactionTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionTagsBadRequest creates a CreateTransactionTagsBadRequest with default headers values
func NewCreateTransactionTagsBadRequest() *CreateTransactionTagsBadRequest {
	return &CreateTransactionTagsBadRequest{}
}

/* CreateTransactionTagsBadRequest describes a response with status code 400, with default header values.

Invalid transaction id supplied
*/
type CreateTransactionTagsBadRequest struct {
}

func (o *CreateTransactionTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/tags][%d] createTransactionTagsBadRequest ", 400)
}

func (o *CreateTransactionTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
