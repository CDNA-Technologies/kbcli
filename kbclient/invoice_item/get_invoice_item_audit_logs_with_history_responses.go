// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
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

// NewGetInvoiceItemAuditLogsWithHistoryOK creates a GetInvoiceItemAuditLogsWithHistoryOK with default headers values
func NewGetInvoiceItemAuditLogsWithHistoryOK() *GetInvoiceItemAuditLogsWithHistoryOK {
	return &GetInvoiceItemAuditLogsWithHistoryOK{}
}

/*GetInvoiceItemAuditLogsWithHistoryOK handles this case with default header values.

successful operation
*/
type GetInvoiceItemAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog

	HttpResponse runtime.ClientResponse
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

/*GetInvoiceItemAuditLogsWithHistoryNotFound handles this case with default header values.

Invoice item not found
*/
type GetInvoiceItemAuditLogsWithHistoryNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/auditLogsWithHistory][%d] getInvoiceItemAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetInvoiceItemAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
