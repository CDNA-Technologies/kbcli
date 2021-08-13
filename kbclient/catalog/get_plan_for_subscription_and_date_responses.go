// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// GetPlanForSubscriptionAndDateReader is a Reader for the GetPlanForSubscriptionAndDate structure.
type GetPlanForSubscriptionAndDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlanForSubscriptionAndDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPlanForSubscriptionAndDateOK()
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

// NewGetPlanForSubscriptionAndDateOK creates a GetPlanForSubscriptionAndDateOK with default headers values
func NewGetPlanForSubscriptionAndDateOK() *GetPlanForSubscriptionAndDateOK {
	return &GetPlanForSubscriptionAndDateOK{}
}

/*GetPlanForSubscriptionAndDateOK handles this case with default header values.

successful operation
*/
type GetPlanForSubscriptionAndDateOK struct {
	Payload *kbmodel.Plan

	HttpResponse runtime.ClientResponse
}

func (o *GetPlanForSubscriptionAndDateOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/plan][%d] getPlanForSubscriptionAndDateOK  %+v", 200, o.Payload)
}

func (o *GetPlanForSubscriptionAndDateOK) GetPayload() *kbmodel.Plan {
	return o.Payload
}

func (o *GetPlanForSubscriptionAndDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Plan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
