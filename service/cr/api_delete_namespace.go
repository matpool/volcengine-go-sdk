// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteNamespaceCommon = "DeleteNamespace"

// DeleteNamespaceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteNamespaceCommon operation. The "output" return
// value will be populated with the DeleteNamespaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNamespaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNamespaceCommon Send returns without error.
//
// See DeleteNamespaceCommon for more information on using the DeleteNamespaceCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteNamespaceCommonRequest method.
//    req, resp := client.DeleteNamespaceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) DeleteNamespaceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteNamespaceCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteNamespaceCommon API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation DeleteNamespaceCommon for usage and error information.
func (c *CR) DeleteNamespaceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteNamespaceCommonRequest(input)
	return out, req.Send()
}

// DeleteNamespaceCommonWithContext is the same as DeleteNamespaceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNamespaceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) DeleteNamespaceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteNamespaceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteNamespace = "DeleteNamespace"

// DeleteNamespaceRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteNamespace operation. The "output" return
// value will be populated with the DeleteNamespaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteNamespaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteNamespaceCommon Send returns without error.
//
// See DeleteNamespace for more information on using the DeleteNamespace
// API call, and error handling.
//
//    // Example sending a request using the DeleteNamespaceRequest method.
//    req, resp := client.DeleteNamespaceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) DeleteNamespaceRequest(input *DeleteNamespaceInput) (req *request.Request, output *DeleteNamespaceOutput) {
	op := &request.Operation{
		Name:       opDeleteNamespace,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteNamespaceInput{}
	}

	output = &DeleteNamespaceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteNamespace API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation DeleteNamespace for usage and error information.
func (c *CR) DeleteNamespace(input *DeleteNamespaceInput) (*DeleteNamespaceOutput, error) {
	req, out := c.DeleteNamespaceRequest(input)
	return out, req.Send()
}

// DeleteNamespaceWithContext is the same as DeleteNamespace with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteNamespace for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) DeleteNamespaceWithContext(ctx volcengine.Context, input *DeleteNamespaceInput, opts ...request.Option) (*DeleteNamespaceOutput, error) {
	req, out := c.DeleteNamespaceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteNamespaceInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Registry *string `type:"string"`
}

// String returns the string representation
func (s DeleteNamespaceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNamespaceInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DeleteNamespaceInput) SetName(v string) *DeleteNamespaceInput {
	s.Name = &v
	return s
}

// SetRegistry sets the Registry field's value.
func (s *DeleteNamespaceInput) SetRegistry(v string) *DeleteNamespaceInput {
	s.Registry = &v
	return s
}

type DeleteNamespaceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteNamespaceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteNamespaceOutput) GoString() string {
	return s.String()
}
