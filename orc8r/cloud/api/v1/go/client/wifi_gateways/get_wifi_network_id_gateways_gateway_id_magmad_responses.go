// Code generated by go-swagger; DO NOT EDIT.

package wifi_gateways

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/api/v1/go/models"
)

// GetWifiNetworkIDGatewaysGatewayIDMagmadReader is a Reader for the GetWifiNetworkIDGatewaysGatewayIDMagmad structure.
type GetWifiNetworkIDGatewaysGatewayIDMagmadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWifiNetworkIDGatewaysGatewayIDMagmadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWifiNetworkIDGatewaysGatewayIDMagmadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetWifiNetworkIDGatewaysGatewayIDMagmadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWifiNetworkIDGatewaysGatewayIDMagmadOK creates a GetWifiNetworkIDGatewaysGatewayIDMagmadOK with default headers values
func NewGetWifiNetworkIDGatewaysGatewayIDMagmadOK() *GetWifiNetworkIDGatewaysGatewayIDMagmadOK {
	return &GetWifiNetworkIDGatewaysGatewayIDMagmadOK{}
}

/*GetWifiNetworkIDGatewaysGatewayIDMagmadOK handles this case with default header values.

Magmad agent configuration
*/
type GetWifiNetworkIDGatewaysGatewayIDMagmadOK struct {
	Payload *models.MagmadGatewayConfigs
}

func (o *GetWifiNetworkIDGatewaysGatewayIDMagmadOK) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/gateways/{gateway_id}/magmad][%d] getWifiNetworkIdGatewaysGatewayIdMagmadOK  %+v", 200, o.Payload)
}

func (o *GetWifiNetworkIDGatewaysGatewayIDMagmadOK) GetPayload() *models.MagmadGatewayConfigs {
	return o.Payload
}

func (o *GetWifiNetworkIDGatewaysGatewayIDMagmadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MagmadGatewayConfigs)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWifiNetworkIDGatewaysGatewayIDMagmadDefault creates a GetWifiNetworkIDGatewaysGatewayIDMagmadDefault with default headers values
func NewGetWifiNetworkIDGatewaysGatewayIDMagmadDefault(code int) *GetWifiNetworkIDGatewaysGatewayIDMagmadDefault {
	return &GetWifiNetworkIDGatewaysGatewayIDMagmadDefault{
		_statusCode: code,
	}
}

/*GetWifiNetworkIDGatewaysGatewayIDMagmadDefault handles this case with default header values.

Unexpected Error
*/
type GetWifiNetworkIDGatewaysGatewayIDMagmadDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get wifi network ID gateways gateway ID magmad default response
func (o *GetWifiNetworkIDGatewaysGatewayIDMagmadDefault) Code() int {
	return o._statusCode
}

func (o *GetWifiNetworkIDGatewaysGatewayIDMagmadDefault) Error() string {
	return fmt.Sprintf("[GET /wifi/{network_id}/gateways/{gateway_id}/magmad][%d] GetWifiNetworkIDGatewaysGatewayIDMagmad default  %+v", o._statusCode, o.Payload)
}

func (o *GetWifiNetworkIDGatewaysGatewayIDMagmadDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetWifiNetworkIDGatewaysGatewayIDMagmadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}