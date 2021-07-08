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

// GetTransactionAuditLogsWithHistoryReader is a Reader for the GetTransactionAuditLogsWithHistory structure.
type GetTransactionAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransactionAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransactionAuditLogsWithHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTransactionAuditLogsWithHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTransactionAuditLogsWithHistoryOK creates a GetTransactionAuditLogsWithHistoryOK with default headers values
func NewGetTransactionAuditLogsWithHistoryOK() *GetTransactionAuditLogsWithHistoryOK {
	return &GetTransactionAuditLogsWithHistoryOK{}
}

/* GetTransactionAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTransactionAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog
}

func (o *GetTransactionAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/auditLogsWithHistory][%d] getTransactionAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}
func (o *GetTransactionAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetTransactionAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransactionAuditLogsWithHistoryNotFound creates a GetTransactionAuditLogsWithHistoryNotFound with default headers values
func NewGetTransactionAuditLogsWithHistoryNotFound() *GetTransactionAuditLogsWithHistoryNotFound {
	return &GetTransactionAuditLogsWithHistoryNotFound{}
}

/* GetTransactionAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetTransactionAuditLogsWithHistoryNotFound struct {
}

func (o *GetTransactionAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentTransactions/{transactionId}/auditLogsWithHistory][%d] getTransactionAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetTransactionAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
