// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDetachServerGroupsCommon = "DetachServerGroups"

// DetachServerGroupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachServerGroupsCommon operation. The "output" return
// value will be populated with the DetachServerGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachServerGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachServerGroupsCommon Send returns without error.
//
// See DetachServerGroupsCommon for more information on using the DetachServerGroupsCommon
// API call, and error handling.
//
//	// Example sending a request using the DetachServerGroupsCommonRequest method.
//	req, resp := client.DetachServerGroupsCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DetachServerGroupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachServerGroupsCommon,
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

// DetachServerGroupsCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DetachServerGroupsCommon for usage and error information.
func (c *AUTOSCALING) DetachServerGroupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachServerGroupsCommonRequest(input)
	return out, req.Send()
}

// DetachServerGroupsCommonWithContext is the same as DetachServerGroupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachServerGroupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DetachServerGroupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachServerGroupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachServerGroups = "DetachServerGroups"

// DetachServerGroupsRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachServerGroups operation. The "output" return
// value will be populated with the DetachServerGroupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachServerGroupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachServerGroupsCommon Send returns without error.
//
// See DetachServerGroups for more information on using the DetachServerGroups
// API call, and error handling.
//
//	// Example sending a request using the DetachServerGroupsRequest method.
//	req, resp := client.DetachServerGroupsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DetachServerGroupsRequest(input *DetachServerGroupsInput) (req *request.Request, output *DetachServerGroupsOutput) {
	op := &request.Operation{
		Name:       opDetachServerGroups,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachServerGroupsInput{}
	}

	output = &DetachServerGroupsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachServerGroups API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DetachServerGroups for usage and error information.
func (c *AUTOSCALING) DetachServerGroups(input *DetachServerGroupsInput) (*DetachServerGroupsOutput, error) {
	req, out := c.DetachServerGroupsRequest(input)
	return out, req.Send()
}

// DetachServerGroupsWithContext is the same as DetachServerGroups with the addition of
// the ability to pass a context and additional request options.
//
// See DetachServerGroups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DetachServerGroupsWithContext(ctx volcengine.Context, input *DetachServerGroupsInput, opts ...request.Option) (*DetachServerGroupsOutput, error) {
	req, out := c.DetachServerGroupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachServerGroupsInput struct {
	_ struct{} `type:"structure"`

	// ScalingGroupId is a required field
	ScalingGroupId *string `type:"string" required:"true"`

	ServerGroupAttributes []*ServerGroupAttributeForDetachServerGroupsInput `type:"list"`
}

// String returns the string representation
func (s DetachServerGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachServerGroupsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DetachServerGroupsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DetachServerGroupsInput"}
	if s.ScalingGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("ScalingGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DetachServerGroupsInput) SetScalingGroupId(v string) *DetachServerGroupsInput {
	s.ScalingGroupId = &v
	return s
}

// SetServerGroupAttributes sets the ServerGroupAttributes field's value.
func (s *DetachServerGroupsInput) SetServerGroupAttributes(v []*ServerGroupAttributeForDetachServerGroupsInput) *DetachServerGroupsInput {
	s.ServerGroupAttributes = v
	return s
}

type DetachServerGroupsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s DetachServerGroupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachServerGroupsOutput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DetachServerGroupsOutput) SetScalingGroupId(v string) *DetachServerGroupsOutput {
	s.ScalingGroupId = &v
	return s
}

type ServerGroupAttributeForDetachServerGroupsInput struct {
	_ struct{} `type:"structure"`

	ServerGroupId *string `type:"string"`
}

// String returns the string representation
func (s ServerGroupAttributeForDetachServerGroupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerGroupAttributeForDetachServerGroupsInput) GoString() string {
	return s.String()
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ServerGroupAttributeForDetachServerGroupsInput) SetServerGroupId(v string) *ServerGroupAttributeForDetachServerGroupsInput {
	s.ServerGroupId = &v
	return s
}
