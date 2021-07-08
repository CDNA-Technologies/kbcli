// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// UpdateAccountReader is a Reader for the UpdateAccount structure.
type UpdateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateAccountNoContent creates a UpdateAccountNoContent with default headers values
func NewUpdateAccountNoContent() *UpdateAccountNoContent {
	return &UpdateAccountNoContent{}
}

/* UpdateAccountNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type UpdateAccountNoContent struct {
}

func (o *UpdateAccountNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}][%d] updateAccountNoContent ", 204)
}

func (o *UpdateAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateAccountBadRequest creates a UpdateAccountBadRequest with default headers values
func NewUpdateAccountBadRequest() *UpdateAccountBadRequest {
	return &UpdateAccountBadRequest{}
}

/* UpdateAccountBadRequest describes a response with status code 400, with default header values.

Invalid account data supplied
*/
type UpdateAccountBadRequest struct {
}

func (o *UpdateAccountBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}][%d] updateAccountBadRequest ", 400)
}

func (o *UpdateAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
