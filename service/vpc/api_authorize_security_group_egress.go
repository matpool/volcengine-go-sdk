// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAuthorizeSecurityGroupEgressCommon = "AuthorizeSecurityGroupEgress"

// AuthorizeSecurityGroupEgressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AuthorizeSecurityGroupEgressCommon operation. The "output" return
// value will be populated with the AuthorizeSecurityGroupEgressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AuthorizeSecurityGroupEgressCommon Request to send the API call to the service.
// the "output" return value is not valid until after AuthorizeSecurityGroupEgressCommon Send returns without error.
//
// See AuthorizeSecurityGroupEgressCommon for more information on using the AuthorizeSecurityGroupEgressCommon
// API call, and error handling.
//
//    // Example sending a request using the AuthorizeSecurityGroupEgressCommonRequest method.
//    req, resp := client.AuthorizeSecurityGroupEgressCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) AuthorizeSecurityGroupEgressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAuthorizeSecurityGroupEgressCommon,
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

// AuthorizeSecurityGroupEgressCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation AuthorizeSecurityGroupEgressCommon for usage and error information.
func (c *VPC) AuthorizeSecurityGroupEgressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AuthorizeSecurityGroupEgressCommonRequest(input)
	return out, req.Send()
}

// AuthorizeSecurityGroupEgressCommonWithContext is the same as AuthorizeSecurityGroupEgressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AuthorizeSecurityGroupEgressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AuthorizeSecurityGroupEgressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AuthorizeSecurityGroupEgressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAuthorizeSecurityGroupEgress = "AuthorizeSecurityGroupEgress"

// AuthorizeSecurityGroupEgressRequest generates a "volcengine/request.Request" representing the
// client's request for the AuthorizeSecurityGroupEgress operation. The "output" return
// value will be populated with the AuthorizeSecurityGroupEgressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AuthorizeSecurityGroupEgressCommon Request to send the API call to the service.
// the "output" return value is not valid until after AuthorizeSecurityGroupEgressCommon Send returns without error.
//
// See AuthorizeSecurityGroupEgress for more information on using the AuthorizeSecurityGroupEgress
// API call, and error handling.
//
//    // Example sending a request using the AuthorizeSecurityGroupEgressRequest method.
//    req, resp := client.AuthorizeSecurityGroupEgressRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) AuthorizeSecurityGroupEgressRequest(input *AuthorizeSecurityGroupEgressInput) (req *request.Request, output *AuthorizeSecurityGroupEgressOutput) {
	op := &request.Operation{
		Name:       opAuthorizeSecurityGroupEgress,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AuthorizeSecurityGroupEgressInput{}
	}

	output = &AuthorizeSecurityGroupEgressOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AuthorizeSecurityGroupEgress API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation AuthorizeSecurityGroupEgress for usage and error information.
func (c *VPC) AuthorizeSecurityGroupEgress(input *AuthorizeSecurityGroupEgressInput) (*AuthorizeSecurityGroupEgressOutput, error) {
	req, out := c.AuthorizeSecurityGroupEgressRequest(input)
	return out, req.Send()
}

// AuthorizeSecurityGroupEgressWithContext is the same as AuthorizeSecurityGroupEgress with the addition of
// the ability to pass a context and additional request options.
//
// See AuthorizeSecurityGroupEgress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AuthorizeSecurityGroupEgressWithContext(ctx volcengine.Context, input *AuthorizeSecurityGroupEgressInput, opts ...request.Option) (*AuthorizeSecurityGroupEgressOutput, error) {
	req, out := c.AuthorizeSecurityGroupEgressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AuthorizeSecurityGroupEgressInput struct {
	_ struct{} `type:"structure"`

	// CidrIp is a required field
	CidrIp *string `type:"string" required:"true"`

	Description *string `type:"string"`

	Policy *string `type:"string"`

	// PortEnd is a required field
	PortEnd *int64 `type:"integer" required:"true"`

	// PortStart is a required field
	PortStart *int64 `type:"integer" required:"true"`

	Priority *int64 `type:"integer"`

	// Protocol is a required field
	Protocol *string `type:"string" required:"true"`

	// SecurityGroupId is a required field
	SecurityGroupId *string `type:"string" required:"true"`

	SourceGroupId *string `type:"string"`
}

// String returns the string representation
func (s AuthorizeSecurityGroupEgressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthorizeSecurityGroupEgressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AuthorizeSecurityGroupEgressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AuthorizeSecurityGroupEgressInput"}
	if s.CidrIp == nil {
		invalidParams.Add(request.NewErrParamRequired("CidrIp"))
	}
	if s.PortEnd == nil {
		invalidParams.Add(request.NewErrParamRequired("PortEnd"))
	}
	if s.PortStart == nil {
		invalidParams.Add(request.NewErrParamRequired("PortStart"))
	}
	if s.Protocol == nil {
		invalidParams.Add(request.NewErrParamRequired("Protocol"))
	}
	if s.SecurityGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("SecurityGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCidrIp sets the CidrIp field's value.
func (s *AuthorizeSecurityGroupEgressInput) SetCidrIp(v string) *AuthorizeSecurityGroupEgressInput {
	s.CidrIp = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *AuthorizeSecurityGroupEgressInput) SetDescription(v string) *AuthorizeSecurityGroupEgressInput {
	s.Description = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *AuthorizeSecurityGroupEgressInput) SetPolicy(v string) *AuthorizeSecurityGroupEgressInput {
	s.Policy = &v
	return s
}

// SetPortEnd sets the PortEnd field's value.
func (s *AuthorizeSecurityGroupEgressInput) SetPortEnd(v int64) *AuthorizeSecurityGroupEgressInput {
	s.PortEnd = &v
	return s
}

// SetPortStart sets the PortStart field's value.
func (s *AuthorizeSecurityGroupEgressInput) SetPortStart(v int64) *AuthorizeSecurityGroupEgressInput {
	s.PortStart = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *AuthorizeSecurityGroupEgressInput) SetPriority(v int64) *AuthorizeSecurityGroupEgressInput {
	s.Priority = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *AuthorizeSecurityGroupEgressInput) SetProtocol(v string) *AuthorizeSecurityGroupEgressInput {
	s.Protocol = &v
	return s
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *AuthorizeSecurityGroupEgressInput) SetSecurityGroupId(v string) *AuthorizeSecurityGroupEgressInput {
	s.SecurityGroupId = &v
	return s
}

// SetSourceGroupId sets the SourceGroupId field's value.
func (s *AuthorizeSecurityGroupEgressInput) SetSourceGroupId(v string) *AuthorizeSecurityGroupEgressInput {
	s.SourceGroupId = &v
	return s
}

type AuthorizeSecurityGroupEgressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AuthorizeSecurityGroupEgressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AuthorizeSecurityGroupEgressOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *AuthorizeSecurityGroupEgressOutput) SetRequestId(v string) *AuthorizeSecurityGroupEgressOutput {
	s.RequestId = &v
	return s
}
