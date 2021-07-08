// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetSubscriptionByKeyReader is a Reader for the GetSubscriptionByKey structure.
type GetSubscriptionByKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionByKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSubscriptionByKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetSubscriptionByKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSubscriptionByKeyOK creates a GetSubscriptionByKeyOK with default headers values
func NewGetSubscriptionByKeyOK() *GetSubscriptionByKeyOK {
	return &GetSubscriptionByKeyOK{}
}

/* GetSubscriptionByKeyOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSubscriptionByKeyOK struct {
	Payload *kbmodel.Subscription
}

func (o *GetSubscriptionByKeyOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions][%d] getSubscriptionByKeyOK  %+v", 200, o.Payload)
}
func (o *GetSubscriptionByKeyOK) GetPayload() *kbmodel.Subscription {
	return o.Payload
}

func (o *GetSubscriptionByKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionByKeyNotFound creates a GetSubscriptionByKeyNotFound with default headers values
func NewGetSubscriptionByKeyNotFound() *GetSubscriptionByKeyNotFound {
	return &GetSubscriptionByKeyNotFound{}
}

/* GetSubscriptionByKeyNotFound describes a response with status code 404, with default header values.

Subscription not found
*/
type GetSubscriptionByKeyNotFound struct {
}

func (o *GetSubscriptionByKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions][%d] getSubscriptionByKeyNotFound ", 404)
}

func (o *GetSubscriptionByKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
