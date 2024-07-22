// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudV2VolumesPostReader is a Reader for the PcloudV2VolumesPost structure.
type PcloudV2VolumesPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumesPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPcloudV2VolumesPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV2VolumesPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV2VolumesPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudV2VolumesPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudV2VolumesPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudV2VolumesPostConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewPcloudV2VolumesPostUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2VolumesPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes] pcloud.v2.volumes.post", response, response.Code())
	}
}

// NewPcloudV2VolumesPostCreated creates a PcloudV2VolumesPostCreated with default headers values
func NewPcloudV2VolumesPostCreated() *PcloudV2VolumesPostCreated {
	return &PcloudV2VolumesPostCreated{}
}

/*
PcloudV2VolumesPostCreated describes a response with status code 201, with default header values.

Created
*/
type PcloudV2VolumesPostCreated struct {
	Payload *models.Volumes
}

// IsSuccess returns true when this pcloud v2 volumes post created response has a 2xx status code
func (o *PcloudV2VolumesPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud v2 volumes post created response has a 3xx status code
func (o *PcloudV2VolumesPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumes post created response has a 4xx status code
func (o *PcloudV2VolumesPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumes post created response has a 5xx status code
func (o *PcloudV2VolumesPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumes post created response a status code equal to that given
func (o *PcloudV2VolumesPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the pcloud v2 volumes post created response
func (o *PcloudV2VolumesPostCreated) Code() int {
	return 201
}

func (o *PcloudV2VolumesPostCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostCreated %s", 201, payload)
}

func (o *PcloudV2VolumesPostCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostCreated %s", 201, payload)
}

func (o *PcloudV2VolumesPostCreated) GetPayload() *models.Volumes {
	return o.Payload
}

func (o *PcloudV2VolumesPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volumes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesPostBadRequest creates a PcloudV2VolumesPostBadRequest with default headers values
func NewPcloudV2VolumesPostBadRequest() *PcloudV2VolumesPostBadRequest {
	return &PcloudV2VolumesPostBadRequest{}
}

/*
PcloudV2VolumesPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV2VolumesPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumes post bad request response has a 2xx status code
func (o *PcloudV2VolumesPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumes post bad request response has a 3xx status code
func (o *PcloudV2VolumesPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumes post bad request response has a 4xx status code
func (o *PcloudV2VolumesPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumes post bad request response has a 5xx status code
func (o *PcloudV2VolumesPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumes post bad request response a status code equal to that given
func (o *PcloudV2VolumesPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud v2 volumes post bad request response
func (o *PcloudV2VolumesPostBadRequest) Code() int {
	return 400
}

func (o *PcloudV2VolumesPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostBadRequest %s", 400, payload)
}

func (o *PcloudV2VolumesPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostBadRequest %s", 400, payload)
}

func (o *PcloudV2VolumesPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesPostUnauthorized creates a PcloudV2VolumesPostUnauthorized with default headers values
func NewPcloudV2VolumesPostUnauthorized() *PcloudV2VolumesPostUnauthorized {
	return &PcloudV2VolumesPostUnauthorized{}
}

/*
PcloudV2VolumesPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2VolumesPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumes post unauthorized response has a 2xx status code
func (o *PcloudV2VolumesPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumes post unauthorized response has a 3xx status code
func (o *PcloudV2VolumesPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumes post unauthorized response has a 4xx status code
func (o *PcloudV2VolumesPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumes post unauthorized response has a 5xx status code
func (o *PcloudV2VolumesPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumes post unauthorized response a status code equal to that given
func (o *PcloudV2VolumesPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud v2 volumes post unauthorized response
func (o *PcloudV2VolumesPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudV2VolumesPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostUnauthorized %s", 401, payload)
}

func (o *PcloudV2VolumesPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostUnauthorized %s", 401, payload)
}

func (o *PcloudV2VolumesPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesPostForbidden creates a PcloudV2VolumesPostForbidden with default headers values
func NewPcloudV2VolumesPostForbidden() *PcloudV2VolumesPostForbidden {
	return &PcloudV2VolumesPostForbidden{}
}

/*
PcloudV2VolumesPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudV2VolumesPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumes post forbidden response has a 2xx status code
func (o *PcloudV2VolumesPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumes post forbidden response has a 3xx status code
func (o *PcloudV2VolumesPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumes post forbidden response has a 4xx status code
func (o *PcloudV2VolumesPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumes post forbidden response has a 5xx status code
func (o *PcloudV2VolumesPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumes post forbidden response a status code equal to that given
func (o *PcloudV2VolumesPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud v2 volumes post forbidden response
func (o *PcloudV2VolumesPostForbidden) Code() int {
	return 403
}

func (o *PcloudV2VolumesPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostForbidden %s", 403, payload)
}

func (o *PcloudV2VolumesPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostForbidden %s", 403, payload)
}

func (o *PcloudV2VolumesPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesPostNotFound creates a PcloudV2VolumesPostNotFound with default headers values
func NewPcloudV2VolumesPostNotFound() *PcloudV2VolumesPostNotFound {
	return &PcloudV2VolumesPostNotFound{}
}

/*
PcloudV2VolumesPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudV2VolumesPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumes post not found response has a 2xx status code
func (o *PcloudV2VolumesPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumes post not found response has a 3xx status code
func (o *PcloudV2VolumesPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumes post not found response has a 4xx status code
func (o *PcloudV2VolumesPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumes post not found response has a 5xx status code
func (o *PcloudV2VolumesPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumes post not found response a status code equal to that given
func (o *PcloudV2VolumesPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud v2 volumes post not found response
func (o *PcloudV2VolumesPostNotFound) Code() int {
	return 404
}

func (o *PcloudV2VolumesPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostNotFound %s", 404, payload)
}

func (o *PcloudV2VolumesPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostNotFound %s", 404, payload)
}

func (o *PcloudV2VolumesPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesPostConflict creates a PcloudV2VolumesPostConflict with default headers values
func NewPcloudV2VolumesPostConflict() *PcloudV2VolumesPostConflict {
	return &PcloudV2VolumesPostConflict{}
}

/*
PcloudV2VolumesPostConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudV2VolumesPostConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumes post conflict response has a 2xx status code
func (o *PcloudV2VolumesPostConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumes post conflict response has a 3xx status code
func (o *PcloudV2VolumesPostConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumes post conflict response has a 4xx status code
func (o *PcloudV2VolumesPostConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumes post conflict response has a 5xx status code
func (o *PcloudV2VolumesPostConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumes post conflict response a status code equal to that given
func (o *PcloudV2VolumesPostConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud v2 volumes post conflict response
func (o *PcloudV2VolumesPostConflict) Code() int {
	return 409
}

func (o *PcloudV2VolumesPostConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostConflict %s", 409, payload)
}

func (o *PcloudV2VolumesPostConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostConflict %s", 409, payload)
}

func (o *PcloudV2VolumesPostConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesPostConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesPostUnprocessableEntity creates a PcloudV2VolumesPostUnprocessableEntity with default headers values
func NewPcloudV2VolumesPostUnprocessableEntity() *PcloudV2VolumesPostUnprocessableEntity {
	return &PcloudV2VolumesPostUnprocessableEntity{}
}

/*
PcloudV2VolumesPostUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type PcloudV2VolumesPostUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumes post unprocessable entity response has a 2xx status code
func (o *PcloudV2VolumesPostUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumes post unprocessable entity response has a 3xx status code
func (o *PcloudV2VolumesPostUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumes post unprocessable entity response has a 4xx status code
func (o *PcloudV2VolumesPostUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumes post unprocessable entity response has a 5xx status code
func (o *PcloudV2VolumesPostUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumes post unprocessable entity response a status code equal to that given
func (o *PcloudV2VolumesPostUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the pcloud v2 volumes post unprocessable entity response
func (o *PcloudV2VolumesPostUnprocessableEntity) Code() int {
	return 422
}

func (o *PcloudV2VolumesPostUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudV2VolumesPostUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostUnprocessableEntity %s", 422, payload)
}

func (o *PcloudV2VolumesPostUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesPostUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumesPostInternalServerError creates a PcloudV2VolumesPostInternalServerError with default headers values
func NewPcloudV2VolumesPostInternalServerError() *PcloudV2VolumesPostInternalServerError {
	return &PcloudV2VolumesPostInternalServerError{}
}

/*
PcloudV2VolumesPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2VolumesPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumes post internal server error response has a 2xx status code
func (o *PcloudV2VolumesPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumes post internal server error response has a 3xx status code
func (o *PcloudV2VolumesPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumes post internal server error response has a 4xx status code
func (o *PcloudV2VolumesPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumes post internal server error response has a 5xx status code
func (o *PcloudV2VolumesPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud v2 volumes post internal server error response a status code equal to that given
func (o *PcloudV2VolumesPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud v2 volumes post internal server error response
func (o *PcloudV2VolumesPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudV2VolumesPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostInternalServerError %s", 500, payload)
}

func (o *PcloudV2VolumesPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes][%d] pcloudV2VolumesPostInternalServerError %s", 500, payload)
}

func (o *PcloudV2VolumesPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumesPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
