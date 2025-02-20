// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDetachDBInstancesCommon = "DetachDBInstances"

// DetachDBInstancesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachDBInstancesCommon operation. The "output" return
// value will be populated with the DetachDBInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachDBInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachDBInstancesCommon Send returns without error.
//
// See DetachDBInstancesCommon for more information on using the DetachDBInstancesCommon
// API call, and error handling.
//
//	// Example sending a request using the DetachDBInstancesCommonRequest method.
//	req, resp := client.DetachDBInstancesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DetachDBInstancesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachDBInstancesCommon,
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

// DetachDBInstancesCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DetachDBInstancesCommon for usage and error information.
func (c *AUTOSCALING) DetachDBInstancesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachDBInstancesCommonRequest(input)
	return out, req.Send()
}

// DetachDBInstancesCommonWithContext is the same as DetachDBInstancesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachDBInstancesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DetachDBInstancesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachDBInstancesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachDBInstances = "DetachDBInstances"

// DetachDBInstancesRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachDBInstances operation. The "output" return
// value will be populated with the DetachDBInstancesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachDBInstancesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachDBInstancesCommon Send returns without error.
//
// See DetachDBInstances for more information on using the DetachDBInstances
// API call, and error handling.
//
//	// Example sending a request using the DetachDBInstancesRequest method.
//	req, resp := client.DetachDBInstancesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DetachDBInstancesRequest(input *DetachDBInstancesInput) (req *request.Request, output *DetachDBInstancesOutput) {
	op := &request.Operation{
		Name:       opDetachDBInstances,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachDBInstancesInput{}
	}

	output = &DetachDBInstancesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachDBInstances API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DetachDBInstances for usage and error information.
func (c *AUTOSCALING) DetachDBInstances(input *DetachDBInstancesInput) (*DetachDBInstancesOutput, error) {
	req, out := c.DetachDBInstancesRequest(input)
	return out, req.Send()
}

// DetachDBInstancesWithContext is the same as DetachDBInstances with the addition of
// the ability to pass a context and additional request options.
//
// See DetachDBInstances for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DetachDBInstancesWithContext(ctx volcengine.Context, input *DetachDBInstancesInput, opts ...request.Option) (*DetachDBInstancesOutput, error) {
	req, out := c.DetachDBInstancesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachDBInstancesInput struct {
	_ struct{} `type:"structure"`

	DBInstanceIds []*string `type:"list"`

	ForceDetach *bool `type:"boolean"`

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s DetachDBInstancesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachDBInstancesInput) GoString() string {
	return s.String()
}

// SetDBInstanceIds sets the DBInstanceIds field's value.
func (s *DetachDBInstancesInput) SetDBInstanceIds(v []*string) *DetachDBInstancesInput {
	s.DBInstanceIds = v
	return s
}

// SetForceDetach sets the ForceDetach field's value.
func (s *DetachDBInstancesInput) SetForceDetach(v bool) *DetachDBInstancesInput {
	s.ForceDetach = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DetachDBInstancesInput) SetScalingGroupId(v string) *DetachDBInstancesInput {
	s.ScalingGroupId = &v
	return s
}

type DetachDBInstancesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s DetachDBInstancesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachDBInstancesOutput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DetachDBInstancesOutput) SetScalingGroupId(v string) *DetachDBInstancesOutput {
	s.ScalingGroupId = &v
	return s
}
