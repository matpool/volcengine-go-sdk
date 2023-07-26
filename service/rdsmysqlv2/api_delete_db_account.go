// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteDBAccountCommon = "DeleteDBAccount"

// DeleteDBAccountCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDBAccountCommon operation. The "output" return
// value will be populated with the DeleteDBAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDBAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDBAccountCommon Send returns without error.
//
// See DeleteDBAccountCommon for more information on using the DeleteDBAccountCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteDBAccountCommonRequest method.
//    req, resp := client.DeleteDBAccountCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DeleteDBAccountCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDBAccountCommon,
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

// DeleteDBAccountCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DeleteDBAccountCommon for usage and error information.
func (c *RDSMYSQLV2) DeleteDBAccountCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDBAccountCommonRequest(input)
	return out, req.Send()
}

// DeleteDBAccountCommonWithContext is the same as DeleteDBAccountCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDBAccountCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DeleteDBAccountCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDBAccountCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDBAccount = "DeleteDBAccount"

// DeleteDBAccountRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDBAccount operation. The "output" return
// value will be populated with the DeleteDBAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDBAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDBAccountCommon Send returns without error.
//
// See DeleteDBAccount for more information on using the DeleteDBAccount
// API call, and error handling.
//
//    // Example sending a request using the DeleteDBAccountRequest method.
//    req, resp := client.DeleteDBAccountRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DeleteDBAccountRequest(input *DeleteDBAccountInput) (req *request.Request, output *DeleteDBAccountOutput) {
	op := &request.Operation{
		Name:       opDeleteDBAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBAccountInput{}
	}

	output = &DeleteDBAccountOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteDBAccount API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DeleteDBAccount for usage and error information.
func (c *RDSMYSQLV2) DeleteDBAccount(input *DeleteDBAccountInput) (*DeleteDBAccountOutput, error) {
	req, out := c.DeleteDBAccountRequest(input)
	return out, req.Send()
}

// DeleteDBAccountWithContext is the same as DeleteDBAccount with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDBAccount for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DeleteDBAccountWithContext(ctx volcengine.Context, input *DeleteDBAccountInput, opts ...request.Option) (*DeleteDBAccountOutput, error) {
	req, out := c.DeleteDBAccountRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteDBAccountInput struct {
	_ struct{} `type:"structure"`

	// AccountName is a required field
	AccountName *string `min:"2" max:"32" type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteDBAccountInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDBAccountInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBAccountInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteDBAccountInput"}
	if s.AccountName == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountName"))
	}
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
func (s *DeleteDBAccountInput) SetAccountName(v string) *DeleteDBAccountInput {
	s.AccountName = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DeleteDBAccountInput) SetInstanceId(v string) *DeleteDBAccountInput {
	s.InstanceId = &v
	return s
}

type DeleteDBAccountOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteDBAccountOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDBAccountOutput) GoString() string {
	return s.String()
}
