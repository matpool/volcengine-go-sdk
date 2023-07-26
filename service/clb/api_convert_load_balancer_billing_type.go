// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opConvertLoadBalancerBillingTypeCommon = "ConvertLoadBalancerBillingType"

// ConvertLoadBalancerBillingTypeCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ConvertLoadBalancerBillingTypeCommon operation. The "output" return
// value will be populated with the ConvertLoadBalancerBillingTypeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ConvertLoadBalancerBillingTypeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ConvertLoadBalancerBillingTypeCommon Send returns without error.
//
// See ConvertLoadBalancerBillingTypeCommon for more information on using the ConvertLoadBalancerBillingTypeCommon
// API call, and error handling.
//
//	// Example sending a request using the ConvertLoadBalancerBillingTypeCommonRequest method.
//	req, resp := client.ConvertLoadBalancerBillingTypeCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) ConvertLoadBalancerBillingTypeCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opConvertLoadBalancerBillingTypeCommon,
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

// ConvertLoadBalancerBillingTypeCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ConvertLoadBalancerBillingTypeCommon for usage and error information.
func (c *CLB) ConvertLoadBalancerBillingTypeCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ConvertLoadBalancerBillingTypeCommonRequest(input)
	return out, req.Send()
}

// ConvertLoadBalancerBillingTypeCommonWithContext is the same as ConvertLoadBalancerBillingTypeCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ConvertLoadBalancerBillingTypeCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) ConvertLoadBalancerBillingTypeCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ConvertLoadBalancerBillingTypeCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opConvertLoadBalancerBillingType = "ConvertLoadBalancerBillingType"

// ConvertLoadBalancerBillingTypeRequest generates a "volcengine/request.Request" representing the
// client's request for the ConvertLoadBalancerBillingType operation. The "output" return
// value will be populated with the ConvertLoadBalancerBillingTypeCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ConvertLoadBalancerBillingTypeCommon Request to send the API call to the service.
// the "output" return value is not valid until after ConvertLoadBalancerBillingTypeCommon Send returns without error.
//
// See ConvertLoadBalancerBillingType for more information on using the ConvertLoadBalancerBillingType
// API call, and error handling.
//
//	// Example sending a request using the ConvertLoadBalancerBillingTypeRequest method.
//	req, resp := client.ConvertLoadBalancerBillingTypeRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) ConvertLoadBalancerBillingTypeRequest(input *ConvertLoadBalancerBillingTypeInput) (req *request.Request, output *ConvertLoadBalancerBillingTypeOutput) {
	op := &request.Operation{
		Name:       opConvertLoadBalancerBillingType,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ConvertLoadBalancerBillingTypeInput{}
	}

	output = &ConvertLoadBalancerBillingTypeOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ConvertLoadBalancerBillingType API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ConvertLoadBalancerBillingType for usage and error information.
func (c *CLB) ConvertLoadBalancerBillingType(input *ConvertLoadBalancerBillingTypeInput) (*ConvertLoadBalancerBillingTypeOutput, error) {
	req, out := c.ConvertLoadBalancerBillingTypeRequest(input)
	return out, req.Send()
}

// ConvertLoadBalancerBillingTypeWithContext is the same as ConvertLoadBalancerBillingType with the addition of
// the ability to pass a context and additional request options.
//
// See ConvertLoadBalancerBillingType for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) ConvertLoadBalancerBillingTypeWithContext(ctx volcengine.Context, input *ConvertLoadBalancerBillingTypeInput, opts ...request.Option) (*ConvertLoadBalancerBillingTypeOutput, error) {
	req, out := c.ConvertLoadBalancerBillingTypeRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConvertLoadBalancerBillingTypeInput struct {
	_ struct{} `type:"structure"`

	// LoadBalancerBillingType is a required field
	LoadBalancerBillingType *int64 `type:"integer" required:"true"`

	// LoadBalancerId is a required field
	LoadBalancerId *string `type:"string" required:"true"`

	Period *int64 `type:"integer"`

	PeriodUnit *string `type:"string"`
}

// String returns the string representation
func (s ConvertLoadBalancerBillingTypeInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConvertLoadBalancerBillingTypeInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConvertLoadBalancerBillingTypeInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ConvertLoadBalancerBillingTypeInput"}
	if s.LoadBalancerBillingType == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerBillingType"))
	}
	if s.LoadBalancerId == nil {
		invalidParams.Add(request.NewErrParamRequired("LoadBalancerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLoadBalancerBillingType sets the LoadBalancerBillingType field's value.
func (s *ConvertLoadBalancerBillingTypeInput) SetLoadBalancerBillingType(v int64) *ConvertLoadBalancerBillingTypeInput {
	s.LoadBalancerBillingType = &v
	return s
}

// SetLoadBalancerId sets the LoadBalancerId field's value.
func (s *ConvertLoadBalancerBillingTypeInput) SetLoadBalancerId(v string) *ConvertLoadBalancerBillingTypeInput {
	s.LoadBalancerId = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ConvertLoadBalancerBillingTypeInput) SetPeriod(v int64) *ConvertLoadBalancerBillingTypeInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ConvertLoadBalancerBillingTypeInput) SetPeriodUnit(v string) *ConvertLoadBalancerBillingTypeInput {
	s.PeriodUnit = &v
	return s
}

type ConvertLoadBalancerBillingTypeOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OrderId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ConvertLoadBalancerBillingTypeOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConvertLoadBalancerBillingTypeOutput) GoString() string {
	return s.String()
}

// SetOrderId sets the OrderId field's value.
func (s *ConvertLoadBalancerBillingTypeOutput) SetOrderId(v string) *ConvertLoadBalancerBillingTypeOutput {
	s.OrderId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *ConvertLoadBalancerBillingTypeOutput) SetRequestId(v string) *ConvertLoadBalancerBillingTypeOutput {
	s.RequestId = &v
	return s
}
