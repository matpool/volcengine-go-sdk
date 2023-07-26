// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opAddBandwidthPackageIpCommon = "AddBandwidthPackageIp"

// AddBandwidthPackageIpCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the AddBandwidthPackageIpCommon operation. The "output" return
// value will be populated with the AddBandwidthPackageIpCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddBandwidthPackageIpCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddBandwidthPackageIpCommon Send returns without error.
//
// See AddBandwidthPackageIpCommon for more information on using the AddBandwidthPackageIpCommon
// API call, and error handling.
//
//	// Example sending a request using the AddBandwidthPackageIpCommonRequest method.
//	req, resp := client.AddBandwidthPackageIpCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AddBandwidthPackageIpCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opAddBandwidthPackageIpCommon,
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

// AddBandwidthPackageIpCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AddBandwidthPackageIpCommon for usage and error information.
func (c *VPC) AddBandwidthPackageIpCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.AddBandwidthPackageIpCommonRequest(input)
	return out, req.Send()
}

// AddBandwidthPackageIpCommonWithContext is the same as AddBandwidthPackageIpCommon with the addition of
// the ability to pass a context and additional request options.
//
// See AddBandwidthPackageIpCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AddBandwidthPackageIpCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.AddBandwidthPackageIpCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opAddBandwidthPackageIp = "AddBandwidthPackageIp"

// AddBandwidthPackageIpRequest generates a "volcengine/request.Request" representing the
// client's request for the AddBandwidthPackageIp operation. The "output" return
// value will be populated with the AddBandwidthPackageIpCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned AddBandwidthPackageIpCommon Request to send the API call to the service.
// the "output" return value is not valid until after AddBandwidthPackageIpCommon Send returns without error.
//
// See AddBandwidthPackageIp for more information on using the AddBandwidthPackageIp
// API call, and error handling.
//
//	// Example sending a request using the AddBandwidthPackageIpRequest method.
//	req, resp := client.AddBandwidthPackageIpRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) AddBandwidthPackageIpRequest(input *AddBandwidthPackageIpInput) (req *request.Request, output *AddBandwidthPackageIpOutput) {
	op := &request.Operation{
		Name:       opAddBandwidthPackageIp,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddBandwidthPackageIpInput{}
	}

	output = &AddBandwidthPackageIpOutput{}
	req = c.newRequest(op, input, output)

	return
}

// AddBandwidthPackageIp API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation AddBandwidthPackageIp for usage and error information.
func (c *VPC) AddBandwidthPackageIp(input *AddBandwidthPackageIpInput) (*AddBandwidthPackageIpOutput, error) {
	req, out := c.AddBandwidthPackageIpRequest(input)
	return out, req.Send()
}

// AddBandwidthPackageIpWithContext is the same as AddBandwidthPackageIp with the addition of
// the ability to pass a context and additional request options.
//
// See AddBandwidthPackageIp for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) AddBandwidthPackageIpWithContext(ctx volcengine.Context, input *AddBandwidthPackageIpInput, opts ...request.Option) (*AddBandwidthPackageIpOutput, error) {
	req, out := c.AddBandwidthPackageIpRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddBandwidthPackageIpInput struct {
	_ struct{} `type:"structure"`

	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`

	// BandwidthPackageId is a required field
	BandwidthPackageId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddBandwidthPackageIpInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddBandwidthPackageIpInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddBandwidthPackageIpInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddBandwidthPackageIpInput"}
	if s.AllocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllocationId"))
	}
	if s.BandwidthPackageId == nil {
		invalidParams.Add(request.NewErrParamRequired("BandwidthPackageId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllocationId sets the AllocationId field's value.
func (s *AddBandwidthPackageIpInput) SetAllocationId(v string) *AddBandwidthPackageIpInput {
	s.AllocationId = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *AddBandwidthPackageIpInput) SetBandwidthPackageId(v string) *AddBandwidthPackageIpInput {
	s.BandwidthPackageId = &v
	return s
}

type AddBandwidthPackageIpOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s AddBandwidthPackageIpOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AddBandwidthPackageIpOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *AddBandwidthPackageIpOutput) SetRequestId(v string) *AddBandwidthPackageIpOutput {
	s.RequestId = &v
	return s
}
