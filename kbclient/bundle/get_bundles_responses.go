// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// GetBundlesReader is a Reader for the GetBundles structure.
type GetBundlesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBundlesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBundlesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetBundlesOK creates a GetBundlesOK with default headers values
func NewGetBundlesOK() *GetBundlesOK {
	return &GetBundlesOK{}
}

/* GetBundlesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetBundlesOK struct {
	Payload []*kbmodel.Bundle
}

func (o *GetBundlesOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/pagination][%d] getBundlesOK  %+v", 200, o.Payload)
}
func (o *GetBundlesOK) GetPayload() []*kbmodel.Bundle {
	return o.Payload
}

func (o *GetBundlesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
