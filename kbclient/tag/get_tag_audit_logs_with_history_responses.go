// Code generated by go-swagger; DO NOT EDIT.

package tag

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetTagAuditLogsWithHistoryReader is a Reader for the GetTagAuditLogsWithHistory structure.
type GetTagAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTagAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTagAuditLogsWithHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetTagAuditLogsWithHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTagAuditLogsWithHistoryOK creates a GetTagAuditLogsWithHistoryOK with default headers values
func NewGetTagAuditLogsWithHistoryOK() *GetTagAuditLogsWithHistoryOK {
	return &GetTagAuditLogsWithHistoryOK{}
}

/* GetTagAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetTagAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog
}

func (o *GetTagAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tags/{tagId}/auditLogsWithHistory][%d] getTagAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}
func (o *GetTagAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetTagAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTagAuditLogsWithHistoryNotFound creates a GetTagAuditLogsWithHistoryNotFound with default headers values
func NewGetTagAuditLogsWithHistoryNotFound() *GetTagAuditLogsWithHistoryNotFound {
	return &GetTagAuditLogsWithHistoryNotFound{}
}

/* GetTagAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetTagAuditLogsWithHistoryNotFound struct {
}

func (o *GetTagAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tags/{tagId}/auditLogsWithHistory][%d] getTagAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetTagAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
