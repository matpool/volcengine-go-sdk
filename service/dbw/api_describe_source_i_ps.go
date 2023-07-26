// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package dbw

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSourceIPsCommon = "DescribeSourceIPs"

// DescribeSourceIPsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSourceIPsCommon operation. The "output" return
// value will be populated with the DescribeSourceIPsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSourceIPsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSourceIPsCommon Send returns without error.
//
// See DescribeSourceIPsCommon for more information on using the DescribeSourceIPsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSourceIPsCommonRequest method.
//    req, resp := client.DescribeSourceIPsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DescribeSourceIPsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSourceIPsCommon,
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

// DescribeSourceIPsCommon API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DescribeSourceIPsCommon for usage and error information.
func (c *DBW) DescribeSourceIPsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSourceIPsCommonRequest(input)
	return out, req.Send()
}

// DescribeSourceIPsCommonWithContext is the same as DescribeSourceIPsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSourceIPsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DescribeSourceIPsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSourceIPsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSourceIPs = "DescribeSourceIPs"

// DescribeSourceIPsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSourceIPs operation. The "output" return
// value will be populated with the DescribeSourceIPsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSourceIPsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSourceIPsCommon Send returns without error.
//
// See DescribeSourceIPs for more information on using the DescribeSourceIPs
// API call, and error handling.
//
//    // Example sending a request using the DescribeSourceIPsRequest method.
//    req, resp := client.DescribeSourceIPsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DBW) DescribeSourceIPsRequest(input *DescribeSourceIPsInput) (req *request.Request, output *DescribeSourceIPsOutput) {
	op := &request.Operation{
		Name:       opDescribeSourceIPs,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSourceIPsInput{}
	}

	output = &DescribeSourceIPsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeSourceIPs API operation for DBW.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DBW's
// API operation DescribeSourceIPs for usage and error information.
func (c *DBW) DescribeSourceIPs(input *DescribeSourceIPsInput) (*DescribeSourceIPsOutput, error) {
	req, out := c.DescribeSourceIPsRequest(input)
	return out, req.Send()
}

// DescribeSourceIPsWithContext is the same as DescribeSourceIPs with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSourceIPs for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DBW) DescribeSourceIPsWithContext(ctx volcengine.Context, input *DescribeSourceIPsInput, opts ...request.Option) (*DescribeSourceIPsOutput, error) {
	req, out := c.DescribeSourceIPsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSourceIPsInput struct {
	_ struct{} `type:"structure"`

	DSType *string `type:"string" enum:"EnumOfDSTypeForDescribeSourceIPsInput"`

	EndTime *int32 `type:"int32"`

	InstanceId *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	RegionId *string `type:"string"`

	SearchParam *SearchParamForDescribeSourceIPsInput `type:"structure"`

	StartTime *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeSourceIPsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSourceIPsInput) GoString() string {
	return s.String()
}

// SetDSType sets the DSType field's value.
func (s *DescribeSourceIPsInput) SetDSType(v string) *DescribeSourceIPsInput {
	s.DSType = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeSourceIPsInput) SetEndTime(v int32) *DescribeSourceIPsInput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeSourceIPsInput) SetInstanceId(v string) *DescribeSourceIPsInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSourceIPsInput) SetPageNumber(v int32) *DescribeSourceIPsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSourceIPsInput) SetPageSize(v int32) *DescribeSourceIPsInput {
	s.PageSize = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *DescribeSourceIPsInput) SetRegionId(v string) *DescribeSourceIPsInput {
	s.RegionId = &v
	return s
}

// SetSearchParam sets the SearchParam field's value.
func (s *DescribeSourceIPsInput) SetSearchParam(v *SearchParamForDescribeSourceIPsInput) *DescribeSourceIPsInput {
	s.SearchParam = v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeSourceIPsInput) SetStartTime(v int32) *DescribeSourceIPsInput {
	s.StartTime = &v
	return s
}

type DescribeSourceIPsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	SourceIPs []*string `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeSourceIPsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSourceIPsOutput) GoString() string {
	return s.String()
}

// SetSourceIPs sets the SourceIPs field's value.
func (s *DescribeSourceIPsOutput) SetSourceIPs(v []*string) *DescribeSourceIPsOutput {
	s.SourceIPs = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeSourceIPsOutput) SetTotal(v int32) *DescribeSourceIPsOutput {
	s.Total = &v
	return s
}

type SearchParamForDescribeSourceIPsInput struct {
	_ struct{} `type:"structure"`

	DBs []*string `type:"list"`

	MaxQueryTime *float64 `type:"double"`

	MinQueryTime *float64 `type:"double"`

	SQLTemplate *string `type:"string"`

	SourceIPs []*string `type:"list"`

	Users []*string `type:"list"`
}

// String returns the string representation
func (s SearchParamForDescribeSourceIPsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SearchParamForDescribeSourceIPsInput) GoString() string {
	return s.String()
}

// SetDBs sets the DBs field's value.
func (s *SearchParamForDescribeSourceIPsInput) SetDBs(v []*string) *SearchParamForDescribeSourceIPsInput {
	s.DBs = v
	return s
}

// SetMaxQueryTime sets the MaxQueryTime field's value.
func (s *SearchParamForDescribeSourceIPsInput) SetMaxQueryTime(v float64) *SearchParamForDescribeSourceIPsInput {
	s.MaxQueryTime = &v
	return s
}

// SetMinQueryTime sets the MinQueryTime field's value.
func (s *SearchParamForDescribeSourceIPsInput) SetMinQueryTime(v float64) *SearchParamForDescribeSourceIPsInput {
	s.MinQueryTime = &v
	return s
}

// SetSQLTemplate sets the SQLTemplate field's value.
func (s *SearchParamForDescribeSourceIPsInput) SetSQLTemplate(v string) *SearchParamForDescribeSourceIPsInput {
	s.SQLTemplate = &v
	return s
}

// SetSourceIPs sets the SourceIPs field's value.
func (s *SearchParamForDescribeSourceIPsInput) SetSourceIPs(v []*string) *SearchParamForDescribeSourceIPsInput {
	s.SourceIPs = v
	return s
}

// SetUsers sets the Users field's value.
func (s *SearchParamForDescribeSourceIPsInput) SetUsers(v []*string) *SearchParamForDescribeSourceIPsInput {
	s.Users = v
	return s
}

const (
	// EnumOfDSTypeForDescribeSourceIPsInputMongo is a EnumOfDSTypeForDescribeSourceIPsInput enum value
	EnumOfDSTypeForDescribeSourceIPsInputMongo = "Mongo"

	// EnumOfDSTypeForDescribeSourceIPsInputMySql is a EnumOfDSTypeForDescribeSourceIPsInput enum value
	EnumOfDSTypeForDescribeSourceIPsInputMySql = "MySQL"

	// EnumOfDSTypeForDescribeSourceIPsInputPostgres is a EnumOfDSTypeForDescribeSourceIPsInput enum value
	EnumOfDSTypeForDescribeSourceIPsInputPostgres = "Postgres"

	// EnumOfDSTypeForDescribeSourceIPsInputRedis is a EnumOfDSTypeForDescribeSourceIPsInput enum value
	EnumOfDSTypeForDescribeSourceIPsInputRedis = "Redis"

	// EnumOfDSTypeForDescribeSourceIPsInputVeDbmySql is a EnumOfDSTypeForDescribeSourceIPsInput enum value
	EnumOfDSTypeForDescribeSourceIPsInputVeDbmySql = "VeDBMySQL"
)
