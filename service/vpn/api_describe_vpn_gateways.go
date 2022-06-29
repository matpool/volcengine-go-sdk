// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeVpnGatewaysCommon = "DescribeVpnGateways"

// DescribeVpnGatewaysCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpnGatewaysCommon operation. The "output" return
// value will be populated with the DescribeVpnGatewaysCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpnGatewaysCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpnGatewaysCommon Send returns without error.
//
// See DescribeVpnGatewaysCommon for more information on using the DescribeVpnGatewaysCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpnGatewaysCommonRequest method.
//    req, resp := client.DescribeVpnGatewaysCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeVpnGatewaysCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeVpnGatewaysCommon,
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

// DescribeVpnGatewaysCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPN's
// API operation DescribeVpnGatewaysCommon for usage and error information.
func (c *VPN) DescribeVpnGatewaysCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeVpnGatewaysCommonRequest(input)
	return out, req.Send()
}

// DescribeVpnGatewaysCommonWithContext is the same as DescribeVpnGatewaysCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpnGatewaysCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeVpnGatewaysCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeVpnGatewaysCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeVpnGateways = "DescribeVpnGateways"

// DescribeVpnGatewaysRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeVpnGateways operation. The "output" return
// value will be populated with the DescribeVpnGatewaysCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeVpnGatewaysCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeVpnGatewaysCommon Send returns without error.
//
// See DescribeVpnGateways for more information on using the DescribeVpnGateways
// API call, and error handling.
//
//    // Example sending a request using the DescribeVpnGatewaysRequest method.
//    req, resp := client.DescribeVpnGatewaysRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) DescribeVpnGatewaysRequest(input *DescribeVpnGatewaysInput) (req *request.Request, output *DescribeVpnGatewaysOutput) {
	op := &request.Operation{
		Name:       opDescribeVpnGateways,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeVpnGatewaysInput{}
	}

	output = &DescribeVpnGatewaysOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeVpnGateways API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPN's
// API operation DescribeVpnGateways for usage and error information.
func (c *VPN) DescribeVpnGateways(input *DescribeVpnGatewaysInput) (*DescribeVpnGatewaysOutput, error) {
	req, out := c.DescribeVpnGatewaysRequest(input)
	return out, req.Send()
}

// DescribeVpnGatewaysWithContext is the same as DescribeVpnGateways with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeVpnGateways for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) DescribeVpnGatewaysWithContext(ctx volcengine.Context, input *DescribeVpnGatewaysInput, opts ...request.Option) (*DescribeVpnGatewaysOutput, error) {
	req, out := c.DescribeVpnGatewaysRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeVpnGatewaysInput struct {
	_ struct{} `type:"structure"`

	IpAddress *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	VpcId *string `type:"string"`

	VpnGatewayIds []*string `type:"list"`

	VpnGatewayName *string `type:"string"`
}

// String returns the string representation
func (s DescribeVpnGatewaysInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpnGatewaysInput) GoString() string {
	return s.String()
}

// SetIpAddress sets the IpAddress field's value.
func (s *DescribeVpnGatewaysInput) SetIpAddress(v string) *DescribeVpnGatewaysInput {
	s.IpAddress = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpnGatewaysInput) SetPageNumber(v int64) *DescribeVpnGatewaysInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpnGatewaysInput) SetPageSize(v int64) *DescribeVpnGatewaysInput {
	s.PageSize = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeVpnGatewaysInput) SetStatus(v string) *DescribeVpnGatewaysInput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeVpnGatewaysInput) SetSubnetId(v string) *DescribeVpnGatewaysInput {
	s.SubnetId = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeVpnGatewaysInput) SetVpcId(v string) *DescribeVpnGatewaysInput {
	s.VpcId = &v
	return s
}

// SetVpnGatewayIds sets the VpnGatewayIds field's value.
func (s *DescribeVpnGatewaysInput) SetVpnGatewayIds(v []*string) *DescribeVpnGatewaysInput {
	s.VpnGatewayIds = v
	return s
}

// SetVpnGatewayName sets the VpnGatewayName field's value.
func (s *DescribeVpnGatewaysInput) SetVpnGatewayName(v string) *DescribeVpnGatewaysInput {
	s.VpnGatewayName = &v
	return s
}

type DescribeVpnGatewaysOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`

	VpnGateways []*VpnGatewayForDescribeVpnGatewaysOutput `type:"list"`
}

// String returns the string representation
func (s DescribeVpnGatewaysOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeVpnGatewaysOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeVpnGatewaysOutput) SetPageNumber(v int64) *DescribeVpnGatewaysOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeVpnGatewaysOutput) SetPageSize(v int64) *DescribeVpnGatewaysOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeVpnGatewaysOutput) SetRequestId(v string) *DescribeVpnGatewaysOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeVpnGatewaysOutput) SetTotalCount(v int64) *DescribeVpnGatewaysOutput {
	s.TotalCount = &v
	return s
}

// SetVpnGateways sets the VpnGateways field's value.
func (s *DescribeVpnGatewaysOutput) SetVpnGateways(v []*VpnGatewayForDescribeVpnGatewaysOutput) *DescribeVpnGatewaysOutput {
	s.VpnGateways = v
	return s
}

type VpnGatewayForDescribeVpnGatewaysOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	ConnectionCount *int64 `type:"integer"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	ExpiredTime *string `type:"string"`

	IpAddress *string `type:"string"`

	LockReason *string `type:"string"`

	RouteCount *int64 `type:"integer"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`

	VpnGatewayId *string `type:"string"`

	VpnGatewayName *string `type:"string"`
}

// String returns the string representation
func (s VpnGatewayForDescribeVpnGatewaysOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpnGatewayForDescribeVpnGatewaysOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetAccountId(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.AccountId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetBandwidth(v int64) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetBillingType(v int64) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetBusinessStatus(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.BusinessStatus = &v
	return s
}

// SetConnectionCount sets the ConnectionCount field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetConnectionCount(v int64) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.ConnectionCount = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetCreationTime(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetDeletedTime(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetDescription(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.Description = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetExpiredTime(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.ExpiredTime = &v
	return s
}

// SetIpAddress sets the IpAddress field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetIpAddress(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.IpAddress = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetLockReason(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.LockReason = &v
	return s
}

// SetRouteCount sets the RouteCount field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetRouteCount(v int64) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.RouteCount = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetStatus(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetSubnetId(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.SubnetId = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetUpdateTime(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetVpcId(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.VpcId = &v
	return s
}

// SetVpnGatewayId sets the VpnGatewayId field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetVpnGatewayId(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.VpnGatewayId = &v
	return s
}

// SetVpnGatewayName sets the VpnGatewayName field's value.
func (s *VpnGatewayForDescribeVpnGatewaysOutput) SetVpnGatewayName(v string) *VpnGatewayForDescribeVpnGatewaysOutput {
	s.VpnGatewayName = &v
	return s
}
