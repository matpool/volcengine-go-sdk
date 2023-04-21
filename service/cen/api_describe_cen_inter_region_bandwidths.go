// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCenInterRegionBandwidthsCommon = "DescribeCenInterRegionBandwidths"

// DescribeCenInterRegionBandwidthsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenInterRegionBandwidthsCommon operation. The "output" return
// value will be populated with the DescribeCenInterRegionBandwidthsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenInterRegionBandwidthsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenInterRegionBandwidthsCommon Send returns without error.
//
// See DescribeCenInterRegionBandwidthsCommon for more information on using the DescribeCenInterRegionBandwidthsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenInterRegionBandwidthsCommonRequest method.
//    req, resp := client.DescribeCenInterRegionBandwidthsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenInterRegionBandwidthsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCenInterRegionBandwidthsCommon,
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

// DescribeCenInterRegionBandwidthsCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCenInterRegionBandwidthsCommon for usage and error information.
func (c *CEN) DescribeCenInterRegionBandwidthsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCenInterRegionBandwidthsCommonRequest(input)
	return out, req.Send()
}

// DescribeCenInterRegionBandwidthsCommonWithContext is the same as DescribeCenInterRegionBandwidthsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenInterRegionBandwidthsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenInterRegionBandwidthsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCenInterRegionBandwidthsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCenInterRegionBandwidths = "DescribeCenInterRegionBandwidths"

// DescribeCenInterRegionBandwidthsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenInterRegionBandwidths operation. The "output" return
// value will be populated with the DescribeCenInterRegionBandwidthsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenInterRegionBandwidthsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenInterRegionBandwidthsCommon Send returns without error.
//
// See DescribeCenInterRegionBandwidths for more information on using the DescribeCenInterRegionBandwidths
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenInterRegionBandwidthsRequest method.
//    req, resp := client.DescribeCenInterRegionBandwidthsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenInterRegionBandwidthsRequest(input *DescribeCenInterRegionBandwidthsInput) (req *request.Request, output *DescribeCenInterRegionBandwidthsOutput) {
	op := &request.Operation{
		Name:       opDescribeCenInterRegionBandwidths,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCenInterRegionBandwidthsInput{}
	}

	output = &DescribeCenInterRegionBandwidthsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCenInterRegionBandwidths API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CEN's
// API operation DescribeCenInterRegionBandwidths for usage and error information.
func (c *CEN) DescribeCenInterRegionBandwidths(input *DescribeCenInterRegionBandwidthsInput) (*DescribeCenInterRegionBandwidthsOutput, error) {
	req, out := c.DescribeCenInterRegionBandwidthsRequest(input)
	return out, req.Send()
}

// DescribeCenInterRegionBandwidthsWithContext is the same as DescribeCenInterRegionBandwidths with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenInterRegionBandwidths for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenInterRegionBandwidthsWithContext(ctx volcengine.Context, input *DescribeCenInterRegionBandwidthsInput, opts ...request.Option) (*DescribeCenInterRegionBandwidthsOutput, error) {
	req, out := c.DescribeCenInterRegionBandwidthsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeCenInterRegionBandwidthsInput struct {
	_ struct{} `type:"structure"`

	InterRegionBandwidthIds []*string `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *string `type:"string"`
}

// String returns the string representation
func (s DescribeCenInterRegionBandwidthsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenInterRegionBandwidthsInput) GoString() string {
	return s.String()
}

// SetInterRegionBandwidthIds sets the InterRegionBandwidthIds field's value.
func (s *DescribeCenInterRegionBandwidthsInput) SetInterRegionBandwidthIds(v []*string) *DescribeCenInterRegionBandwidthsInput {
	s.InterRegionBandwidthIds = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenInterRegionBandwidthsInput) SetPageNumber(v int64) *DescribeCenInterRegionBandwidthsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenInterRegionBandwidthsInput) SetPageSize(v string) *DescribeCenInterRegionBandwidthsInput {
	s.PageSize = &v
	return s
}

type DescribeCenInterRegionBandwidthsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InterRegionBandwidths []*InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCenInterRegionBandwidthsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenInterRegionBandwidthsOutput) GoString() string {
	return s.String()
}

// SetInterRegionBandwidths sets the InterRegionBandwidths field's value.
func (s *DescribeCenInterRegionBandwidthsOutput) SetInterRegionBandwidths(v []*InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) *DescribeCenInterRegionBandwidthsOutput {
	s.InterRegionBandwidths = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenInterRegionBandwidthsOutput) SetPageNumber(v int64) *DescribeCenInterRegionBandwidthsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenInterRegionBandwidthsOutput) SetPageSize(v int64) *DescribeCenInterRegionBandwidthsOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCenInterRegionBandwidthsOutput) SetTotalCount(v int64) *DescribeCenInterRegionBandwidthsOutput {
	s.TotalCount = &v
	return s
}

type InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput struct {
	_ struct{} `type:"structure"`

	Bandwidth *int64 `type:"integer"`

	CenBandwidthPackageId *string `type:"string"`

	CenId *string `type:"string"`

	CreationTime *string `type:"string"`

	InterRegionBandwidthId *string `type:"string"`

	LocalRegionId *string `type:"string"`

	PeerRegionId *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) GoString() string {
	return s.String()
}

// SetBandwidth sets the Bandwidth field's value.
func (s *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) SetBandwidth(v int64) *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput {
	s.Bandwidth = &v
	return s
}

// SetCenBandwidthPackageId sets the CenBandwidthPackageId field's value.
func (s *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) SetCenBandwidthPackageId(v string) *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput {
	s.CenBandwidthPackageId = &v
	return s
}

// SetCenId sets the CenId field's value.
func (s *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) SetCenId(v string) *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput {
	s.CenId = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) SetCreationTime(v string) *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput {
	s.CreationTime = &v
	return s
}

// SetInterRegionBandwidthId sets the InterRegionBandwidthId field's value.
func (s *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) SetInterRegionBandwidthId(v string) *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput {
	s.InterRegionBandwidthId = &v
	return s
}

// SetLocalRegionId sets the LocalRegionId field's value.
func (s *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) SetLocalRegionId(v string) *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput {
	s.LocalRegionId = &v
	return s
}

// SetPeerRegionId sets the PeerRegionId field's value.
func (s *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) SetPeerRegionId(v string) *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput {
	s.PeerRegionId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) SetStatus(v string) *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput) SetUpdateTime(v string) *InterRegionBandwidthForDescribeCenInterRegionBandwidthsOutput {
	s.UpdateTime = &v
	return s
}
