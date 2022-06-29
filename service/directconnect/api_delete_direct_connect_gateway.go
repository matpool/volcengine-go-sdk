// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteDirectConnectGatewayCommon = "DeleteDirectConnectGateway"

// DeleteDirectConnectGatewayCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDirectConnectGatewayCommon operation. The "output" return
// value will be populated with the DeleteDirectConnectGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDirectConnectGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDirectConnectGatewayCommon Send returns without error.
//
// See DeleteDirectConnectGatewayCommon for more information on using the DeleteDirectConnectGatewayCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteDirectConnectGatewayCommonRequest method.
//    req, resp := client.DeleteDirectConnectGatewayCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDirectConnectGatewayCommon,
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

// DeleteDirectConnectGatewayCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation DeleteDirectConnectGatewayCommon for usage and error information.
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDirectConnectGatewayCommonRequest(input)
	return out, req.Send()
}

// DeleteDirectConnectGatewayCommonWithContext is the same as DeleteDirectConnectGatewayCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDirectConnectGatewayCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDirectConnectGatewayCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDirectConnectGateway = "DeleteDirectConnectGateway"

// DeleteDirectConnectGatewayRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDirectConnectGateway operation. The "output" return
// value will be populated with the DeleteDirectConnectGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDirectConnectGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDirectConnectGatewayCommon Send returns without error.
//
// See DeleteDirectConnectGateway for more information on using the DeleteDirectConnectGateway
// API call, and error handling.
//
//    // Example sending a request using the DeleteDirectConnectGatewayRequest method.
//    req, resp := client.DeleteDirectConnectGatewayRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayRequest(input *DeleteDirectConnectGatewayInput) (req *request.Request, output *DeleteDirectConnectGatewayOutput) {
	op := &request.Operation{
		Name:       opDeleteDirectConnectGateway,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDirectConnectGatewayInput{}
	}

	output = &DeleteDirectConnectGatewayOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteDirectConnectGateway API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation DeleteDirectConnectGateway for usage and error information.
func (c *DIRECTCONNECT) DeleteDirectConnectGateway(input *DeleteDirectConnectGatewayInput) (*DeleteDirectConnectGatewayOutput, error) {
	req, out := c.DeleteDirectConnectGatewayRequest(input)
	return out, req.Send()
}

// DeleteDirectConnectGatewayWithContext is the same as DeleteDirectConnectGateway with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDirectConnectGateway for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DeleteDirectConnectGatewayWithContext(ctx volcengine.Context, input *DeleteDirectConnectGatewayInput, opts ...request.Option) (*DeleteDirectConnectGatewayOutput, error) {
	req, out := c.DeleteDirectConnectGatewayRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteDirectConnectGatewayInput struct {
	_ struct{} `type:"structure"`

	// DirectConnectGatewayId is a required field
	DirectConnectGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDirectConnectGatewayInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDirectConnectGatewayInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteDirectConnectGatewayInput"}
	if s.DirectConnectGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *DeleteDirectConnectGatewayInput) SetDirectConnectGatewayId(v string) *DeleteDirectConnectGatewayInput {
	s.DirectConnectGatewayId = &v
	return s
}

type DeleteDirectConnectGatewayOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PreOrderNumber *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteDirectConnectGatewayOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDirectConnectGatewayOutput) GoString() string {
	return s.String()
}

// SetPreOrderNumber sets the PreOrderNumber field's value.
func (s *DeleteDirectConnectGatewayOutput) SetPreOrderNumber(v string) *DeleteDirectConnectGatewayOutput {
	s.PreOrderNumber = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteDirectConnectGatewayOutput) SetRequestId(v string) *DeleteDirectConnectGatewayOutput {
	s.RequestId = &v
	return s
}
