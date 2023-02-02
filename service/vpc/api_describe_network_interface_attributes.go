// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeNetworkInterfaceAttributesCommon = "DescribeNetworkInterfaceAttributes"

// DescribeNetworkInterfaceAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNetworkInterfaceAttributesCommon operation. The "output" return
// value will be populated with the DescribeNetworkInterfaceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNetworkInterfaceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNetworkInterfaceAttributesCommon Send returns without error.
//
// See DescribeNetworkInterfaceAttributesCommon for more information on using the DescribeNetworkInterfaceAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeNetworkInterfaceAttributesCommonRequest method.
//    req, resp := client.DescribeNetworkInterfaceAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeNetworkInterfaceAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeNetworkInterfaceAttributesCommon,
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

// DescribeNetworkInterfaceAttributesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeNetworkInterfaceAttributesCommon for usage and error information.
func (c *VPC) DescribeNetworkInterfaceAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeNetworkInterfaceAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeNetworkInterfaceAttributesCommonWithContext is the same as DescribeNetworkInterfaceAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNetworkInterfaceAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeNetworkInterfaceAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeNetworkInterfaceAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeNetworkInterfaceAttributes = "DescribeNetworkInterfaceAttributes"

// DescribeNetworkInterfaceAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNetworkInterfaceAttributes operation. The "output" return
// value will be populated with the DescribeNetworkInterfaceAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNetworkInterfaceAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNetworkInterfaceAttributesCommon Send returns without error.
//
// See DescribeNetworkInterfaceAttributes for more information on using the DescribeNetworkInterfaceAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeNetworkInterfaceAttributesRequest method.
//    req, resp := client.DescribeNetworkInterfaceAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeNetworkInterfaceAttributesRequest(input *DescribeNetworkInterfaceAttributesInput) (req *request.Request, output *DescribeNetworkInterfaceAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeNetworkInterfaceAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNetworkInterfaceAttributesInput{}
	}

	output = &DescribeNetworkInterfaceAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeNetworkInterfaceAttributes API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeNetworkInterfaceAttributes for usage and error information.
func (c *VPC) DescribeNetworkInterfaceAttributes(input *DescribeNetworkInterfaceAttributesInput) (*DescribeNetworkInterfaceAttributesOutput, error) {
	req, out := c.DescribeNetworkInterfaceAttributesRequest(input)
	return out, req.Send()
}

// DescribeNetworkInterfaceAttributesWithContext is the same as DescribeNetworkInterfaceAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNetworkInterfaceAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeNetworkInterfaceAttributesWithContext(ctx volcengine.Context, input *DescribeNetworkInterfaceAttributesInput, opts ...request.Option) (*DescribeNetworkInterfaceAttributesOutput, error) {
	req, out := c.DescribeNetworkInterfaceAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput struct {
	_ struct{} `type:"structure"`

	AllocationId *string `type:"string"`

	EipAddress *string `type:"string"`
}

// String returns the string representation
func (s AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput) SetAllocationId(v string) *AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput {
	s.AllocationId = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput) SetEipAddress(v string) *AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput {
	s.EipAddress = &v
	return s
}

type DescribeNetworkInterfaceAttributesInput struct {
	_ struct{} `type:"structure"`

	// NetworkInterfaceId is a required field
	NetworkInterfaceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeNetworkInterfaceAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNetworkInterfaceAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNetworkInterfaceAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeNetworkInterfaceAttributesInput"}
	if s.NetworkInterfaceId == nil {
		invalidParams.Add(request.NewErrParamRequired("NetworkInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *DescribeNetworkInterfaceAttributesInput) SetNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributesInput {
	s.NetworkInterfaceId = &v
	return s
}

type DescribeNetworkInterfaceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string"`

	AssociatedElasticIp *AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput `type:"structure"`

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	DeviceId *string `type:"string"`

	MacAddress *string `type:"string"`

	NetworkInterfaceId *string `type:"string"`

	NetworkInterfaceName *string `type:"string"`

	PortSecurityEnabled *bool `type:"boolean"`

	PrimaryIpAddress *string `type:"string"`

	PrivateIpSets *PrivateIpSetsForDescribeNetworkInterfaceAttributesOutput `type:"structure"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	SecurityGroupIds []*string `type:"list"`

	ServiceManaged *bool `type:"boolean"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	Tags []*TagForDescribeNetworkInterfaceAttributesOutput `type:"list"`

	Type *string `type:"string"`

	UpdatedAt *string `type:"string"`

	VpcId *string `type:"string"`

	VpcName *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeNetworkInterfaceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNetworkInterfaceAttributesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetAccountId(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.AccountId = &v
	return s
}

// SetAssociatedElasticIp sets the AssociatedElasticIp field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetAssociatedElasticIp(v *AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput) *DescribeNetworkInterfaceAttributesOutput {
	s.AssociatedElasticIp = v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetCreatedAt(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetDescription(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.Description = &v
	return s
}

// SetDeviceId sets the DeviceId field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetDeviceId(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.DeviceId = &v
	return s
}

// SetMacAddress sets the MacAddress field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetMacAddress(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.MacAddress = &v
	return s
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.NetworkInterfaceId = &v
	return s
}

// SetNetworkInterfaceName sets the NetworkInterfaceName field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetNetworkInterfaceName(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.NetworkInterfaceName = &v
	return s
}

// SetPortSecurityEnabled sets the PortSecurityEnabled field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetPortSecurityEnabled(v bool) *DescribeNetworkInterfaceAttributesOutput {
	s.PortSecurityEnabled = &v
	return s
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetPrimaryIpAddress(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.PrimaryIpAddress = &v
	return s
}

// SetPrivateIpSets sets the PrivateIpSets field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetPrivateIpSets(v *PrivateIpSetsForDescribeNetworkInterfaceAttributesOutput) *DescribeNetworkInterfaceAttributesOutput {
	s.PrivateIpSets = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetProjectName(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetRequestId(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.RequestId = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetSecurityGroupIds(v []*string) *DescribeNetworkInterfaceAttributesOutput {
	s.SecurityGroupIds = v
	return s
}

// SetServiceManaged sets the ServiceManaged field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetServiceManaged(v bool) *DescribeNetworkInterfaceAttributesOutput {
	s.ServiceManaged = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetStatus(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetSubnetId(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetTags(v []*TagForDescribeNetworkInterfaceAttributesOutput) *DescribeNetworkInterfaceAttributesOutput {
	s.Tags = v
	return s
}

// SetType sets the Type field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetType(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.Type = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetUpdatedAt(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.UpdatedAt = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetVpcId(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.VpcId = &v
	return s
}

// SetVpcName sets the VpcName field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetVpcName(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.VpcName = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeNetworkInterfaceAttributesOutput) SetZoneId(v string) *DescribeNetworkInterfaceAttributesOutput {
	s.ZoneId = &v
	return s
}

type PrivateIpSetForDescribeNetworkInterfaceAttributesOutput struct {
	_ struct{} `type:"structure"`

	AssociatedElasticIp *AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput `type:"structure"`

	Primary *bool `type:"boolean"`

	PrivateIpAddress *string `type:"string"`
}

// String returns the string representation
func (s PrivateIpSetForDescribeNetworkInterfaceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrivateIpSetForDescribeNetworkInterfaceAttributesOutput) GoString() string {
	return s.String()
}

// SetAssociatedElasticIp sets the AssociatedElasticIp field's value.
func (s *PrivateIpSetForDescribeNetworkInterfaceAttributesOutput) SetAssociatedElasticIp(v *AssociatedElasticIpForDescribeNetworkInterfaceAttributesOutput) *PrivateIpSetForDescribeNetworkInterfaceAttributesOutput {
	s.AssociatedElasticIp = v
	return s
}

// SetPrimary sets the Primary field's value.
func (s *PrivateIpSetForDescribeNetworkInterfaceAttributesOutput) SetPrimary(v bool) *PrivateIpSetForDescribeNetworkInterfaceAttributesOutput {
	s.Primary = &v
	return s
}

// SetPrivateIpAddress sets the PrivateIpAddress field's value.
func (s *PrivateIpSetForDescribeNetworkInterfaceAttributesOutput) SetPrivateIpAddress(v string) *PrivateIpSetForDescribeNetworkInterfaceAttributesOutput {
	s.PrivateIpAddress = &v
	return s
}

type PrivateIpSetsForDescribeNetworkInterfaceAttributesOutput struct {
	_ struct{} `type:"structure"`

	PrivateIpSet []*PrivateIpSetForDescribeNetworkInterfaceAttributesOutput `type:"list"`
}

// String returns the string representation
func (s PrivateIpSetsForDescribeNetworkInterfaceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrivateIpSetsForDescribeNetworkInterfaceAttributesOutput) GoString() string {
	return s.String()
}

// SetPrivateIpSet sets the PrivateIpSet field's value.
func (s *PrivateIpSetsForDescribeNetworkInterfaceAttributesOutput) SetPrivateIpSet(v []*PrivateIpSetForDescribeNetworkInterfaceAttributesOutput) *PrivateIpSetsForDescribeNetworkInterfaceAttributesOutput {
	s.PrivateIpSet = v
	return s
}

type TagForDescribeNetworkInterfaceAttributesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeNetworkInterfaceAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeNetworkInterfaceAttributesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeNetworkInterfaceAttributesOutput) SetKey(v string) *TagForDescribeNetworkInterfaceAttributesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeNetworkInterfaceAttributesOutput) SetValue(v string) *TagForDescribeNetworkInterfaceAttributesOutput {
	s.Value = &v
	return s
}
