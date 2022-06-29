// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAttachServerGroupsCommon = "AttachServerGroups"

// AttachServerGroupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachServerGroupsCommon operation. The "output" return
// value will be populated with the AttachServerGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachServerGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachServerGroupsCommon Send returns without error.
//
// See AttachServerGroupsCommon for more information on using the AttachServerGroupsCommon
// API call, and error handling.
//
//    // Example sending a request using the AttachServerGroupsCommonRequest method.
//    req, resp := client.AttachServerGroupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) AttachServerGroupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachServerGroupsCommon,
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

// AttachServerGroupsCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation AttachServerGroupsCommon for usage and error information.
func (c *AUTOSCALING) AttachServerGroupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachServerGroupsCommonRequest(input)
	return out, req.Send()
}

// AttachServerGroupsCommonWithContext is the same as AttachServerGroupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachServerGroupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) AttachServerGroupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachServerGroupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachServerGroups = "AttachServerGroups"

// AttachServerGroupsRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachServerGroups operation. The "output" return
// value will be populated with the AttachServerGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachServerGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachServerGroupsCommon Send returns without error.
//
// See AttachServerGroups for more information on using the AttachServerGroups
// API call, and error handling.
//
//    // Example sending a request using the AttachServerGroupsRequest method.
//    req, resp := client.AttachServerGroupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) AttachServerGroupsRequest(input *AttachServerGroupsInput) (req *request.Request, output *AttachServerGroupsOutput) {
	op := &request.Operation{
		Name:       opAttachServerGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachServerGroupsInput{}
	}

	output = &AttachServerGroupsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachServerGroups API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for AUTO_SCALING's
// API operation AttachServerGroups for usage and error information.
func (c *AUTOSCALING) AttachServerGroups(input *AttachServerGroupsInput) (*AttachServerGroupsOutput, error) {
	req, out := c.AttachServerGroupsRequest(input)
	return out, req.Send()
}

// AttachServerGroupsWithContext is the same as AttachServerGroups with the addition of
// the ability to pass a context and additional request options.
//
// See AttachServerGroups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) AttachServerGroupsWithContext(ctx volcengine.Context, input *AttachServerGroupsInput, opts ...request.Option) (*AttachServerGroupsOutput, error) {
	req, out := c.AttachServerGroupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachServerGroupsInput struct {
	_ struct{} `type:"structure"`

	ScalingGroupId *string `type:"string"`

	ServerGroupAttributes []*ServerGroupAttributeForAttachServerGroupsInput `type:"list"`
}

// String returns the string representation
func (s AttachServerGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachServerGroupsInput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *AttachServerGroupsInput) SetScalingGroupId(v string) *AttachServerGroupsInput {
	s.ScalingGroupId = &v
	return s
}

// SetServerGroupAttributes sets the ServerGroupAttributes field's value.
func (s *AttachServerGroupsInput) SetServerGroupAttributes(v []*ServerGroupAttributeForAttachServerGroupsInput) *AttachServerGroupsInput {
	s.ServerGroupAttributes = v
	return s
}

type AttachServerGroupsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s AttachServerGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachServerGroupsOutput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *AttachServerGroupsOutput) SetScalingGroupId(v string) *AttachServerGroupsOutput {
	s.ScalingGroupId = &v
	return s
}

type ServerGroupAttributeForAttachServerGroupsInput struct {
	_ struct{} `type:"structure"`

	Port *int32 `type:"int32"`

	ServerGroupId *string `type:"string"`

	Weight *int32 `type:"int32"`
}

// String returns the string representation
func (s ServerGroupAttributeForAttachServerGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerGroupAttributeForAttachServerGroupsInput) GoString() string {
	return s.String()
}

// SetPort sets the Port field's value.
func (s *ServerGroupAttributeForAttachServerGroupsInput) SetPort(v int32) *ServerGroupAttributeForAttachServerGroupsInput {
	s.Port = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ServerGroupAttributeForAttachServerGroupsInput) SetServerGroupId(v string) *ServerGroupAttributeForAttachServerGroupsInput {
	s.ServerGroupId = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ServerGroupAttributeForAttachServerGroupsInput) SetWeight(v int32) *ServerGroupAttributeForAttachServerGroupsInput {
	s.Weight = &v
	return s
}
