// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opEnableAccessLogCommon = "EnableAccessLog"

// EnableAccessLogCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableAccessLogCommon operation. The "output" return
// value will be populated with the EnableAccessLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableAccessLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableAccessLogCommon Send returns without error.
//
// See EnableAccessLogCommon for more information on using the EnableAccessLogCommon
// API call, and error handling.
//
//    // Example sending a request using the EnableAccessLogCommonRequest method.
//    req, resp := client.EnableAccessLogCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) EnableAccessLogCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opEnableAccessLogCommon,
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

// EnableAccessLogCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation EnableAccessLogCommon for usage and error information.
func (c *CLB) EnableAccessLogCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.EnableAccessLogCommonRequest(input)
	return out, req.Send()
}

// EnableAccessLogCommonWithContext is the same as EnableAccessLogCommon with the addition of
// the ability to pass a context and additional request options.
//
// See EnableAccessLogCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) EnableAccessLogCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.EnableAccessLogCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opEnableAccessLog = "EnableAccessLog"

// EnableAccessLogRequest generates a "volcengine/request.Request" representing the
// client's request for the EnableAccessLog operation. The "output" return
// value will be populated with the EnableAccessLogCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned EnableAccessLogCommon Request to send the API call to the service.
// the "output" return value is not valid until after EnableAccessLogCommon Send returns without error.
//
// See EnableAccessLog for more information on using the EnableAccessLog
// API call, and error handling.
//
//    // Example sending a request using the EnableAccessLogRequest method.
//    req, resp := client.EnableAccessLogRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) EnableAccessLogRequest(input *EnableAccessLogInput) (req *request.Request, output *EnableAccessLogOutput) {
	op := &request.Operation{
		Name:       opEnableAccessLog,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &EnableAccessLogInput{}
	}

	output = &EnableAccessLogOutput{}
	req = c.newRequest(op, input, output)

	return
}

// EnableAccessLog API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation EnableAccessLog for usage and error information.
func (c *CLB) EnableAccessLog(input *EnableAccessLogInput) (*EnableAccessLogOutput, error) {
	req, out := c.EnableAccessLogRequest(input)
	return out, req.Send()
}

// EnableAccessLogWithContext is the same as EnableAccessLog with the addition of
// the ability to pass a context and additional request options.
//
// See EnableAccessLog for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) EnableAccessLogWithContext(ctx volcengine.Context, input *EnableAccessLogInput, opts ...request.Option) (*EnableAccessLogOutput, error) {
	req, out := c.EnableAccessLogRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type EnableAccessLogInput struct {
	_ struct{} `type:"structure"`

	// BucketName is a required field
	BucketName *string `type:"string" required:"true"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s EnableAccessLogInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableAccessLogInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *EnableAccessLogInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "EnableAccessLogInput"}
	if s.BucketName == nil {
		invalidParams.Add(request.NewErrParamRequired("BucketName"))
	}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBucketName sets the BucketName field's value.
func (s *EnableAccessLogInput) SetBucketName(v string) *EnableAccessLogInput {
	s.BucketName = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *EnableAccessLogInput) SetLoadBalancerId(v string) *EnableAccessLogInput {
	s.LoadBalancerId = &v
	return s
}

type EnableAccessLogOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s EnableAccessLogOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EnableAccessLogOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *EnableAccessLogOutput) SetRequestId(v string) *EnableAccessLogOutput {
	s.RequestId = &v
	return s
}
