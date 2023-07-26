// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dbw

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDataCancelExecCommon = "DataCancelExec"

// DataCancelExecCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DataCancelExecCommon operation. The "output" return
// value will be populated with the DataCancelExecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DataCancelExecCommon Request to send the API call to the service.
// the "output" return value is not valid until after DataCancelExecCommon Send returns without error.
//
// See DataCancelExecCommon for more information on using the DataCancelExecCommon
// API call, and error handling.
//
//    // Example sending a request using the DataCancelExecCommonRequest method.
//    req, resp := client.DataCancelExecCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DataCancelExecCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDataCancelExecCommon,
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

// DataCancelExecCommon API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DataCancelExecCommon for usage and error information.
func (c *DBW) DataCancelExecCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DataCancelExecCommonRequest(input)
	return out, req.Send()
}

// DataCancelExecCommonWithContext is the same as DataCancelExecCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DataCancelExecCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DataCancelExecCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DataCancelExecCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDataCancelExec = "DataCancelExec"

// DataCancelExecRequest generates a "volcengine/request.Request" representing the
// client's request for the DataCancelExec operation. The "output" return
// value will be populated with the DataCancelExecCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DataCancelExecCommon Request to send the API call to the service.
// the "output" return value is not valid until after DataCancelExecCommon Send returns without error.
//
// See DataCancelExec for more information on using the DataCancelExec
// API call, and error handling.
//
//    // Example sending a request using the DataCancelExecRequest method.
//    req, resp := client.DataCancelExecRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DataCancelExecRequest(input *DataCancelExecInput) (req *request.Request, output *DataCancelExecOutput) {
	op := &request.Operation{
		Name:       opDataCancelExec,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DataCancelExecInput{}
	}

	output = &DataCancelExecOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DataCancelExec API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DataCancelExec for usage and error information.
func (c *DBW) DataCancelExec(input *DataCancelExecInput) (*DataCancelExecOutput, error) {
	req, out := c.DataCancelExecRequest(input)
	return out, req.Send()
}

// DataCancelExecWithContext is the same as DataCancelExec with the addition of
// the ability to pass a context and additional request options.
//
// See DataCancelExec for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DataCancelExecWithContext(ctx volcengine.Context, input *DataCancelExecInput, opts ...request.Option) (*DataCancelExecOutput, error) {
	req, out := c.DataCancelExecRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataCancelExecInput struct {
	_ struct{} `type:"structure"`

	SessionId *string `type:"string"`
}

// String returns the string representation
func (s DataCancelExecInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataCancelExecInput) GoString() string {
	return s.String()
}

// SetSessionId sets the SessionId field's value.
func (s *DataCancelExecInput) SetSessionId(v string) *DataCancelExecInput {
	s.SessionId = &v
	return s
}

type DataCancelExecOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DataCancelExecOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataCancelExecOutput) GoString() string {
	return s.String()
}
