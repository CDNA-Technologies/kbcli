// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeletePaymentCustomFieldsReader is a Reader for the DeletePaymentCustomFields structure.
type DeletePaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePaymentCustomFieldsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeletePaymentCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePaymentCustomFieldsNoContent creates a DeletePaymentCustomFieldsNoContent with default headers values
func NewDeletePaymentCustomFieldsNoContent() *DeletePaymentCustomFieldsNoContent {
	return &DeletePaymentCustomFieldsNoContent{}
}

/* DeletePaymentCustomFieldsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeletePaymentCustomFieldsNoContent struct {
}

func (o *DeletePaymentCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/payments/{paymentId}/customFields][%d] deletePaymentCustomFieldsNoContent ", 204)
}

func (o *DeletePaymentCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePaymentCustomFieldsBadRequest creates a DeletePaymentCustomFieldsBadRequest with default headers values
func NewDeletePaymentCustomFieldsBadRequest() *DeletePaymentCustomFieldsBadRequest {
	return &DeletePaymentCustomFieldsBadRequest{}
}

/* DeletePaymentCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type DeletePaymentCustomFieldsBadRequest struct {
}

func (o *DeletePaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/payments/{paymentId}/customFields][%d] deletePaymentCustomFieldsBadRequest ", 400)
}

func (o *DeletePaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
