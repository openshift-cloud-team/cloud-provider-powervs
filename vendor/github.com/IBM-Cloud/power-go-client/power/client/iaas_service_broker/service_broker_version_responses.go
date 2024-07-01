// Code generated by go-swagger; DO NOT EDIT.

package iaas_service_broker

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

// ServiceBrokerVersionReader is a Reader for the ServiceBrokerVersion structure.
type ServiceBrokerVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBrokerVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBrokerVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /broker/v1/version] serviceBroker.version", response, response.Code())
	}
}

// NewServiceBrokerVersionOK creates a ServiceBrokerVersionOK with default headers values
func NewServiceBrokerVersionOK() *ServiceBrokerVersionOK {
	return &ServiceBrokerVersionOK{}
}

/*
ServiceBrokerVersionOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerVersionOK struct {
	Payload *models.Version
}

// IsSuccess returns true when this service broker version o k response has a 2xx status code
func (o *ServiceBrokerVersionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker version o k response has a 3xx status code
func (o *ServiceBrokerVersionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker version o k response has a 4xx status code
func (o *ServiceBrokerVersionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker version o k response has a 5xx status code
func (o *ServiceBrokerVersionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker version o k response a status code equal to that given
func (o *ServiceBrokerVersionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker version o k response
func (o *ServiceBrokerVersionOK) Code() int {
	return 200
}

func (o *ServiceBrokerVersionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionOK %s", 200, payload)
}

func (o *ServiceBrokerVersionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionOK %s", 200, payload)
}

func (o *ServiceBrokerVersionOK) GetPayload() *models.Version {
	return o.Payload
}

func (o *ServiceBrokerVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Version)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerVersionBadRequest creates a ServiceBrokerVersionBadRequest with default headers values
func NewServiceBrokerVersionBadRequest() *ServiceBrokerVersionBadRequest {
	return &ServiceBrokerVersionBadRequest{}
}

/*
ServiceBrokerVersionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerVersionBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker version bad request response has a 2xx status code
func (o *ServiceBrokerVersionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker version bad request response has a 3xx status code
func (o *ServiceBrokerVersionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker version bad request response has a 4xx status code
func (o *ServiceBrokerVersionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker version bad request response has a 5xx status code
func (o *ServiceBrokerVersionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker version bad request response a status code equal to that given
func (o *ServiceBrokerVersionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker version bad request response
func (o *ServiceBrokerVersionBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerVersionBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionBadRequest %s", 400, payload)
}

func (o *ServiceBrokerVersionBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionBadRequest %s", 400, payload)
}

func (o *ServiceBrokerVersionBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerVersionUnauthorized creates a ServiceBrokerVersionUnauthorized with default headers values
func NewServiceBrokerVersionUnauthorized() *ServiceBrokerVersionUnauthorized {
	return &ServiceBrokerVersionUnauthorized{}
}

/*
ServiceBrokerVersionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerVersionUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker version unauthorized response has a 2xx status code
func (o *ServiceBrokerVersionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker version unauthorized response has a 3xx status code
func (o *ServiceBrokerVersionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker version unauthorized response has a 4xx status code
func (o *ServiceBrokerVersionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker version unauthorized response has a 5xx status code
func (o *ServiceBrokerVersionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker version unauthorized response a status code equal to that given
func (o *ServiceBrokerVersionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service broker version unauthorized response
func (o *ServiceBrokerVersionUnauthorized) Code() int {
	return 401
}

func (o *ServiceBrokerVersionUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionUnauthorized %s", 401, payload)
}

func (o *ServiceBrokerVersionUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionUnauthorized %s", 401, payload)
}

func (o *ServiceBrokerVersionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerVersionForbidden creates a ServiceBrokerVersionForbidden with default headers values
func NewServiceBrokerVersionForbidden() *ServiceBrokerVersionForbidden {
	return &ServiceBrokerVersionForbidden{}
}

/*
ServiceBrokerVersionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBrokerVersionForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker version forbidden response has a 2xx status code
func (o *ServiceBrokerVersionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker version forbidden response has a 3xx status code
func (o *ServiceBrokerVersionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker version forbidden response has a 4xx status code
func (o *ServiceBrokerVersionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker version forbidden response has a 5xx status code
func (o *ServiceBrokerVersionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker version forbidden response a status code equal to that given
func (o *ServiceBrokerVersionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service broker version forbidden response
func (o *ServiceBrokerVersionForbidden) Code() int {
	return 403
}

func (o *ServiceBrokerVersionForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionForbidden %s", 403, payload)
}

func (o *ServiceBrokerVersionForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionForbidden %s", 403, payload)
}

func (o *ServiceBrokerVersionForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerVersionNotFound creates a ServiceBrokerVersionNotFound with default headers values
func NewServiceBrokerVersionNotFound() *ServiceBrokerVersionNotFound {
	return &ServiceBrokerVersionNotFound{}
}

/*
ServiceBrokerVersionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerVersionNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker version not found response has a 2xx status code
func (o *ServiceBrokerVersionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker version not found response has a 3xx status code
func (o *ServiceBrokerVersionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker version not found response has a 4xx status code
func (o *ServiceBrokerVersionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker version not found response has a 5xx status code
func (o *ServiceBrokerVersionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker version not found response a status code equal to that given
func (o *ServiceBrokerVersionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker version not found response
func (o *ServiceBrokerVersionNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerVersionNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionNotFound %s", 404, payload)
}

func (o *ServiceBrokerVersionNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /broker/v1/version][%d] serviceBrokerVersionNotFound %s", 404, payload)
}

func (o *ServiceBrokerVersionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
