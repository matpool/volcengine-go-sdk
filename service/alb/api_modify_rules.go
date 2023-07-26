// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

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
func (c *ALB) ModifyRulesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
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

// ModifyRulesCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyRulesCommon for usage and error information.
func (c *ALB) ModifyRulesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
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
func (c *ALB) ModifyRulesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
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
func (c *ALB) ModifyRulesRequest(input *ModifyRulesInput) (req *request.Request, output *ModifyRulesOutput) {
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

// ModifyRules API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ModifyRules for usage and error information.
func (c *ALB) ModifyRules(input *ModifyRulesInput) (*ModifyRulesOutput, error) {
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
func (c *ALB) ModifyRulesWithContext(ctx volcengine.Context, input *ModifyRulesInput, opts ...request.Option) (*ModifyRulesOutput, error) {
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

type RedirectConfigForModifyRulesInput struct {
	_ struct{} `type:"structure"`

	RedirectDomain *string `min:"1" max:"128" type:"string"`

	RedirectHttpCode *string `type:"string"`

	RedirectPort *string `type:"string"`

	RedirectProtocol *string `type:"string"`

	RedirectUri *string `min:"1" max:"128" type:"string"`
}

// String returns the string representation
func (s RedirectConfigForModifyRulesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s RedirectConfigForModifyRulesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RedirectConfigForModifyRulesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RedirectConfigForModifyRulesInput"}
	if s.RedirectDomain != nil && len(*s.RedirectDomain) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RedirectDomain", 1))
	}
	if s.RedirectDomain != nil && len(*s.RedirectDomain) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("RedirectDomain", 128, *s.RedirectDomain))
	}
	if s.RedirectUri != nil && len(*s.RedirectUri) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("RedirectUri", 1))
	}
	if s.RedirectUri != nil && len(*s.RedirectUri) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("RedirectUri", 128, *s.RedirectUri))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetRedirectDomain sets the RedirectDomain field's value.
func (s *RedirectConfigForModifyRulesInput) SetRedirectDomain(v string) *RedirectConfigForModifyRulesInput {
	s.RedirectDomain = &v
	return s
}

// SetRedirectHttpCode sets the RedirectHttpCode field's value.
func (s *RedirectConfigForModifyRulesInput) SetRedirectHttpCode(v string) *RedirectConfigForModifyRulesInput {
	s.RedirectHttpCode = &v
	return s
}

// SetRedirectPort sets the RedirectPort field's value.
func (s *RedirectConfigForModifyRulesInput) SetRedirectPort(v string) *RedirectConfigForModifyRulesInput {
	s.RedirectPort = &v
	return s
}

// SetRedirectProtocol sets the RedirectProtocol field's value.
func (s *RedirectConfigForModifyRulesInput) SetRedirectProtocol(v string) *RedirectConfigForModifyRulesInput {
	s.RedirectProtocol = &v
	return s
}

// SetRedirectUri sets the RedirectUri field's value.
func (s *RedirectConfigForModifyRulesInput) SetRedirectUri(v string) *RedirectConfigForModifyRulesInput {
	s.RedirectUri = &v
	return s
}

type RuleForModifyRulesInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	RedirectConfig *RedirectConfigForModifyRulesInput `type:"structure"`

	RuleAction *string `type:"string"`

	// RuleId is a required field
	RuleId *string `type:"string" required:"true"`

	ServerGroupId *string `type:"string"`

	TrafficLimitEnabled *string `type:"string"`

	TrafficLimitQPS *int64 `type:"integer"`
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
	if s.RedirectConfig != nil {
		if err := s.RedirectConfig.Validate(); err != nil {
			invalidParams.AddNested("RedirectConfig", err.(request.ErrInvalidParams))
		}
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

// SetRedirectConfig sets the RedirectConfig field's value.
func (s *RuleForModifyRulesInput) SetRedirectConfig(v *RedirectConfigForModifyRulesInput) *RuleForModifyRulesInput {
	s.RedirectConfig = v
	return s
}

// SetRuleAction sets the RuleAction field's value.
func (s *RuleForModifyRulesInput) SetRuleAction(v string) *RuleForModifyRulesInput {
	s.RuleAction = &v
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

// SetTrafficLimitEnabled sets the TrafficLimitEnabled field's value.
func (s *RuleForModifyRulesInput) SetTrafficLimitEnabled(v string) *RuleForModifyRulesInput {
	s.TrafficLimitEnabled = &v
	return s
}

// SetTrafficLimitQPS sets the TrafficLimitQPS field's value.
func (s *RuleForModifyRulesInput) SetTrafficLimitQPS(v int64) *RuleForModifyRulesInput {
	s.TrafficLimitQPS = &v
	return s
}
