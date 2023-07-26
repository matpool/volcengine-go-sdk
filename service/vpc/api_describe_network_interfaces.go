// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeNetworkInterfacesCommon = "DescribeNetworkInterfaces"

// DescribeNetworkInterfacesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNetworkInterfacesCommon operation. The "output" return
// value will be populated with the DescribeNetworkInterfacesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNetworkInterfacesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNetworkInterfacesCommon Send returns without error.
//
// See DescribeNetworkInterfacesCommon for more information on using the DescribeNetworkInterfacesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeNetworkInterfacesCommonRequest method.
//	req, resp := client.DescribeNetworkInterfacesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeNetworkInterfacesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeNetworkInterfacesCommon,
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

// DescribeNetworkInterfacesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeNetworkInterfacesCommon for usage and error information.
func (c *VPC) DescribeNetworkInterfacesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeNetworkInterfacesCommonRequest(input)
	return out, req.Send()
}

// DescribeNetworkInterfacesCommonWithContext is the same as DescribeNetworkInterfacesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNetworkInterfacesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeNetworkInterfacesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeNetworkInterfacesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeNetworkInterfaces = "DescribeNetworkInterfaces"

// DescribeNetworkInterfacesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNetworkInterfaces operation. The "output" return
// value will be populated with the DescribeNetworkInterfacesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNetworkInterfacesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNetworkInterfacesCommon Send returns without error.
//
// See DescribeNetworkInterfaces for more information on using the DescribeNetworkInterfaces
// API call, and error handling.
//
//	// Example sending a request using the DescribeNetworkInterfacesRequest method.
//	req, resp := client.DescribeNetworkInterfacesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DescribeNetworkInterfacesRequest(input *DescribeNetworkInterfacesInput) (req *request.Request, output *DescribeNetworkInterfacesOutput) {
	op := &request.Operation{
		Name:       opDescribeNetworkInterfaces,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNetworkInterfacesInput{}
	}

	output = &DescribeNetworkInterfacesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeNetworkInterfaces API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeNetworkInterfaces for usage and error information.
func (c *VPC) DescribeNetworkInterfaces(input *DescribeNetworkInterfacesInput) (*DescribeNetworkInterfacesOutput, error) {
	req, out := c.DescribeNetworkInterfacesRequest(input)
	return out, req.Send()
}

// DescribeNetworkInterfacesWithContext is the same as DescribeNetworkInterfaces with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNetworkInterfaces for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeNetworkInterfacesWithContext(ctx volcengine.Context, input *DescribeNetworkInterfacesInput, opts ...request.Option) (*DescribeNetworkInterfacesOutput, error) {
	req, out := c.DescribeNetworkInterfacesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociatedElasticIpForDescribeNetworkInterfacesOutput struct {
	_ struct{} `type:"structure"`

	AllocationId *string `type:"string"`

	EipAddress *string `type:"string"`
}

// String returns the string representation
func (s AssociatedElasticIpForDescribeNetworkInterfacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociatedElasticIpForDescribeNetworkInterfacesOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *AssociatedElasticIpForDescribeNetworkInterfacesOutput) SetAllocationId(v string) *AssociatedElasticIpForDescribeNetworkInterfacesOutput {
	s.AllocationId = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *AssociatedElasticIpForDescribeNetworkInterfacesOutput) SetEipAddress(v string) *AssociatedElasticIpForDescribeNetworkInterfacesOutput {
	s.EipAddress = &v
	return s
}

type DescribeNetworkInterfacesInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	NetworkInterfaceIds []*string `type:"list"`

	NetworkInterfaceName *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	PrimaryIpAddresses []*string `type:"list"`

	PrivateIpAddresses []*string `type:"list"`

	ProjectName *string `type:"string"`

	SecurityGroupId *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	TagFilters []*TagFilterForDescribeNetworkInterfacesInput `type:"list"`

	Type *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s DescribeNetworkInterfacesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNetworkInterfacesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNetworkInterfacesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeNetworkInterfacesInput"}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeNetworkInterfacesInput) SetInstanceId(v string) *DescribeNetworkInterfacesInput {
	s.InstanceId = &v
	return s
}

// SetNetworkInterfaceIds sets the NetworkInterfaceIds field's value.
func (s *DescribeNetworkInterfacesInput) SetNetworkInterfaceIds(v []*string) *DescribeNetworkInterfacesInput {
	s.NetworkInterfaceIds = v
	return s
}

// SetNetworkInterfaceName sets the NetworkInterfaceName field's value.
func (s *DescribeNetworkInterfacesInput) SetNetworkInterfaceName(v string) *DescribeNetworkInterfacesInput {
	s.NetworkInterfaceName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeNetworkInterfacesInput) SetPageNumber(v int64) *DescribeNetworkInterfacesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeNetworkInterfacesInput) SetPageSize(v int64) *DescribeNetworkInterfacesInput {
	s.PageSize = &v
	return s
}

// SetPrimaryIpAddresses sets the PrimaryIpAddresses field's value.
func (s *DescribeNetworkInterfacesInput) SetPrimaryIpAddresses(v []*string) *DescribeNetworkInterfacesInput {
	s.PrimaryIpAddresses = v
	return s
}

// SetPrivateIpAddresses sets the PrivateIpAddresses field's value.
func (s *DescribeNetworkInterfacesInput) SetPrivateIpAddresses(v []*string) *DescribeNetworkInterfacesInput {
	s.PrivateIpAddresses = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeNetworkInterfacesInput) SetProjectName(v string) *DescribeNetworkInterfacesInput {
	s.ProjectName = &v
	return s
}

// SetSecurityGroupId sets the SecurityGroupId field's value.
func (s *DescribeNetworkInterfacesInput) SetSecurityGroupId(v string) *DescribeNetworkInterfacesInput {
	s.SecurityGroupId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeNetworkInterfacesInput) SetStatus(v string) *DescribeNetworkInterfacesInput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeNetworkInterfacesInput) SetSubnetId(v string) *DescribeNetworkInterfacesInput {
	s.SubnetId = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeNetworkInterfacesInput) SetTagFilters(v []*TagFilterForDescribeNetworkInterfacesInput) *DescribeNetworkInterfacesInput {
	s.TagFilters = v
	return s
}

// SetType sets the Type field's value.
func (s *DescribeNetworkInterfacesInput) SetType(v string) *DescribeNetworkInterfacesInput {
	s.Type = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeNetworkInterfacesInput) SetVpcId(v string) *DescribeNetworkInterfacesInput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *DescribeNetworkInterfacesInput) SetZoneId(v string) *DescribeNetworkInterfacesInput {
	s.ZoneId = &v
	return s
}

type DescribeNetworkInterfacesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NetworkInterfaceSets []*NetworkInterfaceSetForDescribeNetworkInterfacesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeNetworkInterfacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNetworkInterfacesOutput) GoString() string {
	return s.String()
}

// SetNetworkInterfaceSets sets the NetworkInterfaceSets field's value.
func (s *DescribeNetworkInterfacesOutput) SetNetworkInterfaceSets(v []*NetworkInterfaceSetForDescribeNetworkInterfacesOutput) *DescribeNetworkInterfacesOutput {
	s.NetworkInterfaceSets = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeNetworkInterfacesOutput) SetPageNumber(v int64) *DescribeNetworkInterfacesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeNetworkInterfacesOutput) SetPageSize(v int64) *DescribeNetworkInterfacesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeNetworkInterfacesOutput) SetRequestId(v string) *DescribeNetworkInterfacesOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeNetworkInterfacesOutput) SetTotalCount(v int64) *DescribeNetworkInterfacesOutput {
	s.TotalCount = &v
	return s
}

type NetworkInterfaceSetForDescribeNetworkInterfacesOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	AssociatedElasticIp *AssociatedElasticIpForDescribeNetworkInterfacesOutput `type:"structure"`

	CreatedAt *string `type:"string"`

	Description *string `type:"string"`

	DeviceId *string `type:"string"`

	IPv6Sets []*string `type:"list"`

	MacAddress *string `type:"string"`

	NetworkInterfaceId *string `type:"string"`

	NetworkInterfaceName *string `type:"string"`

	PortSecurityEnabled *bool `type:"boolean"`

	PrimaryIpAddress *string `type:"string"`

	PrivateIpSets *PrivateIpSetsForDescribeNetworkInterfacesOutput `type:"structure"`

	ProjectName *string `type:"string"`

	SecurityGroupIds []*string `type:"list"`

	ServiceManaged *bool `type:"boolean"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	Tags []*TagForDescribeNetworkInterfacesOutput `type:"list"`

	Type *string `type:"string"`

	UpdatedAt *string `type:"string"`

	VpcId *string `type:"string"`

	VpcName *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s NetworkInterfaceSetForDescribeNetworkInterfacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NetworkInterfaceSetForDescribeNetworkInterfacesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetAccountId(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.AccountId = &v
	return s
}

// SetAssociatedElasticIp sets the AssociatedElasticIp field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetAssociatedElasticIp(v *AssociatedElasticIpForDescribeNetworkInterfacesOutput) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.AssociatedElasticIp = v
	return s
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetCreatedAt(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.CreatedAt = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetDescription(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.Description = &v
	return s
}

// SetDeviceId sets the DeviceId field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetDeviceId(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.DeviceId = &v
	return s
}

// SetIPv6Sets sets the IPv6Sets field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetIPv6Sets(v []*string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.IPv6Sets = v
	return s
}

// SetMacAddress sets the MacAddress field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetMacAddress(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.MacAddress = &v
	return s
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetNetworkInterfaceId(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.NetworkInterfaceId = &v
	return s
}

// SetNetworkInterfaceName sets the NetworkInterfaceName field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetNetworkInterfaceName(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.NetworkInterfaceName = &v
	return s
}

// SetPortSecurityEnabled sets the PortSecurityEnabled field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetPortSecurityEnabled(v bool) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.PortSecurityEnabled = &v
	return s
}

// SetPrimaryIpAddress sets the PrimaryIpAddress field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetPrimaryIpAddress(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.PrimaryIpAddress = &v
	return s
}

// SetPrivateIpSets sets the PrivateIpSets field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetPrivateIpSets(v *PrivateIpSetsForDescribeNetworkInterfacesOutput) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.PrivateIpSets = v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetProjectName(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.ProjectName = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetSecurityGroupIds(v []*string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.SecurityGroupIds = v
	return s
}

// SetServiceManaged sets the ServiceManaged field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetServiceManaged(v bool) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.ServiceManaged = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetStatus(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetSubnetId(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetTags(v []*TagForDescribeNetworkInterfacesOutput) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.Tags = v
	return s
}

// SetType sets the Type field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetType(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.Type = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetUpdatedAt(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.UpdatedAt = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetVpcId(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.VpcId = &v
	return s
}

// SetVpcName sets the VpcName field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetVpcName(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.VpcName = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NetworkInterfaceSetForDescribeNetworkInterfacesOutput) SetZoneId(v string) *NetworkInterfaceSetForDescribeNetworkInterfacesOutput {
	s.ZoneId = &v
	return s
}

type PrivateIpSetForDescribeNetworkInterfacesOutput struct {
	_ struct{} `type:"structure"`

	AssociatedElasticIp *AssociatedElasticIpForDescribeNetworkInterfacesOutput `type:"structure"`

	Primary *bool `type:"boolean"`

	PrivateIpAddress *string `type:"string"`
}

// String returns the string representation
func (s PrivateIpSetForDescribeNetworkInterfacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrivateIpSetForDescribeNetworkInterfacesOutput) GoString() string {
	return s.String()
}

// SetAssociatedElasticIp sets the AssociatedElasticIp field's value.
func (s *PrivateIpSetForDescribeNetworkInterfacesOutput) SetAssociatedElasticIp(v *AssociatedElasticIpForDescribeNetworkInterfacesOutput) *PrivateIpSetForDescribeNetworkInterfacesOutput {
	s.AssociatedElasticIp = v
	return s
}

// SetPrimary sets the Primary field's value.
func (s *PrivateIpSetForDescribeNetworkInterfacesOutput) SetPrimary(v bool) *PrivateIpSetForDescribeNetworkInterfacesOutput {
	s.Primary = &v
	return s
}

// SetPrivateIpAddress sets the PrivateIpAddress field's value.
func (s *PrivateIpSetForDescribeNetworkInterfacesOutput) SetPrivateIpAddress(v string) *PrivateIpSetForDescribeNetworkInterfacesOutput {
	s.PrivateIpAddress = &v
	return s
}

type PrivateIpSetsForDescribeNetworkInterfacesOutput struct {
	_ struct{} `type:"structure"`

	PrivateIpSet []*PrivateIpSetForDescribeNetworkInterfacesOutput `type:"list"`
}

// String returns the string representation
func (s PrivateIpSetsForDescribeNetworkInterfacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PrivateIpSetsForDescribeNetworkInterfacesOutput) GoString() string {
	return s.String()
}

// SetPrivateIpSet sets the PrivateIpSet field's value.
func (s *PrivateIpSetsForDescribeNetworkInterfacesOutput) SetPrivateIpSet(v []*PrivateIpSetForDescribeNetworkInterfacesOutput) *PrivateIpSetsForDescribeNetworkInterfacesOutput {
	s.PrivateIpSet = v
	return s
}

type TagFilterForDescribeNetworkInterfacesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeNetworkInterfacesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeNetworkInterfacesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeNetworkInterfacesInput) SetKey(v string) *TagFilterForDescribeNetworkInterfacesInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeNetworkInterfacesInput) SetValues(v []*string) *TagFilterForDescribeNetworkInterfacesInput {
	s.Values = v
	return s
}

type TagForDescribeNetworkInterfacesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeNetworkInterfacesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeNetworkInterfacesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeNetworkInterfacesOutput) SetKey(v string) *TagForDescribeNetworkInterfacesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeNetworkInterfacesOutput) SetValue(v string) *TagForDescribeNetworkInterfacesOutput {
	s.Value = &v
	return s
}
