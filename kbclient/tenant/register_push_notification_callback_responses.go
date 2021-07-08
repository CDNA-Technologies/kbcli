// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// RegisterPushNotificationCallbackReader is a Reader for the RegisterPushNotificationCallback structure.
type RegisterPushNotificationCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterPushNotificationCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewRegisterPushNotificationCallbackCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegisterPushNotificationCallbackBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterPushNotificationCallbackCreated creates a RegisterPushNotificationCallbackCreated with default headers values
func NewRegisterPushNotificationCallbackCreated() *RegisterPushNotificationCallbackCreated {
	return &RegisterPushNotificationCallbackCreated{}
}

/* RegisterPushNotificationCallbackCreated describes a response with status code 201, with default header values.

Push notification registered successfully
*/
type RegisterPushNotificationCallbackCreated struct {
	Payload *kbmodel.TenantKeyValue
}

func (o *RegisterPushNotificationCallbackCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/registerNotificationCallback][%d] registerPushNotificationCallbackCreated  %+v", 201, o.Payload)
}
func (o *RegisterPushNotificationCallbackCreated) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *RegisterPushNotificationCallbackCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterPushNotificationCallbackBadRequest creates a RegisterPushNotificationCallbackBadRequest with default headers values
func NewRegisterPushNotificationCallbackBadRequest() *RegisterPushNotificationCallbackBadRequest {
	return &RegisterPushNotificationCallbackBadRequest{}
}

/* RegisterPushNotificationCallbackBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type RegisterPushNotificationCallbackBadRequest struct {
}

func (o *RegisterPushNotificationCallbackBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/registerNotificationCallback][%d] registerPushNotificationCallbackBadRequest ", 400)
}

func (o *RegisterPushNotificationCallbackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
