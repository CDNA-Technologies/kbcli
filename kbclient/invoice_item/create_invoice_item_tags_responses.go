// Code generated by go-swagger; DO NOT EDIT.

package invoice_item

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

// CreateInvoiceItemTagsReader is a Reader for the CreateInvoiceItemTags structure.
type CreateInvoiceItemTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateInvoiceItemTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateInvoiceItemTagsCreated()
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

// NewCreateInvoiceItemTagsCreated creates a CreateInvoiceItemTagsCreated with default headers values
func NewCreateInvoiceItemTagsCreated() *CreateInvoiceItemTagsCreated {
	return &CreateInvoiceItemTagsCreated{}
}

/*CreateInvoiceItemTagsCreated handles this case with default header values.

Tag created successfully
*/
type CreateInvoiceItemTagsCreated struct {
	Payload []*kbmodel.Tag

	HttpResponse runtime.ClientResponse
}

func (o *CreateInvoiceItemTagsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoiceItems/{invoiceItemId}/tags][%d] createInvoiceItemTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateInvoiceItemTagsCreated) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *CreateInvoiceItemTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateInvoiceItemTagsBadRequest creates a CreateInvoiceItemTagsBadRequest with default headers values
func NewCreateInvoiceItemTagsBadRequest() *CreateInvoiceItemTagsBadRequest {
	return &CreateInvoiceItemTagsBadRequest{}
}

/*CreateInvoiceItemTagsBadRequest handles this case with default header values.

Invalid invoice item id supplied
*/
type CreateInvoiceItemTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateInvoiceItemTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoiceItems/{invoiceItemId}/tags][%d] createInvoiceItemTagsBadRequest ", 400)
}

func (o *CreateInvoiceItemTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
