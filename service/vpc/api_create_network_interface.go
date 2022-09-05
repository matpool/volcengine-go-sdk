// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateNetworkInterfaceCommon = "CreateNetworkInterface"

// CreateNetworkInterfaceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNetworkInterfaceCommon operation. The "output" return
// value will be populated with the CreateNetworkInterfaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNetworkInterfaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNetworkInterfaceCommon Send returns without error.
//
// See CreateNetworkInterfaceCommon for more information on using the CreateNetworkInterfaceCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateNetworkInterfaceCommonRequest method.
//    req, resp := client.CreateNetworkInterfaceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateNetworkInterfaceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateNetworkInterfaceCommon,
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

// CreateNetworkInterfaceCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateNetworkInterfaceCommon for usage and error information.
func (c *VPC) CreateNetworkInterfaceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateNetworkInterfaceCommonRequest(input)
	return out, req.Send()
}

// CreateNetworkInterfaceCommonWithContext is the same as CreateNetworkInterfaceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNetworkInterfaceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateNetworkInterfaceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateNetworkInterfaceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateNetworkInterface = "CreateNetworkInterface"

// CreateNetworkInterfaceRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNetworkInterface operation. The "output" return
// value will be populated with the CreateNetworkInterfaceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNetworkInterfaceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNetworkInterfaceCommon Send returns without error.
//
// See CreateNetworkInterface for more information on using the CreateNetworkInterface
// API call, and error handling.
//
//    // Example sending a request using the CreateNetworkInterfaceRequest method.
//    req, resp := client.CreateNetworkInterfaceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) CreateNetworkInterfaceRequest(input *CreateNetworkInterfaceInput) (req *request.Request, output *CreateNetworkInterfaceOutput) {
	op := &request.Operation{
		Name:       opCreateNetworkInterface,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNetworkInterfaceInput{}
	}

	output = &CreateNetworkInterfaceOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateNetworkInterface API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation CreateNetworkInterface for usage and error information.
func (c *VPC) CreateNetworkInterface(input *CreateNetworkInterfaceInput) (*CreateNetworkInterfaceOutput, error) {
	req, out := c.CreateNetworkInterfaceRequest(input)
	return out, req.Send()
}

// CreateNetworkInterfaceWithContext is the same as CreateNetworkInterface with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNetworkInterface for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) CreateNetworkInterfaceWithContext(ctx volcengine.Context, input *CreateNetworkInterfaceInput, opts ...request.Option) (*CreateNetworkInterfaceOutput, error) {
	req, out := c.CreateNetworkInterfaceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateNetworkInterfaceInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `min:"1" max:"255" type:"string"`

	NetworkInterfaceName *string `min:"1" max:"128" type:"string"`

	PortSecurityEnabled *bool `type:"boolean"`

	PrimaryIpAddress *string `type:"string"`

	PrivateIpAddress []*string `type:"list"`

	ProjectName *string `type:"string"`

	SecondaryPrivateIpAddressCount *int64 `type:"integer"`

	// SecurityGroupIds is a required field
	SecurityGroupIds []*string `type:"list" required:"true"`

	// SubnetId is a required field
	SubnetId *string `type:"string" required:"true"`

	Tags []*TagForCreateNetworkInterfaceInput `type:"list"`
}

// String returns the string representation
func (s CreateNetworkInterfaceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNetworkInterfaceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateNetworkInterfaceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateNetworkInterfaceInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.NetworkInterfaceName != nil && len(*s.NetworkInterfaceName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("NetworkInterfaceName", 1))
	}
	if s.NetworkInterfaceName != nil && len(*s.NetworkInterfaceName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("NetworkInterfaceName", 128, *s.NetworkInterfaceName))
	}
	if s.SecurityGroupIds == nil {
		invalidParams.Add(request.NewErrParamRequired("SecurityGroupIds"))
	}
	if s.SubnetId == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateNetworkInterfaceInput) SetClientToken(v string) *CreateNetworkInterfaceInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateNetworkInterfaceInput) SetDescription(v string) *CreateNetworkInterfaceInput {
	s.Description = &v
	return s
}

// SetNetworkInterfaceName sets the NetworkInterfaceName field's value.
func (s *CreateNetworkInterfaceInput) SetNetworkInterfaceName(v string) *CreateNetworkInterfaceInput {
	s.NetworkInterfaceName = &v
	return s
}

// SetPortSecurityEnabled sets the PortSecurityEnabled field's value.
func (s *CreateNetworkInterfaceInput) SetPortSecurityEnabled(v bool) *CreateNetworkInterfaceInput {
	s.PortSecurityEnabled = &v
	return s
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *CreateNetworkInterfaceInput) SetPrimaryIpAddress(v string) *CreateNetworkInterfaceInput {
	s.PrimaryIpAddress = &v
	return s
}

// SetPrivateIpAddress sets the PrivateIpAddress field's value.
func (s *CreateNetworkInterfaceInput) SetPrivateIpAddress(v []*string) *CreateNetworkInterfaceInput {
	s.PrivateIpAddress = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateNetworkInterfaceInput) SetProjectName(v string) *CreateNetworkInterfaceInput {
	s.ProjectName = &v
	return s
}

// SetSecondaryPrivateIpAddressCount sets the SecondaryPrivateIpAddressCount field's value.
func (s *CreateNetworkInterfaceInput) SetSecondaryPrivateIpAddressCount(v int64) *CreateNetworkInterfaceInput {
	s.SecondaryPrivateIpAddressCount = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *CreateNetworkInterfaceInput) SetSecurityGroupIds(v []*string) *CreateNetworkInterfaceInput {
	s.SecurityGroupIds = v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateNetworkInterfaceInput) SetSubnetId(v string) *CreateNetworkInterfaceInput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateNetworkInterfaceInput) SetTags(v []*TagForCreateNetworkInterfaceInput) *CreateNetworkInterfaceInput {
	s.Tags = v
	return s
}

type CreateNetworkInterfaceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NetworkInterfaceId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateNetworkInterfaceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNetworkInterfaceOutput) GoString() string {
	return s.String()
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *CreateNetworkInterfaceOutput) SetNetworkInterfaceId(v string) *CreateNetworkInterfaceOutput {
	s.NetworkInterfaceId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateNetworkInterfaceOutput) SetRequestId(v string) *CreateNetworkInterfaceOutput {
	s.RequestId = &v
	return s
}

type TagForCreateNetworkInterfaceInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateNetworkInterfaceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateNetworkInterfaceInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateNetworkInterfaceInput) SetKey(v string) *TagForCreateNetworkInterfaceInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateNetworkInterfaceInput) SetValue(v string) *TagForCreateNetworkInterfaceInput {
	s.Value = &v
	return s
}
