// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyBandwidthPackageAttributesCommon = "ModifyBandwidthPackageAttributes"

// ModifyBandwidthPackageAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyBandwidthPackageAttributesCommon operation. The "output" return
// value will be populated with the ModifyBandwidthPackageAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyBandwidthPackageAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyBandwidthPackageAttributesCommon Send returns without error.
//
// See ModifyBandwidthPackageAttributesCommon for more information on using the ModifyBandwidthPackageAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyBandwidthPackageAttributesCommonRequest method.
//	req, resp := client.ModifyBandwidthPackageAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ModifyBandwidthPackageAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyBandwidthPackageAttributesCommon,
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

// ModifyBandwidthPackageAttributesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyBandwidthPackageAttributesCommon for usage and error information.
func (c *VPC) ModifyBandwidthPackageAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyBandwidthPackageAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyBandwidthPackageAttributesCommonWithContext is the same as ModifyBandwidthPackageAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyBandwidthPackageAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyBandwidthPackageAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyBandwidthPackageAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyBandwidthPackageAttributes = "ModifyBandwidthPackageAttributes"

// ModifyBandwidthPackageAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyBandwidthPackageAttributes operation. The "output" return
// value will be populated with the ModifyBandwidthPackageAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyBandwidthPackageAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyBandwidthPackageAttributesCommon Send returns without error.
//
// See ModifyBandwidthPackageAttributes for more information on using the ModifyBandwidthPackageAttributes
// API call, and error handling.
//
//	// Example sending a request using the ModifyBandwidthPackageAttributesRequest method.
//	req, resp := client.ModifyBandwidthPackageAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) ModifyBandwidthPackageAttributesRequest(input *ModifyBandwidthPackageAttributesInput) (req *request.Request, output *ModifyBandwidthPackageAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyBandwidthPackageAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyBandwidthPackageAttributesInput{}
	}

	output = &ModifyBandwidthPackageAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyBandwidthPackageAttributes API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation ModifyBandwidthPackageAttributes for usage and error information.
func (c *VPC) ModifyBandwidthPackageAttributes(input *ModifyBandwidthPackageAttributesInput) (*ModifyBandwidthPackageAttributesOutput, error) {
	req, out := c.ModifyBandwidthPackageAttributesRequest(input)
	return out, req.Send()
}

// ModifyBandwidthPackageAttributesWithContext is the same as ModifyBandwidthPackageAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyBandwidthPackageAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) ModifyBandwidthPackageAttributesWithContext(ctx volcengine.Context, input *ModifyBandwidthPackageAttributesInput, opts ...request.Option) (*ModifyBandwidthPackageAttributesOutput, error) {
	req, out := c.ModifyBandwidthPackageAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyBandwidthPackageAttributesInput struct {
	_ struct{} `type:"structure"`

	// BandwidthPackageId is a required field
	BandwidthPackageId *string `type:"string" required:"true"`

	BandwidthPackageName *string `min:"1" max:"128" type:"string"`

	Description *string `min:"1" max:"255" type:"string"`
}

// String returns the string representation
func (s ModifyBandwidthPackageAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyBandwidthPackageAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyBandwidthPackageAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyBandwidthPackageAttributesInput"}
	if s.BandwidthPackageId == nil {
		invalidParams.Add(request.NewErrParamRequired("BandwidthPackageId"))
	}
	if s.BandwidthPackageName != nil && len(*s.BandwidthPackageName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("BandwidthPackageName", 1))
	}
	if s.BandwidthPackageName != nil && len(*s.BandwidthPackageName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("BandwidthPackageName", 128, *s.BandwidthPackageName))
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

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *ModifyBandwidthPackageAttributesInput) SetBandwidthPackageId(v string) *ModifyBandwidthPackageAttributesInput {
	s.BandwidthPackageId = &v
	return s
}

// SetBandwidthPackageName sets the BandwidthPackageName field's value.
func (s *ModifyBandwidthPackageAttributesInput) SetBandwidthPackageName(v string) *ModifyBandwidthPackageAttributesInput {
	s.BandwidthPackageName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ModifyBandwidthPackageAttributesInput) SetDescription(v string) *ModifyBandwidthPackageAttributesInput {
	s.Description = &v
	return s
}

type ModifyBandwidthPackageAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyBandwidthPackageAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyBandwidthPackageAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyBandwidthPackageAttributesOutput) SetRequestId(v string) *ModifyBandwidthPackageAttributesOutput {
	s.RequestId = &v
	return s
}
