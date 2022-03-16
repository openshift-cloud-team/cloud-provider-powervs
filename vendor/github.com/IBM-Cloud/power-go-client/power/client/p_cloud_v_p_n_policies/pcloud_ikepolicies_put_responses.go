// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudIkepoliciesPutReader is a Reader for the PcloudIkepoliciesPut structure.
type PcloudIkepoliciesPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudIkepoliciesPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudIkepoliciesPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudIkepoliciesPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudIkepoliciesPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudIkepoliciesPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudIkepoliciesPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudIkepoliciesPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudIkepoliciesPutOK creates a PcloudIkepoliciesPutOK with default headers values
func NewPcloudIkepoliciesPutOK() *PcloudIkepoliciesPutOK {
	return &PcloudIkepoliciesPutOK{}
}

/* PcloudIkepoliciesPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudIkepoliciesPutOK struct {
	Payload *models.IKEPolicy
}

func (o *PcloudIkepoliciesPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesPutOK  %+v", 200, o.Payload)
}
func (o *PcloudIkepoliciesPutOK) GetPayload() *models.IKEPolicy {
	return o.Payload
}

func (o *PcloudIkepoliciesPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IKEPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPutBadRequest creates a PcloudIkepoliciesPutBadRequest with default headers values
func NewPcloudIkepoliciesPutBadRequest() *PcloudIkepoliciesPutBadRequest {
	return &PcloudIkepoliciesPutBadRequest{}
}

/* PcloudIkepoliciesPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudIkepoliciesPutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesPutBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudIkepoliciesPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPutUnauthorized creates a PcloudIkepoliciesPutUnauthorized with default headers values
func NewPcloudIkepoliciesPutUnauthorized() *PcloudIkepoliciesPutUnauthorized {
	return &PcloudIkepoliciesPutUnauthorized{}
}

/* PcloudIkepoliciesPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudIkepoliciesPutUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesPutUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudIkepoliciesPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPutForbidden creates a PcloudIkepoliciesPutForbidden with default headers values
func NewPcloudIkepoliciesPutForbidden() *PcloudIkepoliciesPutForbidden {
	return &PcloudIkepoliciesPutForbidden{}
}

/* PcloudIkepoliciesPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudIkepoliciesPutForbidden struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesPutForbidden  %+v", 403, o.Payload)
}
func (o *PcloudIkepoliciesPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPutUnprocessableEntity creates a PcloudIkepoliciesPutUnprocessableEntity with default headers values
func NewPcloudIkepoliciesPutUnprocessableEntity() *PcloudIkepoliciesPutUnprocessableEntity {
	return &PcloudIkepoliciesPutUnprocessableEntity{}
}

/* PcloudIkepoliciesPutUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudIkepoliciesPutUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesPutUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudIkepoliciesPutUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudIkepoliciesPutInternalServerError creates a PcloudIkepoliciesPutInternalServerError with default headers values
func NewPcloudIkepoliciesPutInternalServerError() *PcloudIkepoliciesPutInternalServerError {
	return &PcloudIkepoliciesPutInternalServerError{}
}

/* PcloudIkepoliciesPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudIkepoliciesPutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudIkepoliciesPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/ike-policies/{ike_policy_id}][%d] pcloudIkepoliciesPutInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudIkepoliciesPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudIkepoliciesPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
