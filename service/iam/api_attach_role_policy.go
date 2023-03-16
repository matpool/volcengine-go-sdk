// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAttachRolePolicyCommon = "AttachRolePolicy"

// AttachRolePolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachRolePolicyCommon operation. The "output" return
// value will be populated with the AttachRolePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachRolePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachRolePolicyCommon Send returns without error.
//
// See AttachRolePolicyCommon for more information on using the AttachRolePolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the AttachRolePolicyCommonRequest method.
//    req, resp := client.AttachRolePolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) AttachRolePolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAttachRolePolicyCommon,
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

// AttachRolePolicyCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AttachRolePolicyCommon for usage and error information.
func (c *IAM) AttachRolePolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AttachRolePolicyCommonRequest(input)
	return out, req.Send()
}

// AttachRolePolicyCommonWithContext is the same as AttachRolePolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AttachRolePolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AttachRolePolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AttachRolePolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAttachRolePolicy = "AttachRolePolicy"

// AttachRolePolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the AttachRolePolicy operation. The "output" return
// value will be populated with the AttachRolePolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AttachRolePolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after AttachRolePolicyCommon Send returns without error.
//
// See AttachRolePolicy for more information on using the AttachRolePolicy
// API call, and error handling.
//
//    // Example sending a request using the AttachRolePolicyRequest method.
//    req, resp := client.AttachRolePolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) AttachRolePolicyRequest(input *AttachRolePolicyInput) (req *request.Request, output *AttachRolePolicyOutput) {
	op := &request.Operation{
		Name:       opAttachRolePolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AttachRolePolicyInput{}
	}

	output = &AttachRolePolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AttachRolePolicy API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation AttachRolePolicy for usage and error information.
func (c *IAM) AttachRolePolicy(input *AttachRolePolicyInput) (*AttachRolePolicyOutput, error) {
	req, out := c.AttachRolePolicyRequest(input)
	return out, req.Send()
}

// AttachRolePolicyWithContext is the same as AttachRolePolicy with the addition of
// the ability to pass a context and additional request options.
//
// See AttachRolePolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) AttachRolePolicyWithContext(ctx volcengine.Context, input *AttachRolePolicyInput, opts ...request.Option) (*AttachRolePolicyOutput, error) {
	req, out := c.AttachRolePolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AttachRolePolicyInput struct {
	_ struct{} `type:"structure"`

	// PolicyName is a required field
	PolicyName *string `type:"string" required:"true"`

	// PolicyType is a required field
	PolicyType *string `type:"string" required:"true"`

	// RoleName is a required field
	RoleName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AttachRolePolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachRolePolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachRolePolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AttachRolePolicyInput"}
	if s.PolicyName == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyType == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyType"))
	}
	if s.RoleName == nil {
		invalidParams.Add(request.NewErrParamRequired("RoleName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPolicyName sets the PolicyName field's value.
func (s *AttachRolePolicyInput) SetPolicyName(v string) *AttachRolePolicyInput {
	s.PolicyName = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *AttachRolePolicyInput) SetPolicyType(v string) *AttachRolePolicyInput {
	s.PolicyType = &v
	return s
}

// SetRoleName sets the RoleName field's value.
func (s *AttachRolePolicyInput) SetRoleName(v string) *AttachRolePolicyInput {
	s.RoleName = &v
	return s
}

type AttachRolePolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s AttachRolePolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AttachRolePolicyOutput) GoString() string {
	return s.String()
}
