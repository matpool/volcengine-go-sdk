// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package iam

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateUserCommon = "CreateUser"

// CreateUserCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateUserCommon operation. The "output" return
// value will be populated with the CreateUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateUserCommon Send returns without error.
//
// See CreateUserCommon for more information on using the CreateUserCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateUserCommonRequest method.
//    req, resp := client.CreateUserCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateUserCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateUserCommon,
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

// CreateUserCommon API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateUserCommon for usage and error information.
func (c *IAM) CreateUserCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateUserCommonRequest(input)
	return out, req.Send()
}

// CreateUserCommonWithContext is the same as CreateUserCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateUserCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateUserCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateUserCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateUser = "CreateUser"

// CreateUserRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateUser operation. The "output" return
// value will be populated with the CreateUserCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateUserCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateUserCommon Send returns without error.
//
// See CreateUser for more information on using the CreateUser
// API call, and error handling.
//
//    // Example sending a request using the CreateUserRequest method.
//    req, resp := client.CreateUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *IAM) CreateUserRequest(input *CreateUserInput) (req *request.Request, output *CreateUserOutput) {
	op := &request.Operation{
		Name:       opCreateUser,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateUserInput{}
	}

	output = &CreateUserOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateUser API operation for IAM.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for IAM's
// API operation CreateUser for usage and error information.
func (c *IAM) CreateUser(input *CreateUserInput) (*CreateUserOutput, error) {
	req, out := c.CreateUserRequest(input)
	return out, req.Send()
}

// CreateUserWithContext is the same as CreateUser with the addition of
// the ability to pass a context and additional request options.
//
// See CreateUser for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *IAM) CreateUserWithContext(ctx volcengine.Context, input *CreateUserInput, opts ...request.Option) (*CreateUserOutput, error) {
	req, out := c.CreateUserRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateUserInput struct {
	_ struct{} `type:"structure"`

	Description *string `type:"string"`

	DisplayName *string `type:"string"`

	Email *string `type:"string"`

	MobilePhone *string `type:"string"`

	// UserName is a required field
	UserName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateUserInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateUserInput"}
	if s.UserName == nil {
		invalidParams.Add(request.NewErrParamRequired("UserName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateUserInput) SetDescription(v string) *CreateUserInput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *CreateUserInput) SetDisplayName(v string) *CreateUserInput {
	s.DisplayName = &v
	return s
}

// SetEmail sets the Email field's value.
func (s *CreateUserInput) SetEmail(v string) *CreateUserInput {
	s.Email = &v
	return s
}

// SetMobilePhone sets the MobilePhone field's value.
func (s *CreateUserInput) SetMobilePhone(v string) *CreateUserInput {
	s.MobilePhone = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *CreateUserInput) SetUserName(v string) *CreateUserInput {
	s.UserName = &v
	return s
}

type CreateUserOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	User *UserForCreateUserOutput `type:"structure"`
}

// String returns the string representation
func (s CreateUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateUserOutput) GoString() string {
	return s.String()
}

// SetUser sets the User field's value.
func (s *CreateUserOutput) SetUser(v *UserForCreateUserOutput) *CreateUserOutput {
	s.User = v
	return s
}

type UserForCreateUserOutput struct {
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
func (s UserForCreateUserOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s UserForCreateUserOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *UserForCreateUserOutput) SetAccountId(v string) *UserForCreateUserOutput {
	s.AccountId = &v
	return s
}

// SetCreateDate sets the CreateDate field's value.
func (s *UserForCreateUserOutput) SetCreateDate(v string) *UserForCreateUserOutput {
	s.CreateDate = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *UserForCreateUserOutput) SetDescription(v string) *UserForCreateUserOutput {
	s.Description = &v
	return s
}

// SetDisplayName sets the DisplayName field's value.
func (s *UserForCreateUserOutput) SetDisplayName(v string) *UserForCreateUserOutput {
	s.DisplayName = &v
	return s
}

// SetEmail sets the Email field's value.
func (s *UserForCreateUserOutput) SetEmail(v string) *UserForCreateUserOutput {
	s.Email = &v
	return s
}

// SetEmailIsVerify sets the EmailIsVerify field's value.
func (s *UserForCreateUserOutput) SetEmailIsVerify(v string) *UserForCreateUserOutput {
	s.EmailIsVerify = &v
	return s
}

// SetMobilePhone sets the MobilePhone field's value.
func (s *UserForCreateUserOutput) SetMobilePhone(v string) *UserForCreateUserOutput {
	s.MobilePhone = &v
	return s
}

// SetMobilePhoneIsVerify sets the MobilePhoneIsVerify field's value.
func (s *UserForCreateUserOutput) SetMobilePhoneIsVerify(v string) *UserForCreateUserOutput {
	s.MobilePhoneIsVerify = &v
	return s
}

// SetTrn sets the Trn field's value.
func (s *UserForCreateUserOutput) SetTrn(v string) *UserForCreateUserOutput {
	s.Trn = &v
	return s
}

// SetUpdateDate sets the UpdateDate field's value.
func (s *UserForCreateUserOutput) SetUpdateDate(v string) *UserForCreateUserOutput {
	s.UpdateDate = &v
	return s
}

// SetUserName sets the UserName field's value.
func (s *UserForCreateUserOutput) SetUserName(v string) *UserForCreateUserOutput {
	s.UserName = &v
	return s
}
