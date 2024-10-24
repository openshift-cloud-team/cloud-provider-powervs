// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_pod_capacity

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

// NewPcloudPodcapacityGetParams creates a new PcloudPodcapacityGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPodcapacityGetParams() *PcloudPodcapacityGetParams {
	return &PcloudPodcapacityGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPodcapacityGetParamsWithTimeout creates a new PcloudPodcapacityGetParams object
// with the ability to set a timeout on a request.
func NewPcloudPodcapacityGetParamsWithTimeout(timeout time.Duration) *PcloudPodcapacityGetParams {
	return &PcloudPodcapacityGetParams{
		timeout: timeout,
	}
}

// NewPcloudPodcapacityGetParamsWithContext creates a new PcloudPodcapacityGetParams object
// with the ability to set a context for a request.
func NewPcloudPodcapacityGetParamsWithContext(ctx context.Context) *PcloudPodcapacityGetParams {
	return &PcloudPodcapacityGetParams{
		Context: ctx,
	}
}

// NewPcloudPodcapacityGetParamsWithHTTPClient creates a new PcloudPodcapacityGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPodcapacityGetParamsWithHTTPClient(client *http.Client) *PcloudPodcapacityGetParams {
	return &PcloudPodcapacityGetParams{
		HTTPClient: client,
	}
}

/*
PcloudPodcapacityGetParams contains all the parameters to send to the API endpoint

	for the pcloud podcapacity get operation.

	Typically these are written to a http.Request.
*/
type PcloudPodcapacityGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud podcapacity get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPodcapacityGetParams) WithDefaults() *PcloudPodcapacityGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud podcapacity get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPodcapacityGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud podcapacity get params
func (o *PcloudPodcapacityGetParams) WithTimeout(timeout time.Duration) *PcloudPodcapacityGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud podcapacity get params
func (o *PcloudPodcapacityGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud podcapacity get params
func (o *PcloudPodcapacityGetParams) WithContext(ctx context.Context) *PcloudPodcapacityGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud podcapacity get params
func (o *PcloudPodcapacityGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud podcapacity get params
func (o *PcloudPodcapacityGetParams) WithHTTPClient(client *http.Client) *PcloudPodcapacityGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud podcapacity get params
func (o *PcloudPodcapacityGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud podcapacity get params
func (o *PcloudPodcapacityGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPodcapacityGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud podcapacity get params
func (o *PcloudPodcapacityGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPodcapacityGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}