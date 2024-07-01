// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

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

// PcloudCloudconnectionsNetworksDeleteReader is a Reader for the PcloudCloudconnectionsNetworksDelete structure.
type PcloudCloudconnectionsNetworksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsNetworksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsNetworksDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudCloudconnectionsNetworksDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsNetworksDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsNetworksDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudconnectionsNetworksDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsNetworksDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsNetworksDeleteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudCloudconnectionsNetworksDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsNetworksDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}] pcloud.cloudconnections.networks.delete", response, response.Code())
	}
}

// NewPcloudCloudconnectionsNetworksDeleteOK creates a PcloudCloudconnectionsNetworksDeleteOK with default headers values
func NewPcloudCloudconnectionsNetworksDeleteOK() *PcloudCloudconnectionsNetworksDeleteOK {
	return &PcloudCloudconnectionsNetworksDeleteOK{}
}

/*
PcloudCloudconnectionsNetworksDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsNetworksDeleteOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud cloudconnections networks delete o k response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections networks delete o k response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks delete o k response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections networks delete o k response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks delete o k response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud cloudconnections networks delete o k response
func (o *PcloudCloudconnectionsNetworksDeleteOK) Code() int {
	return 200
}

func (o *PcloudCloudconnectionsNetworksDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteOK %s", 200, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteOK %s", 200, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteAccepted creates a PcloudCloudconnectionsNetworksDeleteAccepted with default headers values
func NewPcloudCloudconnectionsNetworksDeleteAccepted() *PcloudCloudconnectionsNetworksDeleteAccepted {
	return &PcloudCloudconnectionsNetworksDeleteAccepted{}
}

/*
PcloudCloudconnectionsNetworksDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudconnectionsNetworksDeleteAccepted struct {
	Payload *models.JobReference
}

// IsSuccess returns true when this pcloud cloudconnections networks delete accepted response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudconnections networks delete accepted response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks delete accepted response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections networks delete accepted response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks delete accepted response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the pcloud cloudconnections networks delete accepted response
func (o *PcloudCloudconnectionsNetworksDeleteAccepted) Code() int {
	return 202
}

func (o *PcloudCloudconnectionsNetworksDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteAccepted %s", 202, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteAccepted %s", 202, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteBadRequest creates a PcloudCloudconnectionsNetworksDeleteBadRequest with default headers values
func NewPcloudCloudconnectionsNetworksDeleteBadRequest() *PcloudCloudconnectionsNetworksDeleteBadRequest {
	return &PcloudCloudconnectionsNetworksDeleteBadRequest{}
}

/*
PcloudCloudconnectionsNetworksDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsNetworksDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks delete bad request response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks delete bad request response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks delete bad request response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks delete bad request response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks delete bad request response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud cloudconnections networks delete bad request response
func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) Code() int {
	return 400
}

func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteBadRequest %s", 400, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteBadRequest %s", 400, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteUnauthorized creates a PcloudCloudconnectionsNetworksDeleteUnauthorized with default headers values
func NewPcloudCloudconnectionsNetworksDeleteUnauthorized() *PcloudCloudconnectionsNetworksDeleteUnauthorized {
	return &PcloudCloudconnectionsNetworksDeleteUnauthorized{}
}

/*
PcloudCloudconnectionsNetworksDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsNetworksDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks delete unauthorized response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks delete unauthorized response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks delete unauthorized response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks delete unauthorized response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks delete unauthorized response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud cloudconnections networks delete unauthorized response
func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) Code() int {
	return 401
}

func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteUnauthorized %s", 401, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteForbidden creates a PcloudCloudconnectionsNetworksDeleteForbidden with default headers values
func NewPcloudCloudconnectionsNetworksDeleteForbidden() *PcloudCloudconnectionsNetworksDeleteForbidden {
	return &PcloudCloudconnectionsNetworksDeleteForbidden{}
}

/*
PcloudCloudconnectionsNetworksDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudconnectionsNetworksDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks delete forbidden response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks delete forbidden response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks delete forbidden response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks delete forbidden response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks delete forbidden response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud cloudconnections networks delete forbidden response
func (o *PcloudCloudconnectionsNetworksDeleteForbidden) Code() int {
	return 403
}

func (o *PcloudCloudconnectionsNetworksDeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteForbidden %s", 403, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteForbidden %s", 403, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteNotFound creates a PcloudCloudconnectionsNetworksDeleteNotFound with default headers values
func NewPcloudCloudconnectionsNetworksDeleteNotFound() *PcloudCloudconnectionsNetworksDeleteNotFound {
	return &PcloudCloudconnectionsNetworksDeleteNotFound{}
}

/*
PcloudCloudconnectionsNetworksDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsNetworksDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks delete not found response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks delete not found response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks delete not found response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks delete not found response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks delete not found response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud cloudconnections networks delete not found response
func (o *PcloudCloudconnectionsNetworksDeleteNotFound) Code() int {
	return 404
}

func (o *PcloudCloudconnectionsNetworksDeleteNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteNotFound %s", 404, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteNotFound %s", 404, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteRequestTimeout creates a PcloudCloudconnectionsNetworksDeleteRequestTimeout with default headers values
func NewPcloudCloudconnectionsNetworksDeleteRequestTimeout() *PcloudCloudconnectionsNetworksDeleteRequestTimeout {
	return &PcloudCloudconnectionsNetworksDeleteRequestTimeout{}
}

/*
PcloudCloudconnectionsNetworksDeleteRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsNetworksDeleteRequestTimeout struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks delete request timeout response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks delete request timeout response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks delete request timeout response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks delete request timeout response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks delete request timeout response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) IsCode(code int) bool {
	return code == 408
}

// Code gets the status code for the pcloud cloudconnections networks delete request timeout response
func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) Code() int {
	return 408
}

func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteRequestTimeout %s", 408, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteRequestTimeout %s", 408, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteGone creates a PcloudCloudconnectionsNetworksDeleteGone with default headers values
func NewPcloudCloudconnectionsNetworksDeleteGone() *PcloudCloudconnectionsNetworksDeleteGone {
	return &PcloudCloudconnectionsNetworksDeleteGone{}
}

/*
PcloudCloudconnectionsNetworksDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudCloudconnectionsNetworksDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks delete gone response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks delete gone response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks delete gone response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudconnections networks delete gone response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudconnections networks delete gone response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the pcloud cloudconnections networks delete gone response
func (o *PcloudCloudconnectionsNetworksDeleteGone) Code() int {
	return 410
}

func (o *PcloudCloudconnectionsNetworksDeleteGone) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteGone %s", 410, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteGone) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteGone %s", 410, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteInternalServerError creates a PcloudCloudconnectionsNetworksDeleteInternalServerError with default headers values
func NewPcloudCloudconnectionsNetworksDeleteInternalServerError() *PcloudCloudconnectionsNetworksDeleteInternalServerError {
	return &PcloudCloudconnectionsNetworksDeleteInternalServerError{}
}

/*
PcloudCloudconnectionsNetworksDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsNetworksDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudconnections networks delete internal server error response has a 2xx status code
func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudconnections networks delete internal server error response has a 3xx status code
func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudconnections networks delete internal server error response has a 4xx status code
func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudconnections networks delete internal server error response has a 5xx status code
func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudconnections networks delete internal server error response a status code equal to that given
func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud cloudconnections networks delete internal server error response
func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) Code() int {
	return 500
}

func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteInternalServerError %s", 500, payload)
}

func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
