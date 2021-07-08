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

// CreateSubscriptionsWithAddOnsReader is a Reader for the CreateSubscriptionsWithAddOns structure.
type CreateSubscriptionsWithAddOnsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateSubscriptionsWithAddOnsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateSubscriptionsWithAddOnsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateSubscriptionsWithAddOnsCreated creates a CreateSubscriptionsWithAddOnsCreated with default headers values
func NewCreateSubscriptionsWithAddOnsCreated() *CreateSubscriptionsWithAddOnsCreated {
	return &CreateSubscriptionsWithAddOnsCreated{}
}

/* CreateSubscriptionsWithAddOnsCreated describes a response with status code 201, with default header values.

Subscriptions created successfully
*/
type CreateSubscriptionsWithAddOnsCreated struct {
	Payload []*kbmodel.Bundle
}

func (o *CreateSubscriptionsWithAddOnsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/createSubscriptionsWithAddOns][%d] createSubscriptionsWithAddOnsCreated  %+v", 201, o.Payload)
}
func (o *CreateSubscriptionsWithAddOnsCreated) GetPayload() []*kbmodel.Bundle {
	return o.Payload
}

func (o *CreateSubscriptionsWithAddOnsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
