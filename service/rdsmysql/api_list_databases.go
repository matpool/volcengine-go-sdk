// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListDatabasesCommon = "ListDatabases"

// ListDatabasesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListDatabasesCommon operation. The "output" return
// value will be populated with the ListDatabasesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListDatabasesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListDatabasesCommon Send returns without error.
//
// See ListDatabasesCommon for more information on using the ListDatabasesCommon
// API call, and error handling.
//
//    // Example sending a request using the ListDatabasesCommonRequest method.
//    req, resp := client.ListDatabasesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListDatabasesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListDatabasesCommon,
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

// ListDatabasesCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListDatabasesCommon for usage and error information.
func (c *RDSMYSQL) ListDatabasesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListDatabasesCommonRequest(input)
	return out, req.Send()
}

// ListDatabasesCommonWithContext is the same as ListDatabasesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListDatabasesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListDatabasesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListDatabasesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListDatabases = "ListDatabases"

// ListDatabasesRequest generates a "volcengine/request.Request" representing the
// client's request for the ListDatabases operation. The "output" return
// value will be populated with the ListDatabasesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListDatabasesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListDatabasesCommon Send returns without error.
//
// See ListDatabases for more information on using the ListDatabases
// API call, and error handling.
//
//    // Example sending a request using the ListDatabasesRequest method.
//    req, resp := client.ListDatabasesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListDatabasesRequest(input *ListDatabasesInput) (req *request.Request, output *ListDatabasesOutput) {
	op := &request.Operation{
		Name:       opListDatabases,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListDatabasesInput{}
	}

	output = &ListDatabasesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListDatabases API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListDatabases for usage and error information.
func (c *RDSMYSQL) ListDatabases(input *ListDatabasesInput) (*ListDatabasesOutput, error) {
	req, out := c.ListDatabasesRequest(input)
	return out, req.Send()
}

// ListDatabasesWithContext is the same as ListDatabases with the addition of
// the ability to pass a context and additional request options.
//
// See ListDatabases for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListDatabasesWithContext(ctx volcengine.Context, input *ListDatabasesInput, opts ...request.Option) (*ListDatabasesOutput, error) {
	req, out := c.ListDatabasesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ListDatabasesInput struct {
	_ struct{} `type:"structure"`

	DBStatus *string `type:"string" enum:"EnumOfDBStatusForListDatabasesInput"`

	InstanceId *string `type:"string"`

	Limit *int32 `type:"int32"`

	Offset *int32 `type:"int32"`
}

// String returns the string representation
func (s ListDatabasesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDatabasesInput) GoString() string {
	return s.String()
}

// SetDBStatus sets the DBStatus field's value.
func (s *ListDatabasesInput) SetDBStatus(v string) *ListDatabasesInput {
	s.DBStatus = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ListDatabasesInput) SetInstanceId(v string) *ListDatabasesInput {
	s.InstanceId = &v
	return s
}

// SetLimit sets the Limit field's value.
func (s *ListDatabasesInput) SetLimit(v int32) *ListDatabasesInput {
	s.Limit = &v
	return s
}

// SetOffset sets the Offset field's value.
func (s *ListDatabasesInput) SetOffset(v int32) *ListDatabasesInput {
	s.Offset = &v
	return s
}

type ListDatabasesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListDatabasesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDatabasesOutput) GoString() string {
	return s.String()
}

// SetTotal sets the Total field's value.
func (s *ListDatabasesOutput) SetTotal(v int32) *ListDatabasesOutput {
	s.Total = &v
	return s
}

const (
	// EnumOfDBStatusForListDatabasesInputCreating is a EnumOfDBStatusForListDatabasesInput enum value
	EnumOfDBStatusForListDatabasesInputCreating = "Creating"

	// EnumOfDBStatusForListDatabasesInputDeleting is a EnumOfDBStatusForListDatabasesInput enum value
	EnumOfDBStatusForListDatabasesInputDeleting = "Deleting"

	// EnumOfDBStatusForListDatabasesInputRunning is a EnumOfDBStatusForListDatabasesInput enum value
	EnumOfDBStatusForListDatabasesInputRunning = "Running"
)
