// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// CreateBundleTagsReader is a Reader for the CreateBundleTags structure.
type CreateBundleTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBundleTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateBundleTagsCreated()
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

// NewCreateBundleTagsCreated creates a CreateBundleTagsCreated with default headers values
func NewCreateBundleTagsCreated() *CreateBundleTagsCreated {
	return &CreateBundleTagsCreated{}
}

/*CreateBundleTagsCreated handles this case with default header values.

Tag created successfully
*/
type CreateBundleTagsCreated struct {
	Payload []*kbmodel.Tag

	HttpResponse runtime.ClientResponse
}

func (o *CreateBundleTagsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/tags][%d] createBundleTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateBundleTagsCreated) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *CreateBundleTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleTagsBadRequest creates a CreateBundleTagsBadRequest with default headers values
func NewCreateBundleTagsBadRequest() *CreateBundleTagsBadRequest {
	return &CreateBundleTagsBadRequest{}
}

/*CreateBundleTagsBadRequest handles this case with default header values.

Invalid bundle id supplied
*/
type CreateBundleTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateBundleTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/tags][%d] createBundleTagsBadRequest ", 400)
}

func (o *CreateBundleTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
