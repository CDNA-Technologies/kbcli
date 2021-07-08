// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// GetInvoiceAuditLogsWithHistoryReader is a Reader for the GetInvoiceAuditLogsWithHistory structure.
type GetInvoiceAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvoiceAuditLogsWithHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetInvoiceAuditLogsWithHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvoiceAuditLogsWithHistoryOK creates a GetInvoiceAuditLogsWithHistoryOK with default headers values
func NewGetInvoiceAuditLogsWithHistoryOK() *GetInvoiceAuditLogsWithHistoryOK {
	return &GetInvoiceAuditLogsWithHistoryOK{}
}

/* GetInvoiceAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoiceAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog
}

func (o *GetInvoiceAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/auditLogsWithHistory][%d] getInvoiceAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}
func (o *GetInvoiceAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetInvoiceAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceAuditLogsWithHistoryNotFound creates a GetInvoiceAuditLogsWithHistoryNotFound with default headers values
func NewGetInvoiceAuditLogsWithHistoryNotFound() *GetInvoiceAuditLogsWithHistoryNotFound {
	return &GetInvoiceAuditLogsWithHistoryNotFound{}
}

/* GetInvoiceAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Invoice not found
*/
type GetInvoiceAuditLogsWithHistoryNotFound struct {
}

func (o *GetInvoiceAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/auditLogsWithHistory][%d] getInvoiceAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetInvoiceAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
