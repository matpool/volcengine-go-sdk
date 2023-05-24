// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opForwardKubernetesApiCommon = "ForwardKubernetesApi"

// ForwardKubernetesApiCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ForwardKubernetesApiCommon operation. The "output" return
// value will be populated with the ForwardKubernetesApiCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ForwardKubernetesApiCommon Request to send the API call to the service.
// the "output" return value is not valid until after ForwardKubernetesApiCommon Send returns without error.
//
// See ForwardKubernetesApiCommon for more information on using the ForwardKubernetesApiCommon
// API call, and error handling.
//
//	// Example sending a request using the ForwardKubernetesApiCommonRequest method.
//	req, resp := client.ForwardKubernetesApiCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VKE) ForwardKubernetesApiCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opForwardKubernetesApiCommon,
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

// ForwardKubernetesApiCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ForwardKubernetesApiCommon for usage and error information.
func (c *VKE) ForwardKubernetesApiCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ForwardKubernetesApiCommonRequest(input)
	return out, req.Send()
}

// ForwardKubernetesApiCommonWithContext is the same as ForwardKubernetesApiCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ForwardKubernetesApiCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ForwardKubernetesApiCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ForwardKubernetesApiCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opForwardKubernetesApi = "ForwardKubernetesApi"

// ForwardKubernetesApiRequest generates a "volcengine/request.Request" representing the
// client's request for the ForwardKubernetesApi operation. The "output" return
// value will be populated with the ForwardKubernetesApiCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ForwardKubernetesApiCommon Request to send the API call to the service.
// the "output" return value is not valid until after ForwardKubernetesApiCommon Send returns without error.
//
// See ForwardKubernetesApi for more information on using the ForwardKubernetesApi
// API call, and error handling.
//
//	// Example sending a request using the ForwardKubernetesApiRequest method.
//	req, resp := client.ForwardKubernetesApiRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VKE) ForwardKubernetesApiRequest(input *ForwardKubernetesApiInput) (req *request.Request, output *ForwardKubernetesApiOutput) {
	op := &request.Operation{
		Name:       opForwardKubernetesApi,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ForwardKubernetesApiInput{}
	}

	output = &ForwardKubernetesApiOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ForwardKubernetesApi API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation ForwardKubernetesApi for usage and error information.
func (c *VKE) ForwardKubernetesApi(input *ForwardKubernetesApiInput) (*ForwardKubernetesApiOutput, error) {
	req, out := c.ForwardKubernetesApiRequest(input)
	return out, req.Send()
}

// ForwardKubernetesApiWithContext is the same as ForwardKubernetesApi with the addition of
// the ability to pass a context and additional request options.
//
// See ForwardKubernetesApi for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) ForwardKubernetesApiWithContext(ctx volcengine.Context, input *ForwardKubernetesApiInput, opts ...request.Option) (*ForwardKubernetesApiOutput, error) {
	req, out := c.ForwardKubernetesApiRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ForwardKubernetesApiInput struct {
	_ struct{} `type:"structure"`

	Body *string `type:"string"`

	ClusterId *string `type:"string"`

	Headers []*HeaderForForwardKubernetesApiInput `type:"list"`

	Method *string `type:"string"`

	Path *string `type:"string"`
}

// String returns the string representation
func (s ForwardKubernetesApiInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ForwardKubernetesApiInput) GoString() string {
	return s.String()
}

// SetBody sets the Body field's value.
func (s *ForwardKubernetesApiInput) SetBody(v string) *ForwardKubernetesApiInput {
	s.Body = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *ForwardKubernetesApiInput) SetClusterId(v string) *ForwardKubernetesApiInput {
	s.ClusterId = &v
	return s
}

// SetHeaders sets the Headers field's value.
func (s *ForwardKubernetesApiInput) SetHeaders(v []*HeaderForForwardKubernetesApiInput) *ForwardKubernetesApiInput {
	s.Headers = v
	return s
}

// SetMethod sets the Method field's value.
func (s *ForwardKubernetesApiInput) SetMethod(v string) *ForwardKubernetesApiInput {
	s.Method = &v
	return s
}

// SetPath sets the Path field's value.
func (s *ForwardKubernetesApiInput) SetPath(v string) *ForwardKubernetesApiInput {
	s.Path = &v
	return s
}

type ForwardKubernetesApiOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Body *string `type:"string"`

	Code *int32 `type:"int32"`
}

// String returns the string representation
func (s ForwardKubernetesApiOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ForwardKubernetesApiOutput) GoString() string {
	return s.String()
}

// SetBody sets the Body field's value.
func (s *ForwardKubernetesApiOutput) SetBody(v string) *ForwardKubernetesApiOutput {
	s.Body = &v
	return s
}

// SetCode sets the Code field's value.
func (s *ForwardKubernetesApiOutput) SetCode(v int32) *ForwardKubernetesApiOutput {
	s.Code = &v
	return s
}

type HeaderForForwardKubernetesApiInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s HeaderForForwardKubernetesApiInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HeaderForForwardKubernetesApiInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *HeaderForForwardKubernetesApiInput) SetKey(v string) *HeaderForForwardKubernetesApiInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *HeaderForForwardKubernetesApiInput) SetValue(v string) *HeaderForForwardKubernetesApiInput {
	s.Value = &v
	return s
}
