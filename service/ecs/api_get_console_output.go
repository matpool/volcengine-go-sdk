// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetConsoleOutputCommon = "GetConsoleOutput"

// GetConsoleOutputCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetConsoleOutputCommon operation. The "output" return
// value will be populated with the GetConsoleOutputCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetConsoleOutputCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetConsoleOutputCommon Send returns without error.
//
// See GetConsoleOutputCommon for more information on using the GetConsoleOutputCommon
// API call, and error handling.
//
//	// Example sending a request using the GetConsoleOutputCommonRequest method.
//	req, resp := client.GetConsoleOutputCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) GetConsoleOutputCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetConsoleOutputCommon,
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

// GetConsoleOutputCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation GetConsoleOutputCommon for usage and error information.
func (c *ECS) GetConsoleOutputCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetConsoleOutputCommonRequest(input)
	return out, req.Send()
}

// GetConsoleOutputCommonWithContext is the same as GetConsoleOutputCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetConsoleOutputCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) GetConsoleOutputCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetConsoleOutputCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetConsoleOutput = "GetConsoleOutput"

// GetConsoleOutputRequest generates a "volcengine/request.Request" representing the
// client's request for the GetConsoleOutput operation. The "output" return
// value will be populated with the GetConsoleOutputCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetConsoleOutputCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetConsoleOutputCommon Send returns without error.
//
// See GetConsoleOutput for more information on using the GetConsoleOutput
// API call, and error handling.
//
//	// Example sending a request using the GetConsoleOutputRequest method.
//	req, resp := client.GetConsoleOutputRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) GetConsoleOutputRequest(input *GetConsoleOutputInput) (req *request.Request, output *GetConsoleOutputOutput) {
	op := &request.Operation{
		Name:       opGetConsoleOutput,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConsoleOutputInput{}
	}

	output = &GetConsoleOutputOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetConsoleOutput API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation GetConsoleOutput for usage and error information.
func (c *ECS) GetConsoleOutput(input *GetConsoleOutputInput) (*GetConsoleOutputOutput, error) {
	req, out := c.GetConsoleOutputRequest(input)
	return out, req.Send()
}

// GetConsoleOutputWithContext is the same as GetConsoleOutput with the addition of
// the ability to pass a context and additional request options.
//
// See GetConsoleOutput for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) GetConsoleOutputWithContext(ctx volcengine.Context, input *GetConsoleOutputInput, opts ...request.Option) (*GetConsoleOutputOutput, error) {
	req, out := c.GetConsoleOutputRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetConsoleOutputInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s GetConsoleOutputInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetConsoleOutputInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *GetConsoleOutputInput) SetInstanceId(v string) *GetConsoleOutputInput {
	s.InstanceId = &v
	return s
}

type GetConsoleOutputOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`

	LastUpdateAt *string `type:"string"`

	Output *string `type:"string"`
}

// String returns the string representation
func (s GetConsoleOutputOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetConsoleOutputOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *GetConsoleOutputOutput) SetInstanceId(v string) *GetConsoleOutputOutput {
	s.InstanceId = &v
	return s
}

// SetLastUpdateAt sets the LastUpdateAt field's value.
func (s *GetConsoleOutputOutput) SetLastUpdateAt(v string) *GetConsoleOutputOutput {
	s.LastUpdateAt = &v
	return s
}

// SetOutput sets the Output field's value.
func (s *GetConsoleOutputOutput) SetOutput(v string) *GetConsoleOutputOutput {
	s.Output = &v
	return s
}
