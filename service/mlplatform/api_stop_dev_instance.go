// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mlplatform

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opStopDevInstanceCommon = "StopDevInstance"

// StopDevInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the StopDevInstanceCommon operation. The "output" return
// value will be populated with the StopDevInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopDevInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopDevInstanceCommon Send returns without error.
//
// See StopDevInstanceCommon for more information on using the StopDevInstanceCommon
// API call, and error handling.
//
//	// Example sending a request using the StopDevInstanceCommonRequest method.
//	req, resp := client.StopDevInstanceCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *MLPLATFORM) StopDevInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opStopDevInstanceCommon,
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

// StopDevInstanceCommon API operation for ML_PLATFORM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM's
// API operation StopDevInstanceCommon for usage and error information.
func (c *MLPLATFORM) StopDevInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.StopDevInstanceCommonRequest(input)
	return out, req.Send()
}

// StopDevInstanceCommonWithContext is the same as StopDevInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See StopDevInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM) StopDevInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.StopDevInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opStopDevInstance = "StopDevInstance"

// StopDevInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the StopDevInstance operation. The "output" return
// value will be populated with the StopDevInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned StopDevInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after StopDevInstanceCommon Send returns without error.
//
// See StopDevInstance for more information on using the StopDevInstance
// API call, and error handling.
//
//	// Example sending a request using the StopDevInstanceRequest method.
//	req, resp := client.StopDevInstanceRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *MLPLATFORM) StopDevInstanceRequest(input *StopDevInstanceInput) (req *request.Request, output *StopDevInstanceOutput) {
	op := &request.Operation{
		Name:       opStopDevInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &StopDevInstanceInput{}
	}

	output = &StopDevInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// StopDevInstance API operation for ML_PLATFORM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ML_PLATFORM's
// API operation StopDevInstance for usage and error information.
func (c *MLPLATFORM) StopDevInstance(input *StopDevInstanceInput) (*StopDevInstanceOutput, error) {
	req, out := c.StopDevInstanceRequest(input)
	return out, req.Send()
}

// StopDevInstanceWithContext is the same as StopDevInstance with the addition of
// the ability to pass a context and additional request options.
//
// See StopDevInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MLPLATFORM) StopDevInstanceWithContext(ctx volcengine.Context, input *StopDevInstanceInput, opts ...request.Option) (*StopDevInstanceOutput, error) {
	req, out := c.StopDevInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type StopDevInstanceInput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StopDevInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StopDevInstanceInput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *StopDevInstanceInput) SetId(v string) *StopDevInstanceInput {
	s.Id = &v
	return s
}

type StopDevInstanceOutput struct {
	_ struct{} `type:"structure" json:",omitempty"`

	Metadata *response.ResponseMetadata

	Id *string `type:"string" json:",omitempty"`
}

// String returns the string representation
func (s StopDevInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s StopDevInstanceOutput) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *StopDevInstanceOutput) SetId(v string) *StopDevInstanceOutput {
	s.Id = &v
	return s
}
