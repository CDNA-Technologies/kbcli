// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// GetInvoiceTagsReader is a Reader for the GetInvoiceTags structure.
type GetInvoiceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoiceTagsOK()
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

// NewGetInvoiceTagsOK creates a GetInvoiceTagsOK with default headers values
func NewGetInvoiceTagsOK() *GetInvoiceTagsOK {
	return &GetInvoiceTagsOK{}
}

/*GetInvoiceTagsOK handles this case with default header values.

successful operation
*/
type GetInvoiceTagsOK struct {
	Payload []*kbmodel.Tag

	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceTagsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/tags][%d] getInvoiceTagsOK  %+v", 200, o.Payload)
}

func (o *GetInvoiceTagsOK) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *GetInvoiceTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceTagsBadRequest creates a GetInvoiceTagsBadRequest with default headers values
func NewGetInvoiceTagsBadRequest() *GetInvoiceTagsBadRequest {
	return &GetInvoiceTagsBadRequest{}
}

/*GetInvoiceTagsBadRequest handles this case with default header values.

Invalid invoice id supplied
*/
type GetInvoiceTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/tags][%d] getInvoiceTagsBadRequest ", 400)
}

func (o *GetInvoiceTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetInvoiceTagsNotFound creates a GetInvoiceTagsNotFound with default headers values
func NewGetInvoiceTagsNotFound() *GetInvoiceTagsNotFound {
	return &GetInvoiceTagsNotFound{}
}

/*GetInvoiceTagsNotFound handles this case with default header values.

Invoice not found
*/
type GetInvoiceTagsNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetInvoiceTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/{invoiceId}/tags][%d] getInvoiceTagsNotFound ", 404)
}

func (o *GetInvoiceTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
