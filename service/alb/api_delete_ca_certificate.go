// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteCACertificateCommon = "DeleteCACertificate"

// DeleteCACertificateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCACertificateCommon operation. The "output" return
// value will be populated with the DeleteCACertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCACertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCACertificateCommon Send returns without error.
//
// See DeleteCACertificateCommon for more information on using the DeleteCACertificateCommon
// API call, and error handling.
//
//	// Example sending a request using the DeleteCACertificateCommonRequest method.
//	req, resp := client.DeleteCACertificateCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) DeleteCACertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteCACertificateCommon,
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

// DeleteCACertificateCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DeleteCACertificateCommon for usage and error information.
func (c *ALB) DeleteCACertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteCACertificateCommonRequest(input)
	return out, req.Send()
}

// DeleteCACertificateCommonWithContext is the same as DeleteCACertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCACertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DeleteCACertificateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteCACertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteCACertificate = "DeleteCACertificate"

// DeleteCACertificateRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteCACertificate operation. The "output" return
// value will be populated with the DeleteCACertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteCACertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteCACertificateCommon Send returns without error.
//
// See DeleteCACertificate for more information on using the DeleteCACertificate
// API call, and error handling.
//
//	// Example sending a request using the DeleteCACertificateRequest method.
//	req, resp := client.DeleteCACertificateRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) DeleteCACertificateRequest(input *DeleteCACertificateInput) (req *request.Request, output *DeleteCACertificateOutput) {
	op := &request.Operation{
		Name:       opDeleteCACertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteCACertificateInput{}
	}

	output = &DeleteCACertificateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteCACertificate API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation DeleteCACertificate for usage and error information.
func (c *ALB) DeleteCACertificate(input *DeleteCACertificateInput) (*DeleteCACertificateOutput, error) {
	req, out := c.DeleteCACertificateRequest(input)
	return out, req.Send()
}

// DeleteCACertificateWithContext is the same as DeleteCACertificate with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteCACertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) DeleteCACertificateWithContext(ctx volcengine.Context, input *DeleteCACertificateInput, opts ...request.Option) (*DeleteCACertificateOutput, error) {
	req, out := c.DeleteCACertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteCACertificateInput struct {
	_ struct{} `type:"structure"`

	// CACertificateId is a required field
	CACertificateId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteCACertificateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCACertificateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteCACertificateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteCACertificateInput"}
	if s.CACertificateId == nil {
		invalidParams.Add(request.NewErrParamRequired("CACertificateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCACertificateId sets the CACertificateId field's value.
func (s *DeleteCACertificateInput) SetCACertificateId(v string) *DeleteCACertificateInput {
	s.CACertificateId = &v
	return s
}

type DeleteCACertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s DeleteCACertificateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteCACertificateOutput) GoString() string {
	return s.String()
}

// SetRequestId sets the RequestId field's value.
func (s *DeleteCACertificateOutput) SetRequestId(v string) *DeleteCACertificateOutput {
	s.RequestId = &v
	return s
}
