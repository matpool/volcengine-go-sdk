// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyVpnConnectionAttributesCommon = "ModifyVpnConnectionAttributes"

// ModifyVpnConnectionAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpnConnectionAttributesCommon operation. The "output" return
// value will be populated with the ModifyVpnConnectionAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpnConnectionAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpnConnectionAttributesCommon Send returns without error.
//
// See ModifyVpnConnectionAttributesCommon for more information on using the ModifyVpnConnectionAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyVpnConnectionAttributesCommonRequest method.
//	req, resp := client.ModifyVpnConnectionAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPN) ModifyVpnConnectionAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyVpnConnectionAttributesCommon,
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

// ModifyVpnConnectionAttributesCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation ModifyVpnConnectionAttributesCommon for usage and error information.
func (c *VPN) ModifyVpnConnectionAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyVpnConnectionAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyVpnConnectionAttributesCommonWithContext is the same as ModifyVpnConnectionAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpnConnectionAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) ModifyVpnConnectionAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyVpnConnectionAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyVpnConnectionAttributes = "ModifyVpnConnectionAttributes"

// ModifyVpnConnectionAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyVpnConnectionAttributes operation. The "output" return
// value will be populated with the ModifyVpnConnectionAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyVpnConnectionAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyVpnConnectionAttributesCommon Send returns without error.
//
// See ModifyVpnConnectionAttributes for more information on using the ModifyVpnConnectionAttributes
// API call, and error handling.
//
//	// Example sending a request using the ModifyVpnConnectionAttributesRequest method.
//	req, resp := client.ModifyVpnConnectionAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPN) ModifyVpnConnectionAttributesRequest(input *ModifyVpnConnectionAttributesInput) (req *request.Request, output *ModifyVpnConnectionAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyVpnConnectionAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpnConnectionAttributesInput{}
	}

	output = &ModifyVpnConnectionAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyVpnConnectionAttributes API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation ModifyVpnConnectionAttributes for usage and error information.
func (c *VPN) ModifyVpnConnectionAttributes(input *ModifyVpnConnectionAttributesInput) (*ModifyVpnConnectionAttributesOutput, error) {
	req, out := c.ModifyVpnConnectionAttributesRequest(input)
	return out, req.Send()
}

// ModifyVpnConnectionAttributesWithContext is the same as ModifyVpnConnectionAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyVpnConnectionAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) ModifyVpnConnectionAttributesWithContext(ctx volcengine.Context, input *ModifyVpnConnectionAttributesInput, opts ...request.Option) (*ModifyVpnConnectionAttributesOutput, error) {
	req, out := c.ModifyVpnConnectionAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyVpnConnectionAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `min:"1" max:"255" type:"string"`

	DpdAction *string `type:"string" enum:"DpdActionForModifyVpnConnectionAttributesInput"`

	IkeConfig *string `type:"string"`

	IpsecConfig *string `type:"string"`

	LocalSubnet []*string `type:"list"`

	LogEnabled *bool `type:"boolean"`

	NatTraversal *bool `type:"boolean"`

	NegotiateInstantly *bool `type:"boolean"`

	RemoteSubnet []*string `type:"list"`

	// VpnConnectionId is a required field
	VpnConnectionId *string `type:"string" required:"true"`

	VpnConnectionName *string `min:"1" max:"128" type:"string"`
}

// String returns the string representation
func (s ModifyVpnConnectionAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpnConnectionAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpnConnectionAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyVpnConnectionAttributesInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}
	if s.VpnConnectionId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpnConnectionId"))
	}
	if s.VpnConnectionName != nil && len(*s.VpnConnectionName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("VpnConnectionName", 1))
	}
	if s.VpnConnectionName != nil && len(*s.VpnConnectionName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("VpnConnectionName", 128, *s.VpnConnectionName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyVpnConnectionAttributesInput) SetDescription(v string) *ModifyVpnConnectionAttributesInput {
	s.Description = &v
	return s
}

// SetDpdAction sets the DpdAction field's value.
func (s *ModifyVpnConnectionAttributesInput) SetDpdAction(v string) *ModifyVpnConnectionAttributesInput {
	s.DpdAction = &v
	return s
}

// SetIkeConfig sets the IkeConfig field's value.
func (s *ModifyVpnConnectionAttributesInput) SetIkeConfig(v string) *ModifyVpnConnectionAttributesInput {
	s.IkeConfig = &v
	return s
}

// SetIpsecConfig sets the IpsecConfig field's value.
func (s *ModifyVpnConnectionAttributesInput) SetIpsecConfig(v string) *ModifyVpnConnectionAttributesInput {
	s.IpsecConfig = &v
	return s
}

// SetLocalSubnet sets the LocalSubnet field's value.
func (s *ModifyVpnConnectionAttributesInput) SetLocalSubnet(v []*string) *ModifyVpnConnectionAttributesInput {
	s.LocalSubnet = v
	return s
}

// SetLogEnabled sets the LogEnabled field's value.
func (s *ModifyVpnConnectionAttributesInput) SetLogEnabled(v bool) *ModifyVpnConnectionAttributesInput {
	s.LogEnabled = &v
	return s
}

// SetNatTraversal sets the NatTraversal field's value.
func (s *ModifyVpnConnectionAttributesInput) SetNatTraversal(v bool) *ModifyVpnConnectionAttributesInput {
	s.NatTraversal = &v
	return s
}

// SetNegotiateInstantly sets the NegotiateInstantly field's value.
func (s *ModifyVpnConnectionAttributesInput) SetNegotiateInstantly(v bool) *ModifyVpnConnectionAttributesInput {
	s.NegotiateInstantly = &v
	return s
}

// SetRemoteSubnet sets the RemoteSubnet field's value.
func (s *ModifyVpnConnectionAttributesInput) SetRemoteSubnet(v []*string) *ModifyVpnConnectionAttributesInput {
	s.RemoteSubnet = v
	return s
}

// SetVpnConnectionId sets the VpnConnectionId field's value.
func (s *ModifyVpnConnectionAttributesInput) SetVpnConnectionId(v string) *ModifyVpnConnectionAttributesInput {
	s.VpnConnectionId = &v
	return s
}

// SetVpnConnectionName sets the VpnConnectionName field's value.
func (s *ModifyVpnConnectionAttributesInput) SetVpnConnectionName(v string) *ModifyVpnConnectionAttributesInput {
	s.VpnConnectionName = &v
	return s
}

type ModifyVpnConnectionAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyVpnConnectionAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyVpnConnectionAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyVpnConnectionAttributesOutput) SetRequestId(v string) *ModifyVpnConnectionAttributesOutput {
	s.RequestId = &v
	return s
}

const (
	// DpdActionForModifyVpnConnectionAttributesInputNone is a DpdActionForModifyVpnConnectionAttributesInput enum value
	DpdActionForModifyVpnConnectionAttributesInputNone = "none"

	// DpdActionForModifyVpnConnectionAttributesInputClear is a DpdActionForModifyVpnConnectionAttributesInput enum value
	DpdActionForModifyVpnConnectionAttributesInputClear = "clear"

	// DpdActionForModifyVpnConnectionAttributesInputHold is a DpdActionForModifyVpnConnectionAttributesInput enum value
	DpdActionForModifyVpnConnectionAttributesInputHold = "hold"

	// DpdActionForModifyVpnConnectionAttributesInputRestart is a DpdActionForModifyVpnConnectionAttributesInput enum value
	DpdActionForModifyVpnConnectionAttributesInputRestart = "restart"
)
