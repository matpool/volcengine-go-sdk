// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateVpcCommon = "CreateVpc"

// CreateVpcCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpcCommon operation. The "output" return
// value will be populated with the CreateVpcCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpcCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpcCommon Send returns without error.
//
// See CreateVpcCommon for more information on using the CreateVpcCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateVpcCommonRequest method.
//	req, resp := client.CreateVpcCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) CreateVpcCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateVpcCommon,
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

// CreateVpcCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateVpcCommon for usage and error information.
func (c *VPC) CreateVpcCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateVpcCommonRequest(input)
	return out, req.Send()
}

// CreateVpcCommonWithContext is the same as CreateVpcCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpcCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateVpcCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateVpcCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateVpc = "CreateVpc"

// CreateVpcRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpc operation. The "output" return
// value will be populated with the CreateVpcCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpcCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpcCommon Send returns without error.
//
// See CreateVpc for more information on using the CreateVpc
// API call, and error handling.
//
//	// Example sending a request using the CreateVpcRequest method.
//	req, resp := client.CreateVpcRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) CreateVpcRequest(input *CreateVpcInput) (req *request.Request, output *CreateVpcOutput) {
	op := &request.Operation{
		Name:       opCreateVpc,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpcInput{}
	}

	output = &CreateVpcOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateVpc API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateVpc for usage and error information.
func (c *VPC) CreateVpc(input *CreateVpcInput) (*CreateVpcOutput, error) {
	req, out := c.CreateVpcRequest(input)
	return out, req.Send()
}

// CreateVpcWithContext is the same as CreateVpc with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpc for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateVpcWithContext(ctx volcengine.Context, input *CreateVpcInput, opts ...request.Option) (*CreateVpcOutput, error) {
	req, out := c.CreateVpcRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateVpcInput struct {
	_ struct{} `type:"structure"`

	// CidrBlock is a required field
	CidrBlock *string `type:"string" required:"true"`

	ClientToken *string `type:"string"`

	Description *string `min:"1" max:"255" type:"string"`

	DnsServers []*string `type:"list"`

	ProjectName *string `type:"string"`

	Tags []*TagForCreateVpcInput `type:"list"`

	VpcName *string `min:"1" max:"128" type:"string"`
}

// String returns the string representation
func (s CreateVpcInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpcInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpcInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateVpcInput"}
	if s.CidrBlock == nil {
		invalidParams.Add(request.NewErrParamRequired("CidrBlock"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.VpcName != nil && len(*s.VpcName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("VpcName", 1))
	}
	if s.VpcName != nil && len(*s.VpcName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("VpcName", 128, *s.VpcName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCidrBlock sets the CidrBlock field's value.
func (s *CreateVpcInput) SetCidrBlock(v string) *CreateVpcInput {
	s.CidrBlock = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateVpcInput) SetClientToken(v string) *CreateVpcInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateVpcInput) SetDescription(v string) *CreateVpcInput {
	s.Description = &v
	return s
}

// SetDnsServers sets the DnsServers field's value.
func (s *CreateVpcInput) SetDnsServers(v []*string) *CreateVpcInput {
	s.DnsServers = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateVpcInput) SetProjectName(v string) *CreateVpcInput {
	s.ProjectName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateVpcInput) SetTags(v []*TagForCreateVpcInput) *CreateVpcInput {
	s.Tags = v
	return s
}

// SetVpcName sets the VpcName field's value.
func (s *CreateVpcInput) SetVpcName(v string) *CreateVpcInput {
	s.VpcName = &v
	return s
}

type CreateVpcOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	RouteTableId *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s CreateVpcOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpcOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateVpcOutput) SetRequestId(v string) *CreateVpcOutput {
	s.RequestId = &v
	return s
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *CreateVpcOutput) SetRouteTableId(v string) *CreateVpcOutput {
	s.RouteTableId = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateVpcOutput) SetVpcId(v string) *CreateVpcOutput {
	s.VpcId = &v
	return s
}

type TagForCreateVpcInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateVpcInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateVpcInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateVpcInput) SetKey(v string) *TagForCreateVpcInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateVpcInput) SetValue(v string) *TagForCreateVpcInput {
	s.Value = &v
	return s
}
