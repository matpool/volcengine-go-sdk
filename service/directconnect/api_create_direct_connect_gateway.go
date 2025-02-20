// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDirectConnectGatewayCommon = "CreateDirectConnectGateway"

// CreateDirectConnectGatewayCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDirectConnectGatewayCommon operation. The "output" return
// value will be populated with the CreateDirectConnectGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDirectConnectGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDirectConnectGatewayCommon Send returns without error.
//
// See CreateDirectConnectGatewayCommon for more information on using the CreateDirectConnectGatewayCommon
// API call, and error handling.
//
//	// Example sending a request using the CreateDirectConnectGatewayCommonRequest method.
//	req, resp := client.CreateDirectConnectGatewayCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) CreateDirectConnectGatewayCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDirectConnectGatewayCommon,
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

// CreateDirectConnectGatewayCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation CreateDirectConnectGatewayCommon for usage and error information.
func (c *DIRECTCONNECT) CreateDirectConnectGatewayCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDirectConnectGatewayCommonRequest(input)
	return out, req.Send()
}

// CreateDirectConnectGatewayCommonWithContext is the same as CreateDirectConnectGatewayCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDirectConnectGatewayCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) CreateDirectConnectGatewayCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDirectConnectGatewayCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDirectConnectGateway = "CreateDirectConnectGateway"

// CreateDirectConnectGatewayRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDirectConnectGateway operation. The "output" return
// value will be populated with the CreateDirectConnectGatewayCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDirectConnectGatewayCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDirectConnectGatewayCommon Send returns without error.
//
// See CreateDirectConnectGateway for more information on using the CreateDirectConnectGateway
// API call, and error handling.
//
//	// Example sending a request using the CreateDirectConnectGatewayRequest method.
//	req, resp := client.CreateDirectConnectGatewayRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) CreateDirectConnectGatewayRequest(input *CreateDirectConnectGatewayInput) (req *request.Request, output *CreateDirectConnectGatewayOutput) {
	op := &request.Operation{
		Name:       opCreateDirectConnectGateway,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDirectConnectGatewayInput{}
	}

	output = &CreateDirectConnectGatewayOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateDirectConnectGateway API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation CreateDirectConnectGateway for usage and error information.
func (c *DIRECTCONNECT) CreateDirectConnectGateway(input *CreateDirectConnectGatewayInput) (*CreateDirectConnectGatewayOutput, error) {
	req, out := c.CreateDirectConnectGatewayRequest(input)
	return out, req.Send()
}

// CreateDirectConnectGatewayWithContext is the same as CreateDirectConnectGateway with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDirectConnectGateway for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) CreateDirectConnectGatewayWithContext(ctx volcengine.Context, input *CreateDirectConnectGatewayInput, opts ...request.Option) (*CreateDirectConnectGatewayOutput, error) {
	req, out := c.CreateDirectConnectGatewayRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDirectConnectGatewayInput struct {
	_ struct{} `type:"structure"`

	ClientToken *string `type:"string"`

	Description *string `type:"string"`

	// DirectConnectGatewayName is a required field
	DirectConnectGatewayName *string `type:"string" required:"true"`

	Tags []*TagForCreateDirectConnectGatewayInput `type:"list"`
}

// String returns the string representation
func (s CreateDirectConnectGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDirectConnectGatewayInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDirectConnectGatewayInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDirectConnectGatewayInput"}
	if s.DirectConnectGatewayName == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectGatewayName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetClientToken sets the ClientToken field's value.
func (s *CreateDirectConnectGatewayInput) SetClientToken(v string) *CreateDirectConnectGatewayInput {
	s.ClientToken = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *CreateDirectConnectGatewayInput) SetDescription(v string) *CreateDirectConnectGatewayInput {
	s.Description = &v
	return s
}

// SetDirectConnectGatewayName sets the DirectConnectGatewayName field's value.
func (s *CreateDirectConnectGatewayInput) SetDirectConnectGatewayName(v string) *CreateDirectConnectGatewayInput {
	s.DirectConnectGatewayName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateDirectConnectGatewayInput) SetTags(v []*TagForCreateDirectConnectGatewayInput) *CreateDirectConnectGatewayInput {
	s.Tags = v
	return s
}

type CreateDirectConnectGatewayOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DirectConnectGatewayId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s CreateDirectConnectGatewayOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDirectConnectGatewayOutput) GoString() string {
	return s.String()
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *CreateDirectConnectGatewayOutput) SetDirectConnectGatewayId(v string) *CreateDirectConnectGatewayOutput {
	s.DirectConnectGatewayId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *CreateDirectConnectGatewayOutput) SetRequestId(v string) *CreateDirectConnectGatewayOutput {
	s.RequestId = &v
	return s
}

type TagForCreateDirectConnectGatewayInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForCreateDirectConnectGatewayInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForCreateDirectConnectGatewayInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForCreateDirectConnectGatewayInput) SetKey(v string) *TagForCreateDirectConnectGatewayInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForCreateDirectConnectGatewayInput) SetValue(v string) *TagForCreateDirectConnectGatewayInput {
	s.Value = &v
	return s
}
