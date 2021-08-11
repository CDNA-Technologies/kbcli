// Code generated by go-swagger; DO NOT EDIT.

package account

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

// CreateAccountReader is a Reader for the CreateAccount structure.
type CreateAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateAccountCreated()
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

// NewCreateAccountCreated creates a CreateAccountCreated with default headers values
func NewCreateAccountCreated() *CreateAccountCreated {
	return &CreateAccountCreated{}
}

/*CreateAccountCreated handles this case with default header values.

Account created successfully
*/
type CreateAccountCreated struct {
	Payload *kbmodel.Account

	HttpResponse runtime.ClientResponse
}

func (o *CreateAccountCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts][%d] createAccountCreated  %+v", 201, o.Payload)
}

func (o *CreateAccountCreated) GetPayload() *kbmodel.Account {
	return o.Payload
}

func (o *CreateAccountCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAccountBadRequest creates a CreateAccountBadRequest with default headers values
func NewCreateAccountBadRequest() *CreateAccountBadRequest {
	return &CreateAccountBadRequest{}
}

/*CreateAccountBadRequest handles this case with default header values.

Invalid account data supplied
*/
type CreateAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateAccountBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts][%d] createAccountBadRequest ", 400)
}

func (o *CreateAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
