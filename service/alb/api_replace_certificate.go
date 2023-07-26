// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package alb

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opReplaceCertificateCommon = "ReplaceCertificate"

// ReplaceCertificateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ReplaceCertificateCommon operation. The "output" return
// value will be populated with the ReplaceCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReplaceCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReplaceCertificateCommon Send returns without error.
//
// See ReplaceCertificateCommon for more information on using the ReplaceCertificateCommon
// API call, and error handling.
//
//	// Example sending a request using the ReplaceCertificateCommonRequest method.
//	req, resp := client.ReplaceCertificateCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) ReplaceCertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opReplaceCertificateCommon,
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

// ReplaceCertificateCommon API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ReplaceCertificateCommon for usage and error information.
func (c *ALB) ReplaceCertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ReplaceCertificateCommonRequest(input)
	return out, req.Send()
}

// ReplaceCertificateCommonWithContext is the same as ReplaceCertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ReplaceCertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ReplaceCertificateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ReplaceCertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opReplaceCertificate = "ReplaceCertificate"

// ReplaceCertificateRequest generates a "volcengine/request.Request" representing the
// client's request for the ReplaceCertificate operation. The "output" return
// value will be populated with the ReplaceCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ReplaceCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after ReplaceCertificateCommon Send returns without error.
//
// See ReplaceCertificate for more information on using the ReplaceCertificate
// API call, and error handling.
//
//	// Example sending a request using the ReplaceCertificateRequest method.
//	req, resp := client.ReplaceCertificateRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *ALB) ReplaceCertificateRequest(input *ReplaceCertificateInput) (req *request.Request, output *ReplaceCertificateOutput) {
	op := &request.Operation{
		Name:       opReplaceCertificate,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReplaceCertificateInput{}
	}

	output = &ReplaceCertificateOutput{}
	req = c.newRequest(op, input, output)

	return
}

// ReplaceCertificate API operation for ALB.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ALB's
// API operation ReplaceCertificate for usage and error information.
func (c *ALB) ReplaceCertificate(input *ReplaceCertificateInput) (*ReplaceCertificateOutput, error) {
	req, out := c.ReplaceCertificateRequest(input)
	return out, req.Send()
}

// ReplaceCertificateWithContext is the same as ReplaceCertificate with the addition of
// the ability to pass a context and additional request options.
//
// See ReplaceCertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ALB) ReplaceCertificateWithContext(ctx volcengine.Context, input *ReplaceCertificateInput, opts ...request.Option) (*ReplaceCertificateOutput, error) {
	req, out := c.ReplaceCertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ReplaceCertificateInput struct {
	_ struct{} `type:"structure"`

	CertificateId *string `type:"string"`

	CertificateName *string `min:"1" max:"128" type:"string"`

	Description *string `type:"string"`

	// OldCertificateId is a required field
	OldCertificateId *string `type:"string" required:"true"`

	PrivateKey *string `type:"string"`

	ProjectName *string `type:"string"`

	PublicKey *string `type:"string"`

	// UpdateMode is a required field
	UpdateMode *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ReplaceCertificateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReplaceCertificateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReplaceCertificateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ReplaceCertificateInput"}
	if s.CertificateName != nil && len(*s.CertificateName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("CertificateName", 1))
	}
	if s.CertificateName != nil && len(*s.CertificateName) > 128 {
		invalidParams.Add(request.NewErrParamMaxLen("CertificateName", 128, *s.CertificateName))
	}
	if s.OldCertificateId == nil {
		invalidParams.Add(request.NewErrParamRequired("OldCertificateId"))
	}
	if s.UpdateMode == nil {
		invalidParams.Add(request.NewErrParamRequired("UpdateMode"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCertificateId sets the CertificateId field's value.
func (s *ReplaceCertificateInput) SetCertificateId(v string) *ReplaceCertificateInput {
	s.CertificateId = &v
	return s
}

// SetCertificateName sets the CertificateName field's value.
func (s *ReplaceCertificateInput) SetCertificateName(v string) *ReplaceCertificateInput {
	s.CertificateName = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *ReplaceCertificateInput) SetDescription(v string) *ReplaceCertificateInput {
	s.Description = &v
	return s
}

// SetOldCertificateId sets the OldCertificateId field's value.
func (s *ReplaceCertificateInput) SetOldCertificateId(v string) *ReplaceCertificateInput {
	s.OldCertificateId = &v
	return s
}

// SetPrivateKey sets the PrivateKey field's value.
func (s *ReplaceCertificateInput) SetPrivateKey(v string) *ReplaceCertificateInput {
	s.PrivateKey = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *ReplaceCertificateInput) SetProjectName(v string) *ReplaceCertificateInput {
	s.ProjectName = &v
	return s
}

// SetPublicKey sets the PublicKey field's value.
func (s *ReplaceCertificateInput) SetPublicKey(v string) *ReplaceCertificateInput {
	s.PublicKey = &v
	return s
}

// SetUpdateMode sets the UpdateMode field's value.
func (s *ReplaceCertificateInput) SetUpdateMode(v string) *ReplaceCertificateInput {
	s.UpdateMode = &v
	return s
}

type ReplaceCertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	CertificateId *string `type:"string"`

	RequestId *string `type:"string"`
}

// String returns the string representation
func (s ReplaceCertificateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ReplaceCertificateOutput) GoString() string {
	return s.String()
}

// SetCertificateId sets the CertificateId field's value.
func (s *ReplaceCertificateOutput) SetCertificateId(v string) *ReplaceCertificateOutput {
	s.CertificateId = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *ReplaceCertificateOutput) SetRequestId(v string) *ReplaceCertificateOutput {
	s.RequestId = &v
	return s
}
