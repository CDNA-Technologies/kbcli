// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// ModifyBundleCustomFieldsReader is a Reader for the ModifyBundleCustomFields structure.
type ModifyBundleCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyBundleCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewModifyBundleCustomFieldsNoContent()
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

// NewModifyBundleCustomFieldsNoContent creates a ModifyBundleCustomFieldsNoContent with default headers values
func NewModifyBundleCustomFieldsNoContent() *ModifyBundleCustomFieldsNoContent {
	return &ModifyBundleCustomFieldsNoContent{}
}

/*ModifyBundleCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type ModifyBundleCustomFieldsNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *ModifyBundleCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/customFields][%d] modifyBundleCustomFieldsNoContent ", 204)
}

func (o *ModifyBundleCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyBundleCustomFieldsBadRequest creates a ModifyBundleCustomFieldsBadRequest with default headers values
func NewModifyBundleCustomFieldsBadRequest() *ModifyBundleCustomFieldsBadRequest {
	return &ModifyBundleCustomFieldsBadRequest{}
}

/*ModifyBundleCustomFieldsBadRequest handles this case with default header values.

Invalid bundle id supplied
*/
type ModifyBundleCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ModifyBundleCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/customFields][%d] modifyBundleCustomFieldsBadRequest ", 400)
}

func (o *ModifyBundleCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
