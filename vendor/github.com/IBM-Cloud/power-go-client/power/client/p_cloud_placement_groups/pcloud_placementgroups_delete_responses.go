// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

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

// PcloudPlacementgroupsDeleteReader is a Reader for the PcloudPlacementgroupsDelete structure.
type PcloudPlacementgroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPlacementgroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPlacementgroupsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPlacementgroupsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPlacementgroupsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPlacementgroupsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPlacementgroupsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPlacementgroupsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}] pcloud.placementgroups.delete", response, response.Code())
	}
}

// NewPcloudPlacementgroupsDeleteOK creates a PcloudPlacementgroupsDeleteOK with default headers values
func NewPcloudPlacementgroupsDeleteOK() *PcloudPlacementgroupsDeleteOK {
	return &PcloudPlacementgroupsDeleteOK{}
}

/*
PcloudPlacementgroupsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPlacementgroupsDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud placementgroups delete o k response has a 2xx status code
func (o *PcloudPlacementgroupsDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud placementgroups delete o k response has a 3xx status code
func (o *PcloudPlacementgroupsDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups delete o k response has a 4xx status code
func (o *PcloudPlacementgroupsDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups delete o k response has a 5xx status code
func (o *PcloudPlacementgroupsDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups delete o k response a status code equal to that given
func (o *PcloudPlacementgroupsDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud placementgroups delete o k response
func (o *PcloudPlacementgroupsDeleteOK) Code() int {
	return 200
}

func (o *PcloudPlacementgroupsDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteOK %s", 200, payload)
}

func (o *PcloudPlacementgroupsDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteOK %s", 200, payload)
}

func (o *PcloudPlacementgroupsDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPlacementgroupsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsDeleteBadRequest creates a PcloudPlacementgroupsDeleteBadRequest with default headers values
func NewPcloudPlacementgroupsDeleteBadRequest() *PcloudPlacementgroupsDeleteBadRequest {
	return &PcloudPlacementgroupsDeleteBadRequest{}
}

/*
PcloudPlacementgroupsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPlacementgroupsDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups delete bad request response has a 2xx status code
func (o *PcloudPlacementgroupsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups delete bad request response has a 3xx status code
func (o *PcloudPlacementgroupsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups delete bad request response has a 4xx status code
func (o *PcloudPlacementgroupsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups delete bad request response has a 5xx status code
func (o *PcloudPlacementgroupsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups delete bad request response a status code equal to that given
func (o *PcloudPlacementgroupsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud placementgroups delete bad request response
func (o *PcloudPlacementgroupsDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudPlacementgroupsDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteBadRequest %s", 400, payload)
}

func (o *PcloudPlacementgroupsDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteBadRequest %s", 400, payload)
}

func (o *PcloudPlacementgroupsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsDeleteUnauthorized creates a PcloudPlacementgroupsDeleteUnauthorized with default headers values
func NewPcloudPlacementgroupsDeleteUnauthorized() *PcloudPlacementgroupsDeleteUnauthorized {
	return &PcloudPlacementgroupsDeleteUnauthorized{}
}

/*
PcloudPlacementgroupsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPlacementgroupsDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups delete unauthorized response has a 2xx status code
func (o *PcloudPlacementgroupsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups delete unauthorized response has a 3xx status code
func (o *PcloudPlacementgroupsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups delete unauthorized response has a 4xx status code
func (o *PcloudPlacementgroupsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups delete unauthorized response has a 5xx status code
func (o *PcloudPlacementgroupsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups delete unauthorized response a status code equal to that given
func (o *PcloudPlacementgroupsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud placementgroups delete unauthorized response
func (o *PcloudPlacementgroupsDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudPlacementgroupsDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudPlacementgroupsDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudPlacementgroupsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsDeleteForbidden creates a PcloudPlacementgroupsDeleteForbidden with default headers values
func NewPcloudPlacementgroupsDeleteForbidden() *PcloudPlacementgroupsDeleteForbidden {
	return &PcloudPlacementgroupsDeleteForbidden{}
}

/*
PcloudPlacementgroupsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPlacementgroupsDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups delete forbidden response has a 2xx status code
func (o *PcloudPlacementgroupsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups delete forbidden response has a 3xx status code
func (o *PcloudPlacementgroupsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups delete forbidden response has a 4xx status code
func (o *PcloudPlacementgroupsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups delete forbidden response has a 5xx status code
func (o *PcloudPlacementgroupsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups delete forbidden response a status code equal to that given
func (o *PcloudPlacementgroupsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud placementgroups delete forbidden response
func (o *PcloudPlacementgroupsDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudPlacementgroupsDeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteForbidden %s", 403, payload)
}

func (o *PcloudPlacementgroupsDeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteForbidden %s", 403, payload)
}

func (o *PcloudPlacementgroupsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsDeleteNotFound creates a PcloudPlacementgroupsDeleteNotFound with default headers values
func NewPcloudPlacementgroupsDeleteNotFound() *PcloudPlacementgroupsDeleteNotFound {
	return &PcloudPlacementgroupsDeleteNotFound{}
}

/*
PcloudPlacementgroupsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPlacementgroupsDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups delete not found response has a 2xx status code
func (o *PcloudPlacementgroupsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups delete not found response has a 3xx status code
func (o *PcloudPlacementgroupsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups delete not found response has a 4xx status code
func (o *PcloudPlacementgroupsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud placementgroups delete not found response has a 5xx status code
func (o *PcloudPlacementgroupsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud placementgroups delete not found response a status code equal to that given
func (o *PcloudPlacementgroupsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud placementgroups delete not found response
func (o *PcloudPlacementgroupsDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudPlacementgroupsDeleteNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteNotFound %s", 404, payload)
}

func (o *PcloudPlacementgroupsDeleteNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteNotFound %s", 404, payload)
}

func (o *PcloudPlacementgroupsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPlacementgroupsDeleteInternalServerError creates a PcloudPlacementgroupsDeleteInternalServerError with default headers values
func NewPcloudPlacementgroupsDeleteInternalServerError() *PcloudPlacementgroupsDeleteInternalServerError {
	return &PcloudPlacementgroupsDeleteInternalServerError{}
}

/*
PcloudPlacementgroupsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPlacementgroupsDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud placementgroups delete internal server error response has a 2xx status code
func (o *PcloudPlacementgroupsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud placementgroups delete internal server error response has a 3xx status code
func (o *PcloudPlacementgroupsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud placementgroups delete internal server error response has a 4xx status code
func (o *PcloudPlacementgroupsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud placementgroups delete internal server error response has a 5xx status code
func (o *PcloudPlacementgroupsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud placementgroups delete internal server error response a status code equal to that given
func (o *PcloudPlacementgroupsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud placementgroups delete internal server error response
func (o *PcloudPlacementgroupsDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudPlacementgroupsDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudPlacementgroupsDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/placement-groups/{placement_group_id}][%d] pcloudPlacementgroupsDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudPlacementgroupsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPlacementgroupsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
