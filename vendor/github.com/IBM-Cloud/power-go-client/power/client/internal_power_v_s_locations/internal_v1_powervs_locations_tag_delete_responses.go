// Code generated by go-swagger; DO NOT EDIT.

package internal_power_v_s_locations

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

// InternalV1PowervsLocationsTagDeleteReader is a Reader for the InternalV1PowervsLocationsTagDelete structure.
type InternalV1PowervsLocationsTagDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1PowervsLocationsTagDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalV1PowervsLocationsTagDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalV1PowervsLocationsTagDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInternalV1PowervsLocationsTagDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewInternalV1PowervsLocationsTagDeleteUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1PowervsLocationsTagDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /internal/v1/powervs/locations/tag] internal.v1.powervs.locations.tag.delete", response, response.Code())
	}
}

// NewInternalV1PowervsLocationsTagDeleteOK creates a InternalV1PowervsLocationsTagDeleteOK with default headers values
func NewInternalV1PowervsLocationsTagDeleteOK() *InternalV1PowervsLocationsTagDeleteOK {
	return &InternalV1PowervsLocationsTagDeleteOK{}
}

/*
InternalV1PowervsLocationsTagDeleteOK describes a response with status code 200, with default header values.

OK
*/
type InternalV1PowervsLocationsTagDeleteOK struct {
}

// IsSuccess returns true when this internal v1 powervs locations tag delete o k response has a 2xx status code
func (o *InternalV1PowervsLocationsTagDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 powervs locations tag delete o k response has a 3xx status code
func (o *InternalV1PowervsLocationsTagDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs locations tag delete o k response has a 4xx status code
func (o *InternalV1PowervsLocationsTagDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 powervs locations tag delete o k response has a 5xx status code
func (o *InternalV1PowervsLocationsTagDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 powervs locations tag delete o k response a status code equal to that given
func (o *InternalV1PowervsLocationsTagDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the internal v1 powervs locations tag delete o k response
func (o *InternalV1PowervsLocationsTagDeleteOK) Code() int {
	return 200
}

func (o *InternalV1PowervsLocationsTagDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteOK", 200)
}

func (o *InternalV1PowervsLocationsTagDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteOK", 200)
}

func (o *InternalV1PowervsLocationsTagDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInternalV1PowervsLocationsTagDeleteBadRequest creates a InternalV1PowervsLocationsTagDeleteBadRequest with default headers values
func NewInternalV1PowervsLocationsTagDeleteBadRequest() *InternalV1PowervsLocationsTagDeleteBadRequest {
	return &InternalV1PowervsLocationsTagDeleteBadRequest{}
}

/*
InternalV1PowervsLocationsTagDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InternalV1PowervsLocationsTagDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 powervs locations tag delete bad request response has a 2xx status code
func (o *InternalV1PowervsLocationsTagDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 powervs locations tag delete bad request response has a 3xx status code
func (o *InternalV1PowervsLocationsTagDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs locations tag delete bad request response has a 4xx status code
func (o *InternalV1PowervsLocationsTagDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 powervs locations tag delete bad request response has a 5xx status code
func (o *InternalV1PowervsLocationsTagDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 powervs locations tag delete bad request response a status code equal to that given
func (o *InternalV1PowervsLocationsTagDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal v1 powervs locations tag delete bad request response
func (o *InternalV1PowervsLocationsTagDeleteBadRequest) Code() int {
	return 400
}

func (o *InternalV1PowervsLocationsTagDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteBadRequest %s", 400, payload)
}

func (o *InternalV1PowervsLocationsTagDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteBadRequest %s", 400, payload)
}

func (o *InternalV1PowervsLocationsTagDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1PowervsLocationsTagDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1PowervsLocationsTagDeleteUnauthorized creates a InternalV1PowervsLocationsTagDeleteUnauthorized with default headers values
func NewInternalV1PowervsLocationsTagDeleteUnauthorized() *InternalV1PowervsLocationsTagDeleteUnauthorized {
	return &InternalV1PowervsLocationsTagDeleteUnauthorized{}
}

/*
InternalV1PowervsLocationsTagDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1PowervsLocationsTagDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 powervs locations tag delete unauthorized response has a 2xx status code
func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 powervs locations tag delete unauthorized response has a 3xx status code
func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs locations tag delete unauthorized response has a 4xx status code
func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 powervs locations tag delete unauthorized response has a 5xx status code
func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 powervs locations tag delete unauthorized response a status code equal to that given
func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the internal v1 powervs locations tag delete unauthorized response
func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) Code() int {
	return 401
}

func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteUnauthorized %s", 401, payload)
}

func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteUnauthorized %s", 401, payload)
}

func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1PowervsLocationsTagDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1PowervsLocationsTagDeleteUnprocessableEntity creates a InternalV1PowervsLocationsTagDeleteUnprocessableEntity with default headers values
func NewInternalV1PowervsLocationsTagDeleteUnprocessableEntity() *InternalV1PowervsLocationsTagDeleteUnprocessableEntity {
	return &InternalV1PowervsLocationsTagDeleteUnprocessableEntity{}
}

/*
InternalV1PowervsLocationsTagDeleteUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type InternalV1PowervsLocationsTagDeleteUnprocessableEntity struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 powervs locations tag delete unprocessable entity response has a 2xx status code
func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 powervs locations tag delete unprocessable entity response has a 3xx status code
func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs locations tag delete unprocessable entity response has a 4xx status code
func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 powervs locations tag delete unprocessable entity response has a 5xx status code
func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 powervs locations tag delete unprocessable entity response a status code equal to that given
func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) IsCode(code int) bool {
	return code == 422
}

// Code gets the status code for the internal v1 powervs locations tag delete unprocessable entity response
func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) Code() int {
	return 422
}

func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteUnprocessableEntity %s", 422, payload)
}

func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteUnprocessableEntity %s", 422, payload)
}

func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1PowervsLocationsTagDeleteUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1PowervsLocationsTagDeleteInternalServerError creates a InternalV1PowervsLocationsTagDeleteInternalServerError with default headers values
func NewInternalV1PowervsLocationsTagDeleteInternalServerError() *InternalV1PowervsLocationsTagDeleteInternalServerError {
	return &InternalV1PowervsLocationsTagDeleteInternalServerError{}
}

/*
InternalV1PowervsLocationsTagDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1PowervsLocationsTagDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 powervs locations tag delete internal server error response has a 2xx status code
func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 powervs locations tag delete internal server error response has a 3xx status code
func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 powervs locations tag delete internal server error response has a 4xx status code
func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 powervs locations tag delete internal server error response has a 5xx status code
func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 powervs locations tag delete internal server error response a status code equal to that given
func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 powervs locations tag delete internal server error response
func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) Code() int {
	return 500
}

func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteInternalServerError %s", 500, payload)
}

func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/powervs/locations/tag][%d] internalV1PowervsLocationsTagDeleteInternalServerError %s", 500, payload)
}

func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1PowervsLocationsTagDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}