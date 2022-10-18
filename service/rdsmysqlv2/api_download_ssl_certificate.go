// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDownloadSSLCertificateCommon = "DownloadSSLCertificate"

// DownloadSSLCertificateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DownloadSSLCertificateCommon operation. The "output" return
// value will be populated with the DownloadSSLCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DownloadSSLCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DownloadSSLCertificateCommon Send returns without error.
//
// See DownloadSSLCertificateCommon for more information on using the DownloadSSLCertificateCommon
// API call, and error handling.
//
//    // Example sending a request using the DownloadSSLCertificateCommonRequest method.
//    req, resp := client.DownloadSSLCertificateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DownloadSSLCertificateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDownloadSSLCertificateCommon,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DownloadSSLCertificateCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DownloadSSLCertificateCommon for usage and error information.
func (c *RDSMYSQLV2) DownloadSSLCertificateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DownloadSSLCertificateCommonRequest(input)
	return out, req.Send()
}

// DownloadSSLCertificateCommonWithContext is the same as DownloadSSLCertificateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DownloadSSLCertificateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DownloadSSLCertificateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DownloadSSLCertificateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDownloadSSLCertificate = "DownloadSSLCertificate"

// DownloadSSLCertificateRequest generates a "volcengine/request.Request" representing the
// client's request for the DownloadSSLCertificate operation. The "output" return
// value will be populated with the DownloadSSLCertificateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DownloadSSLCertificateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DownloadSSLCertificateCommon Send returns without error.
//
// See DownloadSSLCertificate for more information on using the DownloadSSLCertificate
// API call, and error handling.
//
//    // Example sending a request using the DownloadSSLCertificateRequest method.
//    req, resp := client.DownloadSSLCertificateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DownloadSSLCertificateRequest(input *DownloadSSLCertificateInput) (req *request.Request, output *DownloadSSLCertificateOutput) {
	op := &request.Operation{
		Name:       opDownloadSSLCertificate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DownloadSSLCertificateInput{}
	}

	output = &DownloadSSLCertificateOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DownloadSSLCertificate API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DownloadSSLCertificate for usage and error information.
func (c *RDSMYSQLV2) DownloadSSLCertificate(input *DownloadSSLCertificateInput) (*DownloadSSLCertificateOutput, error) {
	req, out := c.DownloadSSLCertificateRequest(input)
	return out, req.Send()
}

// DownloadSSLCertificateWithContext is the same as DownloadSSLCertificate with the addition of
// the ability to pass a context and additional request options.
//
// See DownloadSSLCertificate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DownloadSSLCertificateWithContext(ctx volcengine.Context, input *DownloadSSLCertificateInput, opts ...request.Option) (*DownloadSSLCertificateOutput, error) {
	req, out := c.DownloadSSLCertificateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DownloadSSLCertificateInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DownloadSSLCertificateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DownloadSSLCertificateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DownloadSSLCertificateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DownloadSSLCertificateInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DownloadSSLCertificateInput) SetInstanceId(v string) *DownloadSSLCertificateInput {
	s.InstanceId = &v
	return s
}

type DownloadSSLCertificateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Certificate *string `type:"string"`
}

// String returns the string representation
func (s DownloadSSLCertificateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DownloadSSLCertificateOutput) GoString() string {
	return s.String()
}

// SetCertificate sets the Certificate field's value.
func (s *DownloadSSLCertificateOutput) SetCertificate(v string) *DownloadSSLCertificateOutput {
	s.Certificate = &v
	return s
}
