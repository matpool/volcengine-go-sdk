// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUpdateSystemEventsCommon = "UpdateSystemEvents"

// UpdateSystemEventsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateSystemEventsCommon operation. The "output" return
// value will be populated with the UpdateSystemEventsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateSystemEventsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateSystemEventsCommon Send returns without error.
//
// See UpdateSystemEventsCommon for more information on using the UpdateSystemEventsCommon
// API call, and error handling.
//
//    // Example sending a request using the UpdateSystemEventsCommonRequest method.
//    req, resp := client.UpdateSystemEventsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) UpdateSystemEventsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUpdateSystemEventsCommon,
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

// UpdateSystemEventsCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation UpdateSystemEventsCommon for usage and error information.
func (c *ECS) UpdateSystemEventsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UpdateSystemEventsCommonRequest(input)
	return out, req.Send()
}

// UpdateSystemEventsCommonWithContext is the same as UpdateSystemEventsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateSystemEventsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) UpdateSystemEventsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UpdateSystemEventsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUpdateSystemEvents = "UpdateSystemEvents"

// UpdateSystemEventsRequest generates a "volcengine/request.Request" representing the
// client's request for the UpdateSystemEvents operation. The "output" return
// value will be populated with the UpdateSystemEventsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UpdateSystemEventsCommon Request to send the API call to the service.
// the "output" return value is not valid until after UpdateSystemEventsCommon Send returns without error.
//
// See UpdateSystemEvents for more information on using the UpdateSystemEvents
// API call, and error handling.
//
//    // Example sending a request using the UpdateSystemEventsRequest method.
//    req, resp := client.UpdateSystemEventsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) UpdateSystemEventsRequest(input *UpdateSystemEventsInput) (req *request.Request, output *UpdateSystemEventsOutput) {
	op := &request.Operation{
		Name:       opUpdateSystemEvents,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateSystemEventsInput{}
	}

	output = &UpdateSystemEventsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UpdateSystemEvents API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation UpdateSystemEvents for usage and error information.
func (c *ECS) UpdateSystemEvents(input *UpdateSystemEventsInput) (*UpdateSystemEventsOutput, error) {
	req, out := c.UpdateSystemEventsRequest(input)
	return out, req.Send()
}

// UpdateSystemEventsWithContext is the same as UpdateSystemEvents with the addition of
// the ability to pass a context and additional request options.
//
// See UpdateSystemEvents for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) UpdateSystemEventsWithContext(ctx volcengine.Context, input *UpdateSystemEventsInput, opts ...request.Option) (*UpdateSystemEventsOutput, error) {
	req, out := c.UpdateSystemEventsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ErrorForUpdateSystemEventsOutput struct {
	_ struct{} `type:"structure"`

	Code *string `type:"string"`

	Message *string `type:"string"`
}

// String returns the string representation
func (s ErrorForUpdateSystemEventsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ErrorForUpdateSystemEventsOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *ErrorForUpdateSystemEventsOutput) SetCode(v string) *ErrorForUpdateSystemEventsOutput {
	s.Code = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *ErrorForUpdateSystemEventsOutput) SetMessage(v string) *ErrorForUpdateSystemEventsOutput {
	s.Message = &v
	return s
}

type OperationDetailForUpdateSystemEventsOutput struct {
	_ struct{} `type:"structure"`

	Error *ErrorForUpdateSystemEventsOutput `type:"structure"`

	EventId *string `type:"string"`
}

// String returns the string representation
func (s OperationDetailForUpdateSystemEventsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s OperationDetailForUpdateSystemEventsOutput) GoString() string {
	return s.String()
}

// SetError sets the Error field's value.
func (s *OperationDetailForUpdateSystemEventsOutput) SetError(v *ErrorForUpdateSystemEventsOutput) *OperationDetailForUpdateSystemEventsOutput {
	s.Error = v
	return s
}

// SetEventId sets the EventId field's value.
func (s *OperationDetailForUpdateSystemEventsOutput) SetEventId(v string) *OperationDetailForUpdateSystemEventsOutput {
	s.EventId = &v
	return s
}

type UpdateSystemEventsInput struct {
	_ struct{} `type:"structure"`

	EventIds []*string `type:"list"`

	OperatedEndAt *string `type:"string"`

	OperatedStartAt *string `type:"string"`

	Status *string `type:"string"`

	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s UpdateSystemEventsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateSystemEventsInput) GoString() string {
	return s.String()
}

// SetEventIds sets the EventIds field's value.
func (s *UpdateSystemEventsInput) SetEventIds(v []*string) *UpdateSystemEventsInput {
	s.EventIds = v
	return s
}

// SetOperatedEndAt sets the OperatedEndAt field's value.
func (s *UpdateSystemEventsInput) SetOperatedEndAt(v string) *UpdateSystemEventsInput {
	s.OperatedEndAt = &v
	return s
}

// SetOperatedStartAt sets the OperatedStartAt field's value.
func (s *UpdateSystemEventsInput) SetOperatedStartAt(v string) *UpdateSystemEventsInput {
	s.OperatedStartAt = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *UpdateSystemEventsInput) SetStatus(v string) *UpdateSystemEventsInput {
	s.Status = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *UpdateSystemEventsInput) SetUpdatedAt(v string) *UpdateSystemEventsInput {
	s.UpdatedAt = &v
	return s
}

type UpdateSystemEventsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	OperationDetails []*OperationDetailForUpdateSystemEventsOutput `type:"list"`
}

// String returns the string representation
func (s UpdateSystemEventsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateSystemEventsOutput) GoString() string {
	return s.String()
}

// SetOperationDetails sets the OperationDetails field's value.
func (s *UpdateSystemEventsOutput) SetOperationDetails(v []*OperationDetailForUpdateSystemEventsOutput) *UpdateSystemEventsOutput {
	s.OperationDetails = v
	return s
}
