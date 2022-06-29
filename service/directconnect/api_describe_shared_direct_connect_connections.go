// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSharedDirectConnectConnectionsCommon = "DescribeSharedDirectConnectConnections"

// DescribeSharedDirectConnectConnectionsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSharedDirectConnectConnectionsCommon operation. The "output" return
// value will be populated with the DescribeSharedDirectConnectConnectionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSharedDirectConnectConnectionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSharedDirectConnectConnectionsCommon Send returns without error.
//
// See DescribeSharedDirectConnectConnectionsCommon for more information on using the DescribeSharedDirectConnectConnectionsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSharedDirectConnectConnectionsCommonRequest method.
//    req, resp := client.DescribeSharedDirectConnectConnectionsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeSharedDirectConnectConnectionsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSharedDirectConnectConnectionsCommon,
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

// DescribeSharedDirectConnectConnectionsCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation DescribeSharedDirectConnectConnectionsCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeSharedDirectConnectConnectionsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSharedDirectConnectConnectionsCommonRequest(input)
	return out, req.Send()
}

// DescribeSharedDirectConnectConnectionsCommonWithContext is the same as DescribeSharedDirectConnectConnectionsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSharedDirectConnectConnectionsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeSharedDirectConnectConnectionsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSharedDirectConnectConnectionsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSharedDirectConnectConnections = "DescribeSharedDirectConnectConnections"

// DescribeSharedDirectConnectConnectionsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSharedDirectConnectConnections operation. The "output" return
// value will be populated with the DescribeSharedDirectConnectConnectionsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSharedDirectConnectConnectionsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSharedDirectConnectConnectionsCommon Send returns without error.
//
// See DescribeSharedDirectConnectConnections for more information on using the DescribeSharedDirectConnectConnections
// API call, and error handling.
//
//    // Example sending a request using the DescribeSharedDirectConnectConnectionsRequest method.
//    req, resp := client.DescribeSharedDirectConnectConnectionsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *DIRECTCONNECT) DescribeSharedDirectConnectConnectionsRequest(input *DescribeSharedDirectConnectConnectionsInput) (req *request.Request, output *DescribeSharedDirectConnectConnectionsOutput) {
	op := &request.Operation{
		Name:       opDescribeSharedDirectConnectConnections,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSharedDirectConnectConnectionsInput{}
	}

	output = &DescribeSharedDirectConnectConnectionsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSharedDirectConnectConnections API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCSTACK API reference guide for DIRECTCONNECT's
// API operation DescribeSharedDirectConnectConnections for usage and error information.
func (c *DIRECTCONNECT) DescribeSharedDirectConnectConnections(input *DescribeSharedDirectConnectConnectionsInput) (*DescribeSharedDirectConnectConnectionsOutput, error) {
	req, out := c.DescribeSharedDirectConnectConnectionsRequest(input)
	return out, req.Send()
}

// DescribeSharedDirectConnectConnectionsWithContext is the same as DescribeSharedDirectConnectConnections with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSharedDirectConnectConnections for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeSharedDirectConnectConnectionsWithContext(ctx volcengine.Context, input *DescribeSharedDirectConnectConnectionsInput, opts ...request.Option) (*DescribeSharedDirectConnectConnectionsOutput, error) {
	req, out := c.DescribeSharedDirectConnectConnectionsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSharedDirectConnectConnectionsInput struct {
	_ struct{} `type:"structure"`

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	ParentConnectionId *string `type:"string"`

	SharedDirectConnectConnectionIds []*string `type:"list"`

	SharedDirectConnectConnectionName *string `type:"string"`

	Status *string `type:"string"`

	UserAccountId *string `type:"string"`

	VlanId *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeSharedDirectConnectConnectionsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSharedDirectConnectConnectionsInput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSharedDirectConnectConnectionsInput) SetPageNumber(v int64) *DescribeSharedDirectConnectConnectionsInput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSharedDirectConnectConnectionsInput) SetPageSize(v int64) *DescribeSharedDirectConnectConnectionsInput {
	s.PageSize = &v
	return s
}

// SetParentConnectionId sets the ParentConnectionId field's value.
func (s *DescribeSharedDirectConnectConnectionsInput) SetParentConnectionId(v string) *DescribeSharedDirectConnectConnectionsInput {
	s.ParentConnectionId = &v
	return s
}

// SetSharedDirectConnectConnectionIds sets the SharedDirectConnectConnectionIds field's value.
func (s *DescribeSharedDirectConnectConnectionsInput) SetSharedDirectConnectConnectionIds(v []*string) *DescribeSharedDirectConnectConnectionsInput {
	s.SharedDirectConnectConnectionIds = v
	return s
}

// SetSharedDirectConnectConnectionName sets the SharedDirectConnectConnectionName field's value.
func (s *DescribeSharedDirectConnectConnectionsInput) SetSharedDirectConnectConnectionName(v string) *DescribeSharedDirectConnectConnectionsInput {
	s.SharedDirectConnectConnectionName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeSharedDirectConnectConnectionsInput) SetStatus(v string) *DescribeSharedDirectConnectConnectionsInput {
	s.Status = &v
	return s
}

// SetUserAccountId sets the UserAccountId field's value.
func (s *DescribeSharedDirectConnectConnectionsInput) SetUserAccountId(v string) *DescribeSharedDirectConnectConnectionsInput {
	s.UserAccountId = &v
	return s
}

// SetVlanId sets the VlanId field's value.
func (s *DescribeSharedDirectConnectConnectionsInput) SetVlanId(v int64) *DescribeSharedDirectConnectConnectionsInput {
	s.VlanId = &v
	return s
}

type DescribeSharedDirectConnectConnectionsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	PageNumber *int64 `type:"integer"`

	PageSize *int64 `type:"integer"`

	RequestId *string `type:"string"`

	SharedDirectConnectConnections []*SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput `type:"list"`

	TotalCount *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeSharedDirectConnectConnectionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSharedDirectConnectConnectionsOutput) GoString() string {
	return s.String()
}

// SetPageNumber sets the PageNumber field's value.
func (s *DescribeSharedDirectConnectConnectionsOutput) SetPageNumber(v int64) *DescribeSharedDirectConnectConnectionsOutput {
	s.PageNumber = &v
	return s
}

// SetPageSize sets the PageSize field's value.
func (s *DescribeSharedDirectConnectConnectionsOutput) SetPageSize(v int64) *DescribeSharedDirectConnectConnectionsOutput {
	s.PageSize = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeSharedDirectConnectConnectionsOutput) SetRequestId(v string) *DescribeSharedDirectConnectConnectionsOutput {
	s.RequestId = &v
	return s
}

// SetSharedDirectConnectConnections sets the SharedDirectConnectConnections field's value.
func (s *DescribeSharedDirectConnectConnectionsOutput) SetSharedDirectConnectConnections(v []*SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) *DescribeSharedDirectConnectConnectionsOutput {
	s.SharedDirectConnectConnections = v
	return s
}

// SetTotalCount sets the TotalCount field's value.
func (s *DescribeSharedDirectConnectConnectionsOutput) SetTotalCount(v int64) *DescribeSharedDirectConnectConnectionsOutput {
	s.TotalCount = &v
	return s
}

type SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput struct {
	_ struct{} `type:"structure"`

	AccessPoint *string `type:"string"`

	AccountId *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	Operator *string `type:"string"`

	ParentConnectionId *string `type:"string"`

	PortType *string `type:"string"`

	SharedDirectConnectConnectionId *string `type:"string"`

	SharedDirectConnectConnectionName *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`

	UserAccountId *string `type:"string"`

	VlanId *int64 `type:"integer"`
}

// String returns the string representation
func (s SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) GoString() string {
	return s.String()
}

// SetAccessPoint sets the AccessPoint field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetAccessPoint(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.AccessPoint = &v
	return s
}

// SetAccountId sets the AccountId field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetAccountId(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.AccountId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetBandwidth(v int64) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.Bandwidth = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetCreationTime(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetDescription(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.Description = &v
	return s
}

// SetOperator sets the Operator field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetOperator(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.Operator = &v
	return s
}

// SetParentConnectionId sets the ParentConnectionId field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetParentConnectionId(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.ParentConnectionId = &v
	return s
}

// SetPortType sets the PortType field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetPortType(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.PortType = &v
	return s
}

// SetSharedDirectConnectConnectionId sets the SharedDirectConnectConnectionId field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetSharedDirectConnectConnectionId(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.SharedDirectConnectConnectionId = &v
	return s
}

// SetSharedDirectConnectConnectionName sets the SharedDirectConnectConnectionName field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetSharedDirectConnectConnectionName(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.SharedDirectConnectConnectionName = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetStatus(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetUpdateTime(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.UpdateTime = &v
	return s
}

// SetUserAccountId sets the UserAccountId field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetUserAccountId(v string) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.UserAccountId = &v
	return s
}

// SetVlanId sets the VlanId field's value.
func (s *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput) SetVlanId(v int64) *SharedDirectConnectConnectionForDescribeSharedDirectConnectConnectionsOutput {
	s.VlanId = &v
	return s
}
