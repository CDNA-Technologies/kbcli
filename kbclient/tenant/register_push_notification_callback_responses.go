// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// RegisterPushNotificationCallbackReader is a Reader for the RegisterPushNotificationCallback structure.
type RegisterPushNotificationCallbackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterPushNotificationCallbackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewRegisterPushNotificationCallbackCreated()
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

// NewRegisterPushNotificationCallbackCreated creates a RegisterPushNotificationCallbackCreated with default headers values
func NewRegisterPushNotificationCallbackCreated() *RegisterPushNotificationCallbackCreated {
	return &RegisterPushNotificationCallbackCreated{}
}

/*RegisterPushNotificationCallbackCreated handles this case with default header values.

Push notification registered successfully
*/
type RegisterPushNotificationCallbackCreated struct {
	Payload *kbmodel.TenantKeyValue

	HttpResponse runtime.ClientResponse
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

/*RegisterPushNotificationCallbackBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type RegisterPushNotificationCallbackBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *RegisterPushNotificationCallbackBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/registerNotificationCallback][%d] registerPushNotificationCallbackBadRequest ", 400)
}

func (o *RegisterPushNotificationCallbackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
