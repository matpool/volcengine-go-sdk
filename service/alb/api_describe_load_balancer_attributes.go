// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeLoadBalancerAttributesCommon = "DescribeLoadBalancerAttributes"

// DescribeLoadBalancerAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLoadBalancerAttributesCommon operation. The "output" return
// value will be populated with the DescribeLoadBalancerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLoadBalancerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLoadBalancerAttributesCommon Send returns without error.
//
// See DescribeLoadBalancerAttributesCommon for more information on using the DescribeLoadBalancerAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeLoadBalancerAttributesCommonRequest method.
//    req, resp := client.DescribeLoadBalancerAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeLoadBalancerAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeLoadBalancerAttributesCommon,
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

// DescribeLoadBalancerAttributesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeLoadBalancerAttributesCommon for usage and error information.
func (c *ALB) DescribeLoadBalancerAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeLoadBalancerAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeLoadBalancerAttributesCommonWithContext is the same as DescribeLoadBalancerAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLoadBalancerAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeLoadBalancerAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeLoadBalancerAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeLoadBalancerAttributes = "DescribeLoadBalancerAttributes"

// DescribeLoadBalancerAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLoadBalancerAttributes operation. The "output" return
// value will be populated with the DescribeLoadBalancerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLoadBalancerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLoadBalancerAttributesCommon Send returns without error.
//
// See DescribeLoadBalancerAttributes for more information on using the DescribeLoadBalancerAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeLoadBalancerAttributesRequest method.
//    req, resp := client.DescribeLoadBalancerAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) DescribeLoadBalancerAttributesRequest(input *DescribeLoadBalancerAttributesInput) (req *request.Request, output *DescribeLoadBalancerAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeLoadBalancerAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLoadBalancerAttributesInput{}
	}

	output = &DescribeLoadBalancerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeLoadBalancerAttributes API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DescribeLoadBalancerAttributes for usage and error information.
func (c *ALB) DescribeLoadBalancerAttributes(input *DescribeLoadBalancerAttributesInput) (*DescribeLoadBalancerAttributesOutput, error) {
	req, out := c.DescribeLoadBalancerAttributesRequest(input)
	return out, req.Send()
}

// DescribeLoadBalancerAttributesWithContext is the same as DescribeLoadBalancerAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLoadBalancerAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DescribeLoadBalancerAttributesWithContext(ctx volcengine.Context, input *DescribeLoadBalancerAttributesInput, opts ...request.Option) (*DescribeLoadBalancerAttributesOutput, error) {
	req, out := c.DescribeLoadBalancerAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccessLogForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	BucketName *string `type:"string"`

	Enabled *bool `type:"boolean"`
}

// String returns the string representation
func (s AccessLogForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccessLogForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetBucketName sets the BucketName field's value.
func (s *AccessLogForDescribeLoadBalancerAttributesOutput) SetBucketName(v string) *AccessLogForDescribeLoadBalancerAttributesOutput {
	s.BucketName = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *AccessLogForDescribeLoadBalancerAttributesOutput) SetEnabled(v bool) *AccessLogForDescribeLoadBalancerAttributesOutput {
	s.Enabled = &v
	return s
}

type DescribeLoadBalancerAttributesInput struct {
	_ struct{} `type:"structure"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeLoadBalancerAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLoadBalancerAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeLoadBalancerAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeLoadBalancerAttributesInput"}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *DescribeLoadBalancerAttributesInput) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributesInput {
	s.LoadBalancerId = &v
	return s
}

type DescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccessLog *AccessLogForDescribeLoadBalancerAttributesOutput `type:"structure"`

	AddressIpVersion *string `type:"string"`

	BusinessStatus *string `type:"string"`

	CreateTime *string `type:"string"`

	DNSName *string `type:"string"`

	DeleteProtection *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	Eip *EipForDescribeLoadBalancerAttributesOutput `type:"structure"`

	EipAddress *string `type:"string"`

	EipId *string `type:"string"`

	Enabled *bool `type:"boolean"`

	EniAddress *string `type:"string"`

	EniId *string `type:"string"`

	HealthLog *HealthLogForDescribeLoadBalancerAttributesOutput `type:"structure"`

	Listeners []*ListenerForDescribeLoadBalancerAttributesOutput `type:"list"`

	LoadBalancerBillingType *int64 `type:"integer"`

	LoadBalancerId *string `type:"string"`

	LoadBalancerName *string `type:"string"`

	LocalAddresses []*string `type:"list"`

	LockReason *string `type:"string"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`

	TLSAccessLog *TLSAccessLogForDescribeLoadBalancerAttributesOutput `type:"structure"`

	Type *string `type:"string"`

	UpdateTime *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneMappings []*ZoneMappingForDescribeLoadBalancerAttributesOutput `type:"list"`
}

// String returns the string representation
func (s DescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetAccessLog sets the AccessLog field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetAccessLog(v *AccessLogForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.AccessLog = v
	return s
}

// SetAddressIpVersion sets the AddressIpVersion field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetAddressIpVersion(v string) *DescribeLoadBalancerAttributesOutput {
	s.AddressIpVersion = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetBusinessStatus(v string) *DescribeLoadBalancerAttributesOutput {
	s.BusinessStatus = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetCreateTime(v string) *DescribeLoadBalancerAttributesOutput {
	s.CreateTime = &v
	return s
}

// SetDNSName sets the DNSName field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetDNSName(v string) *DescribeLoadBalancerAttributesOutput {
	s.DNSName = &v
	return s
}

// SetDeleteProtection sets the DeleteProtection field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetDeleteProtection(v string) *DescribeLoadBalancerAttributesOutput {
	s.DeleteProtection = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetDeletedTime(v string) *DescribeLoadBalancerAttributesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetDescription(v string) *DescribeLoadBalancerAttributesOutput {
	s.Description = &v
	return s
}

// SetEip sets the Eip field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEip(v *EipForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.Eip = v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEipAddress(v string) *DescribeLoadBalancerAttributesOutput {
	s.EipAddress = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEipId(v string) *DescribeLoadBalancerAttributesOutput {
	s.EipId = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEnabled(v bool) *DescribeLoadBalancerAttributesOutput {
	s.Enabled = &v
	return s
}

// SetEniAddress sets the EniAddress field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEniAddress(v string) *DescribeLoadBalancerAttributesOutput {
	s.EniAddress = &v
	return s
}

// SetEniId sets the EniId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetEniId(v string) *DescribeLoadBalancerAttributesOutput {
	s.EniId = &v
	return s
}

// SetHealthLog sets the HealthLog field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetHealthLog(v *HealthLogForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.HealthLog = v
	return s
}

// SetListeners sets the Listeners field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetListeners(v []*ListenerForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.Listeners = v
	return s
}

// SetLoadBalancerBillingType sets the LoadBalancerBillingType field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLoadBalancerBillingType(v int64) *DescribeLoadBalancerAttributesOutput {
	s.LoadBalancerBillingType = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributesOutput {
	s.LoadBalancerId = &v
	return s
}

// SetLoadBalancerName sets the LoadBalancerName field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLoadBalancerName(v string) *DescribeLoadBalancerAttributesOutput {
	s.LoadBalancerName = &v
	return s
}

// SetLocalAddresses sets the LocalAddresses field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLocalAddresses(v []*string) *DescribeLoadBalancerAttributesOutput {
	s.LocalAddresses = v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetLockReason(v string) *DescribeLoadBalancerAttributesOutput {
	s.LockReason = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetOverdueTime(v string) *DescribeLoadBalancerAttributesOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetProjectName(v string) *DescribeLoadBalancerAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetRequestId(v string) *DescribeLoadBalancerAttributesOutput {
	s.RequestId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetStatus(v string) *DescribeLoadBalancerAttributesOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetSubnetId(v string) *DescribeLoadBalancerAttributesOutput {
	s.SubnetId = &v
	return s
}

// SetTLSAccessLog sets the TLSAccessLog field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetTLSAccessLog(v *TLSAccessLogForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.TLSAccessLog = v
	return s
}

// SetType sets the Type field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetType(v string) *DescribeLoadBalancerAttributesOutput {
	s.Type = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetUpdateTime(v string) *DescribeLoadBalancerAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetVpcId(v string) *DescribeLoadBalancerAttributesOutput {
	s.VpcId = &v
	return s
}

// SetZoneMappings sets the ZoneMappings field's value.
func (s *DescribeLoadBalancerAttributesOutput) SetZoneMappings(v []*ZoneMappingForDescribeLoadBalancerAttributesOutput) *DescribeLoadBalancerAttributesOutput {
	s.ZoneMappings = v
	return s
}

type EipForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	AssociationMode *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	EipAddress *string `type:"string"`

	EipBillingType *int64 `type:"integer"`

	EipType *string `type:"string"`

	ISP *string `type:"string"`

	PopLocations []*PopLocationForDescribeLoadBalancerAttributesOutput `type:"list"`

	SecurityProtectionTypes []*string `type:"list"`
}

// String returns the string representation
func (s EipForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetAssociationMode sets the AssociationMode field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetAssociationMode(v string) *EipForDescribeLoadBalancerAttributesOutput {
	s.AssociationMode = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetBandwidth(v int64) *EipForDescribeLoadBalancerAttributesOutput {
	s.Bandwidth = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetEipAddress(v string) *EipForDescribeLoadBalancerAttributesOutput {
	s.EipAddress = &v
	return s
}

// SetEipBillingType sets the EipBillingType field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetEipBillingType(v int64) *EipForDescribeLoadBalancerAttributesOutput {
	s.EipBillingType = &v
	return s
}

// SetEipType sets the EipType field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetEipType(v string) *EipForDescribeLoadBalancerAttributesOutput {
	s.EipType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetISP(v string) *EipForDescribeLoadBalancerAttributesOutput {
	s.ISP = &v
	return s
}

// SetPopLocations sets the PopLocations field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetPopLocations(v []*PopLocationForDescribeLoadBalancerAttributesOutput) *EipForDescribeLoadBalancerAttributesOutput {
	s.PopLocations = v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *EipForDescribeLoadBalancerAttributesOutput) SetSecurityProtectionTypes(v []*string) *EipForDescribeLoadBalancerAttributesOutput {
	s.SecurityProtectionTypes = v
	return s
}

type HealthLogForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Enabled *bool `type:"boolean"`

	ProjectId *string `type:"string"`

	TopicId *string `type:"string"`
}

// String returns the string representation
func (s HealthLogForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s HealthLogForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetEnabled sets the Enabled field's value.
func (s *HealthLogForDescribeLoadBalancerAttributesOutput) SetEnabled(v bool) *HealthLogForDescribeLoadBalancerAttributesOutput {
	s.Enabled = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *HealthLogForDescribeLoadBalancerAttributesOutput) SetProjectId(v string) *HealthLogForDescribeLoadBalancerAttributesOutput {
	s.ProjectId = &v
	return s
}

// SetTopicId sets the TopicId field's value.
func (s *HealthLogForDescribeLoadBalancerAttributesOutput) SetTopicId(v string) *HealthLogForDescribeLoadBalancerAttributesOutput {
	s.TopicId = &v
	return s
}

type Ipv6EipForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `type:"integer"`

	BillingType *int64 `type:"integer"`

	ISP *string `type:"string"`
}

// String returns the string representation
func (s Ipv6EipForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s Ipv6EipForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *Ipv6EipForDescribeLoadBalancerAttributesOutput) SetBandwidth(v int64) *Ipv6EipForDescribeLoadBalancerAttributesOutput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *Ipv6EipForDescribeLoadBalancerAttributesOutput) SetBillingType(v int64) *Ipv6EipForDescribeLoadBalancerAttributesOutput {
	s.BillingType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *Ipv6EipForDescribeLoadBalancerAttributesOutput) SetISP(v string) *Ipv6EipForDescribeLoadBalancerAttributesOutput {
	s.ISP = &v
	return s
}

type ListenerForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	ListenerId *string `type:"string"`

	ListenerName *string `type:"string"`
}

// String returns the string representation
func (s ListenerForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListenerForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetListenerId sets the ListenerId field's value.
func (s *ListenerForDescribeLoadBalancerAttributesOutput) SetListenerId(v string) *ListenerForDescribeLoadBalancerAttributesOutput {
	s.ListenerId = &v
	return s
}

// SetListenerName sets the ListenerName field's value.
func (s *ListenerForDescribeLoadBalancerAttributesOutput) SetListenerName(v string) *ListenerForDescribeLoadBalancerAttributesOutput {
	s.ListenerName = &v
	return s
}

type LoadBalancerAddressForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Eip *EipForDescribeLoadBalancerAttributesOutput `type:"structure"`

	EipAddress *string `type:"string"`

	EipId *string `type:"string"`

	EniAddress *string `type:"string"`

	EniId *string `type:"string"`

	EniIpv6Address *string `type:"string"`

	Ipv6Eip *Ipv6EipForDescribeLoadBalancerAttributesOutput `type:"structure"`

	Ipv6EipId *string `type:"string"`
}

// String returns the string representation
func (s LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetEip sets the Eip field's value.
func (s *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) SetEip(v *EipForDescribeLoadBalancerAttributesOutput) *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput {
	s.Eip = v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) SetEipAddress(v string) *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput {
	s.EipAddress = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) SetEipId(v string) *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput {
	s.EipId = &v
	return s
}

// SetEniAddress sets the EniAddress field's value.
func (s *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) SetEniAddress(v string) *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput {
	s.EniAddress = &v
	return s
}

// SetEniId sets the EniId field's value.
func (s *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) SetEniId(v string) *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput {
	s.EniId = &v
	return s
}

// SetEniIpv6Address sets the EniIpv6Address field's value.
func (s *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) SetEniIpv6Address(v string) *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput {
	s.EniIpv6Address = &v
	return s
}

// SetIpv6Eip sets the Ipv6Eip field's value.
func (s *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) SetIpv6Eip(v *Ipv6EipForDescribeLoadBalancerAttributesOutput) *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput {
	s.Ipv6Eip = v
	return s
}

// SetIpv6EipId sets the Ipv6EipId field's value.
func (s *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) SetIpv6EipId(v string) *LoadBalancerAddressForDescribeLoadBalancerAttributesOutput {
	s.Ipv6EipId = &v
	return s
}

type PopLocationForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	PopId *string `type:"string"`

	PopName *string `type:"string"`
}

// String returns the string representation
func (s PopLocationForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s PopLocationForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetPopId sets the PopId field's value.
func (s *PopLocationForDescribeLoadBalancerAttributesOutput) SetPopId(v string) *PopLocationForDescribeLoadBalancerAttributesOutput {
	s.PopId = &v
	return s
}

// SetPopName sets the PopName field's value.
func (s *PopLocationForDescribeLoadBalancerAttributesOutput) SetPopName(v string) *PopLocationForDescribeLoadBalancerAttributesOutput {
	s.PopName = &v
	return s
}

type TLSAccessLogForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Enabled *bool `type:"boolean"`

	ProjectId *string `type:"string"`

	TopicId *string `type:"string"`
}

// String returns the string representation
func (s TLSAccessLogForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TLSAccessLogForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetEnabled sets the Enabled field's value.
func (s *TLSAccessLogForDescribeLoadBalancerAttributesOutput) SetEnabled(v bool) *TLSAccessLogForDescribeLoadBalancerAttributesOutput {
	s.Enabled = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *TLSAccessLogForDescribeLoadBalancerAttributesOutput) SetProjectId(v string) *TLSAccessLogForDescribeLoadBalancerAttributesOutput {
	s.ProjectId = &v
	return s
}

// SetTopicId sets the TopicId field's value.
func (s *TLSAccessLogForDescribeLoadBalancerAttributesOutput) SetTopicId(v string) *TLSAccessLogForDescribeLoadBalancerAttributesOutput {
	s.TopicId = &v
	return s
}

type ZoneMappingForDescribeLoadBalancerAttributesOutput struct {
	_ struct{} `type:"structure"`

	LoadBalancerAddresses []*LoadBalancerAddressForDescribeLoadBalancerAttributesOutput `type:"list"`

	SubnetId *string `type:"string"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ZoneMappingForDescribeLoadBalancerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ZoneMappingForDescribeLoadBalancerAttributesOutput) GoString() string {
	return s.String()
}

// SetLoadBalancerAddresses sets the LoadBalancerAddresses field's value.
func (s *ZoneMappingForDescribeLoadBalancerAttributesOutput) SetLoadBalancerAddresses(v []*LoadBalancerAddressForDescribeLoadBalancerAttributesOutput) *ZoneMappingForDescribeLoadBalancerAttributesOutput {
	s.LoadBalancerAddresses = v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *ZoneMappingForDescribeLoadBalancerAttributesOutput) SetSubnetId(v string) *ZoneMappingForDescribeLoadBalancerAttributesOutput {
	s.SubnetId = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ZoneMappingForDescribeLoadBalancerAttributesOutput) SetZoneId(v string) *ZoneMappingForDescribeLoadBalancerAttributesOutput {
	s.ZoneId = &v
	return s
}
