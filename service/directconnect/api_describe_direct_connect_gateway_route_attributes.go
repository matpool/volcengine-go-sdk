// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDirectConnectGatewayRouteAttributesCommon = "DescribeDirectConnectGatewayRouteAttributes"

// DescribeDirectConnectGatewayRouteAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectGatewayRouteAttributesCommon operation. The "output" return
// value will be populated with the DescribeDirectConnectGatewayRouteAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectGatewayRouteAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectGatewayRouteAttributesCommon Send returns without error.
//
// See DescribeDirectConnectGatewayRouteAttributesCommon for more information on using the DescribeDirectConnectGatewayRouteAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeDirectConnectGatewayRouteAttributesCommonRequest method.
//	req, resp := client.DescribeDirectConnectGatewayRouteAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRouteAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectGatewayRouteAttributesCommon,
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

// DescribeDirectConnectGatewayRouteAttributesCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectGatewayRouteAttributesCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRouteAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectGatewayRouteAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectGatewayRouteAttributesCommonWithContext is the same as DescribeDirectConnectGatewayRouteAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectGatewayRouteAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRouteAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectGatewayRouteAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDirectConnectGatewayRouteAttributes = "DescribeDirectConnectGatewayRouteAttributes"

// DescribeDirectConnectGatewayRouteAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectGatewayRouteAttributes operation. The "output" return
// value will be populated with the DescribeDirectConnectGatewayRouteAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectGatewayRouteAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectGatewayRouteAttributesCommon Send returns without error.
//
// See DescribeDirectConnectGatewayRouteAttributes for more information on using the DescribeDirectConnectGatewayRouteAttributes
// API call, and error handling.
//
//	// Example sending a request using the DescribeDirectConnectGatewayRouteAttributesRequest method.
//	req, resp := client.DescribeDirectConnectGatewayRouteAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRouteAttributesRequest(input *DescribeDirectConnectGatewayRouteAttributesInput) (req *request.Request, output *DescribeDirectConnectGatewayRouteAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectGatewayRouteAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDirectConnectGatewayRouteAttributesInput{}
	}

	output = &DescribeDirectConnectGatewayRouteAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDirectConnectGatewayRouteAttributes API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectGatewayRouteAttributes for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRouteAttributes(input *DescribeDirectConnectGatewayRouteAttributesInput) (*DescribeDirectConnectGatewayRouteAttributesOutput, error) {
	req, out := c.DescribeDirectConnectGatewayRouteAttributesRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectGatewayRouteAttributesWithContext is the same as DescribeDirectConnectGatewayRouteAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectGatewayRouteAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRouteAttributesWithContext(ctx volcengine.Context, input *DescribeDirectConnectGatewayRouteAttributesInput, opts ...request.Option) (*DescribeDirectConnectGatewayRouteAttributesOutput, error) {
	req, out := c.DescribeDirectConnectGatewayRouteAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDirectConnectGatewayRouteAttributesInput struct {
	_ struct{} `type:"structure"`

	// DirectConnectGatewayRouteId is a required field
	DirectConnectGatewayRouteId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDirectConnectGatewayRouteAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectGatewayRouteAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDirectConnectGatewayRouteAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDirectConnectGatewayRouteAttributesInput"}
	if s.DirectConnectGatewayRouteId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectGatewayRouteId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDirectConnectGatewayRouteId sets the DirectConnectGatewayRouteId field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesInput) SetDirectConnectGatewayRouteId(v string) *DescribeDirectConnectGatewayRouteAttributesInput {
	s.DirectConnectGatewayRouteId = &v
	return s
}

type DescribeDirectConnectGatewayRouteAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string"`

	CreationTime *string `type:"string"`

	DestinationCidrBlock *string `type:"string"`

	DirectConnectGatewayId *string `type:"string"`

	DirectConnectGatewayRouteId *string `type:"string"`

	NextHopId *string `type:"string"`

	NextHopType *string `type:"string"`

	RequestId *string `type:"string"`

	RouteType *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DescribeDirectConnectGatewayRouteAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectGatewayRouteAttributesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetAccountId(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.AccountId = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetCreationTime(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetDestinationCidrBlock(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.DestinationCidrBlock = &v
	return s
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetDirectConnectGatewayId(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.DirectConnectGatewayId = &v
	return s
}

// SetDirectConnectGatewayRouteId sets the DirectConnectGatewayRouteId field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetDirectConnectGatewayRouteId(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.DirectConnectGatewayRouteId = &v
	return s
}

// SetNextHopId sets the NextHopId field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetNextHopId(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.NextHopId = &v
	return s
}

// SetNextHopType sets the NextHopType field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetNextHopType(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.NextHopType = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetRequestId(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.RequestId = &v
	return s
}

// SetRouteType sets the RouteType field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetRouteType(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.RouteType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeDirectConnectGatewayRouteAttributesOutput) SetStatus(v string) *DescribeDirectConnectGatewayRouteAttributesOutput {
	s.Status = &v
	return s
}
