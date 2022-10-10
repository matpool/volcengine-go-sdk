// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package vpc

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeEipAddressAttributesCommon = "DescribeEipAddressAttributes"

// DescribeEipAddressAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEipAddressAttributesCommon operation. The "output" return
// value will be populated with the DescribeEipAddressAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEipAddressAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEipAddressAttributesCommon Send returns without error.
//
// See DescribeEipAddressAttributesCommon for more information on using the DescribeEipAddressAttributesCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeEipAddressAttributesCommonRequest method.
//    req, resp := client.DescribeEipAddressAttributesCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeEipAddressAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeEipAddressAttributesCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeEipAddressAttributesCommon API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeEipAddressAttributesCommon for usage and error information.
func (c *VPC) DescribeEipAddressAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeEipAddressAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeEipAddressAttributesCommonWithContext is the same as DescribeEipAddressAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEipAddressAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeEipAddressAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeEipAddressAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeEipAddressAttributes = "DescribeEipAddressAttributes"

// DescribeEipAddressAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeEipAddressAttributes operation. The "output" return
// value will be populated with the DescribeEipAddressAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeEipAddressAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeEipAddressAttributesCommon Send returns without error.
//
// See DescribeEipAddressAttributes for more information on using the DescribeEipAddressAttributes
// API call, and error handling.
//
//    // Example sending a request using the DescribeEipAddressAttributesRequest method.
//    req, resp := client.DescribeEipAddressAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *VPC) DescribeEipAddressAttributesRequest(input *DescribeEipAddressAttributesInput) (req *request.Request, output *DescribeEipAddressAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeEipAddressAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeEipAddressAttributesInput{}
	}

	output = &DescribeEipAddressAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeEipAddressAttributes API operation for VPC.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for VPC's
// API operation DescribeEipAddressAttributes for usage and error information.
func (c *VPC) DescribeEipAddressAttributes(input *DescribeEipAddressAttributesInput) (*DescribeEipAddressAttributesOutput, error) {
	req, out := c.DescribeEipAddressAttributesRequest(input)
	return out, req.Send()
}

// DescribeEipAddressAttributesWithContext is the same as DescribeEipAddressAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeEipAddressAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *VPC) DescribeEipAddressAttributesWithContext(ctx volcengine.Context, input *DescribeEipAddressAttributesInput, opts ...request.Option) (*DescribeEipAddressAttributesOutput, error) {
	req, out := c.DescribeEipAddressAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeEipAddressAttributesInput struct {
	_ struct{} `type:"structure"`

	// AllocationId is a required field
	AllocationId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeEipAddressAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEipAddressAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeEipAddressAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeEipAddressAttributesInput"}
	if s.AllocationId == nil {
		invalidParams.Add(request.NewErrParamRequired("AllocationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAllocationId sets the AllocationId field's value.
func (s *DescribeEipAddressAttributesInput) SetAllocationId(v string) *DescribeEipAddressAttributesInput {
	s.AllocationId = &v
	return s
}

type DescribeEipAddressAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AllocationId *string `type:"string"`

	AllocationTime *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	BandwidthPackageId *string `type:"string"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	EipAddress *string `type:"string"`

	ExpiredTime *string `type:"string"`

	ISP *string `type:"string"`

	InstanceId *string `type:"string"`

	InstanceType *string `type:"string"`

	LockReason *string `type:"string"`

	Name *string `type:"string"`

	OverdueTime *string `type:"string"`

	ProjectName *string `type:"string"`

	RequestId *string `type:"string"`

	SecurityProtectionTypes []*string `type:"list"`

	Status *string `type:"string"`

	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s DescribeEipAddressAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeEipAddressAttributesOutput) GoString() string {
	return s.String()
}

// SetAllocationId sets the AllocationId field's value.
func (s *DescribeEipAddressAttributesOutput) SetAllocationId(v string) *DescribeEipAddressAttributesOutput {
	s.AllocationId = &v
	return s
}

// SetAllocationTime sets the AllocationTime field's value.
func (s *DescribeEipAddressAttributesOutput) SetAllocationTime(v string) *DescribeEipAddressAttributesOutput {
	s.AllocationTime = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *DescribeEipAddressAttributesOutput) SetBandwidth(v int64) *DescribeEipAddressAttributesOutput {
	s.Bandwidth = &v
	return s
}

// SetBandwidthPackageId sets the BandwidthPackageId field's value.
func (s *DescribeEipAddressAttributesOutput) SetBandwidthPackageId(v string) *DescribeEipAddressAttributesOutput {
	s.BandwidthPackageId = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *DescribeEipAddressAttributesOutput) SetBillingType(v int64) *DescribeEipAddressAttributesOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *DescribeEipAddressAttributesOutput) SetBusinessStatus(v string) *DescribeEipAddressAttributesOutput {
	s.BusinessStatus = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *DescribeEipAddressAttributesOutput) SetDeletedTime(v string) *DescribeEipAddressAttributesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeEipAddressAttributesOutput) SetDescription(v string) *DescribeEipAddressAttributesOutput {
	s.Description = &v
	return s
}

// SetEipAddress sets the EipAddress field's value.
func (s *DescribeEipAddressAttributesOutput) SetEipAddress(v string) *DescribeEipAddressAttributesOutput {
	s.EipAddress = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *DescribeEipAddressAttributesOutput) SetExpiredTime(v string) *DescribeEipAddressAttributesOutput {
	s.ExpiredTime = &v
	return s
}

// SetISP sets the ISP field's value.
func (s *DescribeEipAddressAttributesOutput) SetISP(v string) *DescribeEipAddressAttributesOutput {
	s.ISP = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *DescribeEipAddressAttributesOutput) SetInstanceId(v string) *DescribeEipAddressAttributesOutput {
	s.InstanceId = &v
	return s
}

// SetInstanceType sets the InstanceType field's value.
func (s *DescribeEipAddressAttributesOutput) SetInstanceType(v string) *DescribeEipAddressAttributesOutput {
	s.InstanceType = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *DescribeEipAddressAttributesOutput) SetLockReason(v string) *DescribeEipAddressAttributesOutput {
	s.LockReason = &v
	return s
}

// SetName sets the Name field's value.
func (s *DescribeEipAddressAttributesOutput) SetName(v string) *DescribeEipAddressAttributesOutput {
	s.Name = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *DescribeEipAddressAttributesOutput) SetOverdueTime(v string) *DescribeEipAddressAttributesOutput {
	s.OverdueTime = &v
	return s
}

// SetProjectName sets the ProjectName field's value.
func (s *DescribeEipAddressAttributesOutput) SetProjectName(v string) *DescribeEipAddressAttributesOutput {
	s.ProjectName = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeEipAddressAttributesOutput) SetRequestId(v string) *DescribeEipAddressAttributesOutput {
	s.RequestId = &v
	return s
}

// SetSecurityProtectionTypes sets the SecurityProtectionTypes field's value.
func (s *DescribeEipAddressAttributesOutput) SetSecurityProtectionTypes(v []*string) *DescribeEipAddressAttributesOutput {
	s.SecurityProtectionTypes = v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeEipAddressAttributesOutput) SetStatus(v string) *DescribeEipAddressAttributesOutput {
	s.Status = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *DescribeEipAddressAttributesOutput) SetUpdatedAt(v string) *DescribeEipAddressAttributesOutput {
	s.UpdatedAt = &v
	return s
}
