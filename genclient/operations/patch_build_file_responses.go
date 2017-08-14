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

// PatchBuildFileReader is a Reader for the PatchBuildFile structure.
type PatchBuildFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchBuildFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchBuildFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPatchBuildFileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPatchBuildFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchBuildFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchBuildFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchBuildFileOK creates a PatchBuildFileOK with default headers values
func NewPatchBuildFileOK() *PatchBuildFileOK {
	return &PatchBuildFileOK{}
}

/*PatchBuildFileOK handles this case with default header values.

Patched build file
*/
type PatchBuildFileOK struct {
	Payload *models.BuildFile
}

func (o *PatchBuildFileOK) Error() string {
	return fmt.Sprintf("[PATCH /buildfiles/{id}][%d] patchBuildFileOK  %+v", 200, o.Payload)
}

func (o *PatchBuildFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BuildFile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBuildFileUnauthorized creates a PatchBuildFileUnauthorized with default headers values
func NewPatchBuildFileUnauthorized() *PatchBuildFileUnauthorized {
	return &PatchBuildFileUnauthorized{}
}

/*PatchBuildFileUnauthorized handles this case with default header values.

Not authorized
*/
type PatchBuildFileUnauthorized struct {
	Payload *models.Error
}

func (o *PatchBuildFileUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /buildfiles/{id}][%d] patchBuildFileUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchBuildFileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBuildFileForbidden creates a PatchBuildFileForbidden with default headers values
func NewPatchBuildFileForbidden() *PatchBuildFileForbidden {
	return &PatchBuildFileForbidden{}
}

/*PatchBuildFileForbidden handles this case with default header values.

Forbidden
*/
type PatchBuildFileForbidden struct {
	Payload *models.Error
}

func (o *PatchBuildFileForbidden) Error() string {
	return fmt.Sprintf("[PATCH /buildfiles/{id}][%d] patchBuildFileForbidden  %+v", 403, o.Payload)
}

func (o *PatchBuildFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBuildFileNotFound creates a PatchBuildFileNotFound with default headers values
func NewPatchBuildFileNotFound() *PatchBuildFileNotFound {
	return &PatchBuildFileNotFound{}
}

/*PatchBuildFileNotFound handles this case with default header values.

Simulation not found (id invalid)
*/
type PatchBuildFileNotFound struct {
	Payload *models.Error
}

func (o *PatchBuildFileNotFound) Error() string {
	return fmt.Sprintf("[PATCH /buildfiles/{id}][%d] patchBuildFileNotFound  %+v", 404, o.Payload)
}

func (o *PatchBuildFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchBuildFileDefault creates a PatchBuildFileDefault with default headers values
func NewPatchBuildFileDefault(code int) *PatchBuildFileDefault {
	return &PatchBuildFileDefault{
		_statusCode: code,
	}
}

/*PatchBuildFileDefault handles this case with default header values.

unexpected error
*/
type PatchBuildFileDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch build file default response
func (o *PatchBuildFileDefault) Code() int {
	return o._statusCode
}

func (o *PatchBuildFileDefault) Error() string {
	return fmt.Sprintf("[PATCH /buildfiles/{id}][%d] patchBuildFile default  %+v", o._statusCode, o.Payload)
}

func (o *PatchBuildFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
