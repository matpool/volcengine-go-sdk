// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeGroupsCommon = "DescribeGroups"

// DescribeGroupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeGroupsCommon operation. The "output" return
// value will be populated with the DescribeGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeGroupsCommon Send returns without error.
//
// See DescribeGroupsCommon for more information on using the DescribeGroupsCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeGroupsCommonRequest method.
//	req, resp := client.DescribeGroupsCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) DescribeGroupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeGroupsCommon,
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

// DescribeGroupsCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeGroupsCommon for usage and error information.
func (c *KAFKA) DescribeGroupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeGroupsCommonRequest(input)
	return out, req.Send()
}

// DescribeGroupsCommonWithContext is the same as DescribeGroupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeGroupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeGroupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeGroupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeGroups = "DescribeGroups"

// DescribeGroupsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeGroups operation. The "output" return
// value will be populated with the DescribeGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeGroupsCommon Send returns without error.
//
// See DescribeGroups for more information on using the DescribeGroups
// API call, and error handling.
//
//	// Example sending a request using the DescribeGroupsRequest method.
//	req, resp := client.DescribeGroupsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) DescribeGroupsRequest(input *DescribeGroupsInput) (req *request.Request, output *DescribeGroupsOutput) {
	op := &request.Operation{
		Name:       opDescribeGroups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeGroupsInput{}
	}

	output = &DescribeGroupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeGroups API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation DescribeGroups for usage and error information.
func (c *KAFKA) DescribeGroups(input *DescribeGroupsInput) (*DescribeGroupsOutput, error) {
	req, out := c.DescribeGroupsRequest(input)
	return out, req.Send()
}

// DescribeGroupsWithContext is the same as DescribeGroups with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeGroups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) DescribeGroupsWithContext(ctx volcengine.Context, input *DescribeGroupsInput, opts ...request.Option) (*DescribeGroupsOutput, error) {
	req, out := c.DescribeGroupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeGroupsInput struct {
	_ struct{} `type:"structure"`

	GroupId *string `type:"string"`

	InstanceId *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeGroupsInput) GoString() string {
	return s.String()
}

// SetGroupId sets the GroupId field's value.
func (s *DescribeGroupsInput) SetGroupId(v string) *DescribeGroupsInput {
	s.GroupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeGroupsInput) SetInstanceId(v string) *DescribeGroupsInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeGroupsInput) SetPageNumber(v int32) *DescribeGroupsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeGroupsInput) SetPageSize(v int32) *DescribeGroupsInput {
	s.PageSize = &v
	return s
}

type DescribeGroupsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	GroupsInfo []*GroupsInfoForDescribeGroupsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeGroupsOutput) GoString() string {
	return s.String()
}

// SetGroupsInfo sets the GroupsInfo field's value.
func (s *DescribeGroupsOutput) SetGroupsInfo(v []*GroupsInfoForDescribeGroupsOutput) *DescribeGroupsOutput {
	s.GroupsInfo = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeGroupsOutput) SetTotal(v int32) *DescribeGroupsOutput {
	s.Total = &v
	return s
}

type GroupsInfoForDescribeGroupsOutput struct {
	_ struct{} `type:"structure"`

	GroupId *string `type:"string"`

	State *string `type:"string"`
}

// String returns the string representation
func (s GroupsInfoForDescribeGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GroupsInfoForDescribeGroupsOutput) GoString() string {
	return s.String()
}

// SetGroupId sets the GroupId field's value.
func (s *GroupsInfoForDescribeGroupsOutput) SetGroupId(v string) *GroupsInfoForDescribeGroupsOutput {
	s.GroupId = &v
	return s
}

// SetState sets the State field's value.
func (s *GroupsInfoForDescribeGroupsOutput) SetState(v string) *GroupsInfoForDescribeGroupsOutput {
	s.State = &v
	return s
}
