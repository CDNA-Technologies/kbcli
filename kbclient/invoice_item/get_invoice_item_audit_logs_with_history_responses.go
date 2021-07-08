// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetInvoiceItemAuditLogsWithHistoryReader is a Reader for the GetInvoiceItemAuditLogsWithHistory structure.
type GetInvoiceItemAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceItemAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvoiceItemAuditLogsWithHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetInvoiceItemAuditLogsWithHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvoiceItemAuditLogsWithHistoryOK creates a GetInvoiceItemAuditLogsWithHistoryOK with default headers values
func NewGetInvoiceItemAuditLogsWithHistoryOK() *GetInvoiceItemAuditLogsWithHistoryOK {
	return &GetInvoiceItemAuditLogsWithHistoryOK{}
}

/* GetInvoiceItemAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoiceItemAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog
}

func (o *GetInvoiceItemAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/auditLogsWithHistory][%d] getInvoiceItemAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}
func (o *GetInvoiceItemAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetInvoiceItemAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceItemAuditLogsWithHistoryNotFound creates a GetInvoiceItemAuditLogsWithHistoryNotFound with default headers values
func NewGetInvoiceItemAuditLogsWithHistoryNotFound() *GetInvoiceItemAuditLogsWithHistoryNotFound {
	return &GetInvoiceItemAuditLogsWithHistoryNotFound{}
}

/* GetInvoiceItemAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Invoice item not found
*/
type GetInvoiceItemAuditLogsWithHistoryNotFound struct {
}

func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/auditLogsWithHistory][%d] getInvoiceItemAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
