// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// CreateSubscriptionWithAddOnsReader is a Reader for the CreateSubscriptionWithAddOns structure.
type CreateSubscriptionWithAddOnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionWithAddOnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSubscriptionWithAddOnsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSubscriptionWithAddOnsCreated creates a CreateSubscriptionWithAddOnsCreated with default headers values
func NewCreateSubscriptionWithAddOnsCreated() *CreateSubscriptionWithAddOnsCreated {
	return &CreateSubscriptionWithAddOnsCreated{}
}

/* CreateSubscriptionWithAddOnsCreated describes a response with status code 201, with default header values.

Subscriptions created successfully
*/
type CreateSubscriptionWithAddOnsCreated struct {
	Payload *kbmodel.Bundle
}

func (o *CreateSubscriptionWithAddOnsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/createSubscriptionWithAddOns][%d] createSubscriptionWithAddOnsCreated  %+v", 201, o.Payload)
}
func (o *CreateSubscriptionWithAddOnsCreated) GetPayload() *kbmodel.Bundle {
	return o.Payload
}

func (o *CreateSubscriptionWithAddOnsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Bundle)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
