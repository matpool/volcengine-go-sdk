// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package natgateway

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSnatEntriesCommon = "DescribeSnatEntries"

// DescribeSnatEntriesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSnatEntriesCommon operation. The "output" return
// value will be populated with the DescribeSnatEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSnatEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSnatEntriesCommon Send returns without error.
//
// See DescribeSnatEntriesCommon for more information on using the DescribeSnatEntriesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSnatEntriesCommonRequest method.
//    req, resp := client.DescribeSnatEntriesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) DescribeSnatEntriesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSnatEntriesCommon,
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

// DescribeSnatEntriesCommon API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for NATGATEWAY's
// API operation DescribeSnatEntriesCommon for usage and error information.
func (c *NATGATEWAY) DescribeSnatEntriesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSnatEntriesCommonRequest(input)
	return out, req.Send()
}

// DescribeSnatEntriesCommonWithContext is the same as DescribeSnatEntriesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSnatEntriesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) DescribeSnatEntriesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSnatEntriesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSnatEntries = "DescribeSnatEntries"

// DescribeSnatEntriesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSnatEntries operation. The "output" return
// value will be populated with the DescribeSnatEntriesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSnatEntriesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSnatEntriesCommon Send returns without error.
//
// See DescribeSnatEntries for more information on using the DescribeSnatEntries
// API call, and error handling.
//
//    // Example sending a request using the DescribeSnatEntriesRequest method.
//    req, resp := client.DescribeSnatEntriesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *NATGATEWAY) DescribeSnatEntriesRequest(input *DescribeSnatEntriesInput) (req *request.Request, output *DescribeSnatEntriesOutput) {
	op := &request.Operation{
		Name:       opDescribeSnatEntries,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSnatEntriesInput{}
	}

	output = &DescribeSnatEntriesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSnatEntries API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for NATGATEWAY's
// API operation DescribeSnatEntries for usage and error information.
func (c *NATGATEWAY) DescribeSnatEntries(input *DescribeSnatEntriesInput) (*DescribeSnatEntriesOutput, error) {
	req, out := c.DescribeSnatEntriesRequest(input)
	return out, req.Send()
}

// DescribeSnatEntriesWithContext is the same as DescribeSnatEntries with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSnatEntries for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) DescribeSnatEntriesWithContext(ctx volcengine.Context, input *DescribeSnatEntriesInput, opts ...request.Option) (*DescribeSnatEntriesOutput, error) {
	req, out := c.DescribeSnatEntriesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSnatEntriesInput struct {
	_ struct{} `type:"structure"`

	EipId *string `type:"string"`

	NatGatewayId *string `type:"string"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `max:"100" type:"integer"`

	SnatEntryIds []*string `type:"list"`

	SnatEntryName *string `type:"string"`

	SubnetId *string `type:"string"`
}

// String returns the string representation
func (s DescribeSnatEntriesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSnatEntriesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeSnatEntriesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeSnatEntriesInput"}
	if s.PageSize != nil && *s.PageSize > 100 {
		invalidParams.Add(request.NewErrParamMaxValue("PageSize", 100))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEipId sets the EipId field's value.
func (s *DescribeSnatEntriesInput) SetEipId(v string) *DescribeSnatEntriesInput {
	s.EipId = &v
	return s
}

// SetNatGatewayId sets the NatGatewayId field's value.
func (s *DescribeSnatEntriesInput) SetNatGatewayId(v string) *DescribeSnatEntriesInput {
	s.NatGatewayId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSnatEntriesInput) SetPageNumber(v int64) *DescribeSnatEntriesInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSnatEntriesInput) SetPageSize(v int64) *DescribeSnatEntriesInput {
	s.PageSize = &v
	return s
}

// SetSnatEntryIds sets the SnatEntryIds field's value.
func (s *DescribeSnatEntriesInput) SetSnatEntryIds(v []*string) *DescribeSnatEntriesInput {
	s.SnatEntryIds = v
	return s
}

// SetSnatEntryName sets the SnatEntryName field's value.
func (s *DescribeSnatEntriesInput) SetSnatEntryName(v string) *DescribeSnatEntriesInput {
	s.SnatEntryName = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *DescribeSnatEntriesInput) SetSubnetId(v string) *DescribeSnatEntriesInput {
	s.SubnetId = &v
	return s
}

type DescribeSnatEntriesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	SnatEntries []*SnatEntryForDescribeSnatEntriesOutput `type:"list"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeSnatEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSnatEntriesOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSnatEntriesOutput) SetPageNumber(v int64) *DescribeSnatEntriesOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSnatEntriesOutput) SetPageSize(v int64) *DescribeSnatEntriesOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeSnatEntriesOutput) SetRequestId(v string) *DescribeSnatEntriesOutput {
	s.RequestId = &v
	return s
}

// SetSnatEntries sets the SnatEntries field's value.
func (s *DescribeSnatEntriesOutput) SetSnatEntries(v []*SnatEntryForDescribeSnatEntriesOutput) *DescribeSnatEntriesOutput {
	s.SnatEntries = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeSnatEntriesOutput) SetTotalCount(v int64) *DescribeSnatEntriesOutput {
	s.TotalCount = &v
	return s
}

type SnatEntryForDescribeSnatEntriesOutput struct {
	_ struct{} `type:"structure"`

	EipAddress *string `type:"string"`

	EipId *string `type:"string"`

	NatGatewayId *string `type:"string"`

	SnatEntryId *string `type:"string"`

	SnatEntryName *string `type:"string"`

	Status *string `type:"string"`

	SubnetId *string `type:"string"`
}

// String returns the string representation
func (s SnatEntryForDescribeSnatEntriesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SnatEntryForDescribeSnatEntriesOutput) GoString() string {
	return s.String()
}

// SetEipAddress sets the EipAddress field's value.
func (s *SnatEntryForDescribeSnatEntriesOutput) SetEipAddress(v string) *SnatEntryForDescribeSnatEntriesOutput {
	s.EipAddress = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *SnatEntryForDescribeSnatEntriesOutput) SetEipId(v string) *SnatEntryForDescribeSnatEntriesOutput {
	s.EipId = &v
	return s
}

// SetNatGatewayId sets the NatGatewayId field's value.
func (s *SnatEntryForDescribeSnatEntriesOutput) SetNatGatewayId(v string) *SnatEntryForDescribeSnatEntriesOutput {
	s.NatGatewayId = &v
	return s
}

// SetSnatEntryId sets the SnatEntryId field's value.
func (s *SnatEntryForDescribeSnatEntriesOutput) SetSnatEntryId(v string) *SnatEntryForDescribeSnatEntriesOutput {
	s.SnatEntryId = &v
	return s
}

// SetSnatEntryName sets the SnatEntryName field's value.
func (s *SnatEntryForDescribeSnatEntriesOutput) SetSnatEntryName(v string) *SnatEntryForDescribeSnatEntriesOutput {
	s.SnatEntryName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *SnatEntryForDescribeSnatEntriesOutput) SetStatus(v string) *SnatEntryForDescribeSnatEntriesOutput {
	s.Status = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *SnatEntryForDescribeSnatEntriesOutput) SetSubnetId(v string) *SnatEntryForDescribeSnatEntriesOutput {
	s.SubnetId = &v
	return s
}
