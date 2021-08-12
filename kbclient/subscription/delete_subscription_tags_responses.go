// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/CDNA-Technologies/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteSubscriptionTagsReader is a Reader for the DeleteSubscriptionTags structure.
type DeleteSubscriptionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSubscriptionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteSubscriptionTagsNoContent()
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

// NewDeleteSubscriptionTagsNoContent creates a DeleteSubscriptionTagsNoContent with default headers values
func NewDeleteSubscriptionTagsNoContent() *DeleteSubscriptionTagsNoContent {
	return &DeleteSubscriptionTagsNoContent{}
}

/*DeleteSubscriptionTagsNoContent handles this case with default header values.

Successful operation
*/
type DeleteSubscriptionTagsNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeleteSubscriptionTagsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/subscriptions/{subscriptionId}/tags][%d] deleteSubscriptionTagsNoContent ", 204)
}

func (o *DeleteSubscriptionTagsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteSubscriptionTagsBadRequest creates a DeleteSubscriptionTagsBadRequest with default headers values
func NewDeleteSubscriptionTagsBadRequest() *DeleteSubscriptionTagsBadRequest {
	return &DeleteSubscriptionTagsBadRequest{}
}

/*DeleteSubscriptionTagsBadRequest handles this case with default header values.

Invalid subscription id supplied
*/
type DeleteSubscriptionTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeleteSubscriptionTagsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/subscriptions/{subscriptionId}/tags][%d] deleteSubscriptionTagsBadRequest ", 400)
}

func (o *DeleteSubscriptionTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
