// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/matpool/volcengine-go-sdk/volcengine"
	"github.com/matpool/volcengine-go-sdk/volcengine/request"
	"github.com/matpool/volcengine-go-sdk/volcengine/response"
	"github.com/matpool/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDirectConnectConnectionAttributesCommon = "DescribeDirectConnectConnectionAttributes"

// DescribeDirectConnectConnectionAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectConnectionAttributesCommon operation. The "output" return
// value will be populated with the DescribeDirectConnectConnectionAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectConnectionAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectConnectionAttributesCommon Send returns without error.
//
// See DescribeDirectConnectConnectionAttributesCommon for more information on using the DescribeDirectConnectConnectionAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeDirectConnectConnectionAttributesCommonRequest method.
//	req, resp := client.DescribeDirectConnectConnectionAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectConnectionAttributesCommon,
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

// DescribeDirectConnectConnectionAttributesCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectConnectionAttributesCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectConnectionAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectConnectionAttributesCommonWithContext is the same as DescribeDirectConnectConnectionAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectConnectionAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectConnectionAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDirectConnectConnectionAttributes = "DescribeDirectConnectConnectionAttributes"

// DescribeDirectConnectConnectionAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectConnectionAttributes operation. The "output" return
// value will be populated with the DescribeDirectConnectConnectionAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectConnectionAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectConnectionAttributesCommon Send returns without error.
//
// See DescribeDirectConnectConnectionAttributes for more information on using the DescribeDirectConnectConnectionAttributes
// API call, and error handling.
//
//	// Example sending a request using the DescribeDirectConnectConnectionAttributesRequest method.
//	req, resp := client.DescribeDirectConnectConnectionAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionAttributesRequest(input *DescribeDirectConnectConnectionAttributesInput) (req *request.Request, output *DescribeDirectConnectConnectionAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectConnectionAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDirectConnectConnectionAttributesInput{}
	}

	output = &DescribeDirectConnectConnectionAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDirectConnectConnectionAttributes API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectConnectionAttributes for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionAttributes(input *DescribeDirectConnectConnectionAttributesInput) (*DescribeDirectConnectConnectionAttributesOutput, error) {
	req, out := c.DescribeDirectConnectConnectionAttributesRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectConnectionAttributesWithContext is the same as DescribeDirectConnectConnectionAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectConnectionAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectConnectionAttributesWithContext(ctx volcengine.Context, input *DescribeDirectConnectConnectionAttributesInput, opts ...request.Option) (*DescribeDirectConnectConnectionAttributesOutput, error) {
	req, out := c.DescribeDirectConnectConnectionAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeDirectConnectConnectionAttributesInput struct {
	_ struct{} `type:"structure"`

	// DirectConnectConnectionId is a required field
	DirectConnectConnectionId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDirectConnectConnectionAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectConnectionAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDirectConnectConnectionAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDirectConnectConnectionAttributesInput"}
	if s.DirectConnectConnectionId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDirectConnectConnectionId sets the DirectConnectConnectionId field's value.
func (s *DescribeDirectConnectConnectionAttributesInput) SetDirectConnectConnectionId(v string) *DescribeDirectConnectConnectionAttributesInput {
	s.DirectConnectConnectionId = &v
	return s
}

type DescribeDirectConnectConnectionAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string"`

	Bandwidth *int64 `type:"integer"`

	BillingType *int64 `type:"integer"`

	BusinessStatus *string `type:"string"`

	ConnectionType *string `type:"string"`

	CreationTime *string `type:"string"`

	CustomerContactEmail *string `type:"string"`

	CustomerContactPhone *string `type:"string"`

	CustomerName *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	DirectConnectAccessPointId *string `type:"string"`

	DirectConnectConnectionId *string `type:"string"`

	DirectConnectConnectionName *string `type:"string"`

	ExpectBandwidth *int64 `type:"integer"`

	ExpiredTime *string `type:"string"`

	LineOperator *string `type:"string"`

	ParentConnectionAccountId *string `type:"string"`

	ParentConnectionId *string `type:"string"`

	PeerLocation *string `type:"string"`

	PortSpec *string `type:"string"`

	PortType *string `type:"string"`

	RequestId *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeDirectConnectConnectionAttributesOutput `type:"list"`

	UpdateTime *string `type:"string"`

	VlanId *int64 `type:"integer"`
}

// String returns the string representation
func (s DescribeDirectConnectConnectionAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectConnectionAttributesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetAccountId(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.AccountId = &v
	return s
}

// SetBandwidth sets the Bandwidth field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetBandwidth(v int64) *DescribeDirectConnectConnectionAttributesOutput {
	s.Bandwidth = &v
	return s
}

// SetBillingType sets the BillingType field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetBillingType(v int64) *DescribeDirectConnectConnectionAttributesOutput {
	s.BillingType = &v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetBusinessStatus(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.BusinessStatus = &v
	return s
}

// SetConnectionType sets the ConnectionType field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetConnectionType(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.ConnectionType = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetCreationTime(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetCustomerContactEmail sets the CustomerContactEmail field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetCustomerContactEmail(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.CustomerContactEmail = &v
	return s
}

// SetCustomerContactPhone sets the CustomerContactPhone field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetCustomerContactPhone(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.CustomerContactPhone = &v
	return s
}

// SetCustomerName sets the CustomerName field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetCustomerName(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.CustomerName = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetDeletedTime(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetDescription(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.Description = &v
	return s
}

// SetDirectConnectAccessPointId sets the DirectConnectAccessPointId field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetDirectConnectAccessPointId(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.DirectConnectAccessPointId = &v
	return s
}

// SetDirectConnectConnectionId sets the DirectConnectConnectionId field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetDirectConnectConnectionId(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.DirectConnectConnectionId = &v
	return s
}

// SetDirectConnectConnectionName sets the DirectConnectConnectionName field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetDirectConnectConnectionName(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.DirectConnectConnectionName = &v
	return s
}

// SetExpectBandwidth sets the ExpectBandwidth field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetExpectBandwidth(v int64) *DescribeDirectConnectConnectionAttributesOutput {
	s.ExpectBandwidth = &v
	return s
}

// SetExpiredTime sets the ExpiredTime field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetExpiredTime(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.ExpiredTime = &v
	return s
}

// SetLineOperator sets the LineOperator field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetLineOperator(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.LineOperator = &v
	return s
}

// SetParentConnectionAccountId sets the ParentConnectionAccountId field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetParentConnectionAccountId(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.ParentConnectionAccountId = &v
	return s
}

// SetParentConnectionId sets the ParentConnectionId field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetParentConnectionId(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.ParentConnectionId = &v
	return s
}

// SetPeerLocation sets the PeerLocation field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetPeerLocation(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.PeerLocation = &v
	return s
}

// SetPortSpec sets the PortSpec field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetPortSpec(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.PortSpec = &v
	return s
}

// SetPortType sets the PortType field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetPortType(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.PortType = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetRequestId(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.RequestId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetStatus(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetTags(v []*TagForDescribeDirectConnectConnectionAttributesOutput) *DescribeDirectConnectConnectionAttributesOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetUpdateTime(v string) *DescribeDirectConnectConnectionAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetVlanId sets the VlanId field's value.
func (s *DescribeDirectConnectConnectionAttributesOutput) SetVlanId(v int64) *DescribeDirectConnectConnectionAttributesOutput {
	s.VlanId = &v
	return s
}

type TagForDescribeDirectConnectConnectionAttributesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeDirectConnectConnectionAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeDirectConnectConnectionAttributesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeDirectConnectConnectionAttributesOutput) SetKey(v string) *TagForDescribeDirectConnectConnectionAttributesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeDirectConnectConnectionAttributesOutput) SetValue(v string) *TagForDescribeDirectConnectConnectionAttributesOutput {
	s.Value = &v
	return s
}
