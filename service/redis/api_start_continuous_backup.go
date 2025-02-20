// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStartContinuousBackupCommon = "StartContinuousBackup"

// StartContinuousBackupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StartContinuousBackupCommon operation. The "output" return
// value will be populated with the StartContinuousBackupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartContinuousBackupCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartContinuousBackupCommon Send returns without error.
//
// See StartContinuousBackupCommon for more information on using the StartContinuousBackupCommon
// API call, and error handling.
//
//    // Example sending a request using the StartContinuousBackupCommonRequest method.
//    req, resp := client.StartContinuousBackupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) StartContinuousBackupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStartContinuousBackupCommon,
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

// StartContinuousBackupCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation StartContinuousBackupCommon for usage and error information.
func (c *REDIS) StartContinuousBackupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StartContinuousBackupCommonRequest(input)
	return out, req.Send()
}

// StartContinuousBackupCommonWithContext is the same as StartContinuousBackupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StartContinuousBackupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) StartContinuousBackupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StartContinuousBackupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStartContinuousBackup = "StartContinuousBackup"

// StartContinuousBackupRequest generates a "volcengine/request.Request" representing the
// client's request for the StartContinuousBackup operation. The "output" return
// value will be populated with the StartContinuousBackupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StartContinuousBackupCommon Request to send the API call to the service.
// the "output" return value is not valid until after StartContinuousBackupCommon Send returns without error.
//
// See StartContinuousBackup for more information on using the StartContinuousBackup
// API call, and error handling.
//
//    // Example sending a request using the StartContinuousBackupRequest method.
//    req, resp := client.StartContinuousBackupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) StartContinuousBackupRequest(input *StartContinuousBackupInput) (req *request.Request, output *StartContinuousBackupOutput) {
	op := &request.Operation{
		Name:       opStartContinuousBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StartContinuousBackupInput{}
	}

	output = &StartContinuousBackupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StartContinuousBackup API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation StartContinuousBackup for usage and error information.
func (c *REDIS) StartContinuousBackup(input *StartContinuousBackupInput) (*StartContinuousBackupOutput, error) {
	req, out := c.StartContinuousBackupRequest(input)
	return out, req.Send()
}

// StartContinuousBackupWithContext is the same as StartContinuousBackup with the addition of
// the ability to pass a context and additional request options.
//
// See StartContinuousBackup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) StartContinuousBackupWithContext(ctx volcengine.Context, input *StartContinuousBackupInput, opts ...request.Option) (*StartContinuousBackupOutput, error) {
	req, out := c.StartContinuousBackupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StartContinuousBackupInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s StartContinuousBackupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartContinuousBackupInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *StartContinuousBackupInput) SetClientToken(v string) *StartContinuousBackupInput {
	s.ClientToken = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *StartContinuousBackupInput) SetInstanceId(v string) *StartContinuousBackupInput {
	s.InstanceId = &v
	return s
}

type StartContinuousBackupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s StartContinuousBackupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StartContinuousBackupOutput) GoString() string {
	return s.String()
}
