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

// PcloudV2VolumescloneCancelPostReader is a Reader for the PcloudV2VolumescloneCancelPost structure.
type PcloudV2VolumescloneCancelPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudV2VolumescloneCancelPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPcloudV2VolumescloneCancelPostAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudV2VolumescloneCancelPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudV2VolumescloneCancelPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudV2VolumescloneCancelPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudV2VolumescloneCancelPostNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudV2VolumescloneCancelPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel] pcloud.v2.volumesclone.cancel.post", response, response.Code())
	}
}

// NewPcloudV2VolumescloneCancelPostAccepted creates a PcloudV2VolumescloneCancelPostAccepted with default headers values
func NewPcloudV2VolumescloneCancelPostAccepted() *PcloudV2VolumescloneCancelPostAccepted {
	return &PcloudV2VolumescloneCancelPostAccepted{}
}

/*
PcloudV2VolumescloneCancelPostAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudV2VolumescloneCancelPostAccepted struct {
	Payload *models.VolumesClone
}

// IsSuccess returns true when this pcloud v2 volumesclone cancel post accepted response has a 2xx status code
func (o *PcloudV2VolumescloneCancelPostAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud v2 volumesclone cancel post accepted response has a 3xx status code
func (o *PcloudV2VolumescloneCancelPostAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone cancel post accepted response has a 4xx status code
func (o *PcloudV2VolumescloneCancelPostAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone cancel post accepted response has a 5xx status code
func (o *PcloudV2VolumescloneCancelPostAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone cancel post accepted response a status code equal to that given
func (o *PcloudV2VolumescloneCancelPostAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud v2 volumesclone cancel post accepted response
func (o *PcloudV2VolumescloneCancelPostAccepted) Code() int {
	return 202
}

func (o *PcloudV2VolumescloneCancelPostAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostAccepted %s", 202, payload)
}

func (o *PcloudV2VolumescloneCancelPostAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostAccepted %s", 202, payload)
}

func (o *PcloudV2VolumescloneCancelPostAccepted) GetPayload() *models.VolumesClone {
	return o.Payload
}

func (o *PcloudV2VolumescloneCancelPostAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumesClone)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneCancelPostBadRequest creates a PcloudV2VolumescloneCancelPostBadRequest with default headers values
func NewPcloudV2VolumescloneCancelPostBadRequest() *PcloudV2VolumescloneCancelPostBadRequest {
	return &PcloudV2VolumescloneCancelPostBadRequest{}
}

/*
PcloudV2VolumescloneCancelPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudV2VolumescloneCancelPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone cancel post bad request response has a 2xx status code
func (o *PcloudV2VolumescloneCancelPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone cancel post bad request response has a 3xx status code
func (o *PcloudV2VolumescloneCancelPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone cancel post bad request response has a 4xx status code
func (o *PcloudV2VolumescloneCancelPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone cancel post bad request response has a 5xx status code
func (o *PcloudV2VolumescloneCancelPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone cancel post bad request response a status code equal to that given
func (o *PcloudV2VolumescloneCancelPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud v2 volumesclone cancel post bad request response
func (o *PcloudV2VolumescloneCancelPostBadRequest) Code() int {
	return 400
}

func (o *PcloudV2VolumescloneCancelPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostBadRequest %s", 400, payload)
}

func (o *PcloudV2VolumescloneCancelPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostBadRequest %s", 400, payload)
}

func (o *PcloudV2VolumescloneCancelPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneCancelPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneCancelPostUnauthorized creates a PcloudV2VolumescloneCancelPostUnauthorized with default headers values
func NewPcloudV2VolumescloneCancelPostUnauthorized() *PcloudV2VolumescloneCancelPostUnauthorized {
	return &PcloudV2VolumescloneCancelPostUnauthorized{}
}

/*
PcloudV2VolumescloneCancelPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudV2VolumescloneCancelPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone cancel post unauthorized response has a 2xx status code
func (o *PcloudV2VolumescloneCancelPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone cancel post unauthorized response has a 3xx status code
func (o *PcloudV2VolumescloneCancelPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone cancel post unauthorized response has a 4xx status code
func (o *PcloudV2VolumescloneCancelPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone cancel post unauthorized response has a 5xx status code
func (o *PcloudV2VolumescloneCancelPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone cancel post unauthorized response a status code equal to that given
func (o *PcloudV2VolumescloneCancelPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud v2 volumesclone cancel post unauthorized response
func (o *PcloudV2VolumescloneCancelPostUnauthorized) Code() int {
	return 401
}

func (o *PcloudV2VolumescloneCancelPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostUnauthorized %s", 401, payload)
}

func (o *PcloudV2VolumescloneCancelPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostUnauthorized %s", 401, payload)
}

func (o *PcloudV2VolumescloneCancelPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneCancelPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneCancelPostForbidden creates a PcloudV2VolumescloneCancelPostForbidden with default headers values
func NewPcloudV2VolumescloneCancelPostForbidden() *PcloudV2VolumescloneCancelPostForbidden {
	return &PcloudV2VolumescloneCancelPostForbidden{}
}

/*
PcloudV2VolumescloneCancelPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudV2VolumescloneCancelPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone cancel post forbidden response has a 2xx status code
func (o *PcloudV2VolumescloneCancelPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone cancel post forbidden response has a 3xx status code
func (o *PcloudV2VolumescloneCancelPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone cancel post forbidden response has a 4xx status code
func (o *PcloudV2VolumescloneCancelPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone cancel post forbidden response has a 5xx status code
func (o *PcloudV2VolumescloneCancelPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone cancel post forbidden response a status code equal to that given
func (o *PcloudV2VolumescloneCancelPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud v2 volumesclone cancel post forbidden response
func (o *PcloudV2VolumescloneCancelPostForbidden) Code() int {
	return 403
}

func (o *PcloudV2VolumescloneCancelPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostForbidden %s", 403, payload)
}

func (o *PcloudV2VolumescloneCancelPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostForbidden %s", 403, payload)
}

func (o *PcloudV2VolumescloneCancelPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneCancelPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneCancelPostNotFound creates a PcloudV2VolumescloneCancelPostNotFound with default headers values
func NewPcloudV2VolumescloneCancelPostNotFound() *PcloudV2VolumescloneCancelPostNotFound {
	return &PcloudV2VolumescloneCancelPostNotFound{}
}

/*
PcloudV2VolumescloneCancelPostNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudV2VolumescloneCancelPostNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone cancel post not found response has a 2xx status code
func (o *PcloudV2VolumescloneCancelPostNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone cancel post not found response has a 3xx status code
func (o *PcloudV2VolumescloneCancelPostNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone cancel post not found response has a 4xx status code
func (o *PcloudV2VolumescloneCancelPostNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud v2 volumesclone cancel post not found response has a 5xx status code
func (o *PcloudV2VolumescloneCancelPostNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud v2 volumesclone cancel post not found response a status code equal to that given
func (o *PcloudV2VolumescloneCancelPostNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud v2 volumesclone cancel post not found response
func (o *PcloudV2VolumescloneCancelPostNotFound) Code() int {
	return 404
}

func (o *PcloudV2VolumescloneCancelPostNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostNotFound %s", 404, payload)
}

func (o *PcloudV2VolumescloneCancelPostNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostNotFound %s", 404, payload)
}

func (o *PcloudV2VolumescloneCancelPostNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneCancelPostNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudV2VolumescloneCancelPostInternalServerError creates a PcloudV2VolumescloneCancelPostInternalServerError with default headers values
func NewPcloudV2VolumescloneCancelPostInternalServerError() *PcloudV2VolumescloneCancelPostInternalServerError {
	return &PcloudV2VolumescloneCancelPostInternalServerError{}
}

/*
PcloudV2VolumescloneCancelPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudV2VolumescloneCancelPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud v2 volumesclone cancel post internal server error response has a 2xx status code
func (o *PcloudV2VolumescloneCancelPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud v2 volumesclone cancel post internal server error response has a 3xx status code
func (o *PcloudV2VolumescloneCancelPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud v2 volumesclone cancel post internal server error response has a 4xx status code
func (o *PcloudV2VolumescloneCancelPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud v2 volumesclone cancel post internal server error response has a 5xx status code
func (o *PcloudV2VolumescloneCancelPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud v2 volumesclone cancel post internal server error response a status code equal to that given
func (o *PcloudV2VolumescloneCancelPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud v2 volumesclone cancel post internal server error response
func (o *PcloudV2VolumescloneCancelPostInternalServerError) Code() int {
	return 500
}

func (o *PcloudV2VolumescloneCancelPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostInternalServerError %s", 500, payload)
}

func (o *PcloudV2VolumescloneCancelPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /pcloud/v2/cloud-instances/{cloud_instance_id}/volumes-clone/{volumes_clone_id}/cancel][%d] pcloudV2VolumescloneCancelPostInternalServerError %s", 500, payload)
}

func (o *PcloudV2VolumescloneCancelPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudV2VolumescloneCancelPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
