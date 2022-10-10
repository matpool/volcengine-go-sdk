// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBackupsCommon = "DescribeBackups"

// DescribeBackupsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackupsCommon operation. The "output" return
// value will be populated with the DescribeBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupsCommon Send returns without error.
//
// See DescribeBackupsCommon for more information on using the DescribeBackupsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupsCommonRequest method.
//    req, resp := client.DescribeBackupsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeBackupsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBackupsCommon,
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

// DescribeBackupsCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeBackupsCommon for usage and error information.
func (c *REDIS) DescribeBackupsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupsCommonRequest(input)
	return out, req.Send()
}

// DescribeBackupsCommonWithContext is the same as DescribeBackupsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackupsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeBackupsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBackupsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBackups = "DescribeBackups"

// DescribeBackupsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBackups operation. The "output" return
// value will be populated with the DescribeBackupsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBackupsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBackupsCommon Send returns without error.
//
// See DescribeBackups for more information on using the DescribeBackups
// API call, and error handling.
//
//    // Example sending a request using the DescribeBackupsRequest method.
//    req, resp := client.DescribeBackupsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeBackupsRequest(input *DescribeBackupsInput) (req *request.Request, output *DescribeBackupsOutput) {
	op := &request.Operation{
		Name:       opDescribeBackups,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBackupsInput{}
	}

	output = &DescribeBackupsOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeBackups API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeBackups for usage and error information.
func (c *REDIS) DescribeBackups(input *DescribeBackupsInput) (*DescribeBackupsOutput, error) {
	req, out := c.DescribeBackupsRequest(input)
	return out, req.Send()
}

// DescribeBackupsWithContext is the same as DescribeBackups with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBackups for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeBackupsWithContext(ctx volcengine.Context, input *DescribeBackupsInput, opts ...request.Option) (*DescribeBackupsOutput, error) {
	req, out := c.DescribeBackupsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type BackupForDescribeBackupsOutput struct {
	_ struct{} `type:"structure"`

	BackupStrategy *string `type:"string"`

	BackupType *string `type:"string"`

	EndTime *string `type:"string"`

	InstanceDetail *InstanceDetailForDescribeBackupsOutput `type:"structure"`

	InstanceId *string `type:"string"`

	Size *int64 `type:"int64"`

	StartTime *string `type:"string"`

	Status *string `type:"string"`
}

// String returns the string representation
func (s BackupForDescribeBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s BackupForDescribeBackupsOutput) GoString() string {
	return s.String()
}

// SetBackupStrategy sets the BackupStrategy field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupStrategy(v string) *BackupForDescribeBackupsOutput {
	s.BackupStrategy = &v
	return s
}

// SetBackupType sets the BackupType field's value.
func (s *BackupForDescribeBackupsOutput) SetBackupType(v string) *BackupForDescribeBackupsOutput {
	s.BackupType = &v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *BackupForDescribeBackupsOutput) SetEndTime(v string) *BackupForDescribeBackupsOutput {
	s.EndTime = &v
	return s
}

// SetInstanceDetail sets the InstanceDetail field's value.
func (s *BackupForDescribeBackupsOutput) SetInstanceDetail(v *InstanceDetailForDescribeBackupsOutput) *BackupForDescribeBackupsOutput {
	s.InstanceDetail = v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *BackupForDescribeBackupsOutput) SetInstanceId(v string) *BackupForDescribeBackupsOutput {
	s.InstanceId = &v
	return s
}

// SetSize sets the Size field's value.
func (s *BackupForDescribeBackupsOutput) SetSize(v int64) *BackupForDescribeBackupsOutput {
	s.Size = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *BackupForDescribeBackupsOutput) SetStartTime(v string) *BackupForDescribeBackupsOutput {
	s.StartTime = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *BackupForDescribeBackupsOutput) SetStatus(v string) *BackupForDescribeBackupsOutput {
	s.Status = &v
	return s
}

type DescribeBackupsInput struct {
	_ struct{} `type:"structure"`

	BackupStrategyList []*string `type:"list"`

	EndTime *string `type:"string"`

	InstanceId *string `type:"string"`

	PageNumber *int32 `type:"int32"`

	PageSize *int32 `type:"int32"`

	StartTime *string `type:"string"`
}

// String returns the string representation
func (s DescribeBackupsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupsInput) GoString() string {
	return s.String()
}

// SetBackupStrategyList sets the BackupStrategyList field's value.
func (s *DescribeBackupsInput) SetBackupStrategyList(v []*string) *DescribeBackupsInput {
	s.BackupStrategyList = v
	return s
}

// SetEndTime sets the EndTime field's value.
func (s *DescribeBackupsInput) SetEndTime(v string) *DescribeBackupsInput {
	s.EndTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeBackupsInput) SetInstanceId(v string) *DescribeBackupsInput {
	s.InstanceId = &v
	return s
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeBackupsInput) SetPageNumber(v int32) *DescribeBackupsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeBackupsInput) SetPageSize(v int32) *DescribeBackupsInput {
	s.PageSize = &v
	return s
}

// SetStartTime sets the StartTime field's value.
func (s *DescribeBackupsInput) SetStartTime(v string) *DescribeBackupsInput {
	s.StartTime = &v
	return s
}

type DescribeBackupsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Backups []*BackupForDescribeBackupsOutput `type:"list"`

	Total *int32 `type:"int32"`
}

// String returns the string representation
func (s DescribeBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBackupsOutput) GoString() string {
	return s.String()
}

// SetBackups sets the Backups field's value.
func (s *DescribeBackupsOutput) SetBackups(v []*BackupForDescribeBackupsOutput) *DescribeBackupsOutput {
	s.Backups = v
	return s
}

// SetTotal sets the Total field's value.
func (s *DescribeBackupsOutput) SetTotal(v int32) *DescribeBackupsOutput {
	s.Total = &v
	return s
}

type InstanceDetailForDescribeBackupsOutput struct {
	_ struct{} `type:"structure"`

	AccountId *int32 `type:"int32"`

	ArchType *string `type:"string"`

	ChargeType *string `type:"string"`

	EngineVersion *string `type:"string"`

	ExpiredTime *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	MaintenanceTime *string `type:"string"`

	NetworkType *string `type:"string"`

	ProjectName *string `type:"string"`

	RegionId *string `type:"string"`

	Replicas *int32 `type:"int32"`

	ServerCpu *int32 `type:"int32"`

	ShardCapacity *int64 `type:"int64"`

	ShardCount *int32 `type:"int32"`

	TotalCapacity *int64 `type:"int64"`

	UsedCapacity *int64 `type:"int64"`

	VpcInfo *VpcInfoForDescribeBackupsOutput `type:"structure"`

	ZoneIds []*string `type:"list"`
}

// String returns the string representation
func (s InstanceDetailForDescribeBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceDetailForDescribeBackupsOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetAccountId(v int32) *InstanceDetailForDescribeBackupsOutput {
	s.AccountId = &v
	return s
}

// SetArchType sets the ArchType field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetArchType(v string) *InstanceDetailForDescribeBackupsOutput {
	s.ArchType = &v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetChargeType(v string) *InstanceDetailForDescribeBackupsOutput {
	s.ChargeType = &v
	return s
}

// SetEngineVersion sets the EngineVersion field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetEngineVersion(v string) *InstanceDetailForDescribeBackupsOutput {
	s.EngineVersion = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetExpiredTime(v string) *InstanceDetailForDescribeBackupsOutput {
	s.ExpiredTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetInstanceId(v string) *InstanceDetailForDescribeBackupsOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetInstanceName(v string) *InstanceDetailForDescribeBackupsOutput {
	s.InstanceName = &v
	return s
}

// SetMaintenanceTime sets the MaintenanceTime field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetMaintenanceTime(v string) *InstanceDetailForDescribeBackupsOutput {
	s.MaintenanceTime = &v
	return s
}

// SetNetworkType sets the NetworkType field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetNetworkType(v string) *InstanceDetailForDescribeBackupsOutput {
	s.NetworkType = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetProjectName(v string) *InstanceDetailForDescribeBackupsOutput {
	s.ProjectName = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetRegionId(v string) *InstanceDetailForDescribeBackupsOutput {
	s.RegionId = &v
	return s
}

// SetReplicas sets the Replicas field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetReplicas(v int32) *InstanceDetailForDescribeBackupsOutput {
	s.Replicas = &v
	return s
}

// SetServerCpu sets the ServerCpu field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetServerCpu(v int32) *InstanceDetailForDescribeBackupsOutput {
	s.ServerCpu = &v
	return s
}

// SetShardCapacity sets the ShardCapacity field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetShardCapacity(v int64) *InstanceDetailForDescribeBackupsOutput {
	s.ShardCapacity = &v
	return s
}

// SetShardCount sets the ShardCount field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetShardCount(v int32) *InstanceDetailForDescribeBackupsOutput {
	s.ShardCount = &v
	return s
}

// SetTotalCapacity sets the TotalCapacity field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetTotalCapacity(v int64) *InstanceDetailForDescribeBackupsOutput {
	s.TotalCapacity = &v
	return s
}

// SetUsedCapacity sets the UsedCapacity field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetUsedCapacity(v int64) *InstanceDetailForDescribeBackupsOutput {
	s.UsedCapacity = &v
	return s
}

// SetVpcInfo sets the VpcInfo field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetVpcInfo(v *VpcInfoForDescribeBackupsOutput) *InstanceDetailForDescribeBackupsOutput {
	s.VpcInfo = v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *InstanceDetailForDescribeBackupsOutput) SetZoneIds(v []*string) *InstanceDetailForDescribeBackupsOutput {
	s.ZoneIds = v
	return s
}

type VpcInfoForDescribeBackupsOutput struct {
	_ struct{} `type:"structure"`

	ID *string `type:"string"`

	Name *string `type:"string"`
}

// String returns the string representation
func (s VpcInfoForDescribeBackupsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VpcInfoForDescribeBackupsOutput) GoString() string {
	return s.String()
}

// SetID sets the ID field's value.
func (s *VpcInfoForDescribeBackupsOutput) SetID(v string) *VpcInfoForDescribeBackupsOutput {
	s.ID = &v
	return s
}

// SetName sets the Name field's value.
func (s *VpcInfoForDescribeBackupsOutput) SetName(v string) *VpcInfoForDescribeBackupsOutput {
	s.Name = &v
	return s
}
