// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeDirectConnectGatewayAttributesCommon = "DescribeDirectConnectGatewayAttributes"

// DescribeDirectConnectGatewayAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectGatewayAttributesCommon operation. The "output" return
// value will be populated with the DescribeDirectConnectGatewayAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectGatewayAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectGatewayAttributesCommon Send returns without error.
//
// See DescribeDirectConnectGatewayAttributesCommon for more information on using the DescribeDirectConnectGatewayAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeDirectConnectGatewayAttributesCommonRequest method.
//	req, resp := client.DescribeDirectConnectGatewayAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectGatewayAttributesCommon,
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

// DescribeDirectConnectGatewayAttributesCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectGatewayAttributesCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectGatewayAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectGatewayAttributesCommonWithContext is the same as DescribeDirectConnectGatewayAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectGatewayAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeDirectConnectGatewayAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeDirectConnectGatewayAttributes = "DescribeDirectConnectGatewayAttributes"

// DescribeDirectConnectGatewayAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeDirectConnectGatewayAttributes operation. The "output" return
// value will be populated with the DescribeDirectConnectGatewayAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeDirectConnectGatewayAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeDirectConnectGatewayAttributesCommon Send returns without error.
//
// See DescribeDirectConnectGatewayAttributes for more information on using the DescribeDirectConnectGatewayAttributes
// API call, and error handling.
//
//	// Example sending a request using the DescribeDirectConnectGatewayAttributesRequest method.
//	req, resp := client.DescribeDirectConnectGatewayAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayAttributesRequest(input *DescribeDirectConnectGatewayAttributesInput) (req *request.Request, output *DescribeDirectConnectGatewayAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeDirectConnectGatewayAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeDirectConnectGatewayAttributesInput{}
	}

	output = &DescribeDirectConnectGatewayAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeDirectConnectGatewayAttributes API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeDirectConnectGatewayAttributes for usage and error information.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayAttributes(input *DescribeDirectConnectGatewayAttributesInput) (*DescribeDirectConnectGatewayAttributesOutput, error) {
	req, out := c.DescribeDirectConnectGatewayAttributesRequest(input)
	return out, req.Send()
}

// DescribeDirectConnectGatewayAttributesWithContext is the same as DescribeDirectConnectGatewayAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeDirectConnectGatewayAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeDirectConnectGatewayAttributesWithContext(ctx volcengine.Context, input *DescribeDirectConnectGatewayAttributesInput, opts ...request.Option) (*DescribeDirectConnectGatewayAttributesOutput, error) {
	req, out := c.DescribeDirectConnectGatewayAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AssociateCenForDescribeDirectConnectGatewayAttributesOutput struct {
	_ struct{} `type:"structure"`

	CenId *string `type:"string"`

	CenOwnerId *string `type:"string"`

	CenStatus *string `type:"string"`
}

// String returns the string representation
func (s AssociateCenForDescribeDirectConnectGatewayAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AssociateCenForDescribeDirectConnectGatewayAttributesOutput) GoString() string {
	return s.String()
}

// SetCenId sets the CenId field's value.
func (s *AssociateCenForDescribeDirectConnectGatewayAttributesOutput) SetCenId(v string) *AssociateCenForDescribeDirectConnectGatewayAttributesOutput {
	s.CenId = &v
	return s
}

// SetCenOwnerId sets the CenOwnerId field's value.
func (s *AssociateCenForDescribeDirectConnectGatewayAttributesOutput) SetCenOwnerId(v string) *AssociateCenForDescribeDirectConnectGatewayAttributesOutput {
	s.CenOwnerId = &v
	return s
}

// SetCenStatus sets the CenStatus field's value.
func (s *AssociateCenForDescribeDirectConnectGatewayAttributesOutput) SetCenStatus(v string) *AssociateCenForDescribeDirectConnectGatewayAttributesOutput {
	s.CenStatus = &v
	return s
}

type DescribeDirectConnectGatewayAttributesInput struct {
	_ struct{} `type:"structure"`

	// DirectConnectGatewayId is a required field
	DirectConnectGatewayId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeDirectConnectGatewayAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectGatewayAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeDirectConnectGatewayAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeDirectConnectGatewayAttributesInput"}
	if s.DirectConnectGatewayId == nil {
		invalidParams.Add(request.NewErrParamRequired("DirectConnectGatewayId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *DescribeDirectConnectGatewayAttributesInput) SetDirectConnectGatewayId(v string) *DescribeDirectConnectGatewayAttributesInput {
	s.DirectConnectGatewayId = &v
	return s
}

type DescribeDirectConnectGatewayAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string"`

	AssociateCens []*AssociateCenForDescribeDirectConnectGatewayAttributesOutput `type:"list"`

	BusinessStatus *string `type:"string"`

	CreationTime *string `type:"string"`

	DeletedTime *string `type:"string"`

	Description *string `type:"string"`

	DirectConnectGatewayId *string `type:"string"`

	DirectConnectGatewayName *string `type:"string"`

	LockReason *string `type:"string"`

	OverdueTime *string `type:"string"`

	RequestId *string `type:"string"`

	Status *string `type:"string"`

	Tags []*TagForDescribeDirectConnectGatewayAttributesOutput `type:"list"`

	UpdateTime *string `type:"string"`
}

// String returns the string representation
func (s DescribeDirectConnectGatewayAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeDirectConnectGatewayAttributesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetAccountId(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.AccountId = &v
	return s
}

// SetAssociateCens sets the AssociateCens field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetAssociateCens(v []*AssociateCenForDescribeDirectConnectGatewayAttributesOutput) *DescribeDirectConnectGatewayAttributesOutput {
	s.AssociateCens = v
	return s
}

// SetBusinessStatus sets the BusinessStatus field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetBusinessStatus(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.BusinessStatus = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetCreationTime(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDeletedTime sets the DeletedTime field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetDeletedTime(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.DeletedTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetDescription(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.Description = &v
	return s
}

// SetDirectConnectGatewayId sets the DirectConnectGatewayId field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetDirectConnectGatewayId(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.DirectConnectGatewayId = &v
	return s
}

// SetDirectConnectGatewayName sets the DirectConnectGatewayName field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetDirectConnectGatewayName(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.DirectConnectGatewayName = &v
	return s
}

// SetLockReason sets the LockReason field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetLockReason(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.LockReason = &v
	return s
}

// SetOverdueTime sets the OverdueTime field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetOverdueTime(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.OverdueTime = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetRequestId(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.RequestId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetStatus(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.Status = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetTags(v []*TagForDescribeDirectConnectGatewayAttributesOutput) *DescribeDirectConnectGatewayAttributesOutput {
	s.Tags = v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeDirectConnectGatewayAttributesOutput) SetUpdateTime(v string) *DescribeDirectConnectGatewayAttributesOutput {
	s.UpdateTime = &v
	return s
}

type TagForDescribeDirectConnectGatewayAttributesOutput struct {
	_ struct{} `type:"structure"`

	Key *string `type:"string"`

	Value *string `type:"string"`
}

// String returns the string representation
func (s TagForDescribeDirectConnectGatewayAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s TagForDescribeDirectConnectGatewayAttributesOutput) GoString() string {
	return s.String()
}

// SetKey sets the Key field's value.
func (s *TagForDescribeDirectConnectGatewayAttributesOutput) SetKey(v string) *TagForDescribeDirectConnectGatewayAttributesOutput {
	s.Key = &v
	return s
}

// SetValue sets the Value field's value.
func (s *TagForDescribeDirectConnectGatewayAttributesOutput) SetValue(v string) *TagForDescribeDirectConnectGatewayAttributesOutput {
	s.Value = &v
	return s
}
