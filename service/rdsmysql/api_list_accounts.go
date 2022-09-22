// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListAccountsCommon = "ListAccounts"

// ListAccountsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAccountsCommon operation. The "output" return
// value will be populated with the ListAccountsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAccountsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAccountsCommon Send returns without error.
//
// See ListAccountsCommon for more information on using the ListAccountsCommon
// API call, and error handling.
//
//    // Example sending a request using the ListAccountsCommonRequest method.
//    req, resp := client.ListAccountsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListAccountsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListAccountsCommon,
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

// ListAccountsCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListAccountsCommon for usage and error information.
func (c *RDSMYSQL) ListAccountsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListAccountsCommonRequest(input)
	return out, req.Send()
}

// ListAccountsCommonWithContext is the same as ListAccountsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListAccountsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListAccountsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListAccountsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListAccounts = "ListAccounts"

// ListAccountsRequest generates a "volcengine/request.Request" representing the
// client's request for the ListAccounts operation. The "output" return
// value will be populated with the ListAccountsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListAccountsCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListAccountsCommon Send returns without error.
//
// See ListAccounts for more information on using the ListAccounts
// API call, and error handling.
//
//    // Example sending a request using the ListAccountsRequest method.
//    req, resp := client.ListAccountsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListAccountsRequest(input *ListAccountsInput) (req *request.Request, output *ListAccountsOutput) {
	op := &request.Operation{
		Name:       opListAccounts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListAccountsInput{}
	}

	output = &ListAccountsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListAccounts API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListAccounts for usage and error information.
func (c *RDSMYSQL) ListAccounts(input *ListAccountsInput) (*ListAccountsOutput, error) {
	req, out := c.ListAccountsRequest(input)
	return out, req.Send()
}

// ListAccountsWithContext is the same as ListAccounts with the addition of
// the ability to pass a context and additional request options.
//
// See ListAccounts for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListAccountsWithContext(ctx volcengine.Context, input *ListAccountsInput, opts ...request.Option) (*ListAccountsOutput, error) {
	req, out := c.ListAccountsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListAccountsInput struct {
	_ struct{} `type:"structure"`

	AccountName *string `min:"2" max:"32" type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`
}

// String returns the string representation
func (s ListAccountsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAccountsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAccountsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListAccountsInput"}
	if s.AccountName != nil && len(*s.AccountName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("AccountName", 2))
	}
	if s.AccountName != nil && len(*s.AccountName) > 32 {
		invalidParams.Add(request.NewErrParamMaxLen("AccountName", 32, *s.AccountName))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountName sets the AccountName field's value.
func (s *ListAccountsInput) SetAccountName(v string) *ListAccountsInput {
	s.AccountName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ListAccountsInput) SetInstanceId(v string) *ListAccountsInput {
	s.InstanceId = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListAccountsInput) SetLimit(v int32) *ListAccountsInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListAccountsInput) SetOffset(v int32) *ListAccountsInput {
	s.Offset = &v
	return s
}

type ListAccountsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListAccountsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListAccountsOutput) GoString() string {
	return s.String()
}

// SetTotal sets the Total field's value.
func (s *ListAccountsOutput) SetTotal(v int32) *ListAccountsOutput {
	s.Total = &v
	return s
}
