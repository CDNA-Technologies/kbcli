// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CloseAccountReader is a Reader for the CloseAccount structure.
type CloseAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloseAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCloseAccountNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCloseAccountBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloseAccountNoContent creates a CloseAccountNoContent with default headers values
func NewCloseAccountNoContent() *CloseAccountNoContent {
	return &CloseAccountNoContent{}
}

/* CloseAccountNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type CloseAccountNoContent struct {
}

func (o *CloseAccountNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}][%d] closeAccountNoContent ", 204)
}

func (o *CloseAccountNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloseAccountBadRequest creates a CloseAccountBadRequest with default headers values
func NewCloseAccountBadRequest() *CloseAccountBadRequest {
	return &CloseAccountBadRequest{}
}

/* CloseAccountBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type CloseAccountBadRequest struct {
}

func (o *CloseAccountBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}][%d] closeAccountBadRequest ", 400)
}

func (o *CloseAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
