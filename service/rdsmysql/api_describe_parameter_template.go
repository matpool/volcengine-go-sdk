// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeParameterTemplateCommon = "DescribeParameterTemplate"

// DescribeParameterTemplateCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeParameterTemplateCommon operation. The "output" return
// value will be populated with the DescribeParameterTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeParameterTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeParameterTemplateCommon Send returns without error.
//
// See DescribeParameterTemplateCommon for more information on using the DescribeParameterTemplateCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeParameterTemplateCommonRequest method.
//    req, resp := client.DescribeParameterTemplateCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) DescribeParameterTemplateCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeParameterTemplateCommon,
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

// DescribeParameterTemplateCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation DescribeParameterTemplateCommon for usage and error information.
func (c *RDSMYSQL) DescribeParameterTemplateCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeParameterTemplateCommonRequest(input)
	return out, req.Send()
}

// DescribeParameterTemplateCommonWithContext is the same as DescribeParameterTemplateCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeParameterTemplateCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) DescribeParameterTemplateCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeParameterTemplateCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeParameterTemplate = "DescribeParameterTemplate"

// DescribeParameterTemplateRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeParameterTemplate operation. The "output" return
// value will be populated with the DescribeParameterTemplateCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeParameterTemplateCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeParameterTemplateCommon Send returns without error.
//
// See DescribeParameterTemplate for more information on using the DescribeParameterTemplate
// API call, and error handling.
//
//    // Example sending a request using the DescribeParameterTemplateRequest method.
//    req, resp := client.DescribeParameterTemplateRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) DescribeParameterTemplateRequest(input *DescribeParameterTemplateInput) (req *request.Request, output *DescribeParameterTemplateOutput) {
	op := &request.Operation{
		Name:       opDescribeParameterTemplate,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeParameterTemplateInput{}
	}

	output = &DescribeParameterTemplateOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeParameterTemplate API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation DescribeParameterTemplate for usage and error information.
func (c *RDSMYSQL) DescribeParameterTemplate(input *DescribeParameterTemplateInput) (*DescribeParameterTemplateOutput, error) {
	req, out := c.DescribeParameterTemplateRequest(input)
	return out, req.Send()
}

// DescribeParameterTemplateWithContext is the same as DescribeParameterTemplate with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeParameterTemplate for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) DescribeParameterTemplateWithContext(ctx volcengine.Context, input *DescribeParameterTemplateInput, opts ...request.Option) (*DescribeParameterTemplateOutput, error) {
	req, out := c.DescribeParameterTemplateRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeParameterTemplateInput struct {
	_ struct{} `type:"structure"`

	// TemplateId is a required field
	TemplateId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeParameterTemplateInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeParameterTemplateInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeParameterTemplateInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeParameterTemplateInput"}
	if s.TemplateId == nil {
		invalidParams.Add(request.NewErrParamRequired("TemplateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetTemplateId sets the TemplateId field's value.
func (s *DescribeParameterTemplateInput) SetTemplateId(v string) *DescribeParameterTemplateInput {
	s.TemplateId = &v
	return s
}

type DescribeParameterTemplateOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	TemplateInfo *TemplateInfoForDescribeParameterTemplateOutput `type:"structure"`
}

// String returns the string representation
func (s DescribeParameterTemplateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeParameterTemplateOutput) GoString() string {
	return s.String()
}

// SetTemplateInfo sets the TemplateInfo field's value.
func (s *DescribeParameterTemplateOutput) SetTemplateInfo(v *TemplateInfoForDescribeParameterTemplateOutput) *DescribeParameterTemplateOutput {
	s.TemplateInfo = v
	return s
}

type TemplateInfoForDescribeParameterTemplateOutput struct {
	_ struct{} `type:"structure"`

	AccountId *string `type:"string"`

	CreateTime *string `type:"string"`

	NeedRestart *bool `type:"boolean"`

	ParameterNum *int64 `type:"int64"`

	TemplateCategory *string `type:"string" enum:"EnumOfTemplateCategoryForDescribeParameterTemplateOutput"`

	TemplateDesc *string `type:"string"`

	TemplateId *string `type:"string"`

	TemplateName *string `type:"string"`

	TemplateParams []*TemplateParamForDescribeParameterTemplateOutput `type:"list"`

	TemplateSource *string `type:"string" enum:"EnumOfTemplateSourceForDescribeParameterTemplateOutput"`

	TemplateType *string `type:"string" enum:"EnumOfTemplateTypeForDescribeParameterTemplateOutput"`

	TemplateTypeVersion *string `type:"string" enum:"EnumOfTemplateTypeVersionForDescribeParameterTemplateOutput"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s TemplateInfoForDescribeParameterTemplateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TemplateInfoForDescribeParameterTemplateOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetAccountId(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.AccountId = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetCreateTime(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.CreateTime = &v
	return s
}

// SetNeedRestart sets the NeedRestart field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetNeedRestart(v bool) *TemplateInfoForDescribeParameterTemplateOutput {
	s.NeedRestart = &v
	return s
}

// SetParameterNum sets the ParameterNum field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetParameterNum(v int64) *TemplateInfoForDescribeParameterTemplateOutput {
	s.ParameterNum = &v
	return s
}

// SetTemplateCategory sets the TemplateCategory field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetTemplateCategory(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.TemplateCategory = &v
	return s
}

// SetTemplateDesc sets the TemplateDesc field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetTemplateDesc(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.TemplateDesc = &v
	return s
}

// SetTemplateId sets the TemplateId field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetTemplateId(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.TemplateId = &v
	return s
}

// SetTemplateName sets the TemplateName field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetTemplateName(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.TemplateName = &v
	return s
}

// SetTemplateParams sets the TemplateParams field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetTemplateParams(v []*TemplateParamForDescribeParameterTemplateOutput) *TemplateInfoForDescribeParameterTemplateOutput {
	s.TemplateParams = v
	return s
}

// SetTemplateSource sets the TemplateSource field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetTemplateSource(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.TemplateSource = &v
	return s
}

// SetTemplateType sets the TemplateType field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetTemplateType(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.TemplateType = &v
	return s
}

// SetTemplateTypeVersion sets the TemplateTypeVersion field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetTemplateTypeVersion(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.TemplateTypeVersion = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *TemplateInfoForDescribeParameterTemplateOutput) SetUpdateTime(v string) *TemplateInfoForDescribeParameterTemplateOutput {
	s.UpdateTime = &v
	return s
}

type TemplateParamForDescribeParameterTemplateOutput struct {
	_ struct{} `type:"structure"`

	DefaultValue *string `type:"string"`

	Description *string `type:"string"`

	ExpectValue *string `type:"string"`

	Name *string `type:"string"`

	Restart *bool `type:"boolean"`

	RunningValue *string `type:"string"`

	ValueRange *string `type:"string"`
}

// String returns the string representation
func (s TemplateParamForDescribeParameterTemplateOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TemplateParamForDescribeParameterTemplateOutput) GoString() string {
	return s.String()
}

// SetDefaultValue sets the DefaultValue field's value.
func (s *TemplateParamForDescribeParameterTemplateOutput) SetDefaultValue(v string) *TemplateParamForDescribeParameterTemplateOutput {
	s.DefaultValue = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *TemplateParamForDescribeParameterTemplateOutput) SetDescription(v string) *TemplateParamForDescribeParameterTemplateOutput {
	s.Description = &v
	return s
}

// SetExpectValue sets the ExpectValue field's value.
func (s *TemplateParamForDescribeParameterTemplateOutput) SetExpectValue(v string) *TemplateParamForDescribeParameterTemplateOutput {
	s.ExpectValue = &v
	return s
}

// SetName sets the Name field's value.
func (s *TemplateParamForDescribeParameterTemplateOutput) SetName(v string) *TemplateParamForDescribeParameterTemplateOutput {
	s.Name = &v
	return s
}

// SetRestart sets the Restart field's value.
func (s *TemplateParamForDescribeParameterTemplateOutput) SetRestart(v bool) *TemplateParamForDescribeParameterTemplateOutput {
	s.Restart = &v
	return s
}

// SetRunningValue sets the RunningValue field's value.
func (s *TemplateParamForDescribeParameterTemplateOutput) SetRunningValue(v string) *TemplateParamForDescribeParameterTemplateOutput {
	s.RunningValue = &v
	return s
}

// SetValueRange sets the ValueRange field's value.
func (s *TemplateParamForDescribeParameterTemplateOutput) SetValueRange(v string) *TemplateParamForDescribeParameterTemplateOutput {
	s.ValueRange = &v
	return s
}

const (
	// EnumOfTemplateCategoryForDescribeParameterTemplateOutputDbengine is a EnumOfTemplateCategoryForDescribeParameterTemplateOutput enum value
	EnumOfTemplateCategoryForDescribeParameterTemplateOutputDbengine = "DBEngine"

	// EnumOfTemplateCategoryForDescribeParameterTemplateOutputProxy is a EnumOfTemplateCategoryForDescribeParameterTemplateOutput enum value
	EnumOfTemplateCategoryForDescribeParameterTemplateOutputProxy = "Proxy"
)

const (
	// EnumOfTemplateSourceForDescribeParameterTemplateOutputSystem is a EnumOfTemplateSourceForDescribeParameterTemplateOutput enum value
	EnumOfTemplateSourceForDescribeParameterTemplateOutputSystem = "System"

	// EnumOfTemplateSourceForDescribeParameterTemplateOutputUser is a EnumOfTemplateSourceForDescribeParameterTemplateOutput enum value
	EnumOfTemplateSourceForDescribeParameterTemplateOutputUser = "User"
)

const (
	// EnumOfTemplateTypeForDescribeParameterTemplateOutputMySql is a EnumOfTemplateTypeForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeForDescribeParameterTemplateOutputMySql = "MySQL"

	// EnumOfTemplateTypeForDescribeParameterTemplateOutputPostgres is a EnumOfTemplateTypeForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeForDescribeParameterTemplateOutputPostgres = "Postgres"

	// EnumOfTemplateTypeForDescribeParameterTemplateOutputSqlserver is a EnumOfTemplateTypeForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeForDescribeParameterTemplateOutputSqlserver = "Sqlserver"
)

const (
	// EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputMySql55 is a EnumOfTemplateTypeVersionForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputMySql55 = "MySQL_5_5"

	// EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputMySql56 is a EnumOfTemplateTypeVersionForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputMySql56 = "MySQL_5_6"

	// EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputMySql80 is a EnumOfTemplateTypeVersionForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputMySql80 = "MySQL_8_0"

	// EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputMySqlCommunity57 is a EnumOfTemplateTypeVersionForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputMySqlCommunity57 = "MySQL_Community_5_7"

	// EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputPostgres12 is a EnumOfTemplateTypeVersionForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputPostgres12 = "Postgres_12"

	// EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputSqlserver2019Ent is a EnumOfTemplateTypeVersionForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputSqlserver2019Ent = "SQLServer_2019_Ent"

	// EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputSqlserver2019Std is a EnumOfTemplateTypeVersionForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputSqlserver2019Std = "SQLServer_2019_Std"

	// EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputSqlserver2019Web is a EnumOfTemplateTypeVersionForDescribeParameterTemplateOutput enum value
	EnumOfTemplateTypeVersionForDescribeParameterTemplateOutputSqlserver2019Web = "SQLServer_2019_Web"
)
