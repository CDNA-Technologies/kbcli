// Code generated by go-swagger; DO NOT EDIT.

package payment_method

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// CreatePaymentMethodCustomFieldsReader is a Reader for the CreatePaymentMethodCustomFields structure.
type CreatePaymentMethodCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentMethodCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreatePaymentMethodCustomFieldsCreated()
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

// NewCreatePaymentMethodCustomFieldsCreated creates a CreatePaymentMethodCustomFieldsCreated with default headers values
func NewCreatePaymentMethodCustomFieldsCreated() *CreatePaymentMethodCustomFieldsCreated {
	return &CreatePaymentMethodCustomFieldsCreated{}
}

/*CreatePaymentMethodCustomFieldsCreated handles this case with default header values.

Custom field created successfully
*/
type CreatePaymentMethodCustomFieldsCreated struct {
	Payload []*kbmodel.CustomField

	HttpResponse runtime.ClientResponse
}

func (o *CreatePaymentMethodCustomFieldsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] createPaymentMethodCustomFieldsCreated  %+v", 201, o.Payload)
}

func (o *CreatePaymentMethodCustomFieldsCreated) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *CreatePaymentMethodCustomFieldsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentMethodCustomFieldsBadRequest creates a CreatePaymentMethodCustomFieldsBadRequest with default headers values
func NewCreatePaymentMethodCustomFieldsBadRequest() *CreatePaymentMethodCustomFieldsBadRequest {
	return &CreatePaymentMethodCustomFieldsBadRequest{}
}

/*CreatePaymentMethodCustomFieldsBadRequest handles this case with default header values.

Invalid payment method id supplied
*/
type CreatePaymentMethodCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreatePaymentMethodCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentMethods/{paymentMethodId}/customFields][%d] createPaymentMethodCustomFieldsBadRequest ", 400)
}

func (o *CreatePaymentMethodCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
