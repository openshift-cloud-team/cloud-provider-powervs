// Code generated by go-swagger; DO NOT EDIT.

package iaas_service_broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// ServiceBrokerTestTimeoutReader is a Reader for the ServiceBrokerTestTimeout structure.
type ServiceBrokerTestTimeoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ServiceBrokerTestTimeoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewServiceBrokerTestTimeoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewServiceBrokerTestTimeoutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewServiceBrokerTestTimeoutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewServiceBrokerTestTimeoutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewServiceBrokerTestTimeoutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /broker/v1/test/timeout] serviceBroker.test.timeout", response, response.Code())
	}
}

// NewServiceBrokerTestTimeoutOK creates a ServiceBrokerTestTimeoutOK with default headers values
func NewServiceBrokerTestTimeoutOK() *ServiceBrokerTestTimeoutOK {
	return &ServiceBrokerTestTimeoutOK{}
}

/*
ServiceBrokerTestTimeoutOK describes a response with status code 200, with default header values.

OK
*/
type ServiceBrokerTestTimeoutOK struct {
	Payload models.Object
}

// IsSuccess returns true when this service broker test timeout o k response has a 2xx status code
func (o *ServiceBrokerTestTimeoutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this service broker test timeout o k response has a 3xx status code
func (o *ServiceBrokerTestTimeoutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker test timeout o k response has a 4xx status code
func (o *ServiceBrokerTestTimeoutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this service broker test timeout o k response has a 5xx status code
func (o *ServiceBrokerTestTimeoutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker test timeout o k response a status code equal to that given
func (o *ServiceBrokerTestTimeoutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the service broker test timeout o k response
func (o *ServiceBrokerTestTimeoutOK) Code() int {
	return 200
}

func (o *ServiceBrokerTestTimeoutOK) Error() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerTestTimeoutOK) String() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutOK  %+v", 200, o.Payload)
}

func (o *ServiceBrokerTestTimeoutOK) GetPayload() models.Object {
	return o.Payload
}

func (o *ServiceBrokerTestTimeoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerTestTimeoutBadRequest creates a ServiceBrokerTestTimeoutBadRequest with default headers values
func NewServiceBrokerTestTimeoutBadRequest() *ServiceBrokerTestTimeoutBadRequest {
	return &ServiceBrokerTestTimeoutBadRequest{}
}

/*
ServiceBrokerTestTimeoutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ServiceBrokerTestTimeoutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker test timeout bad request response has a 2xx status code
func (o *ServiceBrokerTestTimeoutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker test timeout bad request response has a 3xx status code
func (o *ServiceBrokerTestTimeoutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker test timeout bad request response has a 4xx status code
func (o *ServiceBrokerTestTimeoutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker test timeout bad request response has a 5xx status code
func (o *ServiceBrokerTestTimeoutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker test timeout bad request response a status code equal to that given
func (o *ServiceBrokerTestTimeoutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the service broker test timeout bad request response
func (o *ServiceBrokerTestTimeoutBadRequest) Code() int {
	return 400
}

func (o *ServiceBrokerTestTimeoutBadRequest) Error() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerTestTimeoutBadRequest) String() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutBadRequest  %+v", 400, o.Payload)
}

func (o *ServiceBrokerTestTimeoutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerTestTimeoutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerTestTimeoutUnauthorized creates a ServiceBrokerTestTimeoutUnauthorized with default headers values
func NewServiceBrokerTestTimeoutUnauthorized() *ServiceBrokerTestTimeoutUnauthorized {
	return &ServiceBrokerTestTimeoutUnauthorized{}
}

/*
ServiceBrokerTestTimeoutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ServiceBrokerTestTimeoutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker test timeout unauthorized response has a 2xx status code
func (o *ServiceBrokerTestTimeoutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker test timeout unauthorized response has a 3xx status code
func (o *ServiceBrokerTestTimeoutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker test timeout unauthorized response has a 4xx status code
func (o *ServiceBrokerTestTimeoutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker test timeout unauthorized response has a 5xx status code
func (o *ServiceBrokerTestTimeoutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker test timeout unauthorized response a status code equal to that given
func (o *ServiceBrokerTestTimeoutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the service broker test timeout unauthorized response
func (o *ServiceBrokerTestTimeoutUnauthorized) Code() int {
	return 401
}

func (o *ServiceBrokerTestTimeoutUnauthorized) Error() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerTestTimeoutUnauthorized) String() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutUnauthorized  %+v", 401, o.Payload)
}

func (o *ServiceBrokerTestTimeoutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerTestTimeoutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerTestTimeoutForbidden creates a ServiceBrokerTestTimeoutForbidden with default headers values
func NewServiceBrokerTestTimeoutForbidden() *ServiceBrokerTestTimeoutForbidden {
	return &ServiceBrokerTestTimeoutForbidden{}
}

/*
ServiceBrokerTestTimeoutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ServiceBrokerTestTimeoutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker test timeout forbidden response has a 2xx status code
func (o *ServiceBrokerTestTimeoutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker test timeout forbidden response has a 3xx status code
func (o *ServiceBrokerTestTimeoutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker test timeout forbidden response has a 4xx status code
func (o *ServiceBrokerTestTimeoutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker test timeout forbidden response has a 5xx status code
func (o *ServiceBrokerTestTimeoutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker test timeout forbidden response a status code equal to that given
func (o *ServiceBrokerTestTimeoutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the service broker test timeout forbidden response
func (o *ServiceBrokerTestTimeoutForbidden) Code() int {
	return 403
}

func (o *ServiceBrokerTestTimeoutForbidden) Error() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerTestTimeoutForbidden) String() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutForbidden  %+v", 403, o.Payload)
}

func (o *ServiceBrokerTestTimeoutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerTestTimeoutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewServiceBrokerTestTimeoutNotFound creates a ServiceBrokerTestTimeoutNotFound with default headers values
func NewServiceBrokerTestTimeoutNotFound() *ServiceBrokerTestTimeoutNotFound {
	return &ServiceBrokerTestTimeoutNotFound{}
}

/*
ServiceBrokerTestTimeoutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ServiceBrokerTestTimeoutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this service broker test timeout not found response has a 2xx status code
func (o *ServiceBrokerTestTimeoutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this service broker test timeout not found response has a 3xx status code
func (o *ServiceBrokerTestTimeoutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this service broker test timeout not found response has a 4xx status code
func (o *ServiceBrokerTestTimeoutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this service broker test timeout not found response has a 5xx status code
func (o *ServiceBrokerTestTimeoutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this service broker test timeout not found response a status code equal to that given
func (o *ServiceBrokerTestTimeoutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the service broker test timeout not found response
func (o *ServiceBrokerTestTimeoutNotFound) Code() int {
	return 404
}

func (o *ServiceBrokerTestTimeoutNotFound) Error() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerTestTimeoutNotFound) String() string {
	return fmt.Sprintf("[GET /broker/v1/test/timeout][%d] serviceBrokerTestTimeoutNotFound  %+v", 404, o.Payload)
}

func (o *ServiceBrokerTestTimeoutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *ServiceBrokerTestTimeoutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
