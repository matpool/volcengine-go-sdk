// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package natgateway

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeNatGatewaysCommon = "DescribeNatGateways"

// DescribeNatGatewaysCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNatGatewaysCommon operation. The "output" return
// value will be populated with the DescribeNatGatewaysCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNatGatewaysCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNatGatewaysCommon Send returns without error.
//
// See DescribeNatGatewaysCommon for more information on using the DescribeNatGatewaysCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeNatGatewaysCommonRequest method.
//	req, resp := client.DescribeNatGatewaysCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *NATGATEWAY) DescribeNatGatewaysCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeNatGatewaysCommon,
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

// DescribeNatGatewaysCommon API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NATGATEWAY's
// API operation DescribeNatGatewaysCommon for usage and error information.
func (c *NATGATEWAY) DescribeNatGatewaysCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeNatGatewaysCommonRequest(input)
	return out, req.Send()
}

// DescribeNatGatewaysCommonWithContext is the same as DescribeNatGatewaysCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNatGatewaysCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) DescribeNatGatewaysCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeNatGatewaysCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeNatGateways = "DescribeNatGateways"

// DescribeNatGatewaysRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeNatGateways operation. The "output" return
// value will be populated with the DescribeNatGatewaysCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeNatGatewaysCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeNatGatewaysCommon Send returns without error.
//
// See DescribeNatGateways for more information on using the DescribeNatGateways
// API call, and error handling.
//
//	// Example sending a request using the DescribeNatGatewaysRequest method.
//	req, resp := client.DescribeNatGatewaysRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *NATGATEWAY) DescribeNatGatewaysRequest(input *DescribeNatGatewaysInput) (req *request.Request, output *DescribeNatGatewaysOutput) {
	op := &request.Operation{
		Name:       opDescribeNatGateways,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeNatGatewaysInput{}
	}

	output = &DescribeNatGatewaysOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeNatGateways API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NATGATEWAY's
// API operation DescribeNatGateways for usage and error information.
func (c *NATGATEWAY) DescribeNatGateways(input *DescribeNatGatewaysInput) (*DescribeNatGatewaysOutput, error) {
	req, out := c.DescribeNatGatewaysRequest(input)
	return out, req.Send()
}

// DescribeNatGatewaysWithContext is the same as DescribeNatGateways with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeNatGateways for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) DescribeNatGatewaysWithContext(ctx volcengine.Context, input *DescribeNatGatewaysInput, opts ...request.Option) (*DescribeNatGatewaysOutput, error) {
	req, out := c.DescribeNatGatewaysRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeNatGatewaysInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	NatGatewayIds []*string `type:"list"`

	NatGatewayName *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	ProjectName *string `type:"string"`

	Spec *string `type:"string" enum:"SpecForDescribeNatGatewaysInput"`

	SubnetId *string `type:"string"`

	TagFilters []*TagFilterForDescribeNatGatewaysInput `type:"list"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s DescribeNatGatewaysInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNatGatewaysInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeNatGatewaysInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeNatGatewaysInput"}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *DescribeNatGatewaysInput) SetDescription(v string) *DescribeNatGatewaysInput {
	s.Description = &v
	return s
}

// SetNatGatewayIds sets the NatGatewayIds field's value.
func (s *DescribeNatGatewaysInput) SetNatGatewayIds(v []*string) *DescribeNatGatewaysInput {
	s.NatGatewayIds = v
	return s
}

// SetNatGatewayName sets the NatGatewayName field's value.
func (s *DescribeNatGatewaysInput) SetNatGatewayName(v string) *DescribeNatGatewaysInput {
	s.NatGatewayName = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeNatGatewaysInput) SetPageNumber(v int64) *DescribeNatGatewaysInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeNatGatewaysInput) SetPageSize(v int64) *DescribeNatGatewaysInput {
	s.PageSize = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeNatGatewaysInput) SetProjectName(v string) *DescribeNatGatewaysInput {
	s.ProjectName = &v
	return s
}

// SetSpec sets the Spec field's value.
func (s *DescribeNatGatewaysInput) SetSpec(v string) *DescribeNatGatewaysInput {
	s.Spec = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeNatGatewaysInput) SetSubnetId(v string) *DescribeNatGatewaysInput {
	s.SubnetId = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeNatGatewaysInput) SetTagFilters(v []*TagFilterForDescribeNatGatewaysInput) *DescribeNatGatewaysInput {
	s.TagFilters = v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeNatGatewaysInput) SetVpcId(v string) *DescribeNatGatewaysInput {
	s.VpcId = &v
	return s
}

type DescribeNatGatewaysOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NatGateways []*NatGatewayForDescribeNatGatewaysOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeNatGatewaysOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeNatGatewaysOutput) GoString() string {
	return s.String()
}

// SetNatGateways sets the NatGateways field's value.
func (s *DescribeNatGatewaysOutput) SetNatGateways(v []*NatGatewayForDescribeNatGatewaysOutput) *DescribeNatGatewaysOutput {
	s.NatGateways = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeNatGatewaysOutput) SetPageNumber(v int64) *DescribeNatGatewaysOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeNatGatewaysOutput) SetPageSize(v int64) *DescribeNatGatewaysOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeNatGatewaysOutput) SetRequestId(v string) *DescribeNatGatewaysOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeNatGatewaysOutput) SetTotalCount(v int64) *DescribeNatGatewaysOutput {
	s.TotalCount = &v
	return s
}

type EipAddressForDescribeNatGatewaysOutput struct {
	_ struct{} `type:"structure"`

	AllocationId *string `type:"string"`

	EipAddress *string `type:"string"`

	UsingStatus *string `type:"string"`
}

// String returns the string representation
func (s EipAddressForDescribeNatGatewaysOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipAddressForDescribeNatGatewaysOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *EipAddressForDescribeNatGatewaysOutput) SetAllocationId(v string) *EipAddressForDescribeNatGatewaysOutput {
	s.AllocationId = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *EipAddressForDescribeNatGatewaysOutput) SetEipAddress(v string) *EipAddressForDescribeNatGatewaysOutput {
	s.EipAddress = &v
	return s
}

// SetUsingStatus sets the UsingStatus field's value.
func (s *EipAddressForDescribeNatGatewaysOutput) SetUsingStatus(v string) *EipAddressForDescribeNatGatewaysOutput {
	s.UsingStatus = &v
	return s
}

type NatGatewayForDescribeNatGatewaysOutput struct {
	_ struct{} `type:"structure"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	EipAddresses []*EipAddressForDescribeNatGatewaysOutput `type:"list"`

	ExpiredTime *string `type:"string"`

	LockReason *string `type:"string"`

	NatGatewayId *string `type:"string"`

	NatGatewayName *string `type:"string"`

	NetworkInterfaceId *string `type:"string"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	Spec *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	Tags []*TagForDescribeNatGatewaysOutput `type:"list"`

	UpdatedAt *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s NatGatewayForDescribeNatGatewaysOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NatGatewayForDescribeNatGatewaysOutput) GoString() string {
	return s.String()
}

// SetBillingType sets the BillingType field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetBillingType(v int64) *NatGatewayForDescribeNatGatewaysOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetBusinessStatus(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.BusinessStatus = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetCreationTime(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetDeletedTime(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetDescription(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.Description = &v
	return s
}

// SetEipAddresses sets the EipAddresses field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetEipAddresses(v []*EipAddressForDescribeNatGatewaysOutput) *NatGatewayForDescribeNatGatewaysOutput {
	s.EipAddresses = v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetExpiredTime(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.ExpiredTime = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetLockReason(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.LockReason = &v
	return s
}

// SetNatGatewayId sets the NatGatewayId field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetNatGatewayId(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.NatGatewayId = &v
	return s
}

// SetNatGatewayName sets the NatGatewayName field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetNatGatewayName(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.NatGatewayName = &v
	return s
}

// SetNetworkInterfaceId sets the NetworkInterfaceId field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetNetworkInterfaceId(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.NetworkInterfaceId = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetOverdueTime(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetProjectName(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.ProjectName = &v
	return s
}

// SetSpec sets the Spec field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetSpec(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.Spec = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetStatus(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetSubnetId(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.SubnetId = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetTags(v []*TagForDescribeNatGatewaysOutput) *NatGatewayForDescribeNatGatewaysOutput {
	s.Tags = v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetUpdatedAt(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.UpdatedAt = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetVpcId(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.VpcId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NatGatewayForDescribeNatGatewaysOutput) SetZoneId(v string) *NatGatewayForDescribeNatGatewaysOutput {
	s.ZoneId = &v
	return s
}

type TagFilterForDescribeNatGatewaysInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeNatGatewaysInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeNatGatewaysInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeNatGatewaysInput) SetKey(v string) *TagFilterForDescribeNatGatewaysInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeNatGatewaysInput) SetValues(v []*string) *TagFilterForDescribeNatGatewaysInput {
	s.Values = v
	return s
}

type TagForDescribeNatGatewaysOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeNatGatewaysOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeNatGatewaysOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeNatGatewaysOutput) SetKey(v string) *TagForDescribeNatGatewaysOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeNatGatewaysOutput) SetValue(v string) *TagForDescribeNatGatewaysOutput {
	s.Value = &v
	return s
}

const (
	// SpecForDescribeNatGatewaysInputSmall is a SpecForDescribeNatGatewaysInput enum value
	SpecForDescribeNatGatewaysInputSmall = "Small"

	// SpecForDescribeNatGatewaysInputMedium is a SpecForDescribeNatGatewaysInput enum value
	SpecForDescribeNatGatewaysInputMedium = "Medium"

	// SpecForDescribeNatGatewaysInputLarge is a SpecForDescribeNatGatewaysInput enum value
	SpecForDescribeNatGatewaysInputLarge = "Large"
)
