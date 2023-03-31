// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vke

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateNodesCommon = "CreateNodes"

// CreateNodesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNodesCommon operation. The "output" return
// value will be populated with the CreateNodesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNodesCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNodesCommon Send returns without error.
//
// See CreateNodesCommon for more information on using the CreateNodesCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateNodesCommonRequest method.
//    req, resp := client.CreateNodesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) CreateNodesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateNodesCommon,
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

// CreateNodesCommon API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation CreateNodesCommon for usage and error information.
func (c *VKE) CreateNodesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateNodesCommonRequest(input)
	return out, req.Send()
}

// CreateNodesCommonWithContext is the same as CreateNodesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNodesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) CreateNodesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateNodesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateNodes = "CreateNodes"

// CreateNodesRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateNodes operation. The "output" return
// value will be populated with the CreateNodesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateNodesCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateNodesCommon Send returns without error.
//
// See CreateNodes for more information on using the CreateNodes
// API call, and error handling.
//
//    // Example sending a request using the CreateNodesRequest method.
//    req, resp := client.CreateNodesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VKE) CreateNodesRequest(input *CreateNodesInput) (req *request.Request, output *CreateNodesOutput) {
	op := &request.Operation{
		Name:       opCreateNodes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateNodesInput{}
	}

	output = &CreateNodesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateNodes API operation for VKE.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VKE's
// API operation CreateNodes for usage and error information.
func (c *VKE) CreateNodes(input *CreateNodesInput) (*CreateNodesOutput, error) {
	req, out := c.CreateNodesRequest(input)
	return out, req.Send()
}

// CreateNodesWithContext is the same as CreateNodes with the addition of
// the ability to pass a context and additional request options.
//
// See CreateNodes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VKE) CreateNodesWithContext(ctx volcengine.Context, input *CreateNodesInput, opts ...request.Option) (*CreateNodesOutput, error) {
	req, out := c.CreateNodesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateNodesInput struct {
	_ struct{} `type:"structure"`

	AdditionalContainerStorageEnabled *bool `type:"boolean"`

	ClientToken *string `type:"string"`

	ClusterId *string `type:"string"`

	ContainerStoragePath *string `type:"string"`

	ImageId *string `type:"string"`

	InitializeScript *string `type:"string"`

	InstanceIds []*string `type:"list"`

	KeepInstanceName *bool `type:"boolean"`

	KubernetesConfig *KubernetesConfigForCreateNodesInput `type:"structure"`
}

// String returns the string representation
func (s CreateNodesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNodesInput) GoString() string {
	return s.String()
}

// SetAdditionalContainerStorageEnabled sets the AdditionalContainerStorageEnabled field's value.
func (s *CreateNodesInput) SetAdditionalContainerStorageEnabled(v bool) *CreateNodesInput {
	s.AdditionalContainerStorageEnabled = &v
	return s
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateNodesInput) SetClientToken(v string) *CreateNodesInput {
	s.ClientToken = &v
	return s
}

// SetClusterId sets the ClusterId field's value.
func (s *CreateNodesInput) SetClusterId(v string) *CreateNodesInput {
	s.ClusterId = &v
	return s
}

// SetContainerStoragePath sets the ContainerStoragePath field's value.
func (s *CreateNodesInput) SetContainerStoragePath(v string) *CreateNodesInput {
	s.ContainerStoragePath = &v
	return s
}

// SetImageId sets the ImageId field's value.
func (s *CreateNodesInput) SetImageId(v string) *CreateNodesInput {
	s.ImageId = &v
	return s
}

// SetInitializeScript sets the InitializeScript field's value.
func (s *CreateNodesInput) SetInitializeScript(v string) *CreateNodesInput {
	s.InitializeScript = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *CreateNodesInput) SetInstanceIds(v []*string) *CreateNodesInput {
	s.InstanceIds = v
	return s
}

// SetKeepInstanceName sets the KeepInstanceName field's value.
func (s *CreateNodesInput) SetKeepInstanceName(v bool) *CreateNodesInput {
	s.KeepInstanceName = &v
	return s
}

// SetKubernetesConfig sets the KubernetesConfig field's value.
func (s *CreateNodesInput) SetKubernetesConfig(v *KubernetesConfigForCreateNodesInput) *CreateNodesInput {
	s.KubernetesConfig = v
	return s
}

type CreateNodesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Ids []*string `type:"list"`
}

// String returns the string representation
func (s CreateNodesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateNodesOutput) GoString() string {
	return s.String()
}

// SetIds sets the Ids field's value.
func (s *CreateNodesOutput) SetIds(v []*string) *CreateNodesOutput {
	s.Ids = v
	return s
}

type KubernetesConfigForCreateNodesInput struct {
	_ struct{} `type:"structure"`

	Cordon *bool `type:"boolean"`

	Labels []*LabelForCreateNodesInput `type:"list"`

	Taints []*TaintForCreateNodesInput `type:"list"`
}

// String returns the string representation
func (s KubernetesConfigForCreateNodesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s KubernetesConfigForCreateNodesInput) GoString() string {
	return s.String()
}

// SetCordon sets the Cordon field's value.
func (s *KubernetesConfigForCreateNodesInput) SetCordon(v bool) *KubernetesConfigForCreateNodesInput {
	s.Cordon = &v
	return s
}

// SetLabels sets the Labels field's value.
func (s *KubernetesConfigForCreateNodesInput) SetLabels(v []*LabelForCreateNodesInput) *KubernetesConfigForCreateNodesInput {
	s.Labels = v
	return s
}

// SetTaints sets the Taints field's value.
func (s *KubernetesConfigForCreateNodesInput) SetTaints(v []*TaintForCreateNodesInput) *KubernetesConfigForCreateNodesInput {
	s.Taints = v
	return s
}

type LabelForCreateNodesInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s LabelForCreateNodesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s LabelForCreateNodesInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *LabelForCreateNodesInput) SetKey(v string) *LabelForCreateNodesInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *LabelForCreateNodesInput) SetValue(v string) *LabelForCreateNodesInput {
	s.Value = &v
	return s
}

type TaintForCreateNodesInput struct {
	_ struct{} `type:"structure"`

	Effect *string `type:"string" enum:"EnumOfEffectForCreateNodesInput"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TaintForCreateNodesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TaintForCreateNodesInput) GoString() string {
	return s.String()
}

// SetEffect sets the Effect field's value.
func (s *TaintForCreateNodesInput) SetEffect(v string) *TaintForCreateNodesInput {
	s.Effect = &v
	return s
}

// SetKey sets the Key field's value.
func (s *TaintForCreateNodesInput) SetKey(v string) *TaintForCreateNodesInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TaintForCreateNodesInput) SetValue(v string) *TaintForCreateNodesInput {
	s.Value = &v
	return s
}

const (
	// EnumOfEffectForCreateNodesInputNoExecute is a EnumOfEffectForCreateNodesInput enum value
	EnumOfEffectForCreateNodesInputNoExecute = "NoExecute"

	// EnumOfEffectForCreateNodesInputNoSchedule is a EnumOfEffectForCreateNodesInput enum value
	EnumOfEffectForCreateNodesInputNoSchedule = "NoSchedule"

	// EnumOfEffectForCreateNodesInputPreferNoSchedule is a EnumOfEffectForCreateNodesInput enum value
	EnumOfEffectForCreateNodesInputPreferNoSchedule = "PreferNoSchedule"
)
