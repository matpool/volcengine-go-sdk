// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpn

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyCustomerGatewayAttributesCommon = "ModifyCustomerGatewayAttributes"

// ModifyCustomerGatewayAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyCustomerGatewayAttributesCommon operation. The "output" return
// value will be populated with the ModifyCustomerGatewayAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyCustomerGatewayAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyCustomerGatewayAttributesCommon Send returns without error.
//
// See ModifyCustomerGatewayAttributesCommon for more information on using the ModifyCustomerGatewayAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyCustomerGatewayAttributesCommonRequest method.
//	req, resp := client.ModifyCustomerGatewayAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPN) ModifyCustomerGatewayAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyCustomerGatewayAttributesCommon,
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

// ModifyCustomerGatewayAttributesCommon API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation ModifyCustomerGatewayAttributesCommon for usage and error information.
func (c *VPN) ModifyCustomerGatewayAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyCustomerGatewayAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyCustomerGatewayAttributesCommonWithContext is the same as ModifyCustomerGatewayAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyCustomerGatewayAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) ModifyCustomerGatewayAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyCustomerGatewayAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyCustomerGatewayAttributes = "ModifyCustomerGatewayAttributes"

// ModifyCustomerGatewayAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyCustomerGatewayAttributes operation. The "output" return
// value will be populated with the ModifyCustomerGatewayAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyCustomerGatewayAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyCustomerGatewayAttributesCommon Send returns without error.
//
// See ModifyCustomerGatewayAttributes for more information on using the ModifyCustomerGatewayAttributes
// API call, and error handling.
//
//	// Example sending a request using the ModifyCustomerGatewayAttributesRequest method.
//	req, resp := client.ModifyCustomerGatewayAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPN) ModifyCustomerGatewayAttributesRequest(input *ModifyCustomerGatewayAttributesInput) (req *request.Request, output *ModifyCustomerGatewayAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyCustomerGatewayAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyCustomerGatewayAttributesInput{}
	}

	output = &ModifyCustomerGatewayAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyCustomerGatewayAttributes API operation for VPN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPN's
// API operation ModifyCustomerGatewayAttributes for usage and error information.
func (c *VPN) ModifyCustomerGatewayAttributes(input *ModifyCustomerGatewayAttributesInput) (*ModifyCustomerGatewayAttributesOutput, error) {
	req, out := c.ModifyCustomerGatewayAttributesRequest(input)
	return out, req.Send()
}

// ModifyCustomerGatewayAttributesWithContext is the same as ModifyCustomerGatewayAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyCustomerGatewayAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPN) ModifyCustomerGatewayAttributesWithContext(ctx volcengine.Context, input *ModifyCustomerGatewayAttributesInput, opts ...request.Option) (*ModifyCustomerGatewayAttributesOutput, error) {
	req, out := c.ModifyCustomerGatewayAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyCustomerGatewayAttributesInput struct {
	_ struct{} `type:"structure"`

	// CustomerGatewayId is a required field
	CustomerGatewayId *string `type:"string" required:"true"`

	CustomerGatewayName *string `min:"1" max:"128" type:"string"`

	Description *string `min:"1" max:"255" type:"string"`
}

// String returns the string representation
func (s ModifyCustomerGatewayAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyCustomerGatewayAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyCustomerGatewayAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyCustomerGatewayAttributesInput"}
	if s.CustomerGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("CustomerGatewayId"))
	}
	if s.CustomerGatewayName != nil && len(*s.CustomerGatewayName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CustomerGatewayName", 1))
	}
	if s.CustomerGatewayName != nil && len(*s.CustomerGatewayName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("CustomerGatewayName", 128, *s.CustomerGatewayName))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Description", 1))
	}
	if s.Description != nil && len(*s.Description) > 255 {
		invalidParams.Add(request.NewErrParamMaxLen("Description", 255, *s.Description))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCustomerGatewayId sets the CustomerGatewayId field's value.
func (s *ModifyCustomerGatewayAttributesInput) SetCustomerGatewayId(v string) *ModifyCustomerGatewayAttributesInput {
	s.CustomerGatewayId = &v
	return s
}

// SetCustomerGatewayName sets the CustomerGatewayName field's value.
func (s *ModifyCustomerGatewayAttributesInput) SetCustomerGatewayName(v string) *ModifyCustomerGatewayAttributesInput {
	s.CustomerGatewayName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyCustomerGatewayAttributesInput) SetDescription(v string) *ModifyCustomerGatewayAttributesInput {
	s.Description = &v
	return s
}

type ModifyCustomerGatewayAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyCustomerGatewayAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyCustomerGatewayAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyCustomerGatewayAttributesOutput) SetRequestId(v string) *ModifyCustomerGatewayAttributesOutput {
	s.RequestId = &v
	return s
}
