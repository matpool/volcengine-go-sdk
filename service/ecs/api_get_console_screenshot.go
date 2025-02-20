// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetConsoleScreenshotCommon = "GetConsoleScreenshot"

// GetConsoleScreenshotCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetConsoleScreenshotCommon operation. The "output" return
// value will be populated with the GetConsoleScreenshotCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetConsoleScreenshotCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetConsoleScreenshotCommon Send returns without error.
//
// See GetConsoleScreenshotCommon for more information on using the GetConsoleScreenshotCommon
// API call, and error handling.
//
//	// Example sending a request using the GetConsoleScreenshotCommonRequest method.
//	req, resp := client.GetConsoleScreenshotCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) GetConsoleScreenshotCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetConsoleScreenshotCommon,
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

// GetConsoleScreenshotCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation GetConsoleScreenshotCommon for usage and error information.
func (c *ECS) GetConsoleScreenshotCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetConsoleScreenshotCommonRequest(input)
	return out, req.Send()
}

// GetConsoleScreenshotCommonWithContext is the same as GetConsoleScreenshotCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetConsoleScreenshotCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) GetConsoleScreenshotCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetConsoleScreenshotCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetConsoleScreenshot = "GetConsoleScreenshot"

// GetConsoleScreenshotRequest generates a "volcengine/request.Request" representing the
// client's request for the GetConsoleScreenshot operation. The "output" return
// value will be populated with the GetConsoleScreenshotCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetConsoleScreenshotCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetConsoleScreenshotCommon Send returns without error.
//
// See GetConsoleScreenshot for more information on using the GetConsoleScreenshot
// API call, and error handling.
//
//	// Example sending a request using the GetConsoleScreenshotRequest method.
//	req, resp := client.GetConsoleScreenshotRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) GetConsoleScreenshotRequest(input *GetConsoleScreenshotInput) (req *request.Request, output *GetConsoleScreenshotOutput) {
	op := &request.Operation{
		Name:       opGetConsoleScreenshot,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetConsoleScreenshotInput{}
	}

	output = &GetConsoleScreenshotOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetConsoleScreenshot API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation GetConsoleScreenshot for usage and error information.
func (c *ECS) GetConsoleScreenshot(input *GetConsoleScreenshotInput) (*GetConsoleScreenshotOutput, error) {
	req, out := c.GetConsoleScreenshotRequest(input)
	return out, req.Send()
}

// GetConsoleScreenshotWithContext is the same as GetConsoleScreenshot with the addition of
// the ability to pass a context and additional request options.
//
// See GetConsoleScreenshot for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) GetConsoleScreenshotWithContext(ctx volcengine.Context, input *GetConsoleScreenshotInput, opts ...request.Option) (*GetConsoleScreenshotOutput, error) {
	req, out := c.GetConsoleScreenshotRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetConsoleScreenshotInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	WakeUp *bool `type:"boolean"`
}

// String returns the string representation
func (s GetConsoleScreenshotInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetConsoleScreenshotInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *GetConsoleScreenshotInput) SetInstanceId(v string) *GetConsoleScreenshotInput {
	s.InstanceId = &v
	return s
}

// SetWakeUp sets the WakeUp field's value.
func (s *GetConsoleScreenshotInput) SetWakeUp(v bool) *GetConsoleScreenshotInput {
	s.WakeUp = &v
	return s
}

type GetConsoleScreenshotOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`

	Screenshot *string `type:"string"`
}

// String returns the string representation
func (s GetConsoleScreenshotOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetConsoleScreenshotOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *GetConsoleScreenshotOutput) SetInstanceId(v string) *GetConsoleScreenshotOutput {
	s.InstanceId = &v
	return s
}

// SetScreenshot sets the Screenshot field's value.
func (s *GetConsoleScreenshotOutput) SetScreenshot(v string) *GetConsoleScreenshotOutput {
	s.Screenshot = &v
	return s
}
