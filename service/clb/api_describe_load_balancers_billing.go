// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeLoadBalancersBillingCommon = "DescribeLoadBalancersBilling"

// DescribeLoadBalancersBillingCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLoadBalancersBillingCommon operation. The "output" return
// value will be populated with the DescribeLoadBalancersBillingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLoadBalancersBillingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLoadBalancersBillingCommon Send returns without error.
//
// See DescribeLoadBalancersBillingCommon for more information on using the DescribeLoadBalancersBillingCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeLoadBalancersBillingCommonRequest method.
//    req, resp := client.DescribeLoadBalancersBillingCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeLoadBalancersBillingCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeLoadBalancersBillingCommon,
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

// DescribeLoadBalancersBillingCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation DescribeLoadBalancersBillingCommon for usage and error information.
func (c *CLB) DescribeLoadBalancersBillingCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeLoadBalancersBillingCommonRequest(input)
	return out, req.Send()
}

// DescribeLoadBalancersBillingCommonWithContext is the same as DescribeLoadBalancersBillingCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLoadBalancersBillingCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeLoadBalancersBillingCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeLoadBalancersBillingCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeLoadBalancersBilling = "DescribeLoadBalancersBilling"

// DescribeLoadBalancersBillingRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLoadBalancersBilling operation. The "output" return
// value will be populated with the DescribeLoadBalancersBillingCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLoadBalancersBillingCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLoadBalancersBillingCommon Send returns without error.
//
// See DescribeLoadBalancersBilling for more information on using the DescribeLoadBalancersBilling
// API call, and error handling.
//
//    // Example sending a request using the DescribeLoadBalancersBillingRequest method.
//    req, resp := client.DescribeLoadBalancersBillingRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) DescribeLoadBalancersBillingRequest(input *DescribeLoadBalancersBillingInput) (req *request.Request, output *DescribeLoadBalancersBillingOutput) {
	op := &request.Operation{
		Name:       opDescribeLoadBalancersBilling,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLoadBalancersBillingInput{}
	}

	output = &DescribeLoadBalancersBillingOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeLoadBalancersBilling API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation DescribeLoadBalancersBilling for usage and error information.
func (c *CLB) DescribeLoadBalancersBilling(input *DescribeLoadBalancersBillingInput) (*DescribeLoadBalancersBillingOutput, error) {
	req, out := c.DescribeLoadBalancersBillingRequest(input)
	return out, req.Send()
}

// DescribeLoadBalancersBillingWithContext is the same as DescribeLoadBalancersBilling with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLoadBalancersBilling for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeLoadBalancersBillingWithContext(ctx volcengine.Context, input *DescribeLoadBalancersBillingInput, opts ...request.Option) (*DescribeLoadBalancersBillingOutput, error) {
	req, out := c.DescribeLoadBalancersBillingRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeLoadBalancersBillingInput struct {
	_ struct{} `type:"structure"`

	LoadBalancerIds []*string `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeLoadBalancersBillingInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLoadBalancersBillingInput) GoString() string {
	return s.String()
}

// SetLoadBalancerIds sets the LoadBalancerIds field's value.
func (s *DescribeLoadBalancersBillingInput) SetLoadBalancerIds(v []*string) *DescribeLoadBalancersBillingInput {
	s.LoadBalancerIds = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLoadBalancersBillingInput) SetPageNumber(v int64) *DescribeLoadBalancersBillingInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLoadBalancersBillingInput) SetPageSize(v int64) *DescribeLoadBalancersBillingInput {
	s.PageSize = &v
	return s
}

type DescribeLoadBalancersBillingOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LoadBalancerBillingConfigs []*LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeLoadBalancersBillingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLoadBalancersBillingOutput) GoString() string {
	return s.String()
}

// SetLoadBalancerBillingConfigs sets the LoadBalancerBillingConfigs field's value.
func (s *DescribeLoadBalancersBillingOutput) SetLoadBalancerBillingConfigs(v []*LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) *DescribeLoadBalancersBillingOutput {
	s.LoadBalancerBillingConfigs = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLoadBalancersBillingOutput) SetPageNumber(v int64) *DescribeLoadBalancersBillingOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLoadBalancersBillingOutput) SetPageSize(v int64) *DescribeLoadBalancersBillingOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeLoadBalancersBillingOutput) SetRequestId(v string) *DescribeLoadBalancersBillingOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeLoadBalancersBillingOutput) SetTotalCount(v int64) *DescribeLoadBalancersBillingOutput {
	s.TotalCount = &v
	return s
}

type LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput struct {
	_ struct{} `type:"structure"`

	BillingType *int64 `type:"integer"`

	ExpiredTime *string `type:"string"`

	InstanceStatus *int64 `type:"integer"`

	LoadBalancerId *string `type:"string"`

	OverdueReclaimTime *string `type:"string"`

	OverdueTime *string `type:"string"`

	ReclaimTime *string `type:"string"`

	RemainRenewTimes *int64 `type:"integer"`

	RenewPeriodTimes *int64 `type:"integer"`

	RenewType *int64 `type:"integer"`
}

// String returns the string representation
func (s LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) GoString() string {
	return s.String()
}

// SetBillingType sets the BillingType field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetBillingType(v int64) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.BillingType = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetExpiredTime(v string) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.ExpiredTime = &v
	return s
}

// SetInstanceStatus sets the InstanceStatus field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetInstanceStatus(v int64) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.InstanceStatus = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetLoadBalancerId(v string) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.LoadBalancerId = &v
	return s
}

// SetOverdueReclaimTime sets the OverdueReclaimTime field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetOverdueReclaimTime(v string) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.OverdueReclaimTime = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetOverdueTime(v string) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.OverdueTime = &v
	return s
}

// SetReclaimTime sets the ReclaimTime field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetReclaimTime(v string) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.ReclaimTime = &v
	return s
}

// SetRemainRenewTimes sets the RemainRenewTimes field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetRemainRenewTimes(v int64) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.RemainRenewTimes = &v
	return s
}

// SetRenewPeriodTimes sets the RenewPeriodTimes field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetRenewPeriodTimes(v int64) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.RenewPeriodTimes = &v
	return s
}

// SetRenewType sets the RenewType field's value.
func (s *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput) SetRenewType(v int64) *LoadBalancerBillingConfigForDescribeLoadBalancersBillingOutput {
	s.RenewType = &v
	return s
}
