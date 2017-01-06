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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUnarchiveMachineParams creates a new UnarchiveMachineParams object
// with the default values initialized.
func NewUnarchiveMachineParams() *UnarchiveMachineParams {
	var ()
	return &UnarchiveMachineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUnarchiveMachineParamsWithTimeout creates a new UnarchiveMachineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUnarchiveMachineParamsWithTimeout(timeout time.Duration) *UnarchiveMachineParams {
	var ()
	return &UnarchiveMachineParams{

		timeout: timeout,
	}
}

// NewUnarchiveMachineParamsWithContext creates a new UnarchiveMachineParams object
// with the default values initialized, and the ability to set a context for a request
func NewUnarchiveMachineParamsWithContext(ctx context.Context) *UnarchiveMachineParams {
	var ()
	return &UnarchiveMachineParams{

		Context: ctx,
	}
}

/*UnarchiveMachineParams contains all the parameters to send to the API endpoint
for the unarchive machine operation typically these are written to a http.Request
*/
type UnarchiveMachineParams struct {

	/*ID
	  machine identifier

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the unarchive machine params
func (o *UnarchiveMachineParams) WithTimeout(timeout time.Duration) *UnarchiveMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unarchive machine params
func (o *UnarchiveMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unarchive machine params
func (o *UnarchiveMachineParams) WithContext(ctx context.Context) *UnarchiveMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unarchive machine params
func (o *UnarchiveMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithID adds the id to the unarchive machine params
func (o *UnarchiveMachineParams) WithID(id int64) *UnarchiveMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the unarchive machine params
func (o *UnarchiveMachineParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UnarchiveMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
