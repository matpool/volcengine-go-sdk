// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opListInstanceParamsHistoryCommon = "ListInstanceParamsHistory"

// ListInstanceParamsHistoryCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ListInstanceParamsHistoryCommon operation. The "output" return
// value will be populated with the ListInstanceParamsHistoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListInstanceParamsHistoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListInstanceParamsHistoryCommon Send returns without error.
//
// See ListInstanceParamsHistoryCommon for more information on using the ListInstanceParamsHistoryCommon
// API call, and error handling.
//
//    // Example sending a request using the ListInstanceParamsHistoryCommonRequest method.
//    req, resp := client.ListInstanceParamsHistoryCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListInstanceParamsHistoryCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opListInstanceParamsHistoryCommon,
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

// ListInstanceParamsHistoryCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListInstanceParamsHistoryCommon for usage and error information.
func (c *RDSMYSQL) ListInstanceParamsHistoryCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ListInstanceParamsHistoryCommonRequest(input)
	return out, req.Send()
}

// ListInstanceParamsHistoryCommonWithContext is the same as ListInstanceParamsHistoryCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ListInstanceParamsHistoryCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListInstanceParamsHistoryCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ListInstanceParamsHistoryCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListInstanceParamsHistory = "ListInstanceParamsHistory"

// ListInstanceParamsHistoryRequest generates a "volcengine/request.Request" representing the
// client's request for the ListInstanceParamsHistory operation. The "output" return
// value will be populated with the ListInstanceParamsHistoryCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ListInstanceParamsHistoryCommon Request to send the API call to the service.
// the "output" return value is not valid until after ListInstanceParamsHistoryCommon Send returns without error.
//
// See ListInstanceParamsHistory for more information on using the ListInstanceParamsHistory
// API call, and error handling.
//
//    // Example sending a request using the ListInstanceParamsHistoryRequest method.
//    req, resp := client.ListInstanceParamsHistoryRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ListInstanceParamsHistoryRequest(input *ListInstanceParamsHistoryInput) (req *request.Request, output *ListInstanceParamsHistoryOutput) {
	op := &request.Operation{
		Name:       opListInstanceParamsHistory,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListInstanceParamsHistoryInput{}
	}

	output = &ListInstanceParamsHistoryOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ListInstanceParamsHistory API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ListInstanceParamsHistory for usage and error information.
func (c *RDSMYSQL) ListInstanceParamsHistory(input *ListInstanceParamsHistoryInput) (*ListInstanceParamsHistoryOutput, error) {
	req, out := c.ListInstanceParamsHistoryRequest(input)
	return out, req.Send()
}

// ListInstanceParamsHistoryWithContext is the same as ListInstanceParamsHistory with the addition of
// the ability to pass a context and additional request options.
//
// See ListInstanceParamsHistory for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ListInstanceParamsHistoryWithContext(ctx volcengine.Context, input *ListInstanceParamsHistoryInput, opts ...request.Option) (*ListInstanceParamsHistoryOutput, error) {
	req, out := c.ListInstanceParamsHistoryRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DataForListInstanceParamsHistoryOutput struct {
	_ struct{} `type:"structure"`

	ModifyTime *string `type:"string"`

	NewParameterValue *string `type:"string"`

	OldParameterValue *string `type:"string"`

	ParameterName *string `type:"string"`

	ParamsStatus *string `type:"string" enum:"EnumOfParamsStatusForListInstanceParamsHistoryOutput"`
}

// String returns the string representation
func (s DataForListInstanceParamsHistoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DataForListInstanceParamsHistoryOutput) GoString() string {
	return s.String()
}

// SetModifyTime sets the ModifyTime field's value.
func (s *DataForListInstanceParamsHistoryOutput) SetModifyTime(v string) *DataForListInstanceParamsHistoryOutput {
	s.ModifyTime = &v
	return s
}

// SetNewParameterValue sets the NewParameterValue field's value.
func (s *DataForListInstanceParamsHistoryOutput) SetNewParameterValue(v string) *DataForListInstanceParamsHistoryOutput {
	s.NewParameterValue = &v
	return s
}

// SetOldParameterValue sets the OldParameterValue field's value.
func (s *DataForListInstanceParamsHistoryOutput) SetOldParameterValue(v string) *DataForListInstanceParamsHistoryOutput {
	s.OldParameterValue = &v
	return s
}

// SetParameterName sets the ParameterName field's value.
func (s *DataForListInstanceParamsHistoryOutput) SetParameterName(v string) *DataForListInstanceParamsHistoryOutput {
	s.ParameterName = &v
	return s
}

// SetParamsStatus sets the ParamsStatus field's value.
func (s *DataForListInstanceParamsHistoryOutput) SetParamsStatus(v string) *DataForListInstanceParamsHistoryOutput {
	s.ParamsStatus = &v
	return s
}

type ListInstanceParamsHistoryInput struct {
	_ struct{} `type:"structure"`

	// EndTime is a required field
	EndTime *string `type:"string" required:"true"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	// StartTime is a required field
	StartTime *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListInstanceParamsHistoryInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListInstanceParamsHistoryInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListInstanceParamsHistoryInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListInstanceParamsHistoryInput"}
	if s.EndTime == nil {
		invalidParams.Add(request.NewErrParamRequired("EndTime"))
	}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.StartTime == nil {
		invalidParams.Add(request.NewErrParamRequired("StartTime"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEndTime sets the EndTime field's value.
func (s *ListInstanceParamsHistoryInput) SetEndTime(v string) *ListInstanceParamsHistoryInput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *ListInstanceParamsHistoryInput) SetInstanceId(v string) *ListInstanceParamsHistoryInput {
	s.InstanceId = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *ListInstanceParamsHistoryInput) SetStartTime(v string) *ListInstanceParamsHistoryInput {
	s.StartTime = &v
	return s
}

type ListInstanceParamsHistoryOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Datas []*DataForListInstanceParamsHistoryOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s ListInstanceParamsHistoryOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ListInstanceParamsHistoryOutput) GoString() string {
	return s.String()
}

// SetDatas sets the Datas field's value.
func (s *ListInstanceParamsHistoryOutput) SetDatas(v []*DataForListInstanceParamsHistoryOutput) *ListInstanceParamsHistoryOutput {
	s.Datas = v
	return s
}

// SetTotal sets the Total field's value.
func (s *ListInstanceParamsHistoryOutput) SetTotal(v int32) *ListInstanceParamsHistoryOutput {
	s.Total = &v
	return s
}

const (
	// EnumOfParamsStatusForListInstanceParamsHistoryOutputApplied is a EnumOfParamsStatusForListInstanceParamsHistoryOutput enum value
	EnumOfParamsStatusForListInstanceParamsHistoryOutputApplied = "Applied"

	// EnumOfParamsStatusForListInstanceParamsHistoryOutputFailed is a EnumOfParamsStatusForListInstanceParamsHistoryOutput enum value
	EnumOfParamsStatusForListInstanceParamsHistoryOutputFailed = "Failed"

	// EnumOfParamsStatusForListInstanceParamsHistoryOutputSyncing is a EnumOfParamsStatusForListInstanceParamsHistoryOutput enum value
	EnumOfParamsStatusForListInstanceParamsHistoryOutputSyncing = "Syncing"
)
