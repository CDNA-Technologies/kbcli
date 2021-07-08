// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// UploadPluginConfigurationReader is a Reader for the UploadPluginConfiguration structure.
type UploadPluginConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadPluginConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUploadPluginConfigurationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadPluginConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUploadPluginConfigurationCreated creates a UploadPluginConfigurationCreated with default headers values
func NewUploadPluginConfigurationCreated() *UploadPluginConfigurationCreated {
	return &UploadPluginConfigurationCreated{}
}

/* UploadPluginConfigurationCreated describes a response with status code 201, with default header values.

Plugin configuration uploaded successfully
*/
type UploadPluginConfigurationCreated struct {
	Payload *kbmodel.TenantKeyValue
}

func (o *UploadPluginConfigurationCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] uploadPluginConfigurationCreated  %+v", 201, o.Payload)
}
func (o *UploadPluginConfigurationCreated) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *UploadPluginConfigurationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadPluginConfigurationBadRequest creates a UploadPluginConfigurationBadRequest with default headers values
func NewUploadPluginConfigurationBadRequest() *UploadPluginConfigurationBadRequest {
	return &UploadPluginConfigurationBadRequest{}
}

/* UploadPluginConfigurationBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type UploadPluginConfigurationBadRequest struct {
}

func (o *UploadPluginConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] uploadPluginConfigurationBadRequest ", 400)
}

func (o *UploadPluginConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
