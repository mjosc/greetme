// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// UserByNameHandlerFunc turns a function with the right signature into a user by name handler
type UserByNameHandlerFunc func(UserByNameParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UserByNameHandlerFunc) Handle(params UserByNameParams) middleware.Responder {
	return fn(params)
}

// UserByNameHandler interface for that can handle valid user by name params
type UserByNameHandler interface {
	Handle(UserByNameParams) middleware.Responder
}

// NewUserByName creates a new http.Handler for the user by name operation
func NewUserByName(ctx *middleware.Context, handler UserByNameHandler) *UserByName {
	return &UserByName{Context: ctx, Handler: handler}
}

/*UserByName swagger:route GET /users/{name} users userByName

UserByName user by name API

*/
type UserByName struct {
	Context *middleware.Context
	Handler UserByNameHandler
}

func (o *UserByName) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUserByNameParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// UserByNameOKBody user by name o k body
// swagger:model UserByNameOKBody
type UserByNameOKBody struct {

	// error
	// Min Length: 1
	Error string `json:"error,omitempty"`

	// The user ID
	// Minimum: 1
	ID int64 `json:"id,omitempty"`

	// message
	// Min Length: 1
	Message string `json:"message,omitempty"`

	// The name of the user
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// valid
	Valid bool `json:"valid,omitempty"`
}

// Validate validates this user by name o k body
func (o *UserByNameOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UserByNameOKBody) validateError(formats strfmt.Registry) error {

	if swag.IsZero(o.Error) { // not required
		return nil
	}

	if err := validate.MinLength("userByNameOK"+"."+"error", "body", string(o.Error), 1); err != nil {
		return err
	}

	return nil
}

func (o *UserByNameOKBody) validateID(formats strfmt.Registry) error {

	if swag.IsZero(o.ID) { // not required
		return nil
	}

	if err := validate.MinimumInt("userByNameOK"+"."+"id", "body", int64(o.ID), 1, false); err != nil {
		return err
	}

	return nil
}

func (o *UserByNameOKBody) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(o.Message) { // not required
		return nil
	}

	if err := validate.MinLength("userByNameOK"+"."+"message", "body", string(o.Message), 1); err != nil {
		return err
	}

	return nil
}

func (o *UserByNameOKBody) validateName(formats strfmt.Registry) error {

	if swag.IsZero(o.Name) { // not required
		return nil
	}

	if err := validate.MinLength("userByNameOK"+"."+"name", "body", string(o.Name), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UserByNameOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UserByNameOKBody) UnmarshalBinary(b []byte) error {
	var res UserByNameOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
