// Code generated by go-swagger; DO NOT EDIT.

package tag_definition

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/CDNA-Technologies/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// GetTagDefinitionReader is a Reader for the GetTagDefinition structure.
type GetTagDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTagDefinitionOK()
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

// NewGetTagDefinitionOK creates a GetTagDefinitionOK with default headers values
func NewGetTagDefinitionOK() *GetTagDefinitionOK {
	return &GetTagDefinitionOK{}
}

/*GetTagDefinitionOK handles this case with default header values.

successful operation
*/
type GetTagDefinitionOK struct {
	Payload *kbmodel.TagDefinition

	HttpResponse runtime.ClientResponse
}

func (o *GetTagDefinitionOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tagDefinitions/{tagDefinitionId}][%d] getTagDefinitionOK  %+v", 200, o.Payload)
}

func (o *GetTagDefinitionOK) GetPayload() *kbmodel.TagDefinition {
	return o.Payload
}

func (o *GetTagDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TagDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagDefinitionBadRequest creates a GetTagDefinitionBadRequest with default headers values
func NewGetTagDefinitionBadRequest() *GetTagDefinitionBadRequest {
	return &GetTagDefinitionBadRequest{}
}

/*GetTagDefinitionBadRequest handles this case with default header values.

Invalid tagDefinitionId supplied
*/
type GetTagDefinitionBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetTagDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tagDefinitions/{tagDefinitionId}][%d] getTagDefinitionBadRequest ", 400)
}

func (o *GetTagDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
