// Code generated by go-swagger; DO NOT EDIT.

package security

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

// AddUserRolesReader is a Reader for the AddUserRoles structure.
type AddUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewAddUserRolesCreated()
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

// NewAddUserRolesCreated creates a AddUserRolesCreated with default headers values
func NewAddUserRolesCreated() *AddUserRolesCreated {
	return &AddUserRolesCreated{}
}

/*AddUserRolesCreated handles this case with default header values.

User role created successfully
*/
type AddUserRolesCreated struct {
	Payload *kbmodel.UserRoles

	HttpResponse runtime.ClientResponse
}

func (o *AddUserRolesCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/security/users][%d] addUserRolesCreated  %+v", 201, o.Payload)
}

func (o *AddUserRolesCreated) GetPayload() *kbmodel.UserRoles {
	return o.Payload
}

func (o *AddUserRolesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.UserRoles)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
