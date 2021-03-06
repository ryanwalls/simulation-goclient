// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/3dsim/simulation-goclient/models"
)

// NewAddPartsParams creates a new AddPartsParams object
// with the default values initialized.
func NewAddPartsParams() *AddPartsParams {
	var ()
	return &AddPartsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddPartsParamsWithTimeout creates a new AddPartsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddPartsParamsWithTimeout(timeout time.Duration) *AddPartsParams {
	var ()
	return &AddPartsParams{

		timeout: timeout,
	}
}

// NewAddPartsParamsWithContext creates a new AddPartsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddPartsParamsWithContext(ctx context.Context) *AddPartsParams {
	var ()
	return &AddPartsParams{

		Context: ctx,
	}
}

// NewAddPartsParamsWithHTTPClient creates a new AddPartsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddPartsParamsWithHTTPClient(client *http.Client) *AddPartsParams {
	var ()
	return &AddPartsParams{
		HTTPClient: client,
	}
}

/*AddPartsParams contains all the parameters to send to the API endpoint
for the add parts operation typically these are written to a http.Request
*/
type AddPartsParams struct {

	/*Part
	  Part to add. First, call parts/geometryurl to get a PartUploadRequest object. PUT the STL file AmazonS3 with the signed URL. Use the s3Key property as the value of fileLocation. A POST with this object will execute the part processing service.

	*/
	Part *models.PartPost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add parts params
func (o *AddPartsParams) WithTimeout(timeout time.Duration) *AddPartsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add parts params
func (o *AddPartsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add parts params
func (o *AddPartsParams) WithContext(ctx context.Context) *AddPartsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add parts params
func (o *AddPartsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add parts params
func (o *AddPartsParams) WithHTTPClient(client *http.Client) *AddPartsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add parts params
func (o *AddPartsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPart adds the part to the add parts params
func (o *AddPartsParams) WithPart(part *models.PartPost) *AddPartsParams {
	o.SetPart(part)
	return o
}

// SetPart adds the part to the add parts params
func (o *AddPartsParams) SetPart(part *models.PartPost) {
	o.Part = part
}

// WriteToRequest writes these params to a swagger request
func (o *AddPartsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Part == nil {
		o.Part = new(models.PartPost)
	}

	if err := r.SetBodyParam(o.Part); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
