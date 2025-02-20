// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCompleteLifecycleActivityCommon = "CompleteLifecycleActivity"

// CompleteLifecycleActivityCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CompleteLifecycleActivityCommon operation. The "output" return
// value will be populated with the CompleteLifecycleActivityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CompleteLifecycleActivityCommon Request to send the API call to the service.
// the "output" return value is not valid until after CompleteLifecycleActivityCommon Send returns without error.
//
// See CompleteLifecycleActivityCommon for more information on using the CompleteLifecycleActivityCommon
// API call, and error handling.
//
//	// Example sending a request using the CompleteLifecycleActivityCommonRequest method.
//	req, resp := client.CompleteLifecycleActivityCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) CompleteLifecycleActivityCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCompleteLifecycleActivityCommon,
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

// CompleteLifecycleActivityCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation CompleteLifecycleActivityCommon for usage and error information.
func (c *AUTOSCALING) CompleteLifecycleActivityCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CompleteLifecycleActivityCommonRequest(input)
	return out, req.Send()
}

// CompleteLifecycleActivityCommonWithContext is the same as CompleteLifecycleActivityCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CompleteLifecycleActivityCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) CompleteLifecycleActivityCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CompleteLifecycleActivityCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCompleteLifecycleActivity = "CompleteLifecycleActivity"

// CompleteLifecycleActivityRequest generates a "volcengine/request.Request" representing the
// client's request for the CompleteLifecycleActivity operation. The "output" return
// value will be populated with the CompleteLifecycleActivityCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CompleteLifecycleActivityCommon Request to send the API call to the service.
// the "output" return value is not valid until after CompleteLifecycleActivityCommon Send returns without error.
//
// See CompleteLifecycleActivity for more information on using the CompleteLifecycleActivity
// API call, and error handling.
//
//	// Example sending a request using the CompleteLifecycleActivityRequest method.
//	req, resp := client.CompleteLifecycleActivityRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) CompleteLifecycleActivityRequest(input *CompleteLifecycleActivityInput) (req *request.Request, output *CompleteLifecycleActivityOutput) {
	op := &request.Operation{
		Name:       opCompleteLifecycleActivity,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CompleteLifecycleActivityInput{}
	}

	output = &CompleteLifecycleActivityOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CompleteLifecycleActivity API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation CompleteLifecycleActivity for usage and error information.
func (c *AUTOSCALING) CompleteLifecycleActivity(input *CompleteLifecycleActivityInput) (*CompleteLifecycleActivityOutput, error) {
	req, out := c.CompleteLifecycleActivityRequest(input)
	return out, req.Send()
}

// CompleteLifecycleActivityWithContext is the same as CompleteLifecycleActivity with the addition of
// the ability to pass a context and additional request options.
//
// See CompleteLifecycleActivity for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) CompleteLifecycleActivityWithContext(ctx volcengine.Context, input *CompleteLifecycleActivityInput, opts ...request.Option) (*CompleteLifecycleActivityOutput, error) {
	req, out := c.CompleteLifecycleActivityRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CompleteLifecycleActivityInput struct {
	_ struct{} `type:"structure"`

	LifecycleActivityId *string `type:"string"`

	LifecycleActivityPolicy *string `type:"string"`
}

// String returns the string representation
func (s CompleteLifecycleActivityInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CompleteLifecycleActivityInput) GoString() string {
	return s.String()
}

// SetLifecycleActivityId sets the LifecycleActivityId field's value.
func (s *CompleteLifecycleActivityInput) SetLifecycleActivityId(v string) *CompleteLifecycleActivityInput {
	s.LifecycleActivityId = &v
	return s
}

// SetLifecycleActivityPolicy sets the LifecycleActivityPolicy field's value.
func (s *CompleteLifecycleActivityInput) SetLifecycleActivityPolicy(v string) *CompleteLifecycleActivityInput {
	s.LifecycleActivityPolicy = &v
	return s
}

type CompleteLifecycleActivityOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`

	LifecycleActivityId *string `type:"string"`
}

// String returns the string representation
func (s CompleteLifecycleActivityOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CompleteLifecycleActivityOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *CompleteLifecycleActivityOutput) SetInstanceId(v string) *CompleteLifecycleActivityOutput {
	s.InstanceId = &v
	return s
}

// SetLifecycleActivityId sets the LifecycleActivityId field's value.
func (s *CompleteLifecycleActivityOutput) SetLifecycleActivityId(v string) *CompleteLifecycleActivityOutput {
	s.LifecycleActivityId = &v
	return s
}
