// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opUploadCACertificateCommon = "UploadCACertificate"

// UploadCACertificateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the UploadCACertificateCommon operation. The "output" return
// value will be populated with the UploadCACertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UploadCACertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after UploadCACertificateCommon Send returns without error.
//
// See UploadCACertificateCommon for more information on using the UploadCACertificateCommon
// API call, and error handling.
//
//    // Example sending a request using the UploadCACertificateCommonRequest method.
//    req, resp := client.UploadCACertificateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) UploadCACertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opUploadCACertificateCommon,
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

// UploadCACertificateCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation UploadCACertificateCommon for usage and error information.
func (c *ALB) UploadCACertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.UploadCACertificateCommonRequest(input)
	return out, req.Send()
}

// UploadCACertificateCommonWithContext is the same as UploadCACertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See UploadCACertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) UploadCACertificateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.UploadCACertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUploadCACertificate = "UploadCACertificate"

// UploadCACertificateRequest generates a "volcengine/request.Request" representing the
// client's request for the UploadCACertificate operation. The "output" return
// value will be populated with the UploadCACertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned UploadCACertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after UploadCACertificateCommon Send returns without error.
//
// See UploadCACertificate for more information on using the UploadCACertificate
// API call, and error handling.
//
//    // Example sending a request using the UploadCACertificateRequest method.
//    req, resp := client.UploadCACertificateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ALB) UploadCACertificateRequest(input *UploadCACertificateInput) (req *request.Request, output *UploadCACertificateOutput) {
	op := &request.Operation{
		Name:       opUploadCACertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UploadCACertificateInput{}
	}

	output = &UploadCACertificateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// UploadCACertificate API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation UploadCACertificate for usage and error information.
func (c *ALB) UploadCACertificate(input *UploadCACertificateInput) (*UploadCACertificateOutput, error) {
	req, out := c.UploadCACertificateRequest(input)
	return out, req.Send()
}

// UploadCACertificateWithContext is the same as UploadCACertificate with the addition of
// the ability to pass a context and additional request options.
//
// See UploadCACertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) UploadCACertificateWithContext(ctx volcengine.Context, input *UploadCACertificateInput, opts ...request.Option) (*UploadCACertificateOutput, error) {
	req, out := c.UploadCACertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type UploadCACertificateInput struct {
	_ struct{} `type:"structure"`

	// CACertificate is a required field
	CACertificate *string `type:"string" required:"true"`

	CACertificateName *string `min:"1" max:"128" type:"string"`

	Description *string `type:"string"`
}

// String returns the string representation
func (s UploadCACertificateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UploadCACertificateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UploadCACertificateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UploadCACertificateInput"}
	if s.CACertificate == nil {
		invalidParams.Add(request.NewErrParamRequired("CACertificate"))
	}
	if s.CACertificateName != nil && len(*s.CACertificateName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CACertificateName", 1))
	}
	if s.CACertificateName != nil && len(*s.CACertificateName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("CACertificateName", 128, *s.CACertificateName))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCACertificate sets the CACertificate field's value.
func (s *UploadCACertificateInput) SetCACertificate(v string) *UploadCACertificateInput {
	s.CACertificate = &v
	return s
}

// SetCACertificateName sets the CACertificateName field's value.
func (s *UploadCACertificateInput) SetCACertificateName(v string) *UploadCACertificateInput {
	s.CACertificateName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UploadCACertificateInput) SetDescription(v string) *UploadCACertificateInput {
	s.Description = &v
	return s
}

type UploadCACertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CACertificateId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s UploadCACertificateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UploadCACertificateOutput) GoString() string {
	return s.String()
}

// SetCACertificateId sets the CACertificateId field's value.
func (s *UploadCACertificateOutput) SetCACertificateId(v string) *UploadCACertificateOutput {
	s.CACertificateId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *UploadCACertificateOutput) SetRequestId(v string) *UploadCACertificateOutput {
	s.RequestId = &v
	return s
}
