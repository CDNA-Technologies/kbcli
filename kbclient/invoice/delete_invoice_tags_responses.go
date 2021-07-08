// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteInvoiceTagsReader is a Reader for the DeleteInvoiceTags structure.
type DeleteInvoiceTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInvoiceTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteInvoiceTagsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteInvoiceTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteInvoiceTagsNoContent creates a DeleteInvoiceTagsNoContent with default headers values
func NewDeleteInvoiceTagsNoContent() *DeleteInvoiceTagsNoContent {
	return &DeleteInvoiceTagsNoContent{}
}

/* DeleteInvoiceTagsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteInvoiceTagsNoContent struct {
}

func (o *DeleteInvoiceTagsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoices/{invoiceId}/tags][%d] deleteInvoiceTagsNoContent ", 204)
}

func (o *DeleteInvoiceTagsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInvoiceTagsBadRequest creates a DeleteInvoiceTagsBadRequest with default headers values
func NewDeleteInvoiceTagsBadRequest() *DeleteInvoiceTagsBadRequest {
	return &DeleteInvoiceTagsBadRequest{}
}

/* DeleteInvoiceTagsBadRequest describes a response with status code 400, with default header values.

Invalid invoice id supplied
*/
type DeleteInvoiceTagsBadRequest struct {
}

func (o *DeleteInvoiceTagsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoices/{invoiceId}/tags][%d] deleteInvoiceTagsBadRequest ", 400)
}

func (o *DeleteInvoiceTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
