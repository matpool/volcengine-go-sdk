// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteVpcCommon = "DeleteVpc"

// DeleteVpcCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpcCommon operation. The "output" return
// value will be populated with the DeleteVpcCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpcCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpcCommon Send returns without error.
//
// See DeleteVpcCommon for more information on using the DeleteVpcCommon
// API call, and error handling.
//
//	// Example sending a request using the DeleteVpcCommonRequest method.
//	req, resp := client.DeleteVpcCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DeleteVpcCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteVpcCommon,
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

// DeleteVpcCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DeleteVpcCommon for usage and error information.
func (c *VPC) DeleteVpcCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteVpcCommonRequest(input)
	return out, req.Send()
}

// DeleteVpcCommonWithContext is the same as DeleteVpcCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpcCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteVpcCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteVpcCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteVpc = "DeleteVpc"

// DeleteVpcRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteVpc operation. The "output" return
// value will be populated with the DeleteVpcCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteVpcCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteVpcCommon Send returns without error.
//
// See DeleteVpc for more information on using the DeleteVpc
// API call, and error handling.
//
//	// Example sending a request using the DeleteVpcRequest method.
//	req, resp := client.DeleteVpcRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *VPC) DeleteVpcRequest(input *DeleteVpcInput) (req *request.Request, output *DeleteVpcOutput) {
	op := &request.Operation{
		Name:       opDeleteVpc,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteVpcInput{}
	}

	output = &DeleteVpcOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteVpc API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DeleteVpc for usage and error information.
func (c *VPC) DeleteVpc(input *DeleteVpcInput) (*DeleteVpcOutput, error) {
	req, out := c.DeleteVpcRequest(input)
	return out, req.Send()
}

// DeleteVpcWithContext is the same as DeleteVpc with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteVpc for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DeleteVpcWithContext(ctx volcengine.Context, input *DeleteVpcInput, opts ...request.Option) (*DeleteVpcOutput, error) {
	req, out := c.DeleteVpcRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteVpcInput struct {
	_ struct{} `type:"structure"`

	// VpcId is a required field
	VpcId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteVpcInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpcInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteVpcInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteVpcInput"}
	if s.VpcId == nil {
		invalidParams.Add(request.NewErrParamRequired("VpcId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetVpcId sets the VpcId field's value.
func (s *DeleteVpcInput) SetVpcId(v string) *DeleteVpcInput {
	s.VpcId = &v
	return s
}

type DeleteVpcOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteVpcOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteVpcOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteVpcOutput) SetRequestId(v string) *DeleteVpcOutput {
	s.RequestId = &v
	return s
}
