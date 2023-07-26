// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package redis

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDBInstanceDetailCommon = "DescribeDBInstanceDetail"

// DescribeDBInstanceDetailCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceDetailCommon operation. The "output" return
// value will be populated with the DescribeDBInstanceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceDetailCommon Send returns without error.
//
// See DescribeDBInstanceDetailCommon for more information on using the DescribeDBInstanceDetailCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceDetailCommonRequest method.
//    req, resp := client.DescribeDBInstanceDetailCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeDBInstanceDetailCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceDetailCommon,
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

// DescribeDBInstanceDetailCommon API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeDBInstanceDetailCommon for usage and error information.
func (c *REDIS) DescribeDBInstanceDetailCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceDetailCommonRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceDetailCommonWithContext is the same as DescribeDBInstanceDetailCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceDetailCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeDBInstanceDetailCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDBInstanceDetailCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDBInstanceDetail = "DescribeDBInstanceDetail"

// DescribeDBInstanceDetailRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDBInstanceDetail operation. The "output" return
// value will be populated with the DescribeDBInstanceDetailCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDBInstanceDetailCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDBInstanceDetailCommon Send returns without error.
//
// See DescribeDBInstanceDetail for more information on using the DescribeDBInstanceDetail
// API call, and error handling.
//
//    // Example sending a request using the DescribeDBInstanceDetailRequest method.
//    req, resp := client.DescribeDBInstanceDetailRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *REDIS) DescribeDBInstanceDetailRequest(input *DescribeDBInstanceDetailInput) (req *request.Request, output *DescribeDBInstanceDetailOutput) {
	op := &request.Operation{
		Name:       opDescribeDBInstanceDetail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDBInstanceDetailInput{}
	}

	output = &DescribeDBInstanceDetailOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// DescribeDBInstanceDetail API operation for REDIS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for REDIS's
// API operation DescribeDBInstanceDetail for usage and error information.
func (c *REDIS) DescribeDBInstanceDetail(input *DescribeDBInstanceDetailInput) (*DescribeDBInstanceDetailOutput, error) {
	req, out := c.DescribeDBInstanceDetailRequest(input)
	return out, req.Send()
}

// DescribeDBInstanceDetailWithContext is the same as DescribeDBInstanceDetail with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDBInstanceDetail for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *REDIS) DescribeDBInstanceDetailWithContext(ctx volcengine.Context, input *DescribeDBInstanceDetailInput, opts ...request.Option) (*DescribeDBInstanceDetailOutput, error) {
	req, out := c.DescribeDBInstanceDetailRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type CapacityForDescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	Total *int64 `type:"int64"`

	Used *int64 `type:"int64"`
}

// String returns the string representation
func (s CapacityForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CapacityForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetTotal sets the Total field's value.
func (s *CapacityForDescribeDBInstanceDetailOutput) SetTotal(v int64) *CapacityForDescribeDBInstanceDetailOutput {
	s.Total = &v
	return s
}

// SetUsed sets the Used field's value.
func (s *CapacityForDescribeDBInstanceDetailOutput) SetUsed(v int64) *CapacityForDescribeDBInstanceDetailOutput {
	s.Used = &v
	return s
}

type DescribeDBInstanceDetailInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDBInstanceDetailInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceDetailInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDBInstanceDetailInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDBInstanceDetailInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstanceDetailInput) SetInstanceId(v string) *DescribeDBInstanceDetailInput {
	s.InstanceId = &v
	return s
}

type DescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	Capacity *CapacityForDescribeDBInstanceDetailOutput `type:"structure"`

	ChargeType *string `type:"string"`

	CreateTime *string `type:"string"`

	DeletionProtection *string `type:"string"`

	EngineVersion *string `type:"string"`

	ExpiredTime *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceName *string `type:"string"`

	MaintenanceTime *string `type:"string"`

	NodeNumber *int32 `type:"int32"`

	ProjectName *string `type:"string"`

	RegionId *string `type:"string"`

	ShardCapacity *float64 `type:"double"`

	ShardNumber *int32 `type:"int32"`

	ShardedCluster *int32 `type:"int32"`

	Status *string `type:"string"`

	VisitAddrs []*VisitAddrForDescribeDBInstanceDetailOutput `type:"list"`

	VpcAuthMode *string `type:"string"`

	VpcId *string `type:"string"`

	ZoneIds []*string `type:"list"`
}

// String returns the string representation
func (s DescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetCapacity sets the Capacity field's value.
func (s *DescribeDBInstanceDetailOutput) SetCapacity(v *CapacityForDescribeDBInstanceDetailOutput) *DescribeDBInstanceDetailOutput {
	s.Capacity = v
	return s
}

// SetChargeType sets the ChargeType field's value.
func (s *DescribeDBInstanceDetailOutput) SetChargeType(v string) *DescribeDBInstanceDetailOutput {
	s.ChargeType = &v
	return s
}

// SetCreateTime sets the CreateTime field's value.
func (s *DescribeDBInstanceDetailOutput) SetCreateTime(v string) *DescribeDBInstanceDetailOutput {
	s.CreateTime = &v
	return s
}

// SetDeletionProtection sets the DeletionProtection field's value.
func (s *DescribeDBInstanceDetailOutput) SetDeletionProtection(v string) *DescribeDBInstanceDetailOutput {
	s.DeletionProtection = &v
	return s
}

// SetEngineVersion sets the EngineVersion field's value.
func (s *DescribeDBInstanceDetailOutput) SetEngineVersion(v string) *DescribeDBInstanceDetailOutput {
	s.EngineVersion = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *DescribeDBInstanceDetailOutput) SetExpiredTime(v string) *DescribeDBInstanceDetailOutput {
	s.ExpiredTime = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeDBInstanceDetailOutput) SetInstanceId(v string) *DescribeDBInstanceDetailOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceName sets the InstanceName field's value.
func (s *DescribeDBInstanceDetailOutput) SetInstanceName(v string) *DescribeDBInstanceDetailOutput {
	s.InstanceName = &v
	return s
}

// SetMaintenanceTime sets the MaintenanceTime field's value.
func (s *DescribeDBInstanceDetailOutput) SetMaintenanceTime(v string) *DescribeDBInstanceDetailOutput {
	s.MaintenanceTime = &v
	return s
}

// SetNodeNumber sets the NodeNumber field's value.
func (s *DescribeDBInstanceDetailOutput) SetNodeNumber(v int32) *DescribeDBInstanceDetailOutput {
	s.NodeNumber = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeDBInstanceDetailOutput) SetProjectName(v string) *DescribeDBInstanceDetailOutput {
	s.ProjectName = &v
	return s
}

// SetRegionId sets the RegionId field's value.
func (s *DescribeDBInstanceDetailOutput) SetRegionId(v string) *DescribeDBInstanceDetailOutput {
	s.RegionId = &v
	return s
}

// SetShardCapacity sets the ShardCapacity field's value.
func (s *DescribeDBInstanceDetailOutput) SetShardCapacity(v float64) *DescribeDBInstanceDetailOutput {
	s.ShardCapacity = &v
	return s
}

// SetShardNumber sets the ShardNumber field's value.
func (s *DescribeDBInstanceDetailOutput) SetShardNumber(v int32) *DescribeDBInstanceDetailOutput {
	s.ShardNumber = &v
	return s
}

// SetShardedCluster sets the ShardedCluster field's value.
func (s *DescribeDBInstanceDetailOutput) SetShardedCluster(v int32) *DescribeDBInstanceDetailOutput {
	s.ShardedCluster = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeDBInstanceDetailOutput) SetStatus(v string) *DescribeDBInstanceDetailOutput {
	s.Status = &v
	return s
}

// SetVisitAddrs sets the VisitAddrs field's value.
func (s *DescribeDBInstanceDetailOutput) SetVisitAddrs(v []*VisitAddrForDescribeDBInstanceDetailOutput) *DescribeDBInstanceDetailOutput {
	s.VisitAddrs = v
	return s
}

// SetVpcAuthMode sets the VpcAuthMode field's value.
func (s *DescribeDBInstanceDetailOutput) SetVpcAuthMode(v string) *DescribeDBInstanceDetailOutput {
	s.VpcAuthMode = &v
	return s
}

// SetVpcId sets the VpcId field's value.
func (s *DescribeDBInstanceDetailOutput) SetVpcId(v string) *DescribeDBInstanceDetailOutput {
	s.VpcId = &v
	return s
}

// SetZoneIds sets the ZoneIds field's value.
func (s *DescribeDBInstanceDetailOutput) SetZoneIds(v []*string) *DescribeDBInstanceDetailOutput {
	s.ZoneIds = v
	return s
}

type VisitAddrForDescribeDBInstanceDetailOutput struct {
	_ struct{} `type:"structure"`

	AddrType *string `type:"string"`

	Address *string `type:"string"`

	EipId *string `type:"string"`

	Port *string `type:"string"`
}

// String returns the string representation
func (s VisitAddrForDescribeDBInstanceDetailOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s VisitAddrForDescribeDBInstanceDetailOutput) GoString() string {
	return s.String()
}

// SetAddrType sets the AddrType field's value.
func (s *VisitAddrForDescribeDBInstanceDetailOutput) SetAddrType(v string) *VisitAddrForDescribeDBInstanceDetailOutput {
	s.AddrType = &v
	return s
}

// SetAddress sets the Address field's value.
func (s *VisitAddrForDescribeDBInstanceDetailOutput) SetAddress(v string) *VisitAddrForDescribeDBInstanceDetailOutput {
	s.Address = &v
	return s
}

// SetEipId sets the EipId field's value.
func (s *VisitAddrForDescribeDBInstanceDetailOutput) SetEipId(v string) *VisitAddrForDescribeDBInstanceDetailOutput {
	s.EipId = &v
	return s
}

// SetPort sets the Port field's value.
func (s *VisitAddrForDescribeDBInstanceDetailOutput) SetPort(v string) *VisitAddrForDescribeDBInstanceDetailOutput {
	s.Port = &v
	return s
}
