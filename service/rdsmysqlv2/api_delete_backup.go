// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteBackupCommon = "DeleteBackup"

// DeleteBackupCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteBackupCommon operation. The "output" return
// value will be populated with the DeleteBackupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteBackupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteBackupCommon Send returns without error.
//
// See DeleteBackupCommon for more information on using the DeleteBackupCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteBackupCommonRequest method.
//    req, resp := client.DeleteBackupCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DeleteBackupCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteBackupCommon,
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

// DeleteBackupCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DeleteBackupCommon for usage and error information.
func (c *RDSMYSQLV2) DeleteBackupCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteBackupCommonRequest(input)
	return out, req.Send()
}

// DeleteBackupCommonWithContext is the same as DeleteBackupCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteBackupCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DeleteBackupCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteBackupCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteBackup = "DeleteBackup"

// DeleteBackupRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteBackup operation. The "output" return
// value will be populated with the DeleteBackupCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteBackupCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteBackupCommon Send returns without error.
//
// See DeleteBackup for more information on using the DeleteBackup
// API call, and error handling.
//
//    // Example sending a request using the DeleteBackupRequest method.
//    req, resp := client.DeleteBackupRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) DeleteBackupRequest(input *DeleteBackupInput) (req *request.Request, output *DeleteBackupOutput) {
	op := &request.Operation{
		Name:       opDeleteBackup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteBackupInput{}
	}

	output = &DeleteBackupOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DeleteBackup API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation DeleteBackup for usage and error information.
func (c *RDSMYSQLV2) DeleteBackup(input *DeleteBackupInput) (*DeleteBackupOutput, error) {
	req, out := c.DeleteBackupRequest(input)
	return out, req.Send()
}

// DeleteBackupWithContext is the same as DeleteBackup with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteBackup for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) DeleteBackupWithContext(ctx volcengine.Context, input *DeleteBackupInput, opts ...request.Option) (*DeleteBackupOutput, error) {
	req, out := c.DeleteBackupRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteBackupInput struct {
	_ struct{} `type:"structure"`

	// BackupId is a required field
	BackupId *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteBackupInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteBackupInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteBackupInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteBackupInput"}
	if s.BackupId == nil {
		invalidParams.Add(request.NewErrParamRequired("BackupId"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBackupId sets the BackupId field's value.
func (s *DeleteBackupInput) SetBackupId(v string) *DeleteBackupInput {
	s.BackupId = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DeleteBackupInput) SetInstanceId(v string) *DeleteBackupInput {
	s.InstanceId = &v
	return s
}

type DeleteBackupOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata
}

// String returns the string representation
func (s DeleteBackupOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteBackupOutput) GoString() string {
	return s.String()
}
