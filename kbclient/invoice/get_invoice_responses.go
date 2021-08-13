// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// GetInvoiceReader is a Reader for the GetInvoice structure.
type GetInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceOK()
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

// NewGetInvoiceOK creates a GetInvoiceOK with default headers values
func NewGetInvoiceOK() *GetInvoiceOK {
	return &GetInvoiceOK{}
}

/*GetInvoiceOK handles this case with default header values.

successful operation
*/
type GetInvoiceOK struct {
	Payload *kbmodel.Invoice

	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}][%d] getInvoiceOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceOK) GetPayload() *kbmodel.Invoice {
	return o.Payload
}

func (o *GetInvoiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Invoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceBadRequest creates a GetInvoiceBadRequest with default headers values
func NewGetInvoiceBadRequest() *GetInvoiceBadRequest {
	return &GetInvoiceBadRequest{}
}

/*GetInvoiceBadRequest handles this case with default header values.

Invalid invoice id supplied
*/
type GetInvoiceBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}][%d] getInvoiceBadRequest ", 400)
}

func (o *GetInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInvoiceNotFound creates a GetInvoiceNotFound with default headers values
func NewGetInvoiceNotFound() *GetInvoiceNotFound {
	return &GetInvoiceNotFound{}
}

/*GetInvoiceNotFound handles this case with default header values.

Invoice not found
*/
type GetInvoiceNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}][%d] getInvoiceNotFound ", 404)
}

func (o *GetInvoiceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
