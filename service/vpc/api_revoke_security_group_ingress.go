// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRevokeSecurityGroupIngressCommon = "RevokeSecurityGroupIngress"

// RevokeSecurityGroupIngressCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RevokeSecurityGroupIngressCommon operation. The "output" return
// value will be populated with the RevokeSecurityGroupIngressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RevokeSecurityGroupIngressCommon Request to send the API call to the service.
// the "output" return value is not valid until after RevokeSecurityGroupIngressCommon Send returns without error.
//
// See RevokeSecurityGroupIngressCommon for more information on using the RevokeSecurityGroupIngressCommon
// API call, and error handling.
//
//	// Example sending a request using the RevokeSecurityGroupIngressCommonRequest method.
//	req, resp := client.RevokeSecurityGroupIngressCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) RevokeSecurityGroupIngressCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRevokeSecurityGroupIngressCommon,
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

// RevokeSecurityGroupIngressCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation RevokeSecurityGroupIngressCommon for usage and error information.
func (c *VPC) RevokeSecurityGroupIngressCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RevokeSecurityGroupIngressCommonRequest(input)
	return out, req.Send()
}

// RevokeSecurityGroupIngressCommonWithContext is the same as RevokeSecurityGroupIngressCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RevokeSecurityGroupIngressCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) RevokeSecurityGroupIngressCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RevokeSecurityGroupIngressCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRevokeSecurityGroupIngress = "RevokeSecurityGroupIngress"

// RevokeSecurityGroupIngressRequest generates a "volcengine/request.Request" representing the
// client's request for the RevokeSecurityGroupIngress operation. The "output" return
// value will be populated with the RevokeSecurityGroupIngressCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RevokeSecurityGroupIngressCommon Request to send the API call to the service.
// the "output" return value is not valid until after RevokeSecurityGroupIngressCommon Send returns without error.
//
// See RevokeSecurityGroupIngress for more information on using the RevokeSecurityGroupIngress
// API call, and error handling.
//
//	// Example sending a request using the RevokeSecurityGroupIngressRequest method.
//	req, resp := client.RevokeSecurityGroupIngressRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) RevokeSecurityGroupIngressRequest(input *RevokeSecurityGroupIngressInput) (req *request.Request, output *RevokeSecurityGroupIngressOutput) {
	op := &request.Operation{
		Name:       opRevokeSecurityGroupIngress,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RevokeSecurityGroupIngressInput{}
	}

	output = &RevokeSecurityGroupIngressOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RevokeSecurityGroupIngress API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation RevokeSecurityGroupIngress for usage and error information.
func (c *VPC) RevokeSecurityGroupIngress(input *RevokeSecurityGroupIngressInput) (*RevokeSecurityGroupIngressOutput, error) {
	req, out := c.RevokeSecurityGroupIngressRequest(input)
	return out, req.Send()
}

// RevokeSecurityGroupIngressWithContext is the same as RevokeSecurityGroupIngress with the addition of
// the ability to pass a context and additional request options.
//
// See RevokeSecurityGroupIngress for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) RevokeSecurityGroupIngressWithContext(ctx volcengine.Context, input *RevokeSecurityGroupIngressInput, opts ...request.Option) (*RevokeSecurityGroupIngressOutput, error) {
	req, out := c.RevokeSecurityGroupIngressRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RevokeSecurityGroupIngressInput struct {
	_ struct{} `type:"structure"`

	CidrIp *string `type:"string"`

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
func (s RevokeSecurityGroupIngressInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RevokeSecurityGroupIngressInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RevokeSecurityGroupIngressInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RevokeSecurityGroupIngressInput"}
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
func (s *RevokeSecurityGroupIngressInput) SetCidrIp(v string) *RevokeSecurityGroupIngressInput {
	s.CidrIp = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *RevokeSecurityGroupIngressInput) SetDescription(v string) *RevokeSecurityGroupIngressInput {
	s.Description = &v
	return s
}

// SetPolicy sets the Policy field's value.
func (s *RevokeSecurityGroupIngressInput) SetPolicy(v string) *RevokeSecurityGroupIngressInput {
	s.Policy = &v
	return s
}

// SetPortEnd sets the PortEnd field's value.
func (s *RevokeSecurityGroupIngressInput) SetPortEnd(v int64) *RevokeSecurityGroupIngressInput {
	s.PortEnd = &v
	return s
}

// SetPortStart sets the PortStart field's value.
func (s *RevokeSecurityGroupIngressInput) SetPortStart(v int64) *RevokeSecurityGroupIngressInput {
	s.PortStart = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *RevokeSecurityGroupIngressInput) SetPriority(v int64) *RevokeSecurityGroupIngressInput {
	s.Priority = &v
	return s
}

// SetProtocol sets the Protocol field's value.
func (s *RevokeSecurityGroupIngressInput) SetProtocol(v string) *RevokeSecurityGroupIngressInput {
	s.Protocol = &v
	return s
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *RevokeSecurityGroupIngressInput) SetSecurityGroupId(v string) *RevokeSecurityGroupIngressInput {
	s.SecurityGroupId = &v
	return s
}

// SetSourceGroupId sets the SourceGroupId field's value.
func (s *RevokeSecurityGroupIngressInput) SetSourceGroupId(v string) *RevokeSecurityGroupIngressInput {
	s.SourceGroupId = &v
	return s
}

type RevokeSecurityGroupIngressOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s RevokeSecurityGroupIngressOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RevokeSecurityGroupIngressOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *RevokeSecurityGroupIngressOutput) SetRequestId(v string) *RevokeSecurityGroupIngressOutput {
	s.RequestId = &v
	return s
}
