// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListApplicationHistoryCommon = "listApplicationHistory"

// ListApplicationHistoryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListApplicationHistoryCommon operation. The "output" return
// value will be populated with the ListApplicationHistoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListApplicationHistoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListApplicationHistoryCommon Send returns without error.
//
// See ListApplicationHistoryCommon for more information on using the ListApplicationHistoryCommon
// API call, and error handling.
//
//	// Example sending a request using the ListApplicationHistoryCommonRequest method.
//	req, resp := client.ListApplicationHistoryCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) ListApplicationHistoryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListApplicationHistoryCommon,
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

// ListApplicationHistoryCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation ListApplicationHistoryCommon for usage and error information.
func (c *SPARK) ListApplicationHistoryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListApplicationHistoryCommonRequest(input)
	return out, req.Send()
}

// ListApplicationHistoryCommonWithContext is the same as ListApplicationHistoryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListApplicationHistoryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) ListApplicationHistoryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListApplicationHistoryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListApplicationHistory = "listApplicationHistory"

// ListApplicationHistoryRequest generates a "volcengine/request.Request" representing the
// client's request for the ListApplicationHistory operation. The "output" return
// value will be populated with the ListApplicationHistoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListApplicationHistoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListApplicationHistoryCommon Send returns without error.
//
// See ListApplicationHistory for more information on using the ListApplicationHistory
// API call, and error handling.
//
//	// Example sending a request using the ListApplicationHistoryRequest method.
//	req, resp := client.ListApplicationHistoryRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) ListApplicationHistoryRequest(input *ListApplicationHistoryInput) (req *request.Request, output *ListApplicationHistoryOutput) {
	op := &request.Operation{
		Name:       opListApplicationHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListApplicationHistoryInput{}
	}

	output = &ListApplicationHistoryOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListApplicationHistory API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation ListApplicationHistory for usage and error information.
func (c *SPARK) ListApplicationHistory(input *ListApplicationHistoryInput) (*ListApplicationHistoryOutput, error) {
	req, out := c.ListApplicationHistoryRequest(input)
	return out, req.Send()
}

// ListApplicationHistoryWithContext is the same as ListApplicationHistory with the addition of
// the ability to pass a context and additional request options.
//
// See ListApplicationHistory for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) ListApplicationHistoryWithContext(ctx volcengine.Context, input *ListApplicationHistoryInput, opts ...request.Option) (*ListApplicationHistoryOutput, error) {
	req, out := c.ListApplicationHistoryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataListForlistApplicationHistoryOutput struct {
	_ struct{} `type:"structure"`

	AppCreateTime *string `type:"string"`

	AppFinishTime *string `type:"string"`

	AppStartTime *string `type:"string"`

	ResourceMeterage *ResourceMeterageForlistApplicationHistoryOutput `type:"structure"`

	RetryCount *int32 `type:"int32"`

	State *string `type:"string"`
}

// String returns the string representation
func (s DataListForlistApplicationHistoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataListForlistApplicationHistoryOutput) GoString() string {
	return s.String()
}

// SetAppCreateTime sets the AppCreateTime field's value.
func (s *DataListForlistApplicationHistoryOutput) SetAppCreateTime(v string) *DataListForlistApplicationHistoryOutput {
	s.AppCreateTime = &v
	return s
}

// SetAppFinishTime sets the AppFinishTime field's value.
func (s *DataListForlistApplicationHistoryOutput) SetAppFinishTime(v string) *DataListForlistApplicationHistoryOutput {
	s.AppFinishTime = &v
	return s
}

// SetAppStartTime sets the AppStartTime field's value.
func (s *DataListForlistApplicationHistoryOutput) SetAppStartTime(v string) *DataListForlistApplicationHistoryOutput {
	s.AppStartTime = &v
	return s
}

// SetResourceMeterage sets the ResourceMeterage field's value.
func (s *DataListForlistApplicationHistoryOutput) SetResourceMeterage(v *ResourceMeterageForlistApplicationHistoryOutput) *DataListForlistApplicationHistoryOutput {
	s.ResourceMeterage = v
	return s
}

// SetRetryCount sets the RetryCount field's value.
func (s *DataListForlistApplicationHistoryOutput) SetRetryCount(v int32) *DataListForlistApplicationHistoryOutput {
	s.RetryCount = &v
	return s
}

// SetState sets the State field's value.
func (s *DataListForlistApplicationHistoryOutput) SetState(v string) *DataListForlistApplicationHistoryOutput {
	s.State = &v
	return s
}

type ListApplicationHistoryInput struct {
	_ struct{} `type:"structure"`

	DeploymentId *string `type:"string"`

	MaxResults *int32 `type:"int32"`

	NextToken *string `type:"string"`

	StartTimeLeft *string `type:"string"`

	StartTimeRight *string `type:"string"`
}

// String returns the string representation
func (s ListApplicationHistoryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListApplicationHistoryInput) GoString() string {
	return s.String()
}

// SetDeploymentId sets the DeploymentId field's value.
func (s *ListApplicationHistoryInput) SetDeploymentId(v string) *ListApplicationHistoryInput {
	s.DeploymentId = &v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListApplicationHistoryInput) SetMaxResults(v int32) *ListApplicationHistoryInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListApplicationHistoryInput) SetNextToken(v string) *ListApplicationHistoryInput {
	s.NextToken = &v
	return s
}

// SetStartTimeLeft sets the StartTimeLeft field's value.
func (s *ListApplicationHistoryInput) SetStartTimeLeft(v string) *ListApplicationHistoryInput {
	s.StartTimeLeft = &v
	return s
}

// SetStartTimeRight sets the StartTimeRight field's value.
func (s *ListApplicationHistoryInput) SetStartTimeRight(v string) *ListApplicationHistoryInput {
	s.StartTimeRight = &v
	return s
}

type ListApplicationHistoryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DataList []*DataListForlistApplicationHistoryOutput `type:"list"`

	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListApplicationHistoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListApplicationHistoryOutput) GoString() string {
	return s.String()
}

// SetDataList sets the DataList field's value.
func (s *ListApplicationHistoryOutput) SetDataList(v []*DataListForlistApplicationHistoryOutput) *ListApplicationHistoryOutput {
	s.DataList = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListApplicationHistoryOutput) SetNextToken(v string) *ListApplicationHistoryOutput {
	s.NextToken = &v
	return s
}

type ResourceMeterageForlistApplicationHistoryOutput struct {
	_ struct{} `type:"structure"`

	Cpu *string `type:"string" json:"cpu"`

	Memory *string `type:"string" json:"memory"`
}

// String returns the string representation
func (s ResourceMeterageForlistApplicationHistoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResourceMeterageForlistApplicationHistoryOutput) GoString() string {
	return s.String()
}

// SetCpu sets the Cpu field's value.
func (s *ResourceMeterageForlistApplicationHistoryOutput) SetCpu(v string) *ResourceMeterageForlistApplicationHistoryOutput {
	s.Cpu = &v
	return s
}

// SetMemory sets the Memory field's value.
func (s *ResourceMeterageForlistApplicationHistoryOutput) SetMemory(v string) *ResourceMeterageForlistApplicationHistoryOutput {
	s.Memory = &v
	return s
}
