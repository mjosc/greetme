// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UserByNameOKCode is the HTTP code returned for type UserByNameOK
const UserByNameOKCode int = 200

/*UserByNameOK A success response describing the user if it exists and an error if not

swagger:response userByNameOK
*/
type UserByNameOK struct {

	/*
	  In: Body
	*/
	Payload *UserByNameOKBody `json:"body,omitempty"`
}

// NewUserByNameOK creates UserByNameOK with default headers values
func NewUserByNameOK() *UserByNameOK {

	return &UserByNameOK{}
}

// WithPayload adds the payload to the user by name o k response
func (o *UserByNameOK) WithPayload(payload *UserByNameOKBody) *UserByNameOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the user by name o k response
func (o *UserByNameOK) SetPayload(payload *UserByNameOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UserByNameOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
