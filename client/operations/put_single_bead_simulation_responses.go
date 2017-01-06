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

// PutSingleBeadSimulationReader is a Reader for the PutSingleBeadSimulation structure.
type PutSingleBeadSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutSingleBeadSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutSingleBeadSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPutSingleBeadSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutSingleBeadSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPutSingleBeadSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutSingleBeadSimulationOK creates a PutSingleBeadSimulationOK with default headers values
func NewPutSingleBeadSimulationOK() *PutSingleBeadSimulationOK {
	return &PutSingleBeadSimulationOK{}
}

/*PutSingleBeadSimulationOK handles this case with default header values.

Successfully updated a simulation
*/
type PutSingleBeadSimulationOK struct {
	Payload *models.SingleBeadSimulation
}

func (o *PutSingleBeadSimulationOK) Error() string {
	return fmt.Sprintf("[PUT /singlebeadsimulations/{id}][%d] putSingleBeadSimulationOK  %+v", 200, o.Payload)
}

func (o *PutSingleBeadSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SingleBeadSimulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSingleBeadSimulationUnauthorized creates a PutSingleBeadSimulationUnauthorized with default headers values
func NewPutSingleBeadSimulationUnauthorized() *PutSingleBeadSimulationUnauthorized {
	return &PutSingleBeadSimulationUnauthorized{}
}

/*PutSingleBeadSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type PutSingleBeadSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *PutSingleBeadSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /singlebeadsimulations/{id}][%d] putSingleBeadSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *PutSingleBeadSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSingleBeadSimulationForbidden creates a PutSingleBeadSimulationForbidden with default headers values
func NewPutSingleBeadSimulationForbidden() *PutSingleBeadSimulationForbidden {
	return &PutSingleBeadSimulationForbidden{}
}

/*PutSingleBeadSimulationForbidden handles this case with default header values.

Forbidden
*/
type PutSingleBeadSimulationForbidden struct {
	Payload *models.Error
}

func (o *PutSingleBeadSimulationForbidden) Error() string {
	return fmt.Sprintf("[PUT /singlebeadsimulations/{id}][%d] putSingleBeadSimulationForbidden  %+v", 403, o.Payload)
}

func (o *PutSingleBeadSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutSingleBeadSimulationDefault creates a PutSingleBeadSimulationDefault with default headers values
func NewPutSingleBeadSimulationDefault(code int) *PutSingleBeadSimulationDefault {
	return &PutSingleBeadSimulationDefault{
		_statusCode: code,
	}
}

/*PutSingleBeadSimulationDefault handles this case with default header values.

unexpected error
*/
type PutSingleBeadSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put single bead simulation default response
func (o *PutSingleBeadSimulationDefault) Code() int {
	return o._statusCode
}

func (o *PutSingleBeadSimulationDefault) Error() string {
	return fmt.Sprintf("[PUT /singlebeadsimulations/{id}][%d] putSingleBeadSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *PutSingleBeadSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
