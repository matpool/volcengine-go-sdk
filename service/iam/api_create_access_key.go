// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateAccessKeyCommon = "CreateAccessKey"

// CreateAccessKeyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateAccessKeyCommon operation. The "output" return
// value will be populated with the CreateAccessKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateAccessKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateAccessKeyCommon Send returns without error.
//
// See CreateAccessKeyCommon for more information on using the CreateAccessKeyCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateAccessKeyCommonRequest method.
//    req, resp := client.CreateAccessKeyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateAccessKeyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateAccessKeyCommon,
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

// CreateAccessKeyCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateAccessKeyCommon for usage and error information.
func (c *IAM) CreateAccessKeyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateAccessKeyCommonRequest(input)
	return out, req.Send()
}

// CreateAccessKeyCommonWithContext is the same as CreateAccessKeyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateAccessKeyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateAccessKeyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateAccessKeyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateAccessKey = "CreateAccessKey"

// CreateAccessKeyRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateAccessKey operation. The "output" return
// value will be populated with the CreateAccessKeyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateAccessKeyCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateAccessKeyCommon Send returns without error.
//
// See CreateAccessKey for more information on using the CreateAccessKey
// API call, and error handling.
//
//    // Example sending a request using the CreateAccessKeyRequest method.
//    req, resp := client.CreateAccessKeyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateAccessKeyRequest(input *CreateAccessKeyInput) (req *request.Request, output *CreateAccessKeyOutput) {
	op := &request.Operation{
		Name:       opCreateAccessKey,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAccessKeyInput{}
	}

	output = &CreateAccessKeyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateAccessKey API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateAccessKey for usage and error information.
func (c *IAM) CreateAccessKey(input *CreateAccessKeyInput) (*CreateAccessKeyOutput, error) {
	req, out := c.CreateAccessKeyRequest(input)
	return out, req.Send()
}

// CreateAccessKeyWithContext is the same as CreateAccessKey with the addition of
// the ability to pass a context and additional request options.
//
// See CreateAccessKey for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateAccessKeyWithContext(ctx volcengine.Context, input *CreateAccessKeyInput, opts ...request.Option) (*CreateAccessKeyOutput, error) {
	req, out := c.CreateAccessKeyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AccessKeyForCreateAccessKeyOutput struct {
	_ struct{} `type:"structure"`

	AccessKeyId *string `type:"string"`

	CreateDate *string `type:"string"`

	SecretAccessKey *string `type:"string"`

	Status *string `type:"string"`

	UpdateDate *string `type:"string"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s AccessKeyForCreateAccessKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AccessKeyForCreateAccessKeyOutput) GoString() string {
	return s.String()
}

// SetAccessKeyId sets the AccessKeyId field's value.
func (s *AccessKeyForCreateAccessKeyOutput) SetAccessKeyId(v string) *AccessKeyForCreateAccessKeyOutput {
	s.AccessKeyId = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *AccessKeyForCreateAccessKeyOutput) SetCreateDate(v string) *AccessKeyForCreateAccessKeyOutput {
	s.CreateDate = &v
	return s
}

// SetSecretAccessKey sets the SecretAccessKey field's value.
func (s *AccessKeyForCreateAccessKeyOutput) SetSecretAccessKey(v string) *AccessKeyForCreateAccessKeyOutput {
	s.SecretAccessKey = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *AccessKeyForCreateAccessKeyOutput) SetStatus(v string) *AccessKeyForCreateAccessKeyOutput {
	s.Status = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *AccessKeyForCreateAccessKeyOutput) SetUpdateDate(v string) *AccessKeyForCreateAccessKeyOutput {
	s.UpdateDate = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *AccessKeyForCreateAccessKeyOutput) SetUserName(v string) *AccessKeyForCreateAccessKeyOutput {
	s.UserName = &v
	return s
}

type CreateAccessKeyInput struct {
	_ struct{} `type:"structure"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s CreateAccessKeyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateAccessKeyInput) GoString() string {
	return s.String()
}

// SetUserName sets the UserName field's value.
func (s *CreateAccessKeyInput) SetUserName(v string) *CreateAccessKeyInput {
	s.UserName = &v
	return s
}

type CreateAccessKeyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccessKey *AccessKeyForCreateAccessKeyOutput `type:"structure"`
}

// String returns the string representation
func (s CreateAccessKeyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateAccessKeyOutput) GoString() string {
	return s.String()
}

// SetAccessKey sets the AccessKey field's value.
func (s *CreateAccessKeyOutput) SetAccessKey(v *AccessKeyForCreateAccessKeyOutput) *CreateAccessKeyOutput {
	s.AccessKey = v
	return s
}
