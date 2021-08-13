// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/CDNA-Technologies/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// TransferChildCreditToParentReader is a Reader for the TransferChildCreditToParent structure.
type TransferChildCreditToParentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransferChildCreditToParentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewTransferChildCreditToParentNoContent()
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

// NewTransferChildCreditToParentNoContent creates a TransferChildCreditToParentNoContent with default headers values
func NewTransferChildCreditToParentNoContent() *TransferChildCreditToParentNoContent {
	return &TransferChildCreditToParentNoContent{}
}

/*TransferChildCreditToParentNoContent handles this case with default header values.

Successful operation
*/
type TransferChildCreditToParentNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *TransferChildCreditToParentNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{childAccountId}/transferCredit][%d] transferChildCreditToParentNoContent ", 204)
}

func (o *TransferChildCreditToParentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferChildCreditToParentBadRequest creates a TransferChildCreditToParentBadRequest with default headers values
func NewTransferChildCreditToParentBadRequest() *TransferChildCreditToParentBadRequest {
	return &TransferChildCreditToParentBadRequest{}
}

/*TransferChildCreditToParentBadRequest handles this case with default header values.

Account does not have credit
*/
type TransferChildCreditToParentBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *TransferChildCreditToParentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{childAccountId}/transferCredit][%d] transferChildCreditToParentBadRequest ", 400)
}

func (o *TransferChildCreditToParentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTransferChildCreditToParentNotFound creates a TransferChildCreditToParentNotFound with default headers values
func NewTransferChildCreditToParentNotFound() *TransferChildCreditToParentNotFound {
	return &TransferChildCreditToParentNotFound{}
}

/*TransferChildCreditToParentNotFound handles this case with default header values.

Account not found
*/
type TransferChildCreditToParentNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *TransferChildCreditToParentNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{childAccountId}/transferCredit][%d] transferChildCreditToParentNotFound ", 404)
}

func (o *TransferChildCreditToParentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
