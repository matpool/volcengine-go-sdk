// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package storageebs

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeTagsCommon = "DescribeTags"

// DescribeTagsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTagsCommon operation. The "output" return
// value will be populated with the DescribeTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTagsCommon Send returns without error.
//
// See DescribeTagsCommon for more information on using the DescribeTagsCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeTagsCommonRequest method.
//	req, resp := client.DescribeTagsCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *STORAGEEBS) DescribeTagsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeTagsCommon,
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

// DescribeTagsCommon API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation DescribeTagsCommon for usage and error information.
func (c *STORAGEEBS) DescribeTagsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeTagsCommonRequest(input)
	return out, req.Send()
}

// DescribeTagsCommonWithContext is the same as DescribeTagsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTagsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DescribeTagsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeTagsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeTags = "DescribeTags"

// DescribeTagsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeTags operation. The "output" return
// value will be populated with the DescribeTagsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeTagsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeTagsCommon Send returns without error.
//
// See DescribeTags for more information on using the DescribeTags
// API call, and error handling.
//
//	// Example sending a request using the DescribeTagsRequest method.
//	req, resp := client.DescribeTagsRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *STORAGEEBS) DescribeTagsRequest(input *DescribeTagsInput) (req *request.Request, output *DescribeTagsOutput) {
	op := &request.Operation{
		Name:       opDescribeTags,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeTagsInput{}
	}

	output = &DescribeTagsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeTags API operation for STORAGE_EBS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for STORAGE_EBS's
// API operation DescribeTags for usage and error information.
func (c *STORAGEEBS) DescribeTags(input *DescribeTagsInput) (*DescribeTagsOutput, error) {
	req, out := c.DescribeTagsRequest(input)
	return out, req.Send()
}

// DescribeTagsWithContext is the same as DescribeTags with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeTags for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *STORAGEEBS) DescribeTagsWithContext(ctx volcengine.Context, input *DescribeTagsInput, opts ...request.Option) (*DescribeTagsOutput, error) {
	req, out := c.DescribeTagsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeTagsInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	ResourceIds []*string `type:"list"`

	ResourceType *string `type:"string"`

	SysTagVisible *bool `type:"boolean"`

	TagFilters []*TagFilterForDescribeTagsInput `type:"list"`
}

// String returns the string representation
func (s DescribeTagsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTagsInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTagsInput) SetPageNumber(v int32) *DescribeTagsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTagsInput) SetPageSize(v int32) *DescribeTagsInput {
	s.PageSize = &v
	return s
}

// SetResourceIds sets the ResourceIds field's value.
func (s *DescribeTagsInput) SetResourceIds(v []*string) *DescribeTagsInput {
	s.ResourceIds = v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *DescribeTagsInput) SetResourceType(v string) *DescribeTagsInput {
	s.ResourceType = &v
	return s
}

// SetSysTagVisible sets the SysTagVisible field's value.
func (s *DescribeTagsInput) SetSysTagVisible(v bool) *DescribeTagsInput {
	s.SysTagVisible = &v
	return s
}

// SetTagFilters sets the TagFilters field's value.
func (s *DescribeTagsInput) SetTagFilters(v []*TagFilterForDescribeTagsInput) *DescribeTagsInput {
	s.TagFilters = v
	return s
}

type DescribeTagsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	TagResources []*TagResourceForDescribeTagsOutput `type:"list"`

	TotalCount *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeTagsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeTagsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeTagsOutput) SetPageNumber(v int32) *DescribeTagsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeTagsOutput) SetPageSize(v int32) *DescribeTagsOutput {
	s.PageSize = &v
	return s
}

// SetTagResources sets the TagResources field's value.
func (s *DescribeTagsOutput) SetTagResources(v []*TagResourceForDescribeTagsOutput) *DescribeTagsOutput {
	s.TagResources = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeTagsOutput) SetTotalCount(v int32) *DescribeTagsOutput {
	s.TotalCount = &v
	return s
}

type TagFilterForDescribeTagsInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Values []*string `type:"list"`
}

// String returns the string representation
func (s TagFilterForDescribeTagsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagFilterForDescribeTagsInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagFilterForDescribeTagsInput) SetKey(v string) *TagFilterForDescribeTagsInput {
	s.Key = &v
	return s
}

// SetValues sets the Values field's value.
func (s *TagFilterForDescribeTagsInput) SetValues(v []*string) *TagFilterForDescribeTagsInput {
	s.Values = v
	return s
}

type TagResourceForDescribeTagsOutput struct {
	_ struct{} `type:"structure"`

	ResourceId *string `type:"string"`

	ResourceType *string `type:"string"`

	TagKey *string `type:"string"`

	TagValue *string `type:"string"`
}

// String returns the string representation
func (s TagResourceForDescribeTagsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagResourceForDescribeTagsOutput) GoString() string {
	return s.String()
}

// SetResourceId sets the ResourceId field's value.
func (s *TagResourceForDescribeTagsOutput) SetResourceId(v string) *TagResourceForDescribeTagsOutput {
	s.ResourceId = &v
	return s
}

// SetResourceType sets the ResourceType field's value.
func (s *TagResourceForDescribeTagsOutput) SetResourceType(v string) *TagResourceForDescribeTagsOutput {
	s.ResourceType = &v
	return s
}

// SetTagKey sets the TagKey field's value.
func (s *TagResourceForDescribeTagsOutput) SetTagKey(v string) *TagResourceForDescribeTagsOutput {
	s.TagKey = &v
	return s
}

// SetTagValue sets the TagValue field's value.
func (s *TagResourceForDescribeTagsOutput) SetTagValue(v string) *TagResourceForDescribeTagsOutput {
	s.TagValue = &v
	return s
}
