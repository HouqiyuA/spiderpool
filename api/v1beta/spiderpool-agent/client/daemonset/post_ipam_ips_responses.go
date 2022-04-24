// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package daemonset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostIpamIpsReader is a Reader for the PostIpamIps structure.
type PostIpamIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostIpamIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostIpamIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostIpamIpsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostIpamIpsOK creates a PostIpamIpsOK with default headers values
func NewPostIpamIpsOK() *PostIpamIpsOK {
	return &PostIpamIpsOK{}
}

/* PostIpamIpsOK describes a response with status code 200, with default header values.

Success
*/
type PostIpamIpsOK struct {
}

func (o *PostIpamIpsOK) Error() string {
	return fmt.Sprintf("[POST /ipam/ips][%d] postIpamIpsOK ", 200)
}

func (o *PostIpamIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostIpamIpsInternalServerError creates a PostIpamIpsInternalServerError with default headers values
func NewPostIpamIpsInternalServerError() *PostIpamIpsInternalServerError {
	return &PostIpamIpsInternalServerError{}
}

/* PostIpamIpsInternalServerError describes a response with status code 500, with default header values.

Allocation failure
*/
type PostIpamIpsInternalServerError struct {
}

func (o *PostIpamIpsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ipam/ips][%d] postIpamIpsInternalServerError ", 500)
}

func (o *PostIpamIpsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
