// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVpnconnectionsPostReader is a Reader for the PcloudVpnconnectionsPost structure.
type PcloudVpnconnectionsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudVpnconnectionsPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudVpnconnectionsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudVpnconnectionsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudVpnconnectionsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudVpnconnectionsPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudVpnconnectionsPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudVpnconnectionsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudVpnconnectionsPostAccepted creates a PcloudVpnconnectionsPostAccepted with default headers values
func NewPcloudVpnconnectionsPostAccepted() *PcloudVpnconnectionsPostAccepted {
	return &PcloudVpnconnectionsPostAccepted{}
}

/* PcloudVpnconnectionsPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudVpnconnectionsPostAccepted struct {
	Payload *models.VPNConnectionCreateResponse
}

func (o *PcloudVpnconnectionsPostAccepted) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections][%d] pcloudVpnconnectionsPostAccepted  %+v", 202, o.Payload)
}
func (o *PcloudVpnconnectionsPostAccepted) GetPayload() *models.VPNConnectionCreateResponse {
	return o.Payload
}

func (o *PcloudVpnconnectionsPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VPNConnectionCreateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPostBadRequest creates a PcloudVpnconnectionsPostBadRequest with default headers values
func NewPcloudVpnconnectionsPostBadRequest() *PcloudVpnconnectionsPostBadRequest {
	return &PcloudVpnconnectionsPostBadRequest{}
}

/* PcloudVpnconnectionsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudVpnconnectionsPostBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPostBadRequest) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections][%d] pcloudVpnconnectionsPostBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudVpnconnectionsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPostUnauthorized creates a PcloudVpnconnectionsPostUnauthorized with default headers values
func NewPcloudVpnconnectionsPostUnauthorized() *PcloudVpnconnectionsPostUnauthorized {
	return &PcloudVpnconnectionsPostUnauthorized{}
}

/* PcloudVpnconnectionsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudVpnconnectionsPostUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPostUnauthorized) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections][%d] pcloudVpnconnectionsPostUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudVpnconnectionsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPostForbidden creates a PcloudVpnconnectionsPostForbidden with default headers values
func NewPcloudVpnconnectionsPostForbidden() *PcloudVpnconnectionsPostForbidden {
	return &PcloudVpnconnectionsPostForbidden{}
}

/* PcloudVpnconnectionsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudVpnconnectionsPostForbidden struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPostForbidden) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections][%d] pcloudVpnconnectionsPostForbidden  %+v", 403, o.Payload)
}
func (o *PcloudVpnconnectionsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPostConflict creates a PcloudVpnconnectionsPostConflict with default headers values
func NewPcloudVpnconnectionsPostConflict() *PcloudVpnconnectionsPostConflict {
	return &PcloudVpnconnectionsPostConflict{}
}

/* PcloudVpnconnectionsPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudVpnconnectionsPostConflict struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPostConflict) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections][%d] pcloudVpnconnectionsPostConflict  %+v", 409, o.Payload)
}
func (o *PcloudVpnconnectionsPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPostUnprocessableEntity creates a PcloudVpnconnectionsPostUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsPostUnprocessableEntity() *PcloudVpnconnectionsPostUnprocessableEntity {
	return &PcloudVpnconnectionsPostUnprocessableEntity{}
}

/* PcloudVpnconnectionsPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsPostUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPostUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections][%d] pcloudVpnconnectionsPostUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *PcloudVpnconnectionsPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPostInternalServerError creates a PcloudVpnconnectionsPostInternalServerError with default headers values
func NewPcloudVpnconnectionsPostInternalServerError() *PcloudVpnconnectionsPostInternalServerError {
	return &PcloudVpnconnectionsPostInternalServerError{}
}

/* PcloudVpnconnectionsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsPostInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPostInternalServerError) Error() string {
	return fmt.Sprintf("[POST /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn/vpn-connections][%d] pcloudVpnconnectionsPostInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudVpnconnectionsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudVpnconnectionsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
