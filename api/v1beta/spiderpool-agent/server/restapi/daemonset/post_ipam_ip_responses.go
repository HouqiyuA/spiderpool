// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2022 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package daemonset

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostIpamIPOKCode is the HTTP code returned for type PostIpamIPOK
const PostIpamIPOKCode int = 200

/*PostIpamIPOK Success

swagger:response postIpamIpOK
*/
type PostIpamIPOK struct {
}

// NewPostIpamIPOK creates PostIpamIPOK with default headers values
func NewPostIpamIPOK() *PostIpamIPOK {

	return &PostIpamIPOK{}
}

// WriteResponse to the client
func (o *PostIpamIPOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PostIpamIPInternalServerErrorCode is the HTTP code returned for type PostIpamIPInternalServerError
const PostIpamIPInternalServerErrorCode int = 500

/*PostIpamIPInternalServerError Allocation failure

swagger:response postIpamIpInternalServerError
*/
type PostIpamIPInternalServerError struct {
}

// NewPostIpamIPInternalServerError creates PostIpamIPInternalServerError with default headers values
func NewPostIpamIPInternalServerError() *PostIpamIPInternalServerError {

	return &PostIpamIPInternalServerError{}
}

// WriteResponse to the client
func (o *PostIpamIPInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
