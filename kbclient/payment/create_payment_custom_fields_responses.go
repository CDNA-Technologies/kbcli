// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// CreatePaymentCustomFieldsReader is a Reader for the CreatePaymentCustomFields structure.
type CreatePaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePaymentCustomFieldsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePaymentCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePaymentCustomFieldsCreated creates a CreatePaymentCustomFieldsCreated with default headers values
func NewCreatePaymentCustomFieldsCreated() *CreatePaymentCustomFieldsCreated {
	return &CreatePaymentCustomFieldsCreated{}
}

/* CreatePaymentCustomFieldsCreated describes a response with status code 201, with default header values.

Custom field created successfully
*/
type CreatePaymentCustomFieldsCreated struct {
	Payload []*kbmodel.CustomField
}

func (o *CreatePaymentCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/customFields][%d] createPaymentCustomFieldsCreated  %+v", 201, o.Payload)
}
func (o *CreatePaymentCustomFieldsCreated) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *CreatePaymentCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentCustomFieldsBadRequest creates a CreatePaymentCustomFieldsBadRequest with default headers values
func NewCreatePaymentCustomFieldsBadRequest() *CreatePaymentCustomFieldsBadRequest {
	return &CreatePaymentCustomFieldsBadRequest{}
}

/* CreatePaymentCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type CreatePaymentCustomFieldsBadRequest struct {
}

func (o *CreatePaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments/{paymentId}/customFields][%d] createPaymentCustomFieldsBadRequest ", 400)
}

func (o *CreatePaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
