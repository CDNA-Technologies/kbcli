// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// InsertUserKeyValueReader is a Reader for the InsertUserKeyValue structure.
type InsertUserKeyValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InsertUserKeyValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewInsertUserKeyValueCreated()
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

// NewInsertUserKeyValueCreated creates a InsertUserKeyValueCreated with default headers values
func NewInsertUserKeyValueCreated() *InsertUserKeyValueCreated {
	return &InsertUserKeyValueCreated{}
}

/*InsertUserKeyValueCreated handles this case with default header values.

Per tenant config uploaded successfully
*/
type InsertUserKeyValueCreated struct {
	Payload *kbmodel.TenantKeyValue

	HttpResponse runtime.ClientResponse
}

func (o *InsertUserKeyValueCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/userKeyValue/{keyName}][%d] insertUserKeyValueCreated  %+v", 201, o.Payload)
}

func (o *InsertUserKeyValueCreated) GetPayload() *kbmodel.TenantKeyValue {
	return o.Payload
}

func (o *InsertUserKeyValueCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInsertUserKeyValueBadRequest creates a InsertUserKeyValueBadRequest with default headers values
func NewInsertUserKeyValueBadRequest() *InsertUserKeyValueBadRequest {
	return &InsertUserKeyValueBadRequest{}
}

/*InsertUserKeyValueBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type InsertUserKeyValueBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *InsertUserKeyValueBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/tenants/userKeyValue/{keyName}][%d] insertUserKeyValueBadRequest ", 400)
}

func (o *InsertUserKeyValueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
