// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/CDNA-Technologies/kbcli/v3/kbmodel"
)

// GetAccountAuditLogsWithHistoryReader is a Reader for the GetAccountAuditLogsWithHistory structure.
type GetAccountAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAccountAuditLogsWithHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetAccountAuditLogsWithHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAccountAuditLogsWithHistoryOK creates a GetAccountAuditLogsWithHistoryOK with default headers values
func NewGetAccountAuditLogsWithHistoryOK() *GetAccountAuditLogsWithHistoryOK {
	return &GetAccountAuditLogsWithHistoryOK{}
}

/* GetAccountAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAccountAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog
}

func (o *GetAccountAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/auditLogsWithHistory][%d] getAccountAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}
func (o *GetAccountAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetAccountAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAuditLogsWithHistoryNotFound creates a GetAccountAuditLogsWithHistoryNotFound with default headers values
func NewGetAccountAuditLogsWithHistoryNotFound() *GetAccountAuditLogsWithHistoryNotFound {
	return &GetAccountAuditLogsWithHistoryNotFound{}
}

/* GetAccountAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetAccountAuditLogsWithHistoryNotFound struct {
}

func (o *GetAccountAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/auditLogsWithHistory][%d] getAccountAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetAccountAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
