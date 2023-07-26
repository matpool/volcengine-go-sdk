// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyServerGroupBackendServersCommon = "ModifyServerGroupBackendServers"

// ModifyServerGroupBackendServersCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyServerGroupBackendServersCommon operation. The "output" return
// value will be populated with the ModifyServerGroupBackendServersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyServerGroupBackendServersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyServerGroupBackendServersCommon Send returns without error.
//
// See ModifyServerGroupBackendServersCommon for more information on using the ModifyServerGroupBackendServersCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyServerGroupBackendServersCommonRequest method.
//	req, resp := client.ModifyServerGroupBackendServersCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) ModifyServerGroupBackendServersCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyServerGroupBackendServersCommon,
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

// ModifyServerGroupBackendServersCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyServerGroupBackendServersCommon for usage and error information.
func (c *ALB) ModifyServerGroupBackendServersCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyServerGroupBackendServersCommonRequest(input)
	return out, req.Send()
}

// ModifyServerGroupBackendServersCommonWithContext is the same as ModifyServerGroupBackendServersCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyServerGroupBackendServersCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyServerGroupBackendServersCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyServerGroupBackendServersCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyServerGroupBackendServers = "ModifyServerGroupBackendServers"

// ModifyServerGroupBackendServersRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyServerGroupBackendServers operation. The "output" return
// value will be populated with the ModifyServerGroupBackendServersCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyServerGroupBackendServersCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyServerGroupBackendServersCommon Send returns without error.
//
// See ModifyServerGroupBackendServers for more information on using the ModifyServerGroupBackendServers
// API call, and error handling.
//
//	// Example sending a request using the ModifyServerGroupBackendServersRequest method.
//	req, resp := client.ModifyServerGroupBackendServersRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) ModifyServerGroupBackendServersRequest(input *ModifyServerGroupBackendServersInput) (req *request.Request, output *ModifyServerGroupBackendServersOutput) {
	op := &request.Operation{
		Name:       opModifyServerGroupBackendServers,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyServerGroupBackendServersInput{}
	}

	output = &ModifyServerGroupBackendServersOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyServerGroupBackendServers API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyServerGroupBackendServers for usage and error information.
func (c *ALB) ModifyServerGroupBackendServers(input *ModifyServerGroupBackendServersInput) (*ModifyServerGroupBackendServersOutput, error) {
	req, out := c.ModifyServerGroupBackendServersRequest(input)
	return out, req.Send()
}

// ModifyServerGroupBackendServersWithContext is the same as ModifyServerGroupBackendServers with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyServerGroupBackendServers for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ModifyServerGroupBackendServersWithContext(ctx volcengine.Context, input *ModifyServerGroupBackendServersInput, opts ...request.Option) (*ModifyServerGroupBackendServersOutput, error) {
	req, out := c.ModifyServerGroupBackendServersRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyServerGroupBackendServersInput struct {
	_ struct{} `type:"structure"`

	// ServerGroupId is a required field
	ServerGroupId *string `type:"string" required:"true"`

	// Servers is a required field
	Servers []*ServerForModifyServerGroupBackendServersInput `type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyServerGroupBackendServersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyServerGroupBackendServersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyServerGroupBackendServersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyServerGroupBackendServersInput"}
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

// SetServerGroupId sets the ServerGroupId field's value.
func (s *ModifyServerGroupBackendServersInput) SetServerGroupId(v string) *ModifyServerGroupBackendServersInput {
	s.ServerGroupId = &v
	return s
}

// SetServers sets the Servers field's value.
func (s *ModifyServerGroupBackendServersInput) SetServers(v []*ServerForModifyServerGroupBackendServersInput) *ModifyServerGroupBackendServersInput {
	s.Servers = v
	return s
}

type ModifyServerGroupBackendServersOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyServerGroupBackendServersOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyServerGroupBackendServersOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyServerGroupBackendServersOutput) SetRequestId(v string) *ModifyServerGroupBackendServersOutput {
	s.RequestId = &v
	return s
}

type ServerForModifyServerGroupBackendServersInput struct {
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
func (s ServerForModifyServerGroupBackendServersInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ServerForModifyServerGroupBackendServersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ServerForModifyServerGroupBackendServersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ServerForModifyServerGroupBackendServersInput"}
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
func (s *ServerForModifyServerGroupBackendServersInput) SetDescription(v string) *ServerForModifyServerGroupBackendServersInput {
	s.Description = &v
	return s
}

// SetPort sets the Port field's value.
func (s *ServerForModifyServerGroupBackendServersInput) SetPort(v int64) *ServerForModifyServerGroupBackendServersInput {
	s.Port = &v
	return s
}

// SetServerId sets the ServerId field's value.
func (s *ServerForModifyServerGroupBackendServersInput) SetServerId(v string) *ServerForModifyServerGroupBackendServersInput {
	s.ServerId = &v
	return s
}

// SetWeight sets the Weight field's value.
func (s *ServerForModifyServerGroupBackendServersInput) SetWeight(v int64) *ServerForModifyServerGroupBackendServersInput {
	s.Weight = &v
	return s
}
