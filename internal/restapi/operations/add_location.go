// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// AddLocationHandlerFunc turns a function with the right signature into a add location handler
type AddLocationHandlerFunc func(AddLocationParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddLocationHandlerFunc) Handle(params AddLocationParams) middleware.Responder {
	return fn(params)
}

// AddLocationHandler interface for that can handle valid add location params
type AddLocationHandler interface {
	Handle(AddLocationParams) middleware.Responder
}

// NewAddLocation creates a new http.Handler for the add location operation
func NewAddLocation(ctx *middleware.Context, handler AddLocationHandler) *AddLocation {
	return &AddLocation{Context: ctx, Handler: handler}
}

/*AddLocation swagger:route POST /bbinput addLocation

Accepts GPS coordinates from the mobile and saves to database

*/
type AddLocation struct {
	Context *middleware.Context
	Handler AddLocationHandler
}

func (o *AddLocation) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddLocationParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddLocationBody add location body
// swagger:model AddLocationBody
type AddLocationBody struct {

	// GPS coordinates of the phone's location
	// Required: true
	Coordinates []float32 `json:"coordinates"`

	// Device identificator
	// Required: true
	Imei *string `json:"imei"`

	// Optional IP address
	IP string `json:"ip,omitempty"`

	// Phone
	// Required: true
	Number *string `json:"number"`

	// EET timestamp in "YYYY/MM/DD-hh:mm:ss" format
	Timestamp string `json:"timestamp,omitempty"`
}

// Validate validates this add location body
func (o *AddLocationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCoordinates(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateImei(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddLocationBody) validateCoordinates(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"coordinates", "body", o.Coordinates); err != nil {
		return err
	}

	return nil
}

func (o *AddLocationBody) validateImei(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"imei", "body", o.Imei); err != nil {
		return err
	}

	return nil
}

func (o *AddLocationBody) validateNumber(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"number", "body", o.Number); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddLocationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddLocationBody) UnmarshalBinary(b []byte) error {
	var res AddLocationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}