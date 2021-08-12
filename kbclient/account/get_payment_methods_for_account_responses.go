// Code generated by go-swagger; DO NOT EDIT.

package account

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

// GetPaymentMethodsForAccountReader is a Reader for the GetPaymentMethodsForAccount structure.
type GetPaymentMethodsForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentMethodsForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentMethodsForAccountOK()
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

// NewGetPaymentMethodsForAccountOK creates a GetPaymentMethodsForAccountOK with default headers values
func NewGetPaymentMethodsForAccountOK() *GetPaymentMethodsForAccountOK {
	return &GetPaymentMethodsForAccountOK{}
}

/*GetPaymentMethodsForAccountOK handles this case with default header values.

successful operation
*/
type GetPaymentMethodsForAccountOK struct {
	Payload []*kbmodel.PaymentMethod

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentMethodsForAccountOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/paymentMethods][%d] getPaymentMethodsForAccountOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodsForAccountOK) GetPayload() []*kbmodel.PaymentMethod {
	return o.Payload
}

func (o *GetPaymentMethodsForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentMethodsForAccountBadRequest creates a GetPaymentMethodsForAccountBadRequest with default headers values
func NewGetPaymentMethodsForAccountBadRequest() *GetPaymentMethodsForAccountBadRequest {
	return &GetPaymentMethodsForAccountBadRequest{}
}

/*GetPaymentMethodsForAccountBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetPaymentMethodsForAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentMethodsForAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/paymentMethods][%d] getPaymentMethodsForAccountBadRequest ", 400)
}

func (o *GetPaymentMethodsForAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPaymentMethodsForAccountNotFound creates a GetPaymentMethodsForAccountNotFound with default headers values
func NewGetPaymentMethodsForAccountNotFound() *GetPaymentMethodsForAccountNotFound {
	return &GetPaymentMethodsForAccountNotFound{}
}

/*GetPaymentMethodsForAccountNotFound handles this case with default header values.

Account not found
*/
type GetPaymentMethodsForAccountNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentMethodsForAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/paymentMethods][%d] getPaymentMethodsForAccountNotFound ", 404)
}

func (o *GetPaymentMethodsForAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
