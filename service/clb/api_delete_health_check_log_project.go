// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteHealthCheckLogProjectCommon = "DeleteHealthCheckLogProject"

// DeleteHealthCheckLogProjectCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteHealthCheckLogProjectCommon operation. The "output" return
// value will be populated with the DeleteHealthCheckLogProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteHealthCheckLogProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteHealthCheckLogProjectCommon Send returns without error.
//
// See DeleteHealthCheckLogProjectCommon for more information on using the DeleteHealthCheckLogProjectCommon
// API call, and error handling.
//
//	// Example sending a request using the DeleteHealthCheckLogProjectCommonRequest method.
//	req, resp := client.DeleteHealthCheckLogProjectCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DeleteHealthCheckLogProjectCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteHealthCheckLogProjectCommon,
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

// DeleteHealthCheckLogProjectCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DeleteHealthCheckLogProjectCommon for usage and error information.
func (c *CLB) DeleteHealthCheckLogProjectCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteHealthCheckLogProjectCommonRequest(input)
	return out, req.Send()
}

// DeleteHealthCheckLogProjectCommonWithContext is the same as DeleteHealthCheckLogProjectCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteHealthCheckLogProjectCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DeleteHealthCheckLogProjectCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteHealthCheckLogProjectCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteHealthCheckLogProject = "DeleteHealthCheckLogProject"

// DeleteHealthCheckLogProjectRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteHealthCheckLogProject operation. The "output" return
// value will be populated with the DeleteHealthCheckLogProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteHealthCheckLogProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteHealthCheckLogProjectCommon Send returns without error.
//
// See DeleteHealthCheckLogProject for more information on using the DeleteHealthCheckLogProject
// API call, and error handling.
//
//	// Example sending a request using the DeleteHealthCheckLogProjectRequest method.
//	req, resp := client.DeleteHealthCheckLogProjectRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DeleteHealthCheckLogProjectRequest(input *DeleteHealthCheckLogProjectInput) (req *request.Request, output *DeleteHealthCheckLogProjectOutput) {
	op := &request.Operation{
		Name:       opDeleteHealthCheckLogProject,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteHealthCheckLogProjectInput{}
	}

	output = &DeleteHealthCheckLogProjectOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteHealthCheckLogProject API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DeleteHealthCheckLogProject for usage and error information.
func (c *CLB) DeleteHealthCheckLogProject(input *DeleteHealthCheckLogProjectInput) (*DeleteHealthCheckLogProjectOutput, error) {
	req, out := c.DeleteHealthCheckLogProjectRequest(input)
	return out, req.Send()
}

// DeleteHealthCheckLogProjectWithContext is the same as DeleteHealthCheckLogProject with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteHealthCheckLogProject for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DeleteHealthCheckLogProjectWithContext(ctx volcengine.Context, input *DeleteHealthCheckLogProjectInput, opts ...request.Option) (*DeleteHealthCheckLogProjectOutput, error) {
	req, out := c.DeleteHealthCheckLogProjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteHealthCheckLogProjectInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteHealthCheckLogProjectInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHealthCheckLogProjectInput) GoString() string {
	return s.String()
}

type DeleteHealthCheckLogProjectOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteHealthCheckLogProjectOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteHealthCheckLogProjectOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteHealthCheckLogProjectOutput) SetRequestId(v string) *DeleteHealthCheckLogProjectOutput {
	s.RequestId = &v
	return s
}
