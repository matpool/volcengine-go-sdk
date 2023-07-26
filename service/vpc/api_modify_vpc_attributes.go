// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyVpcAttributesCommon = "ModifyVpcAttributes"

// ModifyVpcAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcAttributesCommon operation. The "output" return
// value will be populated with the ModifyVpcAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcAttributesCommon Send returns without error.
//
// See ModifyVpcAttributesCommon for more information on using the ModifyVpcAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyVpcAttributesCommonRequest method.
//	req, resp := client.ModifyVpcAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ModifyVpcAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyVpcAttributesCommon,
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

// ModifyVpcAttributesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyVpcAttributesCommon for usage and error information.
func (c *VPC) ModifyVpcAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyVpcAttributesCommonWithContext is the same as ModifyVpcAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyVpcAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyVpcAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyVpcAttributes = "ModifyVpcAttributes"

// ModifyVpcAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpcAttributes operation. The "output" return
// value will be populated with the ModifyVpcAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpcAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpcAttributesCommon Send returns without error.
//
// See ModifyVpcAttributes for more information on using the ModifyVpcAttributes
// API call, and error handling.
//
//	// Example sending a request using the ModifyVpcAttributesRequest method.
//	req, resp := client.ModifyVpcAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ModifyVpcAttributesRequest(input *ModifyVpcAttributesInput) (req *request.Request, output *ModifyVpcAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyVpcAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpcAttributesInput{}
	}

	output = &ModifyVpcAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyVpcAttributes API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyVpcAttributes for usage and error information.
func (c *VPC) ModifyVpcAttributes(input *ModifyVpcAttributesInput) (*ModifyVpcAttributesOutput, error) {
	req, out := c.ModifyVpcAttributesRequest(input)
	return out, req.Send()
}

// ModifyVpcAttributesWithContext is the same as ModifyVpcAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpcAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyVpcAttributesWithContext(ctx volcengine.Context, input *ModifyVpcAttributesInput, opts ...request.Option) (*ModifyVpcAttributesOutput, error) {
	req, out := c.ModifyVpcAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyVpcAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `min:"1" max:"255" type:"string"`

	DnsServers []*string `type:"list"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`

	VpcName *string `min:"2" max:"128" type:"string"`
}

// String returns the string representation
func (s ModifyVpcAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpcAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyVpcAttributesInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}
	if s.VpcName != nil && len(*s.VpcName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("VpcName", 2))
	}
	if s.VpcName != nil && len(*s.VpcName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("VpcName", 128, *s.VpcName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyVpcAttributesInput) SetDescription(v string) *ModifyVpcAttributesInput {
	s.Description = &v
	return s
}

// SetDnsServers sets the DnsServers field's value.
func (s *ModifyVpcAttributesInput) SetDnsServers(v []*string) *ModifyVpcAttributesInput {
	s.DnsServers = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *ModifyVpcAttributesInput) SetVpcId(v string) *ModifyVpcAttributesInput {
	s.VpcId = &v
	return s
}

// SetVpcName sets the VpcName field's value.
func (s *ModifyVpcAttributesInput) SetVpcName(v string) *ModifyVpcAttributesInput {
	s.VpcName = &v
	return s
}

type ModifyVpcAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyVpcAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpcAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyVpcAttributesOutput) SetRequestId(v string) *ModifyVpcAttributesOutput {
	s.RequestId = &v
	return s
}
