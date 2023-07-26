// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package clb

import (
	"fmt"

	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyRulesCommon = "ModifyRules"

// ModifyRulesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyRulesCommon operation. The "output" return
// value will be populated with the ModifyRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyRulesCommon Send returns without error.
//
// See ModifyRulesCommon for more information on using the ModifyRulesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifyRulesCommonRequest method.
//	req, resp := client.ModifyRulesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) ModifyRulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyRulesCommon,
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

// ModifyRulesCommon API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ModifyRulesCommon for usage and error information.
func (c *CLB) ModifyRulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyRulesCommonRequest(input)
	return out, req.Send()
}

// ModifyRulesCommonWithContext is the same as ModifyRulesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyRulesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) ModifyRulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyRulesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyRules = "ModifyRules"

// ModifyRulesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyRules operation. The "output" return
// value will be populated with the ModifyRulesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyRulesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyRulesCommon Send returns without error.
//
// See ModifyRules for more information on using the ModifyRules
// API call, and error handling.
//
//	// Example sending a request using the ModifyRulesRequest method.
//	req, resp := client.ModifyRulesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *CLB) ModifyRulesRequest(input *ModifyRulesInput) (req *request.Request, output *ModifyRulesOutput) {
	op := &request.Operation{
		Name:       opModifyRules,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyRulesInput{}
	}

	output = &ModifyRulesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifyRules API operation for CLB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for CLB's
// API operation ModifyRules for usage and error information.
func (c *CLB) ModifyRules(input *ModifyRulesInput) (*ModifyRulesOutput, error) {
	req, out := c.ModifyRulesRequest(input)
	return out, req.Send()
}

// ModifyRulesWithContext is the same as ModifyRules with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyRules for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *CLB) ModifyRulesWithContext(ctx volcengine.Context, input *ModifyRulesInput, opts ...request.Option) (*ModifyRulesOutput, error) {
	req, out := c.ModifyRulesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifyRulesInput struct {
	_ struct{} `type:"structure"`

	// ListenerId is a required field
	ListenerId *string `type:"string" required:"true"`

	// Rules is a required field
	Rules []*RuleForModifyRulesInput `type:"list" required:"true"`
}

// String returns the string representation
func (s ModifyRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyRulesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyRulesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyRulesInput"}
	if s.ListenerId == nil {
		invalidParams.Add(request.NewErrParamRequired("ListenerId"))
	}
	if s.Rules == nil {
		invalidParams.Add(request.NewErrParamRequired("Rules"))
	}
	if s.Rules != nil {
		for i, v := range s.Rules {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Rules", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetListenerId sets the ListenerId field's value.
func (s *ModifyRulesInput) SetListenerId(v string) *ModifyRulesInput {
	s.ListenerId = &v
	return s
}

// SetRules sets the Rules field's value.
func (s *ModifyRulesInput) SetRules(v []*RuleForModifyRulesInput) *ModifyRulesInput {
	s.Rules = v
	return s
}

type ModifyRulesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifyRulesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyRulesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifyRulesOutput) SetRequestId(v string) *ModifyRulesOutput {
	s.RequestId = &v
	return s
}

type RuleForModifyRulesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	// RuleId is a required field
	RuleId *string `type:"string" required:"true"`

	ServerGroupId *string `type:"string"`
}

// String returns the string representation
func (s RuleForModifyRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RuleForModifyRulesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RuleForModifyRulesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RuleForModifyRulesInput"}
	if s.RuleId == nil {
		invalidParams.Add(request.NewErrParamRequired("RuleId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *RuleForModifyRulesInput) SetDescription(v string) *RuleForModifyRulesInput {
	s.Description = &v
	return s
}

// SetRuleId sets the RuleId field's value.
func (s *RuleForModifyRulesInput) SetRuleId(v string) *RuleForModifyRulesInput {
	s.RuleId = &v
	return s
}

// SetServerGroupId sets the ServerGroupId field's value.
func (s *RuleForModifyRulesInput) SetServerGroupId(v string) *RuleForModifyRulesInput {
	s.ServerGroupId = &v
	return s
}
