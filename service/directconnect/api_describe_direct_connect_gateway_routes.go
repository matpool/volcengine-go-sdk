// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDirectConnectGatewayRoutesCommon = "DescribeDirectConnectGatewayRoutes"

// DescribeDirectConnectGatewayRoutesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectGatewayRoutesCommon operation. The "output" return
// value will be populated with the DescribeDirectConnectGatewayRoutesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectGatewayRoutesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectGatewayRoutesCommon Send returns without error.
//
// See DescribeDirectConnectGatewayRoutesCommon for more information on using the DescribeDirectConnectGatewayRoutesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeDirectConnectGatewayRoutesCommonRequest method.
//	req, resp := client.DescribeDirectConnectGatewayRoutesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRoutesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectGatewayRoutesCommon,
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

// DescribeDirectConnectGatewayRoutesCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectGatewayRoutesCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRoutesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectGatewayRoutesCommonRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectGatewayRoutesCommonWithContext is the same as DescribeDirectConnectGatewayRoutesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectGatewayRoutesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRoutesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectGatewayRoutesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDirectConnectGatewayRoutes = "DescribeDirectConnectGatewayRoutes"

// DescribeDirectConnectGatewayRoutesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectGatewayRoutes operation. The "output" return
// value will be populated with the DescribeDirectConnectGatewayRoutesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectGatewayRoutesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectGatewayRoutesCommon Send returns without error.
//
// See DescribeDirectConnectGatewayRoutes for more information on using the DescribeDirectConnectGatewayRoutes
// API call, and error handling.
//
//	// Example sending a request using the DescribeDirectConnectGatewayRoutesRequest method.
//	req, resp := client.DescribeDirectConnectGatewayRoutesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRoutesRequest(input *DescribeDirectConnectGatewayRoutesInput) (req *request.Request, output *DescribeDirectConnectGatewayRoutesOutput) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectGatewayRoutes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDirectConnectGatewayRoutesInput{}
	}

	output = &DescribeDirectConnectGatewayRoutesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDirectConnectGatewayRoutes API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectGatewayRoutes for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRoutes(input *DescribeDirectConnectGatewayRoutesInput) (*DescribeDirectConnectGatewayRoutesOutput, error) {
	req, out := c.DescribeDirectConnectGatewayRoutesRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectGatewayRoutesWithContext is the same as DescribeDirectConnectGatewayRoutes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectGatewayRoutes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayRoutesWithContext(ctx volcengine.Context, input *DescribeDirectConnectGatewayRoutesInput, opts ...request.Option) (*DescribeDirectConnectGatewayRoutesOutput, error) {
	req, out := c.DescribeDirectConnectGatewayRoutesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDirectConnectGatewayRoutesInput struct {
	_ struct{} `type:"structure"`

	DestinationCidrBlock *string `type:"string"`

	DirectConnectGatewayId *string `type:"string"`

	DirectConnectGatewayRouteIds []*string `type:"list"`

	NextHopId *string `type:"string"`

	NextHopType *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RouteType *string `type:"string"`
}

// String returns the string representation
func (s DescribeDirectConnectGatewayRoutesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectGatewayRoutesInput) GoString() string {
	return s.String()
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *DescribeDirectConnectGatewayRoutesInput) SetDestinationCidrBlock(v string) *DescribeDirectConnectGatewayRoutesInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *DescribeDirectConnectGatewayRoutesInput) SetDirectConnectGatewayId(v string) *DescribeDirectConnectGatewayRoutesInput {
	s.DirectConnectGatewayId = &v
	return s
}

// SetDirectConnectGatewayRouteIds sets the DirectConnectGatewayRouteIds field's value.
func (s *DescribeDirectConnectGatewayRoutesInput) SetDirectConnectGatewayRouteIds(v []*string) *DescribeDirectConnectGatewayRoutesInput {
	s.DirectConnectGatewayRouteIds = v
	return s
}

// SetNextHopId sets the NextHopId field's value.
func (s *DescribeDirectConnectGatewayRoutesInput) SetNextHopId(v string) *DescribeDirectConnectGatewayRoutesInput {
	s.NextHopId = &v
	return s
}

// SetNextHopType sets the NextHopType field's value.
func (s *DescribeDirectConnectGatewayRoutesInput) SetNextHopType(v string) *DescribeDirectConnectGatewayRoutesInput {
	s.NextHopType = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDirectConnectGatewayRoutesInput) SetPageNumber(v int64) *DescribeDirectConnectGatewayRoutesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDirectConnectGatewayRoutesInput) SetPageSize(v int64) *DescribeDirectConnectGatewayRoutesInput {
	s.PageSize = &v
	return s
}

// SetRouteType sets the RouteType field's value.
func (s *DescribeDirectConnectGatewayRoutesInput) SetRouteType(v string) *DescribeDirectConnectGatewayRoutesInput {
	s.RouteType = &v
	return s
}

type DescribeDirectConnectGatewayRoutesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DirectConnectGatewayRoutes []*DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDirectConnectGatewayRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectGatewayRoutesOutput) GoString() string {
	return s.String()
}

// SetDirectConnectGatewayRoutes sets the DirectConnectGatewayRoutes field's value.
func (s *DescribeDirectConnectGatewayRoutesOutput) SetDirectConnectGatewayRoutes(v []*DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) *DescribeDirectConnectGatewayRoutesOutput {
	s.DirectConnectGatewayRoutes = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDirectConnectGatewayRoutesOutput) SetPageNumber(v int64) *DescribeDirectConnectGatewayRoutesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDirectConnectGatewayRoutesOutput) SetPageSize(v int64) *DescribeDirectConnectGatewayRoutesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeDirectConnectGatewayRoutesOutput) SetRequestId(v string) *DescribeDirectConnectGatewayRoutesOutput {
	s.RequestId = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeDirectConnectGatewayRoutesOutput) SetTotalCount(v int64) *DescribeDirectConnectGatewayRoutesOutput {
	s.TotalCount = &v
	return s
}

type DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	CreationTime *string `type:"string"`

	DestinationCidrBlock *string `type:"string"`

	DirectConnectGatewayId *string `type:"string"`

	DirectConnectGatewayRouteId *string `type:"string"`

	NextHopId *string `type:"string"`

	NextHopType *string `type:"string"`

	RouteType *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) SetAccountId(v string) *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput {
	s.AccountId = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) SetCreationTime(v string) *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput {
	s.CreationTime = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) SetDestinationCidrBlock(v string) *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput {
	s.DestinationCidrBlock = &v
	return s
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) SetDirectConnectGatewayId(v string) *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput {
	s.DirectConnectGatewayId = &v
	return s
}

// SetDirectConnectGatewayRouteId sets the DirectConnectGatewayRouteId field's value.
func (s *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) SetDirectConnectGatewayRouteId(v string) *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput {
	s.DirectConnectGatewayRouteId = &v
	return s
}

// SetNextHopId sets the NextHopId field's value.
func (s *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) SetNextHopId(v string) *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput {
	s.NextHopId = &v
	return s
}

// SetNextHopType sets the NextHopType field's value.
func (s *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) SetNextHopType(v string) *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput {
	s.NextHopType = &v
	return s
}

// SetRouteType sets the RouteType field's value.
func (s *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) SetRouteType(v string) *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput {
	s.RouteType = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput) SetStatus(v string) *DirectConnectGatewayRouteForDescribeDirectConnectGatewayRoutesOutput {
	s.Status = &v
	return s
}
