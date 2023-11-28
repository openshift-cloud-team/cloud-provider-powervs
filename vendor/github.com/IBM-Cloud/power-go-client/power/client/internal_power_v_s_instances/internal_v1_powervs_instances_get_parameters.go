// Code generated by go-swagger; DO NOT EDIT.

package internal_power_v_s_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewInternalV1PowervsInstancesGetParams creates a new InternalV1PowervsInstancesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1PowervsInstancesGetParams() *InternalV1PowervsInstancesGetParams {
	return &InternalV1PowervsInstancesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1PowervsInstancesGetParamsWithTimeout creates a new InternalV1PowervsInstancesGetParams object
// with the ability to set a timeout on a request.
func NewInternalV1PowervsInstancesGetParamsWithTimeout(timeout time.Duration) *InternalV1PowervsInstancesGetParams {
	return &InternalV1PowervsInstancesGetParams{
		timeout: timeout,
	}
}

// NewInternalV1PowervsInstancesGetParamsWithContext creates a new InternalV1PowervsInstancesGetParams object
// with the ability to set a context for a request.
func NewInternalV1PowervsInstancesGetParamsWithContext(ctx context.Context) *InternalV1PowervsInstancesGetParams {
	return &InternalV1PowervsInstancesGetParams{
		Context: ctx,
	}
}

// NewInternalV1PowervsInstancesGetParamsWithHTTPClient creates a new InternalV1PowervsInstancesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1PowervsInstancesGetParamsWithHTTPClient(client *http.Client) *InternalV1PowervsInstancesGetParams {
	return &InternalV1PowervsInstancesGetParams{
		HTTPClient: client,
	}
}

/*
InternalV1PowervsInstancesGetParams contains all the parameters to send to the API endpoint

	for the internal v1 powervs instances get operation.

	Typically these are written to a http.Request.
*/
type InternalV1PowervsInstancesGetParams struct {

	/* PowervsLocation.

	   The PowerVS Location
	*/
	PowervsLocation *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 powervs instances get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1PowervsInstancesGetParams) WithDefaults() *InternalV1PowervsInstancesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 powervs instances get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1PowervsInstancesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 powervs instances get params
func (o *InternalV1PowervsInstancesGetParams) WithTimeout(timeout time.Duration) *InternalV1PowervsInstancesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 powervs instances get params
func (o *InternalV1PowervsInstancesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 powervs instances get params
func (o *InternalV1PowervsInstancesGetParams) WithContext(ctx context.Context) *InternalV1PowervsInstancesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 powervs instances get params
func (o *InternalV1PowervsInstancesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 powervs instances get params
func (o *InternalV1PowervsInstancesGetParams) WithHTTPClient(client *http.Client) *InternalV1PowervsInstancesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 powervs instances get params
func (o *InternalV1PowervsInstancesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPowervsLocation adds the powervsLocation to the internal v1 powervs instances get params
func (o *InternalV1PowervsInstancesGetParams) WithPowervsLocation(powervsLocation *string) *InternalV1PowervsInstancesGetParams {
	o.SetPowervsLocation(powervsLocation)
	return o
}

// SetPowervsLocation adds the powervsLocation to the internal v1 powervs instances get params
func (o *InternalV1PowervsInstancesGetParams) SetPowervsLocation(powervsLocation *string) {
	o.PowervsLocation = powervsLocation
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1PowervsInstancesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PowervsLocation != nil {

		// query param powervs_location
		var qrPowervsLocation string

		if o.PowervsLocation != nil {
			qrPowervsLocation = *o.PowervsLocation
		}
		qPowervsLocation := qrPowervsLocation
		if qPowervsLocation != "" {

			if err := r.SetQueryParam("powervs_location", qPowervsLocation); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
