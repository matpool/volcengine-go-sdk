// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnableTLSAccessLogCommon = "EnableTLSAccessLog"

// EnableTLSAccessLogCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableTLSAccessLogCommon operation. The "output" return
// value will be populated with the EnableTLSAccessLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableTLSAccessLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableTLSAccessLogCommon Send returns without error.
//
// See EnableTLSAccessLogCommon for more information on using the EnableTLSAccessLogCommon
// API call, and error handling.
//
//    // Example sending a request using the EnableTLSAccessLogCommonRequest method.
//    req, resp := client.EnableTLSAccessLogCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) EnableTLSAccessLogCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnableTLSAccessLogCommon,
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

// EnableTLSAccessLogCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation EnableTLSAccessLogCommon for usage and error information.
func (c *ALB) EnableTLSAccessLogCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnableTLSAccessLogCommonRequest(input)
	return out, req.Send()
}

// EnableTLSAccessLogCommonWithContext is the same as EnableTLSAccessLogCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnableTLSAccessLogCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) EnableTLSAccessLogCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnableTLSAccessLogCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnableTLSAccessLog = "EnableTLSAccessLog"

// EnableTLSAccessLogRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableTLSAccessLog operation. The "output" return
// value will be populated with the EnableTLSAccessLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableTLSAccessLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableTLSAccessLogCommon Send returns without error.
//
// See EnableTLSAccessLog for more information on using the EnableTLSAccessLog
// API call, and error handling.
//
//    // Example sending a request using the EnableTLSAccessLogRequest method.
//    req, resp := client.EnableTLSAccessLogRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) EnableTLSAccessLogRequest(input *EnableTLSAccessLogInput) (req *request.Request, output *EnableTLSAccessLogOutput) {
	op := &request.Operation{
		Name:       opEnableTLSAccessLog,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableTLSAccessLogInput{}
	}

	output = &EnableTLSAccessLogOutput{}
	req = c.newRequest(op, input, output)

	return
}

// EnableTLSAccessLog API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation EnableTLSAccessLog for usage and error information.
func (c *ALB) EnableTLSAccessLog(input *EnableTLSAccessLogInput) (*EnableTLSAccessLogOutput, error) {
	req, out := c.EnableTLSAccessLogRequest(input)
	return out, req.Send()
}

// EnableTLSAccessLogWithContext is the same as EnableTLSAccessLog with the addition of
// the ability to pass a context and additional request options.
//
// See EnableTLSAccessLog for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) EnableTLSAccessLogWithContext(ctx volcengine.Context, input *EnableTLSAccessLogInput, opts ...request.Option) (*EnableTLSAccessLogOutput, error) {
	req, out := c.EnableTLSAccessLogRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EnableTLSAccessLogInput struct {
	_ struct{} `type:"structure"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`

	// ProjectId is a required field
	ProjectId *string `type:"string" required:"true"`

	// TopicId is a required field
	TopicId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableTLSAccessLogInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableTLSAccessLogInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableTLSAccessLogInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EnableTLSAccessLogInput"}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}
	if s.ProjectId == nil {
		invalidParams.Add(request.NewErrParamRequired("ProjectId"))
	}
	if s.TopicId == nil {
		invalidParams.Add(request.NewErrParamRequired("TopicId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *EnableTLSAccessLogInput) SetLoadBalancerId(v string) *EnableTLSAccessLogInput {
	s.LoadBalancerId = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *EnableTLSAccessLogInput) SetProjectId(v string) *EnableTLSAccessLogInput {
	s.ProjectId = &v
	return s
}

// SetTopicId sets the TopicId field's value.
func (s *EnableTLSAccessLogInput) SetTopicId(v string) *EnableTLSAccessLogInput {
	s.TopicId = &v
	return s
}

type EnableTLSAccessLogOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s EnableTLSAccessLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableTLSAccessLogOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *EnableTLSAccessLogOutput) SetRequestId(v string) *EnableTLSAccessLogOutput {
	s.RequestId = &v
	return s
}
