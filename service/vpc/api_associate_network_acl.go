// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAssociateNetworkAclCommon = "AssociateNetworkAcl"

// AssociateNetworkAclCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateNetworkAclCommon operation. The "output" return
// value will be populated with the AssociateNetworkAclCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateNetworkAclCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateNetworkAclCommon Send returns without error.
//
// See AssociateNetworkAclCommon for more information on using the AssociateNetworkAclCommon
// API call, and error handling.
//
//	// Example sending a request using the AssociateNetworkAclCommonRequest method.
//	req, resp := client.AssociateNetworkAclCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AssociateNetworkAclCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAssociateNetworkAclCommon,
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

// AssociateNetworkAclCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AssociateNetworkAclCommon for usage and error information.
func (c *VPC) AssociateNetworkAclCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AssociateNetworkAclCommonRequest(input)
	return out, req.Send()
}

// AssociateNetworkAclCommonWithContext is the same as AssociateNetworkAclCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateNetworkAclCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssociateNetworkAclCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AssociateNetworkAclCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAssociateNetworkAcl = "AssociateNetworkAcl"

// AssociateNetworkAclRequest generates a "volcengine/request.Request" representing the
// client's request for the AssociateNetworkAcl operation. The "output" return
// value will be populated with the AssociateNetworkAclCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AssociateNetworkAclCommon Request to send the API call to the service.
// the "output" return value is not valid until after AssociateNetworkAclCommon Send returns without error.
//
// See AssociateNetworkAcl for more information on using the AssociateNetworkAcl
// API call, and error handling.
//
//	// Example sending a request using the AssociateNetworkAclRequest method.
//	req, resp := client.AssociateNetworkAclRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AssociateNetworkAclRequest(input *AssociateNetworkAclInput) (req *request.Request, output *AssociateNetworkAclOutput) {
	op := &request.Operation{
		Name:       opAssociateNetworkAcl,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateNetworkAclInput{}
	}

	output = &AssociateNetworkAclOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AssociateNetworkAcl API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AssociateNetworkAcl for usage and error information.
func (c *VPC) AssociateNetworkAcl(input *AssociateNetworkAclInput) (*AssociateNetworkAclOutput, error) {
	req, out := c.AssociateNetworkAclRequest(input)
	return out, req.Send()
}

// AssociateNetworkAclWithContext is the same as AssociateNetworkAcl with the addition of
// the ability to pass a context and additional request options.
//
// See AssociateNetworkAcl for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AssociateNetworkAclWithContext(ctx volcengine.Context, input *AssociateNetworkAclInput, opts ...request.Option) (*AssociateNetworkAclOutput, error) {
	req, out := c.AssociateNetworkAclRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateNetworkAclInput struct {
	_ struct{} `type:"structure"`

	// NetworkAclId is a required field
	NetworkAclId *string `type:"string" required:"true"`

	Resource []*ResourceForAssociateNetworkAclInput `type:"list"`
}

// String returns the string representation
func (s AssociateNetworkAclInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateNetworkAclInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateNetworkAclInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AssociateNetworkAclInput"}
	if s.NetworkAclId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkAclId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNetworkAclId sets the NetworkAclId field's value.
func (s *AssociateNetworkAclInput) SetNetworkAclId(v string) *AssociateNetworkAclInput {
	s.NetworkAclId = &v
	return s
}

// SetResource sets the Resource field's value.
func (s *AssociateNetworkAclInput) SetResource(v []*ResourceForAssociateNetworkAclInput) *AssociateNetworkAclInput {
	s.Resource = v
	return s
}

type AssociateNetworkAclOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AssociateNetworkAclOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateNetworkAclOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *AssociateNetworkAclOutput) SetRequestId(v string) *AssociateNetworkAclOutput {
	s.RequestId = &v
	return s
}

type ResourceForAssociateNetworkAclInput struct {
	_ struct{} `type:"structure"`

	ResourceId *string `type:"string"`
}

// String returns the string representation
func (s ResourceForAssociateNetworkAclInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceForAssociateNetworkAclInput) GoString() string {
	return s.String()
}

// SetResourceId sets the ResourceId field's value.
func (s *ResourceForAssociateNetworkAclInput) SetResourceId(v string) *ResourceForAssociateNetworkAclInput {
	s.ResourceId = &v
	return s
}
