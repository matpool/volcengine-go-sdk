// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package cen

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeCenBandwidthPackagesCommon = "DescribeCenBandwidthPackages"

// DescribeCenBandwidthPackagesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenBandwidthPackagesCommon operation. The "output" return
// value will be populated with the DescribeCenBandwidthPackagesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenBandwidthPackagesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenBandwidthPackagesCommon Send returns without error.
//
// See DescribeCenBandwidthPackagesCommon for more information on using the DescribeCenBandwidthPackagesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenBandwidthPackagesCommonRequest method.
//    req, resp := client.DescribeCenBandwidthPackagesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenBandwidthPackagesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeCenBandwidthPackagesCommon,
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

// DescribeCenBandwidthPackagesCommon API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation DescribeCenBandwidthPackagesCommon for usage and error information.
func (c *CEN) DescribeCenBandwidthPackagesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeCenBandwidthPackagesCommonRequest(input)
	return out, req.Send()
}

// DescribeCenBandwidthPackagesCommonWithContext is the same as DescribeCenBandwidthPackagesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenBandwidthPackagesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenBandwidthPackagesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeCenBandwidthPackagesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeCenBandwidthPackages = "DescribeCenBandwidthPackages"

// DescribeCenBandwidthPackagesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeCenBandwidthPackages operation. The "output" return
// value will be populated with the DescribeCenBandwidthPackagesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeCenBandwidthPackagesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeCenBandwidthPackagesCommon Send returns without error.
//
// See DescribeCenBandwidthPackages for more information on using the DescribeCenBandwidthPackages
// API call, and error handling.
//
//    // Example sending a request using the DescribeCenBandwidthPackagesRequest method.
//    req, resp := client.DescribeCenBandwidthPackagesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *CEN) DescribeCenBandwidthPackagesRequest(input *DescribeCenBandwidthPackagesInput) (req *request.Request, output *DescribeCenBandwidthPackagesOutput) {
	op := &request.Operation{
		Name:       opDescribeCenBandwidthPackages,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeCenBandwidthPackagesInput{}
	}

	output = &DescribeCenBandwidthPackagesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeCenBandwidthPackages API operation for CEN.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for CEN's
// API operation DescribeCenBandwidthPackages for usage and error information.
func (c *CEN) DescribeCenBandwidthPackages(input *DescribeCenBandwidthPackagesInput) (*DescribeCenBandwidthPackagesOutput, error) {
	req, out := c.DescribeCenBandwidthPackagesRequest(input)
	return out, req.Send()
}

// DescribeCenBandwidthPackagesWithContext is the same as DescribeCenBandwidthPackages with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeCenBandwidthPackages for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CEN) DescribeCenBandwidthPackagesWithContext(ctx volcengine.Context, input *DescribeCenBandwidthPackagesInput, opts ...request.Option) (*DescribeCenBandwidthPackagesOutput, error) {
	req, out := c.DescribeCenBandwidthPackagesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CenBandwidthPackageForDescribeCenBandwidthPackagesOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	CenBandwidthPackageId *string `type:"string"`

	CenBandwidthPackageName *string `type:"string"`

	CenIds []*string `type:"list"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	ExpiredTime *string `type:"string"`

	LocalGeographicRegionSetId *string `type:"string"`

	PeerGeographicRegionSetId *string `type:"string"`

	RemainingBandwidth *int64 `type:"integer"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetAccountId(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.AccountId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetBandwidth(v int64) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetBillingType(v int64) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetBusinessStatus(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.BusinessStatus = &v
	return s
}

// SetCenBandwidthPackageId sets the CenBandwidthPackageId field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetCenBandwidthPackageId(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.CenBandwidthPackageId = &v
	return s
}

// SetCenBandwidthPackageName sets the CenBandwidthPackageName field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetCenBandwidthPackageName(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.CenBandwidthPackageName = &v
	return s
}

// SetCenIds sets the CenIds field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetCenIds(v []*string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.CenIds = v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetCreationTime(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetDeletedTime(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetDescription(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.Description = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetExpiredTime(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.ExpiredTime = &v
	return s
}

// SetLocalGeographicRegionSetId sets the LocalGeographicRegionSetId field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetLocalGeographicRegionSetId(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.LocalGeographicRegionSetId = &v
	return s
}

// SetPeerGeographicRegionSetId sets the PeerGeographicRegionSetId field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetPeerGeographicRegionSetId(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.PeerGeographicRegionSetId = &v
	return s
}

// SetRemainingBandwidth sets the RemainingBandwidth field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetRemainingBandwidth(v int64) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.RemainingBandwidth = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetStatus(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) SetUpdateTime(v string) *CenBandwidthPackageForDescribeCenBandwidthPackagesOutput {
	s.UpdateTime = &v
	return s
}

type DescribeCenBandwidthPackagesInput struct {
	_ struct{} `type:"structure"`

	CenBandwidthPackageIds []*string `type:"list"`

	CenBandwidthPackageName *string `type:"string"`

	CenId *string `type:"string"`

	LocalGeographicRegionSetId *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	PeerGeographicRegionSetId *string `type:"string"`
}

// String returns the string representation
func (s DescribeCenBandwidthPackagesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenBandwidthPackagesInput) GoString() string {
	return s.String()
}

// SetCenBandwidthPackageIds sets the CenBandwidthPackageIds field's value.
func (s *DescribeCenBandwidthPackagesInput) SetCenBandwidthPackageIds(v []*string) *DescribeCenBandwidthPackagesInput {
	s.CenBandwidthPackageIds = v
	return s
}

// SetCenBandwidthPackageName sets the CenBandwidthPackageName field's value.
func (s *DescribeCenBandwidthPackagesInput) SetCenBandwidthPackageName(v string) *DescribeCenBandwidthPackagesInput {
	s.CenBandwidthPackageName = &v
	return s
}

// SetCenId sets the CenId field's value.
func (s *DescribeCenBandwidthPackagesInput) SetCenId(v string) *DescribeCenBandwidthPackagesInput {
	s.CenId = &v
	return s
}

// SetLocalGeographicRegionSetId sets the LocalGeographicRegionSetId field's value.
func (s *DescribeCenBandwidthPackagesInput) SetLocalGeographicRegionSetId(v string) *DescribeCenBandwidthPackagesInput {
	s.LocalGeographicRegionSetId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenBandwidthPackagesInput) SetPageNumber(v int64) *DescribeCenBandwidthPackagesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenBandwidthPackagesInput) SetPageSize(v int64) *DescribeCenBandwidthPackagesInput {
	s.PageSize = &v
	return s
}

// SetPeerGeographicRegionSetId sets the PeerGeographicRegionSetId field's value.
func (s *DescribeCenBandwidthPackagesInput) SetPeerGeographicRegionSetId(v string) *DescribeCenBandwidthPackagesInput {
	s.PeerGeographicRegionSetId = &v
	return s
}

type DescribeCenBandwidthPackagesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CenBandwidthPackages []*CenBandwidthPackageForDescribeCenBandwidthPackagesOutput `type:"list"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeCenBandwidthPackagesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeCenBandwidthPackagesOutput) GoString() string {
	return s.String()
}

// SetCenBandwidthPackages sets the CenBandwidthPackages field's value.
func (s *DescribeCenBandwidthPackagesOutput) SetCenBandwidthPackages(v []*CenBandwidthPackageForDescribeCenBandwidthPackagesOutput) *DescribeCenBandwidthPackagesOutput {
	s.CenBandwidthPackages = v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeCenBandwidthPackagesOutput) SetPageNumber(v int64) *DescribeCenBandwidthPackagesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeCenBandwidthPackagesOutput) SetPageSize(v int64) *DescribeCenBandwidthPackagesOutput {
	s.PageSize = &v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeCenBandwidthPackagesOutput) SetTotalCount(v int64) *DescribeCenBandwidthPackagesOutput {
	s.TotalCount = &v
	return s
}
