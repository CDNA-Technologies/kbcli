// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/CDNA-Technologies/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdatePaymentTransactionStateReader is a Reader for the UpdatePaymentTransactionState structure.
type UpdatePaymentTransactionStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdatePaymentTransactionStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdatePaymentTransactionStateNoContent()
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

// NewUpdatePaymentTransactionStateNoContent creates a UpdatePaymentTransactionStateNoContent with default headers values
func NewUpdatePaymentTransactionStateNoContent() *UpdatePaymentTransactionStateNoContent {
	return &UpdatePaymentTransactionStateNoContent{}
}

/*UpdatePaymentTransactionStateNoContent handles this case with default header values.

Successful operation
*/
type UpdatePaymentTransactionStateNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *UpdatePaymentTransactionStateNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/admin/payments/{paymentId}/transactions/{paymentTransactionId}][%d] updatePaymentTransactionStateNoContent ", 204)
}

func (o *UpdatePaymentTransactionStateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdatePaymentTransactionStateBadRequest creates a UpdatePaymentTransactionStateBadRequest with default headers values
func NewUpdatePaymentTransactionStateBadRequest() *UpdatePaymentTransactionStateBadRequest {
	return &UpdatePaymentTransactionStateBadRequest{}
}

/*UpdatePaymentTransactionStateBadRequest handles this case with default header values.

Invalid account data supplied
*/
type UpdatePaymentTransactionStateBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *UpdatePaymentTransactionStateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/admin/payments/{paymentId}/transactions/{paymentTransactionId}][%d] updatePaymentTransactionStateBadRequest ", 400)
}

func (o *UpdatePaymentTransactionStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
