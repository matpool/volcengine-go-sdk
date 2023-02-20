// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opGetUserCommon = "GetUser"

// GetUserCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the GetUserCommon operation. The "output" return
// value will be populated with the GetUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetUserCommon Send returns without error.
//
// See GetUserCommon for more information on using the GetUserCommon
// API call, and error handling.
//
//    // Example sending a request using the GetUserCommonRequest method.
//    req, resp := client.GetUserCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetUserCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opGetUserCommon,
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

// GetUserCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetUserCommon for usage and error information.
func (c *IAM) GetUserCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.GetUserCommonRequest(input)
	return out, req.Send()
}

// GetUserCommonWithContext is the same as GetUserCommon with the addition of
// the ability to pass a context and additional request options.
//
// See GetUserCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetUserCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.GetUserCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetUser = "GetUser"

// GetUserRequest generates a "volcengine/request.Request" representing the
// client's request for the GetUser operation. The "output" return
// value will be populated with the GetUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned GetUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after GetUserCommon Send returns without error.
//
// See GetUser for more information on using the GetUser
// API call, and error handling.
//
//    // Example sending a request using the GetUserRequest method.
//    req, resp := client.GetUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) GetUserRequest(input *GetUserInput) (req *request.Request, output *GetUserOutput) {
	op := &request.Operation{
		Name:       opGetUser,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetUserInput{}
	}

	output = &GetUserOutput{}
	req = c.newRequest(op, input, output)

	return
}

// GetUser API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation GetUser for usage and error information.
func (c *IAM) GetUser(input *GetUserInput) (*GetUserOutput, error) {
	req, out := c.GetUserRequest(input)
	return out, req.Send()
}

// GetUserWithContext is the same as GetUser with the addition of
// the ability to pass a context and additional request options.
//
// See GetUser for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) GetUserWithContext(ctx volcengine.Context, input *GetUserInput, opts ...request.Option) (*GetUserOutput, error) {
	req, out := c.GetUserRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type GetUserInput struct {
	_ struct{} `type:"structure"`

	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetUserInput"}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetUserName sets the UserName field's value.
func (s *GetUserInput) SetUserName(v string) *GetUserInput {
	s.UserName = &v
	return s
}

type GetUserOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	User *UserForGetUserOutput `type:"structure"`
}

// String returns the string representation
func (s GetUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s GetUserOutput) GoString() string {
	return s.String()
}

// SetUser sets the User field's value.
func (s *GetUserOutput) SetUser(v *UserForGetUserOutput) *GetUserOutput {
	s.User = v
	return s
}

type UserForGetUserOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	CreateDate *string `type:"string"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	Email *string `type:"string"`

	EmailIsVerify *string `type:"string"`

	MobilePhone *string `type:"string"`

	MobilePhoneIsVerify *string `type:"string"`

	Trn *string `type:"string"`

	UpdateDate *string `type:"string"`

	UserName *string `type:"string"`
}

// String returns the string representation
func (s UserForGetUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UserForGetUserOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *UserForGetUserOutput) SetAccountId(v string) *UserForGetUserOutput {
	s.AccountId = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *UserForGetUserOutput) SetCreateDate(v string) *UserForGetUserOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UserForGetUserOutput) SetDescription(v string) *UserForGetUserOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *UserForGetUserOutput) SetDisplayName(v string) *UserForGetUserOutput {
	s.DisplayName = &v
	return s
}

// SetEmail sets the Email field's value.
func (s *UserForGetUserOutput) SetEmail(v string) *UserForGetUserOutput {
	s.Email = &v
	return s
}

// SetEmailIsVerify sets the EmailIsVerify field's value.
func (s *UserForGetUserOutput) SetEmailIsVerify(v string) *UserForGetUserOutput {
	s.EmailIsVerify = &v
	return s
}

// SetMobilePhone sets the MobilePhone field's value.
func (s *UserForGetUserOutput) SetMobilePhone(v string) *UserForGetUserOutput {
	s.MobilePhone = &v
	return s
}

// SetMobilePhoneIsVerify sets the MobilePhoneIsVerify field's value.
func (s *UserForGetUserOutput) SetMobilePhoneIsVerify(v string) *UserForGetUserOutput {
	s.MobilePhoneIsVerify = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *UserForGetUserOutput) SetTrn(v string) *UserForGetUserOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *UserForGetUserOutput) SetUpdateDate(v string) *UserForGetUserOutput {
	s.UpdateDate = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *UserForGetUserOutput) SetUserName(v string) *UserForGetUserOutput {
	s.UserName = &v
	return s
}
