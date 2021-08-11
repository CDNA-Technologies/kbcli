// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v2/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// GetAccountTagsReader is a Reader for the GetAccountTags structure.
type GetAccountTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountTagsOK()
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

// NewGetAccountTagsOK creates a GetAccountTagsOK with default headers values
func NewGetAccountTagsOK() *GetAccountTagsOK {
	return &GetAccountTagsOK{}
}

/*GetAccountTagsOK handles this case with default header values.

successful operation
*/
type GetAccountTagsOK struct {
	Payload []*kbmodel.Tag

	HttpResponse runtime.ClientResponse
}

func (o *GetAccountTagsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/tags][%d] getAccountTagsOK  %+v", 200, o.Payload)
}

func (o *GetAccountTagsOK) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *GetAccountTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountTagsBadRequest creates a GetAccountTagsBadRequest with default headers values
func NewGetAccountTagsBadRequest() *GetAccountTagsBadRequest {
	return &GetAccountTagsBadRequest{}
}

/*GetAccountTagsBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetAccountTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAccountTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/tags][%d] getAccountTagsBadRequest ", 400)
}

func (o *GetAccountTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountTagsNotFound creates a GetAccountTagsNotFound with default headers values
func NewGetAccountTagsNotFound() *GetAccountTagsNotFound {
	return &GetAccountTagsNotFound{}
}

/*GetAccountTagsNotFound handles this case with default header values.

Account not found
*/
type GetAccountTagsNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAccountTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/tags][%d] getAccountTagsNotFound ", 404)
}

func (o *GetAccountTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
