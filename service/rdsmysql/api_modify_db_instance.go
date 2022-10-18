// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package rdsmysql

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opModifyDBInstanceCommon = "ModifyDBInstance"

// ModifyDBInstanceCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstanceCommon operation. The "output" return
// value will be populated with the ModifyDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceCommon Send returns without error.
//
// See ModifyDBInstanceCommon for more information on using the ModifyDBInstanceCommon
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceCommonRequest method.
//    req, resp := client.ModifyDBInstanceCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ModifyDBInstanceCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opModifyDBInstanceCommon,
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

// ModifyDBInstanceCommon API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ModifyDBInstanceCommon for usage and error information.
func (c *RDSMYSQL) ModifyDBInstanceCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceCommonRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceCommonWithContext is the same as ModifyDBInstanceCommon with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstanceCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ModifyDBInstanceCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.ModifyDBInstanceCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opModifyDBInstance = "ModifyDBInstance"

// ModifyDBInstanceRequest generates a "volcengine/request.Request" representing the
// client's request for the ModifyDBInstance operation. The "output" return
// value will be populated with the ModifyDBInstanceCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned ModifyDBInstanceCommon Request to send the API call to the service.
// the "output" return value is not valid until after ModifyDBInstanceCommon Send returns without error.
//
// See ModifyDBInstance for more information on using the ModifyDBInstance
// API call, and error handling.
//
//    // Example sending a request using the ModifyDBInstanceRequest method.
//    req, resp := client.ModifyDBInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *RDSMYSQL) ModifyDBInstanceRequest(input *ModifyDBInstanceInput) (req *request.Request, output *ModifyDBInstanceOutput) {
	op := &request.Operation{
		Name:       opModifyDBInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyDBInstanceInput{}
	}

	output = &ModifyDBInstanceOutput{}
	req = c.newRequest(op, input, output)

	req.HTTPRequest.Header.Set("Content-Type", "application/json; charset=utf-8")

	return
}

// ModifyDBInstance API operation for RDS_MYSQL.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for RDS_MYSQL's
// API operation ModifyDBInstance for usage and error information.
func (c *RDSMYSQL) ModifyDBInstance(input *ModifyDBInstanceInput) (*ModifyDBInstanceOutput, error) {
	req, out := c.ModifyDBInstanceRequest(input)
	return out, req.Send()
}

// ModifyDBInstanceWithContext is the same as ModifyDBInstance with the addition of
// the ability to pass a context and additional request options.
//
// See ModifyDBInstance for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *RDSMYSQL) ModifyDBInstanceWithContext(ctx volcengine.Context, input *ModifyDBInstanceInput, opts ...request.Option) (*ModifyDBInstanceOutput, error) {
	req, out := c.ModifyDBInstanceRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type InstanceSpecForModifyDBInstanceInput struct {
	_ struct{} `type:"structure"`

	CpuNum *float64 `type:"double"`

	InstanceFamily *string `type:"string" enum:"EnumOfInstanceFamilyForModifyDBInstanceInput"`

	MemInGb *float64 `type:"double"`

	SpecName *string `type:"string"`
}

// String returns the string representation
func (s InstanceSpecForModifyDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceSpecForModifyDBInstanceInput) GoString() string {
	return s.String()
}

// SetCpuNum sets the CpuNum field's value.
func (s *InstanceSpecForModifyDBInstanceInput) SetCpuNum(v float64) *InstanceSpecForModifyDBInstanceInput {
	s.CpuNum = &v
	return s
}

// SetInstanceFamily sets the InstanceFamily field's value.
func (s *InstanceSpecForModifyDBInstanceInput) SetInstanceFamily(v string) *InstanceSpecForModifyDBInstanceInput {
	s.InstanceFamily = &v
	return s
}

// SetMemInGb sets the MemInGb field's value.
func (s *InstanceSpecForModifyDBInstanceInput) SetMemInGb(v float64) *InstanceSpecForModifyDBInstanceInput {
	s.MemInGb = &v
	return s
}

// SetSpecName sets the SpecName field's value.
func (s *InstanceSpecForModifyDBInstanceInput) SetSpecName(v string) *InstanceSpecForModifyDBInstanceInput {
	s.SpecName = &v
	return s
}

type ModifyDBInstanceInput struct {
	_ struct{} `type:"structure"`

	// InstanceId is a required field
	InstanceId *string `type:"string" required:"true"`

	InstanceSpec *InstanceSpecForModifyDBInstanceInput `type:"structure"`

	InstanceType *string `type:"string" enum:"EnumOfInstanceTypeForModifyDBInstanceInput"`

	// StorageSpaceGB is a required field
	StorageSpaceGB *int32 `type:"int32" required:"true"`
}

// String returns the string representation
func (s ModifyDBInstanceInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyDBInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyDBInstanceInput"}
	if s.InstanceId == nil {
		invalidParams.Add(request.NewErrParamRequired("InstanceId"))
	}
	if s.StorageSpaceGB == nil {
		invalidParams.Add(request.NewErrParamRequired("StorageSpaceGB"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceInput) SetInstanceId(v string) *ModifyDBInstanceInput {
	s.InstanceId = &v
	return s
}

// SetInstanceSpec sets the InstanceSpec field's value.
func (s *ModifyDBInstanceInput) SetInstanceSpec(v *InstanceSpecForModifyDBInstanceInput) *ModifyDBInstanceInput {
	s.InstanceSpec = v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *ModifyDBInstanceInput) SetInstanceType(v string) *ModifyDBInstanceInput {
	s.InstanceType = &v
	return s
}

// SetStorageSpaceGB sets the StorageSpaceGB field's value.
func (s *ModifyDBInstanceInput) SetStorageSpaceGB(v int32) *ModifyDBInstanceInput {
	s.StorageSpaceGB = &v
	return s
}

type ModifyDBInstanceOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceId *string `type:"string"`
}

// String returns the string representation
func (s ModifyDBInstanceOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyDBInstanceOutput) GoString() string {
	return s.String()
}

// SetInstanceId sets the InstanceId field's value.
func (s *ModifyDBInstanceOutput) SetInstanceId(v string) *ModifyDBInstanceOutput {
	s.InstanceId = &v
	return s
}

const (
	// EnumOfInstanceFamilyForModifyDBInstanceInputExclusive is a EnumOfInstanceFamilyForModifyDBInstanceInput enum value
	EnumOfInstanceFamilyForModifyDBInstanceInputExclusive = "Exclusive"

	// EnumOfInstanceFamilyForModifyDBInstanceInputExclusiveHost is a EnumOfInstanceFamilyForModifyDBInstanceInput enum value
	EnumOfInstanceFamilyForModifyDBInstanceInputExclusiveHost = "ExclusiveHost"

	// EnumOfInstanceFamilyForModifyDBInstanceInputGeneral is a EnumOfInstanceFamilyForModifyDBInstanceInput enum value
	EnumOfInstanceFamilyForModifyDBInstanceInputGeneral = "General"
)

const (
	// EnumOfInstanceTypeForModifyDBInstanceInputBasic is a EnumOfInstanceTypeForModifyDBInstanceInput enum value
	EnumOfInstanceTypeForModifyDBInstanceInputBasic = "Basic"

	// EnumOfInstanceTypeForModifyDBInstanceInputCluster is a EnumOfInstanceTypeForModifyDBInstanceInput enum value
	EnumOfInstanceTypeForModifyDBInstanceInputCluster = "Cluster"

	// EnumOfInstanceTypeForModifyDBInstanceInputFinance is a EnumOfInstanceTypeForModifyDBInstanceInput enum value
	EnumOfInstanceTypeForModifyDBInstanceInputFinance = "Finance"

	// EnumOfInstanceTypeForModifyDBInstanceInputHa is a EnumOfInstanceTypeForModifyDBInstanceInput enum value
	EnumOfInstanceTypeForModifyDBInstanceInputHa = "HA"
)
