// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetPolicyCommon = "GetPolicy"

// GetPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetPolicyCommon operation. The "output" return
// value will be populated with the GetPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetPolicyCommon Send returns without error.
//
// See GetPolicyCommon for more information on using the GetPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the GetPolicyCommonRequest method.
//    req, resp := client.GetPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetPolicyCommon,
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

// GetPolicyCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetPolicyCommon for usage and error information.
func (c *IAM) GetPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetPolicyCommonRequest(input)
	return out, req.Send()
}

// GetPolicyCommonWithContext is the same as GetPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetPolicy = "GetPolicy"

// GetPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the GetPolicy operation. The "output" return
// value will be populated with the GetPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetPolicyCommon Send returns without error.
//
// See GetPolicy for more information on using the GetPolicy
// API call, and error handling.
//
//    // Example sending a request using the GetPolicyRequest method.
//    req, resp := client.GetPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetPolicyRequest(input *GetPolicyInput) (req *request.Request, output *GetPolicyOutput) {
	op := &request.Operation{
		Name:       opGetPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetPolicyInput{}
	}

	output = &GetPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetPolicy API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetPolicy for usage and error information.
func (c *IAM) GetPolicy(input *GetPolicyInput) (*GetPolicyOutput, error) {
	req, out := c.GetPolicyRequest(input)
	return out, req.Send()
}

// GetPolicyWithContext is the same as GetPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See GetPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetPolicyWithContext(ctx volcengine.Context, input *GetPolicyInput, opts ...request.Option) (*GetPolicyOutput, error) {
	req, out := c.GetPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetPolicyInput struct {
	_ struct{} `type:"structure"`

	// PolicyName is a required field
	PolicyName *string `type:"string" required:"true"`

	// PolicyType is a required field
	PolicyType *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPolicyInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPolicyInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetPolicyInput"}
	if s.PolicyName == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyName"))
	}
	if s.PolicyType == nil {
		invalidParams.Add(request.NewErrParamRequired("PolicyType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPolicyName sets the PolicyName field's value.
func (s *GetPolicyInput) SetPolicyName(v string) *GetPolicyInput {
	s.PolicyName = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *GetPolicyInput) SetPolicyType(v string) *GetPolicyInput {
	s.PolicyType = &v
	return s
}

type GetPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Policy *PolicyForGetPolicyOutput `type:"structure"`
}

// String returns the string representation
func (s GetPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPolicyOutput) GoString() string {
	return s.String()
}

// SetPolicy sets the Policy field's value.
func (s *GetPolicyOutput) SetPolicy(v *PolicyForGetPolicyOutput) *GetPolicyOutput {
	s.Policy = v
	return s
}

type PolicyForGetPolicyOutput struct {
	_ struct{} `type:"structure"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	PolicyDocument *string `type:"string"`

	PolicyName *string `type:"string"`

	PolicyTrn *string `type:"string"`

	PolicyType *string `type:"string"`

	UpdateDate *string `type:"string"`
}

// String returns the string representation
func (s PolicyForGetPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PolicyForGetPolicyOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *PolicyForGetPolicyOutput) SetCreateDate(v string) *PolicyForGetPolicyOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *PolicyForGetPolicyOutput) SetDescription(v string) *PolicyForGetPolicyOutput {
	s.Description = &v
	return s
}

// SetPolicyDocument sets the PolicyDocument field's value.
func (s *PolicyForGetPolicyOutput) SetPolicyDocument(v string) *PolicyForGetPolicyOutput {
	s.PolicyDocument = &v
	return s
}

// SetPolicyName sets the PolicyName field's value.
func (s *PolicyForGetPolicyOutput) SetPolicyName(v string) *PolicyForGetPolicyOutput {
	s.PolicyName = &v
	return s
}

// SetPolicyTrn sets the PolicyTrn field's value.
func (s *PolicyForGetPolicyOutput) SetPolicyTrn(v string) *PolicyForGetPolicyOutput {
	s.PolicyTrn = &v
	return s
}

// SetPolicyType sets the PolicyType field's value.
func (s *PolicyForGetPolicyOutput) SetPolicyType(v string) *PolicyForGetPolicyOutput {
	s.PolicyType = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *PolicyForGetPolicyOutput) SetUpdateDate(v string) *PolicyForGetPolicyOutput {
	s.UpdateDate = &v
	return s
}
