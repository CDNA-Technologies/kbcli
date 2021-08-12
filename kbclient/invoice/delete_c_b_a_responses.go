// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/CDNA-Technologies/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCBAReader is a Reader for the DeleteCBA structure.
type DeleteCBAReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCBAReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCBANoContent()
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

// NewDeleteCBANoContent creates a DeleteCBANoContent with default headers values
func NewDeleteCBANoContent() *DeleteCBANoContent {
	return &DeleteCBANoContent{}
}

/*DeleteCBANoContent handles this case with default header values.

Successful operation
*/
type DeleteCBANoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeleteCBANoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoices/{invoiceId}/{invoiceItemId}/cba][%d] deleteCBANoContent ", 204)
}

func (o *DeleteCBANoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCBABadRequest creates a DeleteCBABadRequest with default headers values
func NewDeleteCBABadRequest() *DeleteCBABadRequest {
	return &DeleteCBABadRequest{}
}

/*DeleteCBABadRequest handles this case with default header values.

Invalid account id, invoice id or invoice item id supplied
*/
type DeleteCBABadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeleteCBABadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoices/{invoiceId}/{invoiceItemId}/cba][%d] deleteCBABadRequest ", 400)
}

func (o *DeleteCBABadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCBANotFound creates a DeleteCBANotFound with default headers values
func NewDeleteCBANotFound() *DeleteCBANotFound {
	return &DeleteCBANotFound{}
}

/*DeleteCBANotFound handles this case with default header values.

Account or invoice not found
*/
type DeleteCBANotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeleteCBANotFound) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoices/{invoiceId}/{invoiceItemId}/cba][%d] deleteCBANotFound ", 404)
}

func (o *DeleteCBANotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
