// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// GetScanPatternSimulationReader is a Reader for the GetScanPatternSimulation structure.
type GetScanPatternSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScanPatternSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetScanPatternSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetScanPatternSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetScanPatternSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetScanPatternSimulationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetScanPatternSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetScanPatternSimulationOK creates a GetScanPatternSimulationOK with default headers values
func NewGetScanPatternSimulationOK() *GetScanPatternSimulationOK {
	return &GetScanPatternSimulationOK{}
}

/*GetScanPatternSimulationOK handles this case with default header values.

Successfully retrieved simulation
*/
type GetScanPatternSimulationOK struct {
	Payload *models.ScanPatternSimulation
}

func (o *GetScanPatternSimulationOK) Error() string {
	return fmt.Sprintf("[GET /scanpatternsimulations/{id}][%d] getScanPatternSimulationOK  %+v", 200, o.Payload)
}

func (o *GetScanPatternSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScanPatternSimulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanPatternSimulationUnauthorized creates a GetScanPatternSimulationUnauthorized with default headers values
func NewGetScanPatternSimulationUnauthorized() *GetScanPatternSimulationUnauthorized {
	return &GetScanPatternSimulationUnauthorized{}
}

/*GetScanPatternSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type GetScanPatternSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *GetScanPatternSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scanpatternsimulations/{id}][%d] getScanPatternSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScanPatternSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanPatternSimulationForbidden creates a GetScanPatternSimulationForbidden with default headers values
func NewGetScanPatternSimulationForbidden() *GetScanPatternSimulationForbidden {
	return &GetScanPatternSimulationForbidden{}
}

/*GetScanPatternSimulationForbidden handles this case with default header values.

Forbidden
*/
type GetScanPatternSimulationForbidden struct {
	Payload *models.Error
}

func (o *GetScanPatternSimulationForbidden) Error() string {
	return fmt.Sprintf("[GET /scanpatternsimulations/{id}][%d] getScanPatternSimulationForbidden  %+v", 403, o.Payload)
}

func (o *GetScanPatternSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanPatternSimulationNotFound creates a GetScanPatternSimulationNotFound with default headers values
func NewGetScanPatternSimulationNotFound() *GetScanPatternSimulationNotFound {
	return &GetScanPatternSimulationNotFound{}
}

/*GetScanPatternSimulationNotFound handles this case with default header values.

Simulation not found (id invalid)
*/
type GetScanPatternSimulationNotFound struct {
	Payload *models.Error
}

func (o *GetScanPatternSimulationNotFound) Error() string {
	return fmt.Sprintf("[GET /scanpatternsimulations/{id}][%d] getScanPatternSimulationNotFound  %+v", 404, o.Payload)
}

func (o *GetScanPatternSimulationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScanPatternSimulationDefault creates a GetScanPatternSimulationDefault with default headers values
func NewGetScanPatternSimulationDefault(code int) *GetScanPatternSimulationDefault {
	return &GetScanPatternSimulationDefault{
		_statusCode: code,
	}
}

/*GetScanPatternSimulationDefault handles this case with default header values.

unexpected error
*/
type GetScanPatternSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get scan pattern simulation default response
func (o *GetScanPatternSimulationDefault) Code() int {
	return o._statusCode
}

func (o *GetScanPatternSimulationDefault) Error() string {
	return fmt.Sprintf("[GET /scanpatternsimulations/{id}][%d] getScanPatternSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *GetScanPatternSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
