// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeConsumedTopicsCommon = "DescribeConsumedTopics"

// DescribeConsumedTopicsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeConsumedTopicsCommon operation. The "output" return
// value will be populated with the DescribeConsumedTopicsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeConsumedTopicsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeConsumedTopicsCommon Send returns without error.
//
// See DescribeConsumedTopicsCommon for more information on using the DescribeConsumedTopicsCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeConsumedTopicsCommonRequest method.
//	req, resp := client.DescribeConsumedTopicsCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) DescribeConsumedTopicsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeConsumedTopicsCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeConsumedTopicsCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeConsumedTopicsCommon for usage and error information.
func (c *KAFKA) DescribeConsumedTopicsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeConsumedTopicsCommonRequest(input)
	return out, req.Send()
}

// DescribeConsumedTopicsCommonWithContext is the same as DescribeConsumedTopicsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeConsumedTopicsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeConsumedTopicsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeConsumedTopicsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeConsumedTopics = "DescribeConsumedTopics"

// DescribeConsumedTopicsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeConsumedTopics operation. The "output" return
// value will be populated with the DescribeConsumedTopicsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeConsumedTopicsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeConsumedTopicsCommon Send returns without error.
//
// See DescribeConsumedTopics for more information on using the DescribeConsumedTopics
// API call, and error handling.
//
//	// Example sending a request using the DescribeConsumedTopicsRequest method.
//	req, resp := client.DescribeConsumedTopicsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) DescribeConsumedTopicsRequest(input *DescribeConsumedTopicsInput) (req *request.Request, output *DescribeConsumedTopicsOutput) {
	op := &request.Operation{
		Name:       opDescribeConsumedTopics,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeConsumedTopicsInput{}
	}

	output = &DescribeConsumedTopicsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeConsumedTopics API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeConsumedTopics for usage and error information.
func (c *KAFKA) DescribeConsumedTopics(input *DescribeConsumedTopicsInput) (*DescribeConsumedTopicsOutput, error) {
	req, out := c.DescribeConsumedTopicsRequest(input)
	return out, req.Send()
}

// DescribeConsumedTopicsWithContext is the same as DescribeConsumedTopics with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeConsumedTopics for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeConsumedTopicsWithContext(ctx volcengine.Context, input *DescribeConsumedTopicsInput, opts ...request.Option) (*DescribeConsumedTopicsOutput, error) {
	req, out := c.DescribeConsumedTopicsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ConsumedTopicsInfoForDescribeConsumedTopicsOutput struct {
	_ struct{} `type:"structure"`

	Accumulation *int64 `type:"int64"`

	TopicName *string `type:"string"`
}

// String returns the string representation
func (s ConsumedTopicsInfoForDescribeConsumedTopicsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ConsumedTopicsInfoForDescribeConsumedTopicsOutput) GoString() string {
	return s.String()
}

// SetAccumulation sets the Accumulation field's value.
func (s *ConsumedTopicsInfoForDescribeConsumedTopicsOutput) SetAccumulation(v int64) *ConsumedTopicsInfoForDescribeConsumedTopicsOutput {
	s.Accumulation = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *ConsumedTopicsInfoForDescribeConsumedTopicsOutput) SetTopicName(v string) *ConsumedTopicsInfoForDescribeConsumedTopicsOutput {
	s.TopicName = &v
	return s
}

type DescribeConsumedTopicsInput struct {
	_ struct{} `type:"structure"`

	GroupId *string `type:"string"`

	InstanceId *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TopicName *string `type:"string"`
}

// String returns the string representation
func (s DescribeConsumedTopicsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeConsumedTopicsInput) GoString() string {
	return s.String()
}

// SetGroupId sets the GroupId field's value.
func (s *DescribeConsumedTopicsInput) SetGroupId(v string) *DescribeConsumedTopicsInput {
	s.GroupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeConsumedTopicsInput) SetInstanceId(v string) *DescribeConsumedTopicsInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeConsumedTopicsInput) SetPageNumber(v int32) *DescribeConsumedTopicsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeConsumedTopicsInput) SetPageSize(v int32) *DescribeConsumedTopicsInput {
	s.PageSize = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *DescribeConsumedTopicsInput) SetTopicName(v string) *DescribeConsumedTopicsInput {
	s.TopicName = &v
	return s
}

type DescribeConsumedTopicsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Accumulation *int64 `type:"int64"`

	ConsumedTopicsInfo []*ConsumedTopicsInfoForDescribeConsumedTopicsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeConsumedTopicsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeConsumedTopicsOutput) GoString() string {
	return s.String()
}

// SetAccumulation sets the Accumulation field's value.
func (s *DescribeConsumedTopicsOutput) SetAccumulation(v int64) *DescribeConsumedTopicsOutput {
	s.Accumulation = &v
	return s
}

// SetConsumedTopicsInfo sets the ConsumedTopicsInfo field's value.
func (s *DescribeConsumedTopicsOutput) SetConsumedTopicsInfo(v []*ConsumedTopicsInfoForDescribeConsumedTopicsOutput) *DescribeConsumedTopicsOutput {
	s.ConsumedTopicsInfo = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeConsumedTopicsOutput) SetTotal(v int32) *DescribeConsumedTopicsOutput {
	s.Total = &v
	return s
}
