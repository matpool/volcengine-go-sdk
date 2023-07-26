// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDBInstanceIPListCommon = "CreateDBInstanceIPList"

// CreateDBInstanceIPListCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBInstanceIPListCommon operation. The "output" return
// value will be populated with the CreateDBInstanceIPListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBInstanceIPListCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBInstanceIPListCommon Send returns without error.
//
// See CreateDBInstanceIPListCommon for more information on using the CreateDBInstanceIPListCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDBInstanceIPListCommonRequest method.
//    req, resp := client.CreateDBInstanceIPListCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) CreateDBInstanceIPListCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDBInstanceIPListCommon,
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

// CreateDBInstanceIPListCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CreateDBInstanceIPListCommon for usage and error information.
func (c *RDSMYSQL) CreateDBInstanceIPListCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDBInstanceIPListCommonRequest(input)
	return out, req.Send()
}

// CreateDBInstanceIPListCommonWithContext is the same as CreateDBInstanceIPListCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBInstanceIPListCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) CreateDBInstanceIPListCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDBInstanceIPListCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDBInstanceIPList = "CreateDBInstanceIPList"

// CreateDBInstanceIPListRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBInstanceIPList operation. The "output" return
// value will be populated with the CreateDBInstanceIPListCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBInstanceIPListCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBInstanceIPListCommon Send returns without error.
//
// See CreateDBInstanceIPList for more information on using the CreateDBInstanceIPList
// API call, and error handling.
//
//    // Example sending a request using the CreateDBInstanceIPListRequest method.
//    req, resp := client.CreateDBInstanceIPListRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) CreateDBInstanceIPListRequest(input *CreateDBInstanceIPListInput) (req *request.Request, output *CreateDBInstanceIPListOutput) {
	op := &request.Operation{
		Name:       opCreateDBInstanceIPList,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBInstanceIPListInput{}
	}

	output = &CreateDBInstanceIPListOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateDBInstanceIPList API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation CreateDBInstanceIPList for usage and error information.
func (c *RDSMYSQL) CreateDBInstanceIPList(input *CreateDBInstanceIPListInput) (*CreateDBInstanceIPListOutput, error) {
	req, out := c.CreateDBInstanceIPListRequest(input)
	return out, req.Send()
}

// CreateDBInstanceIPListWithContext is the same as CreateDBInstanceIPList with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBInstanceIPList for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) CreateDBInstanceIPListWithContext(ctx volcengine.Context, input *CreateDBInstanceIPListInput, opts ...request.Option) (*CreateDBInstanceIPListOutput, error) {
	req, out := c.CreateDBInstanceIPListRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CreateDBInstanceIPListInput struct {
	_ struct{} `type:"structure"`

	GroupName *string `min:"2" max:"120" type:"string"`

	IPList []*string `type:"list"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateDBInstanceIPListInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBInstanceIPListInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBInstanceIPListInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDBInstanceIPListInput"}
	if s.GroupName != nil && len(*s.GroupName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("GroupName", 2))
	}
	if s.GroupName != nil && len(*s.GroupName) > 120 {
		invalidParams.Add(request.NewErrParamMaxLen("GroupName", 120, *s.GroupName))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetGroupName sets the GroupName field's value.
func (s *CreateDBInstanceIPListInput) SetGroupName(v string) *CreateDBInstanceIPListInput {
	s.GroupName = &v
	return s
}

// SetIPList sets the IPList field's value.
func (s *CreateDBInstanceIPListInput) SetIPList(v []*string) *CreateDBInstanceIPListInput {
	s.IPList = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDBInstanceIPListInput) SetInstanceId(v string) *CreateDBInstanceIPListInput {
	s.InstanceId = &v
	return s
}

type CreateDBInstanceIPListOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s CreateDBInstanceIPListOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBInstanceIPListOutput) GoString() string {
	return s.String()
}
