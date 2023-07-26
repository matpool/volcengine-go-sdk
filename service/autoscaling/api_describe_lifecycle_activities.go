// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeLifecycleActivitiesCommon = "DescribeLifecycleActivities"

// DescribeLifecycleActivitiesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLifecycleActivitiesCommon operation. The "output" return
// value will be populated with the DescribeLifecycleActivitiesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLifecycleActivitiesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLifecycleActivitiesCommon Send returns without error.
//
// See DescribeLifecycleActivitiesCommon for more information on using the DescribeLifecycleActivitiesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeLifecycleActivitiesCommonRequest method.
//	req, resp := client.DescribeLifecycleActivitiesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DescribeLifecycleActivitiesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeLifecycleActivitiesCommon,
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

// DescribeLifecycleActivitiesCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DescribeLifecycleActivitiesCommon for usage and error information.
func (c *AUTOSCALING) DescribeLifecycleActivitiesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeLifecycleActivitiesCommonRequest(input)
	return out, req.Send()
}

// DescribeLifecycleActivitiesCommonWithContext is the same as DescribeLifecycleActivitiesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLifecycleActivitiesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeLifecycleActivitiesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeLifecycleActivitiesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeLifecycleActivities = "DescribeLifecycleActivities"

// DescribeLifecycleActivitiesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeLifecycleActivities operation. The "output" return
// value will be populated with the DescribeLifecycleActivitiesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeLifecycleActivitiesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeLifecycleActivitiesCommon Send returns without error.
//
// See DescribeLifecycleActivities for more information on using the DescribeLifecycleActivities
// API call, and error handling.
//
//	// Example sending a request using the DescribeLifecycleActivitiesRequest method.
//	req, resp := client.DescribeLifecycleActivitiesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DescribeLifecycleActivitiesRequest(input *DescribeLifecycleActivitiesInput) (req *request.Request, output *DescribeLifecycleActivitiesOutput) {
	op := &request.Operation{
		Name:       opDescribeLifecycleActivities,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeLifecycleActivitiesInput{}
	}

	output = &DescribeLifecycleActivitiesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeLifecycleActivities API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DescribeLifecycleActivities for usage and error information.
func (c *AUTOSCALING) DescribeLifecycleActivities(input *DescribeLifecycleActivitiesInput) (*DescribeLifecycleActivitiesOutput, error) {
	req, out := c.DescribeLifecycleActivitiesRequest(input)
	return out, req.Send()
}

// DescribeLifecycleActivitiesWithContext is the same as DescribeLifecycleActivities with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeLifecycleActivities for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeLifecycleActivitiesWithContext(ctx volcengine.Context, input *DescribeLifecycleActivitiesInput, opts ...request.Option) (*DescribeLifecycleActivitiesOutput, error) {
	req, out := c.DescribeLifecycleActivitiesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeLifecycleActivitiesInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	LifecycleActivityStatus *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingActivityId *string `type:"string"`
}

// String returns the string representation
func (s DescribeLifecycleActivitiesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLifecycleActivitiesInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeLifecycleActivitiesInput) SetInstanceId(v string) *DescribeLifecycleActivitiesInput {
	s.InstanceId = &v
	return s
}

// SetLifecycleActivityStatus sets the LifecycleActivityStatus field's value.
func (s *DescribeLifecycleActivitiesInput) SetLifecycleActivityStatus(v string) *DescribeLifecycleActivitiesInput {
	s.LifecycleActivityStatus = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLifecycleActivitiesInput) SetPageNumber(v int32) *DescribeLifecycleActivitiesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLifecycleActivitiesInput) SetPageSize(v int32) *DescribeLifecycleActivitiesInput {
	s.PageSize = &v
	return s
}

// SetScalingActivityId sets the ScalingActivityId field's value.
func (s *DescribeLifecycleActivitiesInput) SetScalingActivityId(v string) *DescribeLifecycleActivitiesInput {
	s.ScalingActivityId = &v
	return s
}

type DescribeLifecycleActivitiesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LifecycleActivities []*LifecycleActivityForDescribeLifecycleActivitiesOutput `type:"list"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeLifecycleActivitiesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeLifecycleActivitiesOutput) GoString() string {
	return s.String()
}

// SetLifecycleActivities sets the LifecycleActivities field's value.
func (s *DescribeLifecycleActivitiesOutput) SetLifecycleActivities(v []*LifecycleActivityForDescribeLifecycleActivitiesOutput) *DescribeLifecycleActivitiesOutput {
	s.LifecycleActivities = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeLifecycleActivitiesOutput) SetPageNumber(v int32) *DescribeLifecycleActivitiesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeLifecycleActivitiesOutput) SetPageSize(v int32) *DescribeLifecycleActivitiesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeLifecycleActivitiesOutput) SetTotalCount(v int32) *DescribeLifecycleActivitiesOutput {
	s.TotalCount = &v
	return s
}

type LifecycleActivityForDescribeLifecycleActivitiesOutput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`

	LifecycleActivityId *string `type:"string"`

	LifecycleActivityStatus *string `type:"string"`

	LifecycleHookId *string `type:"string"`

	LifecycleHookPolicy *string `type:"string"`

	ScalingActivityId *string `type:"string"`
}

// String returns the string representation
func (s LifecycleActivityForDescribeLifecycleActivitiesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LifecycleActivityForDescribeLifecycleActivitiesOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *LifecycleActivityForDescribeLifecycleActivitiesOutput) SetInstanceId(v string) *LifecycleActivityForDescribeLifecycleActivitiesOutput {
	s.InstanceId = &v
	return s
}

// SetLifecycleActivityId sets the LifecycleActivityId field's value.
func (s *LifecycleActivityForDescribeLifecycleActivitiesOutput) SetLifecycleActivityId(v string) *LifecycleActivityForDescribeLifecycleActivitiesOutput {
	s.LifecycleActivityId = &v
	return s
}

// SetLifecycleActivityStatus sets the LifecycleActivityStatus field's value.
func (s *LifecycleActivityForDescribeLifecycleActivitiesOutput) SetLifecycleActivityStatus(v string) *LifecycleActivityForDescribeLifecycleActivitiesOutput {
	s.LifecycleActivityStatus = &v
	return s
}

// SetLifecycleHookId sets the LifecycleHookId field's value.
func (s *LifecycleActivityForDescribeLifecycleActivitiesOutput) SetLifecycleHookId(v string) *LifecycleActivityForDescribeLifecycleActivitiesOutput {
	s.LifecycleHookId = &v
	return s
}

// SetLifecycleHookPolicy sets the LifecycleHookPolicy field's value.
func (s *LifecycleActivityForDescribeLifecycleActivitiesOutput) SetLifecycleHookPolicy(v string) *LifecycleActivityForDescribeLifecycleActivitiesOutput {
	s.LifecycleHookPolicy = &v
	return s
}

// SetScalingActivityId sets the ScalingActivityId field's value.
func (s *LifecycleActivityForDescribeLifecycleActivitiesOutput) SetScalingActivityId(v string) *LifecycleActivityForDescribeLifecycleActivitiesOutput {
	s.ScalingActivityId = &v
	return s
}
