// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package natgateway

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifySnatEntryAttributesCommon = "ModifySnatEntryAttributes"

// ModifySnatEntryAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifySnatEntryAttributesCommon operation. The "output" return
// value will be populated with the ModifySnatEntryAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifySnatEntryAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifySnatEntryAttributesCommon Send returns without error.
//
// See ModifySnatEntryAttributesCommon for more information on using the ModifySnatEntryAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the ModifySnatEntryAttributesCommonRequest method.
//	req, resp := client.ModifySnatEntryAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *NATGATEWAY) ModifySnatEntryAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifySnatEntryAttributesCommon,
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

// ModifySnatEntryAttributesCommon API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NATGATEWAY's
// API operation ModifySnatEntryAttributesCommon for usage and error information.
func (c *NATGATEWAY) ModifySnatEntryAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifySnatEntryAttributesCommonRequest(input)
	return out, req.Send()
}

// ModifySnatEntryAttributesCommonWithContext is the same as ModifySnatEntryAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySnatEntryAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) ModifySnatEntryAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifySnatEntryAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifySnatEntryAttributes = "ModifySnatEntryAttributes"

// ModifySnatEntryAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifySnatEntryAttributes operation. The "output" return
// value will be populated with the ModifySnatEntryAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifySnatEntryAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifySnatEntryAttributesCommon Send returns without error.
//
// See ModifySnatEntryAttributes for more information on using the ModifySnatEntryAttributes
// API call, and error handling.
//
//	// Example sending a request using the ModifySnatEntryAttributesRequest method.
//	req, resp := client.ModifySnatEntryAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *NATGATEWAY) ModifySnatEntryAttributesRequest(input *ModifySnatEntryAttributesInput) (req *request.Request, output *ModifySnatEntryAttributesOutput) {
	op := &request.Operation{
		Name:       opModifySnatEntryAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifySnatEntryAttributesInput{}
	}

	output = &ModifySnatEntryAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ModifySnatEntryAttributes API operation for NATGATEWAY.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for NATGATEWAY's
// API operation ModifySnatEntryAttributes for usage and error information.
func (c *NATGATEWAY) ModifySnatEntryAttributes(input *ModifySnatEntryAttributesInput) (*ModifySnatEntryAttributesOutput, error) {
	req, out := c.ModifySnatEntryAttributesRequest(input)
	return out, req.Send()
}

// ModifySnatEntryAttributesWithContext is the same as ModifySnatEntryAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See ModifySnatEntryAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *NATGATEWAY) ModifySnatEntryAttributesWithContext(ctx volcengine.Context, input *ModifySnatEntryAttributesInput, opts ...request.Option) (*ModifySnatEntryAttributesOutput, error) {
	req, out := c.ModifySnatEntryAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ModifySnatEntryAttributesInput struct {
	_ struct{} `type:"structure"`

	EipId *string `type:"string"`

	// SnatEntryId is a required field
	SnatEntryId *string `type:"string" required:"true"`

	SnatEntryName *string `min:"1" max:"128" type:"string"`
}

// String returns the string representation
func (s ModifySnatEntryAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifySnatEntryAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifySnatEntryAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifySnatEntryAttributesInput"}
	if s.SnatEntryId == nil {
		invalidParams.Add(request.NewErrParamRequired("SnatEntryId"))
	}
	if s.SnatEntryName != nil && len(*s.SnatEntryName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("SnatEntryName", 1))
	}
	if s.SnatEntryName != nil && len(*s.SnatEntryName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("SnatEntryName", 128, *s.SnatEntryName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEipId sets the EipId field's value.
func (s *ModifySnatEntryAttributesInput) SetEipId(v string) *ModifySnatEntryAttributesInput {
	s.EipId = &v
	return s
}

// SetSnatEntryId sets the SnatEntryId field's value.
func (s *ModifySnatEntryAttributesInput) SetSnatEntryId(v string) *ModifySnatEntryAttributesInput {
	s.SnatEntryId = &v
	return s
}

// SetSnatEntryName sets the SnatEntryName field's value.
func (s *ModifySnatEntryAttributesInput) SetSnatEntryName(v string) *ModifySnatEntryAttributesInput {
	s.SnatEntryName = &v
	return s
}

type ModifySnatEntryAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ModifySnatEntryAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifySnatEntryAttributesOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *ModifySnatEntryAttributesOutput) SetRequestId(v string) *ModifySnatEntryAttributesOutput {
	s.RequestId = &v
	return s
}
