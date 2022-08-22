// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteRepositoryCommon = "DeleteRepository"

// DeleteRepositoryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteRepositoryCommon operation. The "output" return
// value will be populated with the DeleteRepositoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteRepositoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteRepositoryCommon Send returns without error.
//
// See DeleteRepositoryCommon for more information on using the DeleteRepositoryCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteRepositoryCommonRequest method.
//    req, resp := client.DeleteRepositoryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) DeleteRepositoryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteRepositoryCommon,
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

// DeleteRepositoryCommon API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation DeleteRepositoryCommon for usage and error information.
func (c *CR) DeleteRepositoryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteRepositoryCommonRequest(input)
	return out, req.Send()
}

// DeleteRepositoryCommonWithContext is the same as DeleteRepositoryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteRepositoryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) DeleteRepositoryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteRepositoryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteRepository = "DeleteRepository"

// DeleteRepositoryRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteRepository operation. The "output" return
// value will be populated with the DeleteRepositoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteRepositoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteRepositoryCommon Send returns without error.
//
// See DeleteRepository for more information on using the DeleteRepository
// API call, and error handling.
//
//    // Example sending a request using the DeleteRepositoryRequest method.
//    req, resp := client.DeleteRepositoryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) DeleteRepositoryRequest(input *DeleteRepositoryInput) (req *request.Request, output *DeleteRepositoryOutput) {
	op := &request.Operation{
		Name:       opDeleteRepository,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteRepositoryInput{}
	}

	output = &DeleteRepositoryOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteRepository API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation DeleteRepository for usage and error information.
func (c *CR) DeleteRepository(input *DeleteRepositoryInput) (*DeleteRepositoryOutput, error) {
	req, out := c.DeleteRepositoryRequest(input)
	return out, req.Send()
}

// DeleteRepositoryWithContext is the same as DeleteRepository with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteRepository for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) DeleteRepositoryWithContext(ctx volcengine.Context, input *DeleteRepositoryInput, opts ...request.Option) (*DeleteRepositoryOutput, error) {
	req, out := c.DeleteRepositoryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteRepositoryInput struct {
	_ struct{} `type:"structure"`

	Name *string `type:"string"`

	Namespace *string `type:"string"`

	Registry *string `type:"string"`
}

// String returns the string representation
func (s DeleteRepositoryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteRepositoryInput) GoString() string {
	return s.String()
}

// SetName sets the Name field's value.
func (s *DeleteRepositoryInput) SetName(v string) *DeleteRepositoryInput {
	s.Name = &v
	return s
}

// SetNamespace sets the Namespace field's value.
func (s *DeleteRepositoryInput) SetNamespace(v string) *DeleteRepositoryInput {
	s.Namespace = &v
	return s
}

// SetRegistry sets the Registry field's value.
func (s *DeleteRepositoryInput) SetRegistry(v string) *DeleteRepositoryInput {
	s.Registry = &v
	return s
}

type DeleteRepositoryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteRepositoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteRepositoryOutput) GoString() string {
	return s.String()
}
