// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// GetPhaseForSubscriptionAndDateReader is a Reader for the GetPhaseForSubscriptionAndDate structure.
type GetPhaseForSubscriptionAndDateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPhaseForSubscriptionAndDateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPhaseForSubscriptionAndDateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPhaseForSubscriptionAndDateOK creates a GetPhaseForSubscriptionAndDateOK with default headers values
func NewGetPhaseForSubscriptionAndDateOK() *GetPhaseForSubscriptionAndDateOK {
	return &GetPhaseForSubscriptionAndDateOK{}
}

/* GetPhaseForSubscriptionAndDateOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPhaseForSubscriptionAndDateOK struct {
	Payload *kbmodel.Phase
}

func (o *GetPhaseForSubscriptionAndDateOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/catalog/phase][%d] getPhaseForSubscriptionAndDateOK  %+v", 200, o.Payload)
}
func (o *GetPhaseForSubscriptionAndDateOK) GetPayload() *kbmodel.Phase {
	return o.Payload
}

func (o *GetPhaseForSubscriptionAndDateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Phase)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
