// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_disaster_recovery

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new p cloud disaster recovery API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new p cloud disaster recovery API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new p cloud disaster recovery API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for p cloud disaster recovery API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	PcloudLocationsDisasterrecoveryGet(params *PcloudLocationsDisasterrecoveryGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudLocationsDisasterrecoveryGetOK, error)

	PcloudLocationsDisasterrecoveryGetall(params *PcloudLocationsDisasterrecoveryGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudLocationsDisasterrecoveryGetallOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
PcloudLocationsDisasterrecoveryGet gets the disaster recovery site details for the current location
*/
func (a *Client) PcloudLocationsDisasterrecoveryGet(params *PcloudLocationsDisasterrecoveryGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudLocationsDisasterrecoveryGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudLocationsDisasterrecoveryGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.locations.disasterrecovery.get",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/cloud-instances/{cloud_instance_id}/locations/disaster-recovery",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudLocationsDisasterrecoveryGetReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudLocationsDisasterrecoveryGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.locations.disasterrecovery.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PcloudLocationsDisasterrecoveryGetall gets all disaster recovery locations supported by power virtual server
*/
func (a *Client) PcloudLocationsDisasterrecoveryGetall(params *PcloudLocationsDisasterrecoveryGetallParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PcloudLocationsDisasterrecoveryGetallOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPcloudLocationsDisasterrecoveryGetallParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "pcloud.locations.disasterrecovery.getall",
		Method:             "GET",
		PathPattern:        "/pcloud/v1/locations/disaster-recovery",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PcloudLocationsDisasterrecoveryGetallReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PcloudLocationsDisasterrecoveryGetallOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for pcloud.locations.disasterrecovery.getall: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
