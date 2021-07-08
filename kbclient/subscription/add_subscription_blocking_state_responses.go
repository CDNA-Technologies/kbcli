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

// AddSubscriptionBlockingStateReader is a Reader for the AddSubscriptionBlockingState structure.
type AddSubscriptionBlockingStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSubscriptionBlockingStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddSubscriptionBlockingStateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddSubscriptionBlockingStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddSubscriptionBlockingStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddSubscriptionBlockingStateCreated creates a AddSubscriptionBlockingStateCreated with default headers values
func NewAddSubscriptionBlockingStateCreated() *AddSubscriptionBlockingStateCreated {
	return &AddSubscriptionBlockingStateCreated{}
}

/* AddSubscriptionBlockingStateCreated describes a response with status code 201, with default header values.

Blocking state created successfully
*/
type AddSubscriptionBlockingStateCreated struct {
	Payload []*kbmodel.BlockingState
}

func (o *AddSubscriptionBlockingStateCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/{subscriptionId}/block][%d] addSubscriptionBlockingStateCreated  %+v", 201, o.Payload)
}
func (o *AddSubscriptionBlockingStateCreated) GetPayload() []*kbmodel.BlockingState {
	return o.Payload
}

func (o *AddSubscriptionBlockingStateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddSubscriptionBlockingStateBadRequest creates a AddSubscriptionBlockingStateBadRequest with default headers values
func NewAddSubscriptionBlockingStateBadRequest() *AddSubscriptionBlockingStateBadRequest {
	return &AddSubscriptionBlockingStateBadRequest{}
}

/* AddSubscriptionBlockingStateBadRequest describes a response with status code 400, with default header values.

Invalid subscription id supplied
*/
type AddSubscriptionBlockingStateBadRequest struct {
}

func (o *AddSubscriptionBlockingStateBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/{subscriptionId}/block][%d] addSubscriptionBlockingStateBadRequest ", 400)
}

func (o *AddSubscriptionBlockingStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddSubscriptionBlockingStateNotFound creates a AddSubscriptionBlockingStateNotFound with default headers values
func NewAddSubscriptionBlockingStateNotFound() *AddSubscriptionBlockingStateNotFound {
	return &AddSubscriptionBlockingStateNotFound{}
}

/* AddSubscriptionBlockingStateNotFound describes a response with status code 404, with default header values.

Subscription not found
*/
type AddSubscriptionBlockingStateNotFound struct {
}

func (o *AddSubscriptionBlockingStateNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/subscriptions/{subscriptionId}/block][%d] addSubscriptionBlockingStateNotFound ", 404)
}

func (o *AddSubscriptionBlockingStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
