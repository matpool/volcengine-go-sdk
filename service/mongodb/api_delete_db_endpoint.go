// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package mongodb

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteDBEndpointCommon = "DeleteDBEndpoint"

// DeleteDBEndpointCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDBEndpointCommon operation. The "output" return
// value will be populated with the DeleteDBEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDBEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDBEndpointCommon Send returns without error.
//
// See DeleteDBEndpointCommon for more information on using the DeleteDBEndpointCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteDBEndpointCommonRequest method.
//    req, resp := client.DeleteDBEndpointCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DeleteDBEndpointCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteDBEndpointCommon,
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

// DeleteDBEndpointCommon API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DeleteDBEndpointCommon for usage and error information.
func (c *MONGODB) DeleteDBEndpointCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteDBEndpointCommonRequest(input)
	return out, req.Send()
}

// DeleteDBEndpointCommonWithContext is the same as DeleteDBEndpointCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDBEndpointCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DeleteDBEndpointCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteDBEndpointCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteDBEndpoint = "DeleteDBEndpoint"

// DeleteDBEndpointRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteDBEndpoint operation. The "output" return
// value will be populated with the DeleteDBEndpointCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteDBEndpointCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteDBEndpointCommon Send returns without error.
//
// See DeleteDBEndpoint for more information on using the DeleteDBEndpoint
// API call, and error handling.
//
//    // Example sending a request using the DeleteDBEndpointRequest method.
//    req, resp := client.DeleteDBEndpointRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *MONGODB) DeleteDBEndpointRequest(input *DeleteDBEndpointInput) (req *request.Request, output *DeleteDBEndpointOutput) {
	op := &request.Operation{
		Name:       opDeleteDBEndpoint,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteDBEndpointInput{}
	}

	output = &DeleteDBEndpointOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteDBEndpoint API operation for MONGODB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for MONGODB's
// API operation DeleteDBEndpoint for usage and error information.
func (c *MONGODB) DeleteDBEndpoint(input *DeleteDBEndpointInput) (*DeleteDBEndpointOutput, error) {
	req, out := c.DeleteDBEndpointRequest(input)
	return out, req.Send()
}

// DeleteDBEndpointWithContext is the same as DeleteDBEndpoint with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteDBEndpoint for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MONGODB) DeleteDBEndpointWithContext(ctx volcengine.Context, input *DeleteDBEndpointInput, opts ...request.Option) (*DeleteDBEndpointOutput, error) {
	req, out := c.DeleteDBEndpointRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteDBEndpointInput struct {
	_ struct{} `type:"structure"`

	EndpointId *string `type:"string"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	MongosNodeIds []*string `type:"list"`
}

// String returns the string representation
func (s DeleteDBEndpointInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDBEndpointInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteDBEndpointInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteDBEndpointInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndpointId sets the EndpointId field's value.
func (s *DeleteDBEndpointInput) SetEndpointId(v string) *DeleteDBEndpointInput {
	s.EndpointId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DeleteDBEndpointInput) SetInstanceId(v string) *DeleteDBEndpointInput {
	s.InstanceId = &v
	return s
}

// SetMongosNodeIds sets the MongosNodeIds field's value.
func (s *DeleteDBEndpointInput) SetMongosNodeIds(v []*string) *DeleteDBEndpointInput {
	s.MongosNodeIds = v
	return s
}

type DeleteDBEndpointOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteDBEndpointOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteDBEndpointOutput) GoString() string {
	return s.String()
}
