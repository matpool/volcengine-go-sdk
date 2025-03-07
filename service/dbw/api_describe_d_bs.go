// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dbw

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBsCommon = "DescribeDBs"

// DescribeDBsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBsCommon operation. The "output" return
// value will be populated with the DescribeDBsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBsCommon Send returns without error.
//
// See DescribeDBsCommon for more information on using the DescribeDBsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBsCommonRequest method.
//    req, resp := client.DescribeDBsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DescribeDBsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBsCommon,
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

// DescribeDBsCommon API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DescribeDBsCommon for usage and error information.
func (c *DBW) DescribeDBsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBsCommonRequest(input)
	return out, req.Send()
}

// DescribeDBsCommonWithContext is the same as DescribeDBsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DescribeDBsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBs = "DescribeDBs"

// DescribeDBsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBs operation. The "output" return
// value will be populated with the DescribeDBsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBsCommon Send returns without error.
//
// See DescribeDBs for more information on using the DescribeDBs
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBsRequest method.
//    req, resp := client.DescribeDBsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DescribeDBsRequest(input *DescribeDBsInput) (req *request.Request, output *DescribeDBsOutput) {
	op := &request.Operation{
		Name:       opDescribeDBs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBsInput{}
	}

	output = &DescribeDBsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBs API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DescribeDBs for usage and error information.
func (c *DBW) DescribeDBs(input *DescribeDBsInput) (*DescribeDBsOutput, error) {
	req, out := c.DescribeDBsRequest(input)
	return out, req.Send()
}

// DescribeDBsWithContext is the same as DescribeDBs with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DescribeDBsWithContext(ctx volcengine.Context, input *DescribeDBsInput, opts ...request.Option) (*DescribeDBsOutput, error) {
	req, out := c.DescribeDBsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDBsInput struct {
	_ struct{} `type:"structure"`

	DSType *string `type:"string" enum:"EnumOfDSTypeForDescribeDBsInput"`

	EndTime *int32 `type:"int32"`

	InstanceId *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	RegionId *string `type:"string"`

	SearchParam *SearchParamForDescribeDBsInput `type:"structure"`

	StartTime *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeDBsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBsInput) GoString() string {
	return s.String()
}

// SetDSType sets the DSType field's value.
func (s *DescribeDBsInput) SetDSType(v string) *DescribeDBsInput {
	s.DSType = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeDBsInput) SetEndTime(v int32) *DescribeDBsInput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBsInput) SetInstanceId(v string) *DescribeDBsInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeDBsInput) SetPageNumber(v int32) *DescribeDBsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeDBsInput) SetPageSize(v int32) *DescribeDBsInput {
	s.PageSize = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *DescribeDBsInput) SetRegionId(v string) *DescribeDBsInput {
	s.RegionId = &v
	return s
}

// SetSearchParam sets the SearchParam field's value.
func (s *DescribeDBsInput) SetSearchParam(v *SearchParamForDescribeDBsInput) *DescribeDBsInput {
	s.SearchParam = v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeDBsInput) SetStartTime(v int32) *DescribeDBsInput {
	s.StartTime = &v
	return s
}

type DescribeDBsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	DBs []*string `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeDBsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBsOutput) GoString() string {
	return s.String()
}

// SetDBs sets the DBs field's value.
func (s *DescribeDBsOutput) SetDBs(v []*string) *DescribeDBsOutput {
	s.DBs = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeDBsOutput) SetTotal(v int32) *DescribeDBsOutput {
	s.Total = &v
	return s
}

type SearchParamForDescribeDBsInput struct {
	_ struct{} `type:"structure"`

	DBs []*string `type:"list"`

	MaxQueryTime *float64 `type:"double"`

	MinQueryTime *float64 `type:"double"`

	SQLTemplate *string `type:"string"`

	SourceIPs []*string `type:"list"`

	Users []*string `type:"list"`
}

// String returns the string representation
func (s SearchParamForDescribeDBsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SearchParamForDescribeDBsInput) GoString() string {
	return s.String()
}

// SetDBs sets the DBs field's value.
func (s *SearchParamForDescribeDBsInput) SetDBs(v []*string) *SearchParamForDescribeDBsInput {
	s.DBs = v
	return s
}

// SetMaxQueryTime sets the MaxQueryTime field's value.
func (s *SearchParamForDescribeDBsInput) SetMaxQueryTime(v float64) *SearchParamForDescribeDBsInput {
	s.MaxQueryTime = &v
	return s
}

// SetMinQueryTime sets the MinQueryTime field's value.
func (s *SearchParamForDescribeDBsInput) SetMinQueryTime(v float64) *SearchParamForDescribeDBsInput {
	s.MinQueryTime = &v
	return s
}

// SetSQLTemplate sets the SQLTemplate field's value.
func (s *SearchParamForDescribeDBsInput) SetSQLTemplate(v string) *SearchParamForDescribeDBsInput {
	s.SQLTemplate = &v
	return s
}

// SetSourceIPs sets the SourceIPs field's value.
func (s *SearchParamForDescribeDBsInput) SetSourceIPs(v []*string) *SearchParamForDescribeDBsInput {
	s.SourceIPs = v
	return s
}

// SetUsers sets the Users field's value.
func (s *SearchParamForDescribeDBsInput) SetUsers(v []*string) *SearchParamForDescribeDBsInput {
	s.Users = v
	return s
}

const (
	// EnumOfDSTypeForDescribeDBsInputMongo is a EnumOfDSTypeForDescribeDBsInput enum value
	EnumOfDSTypeForDescribeDBsInputMongo = "Mongo"

	// EnumOfDSTypeForDescribeDBsInputMySql is a EnumOfDSTypeForDescribeDBsInput enum value
	EnumOfDSTypeForDescribeDBsInputMySql = "MySQL"

	// EnumOfDSTypeForDescribeDBsInputPostgres is a EnumOfDSTypeForDescribeDBsInput enum value
	EnumOfDSTypeForDescribeDBsInputPostgres = "Postgres"

	// EnumOfDSTypeForDescribeDBsInputRedis is a EnumOfDSTypeForDescribeDBsInput enum value
	EnumOfDSTypeForDescribeDBsInputRedis = "Redis"

	// EnumOfDSTypeForDescribeDBsInputVeDbmySql is a EnumOfDSTypeForDescribeDBsInput enum value
	EnumOfDSTypeForDescribeDBsInputVeDbmySql = "VeDBMySQL"
)
