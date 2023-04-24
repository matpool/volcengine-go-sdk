// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateNodePoolCommon = "CreateNodePool"

// CreateNodePoolCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNodePoolCommon operation. The "output" return
// value will be populated with the CreateNodePoolCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNodePoolCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNodePoolCommon Send returns without error.
//
// See CreateNodePoolCommon for more information on using the CreateNodePoolCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateNodePoolCommonRequest method.
//    req, resp := client.CreateNodePoolCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) CreateNodePoolCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateNodePoolCommon,
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

// CreateNodePoolCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation CreateNodePoolCommon for usage and error information.
func (c *VKE) CreateNodePoolCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateNodePoolCommonRequest(input)
	return out, req.Send()
}

// CreateNodePoolCommonWithContext is the same as CreateNodePoolCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNodePoolCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) CreateNodePoolCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateNodePoolCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateNodePool = "CreateNodePool"

// CreateNodePoolRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNodePool operation. The "output" return
// value will be populated with the CreateNodePoolCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNodePoolCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNodePoolCommon Send returns without error.
//
// See CreateNodePool for more information on using the CreateNodePool
// API call, and error handling.
//
//    // Example sending a request using the CreateNodePoolRequest method.
//    req, resp := client.CreateNodePoolRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) CreateNodePoolRequest(input *CreateNodePoolInput) (req *request.Request, output *CreateNodePoolOutput) {
	op := &request.Operation{
		Name:       opCreateNodePool,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNodePoolInput{}
	}

	output = &CreateNodePoolOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateNodePool API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation CreateNodePool for usage and error information.
func (c *VKE) CreateNodePool(input *CreateNodePoolInput) (*CreateNodePoolOutput, error) {
	req, out := c.CreateNodePoolRequest(input)
	return out, req.Send()
}

// CreateNodePoolWithContext is the same as CreateNodePool with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNodePool for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) CreateNodePoolWithContext(ctx volcengine.Context, input *CreateNodePoolInput, opts ...request.Option) (*CreateNodePoolOutput, error) {
	req, out := c.CreateNodePoolRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AutoScalingForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	DesiredReplicas *int32 `type:"int32"`

	Enabled *bool `type:"boolean"`

	MaxReplicas *int32 `type:"int32"`

	MinReplicas *int32 `type:"int32"`

	Priority *int32 `type:"int32"`

	SubnetPolicy *string `type:"string" enum:"EnumOfSubnetPolicyForCreateNodePoolInput"`
}

// String returns the string representation
func (s AutoScalingForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AutoScalingForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetDesiredReplicas sets the DesiredReplicas field's value.
func (s *AutoScalingForCreateNodePoolInput) SetDesiredReplicas(v int32) *AutoScalingForCreateNodePoolInput {
	s.DesiredReplicas = &v
	return s
}

// SetEnabled sets the Enabled field's value.
func (s *AutoScalingForCreateNodePoolInput) SetEnabled(v bool) *AutoScalingForCreateNodePoolInput {
	s.Enabled = &v
	return s
}

// SetMaxReplicas sets the MaxReplicas field's value.
func (s *AutoScalingForCreateNodePoolInput) SetMaxReplicas(v int32) *AutoScalingForCreateNodePoolInput {
	s.MaxReplicas = &v
	return s
}

// SetMinReplicas sets the MinReplicas field's value.
func (s *AutoScalingForCreateNodePoolInput) SetMinReplicas(v int32) *AutoScalingForCreateNodePoolInput {
	s.MinReplicas = &v
	return s
}

// SetPriority sets the Priority field's value.
func (s *AutoScalingForCreateNodePoolInput) SetPriority(v int32) *AutoScalingForCreateNodePoolInput {
	s.Priority = &v
	return s
}

// SetSubnetPolicy sets the SubnetPolicy field's value.
func (s *AutoScalingForCreateNodePoolInput) SetSubnetPolicy(v string) *AutoScalingForCreateNodePoolInput {
	s.SubnetPolicy = &v
	return s
}

type CreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	AutoScaling *AutoScalingForCreateNodePoolInput `type:"structure"`

	ClientToken *string `type:"string"`

	ClusterId *string `type:"string"`

	KubernetesConfig *KubernetesConfigForCreateNodePoolInput `type:"structure"`

	Name *string `type:"string"`

	NodeConfig *NodeConfigForCreateNodePoolInput `type:"structure"`

	Tags []*TagForCreateNodePoolInput `type:"list"`
}

// String returns the string representation
func (s CreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNodePoolInput) GoString() string {
	return s.String()
}

// SetAutoScaling sets the AutoScaling field's value.
func (s *CreateNodePoolInput) SetAutoScaling(v *AutoScalingForCreateNodePoolInput) *CreateNodePoolInput {
	s.AutoScaling = v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateNodePoolInput) SetClientToken(v string) *CreateNodePoolInput {
	s.ClientToken = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *CreateNodePoolInput) SetClusterId(v string) *CreateNodePoolInput {
	s.ClusterId = &v
	return s
}

// SetKubernetesConfig sets the KubernetesConfig field's value.
func (s *CreateNodePoolInput) SetKubernetesConfig(v *KubernetesConfigForCreateNodePoolInput) *CreateNodePoolInput {
	s.KubernetesConfig = v
	return s
}

// SetName sets the Name field's value.
func (s *CreateNodePoolInput) SetName(v string) *CreateNodePoolInput {
	s.Name = &v
	return s
}

// SetNodeConfig sets the NodeConfig field's value.
func (s *CreateNodePoolInput) SetNodeConfig(v *NodeConfigForCreateNodePoolInput) *CreateNodePoolInput {
	s.NodeConfig = v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateNodePoolInput) SetTags(v []*TagForCreateNodePoolInput) *CreateNodePoolInput {
	s.Tags = v
	return s
}

type CreateNodePoolOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string"`
}

// String returns the string representation
func (s CreateNodePoolOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNodePoolOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *CreateNodePoolOutput) SetId(v string) *CreateNodePoolOutput {
	s.Id = &v
	return s
}

type DataVolumeForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	MountPoint *string `type:"string"`

	Size *int32 `type:"int32"`

	Type *string `type:"string" enum:"EnumOfTypeForCreateNodePoolInput"`
}

// String returns the string representation
func (s DataVolumeForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataVolumeForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetMountPoint sets the MountPoint field's value.
func (s *DataVolumeForCreateNodePoolInput) SetMountPoint(v string) *DataVolumeForCreateNodePoolInput {
	s.MountPoint = &v
	return s
}

// SetSize sets the Size field's value.
func (s *DataVolumeForCreateNodePoolInput) SetSize(v int32) *DataVolumeForCreateNodePoolInput {
	s.Size = &v
	return s
}

// SetType sets the Type field's value.
func (s *DataVolumeForCreateNodePoolInput) SetType(v string) *DataVolumeForCreateNodePoolInput {
	s.Type = &v
	return s
}

type KubernetesConfigForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	Cordon *bool `type:"boolean"`

	Labels []*LabelForCreateNodePoolInput `type:"list"`

	Taints []*TaintForCreateNodePoolInput `type:"list"`
}

// String returns the string representation
func (s KubernetesConfigForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s KubernetesConfigForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetCordon sets the Cordon field's value.
func (s *KubernetesConfigForCreateNodePoolInput) SetCordon(v bool) *KubernetesConfigForCreateNodePoolInput {
	s.Cordon = &v
	return s
}

// SetLabels sets the Labels field's value.
func (s *KubernetesConfigForCreateNodePoolInput) SetLabels(v []*LabelForCreateNodePoolInput) *KubernetesConfigForCreateNodePoolInput {
	s.Labels = v
	return s
}

// SetTaints sets the Taints field's value.
func (s *KubernetesConfigForCreateNodePoolInput) SetTaints(v []*TaintForCreateNodePoolInput) *KubernetesConfigForCreateNodePoolInput {
	s.Taints = v
	return s
}

type LabelForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s LabelForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LabelForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *LabelForCreateNodePoolInput) SetKey(v string) *LabelForCreateNodePoolInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *LabelForCreateNodePoolInput) SetValue(v string) *LabelForCreateNodePoolInput {
	s.Value = &v
	return s
}

type LoginForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	Password *string `type:"string"`

	SshKeyPairName *string `type:"string"`
}

// String returns the string representation
func (s LoginForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LoginForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetPassword sets the Password field's value.
func (s *LoginForCreateNodePoolInput) SetPassword(v string) *LoginForCreateNodePoolInput {
	s.Password = &v
	return s
}

// SetSshKeyPairName sets the SshKeyPairName field's value.
func (s *LoginForCreateNodePoolInput) SetSshKeyPairName(v string) *LoginForCreateNodePoolInput {
	s.SshKeyPairName = &v
	return s
}

type NodeConfigForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	AdditionalContainerStorageEnabled *bool `type:"boolean"`

	AutoRenew *bool `type:"boolean"`

	AutoRenewPeriod *int32 `type:"int32"`

	DataVolumes []*DataVolumeForCreateNodePoolInput `type:"list"`

	HpcClusterIds []*string `type:"list"`

	ImageId *string `type:"string"`

	InitializeScript *string `type:"string"`

	InstanceChargeType *string `type:"string" enum:"EnumOfInstanceChargeTypeForCreateNodePoolInput"`

	InstanceTypeIds []*string `type:"list"`

	NamePrefix *string `type:"string"`

	Period *int32 `type:"int32"`

	Security *SecurityForCreateNodePoolInput `type:"structure"`

	SubnetIds []*string `type:"list"`

	SystemVolume *SystemVolumeForCreateNodePoolInput `type:"structure"`

	Tags []*TagForCreateNodePoolInput `type:"list"`
}

// String returns the string representation
func (s NodeConfigForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeConfigForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetAdditionalContainerStorageEnabled sets the AdditionalContainerStorageEnabled field's value.
func (s *NodeConfigForCreateNodePoolInput) SetAdditionalContainerStorageEnabled(v bool) *NodeConfigForCreateNodePoolInput {
	s.AdditionalContainerStorageEnabled = &v
	return s
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *NodeConfigForCreateNodePoolInput) SetAutoRenew(v bool) *NodeConfigForCreateNodePoolInput {
	s.AutoRenew = &v
	return s
}

// SetAutoRenewPeriod sets the AutoRenewPeriod field's value.
func (s *NodeConfigForCreateNodePoolInput) SetAutoRenewPeriod(v int32) *NodeConfigForCreateNodePoolInput {
	s.AutoRenewPeriod = &v
	return s
}

// SetDataVolumes sets the DataVolumes field's value.
func (s *NodeConfigForCreateNodePoolInput) SetDataVolumes(v []*DataVolumeForCreateNodePoolInput) *NodeConfigForCreateNodePoolInput {
	s.DataVolumes = v
	return s
}

// SetHpcClusterIds sets the HpcClusterIds field's value.
func (s *NodeConfigForCreateNodePoolInput) SetHpcClusterIds(v []*string) *NodeConfigForCreateNodePoolInput {
	s.HpcClusterIds = v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *NodeConfigForCreateNodePoolInput) SetImageId(v string) *NodeConfigForCreateNodePoolInput {
	s.ImageId = &v
	return s
}

// SetInitializeScript sets the InitializeScript field's value.
func (s *NodeConfigForCreateNodePoolInput) SetInitializeScript(v string) *NodeConfigForCreateNodePoolInput {
	s.InitializeScript = &v
	return s
}

// SetInstanceChargeType sets the InstanceChargeType field's value.
func (s *NodeConfigForCreateNodePoolInput) SetInstanceChargeType(v string) *NodeConfigForCreateNodePoolInput {
	s.InstanceChargeType = &v
	return s
}

// SetInstanceTypeIds sets the InstanceTypeIds field's value.
func (s *NodeConfigForCreateNodePoolInput) SetInstanceTypeIds(v []*string) *NodeConfigForCreateNodePoolInput {
	s.InstanceTypeIds = v
	return s
}

// SetNamePrefix sets the NamePrefix field's value.
func (s *NodeConfigForCreateNodePoolInput) SetNamePrefix(v string) *NodeConfigForCreateNodePoolInput {
	s.NamePrefix = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *NodeConfigForCreateNodePoolInput) SetPeriod(v int32) *NodeConfigForCreateNodePoolInput {
	s.Period = &v
	return s
}

// SetSecurity sets the Security field's value.
func (s *NodeConfigForCreateNodePoolInput) SetSecurity(v *SecurityForCreateNodePoolInput) *NodeConfigForCreateNodePoolInput {
	s.Security = v
	return s
}

// SetSubnetIds sets the SubnetIds field's value.
func (s *NodeConfigForCreateNodePoolInput) SetSubnetIds(v []*string) *NodeConfigForCreateNodePoolInput {
	s.SubnetIds = v
	return s
}

// SetSystemVolume sets the SystemVolume field's value.
func (s *NodeConfigForCreateNodePoolInput) SetSystemVolume(v *SystemVolumeForCreateNodePoolInput) *NodeConfigForCreateNodePoolInput {
	s.SystemVolume = v
	return s
}

// SetTags sets the Tags field's value.
func (s *NodeConfigForCreateNodePoolInput) SetTags(v []*TagForCreateNodePoolInput) *NodeConfigForCreateNodePoolInput {
	s.Tags = v
	return s
}

type SecurityForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	Login *LoginForCreateNodePoolInput `type:"structure"`

	SecurityGroupIds []*string `type:"list"`

	SecurityStrategies []*string `type:"list"`
}

// String returns the string representation
func (s SecurityForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SecurityForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetLogin sets the Login field's value.
func (s *SecurityForCreateNodePoolInput) SetLogin(v *LoginForCreateNodePoolInput) *SecurityForCreateNodePoolInput {
	s.Login = v
	return s
}

// SetSecurityGroupIds sets the SecurityGroupIds field's value.
func (s *SecurityForCreateNodePoolInput) SetSecurityGroupIds(v []*string) *SecurityForCreateNodePoolInput {
	s.SecurityGroupIds = v
	return s
}

// SetSecurityStrategies sets the SecurityStrategies field's value.
func (s *SecurityForCreateNodePoolInput) SetSecurityStrategies(v []*string) *SecurityForCreateNodePoolInput {
	s.SecurityStrategies = v
	return s
}

type SystemVolumeForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	Size *int32 `type:"int32"`

	Type *string `type:"string" enum:"EnumOfTypeForCreateNodePoolInput"`
}

// String returns the string representation
func (s SystemVolumeForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SystemVolumeForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetSize sets the Size field's value.
func (s *SystemVolumeForCreateNodePoolInput) SetSize(v int32) *SystemVolumeForCreateNodePoolInput {
	s.Size = &v
	return s
}

// SetType sets the Type field's value.
func (s *SystemVolumeForCreateNodePoolInput) SetType(v string) *SystemVolumeForCreateNodePoolInput {
	s.Type = &v
	return s
}

type TagForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateNodePoolInput) SetKey(v string) *TagForCreateNodePoolInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateNodePoolInput) SetValue(v string) *TagForCreateNodePoolInput {
	s.Value = &v
	return s
}

type TaintForCreateNodePoolInput struct {
	_ struct{} `type:"structure"`

	Effect *string `type:"string" enum:"EnumOfEffectForCreateNodePoolInput"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TaintForCreateNodePoolInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TaintForCreateNodePoolInput) GoString() string {
	return s.String()
}

// SetEffect sets the Effect field's value.
func (s *TaintForCreateNodePoolInput) SetEffect(v string) *TaintForCreateNodePoolInput {
	s.Effect = &v
	return s
}

// SetKey sets the Key field's value.
func (s *TaintForCreateNodePoolInput) SetKey(v string) *TaintForCreateNodePoolInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TaintForCreateNodePoolInput) SetValue(v string) *TaintForCreateNodePoolInput {
	s.Value = &v
	return s
}

const (
	// EnumOfEffectForCreateNodePoolInputNoExecute is a EnumOfEffectForCreateNodePoolInput enum value
	EnumOfEffectForCreateNodePoolInputNoExecute = "NoExecute"

	// EnumOfEffectForCreateNodePoolInputNoSchedule is a EnumOfEffectForCreateNodePoolInput enum value
	EnumOfEffectForCreateNodePoolInputNoSchedule = "NoSchedule"

	// EnumOfEffectForCreateNodePoolInputPreferNoSchedule is a EnumOfEffectForCreateNodePoolInput enum value
	EnumOfEffectForCreateNodePoolInputPreferNoSchedule = "PreferNoSchedule"
)

const (
	// EnumOfInstanceChargeTypeForCreateNodePoolInputPostPaid is a EnumOfInstanceChargeTypeForCreateNodePoolInput enum value
	EnumOfInstanceChargeTypeForCreateNodePoolInputPostPaid = "PostPaid"

	// EnumOfInstanceChargeTypeForCreateNodePoolInputPrePaid is a EnumOfInstanceChargeTypeForCreateNodePoolInput enum value
	EnumOfInstanceChargeTypeForCreateNodePoolInputPrePaid = "PrePaid"
)

const (
	// EnumOfSecurityStrategyListForCreateNodePoolInputHids is a EnumOfSecurityStrategyListForCreateNodePoolInput enum value
	EnumOfSecurityStrategyListForCreateNodePoolInputHids = "Hids"
)

const (
	// EnumOfSubnetPolicyForCreateNodePoolInputPriority is a EnumOfSubnetPolicyForCreateNodePoolInput enum value
	EnumOfSubnetPolicyForCreateNodePoolInputPriority = "Priority"

	// EnumOfSubnetPolicyForCreateNodePoolInputZoneBalance is a EnumOfSubnetPolicyForCreateNodePoolInput enum value
	EnumOfSubnetPolicyForCreateNodePoolInputZoneBalance = "ZoneBalance"
)

const (
	// EnumOfTypeForCreateNodePoolInputEssd is a EnumOfTypeForCreateNodePoolInput enum value
	EnumOfTypeForCreateNodePoolInputEssd = "ESSD"

	// EnumOfTypeForCreateNodePoolInputEssdFlexPl is a EnumOfTypeForCreateNodePoolInput enum value
	EnumOfTypeForCreateNodePoolInputEssdFlexPl = "ESSD_FlexPL"

	// EnumOfTypeForCreateNodePoolInputEssdPl0 is a EnumOfTypeForCreateNodePoolInput enum value
	EnumOfTypeForCreateNodePoolInputEssdPl0 = "ESSD_PL0"

	// EnumOfTypeForCreateNodePoolInputEssdPl1 is a EnumOfTypeForCreateNodePoolInput enum value
	EnumOfTypeForCreateNodePoolInputEssdPl1 = "ESSD_PL1"

	// EnumOfTypeForCreateNodePoolInputPtssd is a EnumOfTypeForCreateNodePoolInput enum value
	EnumOfTypeForCreateNodePoolInputPtssd = "PTSSD"
)
