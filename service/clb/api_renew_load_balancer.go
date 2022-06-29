// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opRenewLoadBalancerCommon = "RenewLoadBalancer"

// RenewLoadBalancerCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the RenewLoadBalancerCommon operation. The "output" return
// value will be populated with the RenewLoadBalancerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RenewLoadBalancerCommon Request to send the API call to the service.
// the "output" return value is not valid until after RenewLoadBalancerCommon Send returns without error.
//
// See RenewLoadBalancerCommon for more information on using the RenewLoadBalancerCommon
// API call, and error handling.
//
//    // Example sending a request using the RenewLoadBalancerCommonRequest method.
//    req, resp := client.RenewLoadBalancerCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) RenewLoadBalancerCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opRenewLoadBalancerCommon,
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

// RenewLoadBalancerCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation RenewLoadBalancerCommon for usage and error information.
func (c *CLB) RenewLoadBalancerCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.RenewLoadBalancerCommonRequest(input)
	return out, req.Send()
}

// RenewLoadBalancerCommonWithContext is the same as RenewLoadBalancerCommon with the addition of
// the ability to pass a context and additional request options.
//
// See RenewLoadBalancerCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) RenewLoadBalancerCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.RenewLoadBalancerCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRenewLoadBalancer = "RenewLoadBalancer"

// RenewLoadBalancerRequest generates a "volcengine/request.Request" representing the
// client's request for the RenewLoadBalancer operation. The "output" return
// value will be populated with the RenewLoadBalancerCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned RenewLoadBalancerCommon Request to send the API call to the service.
// the "output" return value is not valid until after RenewLoadBalancerCommon Send returns without error.
//
// See RenewLoadBalancer for more information on using the RenewLoadBalancer
// API call, and error handling.
//
//    // Example sending a request using the RenewLoadBalancerRequest method.
//    req, resp := client.RenewLoadBalancerRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CLB) RenewLoadBalancerRequest(input *RenewLoadBalancerInput) (req *request.Request, output *RenewLoadBalancerOutput) {
	op := &request.Operation{
		Name:       opRenewLoadBalancer,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RenewLoadBalancerInput{}
	}

	output = &RenewLoadBalancerOutput{}
	req = c.newRequest(op, input, output)

	return
}

// RenewLoadBalancer API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CLB's
// API operation RenewLoadBalancer for usage and error information.
func (c *CLB) RenewLoadBalancer(input *RenewLoadBalancerInput) (*RenewLoadBalancerOutput, error) {
	req, out := c.RenewLoadBalancerRequest(input)
	return out, req.Send()
}

// RenewLoadBalancerWithContext is the same as RenewLoadBalancer with the addition of
// the ability to pass a context and additional request options.
//
// See RenewLoadBalancer for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) RenewLoadBalancerWithContext(ctx volcengine.Context, input *RenewLoadBalancerInput, opts ...request.Option) (*RenewLoadBalancerOutput, error) {
	req, out := c.RenewLoadBalancerRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type RenewLoadBalancerInput struct {
	_ struct{} `type:"structure"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`

	Period *int64 `type:"integer"`

	PeriodUnit *string `type:"string"`
}

// String returns the string representation
func (s RenewLoadBalancerInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RenewLoadBalancerInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RenewLoadBalancerInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RenewLoadBalancerInput"}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *RenewLoadBalancerInput) SetLoadBalancerId(v string) *RenewLoadBalancerInput {
	s.LoadBalancerId = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *RenewLoadBalancerInput) SetPeriod(v int64) *RenewLoadBalancerInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *RenewLoadBalancerInput) SetPeriodUnit(v string) *RenewLoadBalancerInput {
	s.PeriodUnit = &v
	return s
}

type RenewLoadBalancerOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s RenewLoadBalancerOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RenewLoadBalancerOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *RenewLoadBalancerOutput) SetRequestId(v string) *RenewLoadBalancerOutput {
	s.RequestId = &v
	return s
}
