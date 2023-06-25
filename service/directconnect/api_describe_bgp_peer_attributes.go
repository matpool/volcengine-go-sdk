// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package directconnect

import (
	"github.com/volcengine/volcengine-go-sdk/volcengine"
	"github.com/volcengine/volcengine-go-sdk/volcengine/request"
	"github.com/volcengine/volcengine-go-sdk/volcengine/response"
	"github.com/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeBgpPeerAttributesCommon = "DescribeBgpPeerAttributes"

// DescribeBgpPeerAttributesCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBgpPeerAttributesCommon operation. The "output" return
// value will be populated with the DescribeBgpPeerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBgpPeerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBgpPeerAttributesCommon Send returns without error.
//
// See DescribeBgpPeerAttributesCommon for more information on using the DescribeBgpPeerAttributesCommon
// API call, and error handling.
//
//	// Example sending a request using the DescribeBgpPeerAttributesCommonRequest method.
//	req, resp := client.DescribeBgpPeerAttributesCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeBgpPeerAttributesCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeBgpPeerAttributesCommon,
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

// DescribeBgpPeerAttributesCommon API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeBgpPeerAttributesCommon for usage and error information.
func (c *DIRECTCONNECT) DescribeBgpPeerAttributesCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeBgpPeerAttributesCommonRequest(input)
	return out, req.Send()
}

// DescribeBgpPeerAttributesCommonWithContext is the same as DescribeBgpPeerAttributesCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBgpPeerAttributesCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeBgpPeerAttributesCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeBgpPeerAttributesCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeBgpPeerAttributes = "DescribeBgpPeerAttributes"

// DescribeBgpPeerAttributesRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeBgpPeerAttributes operation. The "output" return
// value will be populated with the DescribeBgpPeerAttributesCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeBgpPeerAttributesCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeBgpPeerAttributesCommon Send returns without error.
//
// See DescribeBgpPeerAttributes for more information on using the DescribeBgpPeerAttributes
// API call, and error handling.
//
//	// Example sending a request using the DescribeBgpPeerAttributesRequest method.
//	req, resp := client.DescribeBgpPeerAttributesRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *DIRECTCONNECT) DescribeBgpPeerAttributesRequest(input *DescribeBgpPeerAttributesInput) (req *request.Request, output *DescribeBgpPeerAttributesOutput) {
	op := &request.Operation{
		Name:       opDescribeBgpPeerAttributes,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeBgpPeerAttributesInput{}
	}

	output = &DescribeBgpPeerAttributesOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeBgpPeerAttributes API operation for DIRECTCONNECT.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for DIRECTCONNECT's
// API operation DescribeBgpPeerAttributes for usage and error information.
func (c *DIRECTCONNECT) DescribeBgpPeerAttributes(input *DescribeBgpPeerAttributesInput) (*DescribeBgpPeerAttributesOutput, error) {
	req, out := c.DescribeBgpPeerAttributesRequest(input)
	return out, req.Send()
}

// DescribeBgpPeerAttributesWithContext is the same as DescribeBgpPeerAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeBgpPeerAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DIRECTCONNECT) DescribeBgpPeerAttributesWithContext(ctx volcengine.Context, input *DescribeBgpPeerAttributesInput, opts ...request.Option) (*DescribeBgpPeerAttributesOutput, error) {
	req, out := c.DescribeBgpPeerAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeBgpPeerAttributesInput struct {
	_ struct{} `type:"structure"`

	// BgpPeerId is a required field
	BgpPeerId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeBgpPeerAttributesInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBgpPeerAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeBgpPeerAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeBgpPeerAttributesInput"}
	if s.BgpPeerId == nil {
		invalidParams.Add(request.NewErrParamRequired("BgpPeerId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBgpPeerId sets the BgpPeerId field's value.
func (s *DescribeBgpPeerAttributesInput) SetBgpPeerId(v string) *DescribeBgpPeerAttributesInput {
	s.BgpPeerId = &v
	return s
}

type DescribeBgpPeerAttributesOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	AccountId *string `type:"string"`

	AuthKey *string `type:"string"`

	BgpPeerId *string `type:"string"`

	BgpPeerName *string `type:"string"`

	CreationTime *string `type:"string"`

	Description *string `type:"string"`

	LocalAsn *int64 `type:"integer"`

	RemoteAsn *int64 `type:"integer"`

	RequestId *string `type:"string"`

	SessionStatus *string `type:"string"`

	Status *string `type:"string"`

	UpdateTime *string `type:"string"`

	VirtualInterfaceId *string `type:"string"`
}

// String returns the string representation
func (s DescribeBgpPeerAttributesOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeBgpPeerAttributesOutput) GoString() string {
	return s.String()
}

// SetAccountId sets the AccountId field's value.
func (s *DescribeBgpPeerAttributesOutput) SetAccountId(v string) *DescribeBgpPeerAttributesOutput {
	s.AccountId = &v
	return s
}

// SetAuthKey sets the AuthKey field's value.
func (s *DescribeBgpPeerAttributesOutput) SetAuthKey(v string) *DescribeBgpPeerAttributesOutput {
	s.AuthKey = &v
	return s
}

// SetBgpPeerId sets the BgpPeerId field's value.
func (s *DescribeBgpPeerAttributesOutput) SetBgpPeerId(v string) *DescribeBgpPeerAttributesOutput {
	s.BgpPeerId = &v
	return s
}

// SetBgpPeerName sets the BgpPeerName field's value.
func (s *DescribeBgpPeerAttributesOutput) SetBgpPeerName(v string) *DescribeBgpPeerAttributesOutput {
	s.BgpPeerName = &v
	return s
}

// SetCreationTime sets the CreationTime field's value.
func (s *DescribeBgpPeerAttributesOutput) SetCreationTime(v string) *DescribeBgpPeerAttributesOutput {
	s.CreationTime = &v
	return s
}

// SetDescription sets the Description field's value.
func (s *DescribeBgpPeerAttributesOutput) SetDescription(v string) *DescribeBgpPeerAttributesOutput {
	s.Description = &v
	return s
}

// SetLocalAsn sets the LocalAsn field's value.
func (s *DescribeBgpPeerAttributesOutput) SetLocalAsn(v int64) *DescribeBgpPeerAttributesOutput {
	s.LocalAsn = &v
	return s
}

// SetRemoteAsn sets the RemoteAsn field's value.
func (s *DescribeBgpPeerAttributesOutput) SetRemoteAsn(v int64) *DescribeBgpPeerAttributesOutput {
	s.RemoteAsn = &v
	return s
}

// SetRequestId sets the RequestId field's value.
func (s *DescribeBgpPeerAttributesOutput) SetRequestId(v string) *DescribeBgpPeerAttributesOutput {
	s.RequestId = &v
	return s
}

// SetSessionStatus sets the SessionStatus field's value.
func (s *DescribeBgpPeerAttributesOutput) SetSessionStatus(v string) *DescribeBgpPeerAttributesOutput {
	s.SessionStatus = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeBgpPeerAttributesOutput) SetStatus(v string) *DescribeBgpPeerAttributesOutput {
	s.Status = &v
	return s
}

// SetUpdateTime sets the UpdateTime field's value.
func (s *DescribeBgpPeerAttributesOutput) SetUpdateTime(v string) *DescribeBgpPeerAttributesOutput {
	s.UpdateTime = &v
	return s
}

// SetVirtualInterfaceId sets the VirtualInterfaceId field's value.
func (s *DescribeBgpPeerAttributesOutput) SetVirtualInterfaceId(v string) *DescribeBgpPeerAttributesOutput {
	s.VirtualInterfaceId = &v
	return s
}
