// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_jobs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesJobsGetReader is a Reader for the PcloudCloudinstancesJobsGet structure.
type PcloudCloudinstancesJobsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesJobsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesJobsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesJobsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesJobsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesJobsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesJobsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesJobsGetOK creates a PcloudCloudinstancesJobsGetOK with default headers values
func NewPcloudCloudinstancesJobsGetOK() *PcloudCloudinstancesJobsGetOK {
	return &PcloudCloudinstancesJobsGetOK{}
}

/* PcloudCloudinstancesJobsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesJobsGetOK struct {
	Payload *models.Job
}

func (o *PcloudCloudinstancesJobsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetOK  %+v", 200, o.Payload)
}
func (o *PcloudCloudinstancesJobsGetOK) GetPayload() *models.Job {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Job)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetBadRequest creates a PcloudCloudinstancesJobsGetBadRequest with default headers values
func NewPcloudCloudinstancesJobsGetBadRequest() *PcloudCloudinstancesJobsGetBadRequest {
	return &PcloudCloudinstancesJobsGetBadRequest{}
}

/* PcloudCloudinstancesJobsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesJobsGetBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesJobsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudinstancesJobsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetUnauthorized creates a PcloudCloudinstancesJobsGetUnauthorized with default headers values
func NewPcloudCloudinstancesJobsGetUnauthorized() *PcloudCloudinstancesJobsGetUnauthorized {
	return &PcloudCloudinstancesJobsGetUnauthorized{}
}

/* PcloudCloudinstancesJobsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesJobsGetUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesJobsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudinstancesJobsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetNotFound creates a PcloudCloudinstancesJobsGetNotFound with default headers values
func NewPcloudCloudinstancesJobsGetNotFound() *PcloudCloudinstancesJobsGetNotFound {
	return &PcloudCloudinstancesJobsGetNotFound{}
}

/* PcloudCloudinstancesJobsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesJobsGetNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesJobsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetNotFound  %+v", 404, o.Payload)
}
func (o *PcloudCloudinstancesJobsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesJobsGetInternalServerError creates a PcloudCloudinstancesJobsGetInternalServerError with default headers values
func NewPcloudCloudinstancesJobsGetInternalServerError() *PcloudCloudinstancesJobsGetInternalServerError {
	return &PcloudCloudinstancesJobsGetInternalServerError{}
}

/* PcloudCloudinstancesJobsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesJobsGetInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudinstancesJobsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/jobs/{job_id}][%d] pcloudCloudinstancesJobsGetInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudinstancesJobsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesJobsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
