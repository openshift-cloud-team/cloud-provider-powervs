// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

// NewPcloudPvminstancesNetworksGetParams creates a new PcloudPvminstancesNetworksGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesNetworksGetParams() *PcloudPvminstancesNetworksGetParams {
	return &PcloudPvminstancesNetworksGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesNetworksGetParamsWithTimeout creates a new PcloudPvminstancesNetworksGetParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesNetworksGetParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesNetworksGetParams {
	return &PcloudPvminstancesNetworksGetParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesNetworksGetParamsWithContext creates a new PcloudPvminstancesNetworksGetParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesNetworksGetParamsWithContext(ctx context.Context) *PcloudPvminstancesNetworksGetParams {
	return &PcloudPvminstancesNetworksGetParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesNetworksGetParamsWithHTTPClient creates a new PcloudPvminstancesNetworksGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesNetworksGetParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesNetworksGetParams {
	return &PcloudPvminstancesNetworksGetParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesNetworksGetParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances networks get operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesNetworksGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances networks get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesNetworksGetParams) WithDefaults() *PcloudPvminstancesNetworksGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances networks get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesNetworksGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesNetworksGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) WithContext(ctx context.Context) *PcloudPvminstancesNetworksGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesNetworksGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesNetworksGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithNetworkID adds the networkID to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) WithNetworkID(networkID string) *PcloudPvminstancesNetworksGetParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesNetworksGetParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances networks get params
func (o *PcloudPvminstancesNetworksGetParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesNetworksGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
