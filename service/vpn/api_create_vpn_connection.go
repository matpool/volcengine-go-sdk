// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateVpnConnectionCommon = "CreateVpnConnection"

// CreateVpnConnectionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpnConnectionCommon operation. The "output" return
// value will be populated with the CreateVpnConnectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpnConnectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpnConnectionCommon Send returns without error.
//
// See CreateVpnConnectionCommon for more information on using the CreateVpnConnectionCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateVpnConnectionCommonRequest method.
//    req, resp := client.CreateVpnConnectionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) CreateVpnConnectionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateVpnConnectionCommon,
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

// CreateVpnConnectionCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPN's
// API operation CreateVpnConnectionCommon for usage and error information.
func (c *VPN) CreateVpnConnectionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateVpnConnectionCommonRequest(input)
	return out, req.Send()
}

// CreateVpnConnectionCommonWithContext is the same as CreateVpnConnectionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpnConnectionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) CreateVpnConnectionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateVpnConnectionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateVpnConnection = "CreateVpnConnection"

// CreateVpnConnectionRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateVpnConnection operation. The "output" return
// value will be populated with the CreateVpnConnectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateVpnConnectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateVpnConnectionCommon Send returns without error.
//
// See CreateVpnConnection for more information on using the CreateVpnConnection
// API call, and error handling.
//
//    // Example sending a request using the CreateVpnConnectionRequest method.
//    req, resp := client.CreateVpnConnectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPN) CreateVpnConnectionRequest(input *CreateVpnConnectionInput) (req *request.Request, output *CreateVpnConnectionOutput) {
	op := &request.Operation{
		Name:       opCreateVpnConnection,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateVpnConnectionInput{}
	}

	output = &CreateVpnConnectionOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateVpnConnection API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPN's
// API operation CreateVpnConnection for usage and error information.
func (c *VPN) CreateVpnConnection(input *CreateVpnConnectionInput) (*CreateVpnConnectionOutput, error) {
	req, out := c.CreateVpnConnectionRequest(input)
	return out, req.Send()
}

// CreateVpnConnectionWithContext is the same as CreateVpnConnection with the addition of
// the ability to pass a context and additional request options.
//
// See CreateVpnConnection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) CreateVpnConnectionWithContext(ctx volcengine.Context, input *CreateVpnConnectionInput, opts ...request.Option) (*CreateVpnConnectionOutput, error) {
	req, out := c.CreateVpnConnectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateVpnConnectionInput struct {
	_ struct{} `type:"structure"`

	// CustomerGatewayId is a required field
	CustomerGatewayId *string `type:"string" required:"true"`

	Description *string `min:"1" max:"255" type:"string"`

	// IkeConfig is a required field
	IkeConfig *IkeConfigForCreateVpnConnectionInput `type:"structure" required:"true"`

	IpsecConfig *IpsecConfigForCreateVpnConnectionInput `type:"structure"`

	// LocalSubnet is a required field
	LocalSubnet *string `type:"string" required:"true"`

	NatTraversal *bool `type:"boolean"`

	// RemoteSubnet is a required field
	RemoteSubnet *string `type:"string" required:"true"`

	VpnConnectionName *string `min:"1" max:"128" type:"string"`

	// VpnGatewayId is a required field
	VpnGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateVpnConnectionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpnConnectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateVpnConnectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateVpnConnectionInput"}
	if s.CustomerGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomerGatewayId"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.IkeConfig == nil {
		invalidParams.Add(request.NewErrParamRequired("IkeConfig"))
	}
	if s.LocalSubnet == nil {
		invalidParams.Add(request.NewErrParamRequired("LocalSubnet"))
	}
	if s.RemoteSubnet == nil {
		invalidParams.Add(request.NewErrParamRequired("RemoteSubnet"))
	}
	if s.VpnConnectionName != nil && len(*s.VpnConnectionName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("VpnConnectionName", 1))
	}
	if s.VpnConnectionName != nil && len(*s.VpnConnectionName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("VpnConnectionName", 128, *s.VpnConnectionName))
	}
	if s.VpnGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpnGatewayId"))
	}
	if s.IkeConfig != nil {
		if err := s.IkeConfig.Validate(); err != nil {
			invalidParams.AddNested("IkeConfig", err.(request.ErrInvalidParams))
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCustomerGatewayId sets the CustomerGatewayId field's value.
func (s *CreateVpnConnectionInput) SetCustomerGatewayId(v string) *CreateVpnConnectionInput {
	s.CustomerGatewayId = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateVpnConnectionInput) SetDescription(v string) *CreateVpnConnectionInput {
	s.Description = &v
	return s
}

// SetIkeConfig sets the IkeConfig field's value.
func (s *CreateVpnConnectionInput) SetIkeConfig(v *IkeConfigForCreateVpnConnectionInput) *CreateVpnConnectionInput {
	s.IkeConfig = v
	return s
}

// SetIpsecConfig sets the IpsecConfig field's value.
func (s *CreateVpnConnectionInput) SetIpsecConfig(v *IpsecConfigForCreateVpnConnectionInput) *CreateVpnConnectionInput {
	s.IpsecConfig = v
	return s
}

// SetLocalSubnet sets the LocalSubnet field's value.
func (s *CreateVpnConnectionInput) SetLocalSubnet(v string) *CreateVpnConnectionInput {
	s.LocalSubnet = &v
	return s
}

// SetNatTraversal sets the NatTraversal field's value.
func (s *CreateVpnConnectionInput) SetNatTraversal(v bool) *CreateVpnConnectionInput {
	s.NatTraversal = &v
	return s
}

// SetRemoteSubnet sets the RemoteSubnet field's value.
func (s *CreateVpnConnectionInput) SetRemoteSubnet(v string) *CreateVpnConnectionInput {
	s.RemoteSubnet = &v
	return s
}

// SetVpnConnectionName sets the VpnConnectionName field's value.
func (s *CreateVpnConnectionInput) SetVpnConnectionName(v string) *CreateVpnConnectionInput {
	s.VpnConnectionName = &v
	return s
}

// SetVpnGatewayId sets the VpnGatewayId field's value.
func (s *CreateVpnConnectionInput) SetVpnGatewayId(v string) *CreateVpnConnectionInput {
	s.VpnGatewayId = &v
	return s
}

type CreateVpnConnectionOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`

	VpnConnectionId *string `type:"string"`
}

// String returns the string representation
func (s CreateVpnConnectionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateVpnConnectionOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *CreateVpnConnectionOutput) SetRequestId(v string) *CreateVpnConnectionOutput {
	s.RequestId = &v
	return s
}

// SetVpnConnectionId sets the VpnConnectionId field's value.
func (s *CreateVpnConnectionOutput) SetVpnConnectionId(v string) *CreateVpnConnectionOutput {
	s.VpnConnectionId = &v
	return s
}

type IkeConfigForCreateVpnConnectionInput struct {
	_ struct{} `type:"structure"`

	AuthAlg *string `type:"string" enum:"IkeConfigAuthAlgForCreateVpnConnectionInput"`

	DhGroup *string `type:"string" enum:"IkeConfigDhGroupForCreateVpnConnectionInput"`

	EncAlg *string `type:"string" enum:"IkeConfigEncAlgForCreateVpnConnectionInput"`

	Lifetime *string `type:"string"`

	LocalId *string `type:"string"`

	Mode *string `type:"string" enum:"IkeConfigModeForCreateVpnConnectionInput"`

	// Psk is a required field
	Psk *string `type:"string" required:"true"`

	RemoteId *string `type:"string"`

	Version *string `type:"string" enum:"IkeConfigVersionForCreateVpnConnectionInput"`
}

// String returns the string representation
func (s IkeConfigForCreateVpnConnectionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IkeConfigForCreateVpnConnectionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *IkeConfigForCreateVpnConnectionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "IkeConfigForCreateVpnConnectionInput"}
	if s.Psk == nil {
		invalidParams.Add(request.NewErrParamRequired("Psk"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAuthAlg sets the AuthAlg field's value.
func (s *IkeConfigForCreateVpnConnectionInput) SetAuthAlg(v string) *IkeConfigForCreateVpnConnectionInput {
	s.AuthAlg = &v
	return s
}

// SetDhGroup sets the DhGroup field's value.
func (s *IkeConfigForCreateVpnConnectionInput) SetDhGroup(v string) *IkeConfigForCreateVpnConnectionInput {
	s.DhGroup = &v
	return s
}

// SetEncAlg sets the EncAlg field's value.
func (s *IkeConfigForCreateVpnConnectionInput) SetEncAlg(v string) *IkeConfigForCreateVpnConnectionInput {
	s.EncAlg = &v
	return s
}

// SetLifetime sets the Lifetime field's value.
func (s *IkeConfigForCreateVpnConnectionInput) SetLifetime(v string) *IkeConfigForCreateVpnConnectionInput {
	s.Lifetime = &v
	return s
}

// SetLocalId sets the LocalId field's value.
func (s *IkeConfigForCreateVpnConnectionInput) SetLocalId(v string) *IkeConfigForCreateVpnConnectionInput {
	s.LocalId = &v
	return s
}

// SetMode sets the Mode field's value.
func (s *IkeConfigForCreateVpnConnectionInput) SetMode(v string) *IkeConfigForCreateVpnConnectionInput {
	s.Mode = &v
	return s
}

// SetPsk sets the Psk field's value.
func (s *IkeConfigForCreateVpnConnectionInput) SetPsk(v string) *IkeConfigForCreateVpnConnectionInput {
	s.Psk = &v
	return s
}

// SetRemoteId sets the RemoteId field's value.
func (s *IkeConfigForCreateVpnConnectionInput) SetRemoteId(v string) *IkeConfigForCreateVpnConnectionInput {
	s.RemoteId = &v
	return s
}

// SetVersion sets the Version field's value.
func (s *IkeConfigForCreateVpnConnectionInput) SetVersion(v string) *IkeConfigForCreateVpnConnectionInput {
	s.Version = &v
	return s
}

type IpsecConfigForCreateVpnConnectionInput struct {
	_ struct{} `type:"structure"`

	AuthAlg *string `type:"string" enum:"IpsecConfigAuthAlgForCreateVpnConnectionInput"`

	DhGroup *string `type:"string" enum:"IpsecConfigDhGroupForCreateVpnConnectionInput"`

	EncAlg *string `type:"string" enum:"IpsecConfigEncAlgForCreateVpnConnectionInput"`

	Lifetime *string `type:"string"`
}

// String returns the string representation
func (s IpsecConfigForCreateVpnConnectionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s IpsecConfigForCreateVpnConnectionInput) GoString() string {
	return s.String()
}

// SetAuthAlg sets the AuthAlg field's value.
func (s *IpsecConfigForCreateVpnConnectionInput) SetAuthAlg(v string) *IpsecConfigForCreateVpnConnectionInput {
	s.AuthAlg = &v
	return s
}

// SetDhGroup sets the DhGroup field's value.
func (s *IpsecConfigForCreateVpnConnectionInput) SetDhGroup(v string) *IpsecConfigForCreateVpnConnectionInput {
	s.DhGroup = &v
	return s
}

// SetEncAlg sets the EncAlg field's value.
func (s *IpsecConfigForCreateVpnConnectionInput) SetEncAlg(v string) *IpsecConfigForCreateVpnConnectionInput {
	s.EncAlg = &v
	return s
}

// SetLifetime sets the Lifetime field's value.
func (s *IpsecConfigForCreateVpnConnectionInput) SetLifetime(v string) *IpsecConfigForCreateVpnConnectionInput {
	s.Lifetime = &v
	return s
}

const (
	// IkeConfigAuthAlgForCreateVpnConnectionInputSha1 is a IkeConfigAuthAlgForCreateVpnConnectionInput enum value
	IkeConfigAuthAlgForCreateVpnConnectionInputSha1 = "sha1"

	// IkeConfigAuthAlgForCreateVpnConnectionInputMd5 is a IkeConfigAuthAlgForCreateVpnConnectionInput enum value
	IkeConfigAuthAlgForCreateVpnConnectionInputMd5 = "md5"

	// IkeConfigAuthAlgForCreateVpnConnectionInputSha256 is a IkeConfigAuthAlgForCreateVpnConnectionInput enum value
	IkeConfigAuthAlgForCreateVpnConnectionInputSha256 = "sha256"

	// IkeConfigAuthAlgForCreateVpnConnectionInputSha384 is a IkeConfigAuthAlgForCreateVpnConnectionInput enum value
	IkeConfigAuthAlgForCreateVpnConnectionInputSha384 = "sha384"

	// IkeConfigAuthAlgForCreateVpnConnectionInputSha512 is a IkeConfigAuthAlgForCreateVpnConnectionInput enum value
	IkeConfigAuthAlgForCreateVpnConnectionInputSha512 = "sha512"
)

const (
	// IkeConfigDhGroupForCreateVpnConnectionInputGroup1 is a IkeConfigDhGroupForCreateVpnConnectionInput enum value
	IkeConfigDhGroupForCreateVpnConnectionInputGroup1 = "group1"

	// IkeConfigDhGroupForCreateVpnConnectionInputGroup2 is a IkeConfigDhGroupForCreateVpnConnectionInput enum value
	IkeConfigDhGroupForCreateVpnConnectionInputGroup2 = "group2"

	// IkeConfigDhGroupForCreateVpnConnectionInputGroup5 is a IkeConfigDhGroupForCreateVpnConnectionInput enum value
	IkeConfigDhGroupForCreateVpnConnectionInputGroup5 = "group5"

	// IkeConfigDhGroupForCreateVpnConnectionInputGroup14 is a IkeConfigDhGroupForCreateVpnConnectionInput enum value
	IkeConfigDhGroupForCreateVpnConnectionInputGroup14 = "group14"
)

const (
	// IkeConfigEncAlgForCreateVpnConnectionInputAes is a IkeConfigEncAlgForCreateVpnConnectionInput enum value
	IkeConfigEncAlgForCreateVpnConnectionInputAes = "aes"

	// IkeConfigEncAlgForCreateVpnConnectionInputAes192 is a IkeConfigEncAlgForCreateVpnConnectionInput enum value
	IkeConfigEncAlgForCreateVpnConnectionInputAes192 = "aes192"

	// IkeConfigEncAlgForCreateVpnConnectionInputAes256 is a IkeConfigEncAlgForCreateVpnConnectionInput enum value
	IkeConfigEncAlgForCreateVpnConnectionInputAes256 = "aes256"

	// IkeConfigEncAlgForCreateVpnConnectionInputDes is a IkeConfigEncAlgForCreateVpnConnectionInput enum value
	IkeConfigEncAlgForCreateVpnConnectionInputDes = "des"

	// IkeConfigEncAlgForCreateVpnConnectionInput3des is a IkeConfigEncAlgForCreateVpnConnectionInput enum value
	IkeConfigEncAlgForCreateVpnConnectionInput3des = "3des"
)

const (
	// IkeConfigModeForCreateVpnConnectionInputMain is a IkeConfigModeForCreateVpnConnectionInput enum value
	IkeConfigModeForCreateVpnConnectionInputMain = "main"

	// IkeConfigModeForCreateVpnConnectionInputAggressive is a IkeConfigModeForCreateVpnConnectionInput enum value
	IkeConfigModeForCreateVpnConnectionInputAggressive = "aggressive"
)

const (
	// IkeConfigVersionForCreateVpnConnectionInputIkev1 is a IkeConfigVersionForCreateVpnConnectionInput enum value
	IkeConfigVersionForCreateVpnConnectionInputIkev1 = "ikev1"

	// IkeConfigVersionForCreateVpnConnectionInputIkev2 is a IkeConfigVersionForCreateVpnConnectionInput enum value
	IkeConfigVersionForCreateVpnConnectionInputIkev2 = "ikev2"
)

const (
	// IpsecConfigAuthAlgForCreateVpnConnectionInputSha1 is a IpsecConfigAuthAlgForCreateVpnConnectionInput enum value
	IpsecConfigAuthAlgForCreateVpnConnectionInputSha1 = "sha1"

	// IpsecConfigAuthAlgForCreateVpnConnectionInputMd5 is a IpsecConfigAuthAlgForCreateVpnConnectionInput enum value
	IpsecConfigAuthAlgForCreateVpnConnectionInputMd5 = "md5"

	// IpsecConfigAuthAlgForCreateVpnConnectionInputSha256 is a IpsecConfigAuthAlgForCreateVpnConnectionInput enum value
	IpsecConfigAuthAlgForCreateVpnConnectionInputSha256 = "sha256"

	// IpsecConfigAuthAlgForCreateVpnConnectionInputSha384 is a IpsecConfigAuthAlgForCreateVpnConnectionInput enum value
	IpsecConfigAuthAlgForCreateVpnConnectionInputSha384 = "sha384"

	// IpsecConfigAuthAlgForCreateVpnConnectionInputSha512 is a IpsecConfigAuthAlgForCreateVpnConnectionInput enum value
	IpsecConfigAuthAlgForCreateVpnConnectionInputSha512 = "sha512"
)

const (
	// IpsecConfigDhGroupForCreateVpnConnectionInputGroup1 is a IpsecConfigDhGroupForCreateVpnConnectionInput enum value
	IpsecConfigDhGroupForCreateVpnConnectionInputGroup1 = "group1"

	// IpsecConfigDhGroupForCreateVpnConnectionInputGroup2 is a IpsecConfigDhGroupForCreateVpnConnectionInput enum value
	IpsecConfigDhGroupForCreateVpnConnectionInputGroup2 = "group2"

	// IpsecConfigDhGroupForCreateVpnConnectionInputGroup5 is a IpsecConfigDhGroupForCreateVpnConnectionInput enum value
	IpsecConfigDhGroupForCreateVpnConnectionInputGroup5 = "group5"

	// IpsecConfigDhGroupForCreateVpnConnectionInputGroup14 is a IpsecConfigDhGroupForCreateVpnConnectionInput enum value
	IpsecConfigDhGroupForCreateVpnConnectionInputGroup14 = "group14"
)

const (
	// IpsecConfigEncAlgForCreateVpnConnectionInputAes is a IpsecConfigEncAlgForCreateVpnConnectionInput enum value
	IpsecConfigEncAlgForCreateVpnConnectionInputAes = "aes"

	// IpsecConfigEncAlgForCreateVpnConnectionInputAes192 is a IpsecConfigEncAlgForCreateVpnConnectionInput enum value
	IpsecConfigEncAlgForCreateVpnConnectionInputAes192 = "aes192"

	// IpsecConfigEncAlgForCreateVpnConnectionInputAes256 is a IpsecConfigEncAlgForCreateVpnConnectionInput enum value
	IpsecConfigEncAlgForCreateVpnConnectionInputAes256 = "aes256"

	// IpsecConfigEncAlgForCreateVpnConnectionInputDes is a IpsecConfigEncAlgForCreateVpnConnectionInput enum value
	IpsecConfigEncAlgForCreateVpnConnectionInputDes = "des"

	// IpsecConfigEncAlgForCreateVpnConnectionInput3des is a IpsecConfigEncAlgForCreateVpnConnectionInput enum value
	IpsecConfigEncAlgForCreateVpnConnectionInput3des = "3des"
)
