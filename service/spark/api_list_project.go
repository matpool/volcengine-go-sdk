// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package spark

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListProjectCommon = "listProject"

// ListProjectCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListProjectCommon operation. The "output" return
// value will be populated with the ListProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListProjectCommon Send returns without error.
//
// See ListProjectCommon for more information on using the ListProjectCommon
// API call, and error handling.
//
//	// Example sending a request using the ListProjectCommonRequest method.
//	req, resp := client.ListProjectCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) ListProjectCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListProjectCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListProjectCommon API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation ListProjectCommon for usage and error information.
func (c *SPARK) ListProjectCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListProjectCommonRequest(input)
	return out, req.Send()
}

// ListProjectCommonWithContext is the same as ListProjectCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListProjectCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) ListProjectCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListProjectCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListProject = "listProject"

// ListProjectRequest generates a "volcengine/request.Request" representing the
// client's request for the ListProject operation. The "output" return
// value will be populated with the ListProjectCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListProjectCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListProjectCommon Send returns without error.
//
// See ListProject for more information on using the ListProject
// API call, and error handling.
//
//	// Example sending a request using the ListProjectRequest method.
//	req, resp := client.ListProjectRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *SPARK) ListProjectRequest(input *ListProjectInput) (req *request.Request, output *ListProjectOutput) {
	op := &request.Operation{
		Name:       opListProject,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListProjectInput{}
	}

	output = &ListProjectOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListProject API operation for SPARK.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for SPARK's
// API operation ListProject for usage and error information.
func (c *SPARK) ListProject(input *ListProjectInput) (*ListProjectOutput, error) {
	req, out := c.ListProjectRequest(input)
	return out, req.Send()
}

// ListProjectWithContext is the same as ListProject with the addition of
// the ability to pass a context and additional request options.
//
// See ListProject for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SPARK) ListProjectWithContext(ctx volcengine.Context, input *ListProjectInput, opts ...request.Option) (*ListProjectOutput, error) {
	req, out := c.ListProjectRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ItemForlistProjectOutput struct {
	_ struct{} `type:"structure"`

	AuthorityType *string `type:"string"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	OwnerId *string `type:"string"`

	OwnerName *string `type:"string"`

	ProjectId *string `type:"string"`

	ProjectName *string `type:"string"`

	RegionId *string `type:"string"`
}

// String returns the string representation
func (s ItemForlistProjectOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ItemForlistProjectOutput) GoString() string {
	return s.String()
}

// SetAuthorityType sets the AuthorityType field's value.
func (s *ItemForlistProjectOutput) SetAuthorityType(v string) *ItemForlistProjectOutput {
	s.AuthorityType = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *ItemForlistProjectOutput) SetCreateDate(v string) *ItemForlistProjectOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ItemForlistProjectOutput) SetDescription(v string) *ItemForlistProjectOutput {
	s.Description = &v
	return s
}

// SetOwnerId sets the OwnerId field's value.
func (s *ItemForlistProjectOutput) SetOwnerId(v string) *ItemForlistProjectOutput {
	s.OwnerId = &v
	return s
}

// SetOwnerName sets the OwnerName field's value.
func (s *ItemForlistProjectOutput) SetOwnerName(v string) *ItemForlistProjectOutput {
	s.OwnerName = &v
	return s
}

// SetProjectId sets the ProjectId field's value.
func (s *ItemForlistProjectOutput) SetProjectId(v string) *ItemForlistProjectOutput {
	s.ProjectId = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ItemForlistProjectOutput) SetProjectName(v string) *ItemForlistProjectOutput {
	s.ProjectName = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *ItemForlistProjectOutput) SetRegionId(v string) *ItemForlistProjectOutput {
	s.RegionId = &v
	return s
}

type ListProjectInput struct {
	_ struct{} `type:"structure"`

	PageNum *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	SearchKey *string `type:"string"`
}

// String returns the string representation
func (s ListProjectInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListProjectInput) GoString() string {
	return s.String()
}

// SetPageNum sets the PageNum field's value.
func (s *ListProjectInput) SetPageNum(v int32) *ListProjectInput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListProjectInput) SetPageSize(v int32) *ListProjectInput {
	s.PageSize = &v
	return s
}

// SetSearchKey sets the SearchKey field's value.
func (s *ListProjectInput) SetSearchKey(v string) *ListProjectInput {
	s.SearchKey = &v
	return s
}

type ListProjectOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Items []*ItemForlistProjectOutput `type:"list"`

	PageNum *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	Pages *int32 `type:"int32"`

	Total *int64 `type:"int64"`
}

// String returns the string representation
func (s ListProjectOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListProjectOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListProjectOutput) SetItems(v []*ItemForlistProjectOutput) *ListProjectOutput {
	s.Items = v
	return s
}

// SetPageNum sets the PageNum field's value.
func (s *ListProjectOutput) SetPageNum(v int32) *ListProjectOutput {
	s.PageNum = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *ListProjectOutput) SetPageSize(v int32) *ListProjectOutput {
	s.PageSize = &v
	return s
}

// SetPages sets the Pages field's value.
func (s *ListProjectOutput) SetPages(v int32) *ListProjectOutput {
	s.Pages = &v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListProjectOutput) SetTotal(v int64) *ListProjectOutput {
	s.Total = &v
	return s
}
