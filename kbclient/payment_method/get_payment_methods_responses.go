// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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

// GetPaymentMethodsReader is a Reader for the GetPaymentMethods structure.
type GetPaymentMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentMethodsOK()
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

// NewGetPaymentMethodsOK creates a GetPaymentMethodsOK with default headers values
func NewGetPaymentMethodsOK() *GetPaymentMethodsOK {
	return &GetPaymentMethodsOK{}
}

/*GetPaymentMethodsOK handles this case with default header values.

successful operation
*/
type GetPaymentMethodsOK struct {
	Payload []*kbmodel.PaymentMethod

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentMethodsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods/pagination][%d] getPaymentMethodsOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodsOK) GetPayload() []*kbmodel.PaymentMethod {
	return o.Payload
}

func (o *GetPaymentMethodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
