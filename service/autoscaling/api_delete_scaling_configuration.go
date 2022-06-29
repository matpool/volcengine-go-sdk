// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteScalingConfigurationCommon = "DeleteScalingConfiguration"

// DeleteScalingConfigurationCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteScalingConfigurationCommon operation. The "output" return
// value will be populated with the DeleteScalingConfigurationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteScalingConfigurationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteScalingConfigurationCommon Send returns without error.
//
// See DeleteScalingConfigurationCommon for more information on using the DeleteScalingConfigurationCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteScalingConfigurationCommonRequest method.
//    req, resp := client.DeleteScalingConfigurationCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DeleteScalingConfigurationCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteScalingConfigurationCommon,
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

// DeleteScalingConfigurationCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation DeleteScalingConfigurationCommon for usage and error information.
func (c *AUTOSCALING) DeleteScalingConfigurationCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteScalingConfigurationCommonRequest(input)
	return out, req.Send()
}

// DeleteScalingConfigurationCommonWithContext is the same as DeleteScalingConfigurationCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteScalingConfigurationCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteScalingConfigurationCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteScalingConfigurationCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteScalingConfiguration = "DeleteScalingConfiguration"

// DeleteScalingConfigurationRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteScalingConfiguration operation. The "output" return
// value will be populated with the DeleteScalingConfigurationCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteScalingConfigurationCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteScalingConfigurationCommon Send returns without error.
//
// See DeleteScalingConfiguration for more information on using the DeleteScalingConfiguration
// API call, and error handling.
//
//    // Example sending a request using the DeleteScalingConfigurationRequest method.
//    req, resp := client.DeleteScalingConfigurationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DeleteScalingConfigurationRequest(input *DeleteScalingConfigurationInput) (req *request.Request, output *DeleteScalingConfigurationOutput) {
	op := &request.Operation{
		Name:       opDeleteScalingConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteScalingConfigurationInput{}
	}

	output = &DeleteScalingConfigurationOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteScalingConfiguration API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation DeleteScalingConfiguration for usage and error information.
func (c *AUTOSCALING) DeleteScalingConfiguration(input *DeleteScalingConfigurationInput) (*DeleteScalingConfigurationOutput, error) {
	req, out := c.DeleteScalingConfigurationRequest(input)
	return out, req.Send()
}

// DeleteScalingConfigurationWithContext is the same as DeleteScalingConfiguration with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteScalingConfiguration for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteScalingConfigurationWithContext(ctx volcengine.Context, input *DeleteScalingConfigurationInput, opts ...request.Option) (*DeleteScalingConfigurationOutput, error) {
	req, out := c.DeleteScalingConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteScalingConfigurationInput struct {
	_ struct{} `type:"structure"`

	ScalingConfigurationId *string `type:"string"`
}

// String returns the string representation
func (s DeleteScalingConfigurationInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteScalingConfigurationInput) GoString() string {
	return s.String()
}

// SetScalingConfigurationId sets the ScalingConfigurationId field's value.
func (s *DeleteScalingConfigurationInput) SetScalingConfigurationId(v string) *DeleteScalingConfigurationInput {
	s.ScalingConfigurationId = &v
	return s
}

type DeleteScalingConfigurationOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingConfigurationId *string `type:"string"`
}

// String returns the string representation
func (s DeleteScalingConfigurationOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteScalingConfigurationOutput) GoString() string {
	return s.String()
}

// SetScalingConfigurationId sets the ScalingConfigurationId field's value.
func (s *DeleteScalingConfigurationOutput) SetScalingConfigurationId(v string) *DeleteScalingConfigurationOutput {
	s.ScalingConfigurationId = &v
	return s
}
