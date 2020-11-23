// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	alternate "github.com/go-swagger/go-swagger/examples/external-types/fred"
)

// NewPutTestParams creates a new PutTestParams object
// no default values defined in spec.
func NewPutTestParams() PutTestParams {

	return PutTestParams{}
}

// PutTestParams contains all the bound params for the put test operation
// typically these are obtained from a http.Request
//
// swagger:parameters PutTest
type PutTestParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A reference to a type defined in the "fred" package, aliased
	as "alternate".

	MyAlternate alternate.MyAlternateType

	  Required: true
	  In: body
	*/
	MyAlternate alternate.MyAlternateType
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewPutTestParams() beforehand.
func (o *PutTestParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body alternate.MyAlternateType
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("myAlternate", "body", ""))
			} else {
				res = append(res, errors.NewParseError("myAlternate", "body", "", err))
			}
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			ctx := validate.WithOperationRequest(context.Background())
			if err := body.ContextValidate(ctx, route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.MyAlternate = body
			}
		}
	} else {
		res = append(res, errors.Required("myAlternate", "body", ""))
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
