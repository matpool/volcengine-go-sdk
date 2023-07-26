// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateAccountCommon = "CreateAccount"

// CreateAccountCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateAccountCommon operation. The "output" return
// value will be populated with the CreateAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateAccountCommon Send returns without error.
//
// See CreateAccountCommon for more information on using the CreateAccountCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateAccountCommonRequest method.
//    req, resp := client.CreateAccountCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) CreateAccountCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateAccountCommon,
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

// CreateAccountCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CreateAccountCommon for usage and error information.
func (c *RDSMYSQL) CreateAccountCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateAccountCommonRequest(input)
	return out, req.Send()
}

// CreateAccountCommonWithContext is the same as CreateAccountCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateAccountCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) CreateAccountCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateAccountCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateAccount = "CreateAccount"

// CreateAccountRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateAccount operation. The "output" return
// value will be populated with the CreateAccountCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateAccountCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateAccountCommon Send returns without error.
//
// See CreateAccount for more information on using the CreateAccount
// API call, and error handling.
//
//    // Example sending a request using the CreateAccountRequest method.
//    req, resp := client.CreateAccountRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) CreateAccountRequest(input *CreateAccountInput) (req *request.Request, output *CreateAccountOutput) {
	op := &request.Operation{
		Name:       opCreateAccount,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateAccountInput{}
	}

	output = &CreateAccountOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateAccount API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CreateAccount for usage and error information.
func (c *RDSMYSQL) CreateAccount(input *CreateAccountInput) (*CreateAccountOutput, error) {
	req, out := c.CreateAccountRequest(input)
	return out, req.Send()
}

// CreateAccountWithContext is the same as CreateAccount with the addition of
// the ability to pass a context and additional request options.
//
// See CreateAccount for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) CreateAccountWithContext(ctx volcengine.Context, input *CreateAccountInput, opts ...request.Option) (*CreateAccountOutput, error) {
	req, out := c.CreateAccountRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateAccountInput struct {
	_ struct{} `type:"structure"`

	AccountName *string `min:"2" max:"32" type:"string"`

	AccountPassword *string `min:"8" max:"32" type:"string"`

	AccountType *string `type:"string" enum:"EnumOfAccountTypeForCreateAccountInput"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateAccountInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateAccountInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateAccountInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateAccountInput"}
	if s.AccountName != nil && len(*s.AccountName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("AccountName", 2))
	}
	if s.AccountName != nil && len(*s.AccountName) > 32 {
		invalidParams.Add(request.NewErrParamMaxLen("AccountName", 32, *s.AccountName))
	}
	if s.AccountPassword != nil && len(*s.AccountPassword) < 8 {
		invalidParams.Add(request.NewErrParamMinLen("AccountPassword", 8))
	}
	if s.AccountPassword != nil && len(*s.AccountPassword) > 32 {
		invalidParams.Add(request.NewErrParamMaxLen("AccountPassword", 32, *s.AccountPassword))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAccountName sets the AccountName field's value.
func (s *CreateAccountInput) SetAccountName(v string) *CreateAccountInput {
	s.AccountName = &v
	return s
}

// SetAccountPassword sets the AccountPassword field's value.
func (s *CreateAccountInput) SetAccountPassword(v string) *CreateAccountInput {
	s.AccountPassword = &v
	return s
}

// SetAccountType sets the AccountType field's value.
func (s *CreateAccountInput) SetAccountType(v string) *CreateAccountInput {
	s.AccountType = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateAccountInput) SetInstanceId(v string) *CreateAccountInput {
	s.InstanceId = &v
	return s
}

type CreateAccountOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateAccountOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateAccountOutput) GoString() string {
	return s.String()
}

const (
	// EnumOfAccountTypeForCreateAccountInputGrant is a EnumOfAccountTypeForCreateAccountInput enum value
	EnumOfAccountTypeForCreateAccountInputGrant = "Grant"

	// EnumOfAccountTypeForCreateAccountInputNormal is a EnumOfAccountTypeForCreateAccountInput enum value
	EnumOfAccountTypeForCreateAccountInputNormal = "Normal"

	// EnumOfAccountTypeForCreateAccountInputSuper is a EnumOfAccountTypeForCreateAccountInput enum value
	EnumOfAccountTypeForCreateAccountInputSuper = "Super"
)
