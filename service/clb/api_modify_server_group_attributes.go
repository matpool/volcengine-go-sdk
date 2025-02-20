// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyServerGroupAttributesCommon = "ModifyServerGroupAttributes"

// ModifyServerGroupAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyServerGroupAttributesCommon operation. The "output" return
// value will be populated with the ModifyServerGroupAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyServerGroupAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyServerGroupAttributesCommon Send returns without error.
//
// See ModifyServerGroupAttributesCommon for more information on using the ModifyServerGroupAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyServerGroupAttributesCommonRequest method.
//	req, resp := client.ModifyServerGroupAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) ModifyServerGroupAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyServerGroupAttributesCommon,
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

// ModifyServerGroupAttributesCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ModifyServerGroupAttributesCommon for usage and error information.
func (c *CLB) ModifyServerGroupAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyServerGroupAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifyServerGroupAttributesCommonWithContext is the same as ModifyServerGroupAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyServerGroupAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) ModifyServerGroupAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyServerGroupAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyServerGroupAttributes = "ModifyServerGroupAttributes"

// ModifyServerGroupAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyServerGroupAttributes operation. The "output" return
// value will be populated with the ModifyServerGroupAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyServerGroupAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyServerGroupAttributesCommon Send returns without error.
//
// See ModifyServerGroupAttributes for more information on using the ModifyServerGroupAttributes
// API call, and error handling.
//
//	// Example sending a request using the ModifyServerGroupAttributesRequest method.
//	req, resp := client.ModifyServerGroupAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) ModifyServerGroupAttributesRequest(input *ModifyServerGroupAttributesInput) (req *request.Request, output *ModifyServerGroupAttributesOutput) {
	op := &request.Operation{
		Name:       opModifyServerGroupAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyServerGroupAttributesInput{}
	}

	output = &ModifyServerGroupAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyServerGroupAttributes API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ModifyServerGroupAttributes for usage and error information.
func (c *CLB) ModifyServerGroupAttributes(input *ModifyServerGroupAttributesInput) (*ModifyServerGroupAttributesOutput, error) {
	req, out := c.ModifyServerGroupAttributesRequest(input)
	return out, req.Send()
}

// ModifyServerGroupAttributesWithContext is the same as ModifyServerGroupAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyServerGroupAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) ModifyServerGroupAttributesWithContext(ctx volcengine.Context, input *ModifyServerGroupAttributesInput, opts ...request.Option) (*ModifyServerGroupAttributesOutput, error) {
	req, out := c.ModifyServerGroupAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyServerGroupAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// ServerGroupId is a required field
	ServerGroupId *string `type:"string" required:"true"`

	ServerGroupName *string `type:"string"`

	// Servers is a required field
	Servers []*ServerForModifyServerGroupAttributesInput `type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyServerGroupAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyServerGroupAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyServerGroupAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyServerGroupAttributesInput"}
	if s.ServerGroupId == nil {
		invalidParams.Add(request.NewErrParamRequired("ServerGroupId"))
	}
	if s.Servers == nil {
		invalidParams.Add(request.NewErrParamRequired("Servers"))
	}
	if s.Servers != nil {
		for i, v := range s.Servers {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Servers", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyServerGroupAttributesInput) SetDescription(v string) *ModifyServerGroupAttributesInput {
	s.Description = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ModifyServerGroupAttributesInput) SetServerGroupId(v string) *ModifyServerGroupAttributesInput {
	s.ServerGroupId = &v
	return s
}

// SetServerGroupName sets the ServerGroupName field's value.
func (s *ModifyServerGroupAttributesInput) SetServerGroupName(v string) *ModifyServerGroupAttributesInput {
	s.ServerGroupName = &v
	return s
}

// SetServers sets the Servers field's value.
func (s *ModifyServerGroupAttributesInput) SetServers(v []*ServerForModifyServerGroupAttributesInput) *ModifyServerGroupAttributesInput {
	s.Servers = v
	return s
}

type ModifyServerGroupAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyServerGroupAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyServerGroupAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyServerGroupAttributesOutput) SetRequestId(v string) *ModifyServerGroupAttributesOutput {
	s.RequestId = &v
	return s
}

type ServerForModifyServerGroupAttributesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// Port is a required field
	Port *int64 `type:"integer" required:"true"`

	// ServerId is a required field
	ServerId *string `type:"string" required:"true"`

	// Weight is a required field
	Weight *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s ServerForModifyServerGroupAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerForModifyServerGroupAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ServerForModifyServerGroupAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ServerForModifyServerGroupAttributesInput"}
	if s.Port == nil {
		invalidParams.Add(request.NewErrParamRequired("Port"))
	}
	if s.ServerId == nil {
		invalidParams.Add(request.NewErrParamRequired("ServerId"))
	}
	if s.Weight == nil {
		invalidParams.Add(request.NewErrParamRequired("Weight"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ServerForModifyServerGroupAttributesInput) SetDescription(v string) *ServerForModifyServerGroupAttributesInput {
	s.Description = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ServerForModifyServerGroupAttributesInput) SetPort(v int64) *ServerForModifyServerGroupAttributesInput {
	s.Port = &v
	return s
}

// SetServerId sets the ServerId field's value.
func (s *ServerForModifyServerGroupAttributesInput) SetServerId(v string) *ServerForModifyServerGroupAttributesInput {
	s.ServerId = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ServerForModifyServerGroupAttributesInput) SetWeight(v int64) *ServerForModifyServerGroupAttributesInput {
	s.Weight = &v
	return s
}
