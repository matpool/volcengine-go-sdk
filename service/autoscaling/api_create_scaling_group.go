// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateScalingGroupCommon = "CreateScalingGroup"

// CreateScalingGroupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateScalingGroupCommon operation. The "output" return
// value will be populated with the CreateScalingGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateScalingGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateScalingGroupCommon Send returns without error.
//
// See CreateScalingGroupCommon for more information on using the CreateScalingGroupCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateScalingGroupCommonRequest method.
//	req, resp := client.CreateScalingGroupCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) CreateScalingGroupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateScalingGroupCommon,
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

// CreateScalingGroupCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation CreateScalingGroupCommon for usage and error information.
func (c *AUTOSCALING) CreateScalingGroupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateScalingGroupCommonRequest(input)
	return out, req.Send()
}

// CreateScalingGroupCommonWithContext is the same as CreateScalingGroupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateScalingGroupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) CreateScalingGroupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateScalingGroupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateScalingGroup = "CreateScalingGroup"

// CreateScalingGroupRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateScalingGroup operation. The "output" return
// value will be populated with the CreateScalingGroupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateScalingGroupCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateScalingGroupCommon Send returns without error.
//
// See CreateScalingGroup for more information on using the CreateScalingGroup
// API call, and error handling.
//
//	// Example sending a request using the CreateScalingGroupRequest method.
//	req, resp := client.CreateScalingGroupRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) CreateScalingGroupRequest(input *CreateScalingGroupInput) (req *request.Request, output *CreateScalingGroupOutput) {
	op := &request.Operation{
		Name:       opCreateScalingGroup,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateScalingGroupInput{}
	}

	output = &CreateScalingGroupOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateScalingGroup API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation CreateScalingGroup for usage and error information.
func (c *AUTOSCALING) CreateScalingGroup(input *CreateScalingGroupInput) (*CreateScalingGroupOutput, error) {
	req, out := c.CreateScalingGroupRequest(input)
	return out, req.Send()
}

// CreateScalingGroupWithContext is the same as CreateScalingGroup with the addition of
// the ability to pass a context and additional request options.
//
// See CreateScalingGroup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) CreateScalingGroupWithContext(ctx volcengine.Context, input *CreateScalingGroupInput, opts ...request.Option) (*CreateScalingGroupOutput, error) {
	req, out := c.CreateScalingGroupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateScalingGroupInput struct {
	_ struct{} `type:"structure"`

	DBInstanceIds []*string `type:"list"`

	DefaultCooldown *int32 `type:"int32"`

	DesireInstanceNumber *int32 `type:"int32"`

	InstanceTerminatePolicy *string `type:"string"`

	MaxInstanceNumber *int32 `type:"int32"`

	MinInstanceNumber *int32 `type:"int32"`

	MultiAZPolicy *string `type:"string"`

	// ScalingGroupName is a required field
	ScalingGroupName *string `type:"string" required:"true"`

	ServerGroupAttributes []*ServerGroupAttributeForCreateScalingGroupInput `type:"list"`

	// SubnetIds is a required field
	SubnetIds []*string `type:"list" required:"true"`
}

// String returns the string representation
func (s CreateScalingGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateScalingGroupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateScalingGroupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateScalingGroupInput"}
	if s.ScalingGroupName == nil {
		invalidParams.Add(request.NewErrParamRequired("ScalingGroupName"))
	}
	if s.SubnetIds == nil {
		invalidParams.Add(request.NewErrParamRequired("SubnetIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDBInstanceIds sets the DBInstanceIds field's value.
func (s *CreateScalingGroupInput) SetDBInstanceIds(v []*string) *CreateScalingGroupInput {
	s.DBInstanceIds = v
	return s
}

// SetDefaultCooldown sets the DefaultCooldown field's value.
func (s *CreateScalingGroupInput) SetDefaultCooldown(v int32) *CreateScalingGroupInput {
	s.DefaultCooldown = &v
	return s
}

// SetDesireInstanceNumber sets the DesireInstanceNumber field's value.
func (s *CreateScalingGroupInput) SetDesireInstanceNumber(v int32) *CreateScalingGroupInput {
	s.DesireInstanceNumber = &v
	return s
}

// SetInstanceTerminatePolicy sets the InstanceTerminatePolicy field's value.
func (s *CreateScalingGroupInput) SetInstanceTerminatePolicy(v string) *CreateScalingGroupInput {
	s.InstanceTerminatePolicy = &v
	return s
}

// SetMaxInstanceNumber sets the MaxInstanceNumber field's value.
func (s *CreateScalingGroupInput) SetMaxInstanceNumber(v int32) *CreateScalingGroupInput {
	s.MaxInstanceNumber = &v
	return s
}

// SetMinInstanceNumber sets the MinInstanceNumber field's value.
func (s *CreateScalingGroupInput) SetMinInstanceNumber(v int32) *CreateScalingGroupInput {
	s.MinInstanceNumber = &v
	return s
}

// SetMultiAZPolicy sets the MultiAZPolicy field's value.
func (s *CreateScalingGroupInput) SetMultiAZPolicy(v string) *CreateScalingGroupInput {
	s.MultiAZPolicy = &v
	return s
}

// SetScalingGroupName sets the ScalingGroupName field's value.
func (s *CreateScalingGroupInput) SetScalingGroupName(v string) *CreateScalingGroupInput {
	s.ScalingGroupName = &v
	return s
}

// SetServerGroupAttributes sets the ServerGroupAttributes field's value.
func (s *CreateScalingGroupInput) SetServerGroupAttributes(v []*ServerGroupAttributeForCreateScalingGroupInput) *CreateScalingGroupInput {
	s.ServerGroupAttributes = v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *CreateScalingGroupInput) SetSubnetIds(v []*string) *CreateScalingGroupInput {
	s.SubnetIds = v
	return s
}

type CreateScalingGroupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s CreateScalingGroupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateScalingGroupOutput) GoString() string {
	return s.String()
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *CreateScalingGroupOutput) SetScalingGroupId(v string) *CreateScalingGroupOutput {
	s.ScalingGroupId = &v
	return s
}

type ServerGroupAttributeForCreateScalingGroupInput struct {
	_ struct{} `type:"structure"`

	Port *int32 `type:"int32"`

	ServerGroupId *string `type:"string"`

	Weight *int32 `type:"int32"`
}

// String returns the string representation
func (s ServerGroupAttributeForCreateScalingGroupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerGroupAttributeForCreateScalingGroupInput) GoString() string {
	return s.String()
}

// SetPort sets the Port field's value.
func (s *ServerGroupAttributeForCreateScalingGroupInput) SetPort(v int32) *ServerGroupAttributeForCreateScalingGroupInput {
	s.Port = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ServerGroupAttributeForCreateScalingGroupInput) SetServerGroupId(v string) *ServerGroupAttributeForCreateScalingGroupInput {
	s.ServerGroupId = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ServerGroupAttributeForCreateScalingGroupInput) SetWeight(v int32) *ServerGroupAttributeForCreateScalingGroupInput {
	s.Weight = &v
	return s
}
