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

// GetCatalogTranslationReader is a Reader for the GetCatalogTranslation structure.
type GetCatalogTranslationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCatalogTranslationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCatalogTranslationOK()
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

// NewGetCatalogTranslationOK creates a GetCatalogTranslationOK with default headers values
func NewGetCatalogTranslationOK() *GetCatalogTranslationOK {
	return &GetCatalogTranslationOK{}
}

/*GetCatalogTranslationOK handles this case with default header values.

successful operation
*/
type GetCatalogTranslationOK struct {
	Payload string

	HttpResponse runtime.ClientResponse
}

func (o *GetCatalogTranslationOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/catalogTranslation/{locale}][%d] getCatalogTranslationOK  %+v", 200, o.Payload)
}

func (o *GetCatalogTranslationOK) GetPayload() string {
	return o.Payload
}

func (o *GetCatalogTranslationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCatalogTranslationBadRequest creates a GetCatalogTranslationBadRequest with default headers values
func NewGetCatalogTranslationBadRequest() *GetCatalogTranslationBadRequest {
	return &GetCatalogTranslationBadRequest{}
}

/*GetCatalogTranslationBadRequest handles this case with default header values.

Invalid locale supplied
*/
type GetCatalogTranslationBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetCatalogTranslationBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/catalogTranslation/{locale}][%d] getCatalogTranslationBadRequest ", 400)
}

func (o *GetCatalogTranslationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetCatalogTranslationNotFound creates a GetCatalogTranslationNotFound with default headers values
func NewGetCatalogTranslationNotFound() *GetCatalogTranslationNotFound {
	return &GetCatalogTranslationNotFound{}
}

/*GetCatalogTranslationNotFound handles this case with default header values.

Template not found
*/
type GetCatalogTranslationNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetCatalogTranslationNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoices/catalogTranslation/{locale}][%d] getCatalogTranslationNotFound ", 404)
}

func (o *GetCatalogTranslationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
