package kv

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
)

// NewGetValueParams creates a new GetValueParams object
// with the default values initialized.
func NewGetValueParams() *GetValueParams {
	var ()
	return &GetValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetValueParamsWithTimeout creates a new GetValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetValueParamsWithTimeout(timeout time.Duration) *GetValueParams {
	var ()
	return &GetValueParams{

		timeout: timeout,
	}
}

// NewGetValueParamsWithContext creates a new GetValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetValueParamsWithContext(ctx context.Context) *GetValueParams {
	var ()
	return &GetValueParams{

		Context: ctx,
	}
}

// NewGetValueParamsWithHTTPClient creates a new GetValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetValueParamsWithHTTPClient(client *http.Client) *GetValueParams {
	var ()
	return &GetValueParams{
		HTTPClient: client,
	}
}

/*GetValueParams contains all the parameters to send to the API endpoint
for the get value operation typically these are written to a http.Request
*/
type GetValueParams struct {

	/*OpID*/
	OpID *string
	/*Key*/
	Key string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get value params
func (o *GetValueParams) WithTimeout(timeout time.Duration) *GetValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get value params
func (o *GetValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get value params
func (o *GetValueParams) WithContext(ctx context.Context) *GetValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get value params
func (o *GetValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get value params
func (o *GetValueParams) WithHTTPClient(client *http.Client) *GetValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get value params
func (o *GetValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpID adds the opID to the get value params
func (o *GetValueParams) WithOpID(opID *string) *GetValueParams {
	o.SetOpID(opID)
	return o
}

// SetOpID adds the opId to the get value params
func (o *GetValueParams) SetOpID(opID *string) {
	o.OpID = opID
}

// WithKey adds the key to the get value params
func (o *GetValueParams) WithKey(key string) *GetValueParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the get value params
func (o *GetValueParams) SetKey(key string) {
	o.Key = key
}

// WriteToRequest writes these params to a swagger request
func (o *GetValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.OpID != nil {

		// header param Op-ID
		if err := r.SetHeaderParam("Op-ID", *o.OpID); err != nil {
			return err
		}

	}

	// path param key
	if err := r.SetPathParam("key", o.Key); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
