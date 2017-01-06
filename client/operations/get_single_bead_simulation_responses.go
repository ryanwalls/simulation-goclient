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

// GetSingleBeadSimulationReader is a Reader for the GetSingleBeadSimulation structure.
type GetSingleBeadSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSingleBeadSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSingleBeadSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSingleBeadSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSingleBeadSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetSingleBeadSimulationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetSingleBeadSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSingleBeadSimulationOK creates a GetSingleBeadSimulationOK with default headers values
func NewGetSingleBeadSimulationOK() *GetSingleBeadSimulationOK {
	return &GetSingleBeadSimulationOK{}
}

/*GetSingleBeadSimulationOK handles this case with default header values.

Successfully retrieved simulation
*/
type GetSingleBeadSimulationOK struct {
	Payload *models.SingleBeadSimulation
}

func (o *GetSingleBeadSimulationOK) Error() string {
	return fmt.Sprintf("[GET /singlebeadsimulations/{id}][%d] getSingleBeadSimulationOK  %+v", 200, o.Payload)
}

func (o *GetSingleBeadSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SingleBeadSimulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleBeadSimulationUnauthorized creates a GetSingleBeadSimulationUnauthorized with default headers values
func NewGetSingleBeadSimulationUnauthorized() *GetSingleBeadSimulationUnauthorized {
	return &GetSingleBeadSimulationUnauthorized{}
}

/*GetSingleBeadSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type GetSingleBeadSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *GetSingleBeadSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /singlebeadsimulations/{id}][%d] getSingleBeadSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSingleBeadSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleBeadSimulationForbidden creates a GetSingleBeadSimulationForbidden with default headers values
func NewGetSingleBeadSimulationForbidden() *GetSingleBeadSimulationForbidden {
	return &GetSingleBeadSimulationForbidden{}
}

/*GetSingleBeadSimulationForbidden handles this case with default header values.

Forbidden
*/
type GetSingleBeadSimulationForbidden struct {
	Payload *models.Error
}

func (o *GetSingleBeadSimulationForbidden) Error() string {
	return fmt.Sprintf("[GET /singlebeadsimulations/{id}][%d] getSingleBeadSimulationForbidden  %+v", 403, o.Payload)
}

func (o *GetSingleBeadSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleBeadSimulationNotFound creates a GetSingleBeadSimulationNotFound with default headers values
func NewGetSingleBeadSimulationNotFound() *GetSingleBeadSimulationNotFound {
	return &GetSingleBeadSimulationNotFound{}
}

/*GetSingleBeadSimulationNotFound handles this case with default header values.

Simulation not found (id invalid)
*/
type GetSingleBeadSimulationNotFound struct {
	Payload *models.Error
}

func (o *GetSingleBeadSimulationNotFound) Error() string {
	return fmt.Sprintf("[GET /singlebeadsimulations/{id}][%d] getSingleBeadSimulationNotFound  %+v", 404, o.Payload)
}

func (o *GetSingleBeadSimulationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSingleBeadSimulationDefault creates a GetSingleBeadSimulationDefault with default headers values
func NewGetSingleBeadSimulationDefault(code int) *GetSingleBeadSimulationDefault {
	return &GetSingleBeadSimulationDefault{
		_statusCode: code,
	}
}

/*GetSingleBeadSimulationDefault handles this case with default header values.

unexpected error
*/
type GetSingleBeadSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get single bead simulation default response
func (o *GetSingleBeadSimulationDefault) Code() int {
	return o._statusCode
}

func (o *GetSingleBeadSimulationDefault) Error() string {
	return fmt.Sprintf("[GET /singlebeadsimulations/{id}][%d] getSingleBeadSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *GetSingleBeadSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
