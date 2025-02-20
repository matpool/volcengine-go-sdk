// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package kafka

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyTopicAttributesCommon = "ModifyTopicAttributes"

// ModifyTopicAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTopicAttributesCommon operation. The "output" return
// value will be populated with the ModifyTopicAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTopicAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTopicAttributesCommon Send returns without error.
//
// See ModifyTopicAttributesCommon for more information on using the ModifyTopicAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyTopicAttributesCommonRequest method.
//	req, resp := client.ModifyTopicAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) ModifyTopicAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyTopicAttributesCommon,
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

// ModifyTopicAttributesCommon API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation ModifyTopicAttributesCommon for usage and error information.
func (c *KAFKA) ModifyTopicAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyTopicAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyTopicAttributesCommonWithContext is the same as ModifyTopicAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTopicAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) ModifyTopicAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyTopicAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyTopicAttributes = "ModifyTopicAttributes"

// ModifyTopicAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyTopicAttributes operation. The "output" return
// value will be populated with the ModifyTopicAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyTopicAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyTopicAttributesCommon Send returns without error.
//
// See ModifyTopicAttributes for more information on using the ModifyTopicAttributes
// API call, and error handling.
//
//	// Example sending a request using the ModifyTopicAttributesRequest method.
//	req, resp := client.ModifyTopicAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *KAFKA) ModifyTopicAttributesRequest(input *ModifyTopicAttributesInput) (req *request.Request, output *ModifyTopicAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyTopicAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyTopicAttributesInput{}
	}

	output = &ModifyTopicAttributesOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyTopicAttributes API operation for KAFKA.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for KAFKA's
// API operation ModifyTopicAttributes for usage and error information.
func (c *KAFKA) ModifyTopicAttributes(input *ModifyTopicAttributesInput) (*ModifyTopicAttributesOutput, error) {
	req, out := c.ModifyTopicAttributesRequest(input)
	return out, req.Send()
}

// ModifyTopicAttributesWithContext is the same as ModifyTopicAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyTopicAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *KAFKA) ModifyTopicAttributesWithContext(ctx volcengine.Context, input *ModifyTopicAttributesInput, opts ...request.Option) (*ModifyTopicAttributesOutput, error) {
	req, out := c.ModifyTopicAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyTopicAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	InstanceId *string `type:"string"`

	TopicName *string `type:"string"`
}

// String returns the string representation
func (s ModifyTopicAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTopicAttributesInput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ModifyTopicAttributesInput) SetDescription(v string) *ModifyTopicAttributesInput {
	s.Description = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyTopicAttributesInput) SetInstanceId(v string) *ModifyTopicAttributesInput {
	s.InstanceId = &v
	return s
}

// SetTopicName sets the TopicName field's value.
func (s *ModifyTopicAttributesInput) SetTopicName(v string) *ModifyTopicAttributesInput {
	s.TopicName = &v
	return s
}

type ModifyTopicAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s ModifyTopicAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyTopicAttributesOutput) GoString() string {
	return s.String()
}
