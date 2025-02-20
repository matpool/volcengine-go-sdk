// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListRolesCommon = "ListRoles"

// ListRolesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRolesCommon operation. The "output" return
// value will be populated with the ListRolesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRolesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRolesCommon Send returns without error.
//
// See ListRolesCommon for more information on using the ListRolesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListRolesCommonRequest method.
//    req, resp := client.ListRolesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListRolesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListRolesCommon,
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

// ListRolesCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListRolesCommon for usage and error information.
func (c *IAM) ListRolesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListRolesCommonRequest(input)
	return out, req.Send()
}

// ListRolesCommonWithContext is the same as ListRolesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListRolesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListRolesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListRolesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListRoles = "ListRoles"

// ListRolesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListRoles operation. The "output" return
// value will be populated with the ListRolesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListRolesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListRolesCommon Send returns without error.
//
// See ListRoles for more information on using the ListRoles
// API call, and error handling.
//
//    // Example sending a request using the ListRolesRequest method.
//    req, resp := client.ListRolesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) ListRolesRequest(input *ListRolesInput) (req *request.Request, output *ListRolesOutput) {
	op := &request.Operation{
		Name:       opListRoles,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListRolesInput{}
	}

	output = &ListRolesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ListRoles API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation ListRoles for usage and error information.
func (c *IAM) ListRoles(input *ListRolesInput) (*ListRolesOutput, error) {
	req, out := c.ListRolesRequest(input)
	return out, req.Send()
}

// ListRolesWithContext is the same as ListRoles with the addition of
// the ability to pass a context and additional request options.
//
// See ListRoles for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) ListRolesWithContext(ctx volcengine.Context, input *ListRolesInput, opts ...request.Option) (*ListRolesOutput, error) {
	req, out := c.ListRolesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListRolesInput struct {
	_ struct{} `type:"structure"`

	Limit *int64 `type:"integer"`

	Offset *int64 `type:"integer"`

	Query *string `type:"string"`
}

// String returns the string representation
func (s ListRolesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRolesInput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListRolesInput) SetLimit(v int64) *ListRolesInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListRolesInput) SetOffset(v int64) *ListRolesInput {
	s.Offset = &v
	return s
}

// SetQuery sets the Query field's value.
func (s *ListRolesInput) SetQuery(v string) *ListRolesInput {
	s.Query = &v
	return s
}

type ListRolesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`

	RoleMetadata []*RoleMetadataForListRolesOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListRolesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListRolesOutput) GoString() string {
	return s.String()
}

// SetLimit sets the Limit field's value.
func (s *ListRolesOutput) SetLimit(v int32) *ListRolesOutput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListRolesOutput) SetOffset(v int32) *ListRolesOutput {
	s.Offset = &v
	return s
}

// SetRoleMetadata sets the RoleMetadata field's value.
func (s *ListRolesOutput) SetRoleMetadata(v []*RoleMetadataForListRolesOutput) *ListRolesOutput {
	s.RoleMetadata = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListRolesOutput) SetTotal(v int32) *ListRolesOutput {
	s.Total = &v
	return s
}

type RoleMetadataForListRolesOutput struct {
	_ struct{} `type:"structure"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	RoleId *int32 `type:"int32"`

	RoleName *string `type:"string"`

	Trn *string `type:"string"`

	TrustPolicyDocument *string `type:"string"`
}

// String returns the string representation
func (s RoleMetadataForListRolesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RoleMetadataForListRolesOutput) GoString() string {
	return s.String()
}

// SetCreateDate sets the CreateDate field's value.
func (s *RoleMetadataForListRolesOutput) SetCreateDate(v string) *RoleMetadataForListRolesOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *RoleMetadataForListRolesOutput) SetDescription(v string) *RoleMetadataForListRolesOutput {
	s.Description = &v
	return s
}

// SetRoleId sets the RoleId field's value.
func (s *RoleMetadataForListRolesOutput) SetRoleId(v int32) *RoleMetadataForListRolesOutput {
	s.RoleId = &v
	return s
}

// SetRoleName sets the RoleName field's value.
func (s *RoleMetadataForListRolesOutput) SetRoleName(v string) *RoleMetadataForListRolesOutput {
	s.RoleName = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *RoleMetadataForListRolesOutput) SetTrn(v string) *RoleMetadataForListRolesOutput {
	s.Trn = &v
	return s
}

// SetTrustPolicyDocument sets the TrustPolicyDocument field's value.
func (s *RoleMetadataForListRolesOutput) SetTrustPolicyDocument(v string) *RoleMetadataForListRolesOutput {
	s.TrustPolicyDocument = &v
	return s
}
