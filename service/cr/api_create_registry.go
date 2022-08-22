// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cr

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateRegistryCommon = "CreateRegistry"

// CreateRegistryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateRegistryCommon operation. The "output" return
// value will be populated with the CreateRegistryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateRegistryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateRegistryCommon Send returns without error.
//
// See CreateRegistryCommon for more information on using the CreateRegistryCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateRegistryCommonRequest method.
//    req, resp := client.CreateRegistryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) CreateRegistryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateRegistryCommon,
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

// CreateRegistryCommon API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation CreateRegistryCommon for usage and error information.
func (c *CR) CreateRegistryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateRegistryCommonRequest(input)
	return out, req.Send()
}

// CreateRegistryCommonWithContext is the same as CreateRegistryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateRegistryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) CreateRegistryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateRegistryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateRegistry = "CreateRegistry"

// CreateRegistryRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateRegistry operation. The "output" return
// value will be populated with the CreateRegistryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateRegistryCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateRegistryCommon Send returns without error.
//
// See CreateRegistry for more information on using the CreateRegistry
// API call, and error handling.
//
//    // Example sending a request using the CreateRegistryRequest method.
//    req, resp := client.CreateRegistryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CR) CreateRegistryRequest(input *CreateRegistryInput) (req *request.Request, output *CreateRegistryOutput) {
	op := &request.Operation{
		Name:       opCreateRegistry,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateRegistryInput{}
	}

	output = &CreateRegistryOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateRegistry API operation for CR.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CR's
// API operation CreateRegistry for usage and error information.
func (c *CR) CreateRegistry(input *CreateRegistryInput) (*CreateRegistryOutput, error) {
	req, out := c.CreateRegistryRequest(input)
	return out, req.Send()
}

// CreateRegistryWithContext is the same as CreateRegistry with the addition of
// the ability to pass a context and additional request options.
//
// See CreateRegistry for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CR) CreateRegistryWithContext(ctx volcengine.Context, input *CreateRegistryInput, opts ...request.Option) (*CreateRegistryOutput, error) {
	req, out := c.CreateRegistryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateRegistryInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s CreateRegistryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateRegistryInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateRegistryInput) SetClientToken(v string) *CreateRegistryInput {
	s.ClientToken = &v
	return s
}

// SetName sets the Name field's value.
func (s *CreateRegistryInput) SetName(v string) *CreateRegistryInput {
	s.Name = &v
	return s
}

type CreateRegistryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateRegistryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateRegistryOutput) GoString() string {
	return s.String()
}
