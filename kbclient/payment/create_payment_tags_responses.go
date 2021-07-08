// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// CreatePaymentTagsReader is a Reader for the CreatePaymentTags structure.
type CreatePaymentTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePaymentTagsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePaymentTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePaymentTagsCreated creates a CreatePaymentTagsCreated with default headers values
func NewCreatePaymentTagsCreated() *CreatePaymentTagsCreated {
	return &CreatePaymentTagsCreated{}
}

/* CreatePaymentTagsCreated describes a response with status code 201, with default header values.

Tag created successfully
*/
type CreatePaymentTagsCreated struct {
	Payload []*kbmodel.Tag
}

func (o *CreatePaymentTagsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/tags][%d] createPaymentTagsCreated  %+v", 201, o.Payload)
}
func (o *CreatePaymentTagsCreated) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *CreatePaymentTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentTagsBadRequest creates a CreatePaymentTagsBadRequest with default headers values
func NewCreatePaymentTagsBadRequest() *CreatePaymentTagsBadRequest {
	return &CreatePaymentTagsBadRequest{}
}

/* CreatePaymentTagsBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type CreatePaymentTagsBadRequest struct {
}

func (o *CreatePaymentTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/tags][%d] createPaymentTagsBadRequest ", 400)
}

func (o *CreatePaymentTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
