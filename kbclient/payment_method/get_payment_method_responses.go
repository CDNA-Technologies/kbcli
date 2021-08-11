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

// GetPaymentMethodReader is a Reader for the GetPaymentMethod structure.
type GetPaymentMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentMethodOK()
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

// NewGetPaymentMethodOK creates a GetPaymentMethodOK with default headers values
func NewGetPaymentMethodOK() *GetPaymentMethodOK {
	return &GetPaymentMethodOK{}
}

/*GetPaymentMethodOK handles this case with default header values.

successful operation
*/
type GetPaymentMethodOK struct {
	Payload *kbmodel.PaymentMethod

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentMethodOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}][%d] getPaymentMethodOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodOK) GetPayload() *kbmodel.PaymentMethod {
	return o.Payload
}

func (o *GetPaymentMethodOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.PaymentMethod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentMethodBadRequest creates a GetPaymentMethodBadRequest with default headers values
func NewGetPaymentMethodBadRequest() *GetPaymentMethodBadRequest {
	return &GetPaymentMethodBadRequest{}
}

/*GetPaymentMethodBadRequest handles this case with default header values.

Invalid paymentMethodId supplied
*/
type GetPaymentMethodBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentMethodBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}][%d] getPaymentMethodBadRequest ", 400)
}

func (o *GetPaymentMethodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPaymentMethodNotFound creates a GetPaymentMethodNotFound with default headers values
func NewGetPaymentMethodNotFound() *GetPaymentMethodNotFound {
	return &GetPaymentMethodNotFound{}
}

/*GetPaymentMethodNotFound handles this case with default header values.

Account or payment method not found
*/
type GetPaymentMethodNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentMethodNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/{paymentMethodId}][%d] getPaymentMethodNotFound ", 404)
}

func (o *GetPaymentMethodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
