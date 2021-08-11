// Code generated by go-swagger; DO NOT EDIT.

package overdue

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

// GetOverdueConfigJSONReader is a Reader for the GetOverdueConfigJSON structure.
type GetOverdueConfigJSONReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOverdueConfigJSONReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOverdueConfigJSONOK()
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

// NewGetOverdueConfigJSONOK creates a GetOverdueConfigJSONOK with default headers values
func NewGetOverdueConfigJSONOK() *GetOverdueConfigJSONOK {
	return &GetOverdueConfigJSONOK{}
}

/*GetOverdueConfigJSONOK handles this case with default header values.

successful operation
*/
type GetOverdueConfigJSONOK struct {
	Payload *kbmodel.Overdue

	HttpResponse runtime.ClientResponse
}

func (o *GetOverdueConfigJSONOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/overdue][%d] getOverdueConfigJsonOK  %+v", 200, o.Payload)
}

func (o *GetOverdueConfigJSONOK) GetPayload() *kbmodel.Overdue {
	return o.Payload
}

func (o *GetOverdueConfigJSONOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Overdue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
