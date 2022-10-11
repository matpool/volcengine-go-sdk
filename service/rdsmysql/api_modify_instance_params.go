// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyInstanceParamsCommon = "ModifyInstanceParams"

// ModifyInstanceParamsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyInstanceParamsCommon operation. The "output" return
// value will be populated with the ModifyInstanceParamsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceParamsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceParamsCommon Send returns without error.
//
// See ModifyInstanceParamsCommon for more information on using the ModifyInstanceParamsCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceParamsCommonRequest method.
//    req, resp := client.ModifyInstanceParamsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ModifyInstanceParamsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyInstanceParamsCommon,
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

// ModifyInstanceParamsCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ModifyInstanceParamsCommon for usage and error information.
func (c *RDSMYSQL) ModifyInstanceParamsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceParamsCommonRequest(input)
	return out, req.Send()
}

// ModifyInstanceParamsCommonWithContext is the same as ModifyInstanceParamsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceParamsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ModifyInstanceParamsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceParamsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyInstanceParams = "ModifyInstanceParams"

// ModifyInstanceParamsRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyInstanceParams operation. The "output" return
// value will be populated with the ModifyInstanceParamsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceParamsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceParamsCommon Send returns without error.
//
// See ModifyInstanceParams for more information on using the ModifyInstanceParams
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceParamsRequest method.
//    req, resp := client.ModifyInstanceParamsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ModifyInstanceParamsRequest(input *ModifyInstanceParamsInput) (req *request.Request, output *ModifyInstanceParamsOutput) {
	op := &request.Operation{
		Name:       opModifyInstanceParams,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstanceParamsInput{}
	}

	output = &ModifyInstanceParamsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyInstanceParams API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ModifyInstanceParams for usage and error information.
func (c *RDSMYSQL) ModifyInstanceParams(input *ModifyInstanceParamsInput) (*ModifyInstanceParamsOutput, error) {
	req, out := c.ModifyInstanceParamsRequest(input)
	return out, req.Send()
}

// ModifyInstanceParamsWithContext is the same as ModifyInstanceParams with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceParams for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ModifyInstanceParamsWithContext(ctx volcengine.Context, input *ModifyInstanceParamsInput, opts ...request.Option) (*ModifyInstanceParamsOutput, error) {
	req, out := c.ModifyInstanceParamsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyInstanceParamsInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	Parameters []*ParameterForModifyInstanceParamsInput `type:"list"`
}

// String returns the string representation
func (s ModifyInstanceParamsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceParamsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyInstanceParamsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyInstanceParamsInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyInstanceParamsInput) SetInstanceId(v string) *ModifyInstanceParamsInput {
	s.InstanceId = &v
	return s
}

// SetParameters sets the Parameters field's value.
func (s *ModifyInstanceParamsInput) SetParameters(v []*ParameterForModifyInstanceParamsInput) *ModifyInstanceParamsInput {
	s.Parameters = v
	return s
}

type ModifyInstanceParamsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyInstanceParamsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceParamsOutput) GoString() string {
	return s.String()
}

type ParameterForModifyInstanceParamsInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	NewValue *string `type:"string"`

	OldValue *string `type:"string"`

	Restart *bool `type:"boolean"`

	Result *string `type:"string"`
}

// String returns the string representation
func (s ParameterForModifyInstanceParamsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ParameterForModifyInstanceParamsInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *ParameterForModifyInstanceParamsInput) SetName(v string) *ParameterForModifyInstanceParamsInput {
	s.Name = &v
	return s
}

// SetNewValue sets the NewValue field's value.
func (s *ParameterForModifyInstanceParamsInput) SetNewValue(v string) *ParameterForModifyInstanceParamsInput {
	s.NewValue = &v
	return s
}

// SetOldValue sets the OldValue field's value.
func (s *ParameterForModifyInstanceParamsInput) SetOldValue(v string) *ParameterForModifyInstanceParamsInput {
	s.OldValue = &v
	return s
}

// SetRestart sets the Restart field's value.
func (s *ParameterForModifyInstanceParamsInput) SetRestart(v bool) *ParameterForModifyInstanceParamsInput {
	s.Restart = &v
	return s
}

// SetResult sets the Result field's value.
func (s *ParameterForModifyInstanceParamsInput) SetResult(v string) *ParameterForModifyInstanceParamsInput {
	s.Result = &v
	return s
}
