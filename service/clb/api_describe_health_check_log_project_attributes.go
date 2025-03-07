// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeHealthCheckLogProjectAttributesCommon = "DescribeHealthCheckLogProjectAttributes"

// DescribeHealthCheckLogProjectAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeHealthCheckLogProjectAttributesCommon operation. The "output" return
// value will be populated with the DescribeHealthCheckLogProjectAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeHealthCheckLogProjectAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeHealthCheckLogProjectAttributesCommon Send returns without error.
//
// See DescribeHealthCheckLogProjectAttributesCommon for more information on using the DescribeHealthCheckLogProjectAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeHealthCheckLogProjectAttributesCommonRequest method.
//	req, resp := client.DescribeHealthCheckLogProjectAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeHealthCheckLogProjectAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeHealthCheckLogProjectAttributesCommon,
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

// DescribeHealthCheckLogProjectAttributesCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeHealthCheckLogProjectAttributesCommon for usage and error information.
func (c *CLB) DescribeHealthCheckLogProjectAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeHealthCheckLogProjectAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeHealthCheckLogProjectAttributesCommonWithContext is the same as DescribeHealthCheckLogProjectAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeHealthCheckLogProjectAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeHealthCheckLogProjectAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeHealthCheckLogProjectAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeHealthCheckLogProjectAttributes = "DescribeHealthCheckLogProjectAttributes"

// DescribeHealthCheckLogProjectAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeHealthCheckLogProjectAttributes operation. The "output" return
// value will be populated with the DescribeHealthCheckLogProjectAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeHealthCheckLogProjectAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeHealthCheckLogProjectAttributesCommon Send returns without error.
//
// See DescribeHealthCheckLogProjectAttributes for more information on using the DescribeHealthCheckLogProjectAttributes
// API call, and error handling.
//
//	// Example sending a request using the DescribeHealthCheckLogProjectAttributesRequest method.
//	req, resp := client.DescribeHealthCheckLogProjectAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeHealthCheckLogProjectAttributesRequest(input *DescribeHealthCheckLogProjectAttributesInput) (req *request.Request, output *DescribeHealthCheckLogProjectAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeHealthCheckLogProjectAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeHealthCheckLogProjectAttributesInput{}
	}

	output = &DescribeHealthCheckLogProjectAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeHealthCheckLogProjectAttributes API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeHealthCheckLogProjectAttributes for usage and error information.
func (c *CLB) DescribeHealthCheckLogProjectAttributes(input *DescribeHealthCheckLogProjectAttributesInput) (*DescribeHealthCheckLogProjectAttributesOutput, error) {
	req, out := c.DescribeHealthCheckLogProjectAttributesRequest(input)
	return out, req.Send()
}

// DescribeHealthCheckLogProjectAttributesWithContext is the same as DescribeHealthCheckLogProjectAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeHealthCheckLogProjectAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeHealthCheckLogProjectAttributesWithContext(ctx volcengine.Context, input *DescribeHealthCheckLogProjectAttributesInput, opts ...request.Option) (*DescribeHealthCheckLogProjectAttributesOutput, error) {
	req, out := c.DescribeHealthCheckLogProjectAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeHealthCheckLogProjectAttributesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeHealthCheckLogProjectAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHealthCheckLogProjectAttributesInput) GoString() string {
	return s.String()
}

type DescribeHealthCheckLogProjectAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LogProjectId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DescribeHealthCheckLogProjectAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeHealthCheckLogProjectAttributesOutput) GoString() string {
	return s.String()
}

// SetLogProjectId sets the LogProjectId field's value.
func (s *DescribeHealthCheckLogProjectAttributesOutput) SetLogProjectId(v string) *DescribeHealthCheckLogProjectAttributesOutput {
	s.LogProjectId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeHealthCheckLogProjectAttributesOutput) SetRequestId(v string) *DescribeHealthCheckLogProjectAttributesOutput {
	s.RequestId = &v
	return s
}
