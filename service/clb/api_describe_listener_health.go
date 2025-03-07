// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeListenerHealthCommon = "DescribeListenerHealth"

// DescribeListenerHealthCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeListenerHealthCommon operation. The "output" return
// value will be populated with the DescribeListenerHealthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeListenerHealthCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeListenerHealthCommon Send returns without error.
//
// See DescribeListenerHealthCommon for more information on using the DescribeListenerHealthCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeListenerHealthCommonRequest method.
//	req, resp := client.DescribeListenerHealthCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeListenerHealthCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeListenerHealthCommon,
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

// DescribeListenerHealthCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeListenerHealthCommon for usage and error information.
func (c *CLB) DescribeListenerHealthCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeListenerHealthCommonRequest(input)
	return out, req.Send()
}

// DescribeListenerHealthCommonWithContext is the same as DescribeListenerHealthCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeListenerHealthCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeListenerHealthCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeListenerHealthCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeListenerHealth = "DescribeListenerHealth"

// DescribeListenerHealthRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeListenerHealth operation. The "output" return
// value will be populated with the DescribeListenerHealthCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeListenerHealthCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeListenerHealthCommon Send returns without error.
//
// See DescribeListenerHealth for more information on using the DescribeListenerHealth
// API call, and error handling.
//
//	// Example sending a request using the DescribeListenerHealthRequest method.
//	req, resp := client.DescribeListenerHealthRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) DescribeListenerHealthRequest(input *DescribeListenerHealthInput) (req *request.Request, output *DescribeListenerHealthOutput) {
	op := &request.Operation{
		Name:       opDescribeListenerHealth,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeListenerHealthInput{}
	}

	output = &DescribeListenerHealthOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeListenerHealth API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation DescribeListenerHealth for usage and error information.
func (c *CLB) DescribeListenerHealth(input *DescribeListenerHealthInput) (*DescribeListenerHealthOutput, error) {
	req, out := c.DescribeListenerHealthRequest(input)
	return out, req.Send()
}

// DescribeListenerHealthWithContext is the same as DescribeListenerHealth with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeListenerHealth for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) DescribeListenerHealthWithContext(ctx volcengine.Context, input *DescribeListenerHealthInput, opts ...request.Option) (*DescribeListenerHealthOutput, error) {
	req, out := c.DescribeListenerHealthRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeListenerHealthInput struct {
	_ struct{} `type:"structure"`

	// ListenerId is a required field
	ListenerId *string `type:"string" required:"true"`

	OnlyUnHealthy *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeListenerHealthInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeListenerHealthInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeListenerHealthInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeListenerHealthInput"}
	if s.ListenerId == nil {
		invalidParams.Add(request.NewErrParamRequired("ListenerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetListenerId sets the ListenerId field's value.
func (s *DescribeListenerHealthInput) SetListenerId(v string) *DescribeListenerHealthInput {
	s.ListenerId = &v
	return s
}

// SetOnlyUnHealthy sets the OnlyUnHealthy field's value.
func (s *DescribeListenerHealthInput) SetOnlyUnHealthy(v string) *DescribeListenerHealthInput {
	s.OnlyUnHealthy = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeListenerHealthInput) SetPageNumber(v int64) *DescribeListenerHealthInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeListenerHealthInput) SetPageSize(v int64) *DescribeListenerHealthInput {
	s.PageSize = &v
	return s
}

type DescribeListenerHealthOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	Results []*ResultForDescribeListenerHealthOutput `type:"list"`

	Status *string `type:"string"`

	TotalCount *int64 `type:"integer"`

	UnHealthyCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeListenerHealthOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeListenerHealthOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeListenerHealthOutput) SetPageNumber(v int64) *DescribeListenerHealthOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeListenerHealthOutput) SetPageSize(v int64) *DescribeListenerHealthOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeListenerHealthOutput) SetRequestId(v string) *DescribeListenerHealthOutput {
	s.RequestId = &v
	return s
}

// SetResults sets the Results field's value.
func (s *DescribeListenerHealthOutput) SetResults(v []*ResultForDescribeListenerHealthOutput) *DescribeListenerHealthOutput {
	s.Results = v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeListenerHealthOutput) SetStatus(v string) *DescribeListenerHealthOutput {
	s.Status = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeListenerHealthOutput) SetTotalCount(v int64) *DescribeListenerHealthOutput {
	s.TotalCount = &v
	return s
}

// SetUnHealthyCount sets the UnHealthyCount field's value.
func (s *DescribeListenerHealthOutput) SetUnHealthyCount(v int64) *DescribeListenerHealthOutput {
	s.UnHealthyCount = &v
	return s
}

type ResultForDescribeListenerHealthOutput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	Ip *string `type:"string"`

	Port *int64 `type:"integer"`

	RuleNumber *int64 `type:"integer"`

	ServerGroupId *string `type:"string"`

	ServerId *string `type:"string"`

	Status *string `type:"string"`

	Type *string `type:"string"`
}

// String returns the string representation
func (s ResultForDescribeListenerHealthOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResultForDescribeListenerHealthOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ResultForDescribeListenerHealthOutput) SetInstanceId(v string) *ResultForDescribeListenerHealthOutput {
	s.InstanceId = &v
	return s
}

// SetIp sets the Ip field's value.
func (s *ResultForDescribeListenerHealthOutput) SetIp(v string) *ResultForDescribeListenerHealthOutput {
	s.Ip = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ResultForDescribeListenerHealthOutput) SetPort(v int64) *ResultForDescribeListenerHealthOutput {
	s.Port = &v
	return s
}

// SetRuleNumber sets the RuleNumber field's value.
func (s *ResultForDescribeListenerHealthOutput) SetRuleNumber(v int64) *ResultForDescribeListenerHealthOutput {
	s.RuleNumber = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ResultForDescribeListenerHealthOutput) SetServerGroupId(v string) *ResultForDescribeListenerHealthOutput {
	s.ServerGroupId = &v
	return s
}

// SetServerId sets the ServerId field's value.
func (s *ResultForDescribeListenerHealthOutput) SetServerId(v string) *ResultForDescribeListenerHealthOutput {
	s.ServerId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *ResultForDescribeListenerHealthOutput) SetStatus(v string) *ResultForDescribeListenerHealthOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *ResultForDescribeListenerHealthOutput) SetType(v string) *ResultForDescribeListenerHealthOutput {
	s.Type = &v
	return s
}
