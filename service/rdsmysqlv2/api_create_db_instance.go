// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysqlv2

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opCreateDBInstanceCommon = "CreateDBInstance"

// CreateDBInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBInstanceCommon operation. The "output" return
// value will be populated with the CreateDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBInstanceCommon Send returns without error.
//
// See CreateDBInstanceCommon for more information on using the CreateDBInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateDBInstanceCommonRequest method.
//    req, resp := client.CreateDBInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) CreateDBInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateDBInstanceCommon,
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

// CreateDBInstanceCommon API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation CreateDBInstanceCommon for usage and error information.
func (c *RDSMYSQLV2) CreateDBInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateDBInstanceCommonRequest(input)
	return out, req.Send()
}

// CreateDBInstanceCommonWithContext is the same as CreateDBInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) CreateDBInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateDBInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateDBInstance = "CreateDBInstance"

// CreateDBInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateDBInstance operation. The "output" return
// value will be populated with the CreateDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateDBInstanceCommon Send returns without error.
//
// See CreateDBInstance for more information on using the CreateDBInstance
// API call, and error handling.
//
//    // Example sending a request using the CreateDBInstanceRequest method.
//    req, resp := client.CreateDBInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQLV2) CreateDBInstanceRequest(input *CreateDBInstanceInput) (req *request.Request, output *CreateDBInstanceOutput) {
	op := &request.Operation{
		Name:       opCreateDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateDBInstanceInput{}
	}

	output = &CreateDBInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// CreateDBInstance API operation for RDS_MYSQL_V2.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL_V2's
// API operation CreateDBInstance for usage and error information.
func (c *RDSMYSQLV2) CreateDBInstance(input *CreateDBInstanceInput) (*CreateDBInstanceOutput, error) {
	req, out := c.CreateDBInstanceRequest(input)
	return out, req.Send()
}

// CreateDBInstanceWithContext is the same as CreateDBInstance with the addition of
// the ability to pass a context and additional request options.
//
// See CreateDBInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQLV2) CreateDBInstanceWithContext(ctx volcengine.Context, input *CreateDBInstanceInput, opts ...request.Option) (*CreateDBInstanceOutput, error) {
	req, out := c.CreateDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type ChargeInfoForCreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	AutoRenew *bool `type:"boolean"`

	ChargeType *string `type:"string" enum:"EnumOfChargeTypeForCreateDBInstanceInput"`

	Number *int32 `type:"int32"`

	Period *int32 `type:"int32"`

	PeriodUnit *string `type:"string" enum:"EnumOfPeriodUnitForCreateDBInstanceInput"`
}

// String returns the string representation
func (s ChargeInfoForCreateDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ChargeInfoForCreateDBInstanceInput) GoString() string {
	return s.String()
}

// SetAutoRenew sets the AutoRenew field's value.
func (s *ChargeInfoForCreateDBInstanceInput) SetAutoRenew(v bool) *ChargeInfoForCreateDBInstanceInput {
	s.AutoRenew = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *ChargeInfoForCreateDBInstanceInput) SetChargeType(v string) *ChargeInfoForCreateDBInstanceInput {
	s.ChargeType = &v
	return s
}

// SetNumber sets the Number field's value.
func (s *ChargeInfoForCreateDBInstanceInput) SetNumber(v int32) *ChargeInfoForCreateDBInstanceInput {
	s.Number = &v
	return s
}

// SetPeriod sets the Period field's value.
func (s *ChargeInfoForCreateDBInstanceInput) SetPeriod(v int32) *ChargeInfoForCreateDBInstanceInput {
	s.Period = &v
	return s
}

// SetPeriodUnit sets the PeriodUnit field's value.
func (s *ChargeInfoForCreateDBInstanceInput) SetPeriodUnit(v string) *ChargeInfoForCreateDBInstanceInput {
	s.PeriodUnit = &v
	return s
}

type CreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	ChargeInfo *ChargeInfoForCreateDBInstanceInput `type:"structure"`

	DBEngineVersion *string `type:"string" enum:"EnumOfDBEngineVersionForCreateDBInstanceInput"`

	DBParamGroupId *string `type:"string"`

	DBTimeZone *string `type:"string"`

	InstanceName *string `type:"string"`

	InstanceTags []*InstanceTagForCreateDBInstanceInput `type:"list"`

	LowerCaseTableNames *string `type:"string"`

	NodeInfo []*NodeInfoForCreateDBInstanceInput `type:"list"`

	Number *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	StorageSpace *int32 `type:"int32"`

	StorageType *string `type:"string" enum:"EnumOfStorageTypeForCreateDBInstanceInput"`

	SubnetId *string `type:"string"`

	SuperAccountName *string `min:"2" max:"16" type:"string"`

	SuperAccountPassword *string `min:"8" max:"32" type:"string"`

	VpcId *string `type:"string"`
}

// String returns the string representation
func (s CreateDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDBInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateDBInstanceInput"}
	if s.SuperAccountName != nil && len(*s.SuperAccountName) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("SuperAccountName", 2))
	}
	if s.SuperAccountName != nil && len(*s.SuperAccountName) > 16 {
		invalidParams.Add(request.NewErrParamMaxLen("SuperAccountName", 16, *s.SuperAccountName))
	}
	if s.SuperAccountPassword != nil && len(*s.SuperAccountPassword) < 8 {
		invalidParams.Add(request.NewErrParamMinLen("SuperAccountPassword", 8))
	}
	if s.SuperAccountPassword != nil && len(*s.SuperAccountPassword) > 32 {
		invalidParams.Add(request.NewErrParamMaxLen("SuperAccountPassword", 32, *s.SuperAccountPassword))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetChargeInfo sets the ChargeInfo field's value.
func (s *CreateDBInstanceInput) SetChargeInfo(v *ChargeInfoForCreateDBInstanceInput) *CreateDBInstanceInput {
	s.ChargeInfo = v
	return s
}

// SetDBEngineVersion sets the DBEngineVersion field's value.
func (s *CreateDBInstanceInput) SetDBEngineVersion(v string) *CreateDBInstanceInput {
	s.DBEngineVersion = &v
	return s
}

// SetDBParamGroupId sets the DBParamGroupId field's value.
func (s *CreateDBInstanceInput) SetDBParamGroupId(v string) *CreateDBInstanceInput {
	s.DBParamGroupId = &v
	return s
}

// SetDBTimeZone sets the DBTimeZone field's value.
func (s *CreateDBInstanceInput) SetDBTimeZone(v string) *CreateDBInstanceInput {
	s.DBTimeZone = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *CreateDBInstanceInput) SetInstanceName(v string) *CreateDBInstanceInput {
	s.InstanceName = &v
	return s
}

// SetInstanceTags sets the InstanceTags field's value.
func (s *CreateDBInstanceInput) SetInstanceTags(v []*InstanceTagForCreateDBInstanceInput) *CreateDBInstanceInput {
	s.InstanceTags = v
	return s
}

// SetLowerCaseTableNames sets the LowerCaseTableNames field's value.
func (s *CreateDBInstanceInput) SetLowerCaseTableNames(v string) *CreateDBInstanceInput {
	s.LowerCaseTableNames = &v
	return s
}

// SetNodeInfo sets the NodeInfo field's value.
func (s *CreateDBInstanceInput) SetNodeInfo(v []*NodeInfoForCreateDBInstanceInput) *CreateDBInstanceInput {
	s.NodeInfo = v
	return s
}

// SetNumber sets the Number field's value.
func (s *CreateDBInstanceInput) SetNumber(v int32) *CreateDBInstanceInput {
	s.Number = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *CreateDBInstanceInput) SetProjectName(v string) *CreateDBInstanceInput {
	s.ProjectName = &v
	return s
}

// SetStorageSpace sets the StorageSpace field's value.
func (s *CreateDBInstanceInput) SetStorageSpace(v int32) *CreateDBInstanceInput {
	s.StorageSpace = &v
	return s
}

// SetStorageType sets the StorageType field's value.
func (s *CreateDBInstanceInput) SetStorageType(v string) *CreateDBInstanceInput {
	s.StorageType = &v
	return s
}

// SetSubnetId sets the SubnetId field's value.
func (s *CreateDBInstanceInput) SetSubnetId(v string) *CreateDBInstanceInput {
	s.SubnetId = &v
	return s
}

// SetSuperAccountName sets the SuperAccountName field's value.
func (s *CreateDBInstanceInput) SetSuperAccountName(v string) *CreateDBInstanceInput {
	s.SuperAccountName = &v
	return s
}

// SetSuperAccountPassword sets the SuperAccountPassword field's value.
func (s *CreateDBInstanceInput) SetSuperAccountPassword(v string) *CreateDBInstanceInput {
	s.SuperAccountPassword = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *CreateDBInstanceInput) SetVpcId(v string) *CreateDBInstanceInput {
	s.VpcId = &v
	return s
}

type CreateDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`

	OrderId *string `type:"string"`
}

// String returns the string representation
func (s CreateDBInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateDBInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *CreateDBInstanceOutput) SetInstanceId(v string) *CreateDBInstanceOutput {
	s.InstanceId = &v
	return s
}

// SetOrderId sets the OrderId field's value.
func (s *CreateDBInstanceOutput) SetOrderId(v string) *CreateDBInstanceOutput {
	s.OrderId = &v
	return s
}

type InstanceTagForCreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s InstanceTagForCreateDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceTagForCreateDBInstanceInput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *InstanceTagForCreateDBInstanceInput) SetKey(v string) *InstanceTagForCreateDBInstanceInput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *InstanceTagForCreateDBInstanceInput) SetValue(v string) *InstanceTagForCreateDBInstanceInput {
	s.Value = &v
	return s
}

type NodeInfoForCreateDBInstanceInput struct {
	_ struct{} `type:"structure"`

	NodeId *string `type:"string"`

	NodeOperateType *string `type:"string" enum:"EnumOfNodeOperateTypeForCreateDBInstanceInput"`

	NodePool *string `type:"string"`

	NodeSpec *string `type:"string"`

	NodeType *string `type:"string" enum:"EnumOfNodeTypeForCreateDBInstanceInput"`

	ZoneId *string `type:"string"`
}

// String returns the string representation
func (s NodeInfoForCreateDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s NodeInfoForCreateDBInstanceInput) GoString() string {
	return s.String()
}

// SetNodeId sets the NodeId field's value.
func (s *NodeInfoForCreateDBInstanceInput) SetNodeId(v string) *NodeInfoForCreateDBInstanceInput {
	s.NodeId = &v
	return s
}

// SetNodeOperateType sets the NodeOperateType field's value.
func (s *NodeInfoForCreateDBInstanceInput) SetNodeOperateType(v string) *NodeInfoForCreateDBInstanceInput {
	s.NodeOperateType = &v
	return s
}

// SetNodePool sets the NodePool field's value.
func (s *NodeInfoForCreateDBInstanceInput) SetNodePool(v string) *NodeInfoForCreateDBInstanceInput {
	s.NodePool = &v
	return s
}

// SetNodeSpec sets the NodeSpec field's value.
func (s *NodeInfoForCreateDBInstanceInput) SetNodeSpec(v string) *NodeInfoForCreateDBInstanceInput {
	s.NodeSpec = &v
	return s
}

// SetNodeType sets the NodeType field's value.
func (s *NodeInfoForCreateDBInstanceInput) SetNodeType(v string) *NodeInfoForCreateDBInstanceInput {
	s.NodeType = &v
	return s
}

// SetZoneId sets the ZoneId field's value.
func (s *NodeInfoForCreateDBInstanceInput) SetZoneId(v string) *NodeInfoForCreateDBInstanceInput {
	s.ZoneId = &v
	return s
}

const (
	// EnumOfChargeTypeForCreateDBInstanceInputNotEnabled is a EnumOfChargeTypeForCreateDBInstanceInput enum value
	EnumOfChargeTypeForCreateDBInstanceInputNotEnabled = "NotEnabled"

	// EnumOfChargeTypeForCreateDBInstanceInputPostPaid is a EnumOfChargeTypeForCreateDBInstanceInput enum value
	EnumOfChargeTypeForCreateDBInstanceInputPostPaid = "PostPaid"

	// EnumOfChargeTypeForCreateDBInstanceInputPrePaid is a EnumOfChargeTypeForCreateDBInstanceInput enum value
	EnumOfChargeTypeForCreateDBInstanceInputPrePaid = "PrePaid"
)

const (
	// EnumOfDBEngineVersionForCreateDBInstanceInputMySql56 is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputMySql56 = "MySQL_5_6"

	// EnumOfDBEngineVersionForCreateDBInstanceInputMySql57 is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputMySql57 = "MySQL_5_7"

	// EnumOfDBEngineVersionForCreateDBInstanceInputMySql80 is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputMySql80 = "MySQL_8_0"

	// EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Ent is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Ent = "SQLServer_2019_Ent"

	// EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Std is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Std = "SQLServer_2019_Std"

	// EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Web is a EnumOfDBEngineVersionForCreateDBInstanceInput enum value
	EnumOfDBEngineVersionForCreateDBInstanceInputSqlserver2019Web = "SQLServer_2019_Web"
)

const (
	// EnumOfNodeOperateTypeForCreateDBInstanceInputCreate is a EnumOfNodeOperateTypeForCreateDBInstanceInput enum value
	EnumOfNodeOperateTypeForCreateDBInstanceInputCreate = "Create"

	// EnumOfNodeOperateTypeForCreateDBInstanceInputDelete is a EnumOfNodeOperateTypeForCreateDBInstanceInput enum value
	EnumOfNodeOperateTypeForCreateDBInstanceInputDelete = "Delete"

	// EnumOfNodeOperateTypeForCreateDBInstanceInputModify is a EnumOfNodeOperateTypeForCreateDBInstanceInput enum value
	EnumOfNodeOperateTypeForCreateDBInstanceInputModify = "Modify"
)

const (
	// EnumOfNodeTypeForCreateDBInstanceInputPrimary is a EnumOfNodeTypeForCreateDBInstanceInput enum value
	EnumOfNodeTypeForCreateDBInstanceInputPrimary = "Primary"

	// EnumOfNodeTypeForCreateDBInstanceInputReadOnly is a EnumOfNodeTypeForCreateDBInstanceInput enum value
	EnumOfNodeTypeForCreateDBInstanceInputReadOnly = "ReadOnly"

	// EnumOfNodeTypeForCreateDBInstanceInputSecondary is a EnumOfNodeTypeForCreateDBInstanceInput enum value
	EnumOfNodeTypeForCreateDBInstanceInputSecondary = "Secondary"
)

const (
	// EnumOfPeriodUnitForCreateDBInstanceInputMonth is a EnumOfPeriodUnitForCreateDBInstanceInput enum value
	EnumOfPeriodUnitForCreateDBInstanceInputMonth = "Month"

	// EnumOfPeriodUnitForCreateDBInstanceInputYear is a EnumOfPeriodUnitForCreateDBInstanceInput enum value
	EnumOfPeriodUnitForCreateDBInstanceInputYear = "Year"
)

const (
	// EnumOfStorageTypeForCreateDBInstanceInputCloudStorage is a EnumOfStorageTypeForCreateDBInstanceInput enum value
	EnumOfStorageTypeForCreateDBInstanceInputCloudStorage = "CloudStorage"

	// EnumOfStorageTypeForCreateDBInstanceInputEssdpl1 is a EnumOfStorageTypeForCreateDBInstanceInput enum value
	EnumOfStorageTypeForCreateDBInstanceInputEssdpl1 = "ESSDPL1"

	// EnumOfStorageTypeForCreateDBInstanceInputEssdpl2 is a EnumOfStorageTypeForCreateDBInstanceInput enum value
	EnumOfStorageTypeForCreateDBInstanceInputEssdpl2 = "ESSDPL2"

	// EnumOfStorageTypeForCreateDBInstanceInputLocalSsd is a EnumOfStorageTypeForCreateDBInstanceInput enum value
	EnumOfStorageTypeForCreateDBInstanceInputLocalSsd = "LocalSSD"
)
