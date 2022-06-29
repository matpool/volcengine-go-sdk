// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeRouteEntryListCommon = "DescribeRouteEntryList"

// DescribeRouteEntryListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeRouteEntryListCommon operation. The "output" return
// value will be populated with the DescribeRouteEntryListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeRouteEntryListCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeRouteEntryListCommon Send returns without error.
//
// See DescribeRouteEntryListCommon for more information on using the DescribeRouteEntryListCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeRouteEntryListCommonRequest method.
//    req, resp := client.DescribeRouteEntryListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeRouteEntryListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeRouteEntryListCommon,
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

// DescribeRouteEntryListCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation DescribeRouteEntryListCommon for usage and error information.
func (c *VPC) DescribeRouteEntryListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeRouteEntryListCommonRequest(input)
	return out, req.Send()
}

// DescribeRouteEntryListCommonWithContext is the same as DescribeRouteEntryListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeRouteEntryListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeRouteEntryListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeRouteEntryListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeRouteEntryList = "DescribeRouteEntryList"

// DescribeRouteEntryListRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeRouteEntryList operation. The "output" return
// value will be populated with the DescribeRouteEntryListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeRouteEntryListCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeRouteEntryListCommon Send returns without error.
//
// See DescribeRouteEntryList for more information on using the DescribeRouteEntryList
// API call, and error handling.
//
//    // Example sending a request using the DescribeRouteEntryListRequest method.
//    req, resp := client.DescribeRouteEntryListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeRouteEntryListRequest(input *DescribeRouteEntryListInput) (req *request.Request, output *DescribeRouteEntryListOutput) {
	op := &request.Operation{
		Name:       opDescribeRouteEntryList,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeRouteEntryListInput{}
	}

	output = &DescribeRouteEntryListOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeRouteEntryList API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for VPC's
// API operation DescribeRouteEntryList for usage and error information.
func (c *VPC) DescribeRouteEntryList(input *DescribeRouteEntryListInput) (*DescribeRouteEntryListOutput, error) {
	req, out := c.DescribeRouteEntryListRequest(input)
	return out, req.Send()
}

// DescribeRouteEntryListWithContext is the same as DescribeRouteEntryList with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeRouteEntryList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeRouteEntryListWithContext(ctx volcengine.Context, input *DescribeRouteEntryListInput, opts ...request.Option) (*DescribeRouteEntryListOutput, error) {
	req, out := c.DescribeRouteEntryListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeRouteEntryListInput struct {
	_ struct{} `type:"structure"`

	DestinationCidrBlock *string `type:"string"`

	NextHopId *string `type:"string"`

	NextHopType *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	RouteEntryId *string `type:"string"`

	RouteEntryName *string `type:"string"`

	RouteEntryType *string `type:"string"`

	RouteTableId *string `type:"string"`
}

// String returns the string representation
func (s DescribeRouteEntryListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRouteEntryListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeRouteEntryListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeRouteEntryListInput"}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *DescribeRouteEntryListInput) SetDestinationCidrBlock(v string) *DescribeRouteEntryListInput {
	s.DestinationCidrBlock = &v
	return s
}

// SetNextHopId sets the NextHopId field's value.
func (s *DescribeRouteEntryListInput) SetNextHopId(v string) *DescribeRouteEntryListInput {
	s.NextHopId = &v
	return s
}

// SetNextHopType sets the NextHopType field's value.
func (s *DescribeRouteEntryListInput) SetNextHopType(v string) *DescribeRouteEntryListInput {
	s.NextHopType = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeRouteEntryListInput) SetPageNumber(v int64) *DescribeRouteEntryListInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeRouteEntryListInput) SetPageSize(v int64) *DescribeRouteEntryListInput {
	s.PageSize = &v
	return s
}

// SetRouteEntryId sets the RouteEntryId field's value.
func (s *DescribeRouteEntryListInput) SetRouteEntryId(v string) *DescribeRouteEntryListInput {
	s.RouteEntryId = &v
	return s
}

// SetRouteEntryName sets the RouteEntryName field's value.
func (s *DescribeRouteEntryListInput) SetRouteEntryName(v string) *DescribeRouteEntryListInput {
	s.RouteEntryName = &v
	return s
}

// SetRouteEntryType sets the RouteEntryType field's value.
func (s *DescribeRouteEntryListInput) SetRouteEntryType(v string) *DescribeRouteEntryListInput {
	s.RouteEntryType = &v
	return s
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *DescribeRouteEntryListInput) SetRouteTableId(v string) *DescribeRouteEntryListInput {
	s.RouteTableId = &v
	return s
}

type DescribeRouteEntryListOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	RouteEntries []*RouteEntryForDescribeRouteEntryListOutput `type:"list"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeRouteEntryListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeRouteEntryListOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeRouteEntryListOutput) SetPageNumber(v int64) *DescribeRouteEntryListOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeRouteEntryListOutput) SetPageSize(v int64) *DescribeRouteEntryListOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeRouteEntryListOutput) SetRequestId(v string) *DescribeRouteEntryListOutput {
	s.RequestId = &v
	return s
}

// SetRouteEntries sets the RouteEntries field's value.
func (s *DescribeRouteEntryListOutput) SetRouteEntries(v []*RouteEntryForDescribeRouteEntryListOutput) *DescribeRouteEntryListOutput {
	s.RouteEntries = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeRouteEntryListOutput) SetTotalCount(v int64) *DescribeRouteEntryListOutput {
	s.TotalCount = &v
	return s
}

type RouteEntryForDescribeRouteEntryListOutput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DestinationCidrBlock *string `type:"string"`

	NextHopId *string `type:"string"`

	NextHopName *string `type:"string"`

	NextHopType *string `type:"string"`

	RouteEntryId *string `type:"string"`

	RouteEntryName *string `type:"string"`

	RouteTableId *string `type:"string"`

	Status *string `type:"string"`

	Type *string `type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s RouteEntryForDescribeRouteEntryListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RouteEntryForDescribeRouteEntryListOutput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetDescription(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.Description = &v
	return s
}

// SetDestinationCidrBlock sets the DestinationCidrBlock field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetDestinationCidrBlock(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.DestinationCidrBlock = &v
	return s
}

// SetNextHopId sets the NextHopId field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetNextHopId(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.NextHopId = &v
	return s
}

// SetNextHopName sets the NextHopName field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetNextHopName(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.NextHopName = &v
	return s
}

// SetNextHopType sets the NextHopType field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetNextHopType(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.NextHopType = &v
	return s
}

// SetRouteEntryId sets the RouteEntryId field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetRouteEntryId(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.RouteEntryId = &v
	return s
}

// SetRouteEntryName sets the RouteEntryName field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetRouteEntryName(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.RouteEntryName = &v
	return s
}

// SetRouteTableId sets the RouteTableId field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetRouteTableId(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.RouteTableId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetStatus(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetType(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.Type = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *RouteEntryForDescribeRouteEntryListOutput) SetVpcId(v string) *RouteEntryForDescribeRouteEntryListOutput {
	s.VpcId = &v
	return s
}
