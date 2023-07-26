// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceParamsCommon = "ModifyDBInstanceParams"

// ModifyDBInstanceParamsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceParamsCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceParamsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceParamsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceParamsCommon Send returns without error.
//
// See ModifyDBInstanceParamsCommon for more information on using the ModifyDBInstanceParamsCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceParamsCommonRequest method.
//    req, resp := client.ModifyDBInstanceParamsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceParamsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceParamsCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceParamsCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceParamsCommon for usage and error information.
func (c *REDIS) ModifyDBInstanceParamsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceParamsCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceParamsCommonWithContext is the same as ModifyDBInstanceParamsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceParamsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceParamsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceParamsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstanceParams = "ModifyDBInstanceParams"

// ModifyDBInstanceParamsRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceParams operation. The "output" return
// value will be populated with the ModifyDBInstanceParamsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceParamsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceParamsCommon Send returns without error.
//
// See ModifyDBInstanceParams for more information on using the ModifyDBInstanceParams
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceParamsRequest method.
//    req, resp := client.ModifyDBInstanceParamsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) ModifyDBInstanceParamsRequest(input *ModifyDBInstanceParamsInput) (req *request.Request, output *ModifyDBInstanceParamsOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstanceParams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceParamsInput{}
	}

	output = &ModifyDBInstanceParamsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstanceParams API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation ModifyDBInstanceParams for usage and error information.
func (c *REDIS) ModifyDBInstanceParams(input *ModifyDBInstanceParamsInput) (*ModifyDBInstanceParamsOutput, error) {
	req, out := c.ModifyDBInstanceParamsRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceParamsWithContext is the same as ModifyDBInstanceParams with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceParams for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) ModifyDBInstanceParamsWithContext(ctx volcengine.Context, input *ModifyDBInstanceParamsInput, opts ...request.Option) (*ModifyDBInstanceParamsOutput, error) {
	req, out := c.ModifyDBInstanceParamsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyDBInstanceParamsInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	ParamValues []*ParamValueForModifyDBInstanceParamsInput `type:"list"`
}

// String returns the string representation
func (s ModifyDBInstanceParamsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceParamsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceParamsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceParamsInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyDBInstanceParamsInput) SetClientToken(v string) *ModifyDBInstanceParamsInput {
	s.ClientToken = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceParamsInput) SetInstanceId(v string) *ModifyDBInstanceParamsInput {
	s.InstanceId = &v
	return s
}

// SetParamValues sets the ParamValues field's value.
func (s *ModifyDBInstanceParamsInput) SetParamValues(v []*ParamValueForModifyDBInstanceParamsInput) *ModifyDBInstanceParamsInput {
	s.ParamValues = v
	return s
}

type ModifyDBInstanceParamsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyDBInstanceParamsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceParamsOutput) GoString() string {
	return s.String()
}

type ParamValueForModifyDBInstanceParamsInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s ParamValueForModifyDBInstanceParamsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ParamValueForModifyDBInstanceParamsInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *ParamValueForModifyDBInstanceParamsInput) SetName(v string) *ParamValueForModifyDBInstanceParamsInput {
	s.Name = &v
	return s
}

// SetValue sets the Value field's value.
func (s *ParamValueForModifyDBInstanceParamsInput) SetValue(v string) *ParamValueForModifyDBInstanceParamsInput {
	s.Value = &v
	return s
}
