// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyInstanceSpecCommon = "ModifyInstanceSpec"

// ModifyInstanceSpecCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyInstanceSpecCommon operation. The "output" return
// value will be populated with the ModifyInstanceSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceSpecCommon Send returns without error.
//
// See ModifyInstanceSpecCommon for more information on using the ModifyInstanceSpecCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceSpecCommonRequest method.
//    req, resp := client.ModifyInstanceSpecCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceSpecCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyInstanceSpecCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyInstanceSpecCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for ECS's
// API operation ModifyInstanceSpecCommon for usage and error information.
func (c *ECS) ModifyInstanceSpecCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceSpecCommonRequest(input)
	return out, req.Send()
}

// ModifyInstanceSpecCommonWithContext is the same as ModifyInstanceSpecCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceSpecCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceSpecCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyInstanceSpecCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyInstanceSpec = "ModifyInstanceSpec"

// ModifyInstanceSpecRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyInstanceSpec operation. The "output" return
// value will be populated with the ModifyInstanceSpecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyInstanceSpecCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyInstanceSpecCommon Send returns without error.
//
// See ModifyInstanceSpec for more information on using the ModifyInstanceSpec
// API call, and error handling.
//
//    // Example sending a request using the ModifyInstanceSpecRequest method.
//    req, resp := client.ModifyInstanceSpecRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) ModifyInstanceSpecRequest(input *ModifyInstanceSpecInput) (req *request.Request, output *ModifyInstanceSpecOutput) {
	op := &request.Operation{
		Name:       opModifyInstanceSpec,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyInstanceSpecInput{}
	}

	output = &ModifyInstanceSpecOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyInstanceSpec API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for ECS's
// API operation ModifyInstanceSpec for usage and error information.
func (c *ECS) ModifyInstanceSpec(input *ModifyInstanceSpecInput) (*ModifyInstanceSpecOutput, error) {
	req, out := c.ModifyInstanceSpecRequest(input)
	return out, req.Send()
}

// ModifyInstanceSpecWithContext is the same as ModifyInstanceSpec with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyInstanceSpec for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) ModifyInstanceSpecWithContext(ctx volcengine.Context, input *ModifyInstanceSpecInput, opts ...request.Option) (*ModifyInstanceSpecOutput, error) {
	req, out := c.ModifyInstanceSpecRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyInstanceSpecInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceType *string `type:"string"`
}

// String returns the string representation
func (s ModifyInstanceSpecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceSpecInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *ModifyInstanceSpecInput) SetClientToken(v string) *ModifyInstanceSpecInput {
	s.ClientToken = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyInstanceSpecInput) SetInstanceId(v string) *ModifyInstanceSpecInput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *ModifyInstanceSpecInput) SetInstanceType(v string) *ModifyInstanceSpecInput {
	s.InstanceType = &v
	return s
}

type ModifyInstanceSpecOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyInstanceSpecOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyInstanceSpecOutput) GoString() string {
	return s.String()
}
