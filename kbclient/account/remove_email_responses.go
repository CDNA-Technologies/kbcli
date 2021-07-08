// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// RemoveEmailReader is a Reader for the RemoveEmail structure.
type RemoveEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRemoveEmailNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRemoveEmailNoContent creates a RemoveEmailNoContent with default headers values
func NewRemoveEmailNoContent() *RemoveEmailNoContent {
	return &RemoveEmailNoContent{}
}

/* RemoveEmailNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type RemoveEmailNoContent struct {
}

func (o *RemoveEmailNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/emails/{email}][%d] removeEmailNoContent ", 204)
}

func (o *RemoveEmailNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveEmailBadRequest creates a RemoveEmailBadRequest with default headers values
func NewRemoveEmailBadRequest() *RemoveEmailBadRequest {
	return &RemoveEmailBadRequest{}
}

/* RemoveEmailBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type RemoveEmailBadRequest struct {
}

func (o *RemoveEmailBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/emails/{email}][%d] removeEmailBadRequest ", 400)
}

func (o *RemoveEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
