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

// PostSingleBeadSimulationReader is a Reader for the PostSingleBeadSimulation structure.
type PostSingleBeadSimulationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSingleBeadSimulationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostSingleBeadSimulationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostSingleBeadSimulationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostSingleBeadSimulationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPostSingleBeadSimulationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPostSingleBeadSimulationOK creates a PostSingleBeadSimulationOK with default headers values
func NewPostSingleBeadSimulationOK() *PostSingleBeadSimulationOK {
	return &PostSingleBeadSimulationOK{}
}

/*PostSingleBeadSimulationOK handles this case with default header values.

Successfully added a simulation
*/
type PostSingleBeadSimulationOK struct {
	Payload *models.SingleBeadSimulation
}

func (o *PostSingleBeadSimulationOK) Error() string {
	return fmt.Sprintf("[POST /singlebeadsimulations][%d] postSingleBeadSimulationOK  %+v", 200, o.Payload)
}

func (o *PostSingleBeadSimulationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SingleBeadSimulation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSingleBeadSimulationUnauthorized creates a PostSingleBeadSimulationUnauthorized with default headers values
func NewPostSingleBeadSimulationUnauthorized() *PostSingleBeadSimulationUnauthorized {
	return &PostSingleBeadSimulationUnauthorized{}
}

/*PostSingleBeadSimulationUnauthorized handles this case with default header values.

Not authorized
*/
type PostSingleBeadSimulationUnauthorized struct {
	Payload *models.Error
}

func (o *PostSingleBeadSimulationUnauthorized) Error() string {
	return fmt.Sprintf("[POST /singlebeadsimulations][%d] postSingleBeadSimulationUnauthorized  %+v", 401, o.Payload)
}

func (o *PostSingleBeadSimulationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSingleBeadSimulationForbidden creates a PostSingleBeadSimulationForbidden with default headers values
func NewPostSingleBeadSimulationForbidden() *PostSingleBeadSimulationForbidden {
	return &PostSingleBeadSimulationForbidden{}
}

/*PostSingleBeadSimulationForbidden handles this case with default header values.

Forbidden
*/
type PostSingleBeadSimulationForbidden struct {
	Payload *models.Error
}

func (o *PostSingleBeadSimulationForbidden) Error() string {
	return fmt.Sprintf("[POST /singlebeadsimulations][%d] postSingleBeadSimulationForbidden  %+v", 403, o.Payload)
}

func (o *PostSingleBeadSimulationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSingleBeadSimulationDefault creates a PostSingleBeadSimulationDefault with default headers values
func NewPostSingleBeadSimulationDefault(code int) *PostSingleBeadSimulationDefault {
	return &PostSingleBeadSimulationDefault{
		_statusCode: code,
	}
}

/*PostSingleBeadSimulationDefault handles this case with default header values.

unexpected error
*/
type PostSingleBeadSimulationDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the post single bead simulation default response
func (o *PostSingleBeadSimulationDefault) Code() int {
	return o._statusCode
}

func (o *PostSingleBeadSimulationDefault) Error() string {
	return fmt.Sprintf("[POST /singlebeadsimulations][%d] postSingleBeadSimulation default  %+v", o._statusCode, o.Payload)
}

func (o *PostSingleBeadSimulationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
