// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteSAMLProviderCommon = "DeleteSAMLProvider"

// DeleteSAMLProviderCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteSAMLProviderCommon operation. The "output" return
// value will be populated with the DeleteSAMLProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteSAMLProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteSAMLProviderCommon Send returns without error.
//
// See DeleteSAMLProviderCommon for more information on using the DeleteSAMLProviderCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteSAMLProviderCommonRequest method.
//    req, resp := client.DeleteSAMLProviderCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DeleteSAMLProviderCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteSAMLProviderCommon,
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

// DeleteSAMLProviderCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DeleteSAMLProviderCommon for usage and error information.
func (c *IAM) DeleteSAMLProviderCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteSAMLProviderCommonRequest(input)
	return out, req.Send()
}

// DeleteSAMLProviderCommonWithContext is the same as DeleteSAMLProviderCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteSAMLProviderCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeleteSAMLProviderCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteSAMLProviderCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteSAMLProvider = "DeleteSAMLProvider"

// DeleteSAMLProviderRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteSAMLProvider operation. The "output" return
// value will be populated with the DeleteSAMLProviderCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteSAMLProviderCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteSAMLProviderCommon Send returns without error.
//
// See DeleteSAMLProvider for more information on using the DeleteSAMLProvider
// API call, and error handling.
//
//    // Example sending a request using the DeleteSAMLProviderRequest method.
//    req, resp := client.DeleteSAMLProviderRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) DeleteSAMLProviderRequest(input *DeleteSAMLProviderInput) (req *request.Request, output *DeleteSAMLProviderOutput) {
	op := &request.Operation{
		Name:       opDeleteSAMLProvider,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteSAMLProviderInput{}
	}

	output = &DeleteSAMLProviderOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteSAMLProvider API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation DeleteSAMLProvider for usage and error information.
func (c *IAM) DeleteSAMLProvider(input *DeleteSAMLProviderInput) (*DeleteSAMLProviderOutput, error) {
	req, out := c.DeleteSAMLProviderRequest(input)
	return out, req.Send()
}

// DeleteSAMLProviderWithContext is the same as DeleteSAMLProvider with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteSAMLProvider for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) DeleteSAMLProviderWithContext(ctx volcengine.Context, input *DeleteSAMLProviderInput, opts ...request.Option) (*DeleteSAMLProviderOutput, error) {
	req, out := c.DeleteSAMLProviderRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteSAMLProviderInput struct {
	_ struct{} `type:"structure"`

	// SAMLProviderName is a required field
	SAMLProviderName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteSAMLProviderInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteSAMLProviderInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteSAMLProviderInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteSAMLProviderInput"}
	if s.SAMLProviderName == nil {
		invalidParams.Add(request.NewErrParamRequired("SAMLProviderName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetSAMLProviderName sets the SAMLProviderName field's value.
func (s *DeleteSAMLProviderInput) SetSAMLProviderName(v string) *DeleteSAMLProviderInput {
	s.SAMLProviderName = &v
	return s
}

type DeleteSAMLProviderOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteSAMLProviderOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteSAMLProviderOutput) GoString() string {
	return s.String()
}
