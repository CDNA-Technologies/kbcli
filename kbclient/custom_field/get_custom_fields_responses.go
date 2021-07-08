// Code generated by go-swagger; DO NOT EDIT.

package custom_field

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// GetCustomFieldsReader is a Reader for the GetCustomFields structure.
type GetCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCustomFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCustomFieldsOK creates a GetCustomFieldsOK with default headers values
func NewGetCustomFieldsOK() *GetCustomFieldsOK {
	return &GetCustomFieldsOK{}
}

/* GetCustomFieldsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetCustomFieldsOK struct {
	Payload []*kbmodel.CustomField
}

func (o *GetCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/customFields/pagination][%d] getCustomFieldsOK  %+v", 200, o.Payload)
}
func (o *GetCustomFieldsOK) GetPayload() []*kbmodel.CustomField {
	return o.Payload
}

func (o *GetCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
