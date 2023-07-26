// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opResetAccountPasswordCommon = "ResetAccountPassword"

// ResetAccountPasswordCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ResetAccountPasswordCommon operation. The "output" return
// value will be populated with the ResetAccountPasswordCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResetAccountPasswordCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResetAccountPasswordCommon Send returns without error.
//
// See ResetAccountPasswordCommon for more information on using the ResetAccountPasswordCommon
// API call, and error handling.
//
//    // Example sending a request using the ResetAccountPasswordCommonRequest method.
//    req, resp := client.ResetAccountPasswordCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ResetAccountPasswordCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opResetAccountPasswordCommon,
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

// ResetAccountPasswordCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ResetAccountPasswordCommon for usage and error information.
func (c *RDSMYSQL) ResetAccountPasswordCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ResetAccountPasswordCommonRequest(input)
	return out, req.Send()
}

// ResetAccountPasswordCommonWithContext is the same as ResetAccountPasswordCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ResetAccountPasswordCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ResetAccountPasswordCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ResetAccountPasswordCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opResetAccountPassword = "ResetAccountPassword"

// ResetAccountPasswordRequest generates a "volcengine/request.Request" representing the
// client's request for the ResetAccountPassword operation. The "output" return
// value will be populated with the ResetAccountPasswordCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ResetAccountPasswordCommon Request to send the API call to the service.
// the "output" return value is not valid until after ResetAccountPasswordCommon Send returns without error.
//
// See ResetAccountPassword for more information on using the ResetAccountPassword
// API call, and error handling.
//
//    // Example sending a request using the ResetAccountPasswordRequest method.
//    req, resp := client.ResetAccountPasswordRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ResetAccountPasswordRequest(input *ResetAccountPasswordInput) (req *request.Request, output *ResetAccountPasswordOutput) {
	op := &request.Operation{
		Name:       opResetAccountPassword,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ResetAccountPasswordInput{}
	}

	output = &ResetAccountPasswordOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ResetAccountPassword API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ResetAccountPassword for usage and error information.
func (c *RDSMYSQL) ResetAccountPassword(input *ResetAccountPasswordInput) (*ResetAccountPasswordOutput, error) {
	req, out := c.ResetAccountPasswordRequest(input)
	return out, req.Send()
}

// ResetAccountPasswordWithContext is the same as ResetAccountPassword with the addition of
// the ability to pass a context and additional request options.
//
// See ResetAccountPassword for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ResetAccountPasswordWithContext(ctx volcengine.Context, input *ResetAccountPasswordInput, opts ...request.Option) (*ResetAccountPasswordOutput, error) {
	req, out := c.ResetAccountPasswordRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ResetAccountPasswordInput struct {
	_ struct{} `type:"structure"`

	// AccountName is a required field
	AccountName *string `min:"2" max:"32" type:"string" required:"true"`

	AccountPassword *string `min:"8" max:"32" type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ResetAccountPasswordInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResetAccountPasswordInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ResetAccountPasswordInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ResetAccountPasswordInput"}
	if s.AccountName == nil {
		invalidParams.Add(request.NewErrParamRequired("AccountName"))
	}
	if s.AccountName != nil && len(*s.AccountName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("AccountName", 2))
	}
	if s.AccountName != nil && len(*s.AccountName) > 32 {
		invalidParams.Add(request.NewErrParamMaxLen("AccountName", 32, *s.AccountName))
	}
	if s.AccountPassword != nil && len(*s.AccountPassword) < 8 {
		invalidParams.Add(request.NewErrParamMinLen("AccountPassword", 8))
	}
	if s.AccountPassword != nil && len(*s.AccountPassword) > 32 {
		invalidParams.Add(request.NewErrParamMaxLen("AccountPassword", 32, *s.AccountPassword))
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
func (s *ResetAccountPasswordInput) SetAccountName(v string) *ResetAccountPasswordInput {
	s.AccountName = &v
	return s
}

// SetAccountPassword sets the AccountPassword field's value.
func (s *ResetAccountPasswordInput) SetAccountPassword(v string) *ResetAccountPasswordInput {
	s.AccountPassword = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ResetAccountPasswordInput) SetInstanceId(v string) *ResetAccountPasswordInput {
	s.InstanceId = &v
	return s
}

type ResetAccountPasswordOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ResetAccountPasswordOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ResetAccountPasswordOutput) GoString() string {
	return s.String()
}
