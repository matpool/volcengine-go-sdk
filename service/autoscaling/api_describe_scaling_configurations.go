// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeScalingConfigurationsCommon = "DescribeScalingConfigurations"

// DescribeScalingConfigurationsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeScalingConfigurationsCommon operation. The "output" return
// value will be populated with the DescribeScalingConfigurationsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingConfigurationsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingConfigurationsCommon Send returns without error.
//
// See DescribeScalingConfigurationsCommon for more information on using the DescribeScalingConfigurationsCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeScalingConfigurationsCommonRequest method.
//	req, resp := client.DescribeScalingConfigurationsCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DescribeScalingConfigurationsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeScalingConfigurationsCommon,
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

// DescribeScalingConfigurationsCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DescribeScalingConfigurationsCommon for usage and error information.
func (c *AUTOSCALING) DescribeScalingConfigurationsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingConfigurationsCommonRequest(input)
	return out, req.Send()
}

// DescribeScalingConfigurationsCommonWithContext is the same as DescribeScalingConfigurationsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingConfigurationsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingConfigurationsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeScalingConfigurationsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeScalingConfigurations = "DescribeScalingConfigurations"

// DescribeScalingConfigurationsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeScalingConfigurations operation. The "output" return
// value will be populated with the DescribeScalingConfigurationsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeScalingConfigurationsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeScalingConfigurationsCommon Send returns without error.
//
// See DescribeScalingConfigurations for more information on using the DescribeScalingConfigurations
// API call, and error handling.
//
//	// Example sending a request using the DescribeScalingConfigurationsRequest method.
//	req, resp := client.DescribeScalingConfigurationsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DescribeScalingConfigurationsRequest(input *DescribeScalingConfigurationsInput) (req *request.Request, output *DescribeScalingConfigurationsOutput) {
	op := &request.Operation{
		Name:       opDescribeScalingConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeScalingConfigurationsInput{}
	}

	output = &DescribeScalingConfigurationsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeScalingConfigurations API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DescribeScalingConfigurations for usage and error information.
func (c *AUTOSCALING) DescribeScalingConfigurations(input *DescribeScalingConfigurationsInput) (*DescribeScalingConfigurationsOutput, error) {
	req, out := c.DescribeScalingConfigurationsRequest(input)
	return out, req.Send()
}

// DescribeScalingConfigurationsWithContext is the same as DescribeScalingConfigurations with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeScalingConfigurations for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DescribeScalingConfigurationsWithContext(ctx volcengine.Context, input *DescribeScalingConfigurationsInput, opts ...request.Option) (*DescribeScalingConfigurationsOutput, error) {
	req, out := c.DescribeScalingConfigurationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeScalingConfigurationsInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingConfigurationIds []*string `type:"list"`

	ScalingConfigurationNames []*string `type:"list"`

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s DescribeScalingConfigurationsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingConfigurationsInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingConfigurationsInput) SetPageNumber(v int32) *DescribeScalingConfigurationsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingConfigurationsInput) SetPageSize(v int32) *DescribeScalingConfigurationsInput {
	s.PageSize = &v
	return s
}

// SetScalingConfigurationIds sets the ScalingConfigurationIds field's value.
func (s *DescribeScalingConfigurationsInput) SetScalingConfigurationIds(v []*string) *DescribeScalingConfigurationsInput {
	s.ScalingConfigurationIds = v
	return s
}

// SetScalingConfigurationNames sets the ScalingConfigurationNames field's value.
func (s *DescribeScalingConfigurationsInput) SetScalingConfigurationNames(v []*string) *DescribeScalingConfigurationsInput {
	s.ScalingConfigurationNames = v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *DescribeScalingConfigurationsInput) SetScalingGroupId(v string) *DescribeScalingConfigurationsInput {
	s.ScalingGroupId = &v
	return s
}

type DescribeScalingConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ScalingConfigurations []*ScalingConfigurationForDescribeScalingConfigurationsOutput `type:"list"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeScalingConfigurationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeScalingConfigurationsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeScalingConfigurationsOutput) SetPageNumber(v int32) *DescribeScalingConfigurationsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeScalingConfigurationsOutput) SetPageSize(v int32) *DescribeScalingConfigurationsOutput {
	s.PageSize = &v
	return s
}

// SetScalingConfigurations sets the ScalingConfigurations field's value.
func (s *DescribeScalingConfigurationsOutput) SetScalingConfigurations(v []*ScalingConfigurationForDescribeScalingConfigurationsOutput) *DescribeScalingConfigurationsOutput {
	s.ScalingConfigurations = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeScalingConfigurationsOutput) SetTotalCount(v int32) *DescribeScalingConfigurationsOutput {
	s.TotalCount = &v
	return s
}

type EipForDescribeScalingConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int32 `type:"int32"`

	BillingType *string `type:"string"`

	ISP *string `type:"string"`
}

// String returns the string representation
func (s EipForDescribeScalingConfigurationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s EipForDescribeScalingConfigurationsOutput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *EipForDescribeScalingConfigurationsOutput) SetBandwidth(v int32) *EipForDescribeScalingConfigurationsOutput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *EipForDescribeScalingConfigurationsOutput) SetBillingType(v string) *EipForDescribeScalingConfigurationsOutput {
	s.BillingType = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *EipForDescribeScalingConfigurationsOutput) SetISP(v string) *EipForDescribeScalingConfigurationsOutput {
	s.ISP = &v
	return s
}

type ScalingConfigurationForDescribeScalingConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	CreatedAt *string `type:"string"`

	Eip *EipForDescribeScalingConfigurationsOutput `type:"structure"`

	HostName *string `type:"string"`

	ImageId *string `type:"string"`

	InstanceDescription *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceTypes []*string `type:"list"`

	KeyPairName *string `type:"string"`

	LifecycleState *string `type:"string"`

	ScalingConfigurationId *string `type:"string"`

	ScalingConfigurationName *string `type:"string"`

	ScalingGroupId *string `type:"string"`

	SecurityEnhancementStrategy *string `type:"string"`

	SecurityGroupIds []*string `type:"list"`

	UpdatedAt *string `type:"string"`

	UserData *string `type:"string"`

	Volumes []*VolumeForDescribeScalingConfigurationsOutput `type:"list"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s ScalingConfigurationForDescribeScalingConfigurationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ScalingConfigurationForDescribeScalingConfigurationsOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetCreatedAt(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.CreatedAt = &v
	return s
}

// SetEip sets the Eip field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetEip(v *EipForDescribeScalingConfigurationsOutput) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.Eip = v
	return s
}

// SetHostName sets the HostName field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetHostName(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.HostName = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetImageId(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.ImageId = &v
	return s
}

// SetInstanceDescription sets the InstanceDescription field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetInstanceDescription(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.InstanceDescription = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetInstanceName(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.InstanceName = &v
	return s
}

// SetInstanceTypes sets the InstanceTypes field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetInstanceTypes(v []*string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.InstanceTypes = v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetKeyPairName(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.KeyPairName = &v
	return s
}

// SetLifecycleState sets the LifecycleState field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetLifecycleState(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.LifecycleState = &v
	return s
}

// SetScalingConfigurationId sets the ScalingConfigurationId field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetScalingConfigurationId(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.ScalingConfigurationId = &v
	return s
}

// SetScalingConfigurationName sets the ScalingConfigurationName field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetScalingConfigurationName(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.ScalingConfigurationName = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetScalingGroupId(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.ScalingGroupId = &v
	return s
}

// SetSecurityEnhancementStrategy sets the SecurityEnhancementStrategy field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetSecurityEnhancementStrategy(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.SecurityEnhancementStrategy = &v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetSecurityGroupIds(v []*string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.SecurityGroupIds = v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetUpdatedAt(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.UpdatedAt = &v
	return s
}

// SetUserData sets the UserData field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetUserData(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.UserData = &v
	return s
}

// SetVolumes sets the Volumes field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetVolumes(v []*VolumeForDescribeScalingConfigurationsOutput) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.Volumes = v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *ScalingConfigurationForDescribeScalingConfigurationsOutput) SetZoneId(v string) *ScalingConfigurationForDescribeScalingConfigurationsOutput {
	s.ZoneId = &v
	return s
}

type VolumeForDescribeScalingConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	DeleteWithInstance *bool `type:"boolean"`

	Size *int32 `type:"int32"`

	VolumeType *string `type:"string"`
}

// String returns the string representation
func (s VolumeForDescribeScalingConfigurationsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VolumeForDescribeScalingConfigurationsOutput) GoString() string {
	return s.String()
}

// SetDeleteWithInstance sets the DeleteWithInstance field's value.
func (s *VolumeForDescribeScalingConfigurationsOutput) SetDeleteWithInstance(v bool) *VolumeForDescribeScalingConfigurationsOutput {
	s.DeleteWithInstance = &v
	return s
}

// SetSize sets the Size field's value.
func (s *VolumeForDescribeScalingConfigurationsOutput) SetSize(v int32) *VolumeForDescribeScalingConfigurationsOutput {
	s.Size = &v
	return s
}

// SetVolumeType sets the VolumeType field's value.
func (s *VolumeForDescribeScalingConfigurationsOutput) SetVolumeType(v string) *VolumeForDescribeScalingConfigurationsOutput {
	s.VolumeType = &v
	return s
}
