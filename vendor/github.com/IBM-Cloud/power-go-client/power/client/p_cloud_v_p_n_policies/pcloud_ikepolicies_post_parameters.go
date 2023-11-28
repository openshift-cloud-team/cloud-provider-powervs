// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

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

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudIkepoliciesPostParams creates a new PcloudIkepoliciesPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudIkepoliciesPostParams() *PcloudIkepoliciesPostParams {
	return &PcloudIkepoliciesPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudIkepoliciesPostParamsWithTimeout creates a new PcloudIkepoliciesPostParams object
// with the ability to set a timeout on a request.
func NewPcloudIkepoliciesPostParamsWithTimeout(timeout time.Duration) *PcloudIkepoliciesPostParams {
	return &PcloudIkepoliciesPostParams{
		timeout: timeout,
	}
}

// NewPcloudIkepoliciesPostParamsWithContext creates a new PcloudIkepoliciesPostParams object
// with the ability to set a context for a request.
func NewPcloudIkepoliciesPostParamsWithContext(ctx context.Context) *PcloudIkepoliciesPostParams {
	return &PcloudIkepoliciesPostParams{
		Context: ctx,
	}
}

// NewPcloudIkepoliciesPostParamsWithHTTPClient creates a new PcloudIkepoliciesPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudIkepoliciesPostParamsWithHTTPClient(client *http.Client) *PcloudIkepoliciesPostParams {
	return &PcloudIkepoliciesPostParams{
		HTTPClient: client,
	}
}

/*
PcloudIkepoliciesPostParams contains all the parameters to send to the API endpoint

	for the pcloud ikepolicies post operation.

	Typically these are written to a http.Request.
*/
type PcloudIkepoliciesPostParams struct {

	/* Body.

	   Parameters for the creation of a new IKE Policy
	*/
	Body *models.IKEPolicyCreate

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud ikepolicies post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIkepoliciesPostParams) WithDefaults() *PcloudIkepoliciesPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud ikepolicies post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudIkepoliciesPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) WithTimeout(timeout time.Duration) *PcloudIkepoliciesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) WithContext(ctx context.Context) *PcloudIkepoliciesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) WithHTTPClient(client *http.Client) *PcloudIkepoliciesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) WithBody(body *models.IKEPolicyCreate) *PcloudIkepoliciesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) SetBody(body *models.IKEPolicyCreate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudIkepoliciesPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud ikepolicies post params
func (o *PcloudIkepoliciesPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudIkepoliciesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
