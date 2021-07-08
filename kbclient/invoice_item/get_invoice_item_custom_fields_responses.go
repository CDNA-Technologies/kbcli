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

// GetInvoiceItemCustomFieldsReader is a Reader for the GetInvoiceItemCustomFields structure.
type GetInvoiceItemCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoiceItemCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInvoiceItemCustomFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInvoiceItemCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetInvoiceItemCustomFieldsOK creates a GetInvoiceItemCustomFieldsOK with default headers values
func NewGetInvoiceItemCustomFieldsOK() *GetInvoiceItemCustomFieldsOK {
	return &GetInvoiceItemCustomFieldsOK{}
}

/* GetInvoiceItemCustomFieldsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoiceItemCustomFieldsOK struct {
	Payload []*kbmodel.CustomField
}

func (o *GetInvoiceItemCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] getInvoiceItemCustomFieldsOK  %+v", 200, o.Payload)
}
func (o *GetInvoiceItemCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetInvoiceItemCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoiceItemCustomFieldsBadRequest creates a GetInvoiceItemCustomFieldsBadRequest with default headers values
func NewGetInvoiceItemCustomFieldsBadRequest() *GetInvoiceItemCustomFieldsBadRequest {
	return &GetInvoiceItemCustomFieldsBadRequest{}
}

/* GetInvoiceItemCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid invoice item id supplied
*/
type GetInvoiceItemCustomFieldsBadRequest struct {
}

func (o *GetInvoiceItemCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoiceItems/{invoiceItemId}/customFields][%d] getInvoiceItemCustomFieldsBadRequest ", 400)
}

func (o *GetInvoiceItemCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
