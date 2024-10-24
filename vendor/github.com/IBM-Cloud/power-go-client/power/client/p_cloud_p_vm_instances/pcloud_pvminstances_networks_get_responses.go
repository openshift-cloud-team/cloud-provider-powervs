// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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

// PcloudPvminstancesNetworksGetReader is a Reader for the PcloudPvminstancesNetworksGet structure.
type PcloudPvminstancesNetworksGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesNetworksGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesNetworksGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesNetworksGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesNetworksGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesNetworksGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesNetworksGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesNetworksGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}] pcloud.pvminstances.networks.get", response, response.Code())
	}
}

// NewPcloudPvminstancesNetworksGetOK creates a PcloudPvminstancesNetworksGetOK with default headers values
func NewPcloudPvminstancesNetworksGetOK() *PcloudPvminstancesNetworksGetOK {
	return &PcloudPvminstancesNetworksGetOK{}
}

/*
PcloudPvminstancesNetworksGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesNetworksGetOK struct {
	Payload *models.PVMInstanceNetworks
}

// IsSuccess returns true when this pcloud pvminstances networks get o k response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances networks get o k response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks get o k response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances networks get o k response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks get o k response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances networks get o k response
func (o *PcloudPvminstancesNetworksGetOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesNetworksGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetOK %s", 200, payload)
}

func (o *PcloudPvminstancesNetworksGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetOK %s", 200, payload)
}

func (o *PcloudPvminstancesNetworksGetOK) GetPayload() *models.PVMInstanceNetworks {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PVMInstanceNetworks)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetBadRequest creates a PcloudPvminstancesNetworksGetBadRequest with default headers values
func NewPcloudPvminstancesNetworksGetBadRequest() *PcloudPvminstancesNetworksGetBadRequest {
	return &PcloudPvminstancesNetworksGetBadRequest{}
}

/*
PcloudPvminstancesNetworksGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesNetworksGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks get bad request response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks get bad request response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks get bad request response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks get bad request response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks get bad request response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances networks get bad request response
func (o *PcloudPvminstancesNetworksGetBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesNetworksGetBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesNetworksGetBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesNetworksGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetUnauthorized creates a PcloudPvminstancesNetworksGetUnauthorized with default headers values
func NewPcloudPvminstancesNetworksGetUnauthorized() *PcloudPvminstancesNetworksGetUnauthorized {
	return &PcloudPvminstancesNetworksGetUnauthorized{}
}

/*
PcloudPvminstancesNetworksGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesNetworksGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks get unauthorized response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks get unauthorized response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks get unauthorized response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks get unauthorized response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks get unauthorized response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances networks get unauthorized response
func (o *PcloudPvminstancesNetworksGetUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesNetworksGetUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesNetworksGetUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesNetworksGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetForbidden creates a PcloudPvminstancesNetworksGetForbidden with default headers values
func NewPcloudPvminstancesNetworksGetForbidden() *PcloudPvminstancesNetworksGetForbidden {
	return &PcloudPvminstancesNetworksGetForbidden{}
}

/*
PcloudPvminstancesNetworksGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesNetworksGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks get forbidden response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks get forbidden response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks get forbidden response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks get forbidden response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks get forbidden response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances networks get forbidden response
func (o *PcloudPvminstancesNetworksGetForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesNetworksGetForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesNetworksGetForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesNetworksGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetNotFound creates a PcloudPvminstancesNetworksGetNotFound with default headers values
func NewPcloudPvminstancesNetworksGetNotFound() *PcloudPvminstancesNetworksGetNotFound {
	return &PcloudPvminstancesNetworksGetNotFound{}
}

/*
PcloudPvminstancesNetworksGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesNetworksGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks get not found response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks get not found response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks get not found response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances networks get not found response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances networks get not found response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances networks get not found response
func (o *PcloudPvminstancesNetworksGetNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesNetworksGetNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesNetworksGetNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesNetworksGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesNetworksGetInternalServerError creates a PcloudPvminstancesNetworksGetInternalServerError with default headers values
func NewPcloudPvminstancesNetworksGetInternalServerError() *PcloudPvminstancesNetworksGetInternalServerError {
	return &PcloudPvminstancesNetworksGetInternalServerError{}
}

/*
PcloudPvminstancesNetworksGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesNetworksGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances networks get internal server error response has a 2xx status code
func (o *PcloudPvminstancesNetworksGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances networks get internal server error response has a 3xx status code
func (o *PcloudPvminstancesNetworksGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances networks get internal server error response has a 4xx status code
func (o *PcloudPvminstancesNetworksGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances networks get internal server error response has a 5xx status code
func (o *PcloudPvminstancesNetworksGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances networks get internal server error response a status code equal to that given
func (o *PcloudPvminstancesNetworksGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances networks get internal server error response
func (o *PcloudPvminstancesNetworksGetInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesNetworksGetInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesNetworksGetInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/networks/{network_id}][%d] pcloudPvminstancesNetworksGetInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesNetworksGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesNetworksGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
