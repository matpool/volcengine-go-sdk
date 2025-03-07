// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDetachKeyPairCommon = "DetachKeyPair"

// DetachKeyPairCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachKeyPairCommon operation. The "output" return
// value will be populated with the DetachKeyPairCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachKeyPairCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachKeyPairCommon Send returns without error.
//
// See DetachKeyPairCommon for more information on using the DetachKeyPairCommon
// API call, and error handling.
//
//	// Example sending a request using the DetachKeyPairCommonRequest method.
//	req, resp := client.DetachKeyPairCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) DetachKeyPairCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDetachKeyPairCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DetachKeyPairCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DetachKeyPairCommon for usage and error information.
func (c *ECS) DetachKeyPairCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DetachKeyPairCommonRequest(input)
	return out, req.Send()
}

// DetachKeyPairCommonWithContext is the same as DetachKeyPairCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DetachKeyPairCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DetachKeyPairCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DetachKeyPairCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDetachKeyPair = "DetachKeyPair"

// DetachKeyPairRequest generates a "volcengine/request.Request" representing the
// client's request for the DetachKeyPair operation. The "output" return
// value will be populated with the DetachKeyPairCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DetachKeyPairCommon Request to send the API call to the service.
// the "output" return value is not valid until after DetachKeyPairCommon Send returns without error.
//
// See DetachKeyPair for more information on using the DetachKeyPair
// API call, and error handling.
//
//	// Example sending a request using the DetachKeyPairRequest method.
//	req, resp := client.DetachKeyPairRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ECS) DetachKeyPairRequest(input *DetachKeyPairInput) (req *request.Request, output *DetachKeyPairOutput) {
	op := &request.Operation{
		Name:       opDetachKeyPair,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DetachKeyPairInput{}
	}

	output = &DetachKeyPairOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DetachKeyPair API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DetachKeyPair for usage and error information.
func (c *ECS) DetachKeyPair(input *DetachKeyPairInput) (*DetachKeyPairOutput, error) {
	req, out := c.DetachKeyPairRequest(input)
	return out, req.Send()
}

// DetachKeyPairWithContext is the same as DetachKeyPair with the addition of
// the ability to pass a context and additional request options.
//
// See DetachKeyPair for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DetachKeyPairWithContext(ctx volcengine.Context, input *DetachKeyPairInput, opts ...request.Option) (*DetachKeyPairOutput, error) {
	req, out := c.DetachKeyPairRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DetachKeyPairInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	InstanceIds []*string `type:"list"`

	KeyPairId *string `type:"string"`

	KeyPairName *string `type:"string"`
}

// String returns the string representation
func (s DetachKeyPairInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachKeyPairInput) GoString() string {
	return s.String()
}

// SetClientToken sets the ClientToken field's value.
func (s *DetachKeyPairInput) SetClientToken(v string) *DetachKeyPairInput {
	s.ClientToken = &v
	return s
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *DetachKeyPairInput) SetInstanceIds(v []*string) *DetachKeyPairInput {
	s.InstanceIds = v
	return s
}

// SetKeyPairId sets the KeyPairId field's value.
func (s *DetachKeyPairInput) SetKeyPairId(v string) *DetachKeyPairInput {
	s.KeyPairId = &v
	return s
}

// SetKeyPairName sets the KeyPairName field's value.
func (s *DetachKeyPairInput) SetKeyPairName(v string) *DetachKeyPairInput {
	s.KeyPairName = &v
	return s
}

type DetachKeyPairOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OperationDetails []*OperationDetailForDetachKeyPairOutput `type:"list"`
}

// String returns the string representation
func (s DetachKeyPairOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DetachKeyPairOutput) GoString() string {
	return s.String()
}

// SetOperationDetails sets the OperationDetails field's value.
func (s *DetachKeyPairOutput) SetOperationDetails(v []*OperationDetailForDetachKeyPairOutput) *DetachKeyPairOutput {
	s.OperationDetails = v
	return s
}

type ErrorForDetachKeyPairOutput struct {
	_ struct{} `type:"structure"`

	Code *string `type:"string"`

	Message *string `type:"string"`
}

// String returns the string representation
func (s ErrorForDetachKeyPairOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForDetachKeyPairOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForDetachKeyPairOutput) SetCode(v string) *ErrorForDetachKeyPairOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ErrorForDetachKeyPairOutput) SetMessage(v string) *ErrorForDetachKeyPairOutput {
	s.Message = &v
	return s
}

type OperationDetailForDetachKeyPairOutput struct {
	_ struct{} `type:"structure"`

	Error *ErrorForDetachKeyPairOutput `type:"structure"`

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s OperationDetailForDetachKeyPairOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OperationDetailForDetachKeyPairOutput) GoString() string {
	return s.String()
}

// SetError sets the Error field's value.
func (s *OperationDetailForDetachKeyPairOutput) SetError(v *ErrorForDetachKeyPairOutput) *OperationDetailForDetachKeyPairOutput {
	s.Error = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *OperationDetailForDetachKeyPairOutput) SetInstanceId(v string) *OperationDetailForDetachKeyPairOutput {
	s.InstanceId = &v
	return s
}
