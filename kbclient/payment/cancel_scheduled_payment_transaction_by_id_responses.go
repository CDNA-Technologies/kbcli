// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// CancelScheduledPaymentTransactionByIDReader is a Reader for the CancelScheduledPaymentTransactionByID structure.
type CancelScheduledPaymentTransactionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelScheduledPaymentTransactionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCancelScheduledPaymentTransactionByIDNoContent()
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

// NewCancelScheduledPaymentTransactionByIDNoContent creates a CancelScheduledPaymentTransactionByIDNoContent with default headers values
func NewCancelScheduledPaymentTransactionByIDNoContent() *CancelScheduledPaymentTransactionByIDNoContent {
	return &CancelScheduledPaymentTransactionByIDNoContent{}
}

/*CancelScheduledPaymentTransactionByIDNoContent handles this case with default header values.

Successful operation
*/
type CancelScheduledPaymentTransactionByIDNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *CancelScheduledPaymentTransactionByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/payments/{paymentTransactionId}/cancelScheduledPaymentTransaction][%d] cancelScheduledPaymentTransactionByIdNoContent ", 204)
}

func (o *CancelScheduledPaymentTransactionByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelScheduledPaymentTransactionByIDBadRequest creates a CancelScheduledPaymentTransactionByIDBadRequest with default headers values
func NewCancelScheduledPaymentTransactionByIDBadRequest() *CancelScheduledPaymentTransactionByIDBadRequest {
	return &CancelScheduledPaymentTransactionByIDBadRequest{}
}

/*CancelScheduledPaymentTransactionByIDBadRequest handles this case with default header values.

Invalid paymentTransactionId supplied
*/
type CancelScheduledPaymentTransactionByIDBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CancelScheduledPaymentTransactionByIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/payments/{paymentTransactionId}/cancelScheduledPaymentTransaction][%d] cancelScheduledPaymentTransactionByIdBadRequest ", 400)
}

func (o *CancelScheduledPaymentTransactionByIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
