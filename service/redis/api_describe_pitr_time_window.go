// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribePitrTimeWindowCommon = "DescribePitrTimeWindow"

// DescribePitrTimeWindowCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribePitrTimeWindowCommon operation. The "output" return
// value will be populated with the DescribePitrTimeWindowCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribePitrTimeWindowCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribePitrTimeWindowCommon Send returns without error.
//
// See DescribePitrTimeWindowCommon for more information on using the DescribePitrTimeWindowCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribePitrTimeWindowCommonRequest method.
//    req, resp := client.DescribePitrTimeWindowCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribePitrTimeWindowCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribePitrTimeWindowCommon,
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

// DescribePitrTimeWindowCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribePitrTimeWindowCommon for usage and error information.
func (c *REDIS) DescribePitrTimeWindowCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribePitrTimeWindowCommonRequest(input)
	return out, req.Send()
}

// DescribePitrTimeWindowCommonWithContext is the same as DescribePitrTimeWindowCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribePitrTimeWindowCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribePitrTimeWindowCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribePitrTimeWindowCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribePitrTimeWindow = "DescribePitrTimeWindow"

// DescribePitrTimeWindowRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribePitrTimeWindow operation. The "output" return
// value will be populated with the DescribePitrTimeWindowCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribePitrTimeWindowCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribePitrTimeWindowCommon Send returns without error.
//
// See DescribePitrTimeWindow for more information on using the DescribePitrTimeWindow
// API call, and error handling.
//
//    // Example sending a request using the DescribePitrTimeWindowRequest method.
//    req, resp := client.DescribePitrTimeWindowRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribePitrTimeWindowRequest(input *DescribePitrTimeWindowInput) (req *request.Request, output *DescribePitrTimeWindowOutput) {
	op := &request.Operation{
		Name:       opDescribePitrTimeWindow,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribePitrTimeWindowInput{}
	}

	output = &DescribePitrTimeWindowOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribePitrTimeWindow API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribePitrTimeWindow for usage and error information.
func (c *REDIS) DescribePitrTimeWindow(input *DescribePitrTimeWindowInput) (*DescribePitrTimeWindowOutput, error) {
	req, out := c.DescribePitrTimeWindowRequest(input)
	return out, req.Send()
}

// DescribePitrTimeWindowWithContext is the same as DescribePitrTimeWindow with the addition of
// the ability to pass a context and additional request options.
//
// See DescribePitrTimeWindow for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribePitrTimeWindowWithContext(ctx volcengine.Context, input *DescribePitrTimeWindowInput, opts ...request.Option) (*DescribePitrTimeWindowOutput, error) {
	req, out := c.DescribePitrTimeWindowRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribePitrTimeWindowInput struct {
	_ struct{} `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s DescribePitrTimeWindowInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribePitrTimeWindowInput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribePitrTimeWindowInput) SetInstanceId(v string) *DescribePitrTimeWindowInput {
	s.InstanceId = &v
	return s
}

type DescribePitrTimeWindowOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	EndTime *string `type:"string"`

	StartTime *string `type:"string"`
}

// String returns the string representation
func (s DescribePitrTimeWindowOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribePitrTimeWindowOutput) GoString() string {
	return s.String()
}

// SetEndTime sets the EndTime field's value.
func (s *DescribePitrTimeWindowOutput) SetEndTime(v string) *DescribePitrTimeWindowOutput {
	s.EndTime = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribePitrTimeWindowOutput) SetStartTime(v string) *DescribePitrTimeWindowOutput {
	s.StartTime = &v
	return s
}
