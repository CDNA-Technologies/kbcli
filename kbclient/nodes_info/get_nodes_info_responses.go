// Code generated by go-swagger; DO NOT EDIT.

package nodes_info

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

// GetNodesInfoReader is a Reader for the GetNodesInfo structure.
type GetNodesInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodesInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodesInfoOK()
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

// NewGetNodesInfoOK creates a GetNodesInfoOK with default headers values
func NewGetNodesInfoOK() *GetNodesInfoOK {
	return &GetNodesInfoOK{}
}

/*GetNodesInfoOK handles this case with default header values.

successful operation
*/
type GetNodesInfoOK struct {
	Payload []*kbmodel.PluginInfo

	HttpResponse runtime.ClientResponse
}

func (o *GetNodesInfoOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/nodesInfo][%d] getNodesInfoOK  %+v", 200, o.Payload)
}

func (o *GetNodesInfoOK) GetPayload() []*kbmodel.PluginInfo {
	return o.Payload
}

func (o *GetNodesInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
